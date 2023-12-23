package common

import (
	"errors"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	sts20150401 "github.com/alibabacloud-go/sts-20150401/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"golang.org/x/net/context"
)

// 生成Sts Token
func CreateStsToken(ctx context.Context, config *openapi.Config, assumeRoleRequest *sts20150401.AssumeRoleRequest) (stsInfo map[string]interface{}, err error) {
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
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		result, errTmp := client.AssumeRoleWithOptions(assumeRoleRequest, &util.RuntimeOptions{})
		if errTmp != nil {
			return errTmp
		}
		if *result.StatusCode != 200 {
			err = errors.New(`Sts Token响应错误`)
			return err
		}
		stsInfo = map[string]interface{}{
			`StatusCode`:      *result.StatusCode,
			`RequestId`:       *result.Body.RequestId,
			`AccessKeyId`:     *result.Body.Credentials.AccessKeyId,
			`AccessKeySecret`: *result.Body.Credentials.AccessKeySecret,
			`Expiration`:      *result.Body.Credentials.Expiration,
			`SecurityToken`:   *result.Body.Credentials.SecurityToken,
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
		errMsg, errTmp := util.AssertAsString(errSDK.Message)
		if errTmp != nil {
			err = errTmp
			return
		}
		err = errors.New(*errMsg)
	}
	return
}
