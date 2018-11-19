package aws

import (
	"testing"

	"github.com/atlassian/escalator/pkg/test"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestInstanceToProviderID(t *testing.T) {
	instance := &autoscaling.Instance{
		AvailabilityZone: aws.String("us-east-1b"),
		InstanceId:       aws.String("abc123"),
	}
	res := instanceToProviderID(instance)
	assert.Equal(t, "aws:///us-east-1b/abc123", res)
}

func newMockCloudProvider(logger *logrus.Logger, nodeGroups []string, service *test.MockAutoscalingService) (*CloudProvider, error) {
	cloudProvider := &CloudProvider{
		logger:     logger,
		service:    service,
		nodeGroups: make(map[string]*NodeGroup, len(nodeGroups)),
	}
	return cloudProvider, cloudProvider.RegisterNodeGroups(nodeGroups...)
}
