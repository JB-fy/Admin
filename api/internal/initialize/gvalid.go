package initialize

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gvalid"
)

func initGvalid(ctx context.Context) {
	myRuleThis := myRule{}

	gvalid.RegisterRule(`distinct`, myRuleThis.Distinct)
}

type myRule struct{}

// 数组不能含有重复值
func (myRule) Distinct(ctx context.Context, in gvalid.RuleFuncInput) (err error) {
	valSet := map[any]struct{}{}
	ok := false
	for _, v := range in.Value.Array() {
		if _, ok = valSet[v]; ok {
			//err = gerror.Newf(`%s字段具有重复值`, in.Field)
			err = gerror.New(in.Message) //这样才会被i18n翻译
			return
		}
		valSet[v] = struct{}{}
	}
	return
}
