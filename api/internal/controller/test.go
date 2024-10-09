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
	/* // HookSelect方法内启用协程且进行多次Value()方法调用或内部调用过Value()方法的Sum()等方法，有概率数据返回为空（当Limit越多时出现概率越高）
	daoAuth.Action.Ctx(ctx).Hook(gdb.HookHandler{
		Select: func(ctx context.Context, in *gdb.HookSelectInput) (result gdb.Result, err error) {
			result, err = in.Next(ctx)
			if err != nil {
				return
			}
			var wg sync.WaitGroup
			for _, record := range result {
				wg.Add(1)
				go func(record gdb.Record) {
					defer wg.Done()
					fmt.Println(daoAuth.Action.Ctx(ctx).Where(daoAuth.Action.Columns().ActionId, record[daoAuth.Action.Columns().ActionId]).Value(daoAuth.Action.Columns().ActionId))
					fmt.Println(daoAuth.Action.Ctx(ctx).Where(daoAuth.Action.Columns().ActionId, record[daoAuth.Action.Columns().ActionId]).Value(daoAuth.Action.Columns().ActionName))
				}(record)
			}
			wg.Wait()
			return
		},
	}).Limit(20).All() */

	// time.Sleep(10 * time.Second) // 睡眠几秒
	// ghttp.RestartAllServer(ctx)  // 重启服务

	/* //生成登录token（测试用）
	sceneInfo, _ := daoAuth.Scene.CtxDaoModel(ctx).Filter(daoAuth.Scene.Columns().SceneCode, `platform`).One()
	tokenInfo := token.TokenInfo{LoginId: `1`}
	token, err := token.NewHandler(ctx, sceneInfo[daoAuth.Scene.Columns().SceneConfig].Map(), sceneInfo[daoAuth.Scene.Columns().SceneCode].String()).Create(tokenInfo)
	fmt.Println(token) */

	/*--------数据库使用示例 开始--------*/
	// gregex.IsMatchString(`1062.*Duplicate.*`, err.Error()) //判断错误是不是唯一索引已存在
	// m = m.Where(m.Builder().Where(`xxxx`).WhereOr(`xxxx`)) // 复杂条件
	/* // 数据库事务
	xxxxTxxxDaoModel := daoXxxx.Txxx.CtxDaoModel(ctx)
	err = xxxxTxxxDaoModel.Transaction(func(ctx context.Context, tx gdb.TX) (err error) {
		list, err := daoXxxx.Txxx.CtxDaoModel(ctx).TX(tx).Filter(`filterKey`, `xxxx`).Field(`xxxx`).LockUpdate().All()                       //查询
		id, err := xxxxTxxxDaoModel.CloneNew().TX(tx).HookInsert(g.Map{`dataKey`: `xxxx`}).InsertAndGetId()                                  //新增
		row, err := xxxxTxxxDaoModel.CloneNew().TX(tx).SetIdArr(g.Map{`id`: id}).HookUpdate(g.Map{`dataKey`: `xxxx`}).UpdateAndGetAffected() //修改
		row, err := xxxxTxxxDaoModel.CloneNew().TX(tx).SetIdArr(g.Map{`filterKey`: `xxxx`}).HookDelete().DeleteAndGetAffected()              //删除
		// _, err = tx.Model(xxxxTxxxDaoModel.DbTable).Data(g.Map{`dataKey`: `xxxx`}).Update()                                                  //不建议用
		return
	}) */
	/*--------数据库使用示例 结束--------*/

	/*--------数据操作示例 开始--------*/
	// g.Validator().Rules(`required|length:1,10|regex:^[\p{L}\p{M}\p{N}_-]+$`).Messages(`必须|最长10个字|昵称不允许特殊字符`).Data(`aaaa`).Run(ctx) // 单独验证

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
	// genv.Set(`X_X`, `xx`)              // key必须由大写和_组成
	// g.Cfg().MustGetWithEnv(ctx, `X_X`) // X_X或x_x或x.x方法都可以读取到
	// g.Cfg().MustGet(ctx, `dev`)        // 获取配置参数
	/*--------配置使用示例 结束--------*/

	/*--------日志使用示例 开始--------*/
	// g.Log().Info(ctx, `日志打印`, "\r\n", g.Map{`aaaa`: `asd`}) // 记录日志
	// g.I18n().T(ctx, `code.99999999`)                        // 多语言
	/*--------日志使用示例 结束--------*/

	/*--------函数结果示例 开始--------*/
	// g.RequestFromCtx(ctx).GetClientIp() // 192.168.0.44。常用于WEB编程，如HTTP
	// g.RequestFromCtx(ctx).GetRemoteIp() // 192.168.0.44。常用于网络编程，如TCP或UDP

	// g.RequestFromCtx(ctx).GetUrl()         // http://192.168.0.200:20080/testMeta?a=1&b=2
	// g.RequestFromCtx(ctx).GetHost()        // 192.168.0.200
	// g.RequestFromCtx(ctx).Host             // 192.168.0.200:20080
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

	/*--------常用协程示例 开始--------*/
	/* //阻塞，等待协程完成。
	mp := sync.Map{}
	var wg sync.WaitGroup
	var mx sync.Mutex
	total := 0
	listRaw := []map[string]any{}
	for k, v := range listRaw {
		wg.Add(1)
		go func(ctx context.Context, key int, value map[string]any) {
			defer wg.Done()
			mp.Store(key, value)

			mx.Lock()
			total += key
			mx.Unlock()
		}(ctx, k, v)
	}
	wg.Wait()
	listResult := []map[string]any{}
	mp.Range(func(key, value any) bool {
		listResult = append(listResult, value.(map[string]any))
		return true
	}) */

	/* //不阻塞，直接返回响应数据。注意：gctx.NeverDone(ctx)必须有
	go func(ctx context.Context) {
		// ctx := g.RequestFromCtx(ctx).GetNeverDoneCtx()
		utils.GetCtxSceneInfo(ctx)
	}(gctx.NeverDone(ctx)) */
	/*--------常用协程示例 结束--------*/

	// g.RequestFromCtx(ctx).Response.Status = http.StatusMultipleChoices
	// err = utils.NewErrorCode(ctx, 99999999, ``)
	// err = errors.New(``)	//与上面返回结果一样。在api/internal/middleware/handler_response.go中会统一处理成99999999错误码
	/* utils.HttpWriteJson(ctx, map[string]any{
		`info`: map[string]any{},
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
	r.Response.WriteJson(map[string]any{
		`code`: 0,
		`msg`:  g.I18n().T(r.GetCtx(), `code.0`),
		`data`: map[string]any{
			`list`: []map[string]any{},
		},
	})
}
