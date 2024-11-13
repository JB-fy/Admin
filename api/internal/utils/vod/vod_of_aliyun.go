package vod

import (
	"api/internal/utils/common"
	"context"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	sts20150401 "github.com/alibabacloud-go/sts-20150401/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/gogf/gf/v2/util/gconv"
)

type VodOfAliyun struct {
	Ctx             context.Context
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	Endpoint        string `json:"endpoint"`
	RoleArn         string `json:"roleArn"`
}

func NewVodOfAliyun(ctx context.Context, config map[string]any) *VodOfAliyun {
	vodObj := &VodOfAliyun{Ctx: ctx}
	gconv.Struct(config, vodObj)
	if vodObj.AccessKeyId == `` || vodObj.AccessKeySecret == `` || vodObj.Endpoint == `` || vodObj.RoleArn == `` {
		panic(`缺少插件配置：视频点播-阿里云`)
	}
	return vodObj
}

// 获取Sts Token
func (vodThis *VodOfAliyun) Sts(param VodParam) (stsInfo map[string]any, err error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(vodThis.AccessKeyId),
		AccessKeySecret: tea.String(vodThis.AccessKeySecret),
		Endpoint:        tea.String(vodThis.Endpoint),
	}
	assumeRoleRequest := &sts20150401.AssumeRoleRequest{
		DurationSeconds: tea.Int64(param.ExpireTime),
		//写入权限：{"Statement": [{"Action": ["oss:PutObject","oss:ListParts","oss:AbortMultipartUpload"],"Effect": "Allow","Resource": ["acs:oss:*:*:$BUCKET_NAME/$OBJECT_PREFIX*"]}],"Version": "1"}
		//读取权限：{"Statement": [{"Action": ["oss:GetObject"],"Effect": "Allow","Resource": ["acs:oss:*:*:$BUCKET_NAME/$OBJECT_PREFIX*"]}],"Version": "1"}
		Policy:          tea.String(`{"Statement": [{"Action": ["vod:*"],"Effect": "Allow","Resource": "*"}],"Version": "1"}`),
		RoleArn:         tea.String(vodThis.RoleArn),
		RoleSessionName: tea.String(`sts_token_to_vod`),
	}
	stsInfo, err = common.CreateStsToken(vodThis.Ctx, config, assumeRoleRequest)
	return
}
