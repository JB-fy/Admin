export default {
	label: {
		websiteConfig: '网站',
		appConfig: 'APP',
		uploadConfig: '上传',
		payConfig: '支付',
		smsConfig: '短信',
		idCardConfig: '实名认证',
		pushConfig: '推送',
		vodConfig: '视频点播',

		android: '安卓',
		ios: '苹果',

		payOfAli: '支付宝',
		payOfWx: '微信',
	},
	name: {
		hotSearch: '热门搜索',
		userAgreement: '用户协议',
		privacyAgreement: '隐私协议',

		packageUrlOfAndroid: '安装包(安卓)',
		packageSizeOfAndroid: '包大小(安卓)',
		packageNameOfAndroid: '包名(安卓)',
		isForceUpdateOfAndroid: '强制更新(安卓)',
		versionNumberOfAndroid: '版本号(安卓)',
		versionNameOfAndroid: '版本名称(安卓)',
		versionIntroOfAndroid: '版本介绍(安卓)',

		packageUrlOfIos: '安装包(苹果)',
		packageSizeOfIos: '包大小(苹果)',
		packageNameOfIos: '包名(苹果)',
		isForceUpdateOfIos: '强制更新(苹果)',
		versionNumberOfIos: '版本号(苹果)',
		versionNameOfIos: '版本名称(苹果)',
		versionIntroOfIos: '版本介绍(苹果)',
		plistUrlOfIos: 'plist文件(苹果)',

		uploadType: '上传方式',
		localUploadUrl: '本地-上传地址',
		localUploadSignKey: '本地-密钥',
		localUploadFileSaveDir: '本地-文件保存目录',
		localUploadFileUrlPrefix: '本地-文件地址前缀',
		aliyunOssHost: '阿里云OSS-域名',
		aliyunOssBucket: '阿里云OSS-Bucket',
		aliyunOssAccessKeyId: '阿里云OSS-AccessKeyId',
		aliyunOssAccessKeySecret: '阿里云OSS-AccessKeySecret',
		aliyunOssCallbackUrl: '阿里云OSS-回调地址',
		aliyunOssEndpoint: '阿里云OSS-Endpoint',
		aliyunOssRoleArn: '阿里云OSS-RoleArn',

		payOfAliAppId: 'AppID',
		payOfAliSignType: '签名方式',
		payOfAliPrivateKey: '私钥',
		payOfAliPublicKey: '公钥',

		payOfWxAppId: 'AppID',
		payOfWxMchid: '商户ID',
		payOfWxSerialNo: '证书序列号',
		payOfWxApiV3Key: 'APIV3密钥',
		payOfWxPrivateKey: '私钥',

		smsType: '短信方式',
		smsOfAliyunAccessKeyId: '阿里云SMS-AccessKeyId',
		smsOfAliyunAccessKeySecret: '阿里云SMS-AccessKeySecret',
		smsOfAliyunEndpoint: '阿里云SMS-Endpoint',
		smsOfAliyunSignName: '阿里云SMS-签名',
		smsOfAliyunTemplateCode: '阿里云SMS-模板标识',

		idCardType: '实名认证方式',
		idCardOfAliyunHost: '阿里云IdCard-域名',
		idCardOfAliyunPath: '阿里云IdCard-请求路径',
		idCardOfAliyunAppcode: '阿里云IdCard-Appcode',

		pushType: '推送方式',
		txTpnsHost: '腾讯移动推送-域名',
		txTpnsAccessIDOfAndroid: '腾讯移动推送-AccessID(安卓)',
		txTpnsSecretKeyOfAndroid: '腾讯移动推送-SecretKey(安卓)',
		txTpnsAccessIDOfIos: '腾讯移动推送-AccessID(苹果)',
		txTpnsSecretKeyOfIos: '腾讯移动推送-SecretKey(苹果)',
		txTpnsAccessIDOfMacOS: '腾讯移动推送-AccessID(苹果电脑)',
		txTpnsSecretKeyOfMacOS: '腾讯移动推送-SecretKey(苹果电脑)',

		vodType: '视频点播方式',
		vodOfAliyunAccessKeyId: '阿里云VOD-AccessKeyId',
		vodOfAliyunAccessKeySecret: '阿里云VOD-AccessKeySecret',
		vodOfAliyunEndpoint: '阿里云VOD-Endpoint',
		vodOfAliyunRoleArn: '阿里云VOD-RoleArn',
	},
	status: {
		uploadType: [
			{ value: `local`, label: '本地' },
			{ value: `aliyunOss`, label: '阿里云' },
		],
		payOfAliSignType: [
			{ value: `RSA2`, label: 'RSA2' },
			{ value: `RSA`, label: 'RSA' },
		],
		smsType: [
			{ value: `smsOfAliyun`, label: '阿里云' },
		],
		idCardType: [
			{ value: `idCardOfAliyun`, label: '阿里云' },
		],
		pushType: [
			{ value: `txTpns`, label: '腾讯移动推送' },
		],
		vodType: [
			{ value: `vodOfAliyun`, label: '阿里云' },
		],
	},
	tip: {
		localUploadFileSaveDir: '根据部署的线上环境设置。一般与nginx中设置的网站对外目录一致',
		localUploadFileUrlPrefix: '根据部署的线上环境设置。与文件保存路径拼接形成文件访问地址',
		aliyunOssHost: '不含Bucket部分',
		aliyunOssCallbackUrl: '设置后开启回调，否则关闭回调',
		aliyunOssEndpoint: 'APP直传需设置，用于生成STS凭证。请参考：<a target="_blank" href="https://api.aliyun.com/product/Sts">https://api.aliyun.com/product/Sts</a>',
		aliyunOssRoleArn: 'APP直传需设置，用于生成STS凭证',

		idCardOfAliyunHost: '购买地址：<a target="_blank" href="https://market.aliyun.com/products/57000002/cmapi014760.html">https://market.aliyun.com/products/57000002/cmapi014760.html</a>（购买其它接口，只需对代码文件做下简单修改即可）',

		txTpnsHost: '参考：<a target="_blank" href="https://cloud.tencent.com/document/product/548/49157">https://cloud.tencent.com/document/product/548/49157</a>',

		vodOfAliyunEndpoint: '用于生成STS凭证。请参考：<a target="_blank" href="https://api.aliyun.com/product/Sts">https://api.aliyun.com/product/Sts</a>',
		vodOfAliyunRoleArn: '用于生成STS凭证',
	},
}