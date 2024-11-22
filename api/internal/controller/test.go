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
	// 数据库中作为配置表的数据，可根据自身需要使用以下方式读取数据
	// daoAuth.Scene.CtxDaoModel(ctx).FilterPri(`platform`).One()                 //不推荐：数据库实时读取
	// cache.DbData.GetOrSet(ctx, &daoAuth.Scene, `platform`, 30*60)              //推荐-数据修改需要立即同步缓存的表：缓存在redis，数据修改可做到立即同步缓存，只需数据修改时删除缓存即可。如在dao层文件的后置更新和后置删除中补充缓存删除代码：cache.DbData.Del(ctx, &daoAuth.Scene, `platform`)
	// cache.DbDataLocal.GetInfo(ctx, daoAuth.Scene.CtxDaoModel(ctx), `platform`) //推荐-数据修改无需立即同步缓存的表：服务启动时，缓存在本机内存中，数据库修改，只能等待各个服务器定时器触发才能同步缓存

	// time.Sleep(10 * time.Second) // 睡眠几秒
	// ghttp.RestartAllServer(ctx)  // 重启服务

	/* //生成登录token（测试用）
	token, err := token.NewHandler(ctx, `platform`).Create(`1`)
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
	// g.Cfg().MustGet(ctx, `logger.http.isRecord`) // 获取配置参数
	/*--------配置使用示例 结束--------*/

	/*--------日志使用示例 开始--------*/
	// g.Log().Info(ctx, `日志打印`, "\r\n", g.Map{`aaaa`: `asd`}) // 记录日志
	// g.I18n().T(ctx, `code.99999999`)                        // 多语言
	/*--------日志使用示例 结束--------*/

	/*--------函数结果示例 开始--------*/
	// g.RequestFromCtx(ctx).GetClientIp() // xxx.xxx.xxx.xxx。常用于WEB编程，如HTTP
	// g.RequestFromCtx(ctx).GetRemoteIp() // xxx.xxx.xxx.xxx。常用于网络编程，如TCP或UDP

	// g.RequestFromCtx(ctx).GetUrl()         // http://jb.admin.com:20080/testMeta?a=1&b=2
	// g.RequestFromCtx(ctx).GetHost()        // jb.admin.com
	// g.RequestFromCtx(ctx).Host             // jb.admin.com:20080
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

	/*--------管道示例 开始--------*/
	/* var ch = make(chan g.Map)
	go func() {
		for {
			item := <-ch
			gutil.Dump(item)
			time.Sleep(3 * time.Second)
		}
	}()
	ch <- g.Map{} */
	/*--------管道示例 结束--------*/

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
