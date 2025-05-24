package sign

// 根据自身需要修改
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

type Handler struct {
	Ctx    context.Context
	sign   model.Sign
	client *gclient.Client
}

var client = g.Client().SetTimeout(30 * time.Second)

func NewHandler(ctx context.Context, sceneIdOpt ...string) *Handler {
	handlerObj := &Handler{Ctx: ctx, client: client}
	var sceneInfo gdb.Record
	if len(sceneIdOpt) == 0 {
		sceneInfo = utils.GetCtxSceneInfo(ctx)
	} else {
		sceneInfo, _ = daoAuth.Scene.CacheGetInfo(ctx, sceneIdOpt[0])
	}
	config, _ := sceneInfo[daoAuth.Scene.Columns().SceneConfig].Map()[`sign_config`].(g.Map)
	gconv.Struct(config, handlerObj)
	handlerObj.sign = NewSign(ctx, gconv.Uint8(config[`sign_type`]), config)
	return handlerObj
}

func (handlerThis *Handler) Create(data map[string]any) (sign string) {
	sign = handlerThis.sign.Create(handlerThis.Ctx, data)
	return
}

func (handlerThis *Handler) Verify(r *ghttp.Request) (err error) {
	if diffSec := gtime.Now().Sub(r.Get(`ts`).GTime()).Seconds(); diffSec < -5 || diffSec > 30 { //误差超过多少秒会报错
		err = errors.New(`时间戳失效`)
		return
	}
	data := r.GetMap()
	delete(data, `sign`)
	err = handlerThis.sign.Verify(handlerThis.Ctx, data, r.Get(`sign`).String())
	return
}

func (handlerThis *Handler) getReqData(param g.Map) (reqData g.Map) {
	reqData = param
	if reqData == nil {
		reqData = g.Map{}
	}
	reqData[`ts`] = gtime.Now().Unix()
	reqData[`sign`] = handlerThis.Create(reqData)
	return
}

func (handlerThis *Handler) getResData(res *gclient.Response) (resData *gjson.Json, err error) {
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

func (handlerThis *Handler) Info(proxyUrl string, id string) (info g.Map, err error) {
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
