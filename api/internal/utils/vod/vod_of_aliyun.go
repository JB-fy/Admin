package vod

import (
	"api/internal/utils"
	"context"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	sts20150401 "github.com/alibabacloud-go/sts-20150401/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/gogf/gf/v2/util/gconv"
)

type VodOfAliyun struct {
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	Endpoint        string `json:"endpoint"`
	RoleArn         string `json:"roleArn"`
	client          *sts20150401.Client
}

func NewVodOfAliyun(ctx context.Context, config map[string]any) *VodOfAliyun {
	obj := &VodOfAliyun{}
	gconv.Struct(config, obj)
	if obj.AccessKeyId == `` || obj.AccessKeySecret == `` || obj.Endpoint == `` || obj.RoleArn == `` {
		panic(`缺少插件配置：视频点播-阿里云`)
	}
	var err error
	obj.client, err = sts20150401.NewClient(&openapi.Config{
		AccessKeyId:     tea.String(obj.AccessKeyId),
		AccessKeySecret: tea.String(obj.AccessKeySecret),
		Endpoint:        tea.String(obj.Endpoint),
	})
	if err != nil {
		panic(err.Error())
	}
	return obj
}

// 获取Sts Token
func (vodThis *VodOfAliyun) Sts(ctx context.Context, param VodParam) (stsInfo map[string]any, err error) {
	stsInfo, err = utils.CreateStsToken(vodThis.client, &sts20150401.AssumeRoleRequest{
		DurationSeconds: tea.Int64(param.ExpireTime),
		//写入权限：{"Statement": [{"Action": ["oss:PutObject","oss:ListParts","oss:AbortMultipartUpload"],"Effect": "Allow","Resource": ["acs:oss:*:*:$BUCKET_NAME/$OBJECT_PREFIX*"]}],"Version": "1"}
		//读取权限：{"Statement": [{"Action": ["oss:GetObject"],"Effect": "Allow","Resource": ["acs:oss:*:*:$BUCKET_NAME/$OBJECT_PREFIX*"]}],"Version": "1"}
		Policy:          tea.String(`{"Statement": [{"Action": ["vod:*"],"Effect": "Allow","Resource": "*"}],"Version": "1"}`),
		RoleArn:         tea.String(vodThis.RoleArn),
		RoleSessionName: tea.String(`sts_token_to_vod`),
	})
	return
}
