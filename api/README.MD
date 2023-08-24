# 常用命令

```bash
# 服务启动
gf run main.go -a "http"

# dao层生成
gf gen dao

# service层生成
gf gen service

# 全自动前后端代码生成（必须先生成main文件。如修改过api/internal/cmd/my_gen.go代码，则需重新生成）
./main myGen
```
# 代码自动生成流程（强烈建议搭配Git使用）

1.新建数据库表。如：auth_test
```
表字段命名需要遵守以下规则，否则只会根据字段类型做默认处理
主键必须在第一个字段。否则需要在dao层重写PrimaryKey方法返回主键字段
表内尽量根据表名设置xxxxId和xxxxName两个字段(这两字段，常用于前端部分组件，服务端请求获取id和label两个字段用于列表展示)
每个字段都必须有注释。以下符号[\n\r.。:：(（]之前的部分或整个注释，将作为字段名称使用
	部分常用字段：
		密码 		命名：password|passwd；类型：char(32)；注意：salt同时存在时，有特殊处理
		加密盐 		命名：salt；类型：char(8)；
		父级		命名：pid；类型：int等类型；注意：level,idPath|id_path同时存在时，有特殊处理
		层级		命名：level；类型：int等类型
		层级路径	命名：idPath|id_path；类型：varchar
		排序		命名：sort；类型：int等类型
		权重		命名：weight；类型：int等类型
		性别		命名：gender；类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（性别：0未设置 1男 2女）
		头像		命名：avatar；类型：varchar
	其他类型字段：
		名称		命名：name后缀；类型：varchar
		标识		命名：code后缀；类型：varchar
		手机		命名：mobile或phone后缀；类型：varchar
		链接		命名：url或link后缀；类型：varchar
		关联ID		命名：id后缀；类型：int等类型
		图片		命名：icon,cover或img,img_list,imgList,img_arr,imgArr或image,image_list,imageList,image_arr,imageArr等后缀；类型：单图片varchar，多图片json或text
		视频		命名：video,video_list,videoList,video_arr,videoArr等后缀；类型：单视频varchar，多视频json或text
		数组		命名：list或arr等后缀；类型：json或text
		IP			命名：IP后缀；类型：varchar
		备注描述	命名：remark或desc后缀；类型：varchar（生成的表单组件：textarea多行文本输入框）
		富文本		命名：intro或content后缀；类型：text（生成的表单组件：tinymce富文本编辑器）
		状态		命名：status后缀；类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回）
		类型		命名：type后缀；类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（类型：0安卓 1苹果）
		是否		命名：is_前缀；类型：int等类型；注释：示例（停用：0否 1是）
		开始时间	命名：start_前缀；类型：timestamp或datetime或date
		结束时间	命名：end_前缀；类型：timestamp或datetime或date
```

2.修改api/hack/config.yaml中gen dao的配置，并使用gf gen dao命令生成dao层代码

3.使用./main myGen命令生成前后端代码（参数参考api/internal/cmd/my_gen.go内结构体MyGenOption的字段注释）
- 未带参数执行，则需根据提示输入全部参数。示例：./main myGen
- 带全部参数执行，则无提示直接生成。示例：./main myGen -sceneCode=platform -dbGroup=default -dbTable=auth_test -removePrefix=auth_ -moduleDir=auth -commonName=测试 -isList=yes -isCount=yes -isInfo=yes -isCreate=yes -isUpdate=yes -isDelete=yes -isApi=yes -isAuthAction=yes -isView=yes -isCover=no

# 常见错误处理

1.遇到dao层互相引用错误import cycle not allowed时，参考文件[allow.go](./internal/dao/auth/allow/allow.go)解决