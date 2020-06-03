package loadbalancer

import (
	"context"
	"strconv"
	"testing"

	"github.com/oscarwin/nrpc/pkg/discovery"
)

// go test -v -count=1 -test.run=TestSimpleHashPick
func TestSimpleHashPick(t *testing.T) {
	l := simpleHashLoadbalancer{}
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
		op := WithKey(strconv.FormatInt(int64(i), 10))
		ins, _ := p.Pick(ctx, op)
		t.Logf("index: %d, addr: %s", i, ins.Addr)
	}
}
