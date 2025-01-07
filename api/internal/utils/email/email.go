package email

import (
	"context"
	"sync"

	"github.com/gogf/gf/v2/crypto/gmd5"
)

type Email interface {
	SendEmail(ctx context.Context, message string, toEmailArr ...string) (err error)
	GetFromEmail() (fromEmail string)
}

var (
	emailTypeDef = `emailOfCommon`
	emailFuncMap = map[string]func(ctx context.Context, config map[string]any) Email{
		`emailOfCommon`: func(ctx context.Context, config map[string]any) Email { return NewEmailOfCommon(ctx, config) },
	}
	emailMap = map[string]Email{} //存放不同配置实例。因初始化只有一次，故重要的是读性能，普通map比sync.Map的读性能好
	emailMu  sync.Mutex
)

func NewEmail(ctx context.Context, emailType string, config map[string]any) (email Email) {
	emailKey := emailType + gmd5.MustEncrypt(config)
	ok := false
	if email, ok = emailMap[emailKey]; ok { //先读一次（不加锁）
		return
	}
	emailMu.Lock()
	defer emailMu.Unlock()
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
