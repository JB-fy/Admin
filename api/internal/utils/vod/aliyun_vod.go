package vod

import (
	"api/internal/utils/common"
	"context"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	sts20150401 "github.com/alibabacloud-go/sts-20150401/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/gogf/gf/v2/util/gconv"
)

type AliyunVod struct {
	Ctx             context.Context
	AccessKeyId     string `json:"aliyunVodAccessKeyId"`
	AccessKeySecret string `json:"aliyunVodAccessKeySecret"`
	RoleArn         string `json:"aliyunVodRoleArn"`
	Endpoint        string `json:"aliyunVodEndpoint"`
}

func NewAliyunVod(ctx context.Context, config map[string]interface{}) *AliyunVod {
	aliyunVodObj := AliyunVod{
		Ctx: ctx,
	}
	gconv.Struct(config, &aliyunVodObj)
	return &aliyunVodObj
}

// 获取Sts Token
func (uploadThis *AliyunVod) Sts(param VodParam) (stsInfo map[string]interface{}, err error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(uploadThis.AccessKeyId),
		AccessKeySecret: tea.String(uploadThis.AccessKeySecret),
		Endpoint:        tea.String(uploadThis.Endpoint),
	}
	assumeRoleRequest := &sts20150401.AssumeRoleRequest{
		DurationSeconds: tea.Int64(param.ExpireTime),
		//写入权限：{"Statement": [{"Action": ["oss:PutObject","oss:ListParts","oss:AbortMultipartUpload"],"Effect": "Allow","Resource": ["acs:oss:*:*:$BUCKET_NAME/$OBJECT_PREFIX*"]}],"Version": "1"}
		//读取权限：{"Statement": [{"Action": ["oss:GetObject"],"Effect": "Allow","Resource": ["acs:oss:*:*:$BUCKET_NAME/$OBJECT_PREFIX*"]}],"Version": "1"}
		Policy:          tea.String(`{"Statement": [{"Action": ["vod:*"],"Effect": "Allow","Resource": "*"}],"Version": "1"}`),
		RoleArn:         tea.String(uploadThis.RoleArn),
		RoleSessionName: tea.String(`sts_token_to_vod`),
	}
	stsInfo, _ = common.CreateStsToken(uploadThis.Ctx, config, assumeRoleRequest)
	return
}
