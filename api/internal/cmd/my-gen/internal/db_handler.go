package internal

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type MyGenDbFuncCode = string

const (
	DbFuncCodeNULLIF    MyGenDbFuncCode = `NULLIF`    //Mysql和Postgresql通用（差别：Postgresql数字字段需加::TEXT转成字符串）
	DbFuncCodeCOALESCE  MyGenDbFuncCode = `COALESCE`  //Mysql和Postgresql通用（差别：Postgresql数字字段需加::TEXT转成字符串）
	DbFuncCodeREPLACE   MyGenDbFuncCode = `REPLACE`   //Mysql和Postgresql通用（差别：Postgresql数字字段需加::TEXT转成字符串）
	DbFuncCodeCONCAT    MyGenDbFuncCode = `CONCAT`    //Mysql和Postgresql通用
	DbFuncCodeCONCAT_WS MyGenDbFuncCode = `CONCAT_WS` //Mysql和Postgresql通用
)

type MyGenDbHandler interface {
	GetTableArr(ctx context.Context, group string) (tableArr []string)                                                    // 获取数据库表数组
	GetFieldList(ctx context.Context, group, table string) (fieldList []MyGenField)                                       // 获取表字段列表
	GetKeyList(ctx context.Context, group, table string) (keyList []MyGenKey)                                             // 获取表索引列表
	GetFieldType(ctx context.Context, field MyGenField, group, table string) (fieldType MyGenFieldType)                   // 获取字段类型
	GetFieldLimitStr(ctx context.Context, field MyGenField, group, table string) (fieldLimitStr string)                   // 获取字符串字段限制
	GetFieldLimitInt(ctx context.Context, field MyGenField, group, table string) (fieldLimitInt MyGenFieldLimitInt)       // 获取整数字段限制
	GetFieldLimitFloat(ctx context.Context, field MyGenField, group, table string) (fieldLimitFloat MyGenFieldLimitFloat) // 获取浮点数字段限制
	GetFuncFieldFormat(dbFuncCode MyGenDbFuncCode, field string) (fieldFormat string)                                     // 获取用于部分函数的字段格式
}

type MyGenField struct {
	FieldRaw     string // 字段（原始）
	FieldTypeRaw string // 字段类型（原始）
	IsNull       bool   // 字段是否可为NULL
	Default      any    // 默认值
	Comment      string // 注释（原始）。
	IsAutoInc    bool   // 是否自增
}

type MyGenKey struct {
	FieldArr  []string // 字段数组。联合索引有多字段，需按顺序存入
	IsPrimary bool     // 是否主键
	IsUnique  bool     // 是否唯一
	IsAutoInc bool     // 是否自增
}

type MyGenFieldLimitInt struct {
	Size int    // 2字节，4字节，8字节
	Min  string // 最小值
	Max  string // 最大值
}

type MyGenFieldLimitFloat struct {
	Size      int    // 位数（含小数位）
	Precision int    // 小数位
	Min       string // 最小值
	Max       string // 最大值
}

type common struct{}

func (common) GetTableArr(ctx context.Context, group string) (tableArr []string) {
	tableArr, _ = g.DB(group).Tables(ctx)
	return
}

func NewMyGenDbHandler(ctx context.Context, dbType string) MyGenDbHandler {
	switch dbType {
	// case `sqlite`:
	// case `mssql`:
	case `pgsql`:
		return pgsql{}
	// case `oracle`:
	// case `mysql`:
	default:
		return mysql{}
	}
}
