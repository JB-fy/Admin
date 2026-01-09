package sign

/* import (
	daoAuth "api/internal/dao/auth"
	"api/internal/utils"
	"api/internal/utils/sign/model"
	"context"
	"errors"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type HandlerOfTemplate struct {
	Ctx    context.Context
	sign   model.Sign
	client *utils.HttpClient
}

func NewHandlerOfTemplate(ctx context.Context, sceneIdOpt ...string) *HandlerOfTemplate {
	handlerObj := &HandlerOfTemplate{Ctx: ctx, client: utils.NewHttpClient(ctx, utils.HttpClientConfig{Timeout: 30 * time.Second})}
	var sceneInfo gdb.Record
	if len(sceneIdOpt) == 0 {
		sceneInfo = jbctx.GetCtxSceneInfo(ctx)
	} else {
		sceneInfo, _ = daoAuth.Scene.CacheGetInfo(ctx, sceneIdOpt[0])
	}
	config, _ := sceneInfo[daoAuth.Scene.Columns().SceneConfig].Map()[`sign_config`].(g.Map)
	gconv.Struct(config, handlerObj)
	handlerObj.sign = NewSign(ctx, gconv.Uint8(config[`sign_type`]), config)
	return handlerObj
}

func (handlerThis *HandlerOfTemplate) Create(data map[string]any) (sign string) {
	sign = handlerThis.sign.Create(handlerThis.Ctx, data)
	return
}

func (handlerThis *HandlerOfTemplate) Verify(r *ghttp.Request) (err error) {
	if diffSec := time.Since(r.Get(`ts`).Time()).Seconds(); diffSec < -10 || diffSec > 10 { //误差超过多少秒会报错
		err = errors.New(`时间戳失效`)
		return
	}
	data := r.GetMap()
	delete(data, `sign`)
	err = handlerThis.sign.Verify(handlerThis.Ctx, data, r.Get(`sign`).String())
	return
}

func (handlerThis *HandlerOfTemplate) getReqData(param g.Map) (reqData g.Map) {
	reqData = param
	if reqData == nil {
		reqData = g.Map{}
	}
	reqData[`ts`] = gtime.Now().Unix()
	reqData[`sign`] = handlerThis.Create(reqData)
	return
}

func (handlerThis *HandlerOfTemplate) getResData(res *gclient.Response) (resData *gjson.Json, err error) {
	resStr := res.ReadAllString()
	resData = gjson.New(resStr)
	if !resData.Contains(`code`) {
		err = errors.New(resStr)
		return
	}
	if resData.Get(`code`).Int() != 0 {
		err = errors.New(resData.Get(`msg`).String())
		return
	}
	return
}

func (handlerThis *HandlerOfTemplate) Info(proxyUrl string, id string) (info g.Map, err error) {
	res, err := handlerThis.client.Post(handlerThis.Ctx, proxyUrl, handlerThis.getReqData(g.Map{
		`id`: id,
	}))
	if err != nil {
		return
	}
	defer res.Close()
	resData, err := handlerThis.getResData(res)
	if err != nil {
		return
	}
	info = resData.Get(`data.info`).Map()
	return
} */
