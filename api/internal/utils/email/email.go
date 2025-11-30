package email

import (
	"api/internal/utils/email/common"
	"api/internal/utils/email/model"
	"context"
	"sync"

	"github.com/gogf/gf/v2/crypto/gmd5"
)

var (
	emailMap     = map[string]model.Email{} //存放不同配置实例。因初始化只有一次，故重要的是读性能，普通map比sync.Map的读性能好
	emailMuMap   sync.Map
	emailTypeDef = `email_of_common`
	emailFuncMap = map[string]model.EmailFunc{
		`email_of_common`: common.NewEmail,
	}
)

func NewEmail(ctx context.Context, emailType string, config map[string]any) (email model.Email) {
	emailKey := emailType + gmd5.MustEncrypt(config)
	ok := false
	if email, ok = emailMap[emailKey]; ok { //先读一次（不加锁）
		return
	}
	muTmp, _ := emailMuMap.LoadOrStore(emailKey, &sync.Mutex{})
	mu := muTmp.(*sync.Mutex)
	mu.Lock()
	defer func() {
		mu.Unlock()
		emailMuMap.Delete(emailKey)
	}()
	if email, ok = emailMap[emailKey]; ok { // 再读一次（加锁），防止重复初始化
		return
	}
	if _, ok = emailFuncMap[emailType]; !ok {
		emailType = emailTypeDef
	}
	email = emailFuncMap[emailType](ctx, config)
	emailMap[emailKey] = email
	return
}
