package utils

import (
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	httpClient         = g.Client()
	httpClientOfUpload = g.Client().Use(func(c *gclient.Client, r *http.Request) (resp *gclient.Response, err error) {
		query := r.URL.Query()
		if formDataContentType := query.Get(`formDataContentType`); formDataContentType != `` {
			r.Header.Set(`Content-Type`, formDataContentType)
			query.Del(`formDataContentType`)
			r.URL.RawQuery = query.Encode()
		}
		resp, err = c.Next(r)
		return
	})
)

func HttpClient() *gclient.Client {
	return httpClient
}

func HttpClientOfUpload() *gclient.Client {
	return httpClientOfUpload
}

func PostFile(ctx context.Context, httpUrl string, param map[string]any, fileBytes []byte, fileFieldName string, fileName string) (res *gclient.Response, err error) {
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

	// ContentType方法会从ClientOfUpload复制一个新客户端，而每次上传只是Content-Type头不一样而已，没必要每次都复制一个新客户端
	// res, err := ClientOfUpload.ContentType(writer.FormDataContentType()).Post(ctx, httpUrl, buf)
	// 将writer.FormDataContentType()当作query传递，在拦截器中识别并设置Content-Type头。减少频繁复制ClientOfUpload客户端
	parseUrl, _ := url.Parse(httpUrl)
	parseUrlQuery := parseUrl.Query()
	parseUrlQuery.Set(`formDataContentType`, writer.FormDataContentType())
	parseUrl.RawQuery = parseUrlQuery.Encode()
	res, err = HttpClientOfUpload().Post(ctx, parseUrl.String(), buf)
	return
}

func IsExistFile(ctx context.Context, httpUrl string) (isExist bool, err error) {
	res, err := HttpClient().Head(ctx, httpUrl)
	if err != nil {
		return
	}
	defer res.Close()
	isExist = res.StatusCode == http.StatusOK
	return
}
