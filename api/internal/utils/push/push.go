package push

import (
	daoPlatform "api/internal/dao/platform"
	"context"
)

/*
//tx_tpns(腾讯移动推送）标签推送规则格式
TagRules: []map[string]any{
	{
		`tag_items`: []map[string]any{
			{
				`tags`:           []string{`aaaa`},
				`is_not`:         false,
				`tags_operator`:  `OR`,
				`items_operator`: `OR`,
				`tag_type`:       `xg_user_define`,
			},
		},
		`operator`: `OR`,
		`is_not`:   false,
	},
} */

// 这里定义统一的参数格式！各插件内部再单独处理
type PushParam struct {
	IsDev      bool     //是否开发环境：false否 true是
	DeviceType uint     //设备类型：0-安卓 1-苹果 2-苹果电脑
	Audience   uint     //推送目标：0-全部 1-token方式 2-tag方式
	TokenList  []string //token列表
	// TagList    []string //tag列表
	TagRules      any           //标签推送规则。这参数较为复杂，不同插件差别极大，格式参考上面示例
	MessageType   uint          //消息类型：0通知消息 1透传消息
	Title         string        //消息标题
	Content       string        //消息内容
	CustomContent CustomContent //自定义数据
}

type CustomContent struct {
	Type string         //类型
	Data map[string]any //数据
}

type TagParam struct {
	OperatorType uint     //设备类型：0-增加 1-删除
	TagList      []string //tag列表
	TokenList    []string //token列表
}

type Push interface {
	Push(param PushParam) (err error)
	TagHandle(param TagParam) (err error)
}

// 设备类型：0安卓 1苹果 2苹果电脑
func NewPush(ctx context.Context, deviceType uint, pushTypeOpt ...string) Push {
	pushType := ``
	if len(pushTypeOpt) > 0 {
		pushType = pushTypeOpt[0]
	} else {
		pushType, _ = daoPlatform.Config.CtxDaoModel(ctx).Filter(daoPlatform.Config.Columns().ConfigKey, `pushType`).ValueStr(daoPlatform.Config.Columns().ConfigValue)
	}

	switch pushType {
	// case `pushOfTx`:	//腾讯移动推送
	default:
		var config map[string]any
		switch deviceType {
		case 1: //IOS
			configTmp, _ := daoPlatform.Config.Get(ctx, []string{`pushOfTxHost`, `pushOfTxIosAccessID`, `pushOfTxIosSecretKey`})
			config = map[string]any{
				`pushOfTxHost`:      configTmp[`pushOfTxHost`],
				`pushOfTxAccessID`:  configTmp[`pushOfTxIosAccessID`],
				`pushOfTxSecretKey`: configTmp[`pushOfTxIosSecretKey`],
			}
		case 2: //MacOS
			configTmp, _ := daoPlatform.Config.Get(ctx, []string{`pushOfTxHost`, `pushOfTxMacOSAccessID`, `pushOfTxMacOSSecretKey`})
			config = map[string]any{
				`pushOfTxHost`:      configTmp[`pushOfTxHost`],
				`pushOfTxAccessID`:  configTmp[`pushOfTxMacOSAccessID`],
				`pushOfTxSecretKey`: configTmp[`pushOfTxMacOSSecretKey`],
			}
		// case 0: //安卓
		default:
			configTmp, _ := daoPlatform.Config.Get(ctx, []string{`pushOfTxHost`, `pushOfTxAndroidAccessID`, `pushOfTxAndroidSecretKey`})
			config = map[string]any{
				`pushOfTxHost`:      configTmp[`pushOfTxHost`],
				`pushOfTxAccessID`:  configTmp[`pushOfTxAndroidAccessID`],
				`pushOfTxSecretKey`: configTmp[`pushOfTxAndroidSecretKey`],
			}
		}
		return NewPushOfTx(ctx, config)
	}
}
