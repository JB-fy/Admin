package controller

import (
	"api/api"
	// daoAuth "api/internal/dao/auth"
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

	/* //生成登录token（测试用）
	claims := utils.CustomClaims{LoginId: 2}
	sceneConfig, _ := daoAuth.Scene.ParseDbCtx(ctx).Where(daoAuth.Scene.Columns().SceneCode, `sceneCode`).Value(daoAuth.Scene.Columns().SceneConfig)
	jwt := utils.NewJWT(ctx, sceneConfig.Map())
	fmt.Println(jwt.CreateToken(claims)
	fmt.Println(token) */

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
