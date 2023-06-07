package api

type ConfigGetReq struct {
	ConfigKeyArr *[]string `c:"configKeyArr,omitempty" p:"configKeyArr" v:"foreach|min-length:1"`
}

type ConfigSaveReq struct {
	AliyunOssAccessId     *string `c:"aliyunOssAccessId,omitempty" p:"aliyunOssAccessId" v:"regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	AliyunOssAccessSecret *string `c:"aliyunOssAccessSecret,omitempty" p:"aliyunOssAccessSecret" v:"regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	AliyunOssHost         *string `c:"aliyunOssHost,omitempty" p:"aliyunOssHost" v:"url"`
	AliyunOssBucket       *string `c:"aliyunOssBucket,omitempty" p:"aliyunOssBucket" v:"regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
}
