package initialize

//时区设置
import "github.com/gogf/gf/v2/os/gtime"

func init() {
	gtime.SetTimeZone(`Asia/Shanghai`)
}
