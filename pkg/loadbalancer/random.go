package loadbalancer

import (
	"context"
	"math/rand"
	"time"

	"github.com/oscarwin/nrpc/pkg/discovery"
)

type randomPicker struct {
	instances []*discovery.Instance
}

func (p *randomPicker) Pick(ctx context.Context, ops ...Option) (*discovery.Instance, bool) {
	if len(p.instances) == 0 {
		return nil, false
	}
	length := len(p.instances)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	pos := r.Int() % length
	return p.instances[pos], true
}

type randomLoadbalancer struct{}

func (l *randomLoadbalancer) Name() string {
	return "randomLoadbalancer"
}

func (l *randomLoadbalancer) NewPicker(ctx context.Context, instances []*discovery.Instance) Picker {
	return &randomPicker{
		instances: instances,
	}
}
