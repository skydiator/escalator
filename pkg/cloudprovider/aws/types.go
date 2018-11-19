package aws

import "fmt"

// NodeNotInAutoScalingGroup is a special error type
// this happens when a node is not inside a specific node group which it's expected to be
type NodeNotInAutoScalingGroup struct {
	NodeName   string
	ProviderID string
	NodeGroup  string
}

func (ne *NodeNotInAutoScalingGroup) Error() string {
	return fmt.Sprintf("node %v, %v belongs in a different asg than %v", ne.NodeName, ne.ProviderID, ne.NodeGroup)
}
