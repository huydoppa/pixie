package certs

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"path"
	"time"

	log "github.com/sirupsen/logrus"

	"pixielabs.ai/pixielabs/src/utils/pl_admin/cmd/k8s"
)

func generateCA(certPath string, bitsize int) (*x509.Certificate, crypto.PrivateKey) {
	ca := &x509.Certificate{
		SerialNumber: big.NewInt(1653),
		Subject: pkix.Name{
			Organization: []string{"Pixie Labs Inc."},
			Country:      []string{"US"},
			Province:     []string{"California"},
			Locality:     []string{"San Francisco"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	caKey, err := rsa.GenerateKey(rand.Reader, bitsize)
	if err != nil {
		log.WithError(err).Fatal("Could not generate key for certificate")
	}

	signCertificate(certPath, "ca", ca, ca, caKey, caKey)

	return ca, caKey
}

func loadCA(caCert string, caKey string) (*x509.Certificate, crypto.PrivateKey) {
	caPair, err := tls.LoadX509KeyPair(caCert, caKey)
	if err != nil {
		log.WithError(err).Fatal("Could not load CA.")
	}
	ca, err := x509.ParseCertificate(caPair.Certificate[0])
	if err != nil {
		log.WithError(err).Fatal("Could not parse CA cert.")
	}

	return ca, caPair.PrivateKey
}

func generateCertificate(certPath string, certName string, caCert *x509.Certificate, caKey crypto.PrivateKey, bitsize int) {
	// Prepare certificate.
	cert := &x509.Certificate{
		SerialNumber: big.NewInt(1658),
		Subject: pkix.Name{
			Organization: []string{"Pixie Labs Inc."},
			Country:      []string{"US"},
			Province:     []string{"California"},
			Locality:     []string{"San Francisco"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0),
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
		DNSNames:              []string{"*.local", "*.plc.svc.cluster.local", "*.pl.svc.cluster.local", "localhost", "pl-nats", "pl-etcd", "*.pl-etcd.pl.svc", "*.pl-etcd.pl.svc.cluster.local"},
	}
	privateKey, err := rsa.GenerateKey(rand.Reader, bitsize)
	if err != nil {
		log.WithError(err).Fatal("Could not generate key for certificate")
	}

	signCertificate(certPath, certName, cert, caCert, caKey, privateKey)
}

func signCertificate(certPath string, certName string, cert *x509.Certificate, ca *x509.Certificate, caKey crypto.PrivateKey, privateKey *rsa.PrivateKey) {
	// Self-sign certificate.
	certB, err := x509.CreateCertificate(rand.Reader, cert, ca, &privateKey.PublicKey, caKey)

	certOut, err := os.Create(path.Join(certPath, fmt.Sprintf("%s.crt", certName)))
	if err != nil {
		log.WithError(err).Fatal(fmt.Sprintf("Could not create %s.crt", certName))
	}
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: certB})
	certOut.Close()
	log.Info(fmt.Sprintf("Created %s.crt", certName))

	// Generate key.
	keyOut, err := os.OpenFile(path.Join(certPath, fmt.Sprintf("%s.key", certName)), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.WithError(err).Fatal(fmt.Sprintf("Could not create %s.key", certName))
	}
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)})
	keyOut.Close()
	log.Info(fmt.Sprintf("Created %s.key", certName))
}

func generateCerts(certPath string, caCertPath string, caKeyPath string, bitsize int) {
	var ca *x509.Certificate
	var caKey crypto.PrivateKey

	if caCertPath == "" {
		log.Info("Generating new CA.")
		ca, caKey = generateCA(certPath, bitsize)
	} else {
		log.Info("Using existing CA.")
		ca, caKey = loadCA(caCertPath, caKeyPath)
	}

	// Generate server certificate.
	generateCertificate(certPath, "server", ca, caKey, bitsize)

	// Generate client certificate.
	generateCertificate(certPath, "client", ca, caKey, bitsize)
}

// InstallCerts generates the necessary certs and installs them in kubernetes.
func InstallCerts(certPath string, caCertPath string, caKeyPath string, namespace string, bitsize int) {
	var err error

	deleteCerts := false
	if certPath == "" {
		certPath, err = ioutil.TempDir("", "certs")
		if err != nil {
			log.WithError(err).Fatal("Could not create temp directory")
		}
		deleteCerts = true
	}

	// Delete generated certs.
	defer func() {
		if deleteCerts {
			log.Info("Deleting generated certs")
			os.RemoveAll(certPath)
		}
	}()

	generateCerts(certPath, caCertPath, caKeyPath, bitsize)

	serverKey := path.Join(certPath, "server.key")
	serverCert := path.Join(certPath, "server.crt")
	caCert := path.Join(certPath, "ca.crt")
	if caCertPath != "" {
		caCert = caCertPath
	}
	clientKey := path.Join(certPath, "client.key")
	clientCert := path.Join(certPath, "client.crt")

	// Authenticate with k8s cluster.
	config := k8s.GetConfig()
	clientset := k8s.GetClientset(config)

	// Delete secrets in k8s.
	k8s.DeleteSecret(clientset, namespace, "proxy-tls-certs")
	k8s.DeleteSecret(clientset, namespace, "service-tls-certs")
	k8s.DeleteSecret(clientset, namespace, "etcd-peer-tls-certs")
	k8s.DeleteSecret(clientset, namespace, "etcd-client-tls-certs")
	k8s.DeleteSecret(clientset, namespace, "etcd-server-tls-certs")

	// Create secrets in k8s.
	k8s.CreateTLSSecret(clientset, namespace, "proxy-tls-certs", serverKey, serverCert)

	k8s.CreateGenericSecret(clientset, namespace, "service-tls-certs", map[string]string{
		"server.key": serverKey,
		"server.crt": serverCert,
		"ca.crt":     caCert,
		"client.key": clientKey,
		"client.crt": clientCert,
	})

	k8s.CreateGenericSecret(clientset, namespace, "etcd-peer-tls-certs", map[string]string{
		"peer.key":    serverKey,
		"peer.crt":    serverCert,
		"peer-ca.crt": caCert,
	})

	k8s.CreateGenericSecret(clientset, namespace, "etcd-client-tls-certs", map[string]string{
		"etcd-client.key":    clientKey,
		"etcd-client.crt":    clientCert,
		"etcd-client-ca.crt": caCert,
	})

	k8s.CreateGenericSecret(clientset, namespace, "etcd-server-tls-certs", map[string]string{
		"server.key":    serverKey,
		"server.crt":    serverCert,
		"server-ca.crt": caCert,
	})
}
