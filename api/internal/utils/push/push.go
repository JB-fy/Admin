package push

import (
	"context"
	"sync"

	"github.com/gogf/gf/v2/crypto/gmd5"
)

/* //tx_tpns(腾讯移动推送）标签推送规则格式
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
	PushMsg(ctx context.Context, param PushParam) (err error)
	TagHandle(ctx context.Context, param TagParam) (err error)
}

var (
	pushTypeDef = `pushOfTx`
	pushFuncMap = map[string]func(ctx context.Context, config map[string]any) Push{
		`pushOfTx`: func(ctx context.Context, config map[string]any) Push { return NewPushOfTx(ctx, config) },
	}
	pushMap = map[string]Push{} //存放不同配置实例。因初始化只有一次，故重要的是读性能，普通map比sync.Map的读性能好
	pushMu  sync.Mutex
)

func NewPush(ctx context.Context, pushType string, config map[string]any) (push Push) {
	pushKey := pushType + gmd5.MustEncrypt(config)
	ok := false
	if push, ok = pushMap[pushKey]; ok { //先读一次（不加锁）
		return
	}
	pushMu.Lock()
	defer pushMu.Unlock()
	if push, ok = pushMap[pushKey]; ok { // 再读一次（加锁），防止重复初始化
		return
	}
	if _, ok = pushFuncMap[pushType]; !ok {
		pushType = pushTypeDef
	}
	push = pushFuncMap[pushType](ctx, config)
	pushMap[pushKey] = push
	return
}
