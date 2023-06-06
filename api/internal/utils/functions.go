package utils

import (
	"api/internal/consts"
	"context"
	"crypto/md5"
	"fmt"
	"math/rand"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

func NewErrorCode(ctx context.Context, code int, msg string, data ...map[string]interface{}) error {
	if msg == "" {
		msg = g.I18n().Tf(ctx, gconv.String(code))
	}
	dataTmp := map[string]interface{}{}
	if len(data) > 0 && data[0] != nil {
		dataTmp = data[0]
	}
	return gerror.NewCode(gcode.New(code, "", dataTmp), msg)
}

func HttpFailJson(r *ghttp.Request, err error) {
	resData := map[string]interface{}{
		"code": 99999999,
		"msg":  err.Error(),
		"data": g.I18n().Tf(r.GetCtx(), "99999999"),
	}
	/* _, ok := err.(*gerror.Error)
	if ok { */
	code := gerror.Code(err)
	if code.Code() > 0 {
		resData["code"] = code.Code()
		resData["data"] = code.Detail()
	}
	r.Response.WriteJson(resData)
}

func HttpSuccessJson(r *ghttp.Request, data map[string]interface{}, code int, msg ...string) (err error) {
	resData := map[string]interface{}{
		"code": code,
		"msg":  "",
		"data": data,
	}
	if len(msg) == 0 || msg[0] == "" {
		resData["msg"] = g.I18n().Tf(r.GetCtx(), gconv.String(code))
	} else {
		resData["msg"] = msg[0]
	}
	r.Response.WriteJson(resData)
	return
}

func RandomStr(length int) string {
	/* var ConstStrArr = []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k",
		"l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v",
		"w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G",
		"H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R",
		"S", "T", "U", "V", "W", "X", "Y", "Z", "0", "1", "2",
		"3", "4", "5", "6", "7", "8", "9",
	} */
	str := ""
	max := len(consts.ConstStrArr)
	for i := 0; i < length; i++ {
		str += consts.ConstStrArr[rand.Intn(max)]
	}
	return str
}

func Md5(rawStr string) string {
	// h := md5.New()
	// h.Write([]byte(rawStr))
	// md5str := fmt.Printf("%x", h.Sum(nil))
	h := md5.Sum([]byte(rawStr))
	md5str := fmt.Sprintf("%x", h) //将[]byte转成16进制
	return md5str
}
