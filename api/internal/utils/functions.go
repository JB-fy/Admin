package utils

import (
	"api/internal/consts"
	"context"
	"os/exec"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// 生成错误码
func NewErrorCode(ctx context.Context, code int, msg string, data ...map[string]interface{}) error {
	dataTmp := map[string]interface{}{}
	if len(data) > 0 && data[0] != nil {
		dataTmp = data[0]
	}
	if msg == `` {
		switch code {
		case 29991062, 89999996:
			msg = g.I18n().Tf(ctx, `code.`+gconv.String(code), gconv.String(dataTmp[`errField`]))
			delete(dataTmp, `errField`)
		default:
			msg = g.I18n().T(ctx, `code.`+gconv.String(code))
		}
	}
	return gerror.NewCode(gcode.New(code, ``, dataTmp), msg)
}

// http接口调用失败返回json
func HttpFailJson(r *ghttp.Request, err error) {
	resData := map[string]interface{}{
		`code`: 99999999,
		`msg`:  err.Error(),
		`data`: map[string]interface{}{},
	}
	code := gerror.Code(err)
	if code.Code() > 0 {
		resData[`code`] = code.Code()
		resData[`data`] = code.Detail()
	}
	r.Response.WriteJsonExit(resData)
}

// http接口调用成功返回json
func HttpSuccessJson(r *ghttp.Request, data map[string]interface{}, code int, msg ...string) {
	resData := map[string]interface{}{
		`code`: code,
		`msg`:  ``,
		`data`: data,
	}
	if len(msg) == 0 || msg[0] == `` {
		resData[`msg`] = g.I18n().T(r.GetCtx(), `code.`+gconv.String(code))
	} else {
		resData[`msg`] = msg[0]
	}
	r.Response.WriteJsonExit(resData)
}

// 设置场景信息
func SetCtxSceneInfo(r *ghttp.Request, info gdb.Record) {
	r.SetCtxVar(consts.ConstCtxSceneInfoName, info)
}

// 获取场景信息
func GetCtxSceneInfo(ctx context.Context) gdb.Record {
	return ctx.Value(consts.ConstCtxSceneInfoName).(gdb.Record)
}

// 获取场景标识
func GetCtxSceneCode(ctx context.Context) string {
	return GetCtxSceneInfo(ctx)[`sceneCode`].String()
}

// 设置登录身份信息
func SetCtxLoginInfo(r *ghttp.Request, info gdb.Record) {
	r.SetCtxVar(consts.ConstCtxLoginInfoName, info)
}

// 获取登录身份信息
func GetCtxLoginInfo(ctx context.Context) gdb.Record {
	return ctx.Value(consts.ConstCtxLoginInfoName).(gdb.Record)
}

// 设置服务器外网ip
func GetServerNetworkIp() string {
	cmd := exec.Command(`/bin/bash`, `-c`, `curl -s ifconfig.me`)
	output, _ := cmd.CombinedOutput()
	return string(output)
}

// 设置服务器内网ip
func GetServerLocalIp() string {
	cmd := exec.Command(`/bin/bash`, `-c`, `hostname -I`)
	output, _ := cmd.CombinedOutput()
	return string(output)
}

// 数据库表按时间做分区（通用，默认以分区最大日期作为分区名）
// 注意：如果表有唯一索引（含主键），则用于建立分区的字段必须和唯一索引字段组成联合索引
// 建议：分区间隔时间，分区数量设置后，两者总时长要比定时器间隔多几天时间，方便分区失败时，有时间让技术人工处理
// dbGroup			db分组
// dbTable			db表
// partitionNumber	当前时间后面，需要新增的分区数量
// partitionTime	间隔多长时间创建一个分区，单位：秒
// partitionField	分区字段，即根据该字段做分区
func DbTablePartition(ctx context.Context, dbGroup string, dbTable string, partitionNumber int64, partitionTime int64, partitionField string) (err error) {
	/* //查看分区
	SELECT PARTITION_NAME, PARTITION_EXPRESSION, PARTITION_DESCRIPTION, TABLE_ROWS
	FROM INFORMATION_SCHEMA.PARTITIONS
	WHERE TABLE_SCHEMA = SCHEMA() AND TABLE_NAME = '表名';
	//修改表为分区表
	ALTER TABLE `表名` PARTITION BY RANGE (TO_DAYS( 分区字段 ))
	(PARTITION `20220115` VALUES LESS THAN (TO_DAYS('2022-01-16')) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
	PARTITION `max` VALUES LESS THAN (MAXVALUE) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 );
	//新增分区
	ALTER TABLE `表名` ADD PARTITION
	(PARTITION `20220115` VALUES LESS THAN (TO_DAYS('2022-01-16')) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 ,
	PARTITION `20220116` VALUES LESS THAN (TO_DAYS('2022-01-17')) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0 );
	//删除分区
	ALTER TABLE `表名` DROP PARTITION 20220115,20220116; */
	/**--------查询分区 开始--------**/
	partitionSelSql := `SELECT MAX(PARTITION_NAME) AS maxPartitionName FROM INFORMATION_SCHEMA.PARTITIONS WHERE TABLE_SCHEMA = SCHEMA() AND TABLE_NAME = '` + dbTable + `'`
	maxPartitionNameTmp, err := g.DB(dbGroup).GetValue(ctx, partitionSelSql) //不是分区表或无分区，查询结果都会有一项，且第一项maxPartitionName值为null
	if err != nil {
		return
	}
	maxPartitionName := maxPartitionNameTmp.String()
	/**--------查询分区 结束--------**/

	/**--------无分区则建立当前分区 开始--------**/
	if maxPartitionName == `` {
		ltTime := gtime.Now().Unix() + partitionTime
		ltDate := gtime.New(ltTime).Format(`Y-m-d`)
		partitionName := gtime.New(ltTime - 24*60*60).Format(`Ymd`) //该分区的最大日期作为分区名
		partitionCreateSql := "PARTITION `" + partitionName + "` VALUES LESS THAN (TO_DAYS( '" + ltDate + "' )) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0"
		partitionCreateSql = "ALTER TABLE `" + dbTable + "` PARTITION BY RANGE (TO_DAYS( " + partitionField + " )) ( " + partitionCreateSql + " )"
		_, err = g.DB(dbGroup).Exec(ctx, partitionCreateSql)
		if err != nil {
			return
		}
		maxPartitionName = partitionName
	}
	/**--------无分区则建立当前分区 结束--------**/

	/**--------检测需要创建的分区是否存在，没有则新增分区 开始--------**/
	partitionAddSqlList := []string{}
	nowTime := gtime.Now().Unix()
	maxPartitionTimeTmp, _ := gtime.StrToTimeFormat(maxPartitionName, `Ymd`)
	maxPartitionTime := maxPartitionTimeTmp.Unix()
	var i int64
	for i = 0; i < partitionNumber; i++ {
		//时间超过最大的分区后才开始新增分区，且以最大分区的时间开始往后计算
		if gtime.New(nowTime+(i+1)*partitionTime-24*60*60).Format(`Ymd`) > maxPartitionName {
			ltTime := maxPartitionTime + (i+1)*partitionTime
			ltDate := gtime.New(ltTime).Format(`Y-m-d`)
			partitionName := gtime.New(ltTime - 24*60*60).Format(`Ymd`) //该分区的最大日期作为分区名
			partitionAddSqlList = append(partitionAddSqlList, "PARTITION `"+partitionName+"` VALUES LESS THAN (TO_DAYS( '"+ltDate+"' )) ENGINE = InnoDB MAX_ROWS = 0 MIN_ROWS = 0")
		}

	}
	if len(partitionAddSqlList) > 0 {
		partitionAddSql := gstr.Join(partitionAddSqlList, `,`)
		partitionAddSql = "ALTER TABLE `" + dbTable + "` ADD PARTITION ( " + partitionAddSql + " )"
		_, err = g.DB(dbGroup).Exec(ctx, partitionAddSql)
		if err != nil {
			return
		}
	}
	/**--------检测需要创建的分区是否存在，没有则新增分区 结束--------**/
	return
}
