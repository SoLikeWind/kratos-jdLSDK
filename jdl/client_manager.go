package jdl

import (
	"context"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
)

type ClientManager struct {
	log     *log.Helper
	clients sync.Map
}

func NewClientManager(logger log.Logger) *ClientManager {
	log := log.NewHelper(log.With(logger, "module", "jdl"))
	return &ClientManager{
		log: log,
	}
}

func (cm *ClientManager) GetOrCreateClient(ctx context.Context, id uint64, createFn func(context.Context, uint64) (*Client, error)) (*Client, error) {
	if client, ok := cm.clients.Load(id); ok {
		return client.(*Client), nil
	}

	// 如果 client 不存在，创建一个新的
	newClient, err := createFn(ctx, id)
	if err != nil {
		log.Errorf("create client err: %s", err)
		return nil, err
	}

	log.Infof("create new client, id: %d", id)

	// 尝试将新创建的 client 存储到 map 中
	// 如果 ID 已经存在，则返回已存在的 client
	if client, loaded := cm.clients.LoadOrStore(id, newClient); loaded {
		return client.(*Client), nil
	}

	return newClient, nil
}

func (cm *ClientManager) ReloadClient(ctx context.Context, id uint64, createFn func(context.Context, uint64) (*Client, error)) (*Client, error) {
	newClient, err := createFn(ctx, id)
	if err != nil {
		log.Errorf("create client err: %s", err)
		return nil, err
	}

	cm.clients.Delete(id)
	cm.clients.Store(id, newClient)

	return newClient, nil
}

func (cm *ClientManager) DeleteClient(ctx context.Context, id uint64) {
	cm.clients.Delete(id)
}
