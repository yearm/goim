package dao

import (
	"fmt"
	"github.com/json-iterator/go"
	"goim/pkg/proto/logic"
	"goim/pkg/util/stringi"
)

// 设备信息
type DevicesInfo struct {
	ServerId string `json:"serverId"` // 节点信息
	Agent    string `json:"agent"`
}

func getDeviceKey(uid string) string {
	return fmt.Sprintf("im_user_device_%s", uid)
}

func (d *Dao) SetDevice(uid string, in *pb_logic.SignInReq) {
	device, err := jsoniter.Marshal(&DevicesInfo{ServerId: in.WsAddr, Agent: in.Agent})
	if err == nil {
		d.rc.HSet(getDeviceKey(uid), in.Token, stringi.Bytes2str(device))
	}
}

func (d *Dao) DelDevice(uid string, in *pb_logic.OfflineReq) {
	d.rc.HDel(getDeviceKey(uid), in.Token)
}
