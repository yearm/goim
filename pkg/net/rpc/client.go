package rpc

import (
	"errors"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"go.etcd.io/etcd/clientv3"
	"goim/pkg/net/rpc/registry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/keepalive"
	"time"
)

const (
	MaxSendMsgSize = 1<<31 - 1
	MaxCallMsgSize = 1<<31 - 1
)

// TODO 熔断器、限流器、链路追踪
type ClientConfig struct {
	EtcdConfig  clientv3.Config
	RegistryDir string
	ServiceName string
	BalanceName string
}

func GetClientConn(c *ClientConfig) *grpc.ClientConn {
	opts, err := addDialOption(c)
	if err != nil {
		grpclog.Fatalf("grpc client get opts: %s", err)
	}
	conn, err := grpc.Dial(fmt.Sprintf("%s:///", registry.Schema), opts...)
	if err != nil {
		grpclog.Fatalf("grpc dial: %s", err)
	}
	return conn
}

func addDialOption(c *ClientConfig) ([]grpc.DialOption, error) {
	if len(c.EtcdConfig.Endpoints) == 0 || c.ServiceName == "" || c.BalanceName == "" {
		return nil, errors.New("client rpc conn config error")
	}
	registry.NewResolver(c.EtcdConfig, c.RegistryDir, c.ServiceName)

	var (
		opts                     []grpc.DialOption
		unaryClientInterceptors  []grpc.UnaryClientInterceptor
		streamClientInterceptors []grpc.StreamClientInterceptor
	)
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBalancerName(c.BalanceName))
	opts = append(opts, grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(MaxCallMsgSize),
		grpc.MaxCallSendMsgSize(MaxSendMsgSize),
	))
	opts = append(opts, grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:                10 * time.Second,
		Timeout:             time.Second,
		PermitWithoutStream: true,
	}))
	if len(unaryClientInterceptors) > 0 {
		opts = append(opts, grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(unaryClientInterceptors...)))
	}
	if len(streamClientInterceptors) > 0 {
		opts = append(opts, grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(streamClientInterceptors...)))
	}
	return opts, nil
}
