package model

import (
	"context"
	"fmt"

	"github.com/gocql/gocql"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/text/gstr"
)

type Observer struct {
	Config *Config
	Log    *glog.Logger
}

func (obThis *Observer) ObserveQuery(ctx context.Context, observedQuery gocql.ObservedQuery) {
	obThis.Log.Debug(ctx, fmt.Sprintf(`[CQL] [%f ms] [%s] [%s] [rows:%d] %s`,
		float64(observedQuery.End.UnixMilli()-observedQuery.Start.UnixMilli())/1000,
		obThis.Config.Group,
		observedQuery.Keyspace,
		observedQuery.Rows,
		fmt.Sprintf(gstr.Replace(observedQuery.Statement, `?`, `%v`), observedQuery.Values...),
	))
	if observedQuery.Err != nil {
		obThis.Log.Error(ctx, observedQuery.Err.Error())
	}
}

func (obThis *Observer) ObserveBatch(ctx context.Context, observedBatch gocql.ObservedBatch) {
	obThis.Log.Debug(ctx, fmt.Sprintf(`[CQL] [BATCH] [%s] [%s] %s`,
		obThis.Config.Group,
		observedBatch.Keyspace,
		`BATCH START`,
	))
	for index, statement := range observedBatch.Statements {
		obThis.Log.Debug(ctx, fmt.Sprintf(`[CQL] [BATCH] [%s] [%s] %s`,
			obThis.Config.Group,
			observedBatch.Keyspace,
			fmt.Sprintf(gstr.Replace(statement, `?`, `%v`), observedBatch.Values[index]...),
		))
	}
	obThis.Log.Debug(ctx, fmt.Sprintf(`[CQL] [BATCH] [%f ms] [%s] [%s] %s`,
		float64(observedBatch.End.UnixMilli()-observedBatch.Start.UnixMilli())/1000,
		obThis.Config.Group,
		observedBatch.Keyspace,
		`BATCH END`,
	))

	if observedBatch.Err != nil {
		obThis.Log.Error(ctx, observedBatch.Err.Error())
	}
}
