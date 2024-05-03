package vod

import (
	daoPlatform "api/internal/dao/platform"
	"api/internal/utils/common"
	"context"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	sts20150401 "github.com/alibabacloud-go/sts-20150401/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/gogf/gf/v2/util/gconv"
)

type VodOfAliyun struct {
	Ctx             context.Context
	AccessKeyId     string `json:"vodOfAliyunAccessKeyId"`
	AccessKeySecret string `json:"vodOfAliyunAccessKeySecret"`
	Endpoint        string `json:"vodOfAliyunEndpoint"`
	RoleArn         string `json:"vodOfAliyunRoleArn"`
}

func NewVodOfAliyun(ctx context.Context, configOpt ...map[string]interface{}) *VodOfAliyun {
	var config map[string]interface{}
	if len(configOpt) > 0 && len(configOpt[0]) > 0 {
		config = configOpt[0]
	} else {
		configTmp, _ := daoPlatform.Config.Get(ctx, []string{`vodOfAliyunAccessKeyId`, `vodOfAliyunAccessKeySecret`, `vodOfAliyunEndpoint`, `vodOfAliyunRoleArn`})
		config = configTmp.Map()
	}

	vodOfAliyunObj := VodOfAliyun{Ctx: ctx}
	gconv.Struct(config, &vodOfAliyunObj)
	return &vodOfAliyunObj
}

// 获取Sts Token
func (uploadThis *VodOfAliyun) Sts(param VodParam) (stsInfo map[string]interface{}, err error) {
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
	stsInfo, err = common.CreateStsToken(uploadThis.Ctx, config, assumeRoleRequest)
	return
}
