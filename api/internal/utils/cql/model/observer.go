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
	obThis.Log.Debug(ctx, fmt.Sprintf(`[%d ms] [CQL] [%s] [%s] [rows:%d] %s`,
		observedQuery.End.UnixMilli()-observedQuery.Start.UnixMilli(),
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
		`BEGIN BATCH`,
	))
	for index, statement := range observedBatch.Statements {
		obThis.Log.Debug(ctx, fmt.Sprintf(`[CQL] [BATCH] [%s] [%s] %s`,
			obThis.Config.Group,
			observedBatch.Keyspace,
			fmt.Sprintf(gstr.Replace(statement, `?`, `%v`), observedBatch.Values[index]...),
		))
	}
	obThis.Log.Debug(ctx, fmt.Sprintf(`[CQL] [BATCH] [%d ms] [%s] [%s] %s`,
		observedBatch.End.UnixMilli()-observedBatch.Start.UnixMilli(),
		obThis.Config.Group,
		observedBatch.Keyspace,
		`APPLY BATCH`,
	))

	if observedBatch.Err != nil {
		obThis.Log.Error(ctx, observedBatch.Err.Error())
	}
}
