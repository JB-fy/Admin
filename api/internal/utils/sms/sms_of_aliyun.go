package sms

import (
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
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	Endpoint        string `json:"endpoint"`
	SignName        string `json:"signName"`
	TemplateCode    string `json:"templateCode"`
}

func NewSmsOfAliyun(ctx context.Context, config map[string]any) *SmsOfAliyun {
	smsObj := &SmsOfAliyun{}
	gconv.Struct(config, smsObj)
	if smsObj.AccessKeyId == `` || smsObj.AccessKeySecret == `` || smsObj.Endpoint == `` {
		panic(`缺少插件配置：短信-阿里云`)
	}
	/* if smsObj.SignName == `` || smsObj.TemplateCode == `` {
		panic(`缺少插件配置：短信-阿里云短信模板`)
	} */
	return smsObj
}

func (smsThis *SmsOfAliyun) SendCode(ctx context.Context, phone string, code string) (err error) {
	if smsThis.SignName == `` || smsThis.TemplateCode == `` {
		err = errors.New(`缺少插件配置：短信-阿里云短信模板`)
		return
	}
	err = smsThis.SendSms(ctx, []string{phone}, `{"code": "`+code+`"}`, smsThis.SignName, smsThis.TemplateCode)
	return
}

func (smsThis *SmsOfAliyun) SendSms(ctx context.Context, phoneArr []string, message string, paramOpt ...any) (err error) {
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
