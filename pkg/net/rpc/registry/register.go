package registry

import (
	"context"
	"github.com/json-iterator/go"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
	"google.golang.org/grpc/grpclog"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Register struct {
	client *clientv3.Client
	key    string
	value  string
	ttl    int64
	ctx    context.Context
	cancel context.CancelFunc
}

type Option struct {
	EtcdConfig  clientv3.Config
	RegistryDir string
	ServiceName string
	NodeID      string
	NData       NodeData
	Ttl         int64
}

type NodeData struct {
	Addr     string
	Metadata map[string]string
}

func NewRegister(option Option) (*Register, error) {
	cli, err := clientv3.New(option.EtcdConfig)
	if err != nil {
		return nil, err
	}
	val, err := jsoniter.Marshal(option.NData)
	if err != nil {
		return nil, err
	}

	nodeId, err := createNodeId(cli, &option)
	if err != nil {
		return nil, err
	}
	option.NodeID = nodeId

	ctx, cancel := context.WithCancel(context.Background())
	registry := &Register{
		client: cli,
		key:    option.RegistryDir + "/" + option.ServiceName + "/" + option.NodeID,
		value:  string(val),
		ttl:    option.Ttl,
		ctx:    ctx,
		cancel: cancel,
	}
	return registry, nil
}

func (r *Register) Register() error {
	withAlive := func() error {
		leaseResp, err := r.client.Grant(r.ctx, r.ttl)
		if err != nil {
			grpclog.Errorf("registry: create grant failed: %s", err.Error())
			return err
		}
		_, err = r.client.Get(r.ctx, r.key)
		if err != nil {
			if err == rpctypes.ErrKeyNotFound {
				if _, err := r.client.Put(r.ctx, r.key, r.value, clientv3.WithLease(leaseResp.ID)); err != nil {
					grpclog.Error("registry: set service '%s' with ttl to etcd3 failed: %s", r.key, err.Error())
					return err
				}
			} else {
				grpclog.Errorf("registry: service '%s' connect to etcd3 failed: %s", r.key, err.Error())
				return err
			}
		} else {
			if _, err := r.client.KeepAlive(r.ctx, leaseResp.ID); err != nil {
				grpclog.Error("registry: keepAlive service '%s' to etcd3 failed: %s", r.key, err.Error())
				return err
			}
		}
		return nil
	}

	if err := withAlive(); err != nil {
		return err
	}
	ticker := time.NewTicker(time.Second * time.Duration(r.ttl))
	go func() {
		for {
			select {
			case <-ticker.C:
				withAlive()
			case <-r.ctx.Done():
				ticker.Stop()
			}
		}
	}()
	return nil
}

func (r *Register) UnRegister() {
	if _, err := r.client.Delete(context.Background(), r.key); err != nil {
		grpclog.Error("registry: unRegister '%s' failed: %s", r.key, err.Error())
	}
	r.cancel()
}

func createNodeId(cli *clientv3.Client, option *Option) (nodeId string, err error) {
	resp, err := cli.KV.Get(context.Background(), option.RegistryDir+"/"+option.ServiceName+"/", clientv3.WithPrefix(), clientv3.WithKeysOnly())
	if err != nil {
		return "", err
	}
	if len(resp.Kvs) == 0 {
		return "1", nil
	}
	ids := make(map[string]struct{})
	for _, v := range resp.Kvs {
		splits := strings.Split(string(v.Key), "/")
		nodeId := splits[len(splits)-1]
		ids[nodeId] = struct{}{}
	}
	for {
		nodeId := strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(resp.Kvs) + 1))
		if _, ok := ids[nodeId]; !ok {
			return nodeId, nil
		}
	}
}
