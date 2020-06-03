package loadbalancer

import (
	"context"

	"github.com/oscarwin/nrpc/pkg/discovery"
)

type simpleRoundRobinPicker struct {
	pos       uint
	instances []*discovery.Instance
}

func (p *simpleRoundRobinPicker) Pick(ctx context.Context, ops ...Option) (*discovery.Instance, bool) {
	if len(p.instances) == 0 {
		return nil, false
	}
	next := p.pos % uint(len(p.instances))
	p.pos += 1
	return p.instances[next], true
}

type simpleRoundRobinLoadbalancer struct {
}

func (l *simpleRoundRobinLoadbalancer) Name() string {
	return "simpleRoundRobinLoadbalancer"
}

func (l *simpleRoundRobinLoadbalancer) NewPicker(ctx context.Context, instances []*discovery.Instance) Picker {
	return &simpleRoundRobinPicker{
		pos:       0,
		instances: instances,
	}
}
