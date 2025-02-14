package id_card

import (
	"api/internal/utils/id-card/aliyun"
	"api/internal/utils/id-card/model"
	"context"
	"sync"

	"github.com/gogf/gf/v2/crypto/gmd5"
)

var (
	idCardMap     = map[string]model.IdCard{} //存放不同配置实例。因初始化只有一次，故重要的是读性能，普通map比sync.Map的读性能好
	idCardMu      sync.Mutex
	idCardTypeDef = `idCardOfAliyun`
	idCardFuncMap = map[string]model.IdCardFunc{
		`idCardOfAliyun`: aliyun.NewIdCard,
	}
)

func NewIdCard(ctx context.Context, idCardType string, config map[string]any) (idCard model.IdCard) {
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
