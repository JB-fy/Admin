package sms

import (
	daoPlatform "api/internal/dao/platform"
	"context"
	"errors"
	"strings"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/gogf/gf/v2/util/gconv"
)

type SmsOfAliyun struct {
	Ctx             context.Context
	AccessKeyId     string `json:"smsOfAliyunAccessKeyId"`
	AccessKeySecret string `json:"smsOfAliyunAccessKeySecret"`
	Endpoint        string `json:"smsOfAliyunEndpoint"`
	SignName        string `json:"smsOfAliyunSignName"`
	TemplateCode    string `json:"smsOfAliyunTemplateCode"`
}

func NewSmsOfAliyun(ctx context.Context, configOpt ...map[string]any) *SmsOfAliyun {
	var config map[string]any
	if len(configOpt) > 0 && len(configOpt[0]) > 0 {
		config = configOpt[0]
	} else {
		configTmp, _ := daoPlatform.Config.Get(ctx, []string{`smsOfAliyunAccessKeyId`, `smsOfAliyunAccessKeySecret`, `smsOfAliyunEndpoint`, `smsOfAliyunSignName`, `smsOfAliyunTemplateCode`})
		config = configTmp.Map()
	}

	smsOfAliyunObj := SmsOfAliyun{Ctx: ctx}
	gconv.Struct(config, &smsOfAliyunObj)
	return &smsOfAliyunObj
}

func (smsThis *SmsOfAliyun) SendCode(phone string, code string) (err error) {
	err = smsThis.SendSms([]string{phone}, `{"code": "`+code+`"}`, smsThis.SignName, smsThis.TemplateCode)
	return
}

func (smsThis *SmsOfAliyun) SendSms(phoneArr []string, message string, paramOpt ...any) (err error) {
	if len(paramOpt) < 2 {
		err = errors.New(`缺少签名和模板参数`)
		return
	}

	client, err := smsThis.CreateClient()
	if err != nil {
		return
	}

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(strings.Join(phoneArr, `,`)),
		SignName:      tea.String(gconv.String(paramOpt[0])),
		TemplateCode:  tea.String(gconv.String(paramOpt[1])),
		TemplateParam: tea.String(message),
	}

	tryErr := func() (_e error) {
		defer func() {
			r := tea.Recover(recover())
			if r != nil {
				_e = r
			}
		}()
		result, err := client.SendSmsWithOptions(sendSmsRequest, &util.RuntimeOptions{})
		if err != nil {
			return err
		}
		if *result.Body.Code != `OK` {
			err = errors.New(*result.Body.Message)
			return err
		}
		return nil
	}()

	if tryErr != nil {
		var errSDK = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			errSDK = _t
		} else {
			errSDK.Message = tea.String(tryErr.Error())
		}
		errMsg, errSms := util.AssertAsString(errSDK.Message)
		if errSms != nil {
			err = errSms
			return
		}
		err = errors.New(*errMsg)
	}
	return
}

func (smsThis *SmsOfAliyun) CreateClient() (client *dysmsapi20170525.Client, err error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(smsThis.AccessKeyId),
		AccessKeySecret: tea.String(smsThis.AccessKeySecret),
		Endpoint:        tea.String(smsThis.Endpoint),
	}
	client, err = dysmsapi20170525.NewClient(config)
	return
}
