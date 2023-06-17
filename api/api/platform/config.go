package api

type ConfigGetReq struct {
	ConfigKeyArr *[]string `c:"configKeyArr,omitempty" p:"configKeyArr" v:"required|distinct|foreach|min-length:1"`
}

type ConfigSaveReq struct {
	AliyunOssAccessKeyId     *string `c:"aliyunOssAccessKeyId,omitempty" p:"aliyunOssAccessKeyId" v:"regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	AliyunOssAccessKeySecret *string `c:"aliyunOssAccessKeySecret,omitempty" p:"aliyunOssAccessKeySecret" v:"regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	AliyunOssHost            *string `c:"aliyunOssHost,omitempty" p:"aliyunOssHost" v:"url"`
	AliyunOssBucket          *string `c:"aliyunOssBucket,omitempty" p:"aliyunOssBucket" v:"regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
}
