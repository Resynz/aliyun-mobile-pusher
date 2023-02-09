/**
 * @Author: Resynz
 * @Date: 2021/6/11 14:56
 */
package aliyun_mobile_pusher

import "encoding/json"

type Config struct {
	AccessKeyId  string `json:"access_key_id"`
	AccessSecret string `json:"access_secret"`
	IsDebug      bool   `json:"is_debug"`
}

// 推送参数
type PushParam struct {
	Platform                         PushPlatform `json:"platform"`
	Type                             PushType     `json:"type"`
	ApnsEnv                          ApnsEnv      `json:"apns_env"`
	AppKey                           string       `json:"app_key"`
	Title                            string       `json:"title"`
	Body                             string       `json:"body"`
	TargetType                       Target       `json:"target_type"`
	TargetValue                      string       `json:"target_value"`
	Ext                              string       `json:"ext"`
	AndroidNotifyChannel             string       `json:"android_notify_channel"`
	AndroidNotificationXiaomiChannel string       `json:"android_notification_xiaomi_channel"`
	AndroidNotificationVivoChannel   string       `json:"android_notification_vivo_channel"`
	AndroidNotificationHuaweiChannel string       `json:"android_notification_huawei_channel"`
	AndroidPopupActivity             string       `json:"android_popup_activity"`
}

func (p *PushParam) toJson() string {
	d, _ := json.Marshal(p)
	return string(d)
}

// 推送响应
type PushResponse struct {
	RequestId string `json:"request_id"`
	MessageId string `json:"message_id"`
}

func (p *PushResponse) toJson() string {
	d, _ := json.Marshal(p)
	return string(d)
}

type PushStat struct {
	AcceptCount   int64 `json:"accept_count"`   // 推送服务端接收到的推送数目
	OpenedCount   int64 `json:"opened_count"`   // 通知在设备上被点击的数目
	DeletedCount  int64 `json:"deleted_count"`  // 通知在设备上被清除的数目
	SentCount     int64 `json:"sent_count"`     // 推送服务端实际发出的数目
	ReceivedCount int64 `json:"received_count"` // 实际送达到设备的数目
}
