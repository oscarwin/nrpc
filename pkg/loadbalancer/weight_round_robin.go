package loadbalancer

import (
	"context"
	"strconv"

	"github.com/oscarwin/nrpc/pkg/discovery"
)

type instanceWithWeight struct {
	instance       *discovery.Instance
	weight         int
	current_weight int
}

type weightRoundRobinPicker struct {
	wInstances []*instanceWithWeight
}

func (p *weightRoundRobinPicker) Pick(ctx context.Context, ops ...Option) (*discovery.Instance, bool) {
	var best *instanceWithWeight
	var totalWeight int
	for _, ins := range p.wInstances {
		// 每个实例的当前权重加上该实例的静态权重
		ins.current_weight += ins.weight
		// 累加所有实例的当前权重
		totalWeight += ins.current_weight
		if best == nil || ins.current_weight > best.current_weight {
			best = ins
		}
	}
	if best == nil {
		return nil, false
	}
	// 被选中的实例降低当前权重
	best.current_weight -= totalWeight
	return best.instance, true
}

type weightRoundRobinLoadbalancer struct{}

func (l *weightRoundRobinLoadbalancer) Name() string {
	return "weightPollingLoadbalancer"
}

func (l *weightRoundRobinLoadbalancer) NewPicker(ctx context.Context, instances []*discovery.Instance) Picker {
	instanceWithWeights := make([]*instanceWithWeight, 0)
	for _, instance := range instances {
		weight := 0
		if w, ok := instance.Tags["weight"]; ok {
			var err error
			weight, err = strconv.Atoi(w)
			if err != nil {
				weight = 0
			}
		}
		ins := &instanceWithWeight{
			instance:       instance,
			weight:         weight,
			current_weight: 0,
		}
		instanceWithWeights = append(instanceWithWeights, ins)
	}
	return &weightRoundRobinPicker{
		wInstances: instanceWithWeights,
	}
}
