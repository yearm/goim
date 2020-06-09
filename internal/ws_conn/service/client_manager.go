package service

import "sync"

var clientManager sync.Map

func store(token string, wsConn *WsClient) {
	clientManager.Store(token, wsConn)
}

func load(token string) *WsClient {
	value, ok := clientManager.Load(token)
	if ok {
		return value.(*WsClient)
	}
	return nil
}

func delete(token string) {
	clientManager.Delete(token)
}
