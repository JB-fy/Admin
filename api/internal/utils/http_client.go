package utils

import (
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/util/gconv"
)

type HttpClient struct {
	*gclient.Client
	config HttpClientConfig
}

type HttpClientConfig struct {
	Timeout      time.Duration     `json:"timeout"`
	Header       map[string]string `json:"header"`
	ProxyUrl     string            `json:"proxy_url"`
	IsFileUpload bool              `json:"is_file_upload"`
}

var (
	httpClientMap   = map[string]*HttpClient{} //存放不同配置实例。因初始化只有一次，故重要的是读性能，普通map比sync.Map的读性能好
	httpClientMuMap sync.Map
)

func NewHttpClient(ctx context.Context, configOpt ...HttpClientConfig) (obj *HttpClient) {
	config := HttpClientConfig{}
	if len(configOpt) > 0 {
		config = configOpt[0]
	}
	key := gmd5.MustEncrypt(config)
	ok := false
	if obj, ok = httpClientMap[key]; ok { //先读一次（不加锁）
		return
	}
	muTmp, _ := httpClientMuMap.LoadOrStore(key, &sync.Mutex{})
	mu := muTmp.(*sync.Mutex)
	mu.Lock()
	defer func() {
		mu.Unlock()
		httpClientMuMap.Delete(key)
	}()
	if obj, ok = httpClientMap[key]; ok { // 再读一次（加锁），防止重复初始化
		return
	}
	client := g.Client()
	client.SetTimeout(config.Timeout)
	client.SetHeaderMap(config.Header)
	client.SetProxy(config.ProxyUrl)
	if config.IsFileUpload {
		client.Use(func(c *gclient.Client, r *http.Request) (resp *gclient.Response, err error) {
			query := r.URL.Query()
			if formDataContentType := query.Get(`formDataContentType`); formDataContentType != `` {
				r.Header.Set(`Content-Type`, formDataContentType)
				query.Del(`formDataContentType`)
				r.URL.RawQuery = query.Encode()
			}
			resp, err = c.Next(r)
			return
		})
	}
	obj = &HttpClient{
		Client: client,
		config: config,
	}
	httpClientMap[key] = obj
	return
}

func (clientThis *HttpClient) PostFile(ctx context.Context, httpUrl string, param map[string]any, fileBytes []byte, fileFieldName string, fileName string) (res *gclient.Response, err error) {
	buf := BytesBufferPoolGet()
	defer BytesBufferPoolPut(buf)
	writer := multipart.NewWriter(buf)
	for k, v := range gconv.MapStrStr(param) {
		err = writer.WriteField(k, v)
		if err != nil {
			return
		}
	}
	var file io.Writer
	file, err = writer.CreateFormFile(fileFieldName, fileName)
	if err != nil {
		return
	}
	_, err = file.Write(fileBytes)
	if err != nil {
		return
	}
	err = writer.Close()
	if err != nil {
		return
	}

	if clientThis.config.IsFileUpload {
		// 将writer.FormDataContentType()当作query传递，在拦截器中识别并设置Content-Type头。减少频繁复制客户端
		parseUrl, _ := url.Parse(httpUrl)
		parseUrlQuery := parseUrl.Query()
		parseUrlQuery.Set(`formDataContentType`, writer.FormDataContentType())
		parseUrl.RawQuery = parseUrlQuery.Encode()
		res, err = clientThis.Post(ctx, parseUrl.String(), buf)
		return
	}
	// ContentType方法会复制一个新客户端，而每次上传只是Content-Type头不一样而已，没必要每次都复制一个新客户端
	res, err = clientThis.ContentType(writer.FormDataContentType()).Post(ctx, httpUrl, buf)
	return
}

func (clientThis *HttpClient) IsExistFile(ctx context.Context, httpUrl string) (isExist bool, err error) {
	res, err := clientThis.Head(ctx, httpUrl)
	if err != nil {
		return
	}
	defer res.Close()
	isExist = res.StatusCode == http.StatusOK
	return
}
