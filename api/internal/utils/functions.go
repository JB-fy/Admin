package utils

import (
	"crypto/md5"
	"fmt"
	"math/rand"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

func HttpFailJson(r *ghttp.Request, code int, msg string, data map[string]interface{}) {
	resData := map[string]interface{}{
		"code": code,
		"data": data,
	}
	if msg == "" {
		resData["msg"] = g.I18n().Tf(r.GetCtx(), gconv.String(code))
	} else {
		resData["msg"] = msg
	}
	r.Response.WriteJson(resData)
}

func HttpSuccessJson(r *ghttp.Request, data map[string]interface{}, code int, msg string) {
	resData := map[string]interface{}{
		"code": code,
		"data": data,
	}
	if msg == "" {
		resData["msg"] = g.I18n().Tf(r.GetCtx(), gconv.String(code))
	} else {
		resData["msg"] = msg
	}
	r.Response.WriteJson(resData)
}

func RandomStr(length int) string {
	var strs = []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k",
		"l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v",
		"w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G",
		"H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R",
		"S", "T", "U", "V", "W", "X", "Y", "Z", "0", "1", "2",
		"3", "4", "5", "6", "7", "8", "9",
	}
	str := ""
	max := len(strs)
	for i := 0; i < length; i++ {
		str += strs[rand.Intn(max)]
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
