package loadbalancer

import (
	"context"

	"github.com/oscarwin/nrpc/pkg/discovery"
)

type Picker interface {
	Pick(ctx context.Context, ops ...Option) (*discovery.Instance, bool)
}
type LoadBalancer interface {
	Name() string
	NewPicker(ctx context.Context, instance *discovery.Instance) Picker
}
