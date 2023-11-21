package push

import (
	daoPlatform "api/internal/dao/platform"
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

// 这里定义统一的参数格式！各插件内部再单独处理
type PushOption struct {
	DeviceType    uint     //设备类型：0-安卓 1-苹果 2-苹果电脑
	Audience      uint     //推送目标：0全部 1单设备(token) 2多设备(token)
	TokenList     []string //token列表
	MessageType   uint     //消息类型：0通知消息 1透传消息
	Title         string   //消息标题
	Content       string   //消息内容
	CustomContent struct {
		Type string                 //类型
		Data map[string]interface{} //数据
	} //自定义数据
	IsDev bool //是否开发环境：false否 true是
}

type Push interface {
	Send(option PushOption) (err error)
}

// 设备类型：0-安卓 1-苹果 2-苹果电脑
func NewPush(ctx context.Context, deviceType uint) Push {
	pushType, _ := daoPlatform.Config.ParseDbCtx(ctx).Where(daoPlatform.Config.Columns().ConfigKey, `pushType`).Value(daoPlatform.Config.Columns().ConfigValue)
	switch pushType.String() {
	// case `txTpns`:	//腾讯移动推送
	default:
		config := g.Map{}
		switch deviceType {
		case 1: //IOS
			configTmp, _ := daoPlatform.Config.Get(ctx, []string{`txTpnsHost`, `txTpnsAccessIDOfIos`, `txTpnsSecretKeyOfIos`})
			config[`txTpnsHost`] = configTmp[`txTpnsHost`]
			config[`txTpnsAccessID`] = configTmp[`txTpnsAccessIDOfIos`]
			config[`txTpnsSecretKey`] = configTmp[`txTpnsSecretKeyOfIos`]
		case 2: //MacOS（暂时不做）
			configTmp, _ := daoPlatform.Config.Get(ctx, []string{`txTpnsHost`, `txTpnsAccessIDOfMacOS`, `txTpnsSecretKeyOfMacOS`})
			config[`txTpnsHost`] = configTmp[`txTpnsHost`]
			config[`txTpnsAccessID`] = configTmp[`txTpnsAccessIDOfMacOS`]
			config[`txTpnsSecretKey`] = configTmp[`txTpnsSecretKeyOfMacOS`]
		// case 0: //安卓
		default:
			configTmp, _ := daoPlatform.Config.Get(ctx, []string{`txTpnsHost`, `txTpnsAccessIDOfAndroid`, `txTpnsSecretKeyOfAndroid`})
			config[`txTpnsHost`] = configTmp[`txTpnsHost`]
			config[`txTpnsAccessID`] = configTmp[`txTpnsAccessIDOfAndroid`]
			config[`txTpnsSecretKey`] = configTmp[`txTpnsSecretKeyOfAndroid`]
		}
		return NewTxTpns(ctx, config)
	}

}
