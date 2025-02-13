package id_card

import (
	"context"
	"sync"

	"github.com/gogf/gf/v2/crypto/gmd5"
)

type IdCardInfo struct {
	Gender uint // 性别：0未设置 1男 2女
	// Birthday *gtime.Time // 生日
	Birthday string // 生日
	Address  string // 详细地址
}

type IdCard interface {
	Auth(ctx context.Context, idCardName string, idCardNo string) (idCardInfo IdCardInfo, err error)
}

var (
	idCardTypeDef = `idCardOfAliyun`
	idCardFuncMap = map[string]func(ctx context.Context, config map[string]any) IdCard{
		`idCardOfAliyun`: func(ctx context.Context, config map[string]any) IdCard { return NewIdCardOfAliyun(ctx, config) },
	}
	idCardMap = map[string]IdCard{} //存放不同配置实例。因初始化只有一次，故重要的是读性能，普通map比sync.Map的读性能好
	idCardMu  sync.Mutex
)

func NewIdCard(ctx context.Context, idCardType string, config map[string]any) (idCard IdCard) {
	idCardKey := idCardType + gmd5.MustEncrypt(config)
	ok := false
	if idCard, ok = idCardMap[idCardKey]; ok { //先读一次（不加锁）
		return
	}
	idCardMu.Lock()
	defer idCardMu.Unlock()
	if idCard, ok = idCardMap[idCardKey]; ok { // 再读一次（加锁），防止重复初始化
		return
	}
	if _, ok = idCardFuncMap[idCardType]; !ok {
		idCardType = idCardTypeDef
	}
	idCard = idCardFuncMap[idCardType](ctx, config)
	idCardMap[idCardKey] = idCard
	return

}
