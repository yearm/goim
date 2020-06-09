package balancer

import (
	"context"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/resolver"
	"math/rand"
	"sync"
)

// 选择指定节点
const AddrPicker = "addr_picker"

const addrKey = "addr"

func init() {
	balancer.Register(newAddrBuilder())
}

func newAddrBuilder() balancer.Builder {
	return base.NewBalancerBuilderWithConfig(AddrPicker, &addrPickerBuilder{}, base.Config{HealthCheck: true})
}

func ContextWithAddr(ctx context.Context, addr string) context.Context {
	return context.WithValue(ctx, addrKey, addr)
}

type addrPickerBuilder struct{}

func (*addrPickerBuilder) Build(readySCs map[resolver.Address]balancer.SubConn) balancer.Picker {
	grpclog.Infof("addrPicker: newPicker called with readySCs: %v", readySCs)

	if len(readySCs) == 0 {
		return base.NewErrPicker(balancer.ErrNoSubConnAvailable)
	}
	var scs = make(map[string]balancer.SubConn)

	for addr, sc := range readySCs {
		scs[addr.Addr] = sc
	}

	return &addrPicker{
		subConns: scs,
		next:     rand.Intn(len(scs)),
	}
}

type addrPicker struct {
	subConns map[string]balancer.SubConn
	mu       sync.Mutex
	next     int
}

func (p *addrPicker) Pick(ctx context.Context, opts balancer.PickOptions) (balancer.SubConn, func(balancer.DoneInfo), error) {
	p.mu.Lock()
	address := ctx.Value(addrKey).(string)
	sc, ok := p.subConns[address]
	if !ok {
		return nil, nil, balancer.ErrNoSubConnAvailable
	}
	p.mu.Unlock()
	return sc, nil, nil
}
