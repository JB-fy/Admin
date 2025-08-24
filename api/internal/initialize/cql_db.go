package initialize

import (
	"api/internal/utils/jbcql"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func initCqlDb(ctx context.Context) {
	for group, config := range g.Cfg().MustGet(ctx, `cqlDb`).Map() {
		jbcql.AddDB(ctx, group, gconv.Map(config))
	}

	/* // 启动自动创建表时使用
	// 字段类型:int,smallint,bigint,decimal,text,blob,TIMESTAMP,DATE,TIME
	// `CREATE TABLE IF NOT EXISTS 表名 (字段 类型, 字段2 类型, PRIMARY KEY (（主键1, 主键2），多聚类键)) WITH default_time_to_live=过期秒数`
	cqlArr := []string{}
	cqlFormat := `CREATE TABLE IF NOT EXISTS %s (created_at TIMESTAMP, test_id text, test_info blob, PRIMARY KEY (test_id)) WITH default_time_to_live=%d`
	cqlArr = append(cqlArr, fmt.Sprintf(cqlFormat, `test`, 7*24*60*60))
	db := jbcql.DB()
	var err error
	for _, cql := range cqlArr {
		err = db.ExecStmt(cql)
		if err != nil {
			panic(err.Error())
		}
	} */
}
