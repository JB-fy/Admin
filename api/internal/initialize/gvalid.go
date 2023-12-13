package initialize

import (
	"context"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gvalid"
)

func initOfGvalid(ctx context.Context) {
	myRuleThis := myRule{}

	gvalid.RegisterRule(`distinct`, myRuleThis.Distinct)
}

type myRule struct{}

// 数组不能含有重复值
func (myRule) Distinct(ctx context.Context, in gvalid.RuleFuncInput) (err error) {
	val := in.Value.Array()
	if len(val) != garray.NewFrom(val).Unique().Len() {
		//err = gerror.Newf(`%s字段具有重复值`, in.Field)
		err = gerror.New(in.Message) //这样才会被i18n翻译
		return
	}
	return
}
