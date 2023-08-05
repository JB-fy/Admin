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
- 表字段命名需要遵守以下规则，否则只会根据字段类型做默认处理
- 主键必须在第一个字段。否则需要在dao层重写PrimaryKey方法返回主键字段
- 表内尽量根据表名设置xxxxId和xxxxName两个字段(这两字段，常用于前端部分组件，服务端请求获取id和label两个字段用于列表展示)
- 每个字段都必须有注释。以下符号[\n\r.。:：(（]之前的部分或整个注释，将作为字段名称使用
    - 部分常用字段：
        - password|passwd	密码（password|passwd,salt同时存在时特殊处理）
        - salt			    加密盐
        - pid				父级（指向本表主键。pid,level,idPath|id_path同时存在时，有特殊处理）
        - level			    层级
        - idPath|id_path	层级路径
        - sort			    排序
        - weight            权重
        - gender 			性别
        - avatar			头像
    - 其他类型字段：
        - 名称和标识字段，命名用name或code后缀
        - 手机号码字段，命名用mobile或phone后缀
        - 链接地址字段，命名用url或link后缀
        - 关联id字段和关联表主键保持一致，命名用id后缀
        - 图片字段，命名用icon,cover或img,img_list,imgList,img_arr,imgArr或image,image_list,imageList,image_arr,imageArr等后缀（多图片时字段类型用json或text，保存格式为JSON格式）
        - 视频字段，命名用video,video_list,videoList,video_arr,videoArr等后缀（多视频时字段类型用json或text，保存格式为JSON格式）
        - ip字段，命名用Ip后缀
        - 备注字段，命名用remark后缀
        - 状态和类型字段，命名用status或type后缀且字段类型必须是int或tinyint。字段注释中多状态之间用[\s,，;；]等字符分隔。如（状态：0待处理 1已处理 2驳回）
        - 是否字段，命名用is_前缀且字段类型必须是int或tinyint。默认：0否 1是

2.修改api/hack/config.yaml中gen dao的配置，并使用gf gen dao命令生成dao层代码

3.使用./main myGen命令生成前后端代码
- 命令参数可以参考api/internal/cmd/my_gen.go内结构体MyGenOption每个字段的注释
- 未带参数启动，则需根据提示输入全部参数。示例：./main myGen
- 带全部参数启动，则无提示直接生成。示例：./main myGen -sceneCode=platform -dbGroup=default -dbTable=auth_test -removePrefix=auth_ -moduleDir=auth -commonName=测试 -isList=yes -isCreate=yes -isUpdate=yes -isDelete=yes -isApi=yes -isAuthAction=yes -isView=yes -isCover=no