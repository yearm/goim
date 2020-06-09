package rpc

import (
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"goim/pkg/net/rpc/registry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	maxSendMsgSize = 1<<31 - 1
	maxCallMsgSize = 1<<31 - 1
)

// TODO 熔断器、Prometheus监控、限流器、链路追踪、健康检查
type ServerConfig struct {
	Addr      string
	RegOption registry.Option
}

type Server struct {
	conf     *ServerConfig
	addr     string
	srv      *grpc.Server
	register *registry.Register
	wg       *sync.WaitGroup
}

func NewServer(c *ServerConfig) *Server {
	server, err := newServer(c)
	if err != nil {
		panic(fmt.Errorf("fatal error NewServer(): %s \n", err))
	}
	return server
}

func newServer(c *ServerConfig) (*Server, error) {
	register, err := registry.NewRegister(c.RegOption)
	if err != nil {
		return nil, err
	}
	if err := register.Register(); err != nil {
		return nil, err
	}

	srv := grpc.NewServer(addServerOption(c)...)
	return &Server{
		conf:     c,
		addr:     c.Addr,
		srv:      srv,
		register: register,
		wg:       new(sync.WaitGroup),
	}, nil
}

func addServerOption(c *ServerConfig) []grpc.ServerOption {
	var (
		opts                     []grpc.ServerOption
		unaryServerInterceptors  []grpc.UnaryServerInterceptor
		streamServerInterceptors []grpc.StreamServerInterceptor
	)
	opts = append(opts, grpc.MaxRecvMsgSize(maxCallMsgSize))
	opts = append(opts, grpc.MaxSendMsgSize(maxSendMsgSize))
	opts = append(opts, grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle:     15 * time.Second,
		MaxConnectionAge:      30 * time.Second,
		MaxConnectionAgeGrace: 5 * time.Second,
		Time:                  5 * time.Second,
		Timeout:               1 * time.Second,
	}))

	if len(unaryServerInterceptors) > 0 {
		opts = append(opts, grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(unaryServerInterceptors...)))
	}
	if len(streamServerInterceptors) > 0 {
		opts = append(opts, grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(streamServerInterceptors...)))
	}
	return opts
}

func (s *Server) Run(registerFunc func(s *grpc.Server)) {
	registerFunc(s.srv)

	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		panic(fmt.Errorf("fatal error net.Listen(): %s \n", err))
	}

	s.wg.Add(1)
	go func() {
		log.Printf("rpc listening on %s", s.addr)
		if err := s.srv.Serve(listener); err != nil {
			panic(fmt.Errorf("fatal error s.srv.Server(): %s \n", err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT)
	q := <-quit
	log.Printf("%s:%s rpc server get a signal %s", s.conf.RegOption.ServiceName, s.addr, q.String())
	s.register.UnRegister()
	s.srv.Stop()
	s.wg.Wait()
}

func (s *Server) Stop() {
	s.srv.GracefulStop()
}
