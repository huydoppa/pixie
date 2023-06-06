// Code generated for package noauth by go-bindata DO NOT EDIT. (@generated)
// sources:
// 01_base_schema.graphql
// 02_unauth_schema.graphql
// 03_auth_schema.graphql
package noauth

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var __01_base_schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\xb1\x4e\xc4\x30\x10\x44\x7b\x7f\xc5\xa0\x14\x54\x5c\x2a\x10\x4a\x49\x4f\x81\xe0\x07\x1c\x7b\x38\x47\x72\xbc\x3e\xef\x46\x47\x84\xf8\x77\x94\xcb\x5d\x77\xd5\x6c\x31\xf3\xb4\x4f\x43\xe2\xec\xf1\xeb\x80\xd3\xc2\xb6\x0e\xf8\xd8\xc2\x01\xf3\x62\xde\x26\x29\x03\xde\xaf\x97\xfb\x73\xae\xc3\x57\x22\xb4\x32\x20\x0a\xb5\x3c\x1a\x7c\xce\x72\x06\xe7\x6a\x2b\x6c\xad\xd4\x83\xeb\xf0\x29\x38\x13\xa1\xd1\x1b\x51\x7d\x0e\x4c\x92\x23\x9b\x22\xb1\x11\xbe\xc4\xeb\xce\x12\x95\xfb\x0e\x26\x18\xe9\x3a\xf0\xc7\x58\x22\x23\xc6\x15\x62\x89\x0d\xdf\x53\xde\xb9\xc9\xac\xea\xd0\xf7\xc7\xc9\xd2\x32\x1e\x82\xcc\xfd\xb1\xf9\x9a\x4e\xf9\x96\x4f\xdb\x73\xfd\xa4\xba\x50\xfb\xe7\x97\x57\xe7\x36\xf8\xae\x75\xf1\x2c\x22\x75\xc0\x9b\x48\xa6\x2f\x0f\x9b\xd4\xa5\x70\xb3\xbc\xdf\xf9\x0f\x00\x00\xff\xff\x6f\xc4\xb8\xef\x28\x01\x00\x00")

func _01_base_schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__01_base_schemaGraphql,
		"01_base_schema.graphql",
	)
}

func _01_base_schemaGraphql() (*asset, error) {
	bytes, err := _01_base_schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "01_base_schema.graphql", size: 296, mode: os.FileMode(436), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __02_unauth_schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8d\xb1\x0a\xc2\x40\x0c\x86\xf7\x7b\x8a\xdc\x56\x5f\xe1\x36\x1d\x84\x0e\x0a\xa2\x9b\x38\x84\x9a\xd6\x60\x2f\x29\x77\xb1\x58\xc4\x77\x17\x0b\xa5\x15\xb7\x9f\xe4\xfb\xbf\x9f\x9e\x46\x72\x05\x1b\x3a\x82\xc3\x83\xd2\x00\x2f\x07\x80\xc9\xb8\xc6\xca\x72\x31\xa5\x3d\x46\x0a\x70\xb4\xc4\xd2\xf8\x55\x80\xf5\x44\x94\x52\xab\x77\x00\x3d\x25\xae\x87\x52\x7a\x36\x3a\xe9\x9d\xa4\xe0\x39\x2f\x9b\x1b\xd5\x96\x50\xbc\x7b\x3b\x37\xce\xfe\xa8\xc6\x79\x36\x8a\x39\xc0\x79\xfa\xf8\xcb\x3f\x3d\x82\x3d\xa5\xcc\x3a\xeb\x1d\x40\x75\x43\x69\xa8\xd5\x66\x79\x34\x8e\x94\x0d\x63\xb7\xcb\x01\xb6\xad\xa2\x7d\x85\x9f\x00\x00\x00\xff\xff\xef\x77\x02\x34\xfc\x00\x00\x00")

func _02_unauth_schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__02_unauth_schemaGraphql,
		"02_unauth_schema.graphql",
	)
}

func _02_unauth_schemaGraphql() (*asset, error) {
	bytes, err := _02_unauth_schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "02_unauth_schema.graphql", size: 252, mode: os.FileMode(436), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __03_auth_schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x5a\xdd\x8f\xe3\xb6\x11\x7f\xf7\x5f\x31\x97\x7b\xc8\x2e\xb0\x3d\x04\x45\x13\x14\xfb\x54\xc5\xd6\xe5\xd4\xdd\xf5\xba\xb6\x37\x69\x10\x1c\x0e\xb4\x34\xb6\x88\x95\x48\x85\xa4\xbc\xeb\x16\xf9\xdf\x8b\x21\x29\x89\x94\xb5\xf7\xd5\xa2\x6f\x16\x3f\x66\x7e\x33\x1c\xce\x17\x8d\xcf\x06\x45\x01\xe6\xd4\x20\xfc\xa3\x45\x75\x82\x7f\xcf\x00\x5a\x8d\xea\x1a\x1e\x34\xaa\x4c\xec\xe5\xab\x19\x80\x54\x87\x6b\xb8\x57\x87\xee\x9b\x56\x6c\xd0\x18\x2e\x0e\xda\xad\xec\xbe\xba\xd9\xc4\x18\xc5\x77\xad\x41\x3f\x3f\x7c\x7b\x7a\x34\xa8\xaf\xe1\xb7\x9e\xcd\x7b\x9a\xc8\xab\x56\x1b\x54\x17\xbc\xb8\x86\x6c\xf1\xea\xf2\x1a\xe6\x6e\xa4\xe3\xec\x17\xfc\x78\x5a\xb2\x1a\x2f\x04\xab\xf1\x1a\x36\x46\x71\x71\x78\x79\x31\xb1\x09\x67\x3c\x27\x9e\x28\xc3\xf7\x2c\x37\x17\xcc\xff\xd8\x9e\x1a\xbc\x86\x24\xf8\xb2\x44\x6f\xb3\x6e\x88\x36\xb2\xd6\xc8\x5c\xd6\x4d\x85\x06\x2f\xb8\x68\x5a\xd3\x21\xb8\x82\xbc\x55\x5a\xaa\x95\xd4\xd7\x90\x09\x73\x05\x2c\x37\x5c\x8a\x6b\x48\x82\x3d\x89\x1d\x23\xe2\x57\x1d\xc0\x87\x6c\xd1\xd1\xb8\x8c\x17\xaf\x51\xb7\xd5\x19\xdb\xb7\x1c\xab\x62\xcc\x7b\x4f\x83\x5e\x82\x60\x6d\x2a\x0c\x37\xa7\x1b\x2e\x8a\xab\x19\x00\x80\xc2\xdf\x5b\xae\xb0\x48\xd4\x81\x16\x93\x6e\xa6\x97\xbf\xff\x0c\x78\x16\x48\x87\x71\x06\xf0\x1a\x36\xb9\xe2\x8d\xa9\x0f\x0a\x50\x14\x8d\xe4\xc2\xe8\x2b\x50\xb8\x47\x05\x46\x42\x21\x73\x0d\x5c\x40\x5e\xc9\xb6\x60\x0d\x7f\xd3\x28\x69\xe4\x0c\xa0\xe2\x47\xfc\x99\xe3\x13\xc1\xb9\xf5\xbf\xef\xd0\xb0\x82\x19\xe6\xce\xab\x5b\x31\x97\xc2\xa0\x30\x3a\x30\x91\xdb\xd1\x14\x2d\xd7\x16\x07\x91\x73\x88\x62\x62\x6e\x76\x82\xd4\x26\x9a\xf0\x32\x2d\xb0\xa9\xe4\x09\x1e\xf1\xa4\x67\x00\x85\xfd\xaa\x51\x98\x1b\x3c\x11\x83\x45\x38\x10\xf3\x89\xd6\x06\x6c\xa2\x2d\x9e\x4b\xb2\xca\x3a\x16\xac\xe1\x9e\x76\xb2\xca\xce\x88\xba\xd9\x80\x9a\x5b\xe4\xc9\xac\xaa\xf6\xc0\xc5\x0c\xa0\xb1\x3f\xf4\xc5\x23\x17\xc5\xb5\x1f\xa6\x73\xbd\xbc\x86\xdf\xdc\x97\x23\xa7\x90\x64\xe5\x52\xb8\x41\xba\x21\x96\xb6\xbf\x55\x57\x9e\xd0\xcf\xa8\xb4\xb5\xe5\xe1\xb6\x0d\x1b\x5e\x59\xd6\xdb\x92\x6b\x78\xe2\x55\x05\x3b\x04\x85\x4d\xc5\x72\x2c\x60\x77\x1a\xb3\x98\x4b\xb1\xe7\x07\x90\x22\x47\x60\x55\x05\xb9\x14\xba\xad\x51\x69\x28\xd9\x11\x61\x87\x28\xa0\x6d\x0a\x66\xb0\x78\xe3\x9c\xc5\x7a\x8a\x40\x88\x72\x10\xca\xcd\x4d\x8a\x36\xb9\x6d\x92\x74\xb4\x79\xd3\x9b\xd2\x3a\x1e\x1a\x31\x71\x83\x23\xf2\x0b\x34\x8c\x57\x58\x8c\xb7\xce\xfe\x98\xcd\x42\xef\x7b\xd7\x1a\x46\xd3\xd6\x01\xcf\x15\x32\x83\xde\x65\x45\x5e\x0d\xfe\x56\x60\xa3\x30\x27\xdd\x5c\x28\x64\x9a\x4e\xe4\x1b\xbf\x40\x03\x53\x08\x42\x3e\x41\x6e\x09\x14\x70\xe4\x0c\x9a\x67\x6f\x86\xdf\x5c\xf6\xa4\x23\xfb\x3b\x33\x47\x80\x05\xd2\xed\x5e\xbc\x60\xbd\x3f\x4a\x59\x21\x13\xaf\x7a\x72\xce\x00\x07\x43\xec\x08\xb8\xef\xe9\x9d\x0f\xf6\x80\xc3\xd0\x71\xa1\xfb\x88\x92\x16\xdc\xb0\x5d\x15\x4d\xd3\xfe\x71\xa4\xd9\xa0\x89\x83\xcb\x05\x0b\xe2\x4e\x48\x25\x88\x3f\x97\x53\x11\xc9\x01\x7e\xb0\x91\x2f\x40\x99\x89\x23\x77\xc3\x17\x58\x33\x5e\x05\xd7\x62\xcf\x95\x36\xcb\x30\x00\x5d\x41\xc5\x46\x43\x97\x5d\x1c\x25\x32\xb1\xd8\x2b\x54\x35\xd7\x74\xa7\xf4\x05\x45\x4c\x72\xb2\xd9\xe2\xd5\x95\x0d\x9f\xc1\x64\x2c\x47\x30\x31\x10\x77\x17\xd0\x9d\xc5\xbd\x3a\x5c\x48\x75\x18\xa3\xc8\x16\x03\xf7\x7b\x75\xe8\x75\x2e\xd5\xa1\x67\x2c\x87\xf1\x81\x69\xb0\x98\xe8\x04\x49\x80\xe3\xe7\x44\xdb\xca\x47\x14\x01\xb1\xcb\x9e\xf7\x0c\x60\x8d\x47\xf9\x88\x49\x55\x05\x6b\x75\xbc\x38\x50\xf9\x1a\x6b\x79\xb4\xb2\xbe\x55\xb2\x26\x71\x02\xed\x84\x4b\x63\x77\xe7\x44\xfb\xa4\xa3\xb8\x02\x14\x24\x56\xd1\x13\xea\x47\x46\x1e\xee\x8a\xfc\xd2\x9e\x87\xba\x08\x89\xea\x49\x7b\x5e\x4f\x38\x03\xab\x5b\x17\x71\x06\x52\xa3\x85\x53\xb7\x6a\x4c\xeb\xd3\x24\xdc\x21\x3b\x53\x7e\x01\x49\xc8\xe8\x8f\xd9\xcc\xfa\x9e\xce\x88\xac\xef\xf1\xeb\x66\x00\x51\x72\x35\x03\x88\x2f\x00\x45\x18\x9e\x9b\x56\x45\x6b\xc6\x96\xe7\x86\x86\xfc\x81\x06\xb8\x4e\x9a\x46\xc9\x63\x70\x06\x03\x96\x6c\x91\xae\x98\x29\x2d\x94\x6c\x91\x8e\x89\x35\xcc\x94\xc3\x77\xb7\xc9\x1b\xe5\x27\xf0\x17\xb2\x66\x5c\x8c\x29\xba\xc3\x77\x88\x58\xa5\xa3\x73\xe0\x05\x12\x18\x72\xfd\x1e\x17\xb9\xfc\x50\x6d\xdd\xd5\xb0\xac\x99\x60\xd5\xc9\xf0\x5c\xdf\x37\x46\x52\x6a\x16\x91\x72\xb0\xc2\xcd\x83\x07\xb2\xdb\x8d\x6c\xd5\x06\x51\xbc\xb4\xcf\xe6\x7b\x2f\x38\xb5\x69\x02\xd3\xbb\x3e\x0b\x73\x0f\x34\xce\x40\x46\x2a\xf6\x71\x26\x31\x77\xfa\x1a\xde\x56\x92\x19\x97\xf5\xe8\xfc\xfc\x90\x1c\xa1\x11\x81\x47\x0a\x1b\xc3\x61\x7c\x09\xbd\xc9\xb4\xeb\xbf\xc0\x17\xd1\xfb\x9f\xc0\x44\xd1\xd6\x13\xb9\xf8\xc6\x30\x83\x96\x41\x92\x6e\x3e\x3c\x2c\x6f\x96\xf7\xbf\x2c\xfd\xd7\x2a\x5d\x2e\xb2\xe5\x4f\xfe\x6b\xfd\xb0\x5c\x0e\x5f\x6f\x93\xec\x36\x5d\xf8\x8f\x6d\xba\xbe\xcb\x96\xc9\x36\x5d\x4c\x72\x1a\x8a\x0c\xc7\x28\xd9\x06\x8c\x5e\x43\x22\x00\x0b\x6e\x7c\x7d\x02\x32\xa7\xc2\x05\xf8\x1e\x98\x8d\x3e\x50\x32\x0d\xb5\x2c\xf8\x9e\x63\x01\xa6\x44\x70\x56\x64\xf0\xd9\x50\x3e\xc7\x85\x46\x45\x36\x04\x52\x41\x41\xee\x86\x7e\xe7\x25\x53\x2c\xa7\x3c\xe4\xcd\x90\x0e\x72\x91\x57\x6d\x81\x9a\xb2\x1c\xbb\x41\x58\x7a\x8f\x78\xda\x49\xa6\x0a\x60\xa2\x80\x86\x69\x47\x40\xd6\x35\x13\x85\xdd\x4e\x88\xd3\x45\xb6\x75\x70\x41\x63\x85\xf9\x80\x57\x54\xa7\x69\xd0\x79\x29\x35\x0a\x60\x22\xaa\x97\x40\xb7\x87\x03\x6a\xda\xfb\xa6\x83\x55\x70\x4a\xa2\x34\xd8\xf2\xe3\xb5\x05\x15\x6d\x31\x25\x33\xc0\x0d\xe8\x52\xb6\x55\x01\x14\x93\xec\x22\x62\xf5\xad\xf6\x95\x1e\xd5\x34\x34\x28\x48\x31\x8c\x7c\x48\xa3\x38\x9d\xae\x61\xbb\x4e\x8a\x4d\x7a\x9b\xce\xb7\x1f\xb1\x07\x4a\xca\xbd\x39\xdc\x44\xe6\x70\xf3\x61\x75\xbf\xf0\xbf\x36\x3f\xcf\xbb\x5f\xf3\x75\xb6\xda\xfa\x8f\x65\x72\x97\x6e\x56\xc9\x3c\xed\xbe\xef\x17\xe9\x70\xe3\x02\x56\x9b\x5e\x03\x96\x95\x2b\x0a\xa6\xb1\x8c\x5c\xa7\xb7\x6c\x0a\x22\x41\x74\x9c\x01\xd4\xcc\xe4\x25\x16\x99\x28\xf0\xd9\xd6\x91\x99\x30\xef\xa9\xb8\x22\xfb\x9e\x22\x6e\x0d\xbf\x47\xb7\x65\xbb\x11\x28\x32\x19\x32\xb5\x02\x9f\x41\xee\xad\x62\x0d\xdb\xb9\x93\x30\x25\xea\xf0\x1c\x5d\xae\xbb\x97\x8a\xd4\x6c\xd8\xce\xa2\xb0\x55\xb7\x25\xf4\x4b\x89\xa6\x44\xe5\xed\x86\x8c\x8b\x05\x9b\x69\x1f\x18\xb2\x03\xa2\xef\x18\xda\xba\xa5\x66\x8f\xee\x94\xbd\x29\x02\x3e\x63\xde\x5a\xcf\x49\x7c\x86\xaf\x64\x6f\xc8\x91\x12\xf1\xc1\x65\x42\x88\x6f\x54\x57\x0f\xa2\xbe\x9f\x3c\x1f\x57\x44\x07\x6a\xd8\x4b\x55\x33\x43\x49\xbc\xbb\x7b\x04\xb6\xbf\x88\xda\x67\x28\x4f\x25\xcf\x4b\x6b\xf8\xb6\x68\x6a\x98\xd2\xae\xe2\x3a\x37\x67\xd9\xdb\xbc\xb3\x77\xb6\xdb\x18\xd9\x40\x23\x35\xb7\x78\x49\xbe\x9e\x67\x16\xb6\x16\x22\x85\x8e\x31\x10\x2e\x06\x47\x56\xf1\xe2\x2a\xd0\x4f\xa7\xc0\x37\x36\xde\xa7\xfd\x78\xa8\xac\xd7\x90\x54\x55\x74\xa4\x74\x2c\xc8\xf2\x32\x38\x7d\x02\xa9\xfd\x19\x6f\x22\xed\x46\xf6\x33\xad\xd4\xa0\x3d\x11\x68\xf6\x05\xcf\xa0\xbd\x55\x74\xf2\x51\x42\xc0\x0b\x2c\x3e\xf7\x58\x5f\x45\x7a\x92\x0a\x84\xb4\x66\x4b\x05\x62\xab\x04\x16\xa0\x2c\x12\x67\xb9\x0d\x53\x86\xb3\x0a\x2e\x8c\x6a\xf1\x92\x96\xf7\x90\x2e\xf6\xac\xd2\x48\xc5\x5a\xc9\x74\x52\x14\xf6\x7c\x58\x75\x67\xaf\x9b\x9e\xc8\x99\xe6\x52\x18\xc6\x05\x2a\xba\x60\xad\x8b\xeb\xe3\xe4\x67\x3a\x64\xf9\xab\x3a\x2c\xab\x51\x6b\x76\x88\x86\xba\x2a\x33\x1c\xd1\x86\x29\x33\x97\xad\x30\xf6\xca\x0d\x50\x6e\xfe\xaa\xd3\x23\x0a\xa7\xee\x09\x62\xb6\x68\xda\xf2\x1a\x23\x18\x54\x36\x8d\x06\x3b\x82\x2b\x59\x7c\x95\x54\xad\xfe\x62\xb1\xf2\x4e\x8d\xb6\x5f\x18\xeb\xd4\x95\xfa\x48\xa2\xd1\x6c\x27\x66\xd7\x01\x98\xd2\x87\xf5\xf6\xbe\x32\x0f\x44\x70\x36\x58\xe0\x9e\x91\x55\xda\x03\xa0\x18\x26\xa4\x29\xfd\x75\x7a\x14\xf2\x49\x90\xc9\xcf\x37\x51\xd0\xa6\x7d\x7e\xbd\x86\x12\x59\x65\xca\x13\x6d\x2d\x91\x29\xb3\x43\xe6\x2d\x4b\x61\x8e\xfc\x88\x05\x85\x5a\x85\x87\xb6\x62\x0a\xb8\x30\xa8\x28\xbd\xb5\xf1\xd6\x94\xce\x07\xf8\xf6\x1e\x91\x53\xa8\x1b\x29\x0a\x42\x60\xa4\xed\x11\xa2\x36\xda\x83\x78\x97\x26\xb7\xdb\x77\xbf\x9e\x83\x68\x45\x00\xc3\xba\xcd\x81\x62\x2e\x85\xc0\x9c\xfc\x97\x91\xb0\xe2\xcf\x1c\x61\x5e\xc9\xd6\x45\x7c\xae\xfd\xf5\xea\xdc\xcb\x20\xc3\x15\xec\xac\xb7\x13\xdf\x1a\xf8\xbd\x45\x75\xb2\xee\x84\xae\xa6\x96\x35\xfa\x63\xf3\x51\x5c\xa1\xc6\x7a\x57\xa1\x86\x77\xdb\xed\xea\x5b\x0d\xdf\x7f\xf7\x9d\x3f\xfd\x5e\x7f\xd3\xe0\xad\xb7\x3f\x48\xdb\x93\xe4\x7a\xc0\xea\xe5\xf8\x69\xbd\x9a\x77\x12\x50\xbc\xd8\x29\x64\x8f\xfa\x8d\x25\x50\xca\x06\x9d\x37\x66\xa6\x4f\x1d\x3a\xc1\x2d\xdd\x9c\x80\xee\x58\xfe\x48\x89\x0a\x17\x68\x45\xa6\xcb\x5f\x93\x6f\x01\x8f\xc8\x21\xf1\x38\x17\xd9\x66\x7e\xbf\x5c\xa6\xf3\xad\xcd\xf0\xc6\x7a\xa6\xda\x92\xce\xe6\xa9\x44\x31\x56\x34\x77\x23\x8d\x92\x39\x6a\x4d\xae\xb3\x5b\xde\xe9\x60\xb5\x48\xb6\x2e\x8d\x74\x74\x8f\xfc\x5f\xbc\xcb\x97\x3a\xc9\x9d\xda\x69\x88\xdc\x96\xa6\x2b\xcc\xc4\x09\xa4\x75\x66\xfb\x56\xb9\x68\xea\xcc\xd8\xf5\xe6\x34\xb0\x9d\x6c\x9d\x0a\x9e\xbc\xd7\xe3\x26\xb4\x4d\xa9\xc6\x50\xce\x65\xf4\x58\x9e\x98\x06\xa3\x4e\xde\xfe\x1c\x03\x07\x69\x6f\xdb\x67\x9d\xd5\x08\xf9\x44\x02\x33\xd8\xb1\x22\x52\xa0\x15\x32\x1d\x72\xe4\x91\x06\x0b\x3c\x28\x56\x0c\x07\x1c\xe8\xaf\xe2\x8f\x58\x9d\x88\xed\x0e\x03\x8b\x23\xde\x35\x3f\x94\x86\x86\x6d\xcb\xc5\x9b\x2a\x95\x19\xdd\xa9\xa5\x3f\xad\x93\x85\x4b\xc1\x9d\x27\x0e\x3a\x75\x71\x05\xd1\xf9\xa4\xc8\x21\x74\xbe\xef\x5d\x67\xfd\x91\x1b\x73\xaa\x19\xf7\x5d\xa9\xa2\x6e\x50\x31\x23\xa7\xa6\xbc\x5c\x2f\xcf\x9c\x15\xd4\x0a\x8d\x39\xcd\xa7\x27\xcf\xfb\xff\x9d\x9f\x54\xb2\x5a\x55\x4c\x60\xef\x9e\x6d\x02\xd8\x7f\x39\xbf\xd8\xbb\x87\x05\x33\xec\xd3\xcb\x45\x5b\x2f\x65\x81\xda\xbb\x50\x3b\x90\x09\x6d\x54\x4b\x45\x19\x16\xf1\xa4\x53\xe8\xdd\xb9\x63\x6f\x14\x1e\xb9\x6c\xf5\x66\x4a\xe3\x67\xf3\x51\xd8\x19\x75\x44\x8e\xdc\xd7\x69\x67\x9d\x0f\x6e\xe7\x6e\xb9\x78\x8c\x0a\xbe\xd7\xb0\xfe\xc4\x4b\x87\xa5\x3e\x7e\xe0\xf8\x54\xdf\x62\x5c\x57\x7e\x21\x9b\xee\x35\xc3\xc7\x64\xc7\xf3\xfa\x0c\x85\xd5\xdd\x73\xd5\xad\x0e\x11\x1c\xb9\xfe\xfb\xe6\x7e\xf9\x35\x20\xe2\xd7\x97\x2f\x92\xd4\xe6\x3f\x1d\xca\x38\xed\xf9\x22\xe6\x2f\xc8\x3f\x7a\x17\xf2\x86\x1d\x8b\xde\x17\x6d\xc1\x93\xa0\x25\x03\x10\x55\xd4\xf6\xf3\x36\x5b\x3e\xfc\xf3\x43\x72\xb7\xf8\xe1\x2f\xdd\xd0\x22\x59\xff\x92\x2d\xe3\xb1\xf9\xfd\x72\x9b\x64\xcb\x74\xfd\x61\x93\x6e\x3f\xfc\x9a\xdc\xdd\x6e\xa6\xa7\x26\xe8\xc5\x0b\xb6\xe9\xdd\xea\x96\xbc\x9e\x23\xd2\x3b\xa1\xe1\xbd\xd2\x3d\xe7\xaa\xc8\x76\x75\xc9\xfe\xfc\xfd\x0f\x91\x8c\xe7\x3d\xa2\xa0\xcd\xec\xce\xec\xbc\x67\x77\xbe\x31\x68\x15\xbb\x6b\xf3\x42\x6b\xcd\x9d\xa0\x6b\xa6\xfe\x49\x61\x65\x9f\x2a\x08\xba\x7e\xd3\x65\x7d\x76\x6e\x32\xe5\x0b\x7a\xb9\xd3\x95\xa9\x75\xaa\xf2\x20\x83\xf2\x85\x38\x68\x33\xe1\x14\x75\xdb\x34\x52\x19\xdd\x37\x4b\xa3\xde\x5b\xff\xb2\x93\x8e\x5a\xc6\x43\xdf\x70\xdc\x34\xee\x2d\x66\x78\x6f\xb3\x52\xac\xc2\xca\x7e\x75\xf3\x61\x9d\x6e\xd3\xe5\x36\xbb\x5f\x0e\x89\x6e\xf8\x34\x36\x25\xf8\x91\x55\x2d\x9e\xb7\xad\x86\x47\x38\xbb\xab\xef\x59\x47\xcf\x61\x9b\xbc\xc4\xba\x7b\x3e\xac\x2a\xf9\x34\x6f\xb5\x91\x75\xfa\x4c\xd2\x3f\xac\x6f\x23\xc9\xec\x82\x4c\x68\xcc\x5b\x85\xdb\xdb\x4d\x34\xe9\x13\xd7\x60\xe7\x34\xa0\x90\xef\xa4\x38\x93\x27\x77\x66\x52\x5f\xa1\x95\x97\x09\xe8\x91\x86\xa6\xd6\xf8\x3f\x07\x8c\xd5\xd3\x9b\x12\x9f\xd0\x4b\x2f\xfa\xe4\xeb\xc3\x47\x8e\xe5\x6b\x99\xbd\x06\x8a\xa8\x83\x7d\xc6\x97\x67\xd4\xf9\xff\x0c\x9f\x3b\x71\x83\xf6\xb6\x06\x10\xf9\x69\xd3\x47\xdc\xf1\xc3\xc9\xe8\x1f\x16\xd9\xc2\x89\xe3\x9e\x8b\xc7\x2d\xff\x15\xe5\xea\x66\xa2\x78\x7d\xe1\x81\xf4\xff\x8d\xfa\xdc\xf7\x4f\x4b\xf2\x91\xc3\x9a\x94\x31\xb6\xc6\x29\x21\x3f\xab\xd7\x36\x12\x6c\x42\xae\x73\xb1\x26\xa4\x9a\x10\xea\x23\x32\xfd\x31\xfb\x4f\x00\x00\x00\xff\xff\xe4\x54\x23\xb9\x20\x24\x00\x00")

func _03_auth_schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		__03_auth_schemaGraphql,
		"03_auth_schema.graphql",
	)
}

func _03_auth_schemaGraphql() (*asset, error) {
	bytes, err := _03_auth_schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "03_auth_schema.graphql", size: 9248, mode: os.FileMode(436), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"01_base_schema.graphql":   _01_base_schemaGraphql,
	"02_unauth_schema.graphql": _02_unauth_schemaGraphql,
	"03_auth_schema.graphql":   _03_auth_schemaGraphql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"01_base_schema.graphql":   &bintree{_01_base_schemaGraphql, map[string]*bintree{}},
	"02_unauth_schema.graphql": &bintree{_02_unauth_schemaGraphql, map[string]*bintree{}},
	"03_auth_schema.graphql":   &bintree{_03_auth_schemaGraphql, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
