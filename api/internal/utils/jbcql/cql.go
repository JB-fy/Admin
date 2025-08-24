package jbcql

import (
	"api/internal/utils/jbcql/internal"
	"api/internal/utils/jbcql/model"
	"context"

	"github.com/scylladb/gocqlx/v3"
)

var sessionMap = map[string]*gocqlx.Session{}

func AddDB(ctx context.Context, group string, configMap map[string]any) {
	config := model.GetConfig(group, configMap)
	session, err := gocqlx.WrapSession(internal.InitSession(ctx, config))
	if err != nil {
		panic(`cql数据库连接失败：` + err.Error())
	}
	sessionMap[config.Group] = &session
}

func DB(opt ...string) (session *gocqlx.Session) {
	group := `default`
	if len(opt) > 0 && opt[0] != `` {
		group = opt[0]
	}
	session, ok := sessionMap[group]
	if !ok {
		panic(`cql数据库连接(分组:` + group + `)不存在`)
	}
	return session
}
