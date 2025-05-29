package aliyun

import (
	"api/internal/utils/sms/model"
	"context"
	"errors"
	"strings"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/gogf/gf/v2/util/gconv"
)

type Sms struct {
	AccessKeyId     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
	Endpoint        string `json:"endpoint"`
	SignName        string `json:"sign_name"`
	TemplateCode    string `json:"template_code"`
	client          *dysmsapi20170525.Client
}

func NewSms(ctx context.Context, config map[string]any) model.Sms {
	obj := &Sms{}
	gconv.Struct(config, obj)
	if obj.AccessKeyId == `` || obj.AccessKeySecret == `` || obj.Endpoint == `` {
		panic(`缺少插件配置：短信-阿里云`)
	}
	var err error
	obj.client, err = dysmsapi20170525.NewClient(&openapi.Config{
		AccessKeyId:     tea.String(obj.AccessKeyId),
		AccessKeySecret: tea.String(obj.AccessKeySecret),
		Endpoint:        tea.String(obj.Endpoint),
	})
	if err != nil {
		panic(err.Error())
	}
	return obj
}

func (smsThis *Sms) SendCode(ctx context.Context, phone string, code string) (err error) {
	if smsThis.SignName == `` || smsThis.TemplateCode == `` {
		err = errors.New(`缺少插件配置：短信-阿里云短信模板`)
		return
	}
	err = smsThis.SendSms(ctx, []string{phone}, `{"code": "`+code+`"}`, smsThis.SignName, smsThis.TemplateCode)
	return
}

func (smsThis *Sms) SendSms(ctx context.Context, phoneArr []string, message string, paramOpt ...any) (err error) {
	if len(paramOpt) < 2 {
		err = errors.New(`缺少签名和模板参数`)
		return
	}

	tryErr := func() (err error) {
		defer func() {
			if errTmp := tea.Recover(recover()); errTmp != nil {
				err = errTmp
			}
		}()
		result, err := smsThis.client.SendSmsWithOptions(&dysmsapi20170525.SendSmsRequest{
			PhoneNumbers:  tea.String(strings.Join(phoneArr, `,`)),
			SignName:      tea.String(gconv.String(paramOpt[0])),
			TemplateCode:  tea.String(gconv.String(paramOpt[1])),
			TemplateParam: tea.String(message),
		}, &util.RuntimeOptions{})
		if err != nil {
			return
		}
		if *result.Body.Code != `OK` {
			err = errors.New(*result.Body.Message)
			return
		}
		return
	}()

	if tryErr != nil {
		var errSDK = &tea.SDKError{Message: tea.String(tryErr.Error())}
		if errSDKTmp, ok := tryErr.(*tea.SDKError); ok {
			errSDK = errSDKTmp
		}
		var errMsg *string
		errMsg, err = util.AssertAsString(errSDK.Message)
		if err != nil {
			return
		}
		err = errors.New(*errMsg)
	}
	return
}
