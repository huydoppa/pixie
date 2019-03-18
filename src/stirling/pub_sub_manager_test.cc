#include <google/protobuf/text_format.h>
#include <google/protobuf/util/message_differencer.h>
#include <gtest/gtest.h>

#include <cstring>
#include <vector>

#include "src/stirling/info_class_manager.h"
#include "src/stirling/pub_sub_manager.h"
#include "src/stirling/source_connector.h"

namespace pl {
namespace stirling {

using google::protobuf::TextFormat;
using google::protobuf::util::MessageDifferencer;
using stirlingpb::Element_State;
using stirlingpb::InfoClass;
using types::DataType;

const char* kInfoClassManager = R"(
  name : "cpu_usage",
  elements {
    name: "user_percentage",
    type: FLOAT64,
    state: SUBSCRIBED,
  }
  elements {
    name: "system_percentage",
    type: FLOAT64,
    state: SUBSCRIBED,
  }
  elements {
    name: "io_percentage",
    type: FLOAT64,
    state: SUBSCRIBED,
  }
  subscribed: false,
  sampling_period_millis: 100,
  push_period_millis: 1000
)";

/**
 * @brief This is a test source connector to be used for testing.
 *
 */
class TestSourceConnector : public SourceConnector {
 public:
  static std::unique_ptr<SourceConnector> Create(const std::string& name) {
    DataElements elements = {DataElement("user_percentage", DataType::FLOAT64),
                             DataElement("system_percentage", DataType::FLOAT64),
                             DataElement("io_percentage", DataType::FLOAT64)};
    return std::unique_ptr<SourceConnector>(new TestSourceConnector(name, elements));
  }

  Status InitImpl() override { return Status::OK(); }

  Status StopImpl() override { return Status::OK(); }

  void TransferDataImpl(ColumnWrapperRecordBatch* /*record_batch*/) override{};

 protected:
  explicit TestSourceConnector(const std::string& name, const DataElements& elements)
      : SourceConnector(SourceType::kUnknown, name, elements) {}
};

class PubSubManagerTest : public ::testing::Test {
 protected:
  PubSubManagerTest() = default;
  void SetUp() override {
    std::string name = "cpu_usage";
    source_ = TestSourceConnector::Create(name);
    info_class_mgrs_.push_back(std::make_unique<InfoClassManager>(name));
    info_class_mgrs_[0]->SetSourceConnector(source_.get());
    ASSERT_OK(info_class_mgrs_[0]->PopulateSchemaFromSource());
    pub_sub_manager_ = std::make_unique<PubSubManager>();
  }
  std::unique_ptr<SourceConnector> source_;
  std::unique_ptr<PubSubManager> pub_sub_manager_;
  InfoClassManagerVec info_class_mgrs_;
};

// This test validates that the Publish proto generated by the PubSubManager
// matches the expected Publish proto message (based on kInfoClassManager proto
// and with some fields added in the test).
TEST_F(PubSubManagerTest, publish_test) {
  // Publish info classes using proto message.
  stirlingpb::Publish actual_publish_pb;
  pub_sub_manager_->GeneratePublishProto(&actual_publish_pb, info_class_mgrs_);

  // Set expectations for the publish message.
  stirlingpb::Publish expected_publish_pb;
  auto expected_info_class = expected_publish_pb.add_published_info_classes();
  EXPECT_TRUE(TextFormat::MergeFromString(kInfoClassManager, expected_info_class));
  expected_info_class->set_id(0);

  EXPECT_FALSE(actual_publish_pb.published_info_classes(0).subscribed());
  EXPECT_EQ(1, expected_publish_pb.published_info_classes_size());
  EXPECT_EQ(0, actual_publish_pb.published_info_classes(0).id());
  EXPECT_EQ(InfoClassManager::kDefaultSamplingPeriod,
            actual_publish_pb.published_info_classes(0).sampling_period_millis());
  EXPECT_EQ(InfoClassManager::kDefaultPushPeriod,
            actual_publish_pb.published_info_classes(0).push_period_millis());

  EXPECT_TRUE(MessageDifferencer::Equals(actual_publish_pb, expected_publish_pb));
}

// This test validates that the InfoClassManager objects have their subscriptions
// updated after the PubSubManager reads a subscribe message (from an agent). The
// subscribe message is created using kInfoClassManager proto message and with fields from
// a Publish proto message.
TEST_F(PubSubManagerTest, subscribe_test) {
  // Do the publish.
  stirlingpb::Publish publish_pb;
  pub_sub_manager_->GeneratePublishProto(&publish_pb, info_class_mgrs_);

  // Get a subscription from an upstream agent.
  stirlingpb::Subscribe subscribe_pb;
  auto info_class = subscribe_pb.add_subscribed_info_classes();
  EXPECT_TRUE(TextFormat::MergeFromString(kInfoClassManager, info_class));

  // The subscribe message needs ids from the publish message and also
  // update the subscription.
  size_t id = publish_pb.published_info_classes(0).id();
  info_class->set_id(id);
  info_class->set_subscribed(true);

  // Update the InfoClassManager objects with the subscribe message.
  EXPECT_EQ(Status::OK(),
            pub_sub_manager_->UpdateSchemaFromSubscribe(subscribe_pb, info_class_mgrs_));
  // Verify updated subscriptions.
  for (auto& info_class_mgr : info_class_mgrs_) {
    EXPECT_TRUE(info_class_mgr->subscribed());
  }
}

}  // namespace stirling
}  // namespace pl
