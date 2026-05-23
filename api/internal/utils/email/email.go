package email

import (
	"api/internal/utils/email/common"
	"api/internal/utils/email/model"
	"context"
	"sync"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"golang.org/x/sync/singleflight"
)

var (
	emailMap     sync.Map
	emailSfg     singleflight.Group
	emailTypeDef = `email_of_common`
	emailFuncMap = map[string]model.EmailFunc{
		`email_of_common`: common.NewEmail,
	}
)

func NewEmail(ctx context.Context, emailType string, config map[string]any) (obj model.Email) {
	if _, ok := emailFuncMap[emailType]; !ok {
		emailType = emailTypeDef
	}
	key := emailType + gmd5.MustEncrypt(config)
	objTmp, ok := emailMap.Load(key)
	if !ok {
		objTmp, _, _ = emailSfg.Do(key, func() (obj any, err error) {
			obj = emailFuncMap[emailType](ctx, config)
			emailMap.Store(key, obj)
			return
		})
	}
	obj = objTmp.(model.Email)
	return
}
