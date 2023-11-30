package controller

import (
	"api/api"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type Test struct{}

func NewTest() *Test {
	return &Test{}
}

func (c *Test) Test(ctx context.Context, req *api.TestReq) (res *api.TestRes, err error) {
	// time.Sleep(10 * time.Second) // 睡眠几秒
	// ghttp.RestartAllServer(ctx) // 重启服务

	/*--------数据库使用示例 开始--------*/
	// g.DB(`default`).Model(`xxxx_txxx`).Safe().Ctx(ctx) // 数据库连接
	// m = m.Where(m.Builder().Where().WhereOr())         // 复杂条件
	// list, err := dao.NewDaoHandler(ctx, &daoXxxx.Txxx).Filter(g.Map{&daoXxxx.Txxx.Columns().Xxxx: `xxxx`}).Field(append(&daoXxxx.Txxx.ColumnArr(), `aaaa`)).JoinGroupByPrimaryKey().GetModel().All() // dao常用方式
	/* // 数据库事务
	err = daoXxxx.Txxx.ParseDbCtx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		// _, err = tx.Model(daoXxxx.Txxx.ParseDbTable(ctx)).Data(g.Map{`Xxxx`: `xxxx`}).Update()
		// id, err = tx.Model(daoXxxx.Txxx.ParseDbTable(ctx)).Handler(daoXxxx.Txxx.ParseInsert(data)).InsertAndGetId()
		return
	}) */
	/*--------数据库使用示例 结束--------*/

	/*--------数据操作示例 开始--------*/
	// g.Validator().Rules(`required|integer`).Data(`aaaa`).Run(ctx) // 单独验证

	// garray.NewStrArrayFrom([]string{`a`, `b`, `c`}).Contains(`a`) // 是否含有元素

	/* // Map并集
	// gmap.NewStrAnyMapFrom(g.Map{`a`: 1, `b`: 2}).Merge(gmap.NewStrAnyMapFrom(g.Map{`a`: 4, `c`: 3})).Map() // 这样直接写会报错。需要分多个步骤编写
	rawMap := gmap.NewStrAnyMapFrom(g.Map{`a`: 1, `b`: 2})
	rawMap.Merge(gmap.NewStrAnyMapFrom(g.Map{`a`: 4, `c`: 3}))
	rawMap.Map() */

	// gset.NewIntSetFrom([]int{1, 2, 3}).Intersect(gset.NewIntSetFrom([]int{1, 3})).Slice() // 交集
	// gset.NewIntSetFrom([]int{1, 2, 3}).Diff(gset.NewIntSetFrom([]int{1, 3})).Slice()      // 差集
	// gset.NewIntSetFrom([]int{1, 2, 3}).Union(gset.NewIntSetFrom([]int{1, 3})).Slice()     // 并集
	// gset.NewIntSetFrom([]int{1, 2, 3}).Merge(gset.NewIntSetFrom([]int{1, 3})).Slice()     // 合并，也是并集
	/*--------数据操作示例 结束--------*/

	/*--------配置使用示例 开始--------*/
	// genv.Set(`X_X`, `xx`)                        // key必须由大写和_组成
	// g.Cfg().MustGetWithEnv(ctx, `X_X`)           // X_X或x_x或x.x方法都可以读取到
	// g.Cfg().MustGet(ctx, `superPlatformAdminId`) // 获取配置参数
	/*--------配置使用示例 结束--------*/

	/*--------日志使用示例 开始--------*/
	// g.Log().Info(ctx, `日志打印`, "\r\n", g.Map{`aaaa`: `asd`}) // 记录日志
	// g.I18n().T(ctx, `code.99999999`)                        // 多语言
	/*--------日志使用示例 结束--------*/

	/*--------函数结果示例 开始--------*/
	// g.RequestFromCtx(ctx).GetClientIp() // 192.168.2.44。常用于WEB编程，如HTTP
	// g.RequestFromCtx(ctx).GetRemoteIp() // 192.168.2.44。常用于网络编程，如TCP或UDP

	// g.RequestFromCtx(ctx).GetUrl()         // http://192.168.2.200:20080/testMeta?a=1&b=2
	// g.RequestFromCtx(ctx).GetHost()        // 192.168.2.200
	// g.RequestFromCtx(ctx).Host             // 192.168.2.200:20080
	// g.RequestFromCtx(ctx).RequestURI       // /testMeta?a=1&b=2
	// g.RequestFromCtx(ctx).URL.String()     // /testMeta?a=1&b=2
	// g.RequestFromCtx(ctx).URL              // /testMeta?a=1&b=2
	// g.RequestFromCtx(ctx).URL.Path         // /testMeta
	// g.RequestFromCtx(ctx).URL.RawQuery     // a=1&b=2
	// g.RequestFromCtx(ctx).URL.RequestURI() // /testMeta?a=1&b=2
	// g.RequestFromCtx(ctx).Router           // &{/testMeta GET default ^/testMeta$ [] 0}
	// g.RequestFromCtx(ctx).Router.Uri       // /testMeta

	// grand.Str(`abcdefg0123456789`, 8) // 指定字符串随机
	// grand.N(1000, 9999)               // 1000~9999
	// grand.Intn(1000)                  // 0~999
	// grand.Digits(8)                   // 0123456789
	// grand.Letters(8)                  // abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ
	// grand.S(8)                        // abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789
	// grand.Symbols(8)                  // !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~
	/*--------函数结果示例 结束--------*/

	/*--------MYSQL数据库常用方法 开始--------*/
	// CASE 字段 WHEN 匹配 THEN 值 ELSE 默认值 END                                    // 当字段匹配时返回对应值，否则返回默认值
	// IF(条件, 值1, 值2)                                                             // 当条件为真，返回值1,否则返回值2
	// IFNULL(值1, 值2)                                                               // 当值1为null时，返回值2
	// NULLIF(值1, 值2)                                                               // 当值1等于值2，返回null，否则返回值1

	// AVG([DISTINCT] 字段)                                                           // 平均值
	// ROUND(值, 2)                                                                   // 保留两位小数
	// CONCAT(字符串,...)                                                             // 拼接字符串
	// CONCAT_WS(分隔符, 字符串,...)                                                  // 拼接字符串。可指定分隔符
	// GROUP_CONCAT([DISTINCT] 字段 [ORDER BY 排序字段 ASC/DESC] [SEPARATOR 分隔符])  // 拼接字符串。一般在GROUP BY语句中使用
	// REPLACE(字符串, ',', '')                                                       // 替换字符串。
	// LENGTH(字符串)                                                                 // 字符串长度。中文根据编码不同，算多个字符
	// CHAR_LENGTH(字符串)                                                            // 字符串长度。中英文都算1个字符

	// UNIX_TIMESTAMP() 或 UNIX_TIMESTAMP('2006-01-02 15:04:05')                      // 时间戳。示例：1136185445
	// FROM_UNIXTIME(1136185445, '%Y-%m-%d %H:%i:%s')                                 // 时间戳转换成指定格式
	// DATE_FORMAT('2006-01-02 15:04:05', '%Y-%m-%d')                                 // 日期格式转换成指定格式
	// STR_TO_DATE('January 02 2016', '%M %d %Y')                                     // 根据指定格式将字符串转变成日期格式。示例：2006-01-02或2006-01-02 15:04:05
	// NOW() 或 UTC_TIMESTAMP()                                                       // 当前日期和时间。示例：2006-01-02 15:04:05
	// CURDATE() 或 UTC_DATE()                                                        // 当前日期。示例：2006-01-02
	// CURTIME() 或 UTC_TIME()                                                        // 当前时间。示例：15:04:05
	// YEAR('2006-01-02 15:04:05')                                                    // 年
	// MONTH('2006-01-02 15:04:05')                                                   // 月
	// DAY('2006-01-02 15:04:05')                                                     // 日
	// HOUR('2006-01-02 15:04:05')                                                    // 时
	// MINUTE('2006-01-02 15:04:05')                                                  // 分
	// SECOND('2006-01-02 15:04:05')                                                  // 秒
	// WEEK('2006-01-02 15:04:05')                                                    // 周。范围0~53
	// WEEKDAY('2006-01-02 15:04:05')                                                 // 周几。0星期一
	// LAST_DAY('2006-01-02 15:04:05')                                                // 返回当前日期月份的最后一天。示例：2006-01-31
	// DATE('2006-01-02 15:04:05')                                                    // 日期。示例：2006-01-02
	// DATE_SUB('2006-01-02 15:04:05', INTERVAL 7 type)                               // 该日期的多少type前。type：YEAR MONTH DAY HOUR MINUTE SECOND
	// DATE_ADD('2006-01-02 15:04:05', INTERVAL 7 type)                               // 该日期的多少type后。type：YEAR MONTH DAY HOUR MINUTE SECOND
	// DATEDIFF('2006-01-03', '2006-01-02')                                           // 相隔天数。示例：1
	// TIMEDIFF('16:05:06', '15:04:05')                                               // 相隔时间。示例：01:01:01

	// POINT(118.585519, 24.914168)                                                   // 字段类型为point时使用
	// ST_X(POINT(118.585519, 24.914168))                                             // 经度，Mysql5中使用X()
	// ST_Y(POINT(118.585519, 24.914168))                                             // 纬度，Mysql5中使用Y()
	// ST_DISTANCE_SPHERE(POINT(118.585519, 24.914168), POINT(118.585519, 24.914168)) // 计算经纬度距离
	/*--------MYSQL数据库常用方法 结束--------*/

	// g.RequestFromCtx(ctx).Response.Status = http.StatusMultipleChoices
	// err = utils.NewErrorCode(ctx, 99999999, ``)
	/* utils.HttpWriteJson(ctx, map[string]interface{}{
		`info`: map[string]interface{}{},
	}, 0, ``) */
	res = &api.TestRes{}
	/* info, _ := g.DB().Model(`auth_scene`).Ctx(ctx).One()
	info.Struct(&res.Info) */
	return
}

func (c *Test) Test1(r *ghttp.Request) {
	/* var req *api.TestReq
	err := r.Parse(&req)
	if err != nil {
		r.Response.Writeln(err.Error())
		return
	} */

	// r.SetError(utils.NewErrorCode(r.GetCtx(), 99999999, ``))
	r.Response.WriteJson(map[string]interface{}{
		`code`: 0,
		`msg`:  g.I18n().T(r.GetCtx(), `code.0`),
		`data`: map[string]interface{}{
			`list`: []map[string]interface{}{},
		},
	})
}
