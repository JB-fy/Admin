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
	Ctx             context.Context
	AccessKeyId     string `json:"smsOfAliyunAccessKeyId"`
	AccessKeySecret string `json:"smsOfAliyunAccessKeySecret"`
	Endpoint        string `json:"smsOfAliyunEndpoint"`
	SignName        string `json:"smsOfAliyunSignName"`
	TemplateCode    string `json:"smsOfAliyunTemplateCode"`
}

func NewSmsOfAliyun(ctx context.Context, config map[string]interface{}) *SmsOfAliyun {
	smsOfAliyunObj := SmsOfAliyun{
		Ctx: ctx,
	}
	gconv.Struct(config, &smsOfAliyunObj)
	return &smsOfAliyunObj
}

func (smsThis *SmsOfAliyun) Send(phone string, code string) (err error) {
	err = smsThis.SendSms([]string{phone}, `{"code": "`+code+`"}`)
	return
}

func (smsThis *SmsOfAliyun) SendSms(phoneArr []string, templateParam string) (err error) {
	client, err := smsThis.CreateClient()
	if err != nil {
		return
	}

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(strings.Join(phoneArr, `,`)),
		SignName:      tea.String(smsThis.SignName),
		TemplateCode:  tea.String(smsThis.TemplateCode),
		TemplateParam: tea.String(templateParam),
		// TemplateParam: tea.String(`{"code": "1234"}`),
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
		_t, ok := tryErr.(*tea.SDKError)
		if ok {
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
