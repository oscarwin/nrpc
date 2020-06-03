package loadbalancer

import (
	"context"
	"testing"

	"github.com/oscarwin/nrpc/pkg/discovery"
)

// go test -v -count=1 -test.run=TestRandomPick
func TestRandomPick(t *testing.T) {
	l := &randomLoadbalancer{}
	ctx := context.Background()
	ins := []*discovery.Instance{
		{
			Addr: "127.0.0.1",
		},
		{
			Addr: "127.0.0.2",
		},
		{
			Addr: "127.0.0.3",
		},
	}
	p := l.NewPicker(ctx, ins)
	for i := 0; i < 10; i++ {
		instance, _ := p.Pick(ctx)
		t.Logf("index: %d, addr: %s", i, instance.Addr)
	}
}
