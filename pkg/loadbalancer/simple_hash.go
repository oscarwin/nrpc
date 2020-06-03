package loadbalancer

import (
	"context"

	"github.com/oscarwin/nrpc/pkg/discovery"
)

type simpleHashPicker struct {
	instances []*discovery.Instance
}

func (p *simpleHashPicker) Pick(ctx context.Context, ops ...Option) (*discovery.Instance, bool) {
	options := newOptions()
	for _, op := range ops {
		op(options)
	}
	if len(p.instances) == 0 {
		return nil, false
	}
	pos := int(options.HashFunc(options.Key))
	pos = pos % len(p.instances)
	return p.instances[pos], true
}

type simpleHashLoadbalancer struct{}

func (l *simpleHashLoadbalancer) Name() string {
	return "simpleHashLoadbalancer"
}

func (l *simpleHashLoadbalancer) NewPicker(ctx context.Context, instances []*discovery.Instance) Picker {
	return &simpleHashPicker{
		instances: instances,
	}
}
