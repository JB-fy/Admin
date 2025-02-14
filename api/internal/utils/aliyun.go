package utils

import (
	"errors"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	sts20150401 "github.com/alibabacloud-go/sts-20150401/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

// 生成Sts Token
func CreateStsToken(config *openapi.Config, assumeRoleRequest *sts20150401.AssumeRoleRequest) (stsInfo map[string]any, err error) {
	/* config := &openapi.Config{
		AccessKeyId:     tea.String(uploadThis.AccessKeyId),
		AccessKeySecret: tea.String(uploadThis.AccessKeySecret),
		Endpoint:        tea.String(uploadThis.Endpoint),
	} */
	client, err := sts20150401.NewClient(config)
	if err != nil {
		return
	}

	/* assumeRoleRequest := &sts20150401.AssumeRoleRequest{
		DurationSeconds: tea.Int64(stsOption.ExpireTime),
		// ExternalId : tea.String(stsOption.ExternalId),
		Policy:          tea.String(stsOption.Policy),
		RoleArn:         tea.String(stsOption.RoleArn),
		RoleSessionName: tea.String(stsOption.SessionName),
	} */
	tryErr := func() (err error) {
		defer func() {
			if errTmp := tea.Recover(recover()); errTmp != nil {
				err = errTmp
			}
		}()
		result, err := client.AssumeRoleWithOptions(assumeRoleRequest, &util.RuntimeOptions{})
		if err != nil {
			return
		}
		if *result.StatusCode != 200 {
			err = errors.New(`Sts Token响应错误`)
			return
		}
		stsInfo = map[string]any{
			`StatusCode`:      *result.StatusCode,
			`RequestId`:       *result.Body.RequestId,
			`AccessKeyId`:     *result.Body.Credentials.AccessKeyId,
			`AccessKeySecret`: *result.Body.Credentials.AccessKeySecret,
			`Expiration`:      *result.Body.Credentials.Expiration,
			`SecurityToken`:   *result.Body.Credentials.SecurityToken,
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
