package loadbalancer

import (
	"context"
	"testing"

	"github.com/oscarwin/nrpc/pkg/discovery"
)

// go test -v -count=1 -test.run TestWeightRoundRobinPick .
func TestWeightRoundRobinPick(t *testing.T) {
	l := weightRoundRobinLoadbalancer{}
	ctx := context.Background()
	ins := []*discovery.Instance{
		{
			Addr: "127.0.0.1",
			Tags: map[string]string{
				"weight": "5",
			},
		},
		{
			Addr: "127.0.0.2",
			Tags: map[string]string{
				"weight": "2",
			},
		},
		{
			Addr: "127.0.0.3",
			Tags: map[string]string{
				"weight": "1",
			},
		},
	}
	p := l.NewPicker(ctx, ins)
	for i := 0; i < 20; i++ {
		instance, ok := p.Pick(ctx)
		if !ok {
			t.Fatalf("pick instance failed")
		}
		t.Logf("index: %d, addr: %s", i, instance.Addr)
	}
}
