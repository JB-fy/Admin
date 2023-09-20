package router

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
)

func InitRouterWebSocket(s *ghttp.Server) {
	s.BindHandler(`/ws`, func(r *ghttp.Request) {
		ws, err := r.WebSocket()
		if err != nil {
			glog.Error(r.GetCtx(), err)
			r.Exit()
		}
		for {
			msgType, msg, err := ws.ReadMessage()
			if err != nil {
				return
			}
			err = ws.WriteMessage(msgType, msg)
			if err != nil {
				return
			}
		}
	})
}
