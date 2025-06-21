package model

import (
	"context"
	"fmt"

	"github.com/gocql/gocql"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/text/gstr"
)

type Log struct {
	Log       *glog.Logger
	CqlConfig *Config
}

func (logThis *Log) ObserveQuery(ctx context.Context, observedQuery gocql.ObservedQuery) {
	logThis.Log.Debug(ctx, fmt.Sprintf(`[%d ms] [CQL] [%s] [%s] [rows:%d] %s`,
		observedQuery.End.UnixMilli()-observedQuery.Start.UnixMilli(),
		logThis.CqlConfig.Group,
		observedQuery.Keyspace,
		observedQuery.Rows,
		fmt.Sprintf(gstr.Replace(observedQuery.Statement, `?`, `%v`), observedQuery.Values...),
	))
	if observedQuery.Err != nil {
		logThis.Log.Error(ctx, observedQuery.Err.Error())
	}
}

func (logThis *Log) ObserveBatch(ctx context.Context, observedBatch gocql.ObservedBatch) {
	logThis.Log.Debug(ctx, fmt.Sprintf(`[CQL] [BATCH] [%s] [%s] %s`,
		logThis.CqlConfig.Group,
		observedBatch.Keyspace,
		`BEGIN BATCH`,
	))
	for index, statement := range observedBatch.Statements {
		logThis.Log.Debug(ctx, fmt.Sprintf(`[CQL] [BATCH] [%s] [%s] %s`,
			logThis.CqlConfig.Group,
			observedBatch.Keyspace,
			fmt.Sprintf(gstr.Replace(statement, `?`, `%v`), observedBatch.Values[index]...),
		))
	}
	logThis.Log.Debug(ctx, fmt.Sprintf(`[CQL] [BATCH] [%d ms] [%s] [%s] %s`,
		observedBatch.End.UnixMilli()-observedBatch.Start.UnixMilli(),
		logThis.CqlConfig.Group,
		observedBatch.Keyspace,
		`APPLY BATCH`,
	))

	if observedBatch.Err != nil {
		logThis.Log.Error(ctx, observedBatch.Err.Error())
	}
}
