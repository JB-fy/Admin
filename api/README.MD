# 常用命令

```bash
# 服务启动（--gf.gcfg.file用于指定配置文件，默认config.yaml）
gf run main.go -a "http --gf.gcfg.file=config.prod.yaml"

# dao层生成
gf gen dao

# service层生成
gf gen service

# 全自动代码生成（需先编译生成可执行文件。如修改api/internal/cmd/my_gen.go代码，需重新编译才会生效）
./main myGen
```
# 代码自动生成流程（强烈建议搭配Git使用）

1.新建数据库表。如：auth_test
```
主键必须在第一个字段。否则需要在dao层重写PrimaryKey方法返回主键字段
表内尽量根据表名设置xxxxId主键和xxxxName名称两个字段（作用1：常用于前端部分组件，如MySelect.vue组件；作用2：当其它表存在与该表主键同名的关联字段时，会自动生成联表查询代码）
每个字段都必须有注释。以下符号[\n\r.。:：(（]之前的部分或整个注释，将作为字段名称使用
表字段按以下规则命名时，会做特殊处理，其它情况根据字段类型做默认处理
	固定命名：
		父级		命名：pid；      		类型：int等类型；		注意：pid,level,idPath|id_path同时存在时，有特殊处理
		层级		命名：level；          	类型：int等类型；		注意：pid,level,idPath|id_path同时存在时，(才)有特殊处理
		层级路径	命名：idPath|id_path；	类型：varchar或text；	注意：pid,level,idPath|id_path同时存在时，(才)有特殊处理
		排序		命名：sort；			类型：int等类型；		注意：pid,level,idPath|id_path|sort同时存在时，(才)有特殊处理
	常用命名(字段含[_of_]时，会忽略[_of_]及其之后的部分)：
		密码		命名：password后缀；			类型：char(32)；
		加密盐 		命名：salt后缀；     			类型：char；	注意：password,salt同时存在时，有特殊处理
		名称		命名：name后缀；				类型：varchar；
		标识		命名：code后缀；				类型：varchar；
		手机		命名：mobile,phone后缀；		类型：varchar；
		链接		命名：url,link后缀；			类型：varchar；
		IP			命名：IP后缀；					类型：varchar；
		关联ID		命名：id后缀；					类型：int等类型；
		排序|权重	命名：sort,weight等后缀；		类型：int等类型；
		状态|类型	命名：status,type,gender等后缀；类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回）
		是否		命名：is_前缀；					类型：int等类型；注释：示例（停用：0否 1是）
		开始时间	命名：start_前缀；				类型：timestamp或datetime或date；
		结束时间	命名：end_前缀；				类型：timestamp或datetime或date；
		(富)文本	命名：remark,desc,msg,message,intro,content后缀；类型：varchar或text；前端生成表单组件：textarea文本输入框或tinymce富文本编辑器
		图片		命名：icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；类型：单图片varchar，多图片json或text
		视频		命名：video,video_list,videoList,video_arr,videoArr等后缀；类型：单视频varchar，多视频json或text
		数组		命名：list,arr等后缀；类型：json或text；
```

2.修改api/hack/config.yaml中gen dao的配置，并使用gf gen dao命令生成dao层代码

3.使用./main myGen命令生成前后端代码（参数参考api/internal/cmd/my_gen.go内结构体MyGenOption的字段注释）
- 未带参数执行，则需根据提示输入全部参数。示例：./main myGen
- 带全部参数执行，则无提示直接生成。示例：./main myGen -sceneCode=platform -dbGroup=default -dbTable=auth_test -removePrefix=auth_ -moduleDir=auth -commonName=权限管理/测试 -isList=1 -isCount=1 -isInfo=1 -isCreate=1 -isUpdate=1 -isDelete=1 -isApi=1 -isAuthAction=1 -isView=1 -isCover=0

# 常见错误处理

1.遇到dao层互相引用错误import cycle not allowed时，参考文件[allow.go](./internal/dao/user/allow/allow.go)解决