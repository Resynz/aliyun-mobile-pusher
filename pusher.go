/**
 * @Author: Resynz
 * @Date: 2021/6/11 14:55
 */
package aliyun_mobile_pusher

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/push"
	"log"
)

type AliMobilePusher struct {
	Config *Config
	client *push.Client
}

func NewAliMobilePusher(conf *Config) (*AliMobilePusher, error) {
	client, err := push.NewClientWithAccessKey("cn-hangzhou", conf.AccessKeyId, conf.AccessSecret)
	if err != nil {
		return nil, err
	}
	pusher := &AliMobilePusher{Config: conf, client: client}
	return pusher, nil
}

func (s *AliMobilePusher) pushMessageIOS(param *PushParam) (*PushResponse, error) {
	request := push.CreatePushMessageToiOSRequest()
	request.Target = string(param.TargetType)
	request.Scheme = "https"
	request.AppKey = requests.Integer(param.AppKey)
	request.Title = param.Title
	request.Body = param.Body
	request.TargetValue = param.TargetValue
	r, err := s.client.PushMessageToiOS(request)

	if err != nil {
		return nil, err
	}
	res := &PushResponse{
		RequestId: r.RequestId,
		MessageId: r.MessageId,
	}
	return res, nil
}

func (s *AliMobilePusher) pushMessageAndroid(param *PushParam) (*PushResponse, error) {
	request := push.CreatePushMessageToAndroidRequest()
	request.Target = string(param.TargetType)
	request.Scheme = "https"
	request.AppKey = requests.Integer(param.AppKey)
	request.Title = param.Title
	request.Body = param.Body
	request.TargetValue = param.TargetValue
	r, err := s.client.PushMessageToAndroid(request)

	if err != nil {
		return nil, err
	}
	res := &PushResponse{
		RequestId: r.RequestId,
		MessageId: r.MessageId,
	}
	return res, nil
}

func (s *AliMobilePusher) pushNoticeIOS(param *PushParam) (*PushResponse, error) {
	request := push.CreatePushNoticeToiOSRequest()
	request.ExtParameters = param.Ext
	request.ApnsEnv = string(param.ApnsEnv)
	request.Title = param.Title
	request.Body = param.Body
	request.Target = string(param.TargetType)
	request.TargetValue = param.TargetValue
	request.AppKey = requests.Integer(param.AppKey)
	r, err := s.client.PushNoticeToiOS(request)
	if err != nil {
		return nil, err
	}
	res := &PushResponse{
		RequestId: r.RequestId,
		MessageId: r.MessageId,
	}
	return res, nil
}

func (s *AliMobilePusher) pushNoticeAndroid(param *PushParam) (*PushResponse, error) {
	//request := push.CreatePushNoticeToAndroidRequest()
	request := push.CreatePushRequest()
	request.PushType = "NOTICE"
	request.AndroidNotificationChannel = param.AndroidNotifyChannel
	request.AndroidExtParameters = param.Ext
	request.Title = param.Title
	request.Body = param.Body
	request.Target = string(param.TargetType)
	request.TargetValue = param.TargetValue
	request.AppKey = requests.Integer(param.AppKey)

	r, err := s.client.Push(request)
	if err != nil {
		return nil, err
	}
	res := &PushResponse{
		RequestId: r.RequestId,
		MessageId: r.MessageId,
	}
	return res, nil
}

func (s *AliMobilePusher) log(txt string) {
	if s.Config.IsDebug {
		log.Printf("[AliMobilePusher] %s\n", txt)
	}
}

// 推送
func (s *AliMobilePusher) Push(param *PushParam) (*PushResponse, error) {
	s.log(fmt.Sprintf("请求参数: %s", param.toJson()))
	var res *PushResponse
	var err error
	if param.Platform == PushPlatformIOS {
		if param.Type == PushMessage {
			res, err = s.pushMessageIOS(param)
		}
		if param.Type == PushNotice {
			res, err = s.pushNoticeIOS(param)
		}
	}

	if param.Platform == PushPlatformAndroid {
		if param.Type == PushMessage {
			res, err = s.pushMessageAndroid(param)
		}
		if param.Type == PushNotice {
			res, err = s.pushNoticeAndroid(param)
		}
	}

	if err != nil {
		s.log(fmt.Sprintf("请求失败！ error:%s", err.Error()))
		return nil, err
	}
	s.log(fmt.Sprintf("响应参数: %s", res.toJson()))
	return res, nil
}
