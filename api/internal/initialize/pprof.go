package initialize

import (
	"context"
	"net/http"
	_ "net/http/pprof"
)

func initPprof(ctx context.Context) {
	/* gtimer.SetInterval(ctx, 5*time.Second, func(ctx context.Context) {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		if m.HeapInuse > 4*1024*1024*1024 { //内存超过多少时，记录pprof
			buf := utils.BytesBufferPoolGet()
			defer utils.BytesBufferPoolPut(buf)
			err := pprof.WriteHeapProfile(buf)
			if err != nil {
				return
			}
			gfile.PutBytes(gfile.SelfDir()+`/log/pprof/heap.pprof`, buf.Bytes())
		}
	}) */
	go func() {
		http.ListenAndServe(`:6060`, nil)
	}()
}
