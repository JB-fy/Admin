package model

import (
	"context"
)

type PushFunc func(ctx context.Context, config map[string]any) Push

type Push interface {
	Push(ctx context.Context, param PushParam) (err error)
	Tag(ctx context.Context, param TagParam) (err error)
}

type PushParam struct {
	IsDev      bool     //是否开发环境：false否 true是
	DeviceType uint     //设备类型：0-安卓 1-苹果 2-苹果电脑
	Audience   uint     //推送目标：0-全部 1-token方式 2-tag方式
	TokenList  []string //token列表
	// TagList    []string //tag列表
	TagRules      any           //标签推送规则。这参数较为复杂，不同推送平台差别应该很大
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
