package service

import "goim/pkg/util/cmap"

var clientManager = cmap.New()

func store(token string, wsConn *WsClient) {
	clientManager.Set(token, wsConn)
}

func load(token string) *WsClient {
	value, ok := clientManager.Get(token)
	if ok {
		return value.(*WsClient)
	}
	return nil
}

func delete(token string) {
	clientManager.Remove(token)
}
