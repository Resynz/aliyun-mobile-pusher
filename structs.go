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
	Platform             PushPlatform `json:"platform"`
	Type                 PushType     `json:"type"`
	ApnsEnv              ApnsEnv      `json:"apns_env"`
	AppKey               string       `json:"app_key"`
	Title                string       `json:"title"`
	Body                 string       `json:"body"`
	TargetType           Target       `json:"target_type"`
	TargetValue          string       `json:"target_value"`
	Ext                  string       `json:"ext"`
	AndroidNotifyChannel string       `json:"android_notify_channel"`
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
