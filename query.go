/**
 * @Author: Resynz
 * @Date: 2023/2/9 16:27
 */
package aliyun_mobile_pusher

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/push"
)

func (s *AliMobilePusher) QueryPushStatByMsg(appKey, msgId string) (*PushStat, error) {
	request := push.CreateQueryPushStatByMsgRequest()
	request.MessageId = requests.Integer(msgId)
	request.AppKey = requests.Integer(appKey)
	resp, err := s.client.QueryPushStatByMsg(request)
	if err != nil {
		return nil, err
	}
	if len(resp.PushStats.PushStat) == 1 {
		ps := resp.PushStats.PushStat[0]
		res := &PushStat{
			AcceptCount:   ps.AcceptCount,
			OpenedCount:   ps.OpenedCount,
			DeletedCount:  ps.DeletedCount,
			SentCount:     ps.SentCount,
			ReceivedCount: ps.ReceivedCount,
		}
		return res, nil
	}
	return nil, nil
}
