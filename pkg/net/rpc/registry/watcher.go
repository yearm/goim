package registry

import (
	"context"
	"encoding/json"
	"go.etcd.io/etcd/clientv3"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/resolver"
	"sync"
)

type watcher struct {
	key    string
	client *clientv3.Client
	addrs  []resolver.Address
	wg     sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
}

func newWatcher(key string, cli *clientv3.Client) *watcher {
	ctx, cancel := context.WithCancel(context.Background())
	w := &watcher{
		key:    key,
		client: cli,
		ctx:    ctx,
		cancel: cancel,
	}
	return w
}

func (w *watcher) Watch() chan []resolver.Address {
	out := make(chan []resolver.Address, 10)
	w.wg.Add(1)
	go func() {
		defer func() {
			close(out)
			w.wg.Done()
		}()
		w.addrs = w.getAllAddresses()
		out <- w.cloneAddresses(w.addrs)

		rch := w.client.Watch(w.ctx, w.key, clientv3.WithPrefix())
		for watchRes := range rch {
			for _, ev := range watchRes.Events {
				switch ev.Type {
				case clientv3.EventTypePut:
					var nodeData NodeData
					err := json.Unmarshal([]byte(ev.Kv.Value), &nodeData)
					if err != nil {
						grpclog.Error("Parse node data error:", err)
						continue
					}
					addr := resolver.Address{Addr: nodeData.Addr, Metadata: &nodeData.Metadata}
					if w.addAddr(addr) {
						out <- w.cloneAddresses(w.addrs)
					}
				case clientv3.EventTypeDelete:
					var nodeData NodeData
					err := json.Unmarshal([]byte(ev.Kv.Value), &nodeData)
					if err != nil {
						grpclog.Error("Parse node data error:", err)
						continue
					}
					addr := resolver.Address{Addr: nodeData.Addr, Metadata: &nodeData.Metadata}
					if w.removeAddr(addr) {
						out <- w.cloneAddresses(w.addrs)
					}
				}
			}
		}
	}()
	return out
}

func (w *watcher) getAllAddresses() []resolver.Address {
	ret := make([]resolver.Address, 0)
	resp, err := w.client.Get(w.ctx, w.key, clientv3.WithPrefix())
	if err == nil {
		addrs := w.extractAddrs(resp)
		if len(addrs) > 0 {
			for _, addr := range addrs {
				v := addr
				ret = append(ret, resolver.Address{
					Addr:     v.Addr,
					Metadata: &v.Metadata,
				})
			}
		}
	}
	return ret
}

func (w *watcher) extractAddrs(resp *clientv3.GetResponse) []NodeData {
	addrs := make([]NodeData, 0)
	if resp == nil || resp.Kvs == nil {
		return addrs
	}

	for i := range resp.Kvs {
		if v := resp.Kvs[i].Value; v != nil {
			var nodeData NodeData
			err := json.Unmarshal(v, &nodeData)
			if err != nil {
				grpclog.Error("Parse node data error:", err)
				continue
			}
			addrs = append(addrs, nodeData)
		}
	}
	return addrs
}

func (w *watcher) cloneAddresses(in []resolver.Address) []resolver.Address {
	out := make([]resolver.Address, len(in))
	for i := 0; i < len(in); i++ {
		out[i] = in[i]
	}
	return out
}

func (w *watcher) addAddr(addr resolver.Address) bool {
	for _, v := range w.addrs {
		if addr.Addr == v.Addr {
			return false
		}
	}
	w.addrs = append(w.addrs, addr)
	return true
}

func (w *watcher) removeAddr(addr resolver.Address) bool {
	for i, v := range w.addrs {
		if addr.Addr == v.Addr {
			w.addrs = append(w.addrs[:i], w.addrs[i+1:]...)
			return true
		}
	}
	return false
}

func (w *watcher) Close() {
	w.cancel()
}
