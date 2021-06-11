/**
 * @Author: Resynz
 * @Date: 2021/6/11 15:10
 */
package aliyun_mobile_pusher

type Target string

const (
	TargetDevice  Target = "DEVICE"
	TargetAccount Target = "ACCOUNT"
	TargetAlias   Target = "ALIAS"
	TargetTag     Target = "TAG"
	TargetAll     Target = "ALL"
)

type ApnsEnv string

const (
	ApnsEnvDev  ApnsEnv = "DEV"
	ApnsEnvProd ApnsEnv = "PRODUCT"
)

type PushPlatform uint8

const (
	PushPlatformIOS PushPlatform = iota
	PushPlatformAndroid
)

type PushType uint8

const (
	PushMessage PushType = iota
	PushNotice
)
