package balancer

import (
	"context"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/resolver"
	"math/rand"
	"strconv"
	"sync"
)

// 轮询调度选择节点
const RoundRobin = "round_robin"

func init() {
	balancer.Register(newRoundRobinBuilder())
}

func newRoundRobinBuilder() balancer.Builder {
	return base.NewBalancerBuilderWithConfig(RoundRobin, &roundRobinPickerBuilder{}, base.Config{HealthCheck: true})
}

type roundRobinPickerBuilder struct{}

func (*roundRobinPickerBuilder) Build(readySCs map[resolver.Address]balancer.SubConn) balancer.Picker {
	grpclog.Infof("round_robinPicker: newPicker called with readySCs: %v", readySCs)

	if len(readySCs) == 0 {
		return base.NewErrPicker(balancer.ErrNoSubConnAvailable)
	}
	var scs []balancer.SubConn
	for addr, sc := range readySCs {
		weight := 1
		m, ok := addr.Metadata.(*map[string]string)
		w, ok := (*m)["weight"]
		if ok {
			n, err := strconv.Atoi(w)
			if err == nil && n > 0 {
				weight = n
			}
		}
		for i := 0; i < weight; i++ {
			scs = append(scs, sc)
		}
	}

	return &roundRobinPicker{
		subConns: scs,
		next:     rand.Intn(len(scs)),
	}
}

type roundRobinPicker struct {
	subConns []balancer.SubConn
	mu       sync.Mutex
	next     int
}

func (p *roundRobinPicker) Pick(ctx context.Context, opts balancer.PickOptions) (balancer.SubConn, func(balancer.DoneInfo), error) {
	p.mu.Lock()
	sc := p.subConns[p.next]
	p.next = (p.next + 1) % len(p.subConns)
	p.mu.Unlock()
	return sc, nil, nil
}
