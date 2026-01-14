package initialize

import (
	"context"
	"net/http"
	_ "net/http/pprof"
)

func initPprof(ctx context.Context) {
	go func() {
		http.ListenAndServe(`:6060`, nil)
	}()
}
