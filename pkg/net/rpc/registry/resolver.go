package registry

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc/resolver"
	"sync"
)

const Schema = "etcd3"

type etcdResolver struct {
	etcdConfig    clientv3.Config
	etcdWatchPath string
	cc            resolver.ClientConn
	watcher       *watcher
	wg            sync.WaitGroup
}

func NewResolver(etcdConf clientv3.Config, registryDir, srvName string) {
	resolver.Register(&etcdResolver{
		etcdConfig:    etcdConf,
		etcdWatchPath: fmt.Sprintf("%s/%s", registryDir, srvName),
	})
}

func (r *etcdResolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOption) (resolver.Resolver, error) {
	cli, err := clientv3.New(r.etcdConfig)
	if err != nil {
		return nil, err
	}
	r.cc = cc
	r.watcher = newWatcher(r.etcdWatchPath, cli)
	r.start()
	return r, nil
}

func (r *etcdResolver) start() {
	r.wg.Add(1)
	go func() {
		defer r.wg.Done()
		out := r.watcher.Watch()
		for addr := range out {
			r.cc.UpdateState(resolver.State{Addresses: addr})
		}
	}()
}

func (r *etcdResolver) Scheme() string {
	return Schema
}

func (r *etcdResolver) ResolveNow(rn resolver.ResolveNowOption) {
}

func (r *etcdResolver) Close() {
	r.watcher.Close()
	r.wg.Wait()
}
