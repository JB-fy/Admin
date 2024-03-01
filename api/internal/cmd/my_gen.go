package cmd

import (
	daoAuth "api/internal/dao/auth"
	"api/internal/utils"
	"context"
	"fmt"
	"os/exec"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

/*
后台常用生成示例：./main myGen -sceneCode=platform -dbGroup=default -dbTable=auth_test -removePrefix=auth_ -moduleDir=auth -commonName=权限管理/测试 -isList=1 -isCount=1 -isInfo=1 -isCreate=1 -isUpdate=1 -isDelete=1 -isApi=1 -isAuthAction=1 -isView=1 -isCover=0
APP常用生成示例：./main myGen -sceneCode=app -dbGroup=xxxx -dbTable=user -removePrefix= -moduleDir=xxxx/user -commonName=用户 -isList=1 -isCount=0 -isInfo=1 -isCreate=0 -isUpdate=0 -isDelete=0 -isApi=1 -isAuthAction=0 -isView=0 -isCover=0

强烈建议搭配Git使用

表名统一使用蛇形命名。不同功能表按以下表规则命名
	主表：正常命名即可。参考以下示例
		platform_admin
		user
		good
		good_category
	扩展表（一对一）：表命名：主表名_xxxx，且存在与主表主键同名的字段，该字段为不递增主键或唯一索引
	扩展表（一对多）：表命名：主表名_xxxx，且存在与主表主键同名的字段，该字段设置普通索引
		参考以下示例
			user_config		说明：存放user主表用户的配置信息
			good_content	说明：存放good主表商品的详情
	中间表（一对一）：表命名使用_rel_to_或_rel_of_关联两表，不同模块两表必须全名，同模块第二个表可全名也可省略前缀。存在与两个关联表主键同名的字段，用_rel_to_做关联时，第一个表的关联字段做主键或唯一索引，用_rel_of_做关联时，第二个表的关联字段做主键或唯一索引。
	中间表（一对多）：表命名使用_rel_to_或_rel_of_关联两表，不同模块两表必须全名，同模块第二个表可全名也可省略前缀。存在与两个关联表主键同名的字段，两关联字段做联合主键或联合唯一索引
		参考以下示例
			auth_role_rel_of_platform_admin	说明：auth_role和platform_admin属不同模块，中间表命名使用两表全名
			good_rel_to_category			说明：good和good_category属同模块，故good_category可省略good_前缀

表字段名统一使用小驼峰或蛇形命名（建议：小驼峰）
	主键必须在第一个字段。否则需要在dao层重写PrimaryKey方法返回主键字段

	尽量根据表名设置xxxxId主键和xxxxName名称两个字段（作用1：常用于前端部分组件，如MySelect.vue组件；作用2：当其它表存在与该表主键同名的关联字段时，会自动生成联表查询代码）

	字段都必须有注释。以下符号[\n\r.。:：(（]之前的部分或整个注释，将作为字段名称使用

	字段按以下规则命名时，会做特殊处理，其它情况根据字段类型做默认处理
		固定命名：
			父级		命名：pid；      		类型：int等类型；		注意：pid,level,idPath|id_path同时存在时，有特殊处理
			层级		命名：level；          	类型：int等类型；		注意：pid,level,idPath|id_path同时存在时，(才)有特殊处理
			层级路径	命名：idPath|id_path；	类型：varchar或text；	注意：pid,level,idPath|id_path同时存在时，(才)有特殊处理
			排序		命名：sort；			类型：int等类型；		注意：pid,level,idPath|id_path|sort同时存在时，(才)有特殊处理

		常用命名(字段含[_of_]时，会忽略[_of_]及其之后的部分)：
			密码		命名：password,passwd后缀；		类型：char(32)；
			加密盐 		命名：salt后缀；     			类型：char；	注意：password,salt同时存在时，有特殊处理
			名称		命名：name后缀；				类型：varchar；
			标识		命名：code后缀；				类型：varchar；
			手机		命名：phone,mobile后缀；		类型：varchar；
			链接		命名：url,link后缀；			类型：varchar；
			IP			命名：IP后缀；					类型：varchar；
			关联ID		命名：id后缀；					类型：int等类型；
			排序|权重	命名：sort,weight等后缀；		类型：int等类型；
			状态|类型	命名：status,type,method,pos,position,gender等后缀；类型：int等类型或varchar或char；注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
			是否		命名：is_前缀；					类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）
			开始时间	命名：start_前缀；				类型：timestamp或datetime或date；
			结束时间	命名：end_前缀；				类型：timestamp或datetime或date；
			(富)文本	命名：remark,desc,msg,message,intro,content后缀；类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
			图片		命名：icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；类型：单图片varchar，多图片json或text
			视频		命名：video,video_list,videoList,video_arr,videoArr等后缀；类型：单视频varchar，多视频json或text
			数组		命名：list,arr等后缀；类型：json或text；
*/

func MyGenFunc(ctx context.Context, parser *gcmd.Parser) (err error) {
	myGenHandlerObj := myGenHandler{ctx: ctx}
	myGenHandlerObj.init(parser)
	myGenHandlerObj.tpl = myGenHandlerObj.setTpl(myGenHandlerObj.option.DbTable, myGenHandlerObj.option.RemovePrefixAlone)

	myGenHandlerObj.genDao()   // dao层存在时，增加或修改部分字段的解析代码
	myGenHandlerObj.genLogic() // logic模板生成（文件不存在时增删改查全部生成，已存在不处理不覆盖）

	if myGenHandlerObj.option.IsApi {
		myGenHandlerObj.genApi()        // api模板生成
		myGenHandlerObj.genController() // controller模板生成
		myGenHandlerObj.genRouter()     // 后端路由生成
	}

	if myGenHandlerObj.option.IsView {
		myGenHandlerObj.genViewIndex()  // 视图模板Index生成
		myGenHandlerObj.genViewList()   // 视图模板List生成
		myGenHandlerObj.genViewQuery()  // 视图模板Query生成
		myGenHandlerObj.genViewSave()   // 视图模板Save生成
		myGenHandlerObj.genViewI18n()   // 视图模板I18n生成
		myGenHandlerObj.genViewRouter() // 前端路由生成
		// 前端代码格式化
		myGenHandlerObj.command(`前端代码格式化`, false, gfile.SelfDir()+`/../view/`+myGenHandlerObj.option.SceneCode,
			`npm`, `run`, `format`)
	}
	return
}

type myGenHandler struct {
	ctx       context.Context
	sceneId   uint     //场景ID
	sceneName string   //场景名称
	dbLink    string   //当前数据库连接配置（gf gen dao命令生成dao时需要）
	db        gdb.DB   //当前数据库连接
	tableArr  []string //当前db全部数据表
	option    myGenOption
	tpl       myGenTpl
}

type myGenOption struct {
	SceneCode          string `json:"sceneCode"`          //场景标识，必须在数据库表auth_scene已存在。示例：platform
	DbGroup            string `json:"dbGroup"`            //db分组。示例：default
	DbTable            string `json:"dbTable"`            //db表。示例：auth_test
	RemovePrefixCommon string `json:"removePrefixCommon"` //要删除的共有前缀，没有可为空。removePrefixCommon + removePrefixAlone必须和hack/config.yaml内removePrefix保持一致
	RemovePrefixAlone  string `json:"removePrefixAlone"`  //要删除的独有前缀。removePrefixCommon + removePrefixAlone必须和hack/config.yaml内removePrefix保持一致，示例：auth_
	// ModuleDir    string `json:"moduleDir"`    //模块目录，支持多目录。必须和hack/config.yaml内daoPath的后面部分保持一致，示例：auth，xxxx/user
	CommonName   string `json:"commonName"`   //公共名称，将同时在swagger文档Tag标签，权限菜单和权限操作中使用。示例：用户，权限管理/测试
	IsList       bool   `json:"isList" `      //是否生成列表接口(0和no为false，1和yes为true)
	IsCount      bool   `json:"isCount" `     //列表接口是否返回总数
	IsInfo       bool   `json:"isInfo" `      //是否生成详情接口
	IsCreate     bool   `json:"isCreate"`     //是否生成创建接口
	IsUpdate     bool   `json:"isUpdate"`     //是否生成更新接口
	IsDelete     bool   `json:"isDelete"`     //是否生成删除接口
	IsApi        bool   `json:"isApi"`        //是否生成后端接口文件
	IsAuthAction bool   `json:"isAuthAction"` //是否判断操作权限，如是，则同时会生成操作权限
	IsView       bool   `json:"isView"`       //是否生成前端视图文件
	IsCover      bool   `json:"isCover"`      //是否覆盖原文件(设置为true时，建议与git一起使用，防止代码覆盖风险)
}

type myGenTpl struct {
	RemovePrefix              string       //要删除的前缀
	TableRaw                  string       //表名（原始，包含前缀）
	TableCaseKebab            string       //表名（横线，已去除前缀）
	TableCaseSnake            string       //表名（蛇形，已去除前缀）
	TableCaseCamel            string       //表名（大驼峰，已去除前缀）
	TableColumnList           gdb.Result   //字段详情
	FieldList                 []myGenField //字段列表
	ModuleDirCaseKebab        string       //模块目录（横线，/会被保留）
	ModuleDirCaseKebabReplace string       //模块目录（横线，/被替换成.）
	ModuleDirCaseCamel        string       //模块目录（大驼峰，/会被去除）
	LogicStructName           string       //logic层结构体名称，也是权限操作前缀（大驼峰，由ModuleDirCaseCamel+TableCaseCamel组成。命名原因：gf gen service只支持logic单层目录，可能导致service层重名）
	FieldPrimary              string       //主键字段
	FieldDeleted              string       //删除时间字段
	FieldUpdated              string       //更新时间字段
	FieldCreated              string       //创建时间字段
	// 以下属性用于对字段做特殊处理
	HandelLabelList   []string                  //label列表,sql查询时可别名为label的字段（常用于前端my-select或my-cascader等组件，也用户关联表查询时获取）。字段存入按以下优先级：表名去掉前缀 + Name > 主键去掉ID + Name > Name > Title > Phone > Email > Account > Nickname
	HandlePasswordMap map[string]handlePassword //password|passwd,salt同时存在时，需特殊处理
	HandlePid         struct {                  //pid,level,idPath|id_path同时存在时，需特殊处理
		IsCoexist   bool   //是否同时存在
		PidField    string //父级字段
		LevelField  string //层级字段
		IdPathField string //层级路径字段
		SortField   string //排序字段
	}
	HandleIdRelTableMap map[string]myGenTpl //id后缀字段关联表信息

	// 以下字段用于对某些表字段做特殊处理
	RelTableMap     map[string]relTableItem     //一对一关联表。id后缀字段，能确定关联表时，会自动生成联表查询代码
	RelTableManyMap map[string]relTableManyItem //一对多关联表。关联表命名必须table_rel_to_table，能确定关联表时，会自动生成联表查询代码
}

type myGenFieldType = uint
type myGenFieldTypeName = string

const (
	//用于结构体中，需从1开始，否则结构体会默认0，即Int
	TypeInt       myGenFieldType = iota + 1 // `int等类型`
	TypeIntU                                // `int等类型（unsigned）`
	TypeFloat                               // `float等类型（unsigned）`
	TypeFloatU                              // `float等类型（unsigned）`
	TypeVarchar                             // `varchar类型`
	TypeChar                                // `char类型`
	TypeText                                // `text类型`
	TypeJson                                // `json类型`
	TypeTimestamp                           // `timestamp类型`
	TypeDatetime                            // `datetime类型`
	TypeDate                                // `date类型`

	TypeNameDeleted        myGenFieldTypeName = `软删除字段`
	TypeNameUpdated        myGenFieldTypeName = `更新时间字段`
	TypeNameCreated        myGenFieldTypeName = `创建时间字段`
	TypeNamePri            myGenFieldTypeName = `主键`
	TypeNamePriAutoInc     myGenFieldTypeName = `主键（自增）`
	TypeNamePid            myGenFieldTypeName = `命名：pid；		类型：int等类型；	注意：pid,level,idPath|id_path同时存在时，有特殊处理`
	TypeNameLevel          myGenFieldTypeName = `命名：level；	类型：int等类型；	注意：pid,level,idPath|id_path同时存在时，(才)有特殊处理`
	TypeNameIdPath         myGenFieldTypeName = `命名：idPath|id_path；	类型：varchar或text；	注意：pid,level,idPath|id_path同时存在时，(才)有特殊处理`
	TypeNameSort           myGenFieldTypeName = `命名：sort；	类型：int等类型；	注意：pid,level,idPath|id_path|sort同时存在时，(才)有特殊处理`
	TypeNamePasswordSuffix myGenFieldTypeName = `命名：password,passwd后缀；		类型：char(32)；`
	TypeNameSaltSuffix     myGenFieldTypeName = `命名：salt后缀；	类型：char；	注意：password,salt同时存在时，有特殊处理`
	TypeNameNameSuffix     myGenFieldTypeName = `命名：name后缀；	类型：varchar；`
	TypeNameCodeSuffix     myGenFieldTypeName = `命名：code后缀；	类型：varchar；`
	TypeNamePhoneSuffix    myGenFieldTypeName = `命名：phone,mobile后缀；	类型：varchar；`
	TypeNameUrlSuffix      myGenFieldTypeName = `命名：url,link后缀；	类型：varchar；`
	TypeNameIpSuffix       myGenFieldTypeName = `命名：IP后缀；	类型：varchar；`
	TypeNameIdSuffix       myGenFieldTypeName = `命名：id后缀；	类型：int等类型；`
	TypeNameSortSuffix     myGenFieldTypeName = `命名：sort,weight等后缀；	类型：int等类型；`
	TypeNameStatusSuffix   myGenFieldTypeName = `命名：status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）`
	TypeNameIsPrefix       myGenFieldTypeName = `命名：is_前缀；		类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）`
	TypeNameStartPrefix    myGenFieldTypeName = `命名：start_前缀；	类型：timestamp或datetime或date；`
	TypeNameEndPrefix      myGenFieldTypeName = `命名：end_前缀；	类型：timestamp或datetime或date；`
	TypeNameRemarkSuffix   myGenFieldTypeName = `命名：remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器`
	TypeNameImageSuffix    myGenFieldTypeName = `命名：icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text`
	TypeNameVideoSuffix    myGenFieldTypeName = `命名：video,video_list,videoList,video_arr,videoArr等后缀；		类型：单视频varchar，多视频json或text`
	TypeNameArrSuffix      myGenFieldTypeName = `命名：list,arr等后缀；	类型：json或text；`
)

type myGenField struct {
	// gdb.TableField
	FieldRaw             string             // 字段（原始）
	FieldCaseSnake       string             // 字段（蛇形）
	FieldCaseCamel       string             // 字段（大驼峰）
	FieldCaseSnakeRemove string             // 字段（蛇形。去除_of_后）
	FieldCaseCamelRemove string             // 字段（大驼峰。去除_of_后）
	FieldTypeRaw         string             // 字段类型（原始）
	FieldType            myGenFieldType     // 字段类型（数据类型）
	FieldTypeName        myGenFieldTypeName // 字段类型（命名类型）
	IndexRaw             string             // 索引类型（原始）。PRI, MUL
	Index                string             // 索引类型
	IsNull               bool               // 字段是否可为NULL
	Default              interface{}        // 默认值
	Extra                string             // 扩展信息： auto_increment自动递增
	Comment              string             // 注释（原始）。
	FieldName            string             // 字段名称。由注释解析出来，前端显示用。符号[\n\r.。:：(（]之前的部分或整个注释，将作为字段名称使用）
	FieldDesc            string             // 字段说明。由注释解析出来，API文档用。符号[\n\r]换成` `，"增加转义换成\"
	FieldTip             string             // 字段提示。由注释解析出来，前端提示用。符号[\n\r]换成` `，"增加转义换成\"
	StatusList           [][2]string        // 状态列表。由注释解析出来，前端显示用。多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
	FieldLimitStr        string             // 字符串字段限制。varchar表示最大长度；char表示长度；
	FieldLimitFloat      [2]string          // 浮点数字段限制。第1个表示整数位，第2个表示小数位
}

type handlePassword struct {
	IsCoexist      bool   //是否同时存在
	PasswordField  string //密码字段
	PasswordLength string //密码字段长度
	SaltField      string //加密盐字段
	SaltLength     string //加密盐字段长度
}

type relTableItem struct {
	RemovePrefix            string     //要删除的前缀
	TableRaw                string     //表名（原始，包含前缀）
	RelTableCaseSnake       string     //表名（蛇形，已去除前缀）
	RelTableCaseCamel       string     //表名（大驼峰，已去除前缀）
	RelTableCaseCamelLower  string     //表名（小驼峰，已去除前缀）
	TableColumnList         gdb.Result //表字段详情
	RelDaoDir               string     //关联表dao层目录
	RelDaoDirCaseCamel      string     //关联表dao层目录（大驼峰，/会被去除）
	RelDaoDirCaseCamelLower string     //关联表dao层目录（小驼峰，/会被保留）
	IsSameDir               bool       //关联表dao层是否与当前生成dao层在相同目录下
	RelTableField           string     //关联表字段
	RelTableFieldName       string     //关联表字段名称
	IsRedundRelNameField    bool       //当前表是否冗余关联表字段
	RelSuffix               string     //关联表字段后缀（原始，大驼峰或蛇形）。字段含[_of_]时，_of_及之后的部分。示例：userIdOfSend对应OfSend；user_id_of_send对应_of_send
	RelSuffixCaseCamel      string     //关联表字段后缀（大驼峰）。字段含[_of_]时，_of_及其之后的部分。示例：userIdOfSend和user_id_of_send都对应OfSend
	RelSuffixCaseSnake      string     //关联表字段后缀（蛇形）。字段含[_of_]时，_of_及其之后的部分。示例：userIdOfSend和user_id_of_send都对应_of_send
	RelTableIsExistPidField bool       //关联表是否pid字段。前端Query和Save视图组件则使用my-cascader组件，否则使用my-select组件
}

// TODO 一对多
type relTableManyItem struct {
	TableRaw                string     //表名（原始，包含前缀）
	RelTableCaseSnake       string     //表名（蛇形，已去除前缀）
	RelTableCaseCamel       string     //表名（大驼峰，已去除前缀）
	RelTableCaseCamelLower  string     //表名（小驼峰，已去除前缀）
	TableColumnList         gdb.Result //表字段详情
	RelDaoDir               string     //关联表dao层目录
	RelDaoDirCaseCamel      string     //关联表dao层目录（大驼峰，/会被去除）
	RelDaoDirCaseCamelLower string     //关联表dao层目录（小驼峰，/会被保留）
	IsSameDir               bool       //关联表dao层是否与当前生成dao层在相同目录下
	RelTableField           string     //关联表字段
	RelTableFieldName       string     //关联表字段名称
	IsRedundRelNameField    bool       //当前表是否冗余关联表字段
	RelSuffix               string     //关联表字段后缀（原始，大驼峰或蛇形）。字段含[_of_]时，_of_及之后的部分。示例：userIdOfSend对应OfSend；user_id_of_send对应_of_send
	RelSuffixCaseCamel      string     //关联表字段后缀（大驼峰）。字段含[_of_]时，_of_及其之后的部分。示例：userIdOfSend和user_id_of_send都对应OfSend
	RelSuffixCaseSnake      string     //关联表字段后缀（蛇形）。字段含[_of_]时，_of_及其之后的部分。示例：userIdOfSend和user_id_of_send都对应_of_send
	RelTableIsExistPidField bool       //关联表是否pid字段。前端Query和Save视图组件则使用my-cascader组件，否则使用my-select组件
}

// 参数处理
func (myGenThis *myGenHandler) init(parser *gcmd.Parser) {
	optionMap := parser.GetOptAll()
	option := myGenOption{}
	gconv.Struct(optionMap, &option)
	defer func() {
		myGenThis.option = option
	}()

	// 场景标识
	if option.SceneCode == `` {
		option.SceneCode = gcmd.Scan("> 请输入场景标识:\n")
	}
	for {
		if option.SceneCode != `` {
			sceneInfo, _ := daoAuth.Scene.CtxDaoModel(myGenThis.ctx).Filter(daoAuth.Scene.Columns().SceneCode, option.SceneCode).One()
			if !sceneInfo.IsEmpty() {
				myGenThis.sceneId = sceneInfo[daoAuth.Scene.Columns().SceneId].Uint()
				myGenThis.sceneName = sceneInfo[daoAuth.Scene.Columns().SceneName].String()
				break
			}
		}
		option.SceneCode = gcmd.Scan("> 场景标识不存在，请重新输入:\n")
	}
	// db分组
	if option.DbGroup == `` {
		option.DbGroup = gcmd.Scan("> 请输入db分组，默认(default):\n")
		if option.DbGroup == `` {
			option.DbGroup = `default`
		}
	}
	for {
		err := g.Try(myGenThis.ctx, func(ctx context.Context) {
			myGenThis.db = g.DB(option.DbGroup)
			myGenThis.dbLink = gconv.String(gconv.SliceMap(g.Cfg().MustGet(myGenThis.ctx, `database`).MapStrAny()[option.DbGroup])[0][`link`])
		})
		if err == nil {
			break
		}
		option.DbGroup = gcmd.Scan("> db分组不存在，请重新输入，默认(default):\n")
		if option.DbGroup == `` {
			option.DbGroup = `default`
		}
	}
	// db表
	myGenThis.tableArr, _ = myGenThis.db.Tables(myGenThis.ctx)
	if option.DbTable == `` {
		option.DbTable = gcmd.Scan("> 请输入db表:\n")
	}
	for {
		if option.DbTable != `` && garray.NewStrArrayFrom(myGenThis.tableArr).Contains(option.DbTable) {
			break
		}
		option.DbTable = gcmd.Scan("> db表不存在，请重新输入:\n")
	}
	// 要删除的共有前缀
	if _, ok := optionMap[`removePrefixCommon`]; !ok {
		option.RemovePrefixCommon = gcmd.Scan("> 请输入要删除的共有前缀，默认(空):\n")
	}
	for {
		if option.RemovePrefixCommon == `` || gstr.Pos(option.DbTable, option.RemovePrefixCommon) == 0 {
			break
		}
		option.RemovePrefixCommon = gcmd.Scan("> 要删除的共有前缀不存在，请重新输入，默认(空):\n")
	}
	// 要删除的独有前缀
	if _, ok := optionMap[`removePrefixAlone`]; !ok {
		option.RemovePrefixAlone = gcmd.Scan("> 请输入要删除的独有前缀，默认(空):\n")
	}
	for {
		if option.RemovePrefixAlone == `` || gstr.Pos(option.DbTable, option.RemovePrefixCommon+option.RemovePrefixAlone) == 0 {
			break
		}
		option.RemovePrefixAlone = gcmd.Scan("> 要删除的独有前缀不存在，请重新输入，默认(空):\n")
	}
	// 公共名称，将同时在swagger文档Tag标签，权限菜单和权限操作中使用。示例：场景
	for {
		if option.CommonName != `` {
			break
		}
		option.CommonName = gcmd.Scan("> 请输入公共名称，将同时在swagger文档Tag标签，权限菜单和权限操作中使用:\n")
	}
noAllRestart:
	// 是否生成列表接口
	isList, ok := optionMap[`isList`]
	if !ok {
		isList = gcmd.Scan("> 是否生成列表接口，默认(yes):\n")
	}
isListEnd:
	for {
		switch isList {
		case ``, `1`, `yes`:
			option.IsList = true
			break isListEnd
		case `0`, `no`:
			option.IsList = false
			break isListEnd
		default:
			isList = gcmd.Scan("> 输入错误，请重新输入，是否生成列表接口，默认(yes):\n")
		}
	}
	// 列表接口是否返回总数
	isCount, ok := optionMap[`isCount`]
	if !ok {
		isCount = gcmd.Scan("> 列表接口是否返回总数，默认(yes):\n")
	}
isCountEnd:
	for {
		switch isCount {
		case ``, `1`, `yes`:
			option.IsCount = true
			break isCountEnd
		case `0`, `no`:
			option.IsCount = false
			break isCountEnd
		default:
			isCount = gcmd.Scan("> 输入错误，请重新输入，列表接口是否返回总数，默认(yes):\n")
		}
	}
	// 是否生成详情接口
	isInfo, ok := optionMap[`isInfo`]
	if !ok {
		isInfo = gcmd.Scan("> 是否生成详情接口，默认(yes):\n")
	}
isInfoEnd:
	for {
		switch isInfo {
		case ``, `1`, `yes`:
			option.IsInfo = true
			break isInfoEnd
		case `0`, `no`:
			option.IsInfo = false
			break isInfoEnd
		default:
			isInfo = gcmd.Scan("> 输入错误，请重新输入，是否生成详情接口，默认(yes):\n")
		}
	}
	// 是否生成创建接口
	isCreate, ok := optionMap[`isCreate`]
	if !ok {
		isCreate = gcmd.Scan("> 是否生成创建接口，默认(yes):\n")
	}
isCreateEnd:
	for {
		switch isCreate {
		case ``, `1`, `yes`:
			option.IsCreate = true
			break isCreateEnd
		case `0`, `no`:
			option.IsCreate = false
			break isCreateEnd
		default:
			isCreate = gcmd.Scan("> 输入错误，请重新输入，是否生成创建接口，默认(yes):\n")
		}
	}
	// 是否生成更新接口
	isUpdate, ok := optionMap[`isUpdate`]
	if !ok {
		isUpdate = gcmd.Scan("> 是否生成更新接口，默认(yes):\n")
	}
isUpdateEnd:
	for {
		switch isUpdate {
		case ``, `1`, `yes`:
			option.IsUpdate = true
			break isUpdateEnd
		case `0`, `no`:
			option.IsUpdate = false
			break isUpdateEnd
		default:
			isUpdate = gcmd.Scan("> 输入错误，请重新输入，是否生成更新接口，默认(yes):\n")
		}
	}
	// 是否生成删除接口
	isDelete, ok := optionMap[`isDelete`]
	if !ok {
		isDelete = gcmd.Scan("> 是否生成删除接口，默认(yes):\n")
	}
isDeleteEnd:
	for {
		switch isDelete {
		case ``, `1`, `yes`:
			option.IsDelete = true
			break isDeleteEnd
		case `0`, `no`:
			option.IsDelete = false
			break isDeleteEnd
		default:
			isDelete = gcmd.Scan("> 输入错误，请重新输入，是否生成删除接口，默认(yes):\n")
		}
	}
	if !(option.IsList || option.IsInfo || option.IsCreate || option.IsUpdate || option.IsDelete) {
		fmt.Println(`请重新选择生成哪些接口，不能全是no！`)
		goto noAllRestart
	}
	// 是否生成后端接口文件
	isApi, ok := optionMap[`isApi`]
	if !ok {
		isApi = gcmd.Scan("> 是否生成后端接口文件，默认(yes):\n")
	}
isApiEnd:
	for {
		switch isApi {
		case ``, `1`, `yes`:
			option.IsApi = true
			break isApiEnd
		case `0`, `no`:
			option.IsApi = false
			break isApiEnd
		default:
			isApi = gcmd.Scan("> 输入错误，请重新输入，是否生成后端接口文件，默认(yes):\n")
		}
	}
	if option.IsApi {
		// 是否判断操作权限，如是，则同时会生成操作权限
		isAuthAction, ok := optionMap[`isAuthAction`]
		if !ok {
			isAuthAction = gcmd.Scan("> 是否判断操作权限，如是，则同时会生成操作权限，默认(yes):\n")
		}
	isAuthActionEnd:
		for {
			switch isAuthAction {
			case ``, `1`, `yes`:
				option.IsAuthAction = true
				break isAuthActionEnd
			case `0`, `no`:
				option.IsAuthAction = false
				break isAuthActionEnd
			default:
				isAuthAction = gcmd.Scan("> 输入错误，请重新输入，是否判断操作权限，如是，则同时会生成操作权限，默认(yes):\n")
			}
		}
	}
	// 是否生成前端视图文件
	isView, ok := optionMap[`isView`]
	if !ok {
		isView = gcmd.Scan("> 是否生成前端视图文件，默认(yes):\n")
	}
isViewEnd:
	for {
		switch isView {
		case ``, `1`, `yes`:
			option.IsView = true
			break isViewEnd
		case `0`, `no`:
			option.IsView = false
			break isViewEnd
		default:
			isView = gcmd.Scan("> 输入错误，请重新输入，是否生成前端视图文件，默认(yes):\n")
		}
	}
	// 是否覆盖原文件
	isCover, ok := optionMap[`isCover`]
	if !ok {
		isCover = gcmd.Scan("> 是否覆盖原文件(设置为yes时，建议与git一起使用，防止代码覆盖风险)，默认(no):\n")
	}
isCoverEnd:
	for {
		switch isCover {
		case `1`, `yes`:
			option.IsCover = true
			break isCoverEnd
		case ``, `0`, `no`:
			option.IsCover = false
			break isCoverEnd
		default:
			isCover = gcmd.Scan("> 输入错误，请重新输入，是否覆盖原文件(设置为yes时，建议与git一起使用，防止代码覆盖风险)，默认(no):\n")
		}
	}
}

// 模板参数处理
func (myGenThis *myGenHandler) setTpl(table, removePrefixAlone string) (tpl myGenTpl) {
	ctx := myGenThis.ctx

	tableColumnList, _ := myGenThis.db.GetAll(ctx, `SHOW FULL COLUMNS FROM `+table)
	removePrefix := myGenThis.option.RemovePrefixCommon + removePrefixAlone
	tableReomve := gstr.Replace(table, removePrefix, ``, 1)
	tpl = myGenTpl{
		RemovePrefix:      removePrefix,
		TableRaw:          table,
		TableCaseKebab:    gstr.CaseKebab(tableReomve),
		TableCaseSnake:    gstr.CaseSnake(tableReomve),
		TableCaseCamel:    gstr.CaseCamel(tableReomve),
		TableColumnList:   tableColumnList,
		HandlePasswordMap: map[string]handlePassword{},
		RelTableMap:       map[string]relTableItem{},
	}

	logicStructName := gstr.TrimLeftStr(table, myGenThis.option.RemovePrefixCommon, 1)
	moduleDirCaseKebab := gstr.CaseKebab(logicStructName)
	moduleDirCaseCamel := gstr.CaseCamel(logicStructName)
	if removePrefixAlone != `` {
		moduleDirCaseKebab = gstr.CaseKebab(gstr.Trim(removePrefixAlone, `_`))
		moduleDirCaseCamel = gstr.CaseCamel(removePrefixAlone)
	}
	if myGenThis.option.DbGroup != `default` {
		logicStructName = myGenThis.option.DbGroup + `_` + logicStructName
		moduleDirCaseKebab = gstr.CaseKebab(myGenThis.option.DbGroup) + `/` + moduleDirCaseKebab
		moduleDirCaseCamel = gstr.CaseCamel(myGenThis.option.DbGroup) + moduleDirCaseCamel
	}
	tpl.LogicStructName = gstr.CaseCamel(logicStructName)
	tpl.ModuleDirCaseKebab = moduleDirCaseKebab
	tpl.ModuleDirCaseKebabReplace = gstr.Replace(moduleDirCaseKebab, `/`, `.`)
	tpl.ModuleDirCaseCamel = moduleDirCaseCamel

	fieldList := make([]myGenField, len(tpl.TableColumnList))

	fieldArr := make([]string, len(tpl.TableColumnList))
	fieldCaseCamelArr := make([]string, len(tpl.TableColumnList))
	for k, v := range tpl.TableColumnList {
		fieldTmp := myGenField{
			FieldRaw:     v[`Field`].String(),
			FieldTypeRaw: v[`Type`].String(),
			IndexRaw:     v[`Key`].String(),
			IsNull:       v[`Null`].Bool(),
			Default:      v[`Default`].Val(),
			Extra:        v[`Extra`].String(),
			Comment:      v[`Comment`].String(),
		}
		fieldTmp.FieldCaseSnake = gstr.CaseSnake(fieldTmp.FieldRaw)
		fieldTmp.FieldCaseCamel = gstr.CaseCamel(fieldTmp.FieldRaw)
		fieldTmp.FieldCaseSnakeRemove = gstr.Split(fieldTmp.FieldCaseSnake, `_of_`)[0]
		fieldTmp.FieldCaseCamelRemove = gstr.CaseCamel(fieldTmp.FieldCaseSnakeRemove)

		tmpFieldName, _ := gregex.MatchString(`[^\n\r\.。:：\(（]*`, fieldTmp.Comment)
		fieldTmp.FieldName = gstr.Trim(tmpFieldName[0])
		fieldTmp.FieldDesc = gstr.Trim(gstr.ReplaceByArray(fieldTmp.Comment, g.SliceStr{
			"\n", ` `,
			"\r", ` `,
			`"`, `\"`,
		}))
		tmpFieldTip := gstr.Replace(fieldTmp.FieldDesc, fieldTmp.FieldName, ``, 1)
		tmp, _ := gregex.MatchString(`\n\r\.。:：\(（`, fieldTmp.Comment)
		if len(tmp) > 0 {
			gstr.TrimLeft(tmpFieldTip, tmp[0])
		}
		for _, v := range []string{"\n", "\r", `.`, `。`, `:`, `：`, `(`, `（`, `)`, `）`, ` `, `,`, `，`, `;`, `；`} {
			tmpFieldTip = gstr.Trim(tmpFieldTip, v)
		}
		fieldTmp.FieldTip = gstr.ReplaceByArray(tmpFieldTip, g.SliceStr{
			`\"`, `"`,
			`}`, `' + "{'}'}" + '`,
			`{"`, `' + "{'{'}" + '"`,
		})

		tmpFieldLimitStr, _ := gregex.MatchString(`.*\((\d*)\)`, fieldTmp.FieldTypeRaw)
		if len(tmpFieldLimitStr) > 1 {
			fieldTmp.FieldLimitStr = tmpFieldLimitStr[1]
		}
		tmpFieldLimitFloat, _ := gregex.MatchString(`.*\((\d*),(\d*)\)`, fieldTmp.FieldTypeRaw)
		if len(tmpFieldLimitFloat) < 3 {
			tmpFieldLimitFloat = []string{``, `10`, `2`}
		}
		fieldTmp.FieldLimitFloat = [2]string{tmpFieldLimitFloat[1], tmpFieldLimitFloat[2]}

		/*--------确定字段数据类型 开始--------*/
		if gstr.Pos(fieldTmp.FieldTypeRaw, `int`) != -1 && gstr.Pos(fieldTmp.FieldTypeRaw, `point`) == -1 { //int等类型
			fieldTmp.FieldType = TypeInt
			if gstr.Pos(fieldTmp.FieldTypeRaw, `unsigned`) != -1 {
				fieldTmp.FieldType = TypeIntU
			}
		} else if gstr.Pos(fieldTmp.FieldTypeRaw, `decimal`) != -1 || gstr.Pos(fieldTmp.FieldTypeRaw, `double`) != -1 || gstr.Pos(fieldTmp.FieldTypeRaw, `float`) != -1 { //float类型
			fieldTmp.FieldType = TypeFloat
			if gstr.Pos(fieldTmp.FieldTypeRaw, `unsigned`) != -1 {
				fieldTmp.FieldType = TypeFloatU
			}
		} else if gstr.Pos(fieldTmp.FieldTypeRaw, `varchar`) != -1 { //varchar类型
			fieldTmp.FieldType = TypeVarchar
		} else if gstr.Pos(fieldTmp.FieldTypeRaw, `char`) != -1 { //char类型
			fieldTmp.FieldType = TypeChar
		} else if gstr.Pos(fieldTmp.FieldTypeRaw, `text`) != -1 { //text类型
			fieldTmp.FieldType = TypeText
		} else if gstr.Pos(fieldTmp.FieldTypeRaw, `json`) != -1 { //json类型
			fieldTmp.FieldType = TypeJson

		} else if gstr.Pos(fieldTmp.FieldTypeRaw, `timestamp`) != -1 || gstr.Pos(fieldTmp.FieldTypeRaw, `date`) != -1 { //timestamp或datetime或date类型
			fieldTmp.FieldType = TypeTimestamp
			if gstr.Pos(fieldTmp.FieldTypeRaw, `datetime`) == -1 {
				fieldTmp.FieldType = TypeDatetime
			} else if gstr.Pos(fieldTmp.FieldTypeRaw, `date`) != -1 {
				fieldTmp.FieldType = TypeDate
			}
		}
		/*--------确定字段数据类型 结束--------*/

		/*--------确定字段命名类型 开始--------*/
		fieldSplitArr := gstr.Split(fieldTmp.FieldCaseSnakeRemove, `_`)
		fieldPrefix := fieldSplitArr[0]
		fieldSuffix := fieldSplitArr[len(fieldSplitArr)-1]
		if garray.NewStrArrayFrom([]string{`DeletedAt`, `DeleteAt`, `DeletedTime`, `DeleteTime`}).Contains(fieldTmp.FieldCaseCamel) {
			fieldTmp.FieldTypeName = TypeNameDeleted
		} else if garray.NewStrArrayFrom([]string{`UpdatedAt`, `UpdateAt`, `UpdatedTime`, `UpdateTime`}).Contains(fieldTmp.FieldCaseCamel) {
			fieldTmp.FieldTypeName = TypeNameUpdated
		} else if garray.NewStrArrayFrom([]string{`CreatedAt`, `CreateAt`, `CreatedTime`, `CreateTime`}).Contains(fieldTmp.FieldCaseCamel) {
			fieldTmp.FieldTypeName = TypeNameCreated
		} else if fieldTmp.IndexRaw == `PRI` {
			fieldTmp.FieldTypeName = TypeNamePri
			if fieldTmp.Extra == `auto_increment` {
				fieldTmp.FieldTypeName = TypeNamePriAutoInc
			}
		} else if garray.NewFrom([]interface{}{TypeVarchar, TypeText}).Contains(fieldTmp.FieldType) && fieldTmp.FieldCaseCamel == `IdPath` { //idPath|id_path
			fieldTmp.FieldTypeName = TypeNameIdPath
		} else if garray.NewFrom([]interface{}{TypeInt, TypeIntU, TypeVarchar, TypeChar}).Contains(fieldTmp.FieldType) && garray.NewStrArrayFrom([]string{`status`, `type`, `method`, `pos`, `position`, `gender`}).Contains(fieldSuffix) { //status,type,method,pos,position,gender等后缀
			fieldTmp.FieldTypeName = TypeNameStatusSuffix
		} else if garray.NewFrom([]interface{}{TypeVarchar, TypeText, TypeJson}).Contains(fieldTmp.FieldType) && (garray.NewStrArrayFrom([]string{`icon`, `cover`, `avatar`, `img`, `image`}).Contains(fieldSuffix) || gstr.SubStr(fieldTmp.FieldCaseCamelRemove, -7) == `ImgList` || gstr.SubStr(fieldTmp.FieldCaseCamelRemove, -6) == `ImgArr` || gstr.SubStr(fieldTmp.FieldCaseCamelRemove, -9) == `ImageList` || gstr.SubStr(fieldTmp.FieldCaseCamelRemove, -8) == `ImageArr`) { //icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀
			fieldTmp.FieldTypeName = TypeNameImageSuffix
		} else if garray.NewFrom([]interface{}{TypeVarchar, TypeText, TypeJson}).Contains(fieldTmp.FieldType) && (garray.NewStrArrayFrom([]string{`video`}).Contains(fieldSuffix) || gstr.SubStr(fieldTmp.FieldCaseCamelRemove, -9) == `VideoList` || gstr.SubStr(fieldTmp.FieldCaseCamelRemove, -8) == `VideoArr`) { //video,video_list,videoList,video_arr,videoArr等后缀
			fieldTmp.FieldTypeName = TypeNameVideoSuffix
		} else if garray.NewFrom([]interface{}{TypeText, TypeJson}).Contains(fieldTmp.FieldType) && garray.NewStrArrayFrom([]string{`list`, `arr`}).Contains(fieldSuffix) { //list,arr等后缀
			fieldTmp.FieldTypeName = TypeNameArrSuffix
		} else if garray.NewFrom([]interface{}{TypeVarchar, TypeText}).Contains(fieldTmp.FieldType) && garray.NewStrArrayFrom([]string{`remark`, `desc`, `msg`, `message`, `intro`, `content`}).Contains(fieldSuffix) { //remark,desc,msg,message,intro,content后缀
			fieldTmp.FieldTypeName = TypeNameRemarkSuffix
		} else if fieldTmp.FieldType == TypeVarchar { //varchar类型
			if garray.NewStrArrayFrom([]string{`name`}).Contains(fieldSuffix) { //name后缀
				fieldTmp.FieldTypeName = TypeNameNameSuffix
			} else if garray.NewStrArrayFrom([]string{`code`}).Contains(fieldSuffix) { //code后缀
				fieldTmp.FieldTypeName = TypeNameCodeSuffix
			} else if garray.NewStrArrayFrom([]string{`phone`, `mobile`}).Contains(fieldSuffix) { //phone,mobile后缀
				fieldTmp.FieldTypeName = TypeNamePhoneSuffix
			} else if garray.NewStrArrayFrom([]string{`url`, `link`}).Contains(fieldSuffix) { //url,link后缀
				fieldTmp.FieldTypeName = TypeNameUrlSuffix
			} else if garray.NewStrArrayFrom([]string{`ip`}).Contains(fieldSuffix) { //IP后缀
				fieldTmp.FieldTypeName = TypeNameIpSuffix
			}
		} else if fieldTmp.FieldType == TypeChar { //char类型
			if garray.NewStrArrayFrom([]string{`password`, `passwd`}).Contains(fieldSuffix) && fieldTmp.FieldTypeRaw == `char(32)` { //password,passwd后缀
				fieldTmp.FieldTypeName = TypeNamePasswordSuffix
			} else if garray.NewStrArrayFrom([]string{`salt`}).Contains(fieldSuffix) { //salt后缀
				fieldTmp.FieldTypeName = TypeNameSaltSuffix
			}
		} else if garray.NewFrom([]interface{}{TypeInt, TypeIntU}).Contains(fieldTmp.FieldType) { //int等类型
			if fieldTmp.FieldRaw == `pid` { //pid
				fieldTmp.FieldTypeName = TypeNamePid
			} else if fieldTmp.FieldRaw == `level` { //level
				fieldTmp.FieldTypeName = TypeNameLevel
			} else if garray.NewStrArrayFrom([]string{`sort`, `weight`}).Contains(fieldSuffix) { //sort,weight等后缀
				fieldTmp.FieldTypeName = TypeNameSortSuffix
				if fieldTmp.FieldRaw == `sort` { //sort
					fieldTmp.FieldTypeName = TypeNameSort
				}
			} else if garray.NewStrArrayFrom([]string{`id`}).Contains(fieldSuffix) { //id后缀
				fieldTmp.FieldTypeName = TypeNameIdSuffix
			} else if garray.NewStrArrayFrom([]string{`is`}).Contains(fieldPrefix) { //is_前缀
				fieldTmp.FieldTypeName = TypeNameIsPrefix
			}
		} else if garray.NewFrom([]interface{}{TypeTimestamp, TypeDatetime, TypeDate}).Contains(fieldTmp.FieldType) { //timestamp或datetime或date类型
			if garray.NewStrArrayFrom([]string{`start`}).Contains(fieldPrefix) { //start_前缀
				fieldTmp.FieldTypeName = TypeNameStartPrefix
			} else if garray.NewStrArrayFrom([]string{`end`}).Contains(fieldPrefix) { //end_前缀
				fieldTmp.FieldTypeName = TypeNameEndPrefix
			}
		}
		/*--------确定字段命名类型 结束--------*/

		/*--------特殊处理字段解析 开始--------*/
		switch fieldTmp.FieldTypeName {
		case TypeNameDeleted: // 软删除字段
			tpl.FieldDeleted = fieldTmp.FieldRaw
		case TypeNameUpdated: // 更新时间字段
			tpl.FieldUpdated = fieldTmp.FieldRaw
		case TypeNameCreated: // 创建时间字段
			tpl.FieldCreated = fieldTmp.FieldRaw
		case TypeNamePri: // 主键
		case TypeNamePriAutoInc: // 主键（自增）
			tpl.FieldPrimary = fieldTmp.FieldRaw
		case TypeNamePid: // pid；		类型：int等类型；	注意：pid,level,idPath|id_path同时存在时，有特殊处理
			tpl.HandlePid.PidField = fieldTmp.FieldRaw
		case TypeNameLevel: // level；	类型：int等类型；	注意：pid,level,idPath|id_path同时存在时，(才)有特殊处理
			tpl.HandlePid.LevelField = fieldTmp.FieldRaw
		case TypeNameIdPath: // idPath|id_path；	类型：varchar或text；	注意：pid,level,idPath|id_path同时存在时，(才)有特殊处理
			tpl.HandlePid.IdPathField = fieldTmp.FieldRaw
		case TypeNameSort: // sort；	类型：int等类型；	注意：pid,level,idPath|id_path|sort同时存在时，(才)有特殊处理
			tpl.HandlePid.SortField = fieldTmp.FieldRaw
		case TypeNamePasswordSuffix: // password,passwd后缀；		类型：char(32)；
			handlePasswordMapKey := myGenThis.genHandlePasswordMapKey(fieldTmp.FieldRaw)
			handlePasswordObj, ok := tpl.HandlePasswordMap[handlePasswordMapKey]
			if ok {
				handlePasswordObj.PasswordField = fieldTmp.FieldRaw
				handlePasswordObj.PasswordLength = fieldTmp.FieldLimitStr
			} else {
				handlePasswordObj = handlePassword{
					PasswordField:  fieldTmp.FieldRaw,
					PasswordLength: fieldTmp.FieldLimitStr,
				}
			}
			tpl.HandlePasswordMap[handlePasswordMapKey] = handlePasswordObj
		case TypeNameSaltSuffix: // salt后缀；	类型：char；	注意：password,salt同时存在时，有特殊处理
			handlePasswordMapKey := myGenThis.genHandlePasswordMapKey(fieldTmp.FieldRaw)
			handlePasswordObj, ok := tpl.HandlePasswordMap[handlePasswordMapKey]
			if ok {
				handlePasswordObj.SaltField = fieldTmp.FieldRaw
				handlePasswordObj.SaltLength = fieldTmp.FieldLimitStr
			} else {
				handlePasswordObj = handlePassword{
					SaltField:  fieldTmp.FieldRaw,
					SaltLength: fieldTmp.FieldLimitStr,
				}
			}
			tpl.HandlePasswordMap[handlePasswordMapKey] = handlePasswordObj
		case TypeNameNameSuffix: // name后缀；	类型：varchar；
		case TypeNameCodeSuffix: // code后缀；	类型：varchar；
		case TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
		case TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
		case TypeNameIpSuffix: // IP后缀；	类型：varchar；
		case TypeNameIdSuffix: // id后缀；	类型：int等类型；
			tpl.RelTableMap[fieldTmp.FieldRaw] = myGenThis.genRelTable(tpl, fieldTmp.FieldRaw, fieldTmp.FieldName)
		case TypeNameSortSuffix: // sort,weight等后缀；	类型：int等类型；
		case TypeNameStatusSuffix: // status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
			isStr := false
			if garray.NewFrom([]interface{}{TypeVarchar, TypeChar}).Contains(fieldTmp.FieldType) {
				isStr = true
			}
			fieldTmp.StatusList = myGenThis.genStatusList(fieldTmp.FieldDesc, isStr)
		case TypeNameIsPrefix: // is_前缀；		类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）
		case TypeNameStartPrefix: // start_前缀；	类型：timestamp或datetime或date；
		case TypeNameEndPrefix: // end_前缀；	类型：timestamp或datetime或date；
		case TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
		case TypeNameImageSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text
		case TypeNameVideoSuffix: // video,video_list,videoList,video_arr,videoArr等后缀；		类型：单视频varchar，多视频json或text
		case TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
		}
		/*--------特殊处理字段解析 结束--------*/

		fieldList[k] = fieldTmp
		fieldArr[k] = fieldTmp.FieldRaw
		fieldCaseCamelArr[k] = fieldTmp.FieldCaseCamel
	}

	fieldCaseCamelArrG := garray.NewStrArrayFrom(fieldCaseCamelArr)

	// TODO 可优化Phone，Account等为后缀也算
	// label列表,sql查询时可别名为label的字段（常用于前端my-select或my-cascader等组件，也用户关联表查询时获取）。字段存入按以下优先级：表名去掉前缀 + Name > 主键去掉ID  + Name > Name > Title > Phone > Email > Account > Nickname
	labelList := []string{tpl.TableCaseCamel + `Name`}
	tmpLabel := gstr.SubStr(gstr.CaseCamel(tpl.FieldPrimary), 0, -2) + `Name`
	if tmpLabel != labelList[0] && tmpLabel != `Name` {
		labelList = append(labelList, tpl.TableCaseCamel+`Name`)
	}
	labelList = append(labelList, `Name`, `Title`, `Phone`, `Email`, `Account`, `Nickname`)
	tpl.HandelLabelList = []string{}
	for _, v := range labelList {
		if index := fieldCaseCamelArrG.Search(v); index != -1 {
			tpl.HandelLabelList = append(tpl.HandelLabelList, fieldArr[index])
		}
	}

	for k, v := range tpl.RelTableMap {
		if garray.NewStrArrayFrom(fieldArr).Contains(v.RelTableField + v.RelSuffix) {
			v.IsRedundRelNameField = true
			tpl.RelTableMap[k] = v
		}
	}

	for k, v := range tpl.HandlePasswordMap {
		if v.PasswordField != `` && v.SaltField != `` {
			v.IsCoexist = true
			tpl.HandlePasswordMap[k] = v
		}
	}

	if tpl.HandlePid.PidField != `` && tpl.HandlePid.LevelField != `` && tpl.HandlePid.IdPathField != `` {
		tpl.HandlePid.IsCoexist = true
	}
	return
}

// dao层存在时，增加或修改部分字段的解析代码
func (myGenThis *myGenHandler) genDao() {
	overwriteDao := `false`
	if myGenThis.option.IsCover {
		overwriteDao = `true`
	}
	myGenThis.command(`当前表dao生成`, true, ``,
		`gf`, `gen`, `dao`,
		`--link`, myGenThis.dbLink,
		`--group`, myGenThis.option.DbGroup,
		`--removePrefix`, myGenThis.option.RemovePrefixCommon+myGenThis.option.RemovePrefixAlone,
		`--daoPath`, `dao/`+myGenThis.tpl.ModuleDirCaseKebab,
		`--doPath`, `model/entity/`+myGenThis.tpl.ModuleDirCaseKebab,
		`--entityPath`, `model/entity/`+myGenThis.tpl.ModuleDirCaseKebab,
		`--tables`, myGenThis.option.DbTable,
		`--tplDaoIndexPath`, `resource/gen/gen_dao_template_dao.txt`,
		`--tplDaoInternalPath`, `resource/gen/gen_dao_template_dao_internal.txt`,
		`--overwriteDao`, overwriteDao)

	tpl := myGenThis.tpl

	saveFile := gfile.SelfDir() + `/internal/dao/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseSnake + `.go`
	tplDao := gfile.GetContents(saveFile)

	daoParseInsert := ``
	daoParseInsertBefore := ``
	daoHookInsert := ``
	daoParseUpdate := ``
	daoHookUpdateBefore := ``
	daoHookUpdateAfter := ``
	daoParseField := ``
	daoHookSelect := ``
	daoParseFilter := ``
	daoParseOrder := ``
	daoParseJoin := ``
	daoFunc := ``
	daoImportOtherDao := ``

	handelLabelListLen := len(tpl.HandelLabelList)
	if handelLabelListLen > 0 {
		daoParseFieldTmp := `
			case ` + "`label`" + `:
				m = m.Fields(daoModel.DbTable + ` + "`.`" + ` + daoThis.Columns().` + gstr.CaseCamel(tpl.HandelLabelList[0]) + ` + ` + "` AS `" + ` + v)`
		daoParseFilterTmp := `
			case ` + "`label`" + `:
				m = m.WhereLike(daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + gstr.CaseCamel(tpl.HandelLabelList[0]) + `, ` + "`%`" + `+gconv.String(v)+` + "`%`" + `)`
		if handelLabelListLen > 1 {
			parseFieldStr := "` + daoModel.DbTable + `.` + daoThis.Columns()." + gstr.CaseCamel(tpl.HandelLabelList[handelLabelListLen-1]) + " + `"
			parseFilterStr := "WhereOrLike(daoModel.DbTable+`.`+daoThis.Columns()." + gstr.CaseCamel(tpl.HandelLabelList[handelLabelListLen-1]) + ", `%`+gconv.String(v)+`%`)"
			for i := handelLabelListLen - 2; i >= 0; i-- {
				parseFieldStr = "IF(` + daoModel.DbTable + `.` + daoThis.Columns()." + gstr.CaseCamel(tpl.HandelLabelList[i]) + " + `, ` + daoModel.DbTable + `.` + daoThis.Columns()." + gstr.CaseCamel(tpl.HandelLabelList[i]) + " + `, " + parseFieldStr + ")"
				if i == 0 {
					parseFilterStr = "WhereLike(daoModel.DbTable+`.`+daoThis.Columns()." + gstr.CaseCamel(tpl.HandelLabelList[i]) + ", `%`+gconv.String(v)+`%`)." + parseFilterStr
				} else {
					parseFilterStr = "WhereOrLike(daoModel.DbTable+`.`+daoThis.Columns()." + gstr.CaseCamel(tpl.HandelLabelList[i]) + ", `%`+gconv.String(v)+`%`)." + parseFilterStr
				}
			}
			daoParseFieldTmp = `
			case ` + "`label`" + `:
				m = m.Fields(` + "`" + parseFieldStr + " AS ` + v)"
			daoParseFilterTmp = `
			case ` + "`label`" + `:
				m = m.Where(m.Builder().` + parseFilterStr + `)`
		}
		if gstr.Pos(tplDao, daoParseFieldTmp) == -1 {
			daoParseField += daoParseFieldTmp
		}
		if gstr.Pos(tplDao, daoParseFilterTmp) == -1 {
			daoParseFilter += daoParseFilterTmp
		}
	}

	for _, column := range tpl.TableColumnList {
		field := column[`Field`].String()
		fieldCaseCamel := gstr.CaseCamel(field)
		fieldCaseSnake := gstr.CaseSnake(field)
		fieldCaseSnakeOfRemove := gstr.Split(fieldCaseSnake, `_of_`)[0]
		// fieldCaseCamelOfRemove := gstr.CaseCamel(fieldCaseSnakeOfRemove)
		fieldSplitArr := gstr.Split(fieldCaseSnakeOfRemove, `_`)
		fieldPrefix := fieldSplitArr[0]
		fieldSuffix := fieldSplitArr[len(fieldSplitArr)-1]

		if garray.NewStrArrayFrom([]string{`DeletedAt`, `DeleteAt`, `DeletedTime`, `DeleteTime`}).Contains(fieldCaseCamel) {
		} else if garray.NewStrArrayFrom([]string{`UpdatedAt`, `UpdateAt`, `UpdatedTime`, `UpdateTime`}).Contains(fieldCaseCamel) {
		} else if garray.NewStrArrayFrom([]string{`CreatedAt`, `CreateAt`, `CreatedTime`, `CreateTime`}).Contains(fieldCaseCamel) {
			daoParseFilterTmp := `
			case ` + "`timeRangeStart`" + `:
				m = m.WhereGTE(daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + fieldCaseCamel + `, v)
			case ` + "`timeRangeEnd`" + `:
				m = m.WhereLTE(daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + fieldCaseCamel + `, v)`
			if gstr.Pos(tplDao, daoParseFilterTmp) == -1 {
				daoParseFilter += daoParseFilterTmp
			}
		} else if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` { //主键
		} else if fieldCaseCamel == `IdPath` && (gstr.Pos(column[`Type`].String(), `varchar`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) && tpl.HandlePid.IsCoexist { //idPath|id_path
		} else if gstr.Pos(column[`Type`].String(), `varchar`) != -1 { //varchar类型
			if garray.NewStrArrayFrom([]string{`name`}).Contains(fieldSuffix) { //name后缀
				daoParseFilterTmp := `
			case daoThis.Columns().` + fieldCaseCamel + `:
				m = m.WhereLike(daoModel.DbTable+` + "`.`" + `+k, ` + "`%`" + `+gconv.String(v)+` + "`%`" + `)`
				if gstr.Pos(tplDao, daoParseFilterTmp) == -1 {
					daoParseFilter += daoParseFilterTmp
				}
			}

			if column[`Key`].String() == `UNI` && column[`Null`].Bool() {
				daoParseInsertTmp := `
			case daoThis.Columns().` + fieldCaseCamel + `:
				insertData[k] = v
				if gconv.String(v) == ` + "``" + ` {
					insertData[k] = nil
				}`
				if gstr.Pos(tplDao, daoParseInsertTmp) == -1 {
					daoParseInsert += daoParseInsertTmp
				}
				daoParseUpdateTmp := `
			case daoThis.Columns().` + fieldCaseCamel + `:
				updateData[daoModel.DbTable+` + "`.`" + `+k] = v
				if gconv.String(v) == ` + "``" + ` {
					updateData[daoModel.DbTable+` + "`.`" + `+k] = nil
				}`
				if gstr.Pos(tplDao, daoParseUpdateTmp) == -1 {
					daoParseUpdate += daoParseUpdateTmp
				}
			}
		} else if gstr.Pos(column[`Type`].String(), `char`) != -1 { //char类型
			if garray.NewStrArrayFrom([]string{`password`, `passwd`}).Contains(fieldSuffix) && column[`Type`].String() == `char(32)` { //password,passwd后缀
				daoParseInsertTmp := `
			case daoThis.Columns().` + fieldCaseCamel + `:
				password := gconv.String(v)
				if len(password) != 32 {
					password = gmd5.MustEncrypt(password)
				}`
				daoParseUpdateTmp := `
			case daoThis.Columns().` + fieldCaseCamel + `:
				password := gconv.String(v)
				if len(password) != 32 {
					password = gmd5.MustEncrypt(password)
				}`
				handlePasswordMapKey := myGenThis.genHandlePasswordMapKey(field)
				if tpl.HandlePasswordMap[handlePasswordMapKey].IsCoexist {
					daoParseInsertTmp += `
				salt := grand.S(` + tpl.HandlePasswordMap[handlePasswordMapKey].SaltLength + `)
				insertData[daoThis.Columns().` + gstr.CaseCamel(tpl.HandlePasswordMap[handlePasswordMapKey].SaltField) + `] = salt
				password = gmd5.MustEncrypt(password + salt)`
					daoParseUpdateTmp += `
				salt := grand.S(` + tpl.HandlePasswordMap[handlePasswordMapKey].SaltLength + `)
				updateData[daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + gstr.CaseCamel(tpl.HandlePasswordMap[handlePasswordMapKey].SaltField) + `] = salt
				password = gmd5.MustEncrypt(password + salt)`
				}
				daoParseInsertTmp += `
				insertData[k] = password`
				daoParseUpdateTmp += `
				updateData[daoModel.DbTable+` + "`.`" + `+k] = password`
				if gstr.Pos(tplDao, daoParseInsertTmp) == -1 {
					daoParseInsert += daoParseInsertTmp
				}
				if gstr.Pos(tplDao, daoParseUpdateTmp) == -1 {
					daoParseUpdate += daoParseUpdateTmp
				}
			} else if garray.NewStrArrayFrom([]string{`salt`}).Contains(fieldSuffix) && tpl.HandlePasswordMap[myGenThis.genHandlePasswordMapKey(field)].IsCoexist { //salt后缀
			} else {
				if column[`Key`].String() == `UNI` && column[`Null`].Bool() {
					daoParseInsertTmp := `
			case daoThis.Columns().` + fieldCaseCamel + `:
				insertData[k] = v
				if gconv.String(v) == ` + "``" + ` {
					insertData[k] = nil
				}`
					if gstr.Pos(tplDao, daoParseInsertTmp) == -1 {
						daoParseInsert += daoParseInsertTmp
					}
					daoParseUpdateTmp := `
			case daoThis.Columns().` + fieldCaseCamel + `:
				updateData[daoModel.DbTable+` + "`.`" + `+k] = v
				if gconv.String(v) == ` + "``" + ` {
					updateData[daoModel.DbTable+` + "`.`" + `+k] = nil
				}`
					if gstr.Pos(tplDao, daoParseUpdateTmp) == -1 {
						daoParseUpdate += daoParseUpdateTmp
					}
				}
			}
		} else if gstr.Pos(column[`Type`].String(), `int`) != -1 && gstr.Pos(column[`Type`].String(), `point`) == -1 { //int等类型
			if field == `pid` { //pid
				if len(tpl.HandelLabelList) > 0 {
					daoParseFieldTmp := `
			case ` + "`p" + gstr.CaseCamel(tpl.HandelLabelList[0]) + "`" + `:
				tableP := ` + "`p_`" + ` + daoModel.DbTable
				m = m.Fields(tableP + ` + "`.`" + ` + daoThis.Columns().` + gstr.CaseCamel(tpl.HandelLabelList[0]) + ` + ` + "` AS `" + ` + v)
				m = m.Handler(daoThis.ParseJoin(tableP, daoModel))`
					if gstr.Pos(tplDao, daoParseFieldTmp) == -1 {
						daoParseField += daoParseFieldTmp
					}
				}
				daoParseFieldTmp := `
			case ` + "`tree`" + `:
				m = m.Fields(daoModel.DbTable + ` + "`.`" + ` + daoThis.PrimaryKey())
				m = m.Fields(daoModel.DbTable + ` + "`.`" + ` + daoThis.Columns().` + fieldCaseCamel + `)
				m = m.Handler(daoThis.ParseOrder([]string{` + "`tree`" + `}, daoModel))`
				if gstr.Pos(tplDao, daoParseFieldTmp) == -1 {
					daoParseField += daoParseFieldTmp
				}
				daoParseOrderTmp := `
			case ` + "`tree`" + `:
				m = m.OrderAsc(daoModel.DbTable + ` + "`.`" + ` + daoThis.Columns().` + fieldCaseCamel + `)`
				if tpl.HandlePid.SortField != `` {
					daoParseOrderTmp += `
				m = m.OrderAsc(daoModel.DbTable + ` + "`.`" + ` + daoThis.Columns().` + gstr.CaseCamel(tpl.HandlePid.SortField) + `)`
				}
				daoParseOrderTmp += `
				m = m.OrderAsc(daoModel.DbTable + ` + "`.`" + ` + daoThis.PrimaryKey())`
				if gstr.Pos(tplDao, daoParseOrderTmp) == -1 {
					daoParseOrder += daoParseOrderTmp
				}
				daoParseJoinTmp := `
		case ` + "`p_`" + ` + daoModel.DbTable:
			m = m.LeftJoin(daoModel.DbTable+` + "` AS `" + `+joinTable, joinTable+` + "`.`" + `+daoThis.PrimaryKey()+` + "` = `" + `+daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + fieldCaseCamel + `)`
				if gstr.Pos(tplDao, daoParseJoinTmp) == -1 {
					daoParseJoin += daoParseJoinTmp
				}

				if tpl.HandlePid.IsCoexist {
					daoParseInsertBeforeTmp := `
		if _, ok := insert[daoThis.Columns().` + gstr.CaseCamel(tpl.HandlePid.PidField) + `]; !ok {
			insert[daoThis.Columns().` + gstr.CaseCamel(tpl.HandlePid.PidField) + `] = 0
		}`
					if gstr.Pos(tplDao, daoParseInsertBeforeTmp) == -1 {
						daoParseInsertBefore += daoParseInsertBeforeTmp
					}
					daoParseInsertTmp := `
			case daoThis.Columns().` + gstr.CaseCamel(tpl.HandlePid.PidField) + `:
				insertData[k] = v
				if gconv.Uint(v) > 0 {
					pInfo, _ := daoThis.CtxDaoModel(m.GetCtx()).Filter(daoThis.PrimaryKey(), v).One()
					daoModel.AfterInsert[` + "`pIdPath`" + `] = pInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.HandlePid.IdPathField) + `].String()
					daoModel.AfterInsert[` + "`pLevel`" + `] = pInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.HandlePid.LevelField) + `].Uint()
				} else {
					daoModel.AfterInsert[` + "`pIdPath`" + `] = ` + "`0`" + `
					daoModel.AfterInsert[` + "`pLevel`" + `] = 0
				}`
					if gstr.Pos(tplDao, daoParseInsertTmp) == -1 {
						daoParseInsert += daoParseInsertTmp
					}
					daoHookInsertTmp := `

			updateSelfData := map[string]interface{}{}
			for k, v := range daoModel.AfterInsert {
				switch k {
				case ` + "`pIdPath`" + `:
					updateSelfData[daoThis.Columns().` + gstr.CaseCamel(tpl.HandlePid.IdPathField) + `] = gconv.String(v) + ` + "`-`" + ` + gconv.String(id)
				case ` + "`pLevel`" + `:
					updateSelfData[daoThis.Columns().` + gstr.CaseCamel(tpl.HandlePid.LevelField) + `] = gconv.Uint(v) + 1
				}
			}
			if len(updateSelfData) > 0 {
				daoModel.CloneNew().Filter(daoThis.PrimaryKey(), id).HookUpdate(updateSelfData).Update()
			}`
					if gstr.Pos(tplDao, daoHookInsertTmp) == -1 {
						daoHookInsert += daoHookInsertTmp
					}
					daoParseUpdateTmp := `
			case daoThis.Columns().` + gstr.CaseCamel(tpl.HandlePid.PidField) + `:
				updateData[daoModel.DbTable+` + "`.`" + `+k] = v
				pIdPath := ` + "`0`" + `
				var pLevel uint = 0
				if gconv.Uint(v) > 0 {
					pInfo, _ := daoThis.CtxDaoModel(m.GetCtx()).Filter(daoThis.PrimaryKey(), v).One()
					pIdPath = pInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.HandlePid.IdPathField) + `].String()
					pLevel = pInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.HandlePid.LevelField) + `].Uint()
				}
				updateData[daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + gstr.CaseCamel(tpl.HandlePid.IdPathField) + `] = gdb.Raw(` + "`CONCAT('`" + ` + pIdPath + ` + "`-', `" + ` + daoThis.PrimaryKey() + ` + "`)`" + `)
				updateData[daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + gstr.CaseCamel(tpl.HandlePid.LevelField) + `] = pLevel + 1
				//更新所有子孙级的idPath和level
				updateChildIdPathAndLevelList := []map[string]interface{}{}
				oldList, _ := daoThis.CtxDaoModel(m.GetCtx()).Filter(daoThis.PrimaryKey(), daoModel.IdArr).All()
				for _, oldInfo := range oldList {
					if gconv.Uint(v) != oldInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.HandlePid.PidField) + `].Uint() {
						updateChildIdPathAndLevelList = append(updateChildIdPathAndLevelList, map[string]interface{}{
							` + "`pIdPathOfOld`" + `: oldInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.HandlePid.IdPathField) + `],
							` + "`pIdPathOfNew`" + `: pIdPath + ` + "`-`" + ` + oldInfo[daoThis.PrimaryKey()].String(),
							` + "`pLevelOfOld`" + `:  oldInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.HandlePid.LevelField) + `],
							` + "`pLevelOfNew`" + `:  pLevel + 1,
						})
					}
				}
				if len(updateChildIdPathAndLevelList) > 0 {
					daoModel.AfterUpdate[` + "`updateChildIdPathAndLevelList`" + `] = updateChildIdPathAndLevelList
				}
			case ` + "`childIdPath`" + `: //更新所有子孙级的idPath。参数：map[string]interface{}{` + "`pIdPathOfOld`" + `: ` + "`父级IdPath（旧）`" + `, ` + "`pIdPathOfNew`" + `: ` + "`父级IdPath（新）`" + `}
				val := gconv.Map(v)
				pIdPathOfOld := gconv.String(val[` + "`pIdPathOfOld`" + `])
				pIdPathOfNew := gconv.String(val[` + "`pIdPathOfNew`" + `])
				updateData[daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + gstr.CaseCamel(tpl.HandlePid.IdPathField) + `] = gdb.Raw(` + "`REPLACE(`" + ` + daoThis.Columns().` + gstr.CaseCamel(tpl.HandlePid.IdPathField) + ` + ` + "`, '`" + ` + pIdPathOfOld + ` + "`', '`" + ` + pIdPathOfNew + ` + "`')`" + `)
			case ` + "`childLevel`" + `: //更新所有子孙级的level。参数：map[string]interface{}{` + "`pLevelOfOld`" + `: ` + "`父级Level（旧）`" + `, ` + "`pLevelOfNew`" + `: ` + "`父级Level（新）`" + `}
				val := gconv.Map(v)
				pLevelOfOld := gconv.Uint(val[` + "`pLevelOfOld`" + `])
				pLevelOfNew := gconv.Uint(val[` + "`pLevelOfNew`" + `])
				updateData[daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + gstr.CaseCamel(tpl.HandlePid.LevelField) + `] = gdb.Raw(daoModel.DbTable + ` + "`.`" + ` + daoThis.Columns().` + gstr.CaseCamel(tpl.HandlePid.LevelField) + ` + ` + "` + `" + ` + gconv.String(pLevelOfNew-pLevelOfOld))
				if pLevelOfNew < pLevelOfOld {
					updateData[daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + gstr.CaseCamel(tpl.HandlePid.LevelField) + `] = gdb.Raw(daoModel.DbTable + ` + "`.`" + ` + daoThis.Columns().` + gstr.CaseCamel(tpl.HandlePid.LevelField) + ` + ` + "` - `" + ` + gconv.String(pLevelOfOld-pLevelOfNew))
				}`
					if gstr.Pos(tplDao, daoParseUpdateTmp) == -1 {
						daoParseUpdate += daoParseUpdateTmp
					}
					daoHookUpdateAfterTmp := `

			for k, v := range daoModel.AfterUpdate {
				switch k {
				case ` + "`updateChildIdPathAndLevelList`" + `: //修改pid时，更新所有子孙级的idPath和level。参数：[]map[string]interface{}{` + "`pIdPathOfOld`" + `: ` + "`父级IdPath（旧）`" + `, ` + "`pIdPathOfNew`" + `: ` + "`父级IdPath（新）`" + `, ` + "`pLevelOfOld`" + `: ` + "`父级Level（旧）`" + `, ` + "`pLevelOfNew`" + `: ` + "`父级Level（新）`" + `}
					val := v.([]map[string]interface{})
					for _, v1 := range val {
						daoModel.CloneNew().Filter(` + "`pIdPathOfOld`" + `, v1[` + "`pIdPathOfOld`" + `]).HookUpdate(g.Map{
							` + "`childIdPath`" + `: g.Map{
								` + "`pIdPathOfOld`" + `: v1[` + "`pIdPathOfOld`" + `],
								` + "`pIdPathOfNew`" + `: v1[` + "`pIdPathOfNew`" + `],
							},
							` + "`childLevel`" + `: g.Map{
								` + "`pLevelOfOld`" + `: v1[` + "`pLevelOfOld`" + `],
								` + "`pLevelOfNew`" + `: v1[` + "`pLevelOfNew`" + `],
							},
						}).Update()
					}
				}
			}`
					if gstr.Pos(tplDao, daoHookUpdateAfterTmp) == -1 {
						daoHookUpdateAfter += daoHookUpdateAfterTmp
					}
					daoParseFilterTmp := `
			case ` + "`pIdPathOfOld`" + `: //父级IdPath（旧）
				m = m.WhereLike(daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + gstr.CaseCamel(tpl.HandlePid.IdPathField) + `, gconv.String(v)+` + "`-%`" + `)`
					if gstr.Pos(tplDao, daoParseFilterTmp) == -1 {
						daoParseFilter += daoParseFilterTmp
					}
				}
			} else if field == `level` && tpl.HandlePid.IsCoexist { //level
				daoParseOrderTmp := `
			case daoThis.Columns().` + fieldCaseCamel + `:
				m = m.Order(daoModel.DbTable + ` + "`.`" + ` + v)
				m = m.OrderDesc(daoModel.DbTable + ` + "`.`" + ` + daoThis.PrimaryKey())` //追加主键倒序。mysql排序字段有重复值时，分页会导致同一条数据可能在不同页都出现
				if gstr.Pos(tplDao, daoParseOrderTmp) == -1 {
					daoParseOrder += daoParseOrderTmp
				}
			} else if garray.NewStrArrayFrom([]string{`sort`, `weight`}).Contains(fieldSuffix) { //sort,weight等后缀
				daoParseOrderTmp := `
			case daoThis.Columns().` + fieldCaseCamel + `:
				m = m.Order(daoModel.DbTable + ` + "`.`" + ` + v)
				m = m.OrderDesc(daoModel.DbTable + ` + "`.`" + ` + daoThis.PrimaryKey())` //追加主键倒序。mysql排序字段有重复值时，分页会导致同一条数据可能在不同页都出现
				if gstr.Pos(tplDao, daoParseOrderTmp) == -1 {
					daoParseOrder += daoParseOrderTmp
				}
			} else if garray.NewStrArrayFrom([]string{`id`}).Contains(fieldSuffix) { //id后缀
				if tpl.RelTableMap[field].TableRaw != `` {
					relTable := tpl.RelTableMap[field]
					daoPath := relTable.RelTableCaseCamel
					if !relTable.IsSameDir {
						daoPath = `dao` + relTable.RelDaoDirCaseCamel + `.` + relTable.RelTableCaseCamel
						daoImportOtherDaoTmp := `
	dao` + relTable.RelDaoDirCaseCamel + ` "api/internal/dao/` + relTable.RelDaoDir + `"`
						if gstr.Pos(tplDao, daoImportOtherDaoTmp) == -1 {
							daoImportOtherDao += daoImportOtherDaoTmp
						}
					}
					if !tpl.RelTableMap[field].IsRedundRelNameField {
						daoParseFieldTmp := `//因前端页面已用该字段名显示，故不存在时改成` + "`" + relTable.RelTableField + relTable.RelSuffix + "`" + `（控制器也要改）。同时下面Fields方法改成m = m.Fields(table` + relTable.RelTableCaseCamel + relTable.RelSuffixCaseCamel + ` + ` + "`.`" + ` + ` + daoPath + `.Columns().Xxxx + ` + "` AS `" + ` + v)`
						if gstr.Pos(tplDao, daoParseFieldTmp) == -1 {
							if relTable.RelSuffix != `` {
								daoParseFieldTmp = `
			case ` + daoPath + `.Columns().` + gstr.CaseCamel(relTable.RelTableField) + " + `" + relTable.RelSuffix + "`: " + daoParseFieldTmp + `
				table` + relTable.RelTableCaseCamel + relTable.RelSuffixCaseCamel + ` := ` + daoPath + `.ParseDbTable(m.GetCtx()) + ` + "`" + relTable.RelSuffixCaseSnake + "`" + `
				m = m.Fields(table` + relTable.RelTableCaseCamel + relTable.RelSuffixCaseCamel + ` + ` + "`.`" + ` + ` + daoPath + `.Columns().` + gstr.CaseCamel(relTable.RelTableField) + ` + ` + "` AS `" + ` + v)
				m = m.Handler(daoThis.ParseJoin(table` + relTable.RelTableCaseCamel + relTable.RelSuffixCaseCamel + `, daoModel))`
							} else {
								daoParseFieldTmp = `
			case ` + daoPath + `.Columns().` + gstr.CaseCamel(relTable.RelTableField) + `: ` + daoParseFieldTmp + `
				table` + relTable.RelTableCaseCamel + ` := ` + daoPath + `.ParseDbTable(m.GetCtx())
				m = m.Fields(table` + relTable.RelTableCaseCamel + ` + ` + "`.`" + ` + v)
				m = m.Handler(daoThis.ParseJoin(table` + relTable.RelTableCaseCamel + `, daoModel))`
							}
							daoParseField += daoParseFieldTmp
						}
					}
					daoParseJoinTmp := `
		case ` + daoPath + `.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+` + "`.`" + `+` + daoPath + `.PrimaryKey()+` + "` = `" + `+daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + fieldCaseCamel + `)`
					if relTable.RelSuffix != `` {
						daoParseJoinTmp = `
		case ` + daoPath + `.ParseDbTable(m.GetCtx()) + ` + "`" + relTable.RelSuffixCaseSnake + "`" + `:
			m = m.LeftJoin(` + daoPath + `.ParseDbTable(m.GetCtx())+` + "` AS `" + `+joinTable, joinTable+` + "`.`" + `+` + daoPath + `.PrimaryKey()+` + "` = `" + `+daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + fieldCaseCamel + `)`
					}
					if gstr.Pos(tplDao, daoParseJoinTmp) == -1 {
						daoParseJoin += daoParseJoinTmp
					}
				}
			}
		} else if gstr.Pos(column[`Type`].String(), `timestamp`) != -1 || gstr.Pos(column[`Type`].String(), `date`) != -1 { //timestamp或datetime或date类型
			if gstr.Pos(column[`Type`].String(), `date`) != -1 && gstr.Pos(column[`Type`].String(), `datetime`) == -1 {
				daoParseOrderTmp := `
			case daoThis.Columns().` + fieldCaseCamel + `:
				m = m.Order(daoModel.DbTable + ` + "`.`" + ` + v)
				m = m.OrderDesc(daoModel.DbTable + ` + "`.`" + ` + daoThis.PrimaryKey())` //追加主键倒序。mysql排序字段有重复值时，分页会导致同一条数据可能在不同页都出现
				if gstr.Pos(tplDao, daoParseOrderTmp) == -1 {
					daoParseOrder += daoParseOrderTmp
				}
			}

			if garray.NewStrArrayFrom([]string{`start`}).Contains(fieldPrefix) { //start_前缀
				daoParseFilterTmp := `
			case daoThis.Columns().` + fieldCaseCamel + `:
				m = m.WhereLTE(daoModel.DbTable+` + "`.`" + `+k, v)`
				if !column[`Null`].Bool() && column[`Default`].String() == `` {
					daoParseFilterTmp = `
			case daoThis.Columns().` + fieldCaseCamel + `:
				m = m.Where(m.Builder().WhereLTE(daoModel.DbTable+` + "`.`" + `+k, v).WhereOrNull(daoModel.DbTable + ` + "`.`" + ` + k))`
				}
				if gstr.Pos(tplDao, daoParseFilterTmp) == -1 {
					daoParseFilter += daoParseFilterTmp
				}
			} else if garray.NewStrArrayFrom([]string{`end`}).Contains(fieldPrefix) { //end_前缀
				daoParseFilterTmp := `
			case daoThis.Columns().` + fieldCaseCamel + `:
				m = m.WhereGTE(daoModel.DbTable+` + "`.`" + `+k, v)`
				if !column[`Null`].Bool() && column[`Default`].String() == `` {
					daoParseFilterTmp = `
			case daoThis.Columns().` + fieldCaseCamel + `:
				m = m.Where(m.Builder().WhereGTE(daoModel.DbTable+` + "`.`" + `+k, v).WhereOrNull(daoModel.DbTable + ` + "`.`" + ` + k))`
				}
				if gstr.Pos(tplDao, daoParseFilterTmp) == -1 {
					daoParseFilter += daoParseFilterTmp
				}
			}
		} else if gstr.Pos(column[`Type`].String(), `json`) != -1 { //json类型
			if column[`Null`].Bool() {
				daoParseInsertTmp := `
			case daoThis.Columns().` + fieldCaseCamel + `:
				insertData[k] = v
				if gconv.String(v) == ` + "``" + ` {
					insertData[k] = nil
				}`
				if gstr.Pos(tplDao, daoParseInsertTmp) == -1 {
					daoParseInsert += daoParseInsertTmp
				}
				daoParseUpdateTmp := `
			case daoThis.Columns().` + fieldCaseCamel + `:
				updateData[daoModel.DbTable+` + "`.`" + `+k] = gvar.New(v)
				if gconv.String(v) == ` + "``" + ` {
					updateData[daoModel.DbTable+` + "`.`" + `+k] = nil
				}`
				if gstr.Pos(tplDao, daoParseUpdateTmp) == -1 {
					daoParseUpdate += daoParseUpdateTmp
				}
			}
		}
	}

	if daoParseInsertBefore != `` {
		daoParseInsertBeforePoint := `
		insertData := map[string]interface{}{}`
		tplDao = gstr.Replace(tplDao, daoParseInsertBeforePoint, daoParseInsertBefore+daoParseInsertBeforePoint, 1)
	}
	if daoParseInsert != `` {
		daoParseInsertPoint := `case ` + "`id`" + `:
				insertData[daoThis.PrimaryKey()] = v`
		tplDao = gstr.Replace(tplDao, daoParseInsertPoint, daoParseInsertPoint+daoParseInsert, 1)
	}
	if daoHookInsert != `` {
		daoHookInsertPoint := `// id, _ := result.LastInsertId()

			/* for k, v := range daoModel.AfterInsert {
				switch k {
				case ` + "`xxxx`" + `:
					daoModel.CloneNew().Filter(daoThis.PrimaryKey(), id).HookUpdate(g.Map{k: v}).Update()
				}
			} */`
		tplDao = gstr.Replace(tplDao, daoHookInsertPoint, `id, _ := result.LastInsertId()`+daoHookInsert, 1)
	}
	if daoParseUpdate != `` {
		daoParseUpdatePoint := `case ` + "`id`" + `:
				updateData[daoModel.DbTable+` + "`.`" + `+daoThis.PrimaryKey()] = v`
		tplDao = gstr.Replace(tplDao, daoParseUpdatePoint, daoParseUpdatePoint+daoParseUpdate, 1)
	}
	if daoHookUpdateBefore != `` || daoHookUpdateAfter != `` {
		daoHookUpdatePoint := `

			/* row, _ := result.RowsAffected()
			if row == 0 {
				return
			} */

			/* for k, v := range daoModel.AfterUpdate {
				switch k {
				case ` + "`xxxx`" + `:
					for _, id := range daoModel.IdArr {
						daoModel.CloneNew().Filter(daoThis.PrimaryKey(), id).HookUpdate(g.Map{k: v}).Update()
					}
				}
			} */`
		if daoHookUpdateBefore != `` {
			tplDao = gstr.Replace(tplDao, daoHookUpdatePoint, daoHookUpdateBefore+daoHookUpdatePoint, 1)
		}
		if daoHookUpdateAfter != `` {
			tplDao = gstr.Replace(tplDao, daoHookUpdatePoint, `

			row, _ := result.RowsAffected()
			if row == 0 {
				return
			}`+daoHookUpdateAfter, 1)
		}
	}
	if daoParseField != `` {
		daoParseFieldPoint := `case ` + "`id`" + `:
				m = m.Fields(daoModel.DbTable + ` + "`.`" + ` + daoThis.PrimaryKey() + ` + "` AS `" + ` + v)`
		tplDao = gstr.Replace(tplDao, daoParseFieldPoint, daoParseFieldPoint+daoParseField, 1)
	}
	if daoHookSelect != `` {
		daoHookSelectPoint := `
					default:
						record[v] = gvar.New(nil)`
		tplDao = gstr.Replace(tplDao, daoHookSelectPoint, daoHookSelect+daoHookSelectPoint, 1)
	}
	if daoParseFilter != `` {
		daoParseFilterPoint := `case ` + "`id`, `idArr`" + `:
				m = m.Where(daoModel.DbTable+` + "`.`" + `+daoThis.PrimaryKey(), v)`
		tplDao = gstr.Replace(tplDao, daoParseFilterPoint, daoParseFilterPoint+daoParseFilter, 1)
	}
	if daoParseOrder != `` {
		daoParseOrderPoint := `case ` + "`id`" + `:
				m = m.Order(daoModel.DbTable + ` + "`.`" + ` + gstr.Replace(v, k, daoThis.PrimaryKey(), 1))`
		tplDao = gstr.Replace(tplDao, daoParseOrderPoint, daoParseOrderPoint+daoParseOrder, 1)
	}
	if daoParseJoin != `` {
		daoParseJoinPoint := `
		/* case Xxxx.ParseDbTable(m.GetCtx()):
		m = m.LeftJoin(joinTable, joinTable+` + "`.`" + `+Xxxx.Columns().XxxxId+` + "` = `" + `+daoModel.DbTable+` + "`.`" + `+daoThis.PrimaryKey())
		// m = m.LeftJoin(Xxxx.ParseDbTable(m.GetCtx())+` + "` AS `" + `+joinTable, joinTable+` + "`.`" + `+Xxxx.Columns().XxxxId+` + "` = `" + `+daoModel.DbTable+` + "`.`" + `+daoThis.PrimaryKey()) */`
		tplDao = gstr.Replace(tplDao, daoParseJoinPoint, daoParseJoinPoint+daoParseJoin, 1)
	}
	if daoFunc != `` {
		tplDao = tplDao + daoFunc
	}
	if daoImportOtherDao != `` {
		daoImportOtherDaoPoint := `"api/internal/dao/` + tpl.ModuleDirCaseKebab + `/internal"`
		tplDao = gstr.Replace(tplDao, daoImportOtherDaoPoint, daoImportOtherDaoPoint+daoImportOtherDao, 1)
	}

	tplDao = gstr.Replace(tplDao, `"github.com/gogf/gf/v2/util/gconv"`, `"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"`, 1)

	gfile.PutContents(saveFile, tplDao)
	utils.GoFileFmt(saveFile)
}

// logic模板生成（文件不存在时增删改查全部生成，已存在不处理不覆盖）
func (myGenThis *myGenHandler) genLogic() {
	option := myGenThis.option
	tpl := myGenThis.tpl

	saveFile := gfile.SelfDir() + `/internal/logic/` + gstr.Replace(tpl.ModuleDirCaseKebab, `/`, `-`) + `/` + tpl.TableCaseSnake + `.go`
	if !option.IsCover && gfile.IsFile(saveFile) {
		return
	}

	tplLogic := `package logic

import (
	dao` + tpl.ModuleDirCaseCamel + ` "api/internal/dao/` + tpl.ModuleDirCaseKebab + `"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type s` + myGenThis.tpl.LogicStructName + ` struct{}

func New` + myGenThis.tpl.LogicStructName + `() *s` + myGenThis.tpl.LogicStructName + ` {
	return &s` + myGenThis.tpl.LogicStructName + `{}
}

func init() {
	service.Register` + myGenThis.tpl.LogicStructName + `(New` + myGenThis.tpl.LogicStructName + `())
}

// 新增
func (logicThis *s` + myGenThis.tpl.LogicStructName + `) Create(ctx context.Context, data map[string]interface{}) (id int64, err error) {
	daoThis := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `
	daoModelThis := daoThis.CtxDaoModel(ctx)
`
	if tpl.HandlePid.PidField != `` {
		tplLogic += `
	if _, ok := data[daoThis.Columns().` + gstr.CaseCamel(tpl.HandlePid.PidField) + `]; ok {
		pid := gconv.Uint(data[daoThis.Columns().` + gstr.CaseCamel(tpl.HandlePid.PidField) + `])
		if pid > 0 {
			pInfo, _ := daoModelThis.CloneNew().Filter(daoThis.PrimaryKey(), pid).One()
			if pInfo.IsEmpty() {
				err = utils.NewErrorCode(ctx, 29999997, ` + "``" + `)
				return
			}
		}
	}
`
	}
	tplLogic += `
	id, err = daoModelThis.HookInsert(data).InsertAndGetId()
	return
}

// 修改
func (logicThis *s` + myGenThis.tpl.LogicStructName + `) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error) {
	daoThis := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `
	daoModelThis := daoThis.CtxDaoModel(ctx)

	daoModelThis.Filters(filter).SetIdArr()
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ` + "``" + `)
		return
	}
`
	if tpl.HandlePid.PidField != `` {
		tplLogic += `
	if _, ok := data[daoThis.Columns().` + gstr.CaseCamel(tpl.HandlePid.PidField) + `]; ok {`
		if tpl.HandlePid.IsCoexist {
			tplLogic += `
		pid := gconv.Uint(data[daoThis.Columns().` + gstr.CaseCamel(tpl.HandlePid.PidField) + `])
		if pid > 0 {
			pInfo, _ := daoModelThis.CloneNew().Filter(daoThis.PrimaryKey(), pid).One()
			if pInfo.IsEmpty() {
				err = utils.NewErrorCode(ctx, 29999997, ` + "``" + `)
				return
			}
			oldList, _ := daoModelThis.CloneNew().Filter(daoThis.PrimaryKey(), daoModelThis.IdArr).All()
			for _, oldInfo := range oldList {
				if pid == oldInfo[daoThis.PrimaryKey()].Uint() { //父级不能是自身
					err = utils.NewErrorCode(ctx, 29999996, ` + "``" + `)
					return
				}
				if pid != oldInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.HandlePid.PidField) + `].Uint() {
					if garray.NewStrArrayFrom(gstr.Split(pInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.HandlePid.IdPathField) + `].String(), ` + "`-`" + `)).Contains(oldInfo[daoThis.PrimaryKey()].String()) { //父级不能是自身的子孙级
						err = utils.NewErrorCode(ctx, 29999995, ` + "``" + `)
						return
					}
				}
			}
		}`
		} else {
			tplLogic += `
		pid := gconv.Uint(data[daoThis.Columns().` + gstr.CaseCamel(tpl.HandlePid.PidField) + `])
		if pid > 0 {
			pInfo, _ := daoModelThis.CloneNew().Filter(daoThis.PrimaryKey(), pid).One()
			if pInfo.IsEmpty() {
				err = utils.NewErrorCode(ctx, 29999997, ` + "``" + `)
				return
			}
		}
		for _, id := range daoModelThis.IdArr {
			if pid == id { //父级不能是自身
				err = utils.NewErrorCode(ctx, 29999996, ` + "``" + `)
				return
			}
		}`
		}
		tplLogic += `
	}
`
	}
	tplLogic += `
	row, err = daoModelThis.HookUpdate(data).UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *s` + myGenThis.tpl.LogicStructName + `) Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error) {
	daoThis := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `
	daoModelThis := daoThis.CtxDaoModel(ctx)

	daoModelThis.Filters(filter).SetIdArr()
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ` + "``" + `)
		return
	}
`
	if tpl.HandlePid.PidField != `` {
		tplLogic += `
	count, _ := daoModelThis.CloneNew().Filter(daoThis.Columns().` + gstr.CaseCamel(tpl.HandlePid.PidField) + `, daoModelThis.IdArr).Count()
	if count > 0 {
		err = utils.NewErrorCode(ctx, 29999994, ` + "``" + `)
		return
	}
`
	}
	tplLogic += `
	row, err = daoModelThis.HookDelete().DeleteAndGetAffected()
	return
}
`

	gfile.PutContents(saveFile, tplLogic)
	utils.GoFileFmt(saveFile)
	myGenThis.command(`service生成`, true, ``,
		`gf`, `gen`, `service`)
}

// api模板生成
func (myGenThis *myGenHandler) genApi() {
	option := myGenThis.option
	tpl := myGenThis.tpl

	saveFile := gfile.SelfDir() + `/api/` + option.SceneCode + `/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseSnake + `.go`
	if !option.IsCover && gfile.IsFile(saveFile) {
		return
	}

	apiReqFilterColumn := ``
	apiReqCreateColumn := ``
	apiReqUpdateColumn := ``
	apiResColumn := ``
	apiResColumnAlloweFieldList := ``
	if len(tpl.HandelLabelList) > 0 {
		apiReqFilterColumn += `Label          string      ` + "`" + `json:"label,omitempty" v:"max-length:30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"标签。常用于前端组件"` + "`\n"
		apiResColumn += `Label       *string     ` + "`" + `json:"label,omitempty" dc:"标签。常用于前端组件"` + "`\n"
	}
	for _, column := range tpl.TableColumnList {
		field := column[`Field`].String()
		fieldCaseCamel := gstr.CaseCamel(field)
		fieldCaseSnake := gstr.CaseSnake(field)
		fieldCaseSnakeOfRemove := gstr.Split(fieldCaseSnake, `_of_`)[0]
		fieldCaseCamelOfRemove := gstr.CaseCamel(fieldCaseSnakeOfRemove)
		fieldSplitArr := gstr.Split(fieldCaseSnakeOfRemove, `_`)
		fieldPrefix := fieldSplitArr[0]
		fieldSuffix := fieldSplitArr[len(fieldSplitArr)-1]
		comment := gstr.Trim(gstr.ReplaceByArray(column[`Comment`].String(), g.SliceStr{
			"\n", ` `,
			"\r", ` `,
			`"`, `\"`,
		}))
		resultStr, _ := gregex.MatchString(`.*\((\d*)\)`, column[`Type`].String())

		typeReqFilter := ``
		typeReqCreate := ``
		typeReqUpdate := ``
		typeRes := ``
		ruleReqFilter := ``
		ruleReqCreate := ``
		ruleReqUpdate := ``
		isRequired := false

		if garray.NewStrArrayFrom([]string{`DeletedAt`, `DeleteAt`, `DeletedTime`, `DeleteTime`}).Contains(fieldCaseCamel) {
			continue
		} else if garray.NewStrArrayFrom([]string{`UpdatedAt`, `UpdateAt`, `UpdatedTime`, `UpdateTime`}).Contains(fieldCaseCamel) {
			typeRes = `*gtime.Time`
		} else if garray.NewStrArrayFrom([]string{`CreatedAt`, `CreateAt`, `CreatedTime`, `CreateTime`}).Contains(fieldCaseCamel) {
			apiReqFilterColumn += `TimeRangeStart *gtime.Time ` + "`" + `json:"timeRangeStart,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"` + "`\n"
			apiReqFilterColumn += `TimeRangeEnd   *gtime.Time ` + "`" + `json:"timeRangeEnd,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"` + "`\n"
			typeRes = `*gtime.Time`
		} else if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` { //主键
			if field == `id` {
				continue
			}
			typeReqFilter = `*uint`
			typeRes = `*uint`
			ruleReqFilter = `min:1`
		} else if fieldCaseCamel == `IdPath` && (gstr.Pos(column[`Type`].String(), `varchar`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) && tpl.HandlePid.IsCoexist { //idPath|id_path
			typeRes = `*string`
		} else if garray.NewStrArrayFrom([]string{`status`, `type`, `method`, `pos`, `position`, `gender`}).Contains(fieldSuffix) && ((gstr.Pos(column[`Type`].String(), `int`) != -1 && gstr.Pos(column[`Type`].String(), `point`) == -1) || gstr.Pos(column[`Type`].String(), `char`) != -1) { //status,type,method,pos,position,gender等后缀
			typeReqFilter = `string`
			typeReqCreate = `*string`
			typeReqUpdate = `*string`
			typeRes = `*string`
			isStr := true
			if gstr.Pos(column[`Type`].String(), `int`) != -1 && gstr.Pos(column[`Type`].String(), `point`) == -1 {
				typeReqFilter = `*int`
				typeReqCreate = `*int`
				typeReqUpdate = `*int`
				typeRes = `*int`
				if gstr.Pos(column[`Type`].String(), `unsigned`) != -1 {
					typeReqFilter = `*uint`
					typeReqCreate = `*uint`
					typeReqUpdate = `*uint`
					typeRes = `*uint`
				}
				isStr = false
			}

			statusList := myGenThis.genStatusList(comment, isStr)
			statusArr := make([]string, len(statusList))
			for index, status := range statusList {
				statusArr[index] = status[0]
			}
			statusStr := gstr.Join(statusArr, `,`)
			ruleReqFilter += `in:` + statusStr
			ruleReqCreate += `in:` + statusStr
			ruleReqUpdate += `in:` + statusStr
		} else if (garray.NewStrArrayFrom([]string{`icon`, `cover`, `avatar`, `img`, `image`}).Contains(fieldSuffix) || gstr.SubStr(fieldCaseCamelOfRemove, -7) == `ImgList` || gstr.SubStr(fieldCaseCamelOfRemove, -6) == `ImgArr` || gstr.SubStr(fieldCaseCamelOfRemove, -9) == `ImageList` || gstr.SubStr(fieldCaseCamelOfRemove, -8) == `ImageArr` || garray.NewStrArrayFrom([]string{`video`}).Contains(fieldSuffix) || gstr.SubStr(fieldCaseCamelOfRemove, -9) == `VideoList` || gstr.SubStr(fieldCaseCamelOfRemove, -8) == `VideoArr`) && (gstr.Pos(column[`Type`].String(), `varchar`) != -1 || gstr.Pos(column[`Type`].String(), `json`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) { //icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀 //video,video_list,videoList,video_arr,videoArr等后缀
			if gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
				typeReqCreate = `*string`
				typeReqUpdate = `*string`
				typeRes = `*string`
				ruleReqCreate = `max-length:` + resultStr[1] + `|url`
				ruleReqUpdate = `max-length:` + resultStr[1] + `|url`
			} else {
				if !column[`Null`].Bool() {
					isRequired = true
				}
				typeReqCreate = `*[]string`
				typeReqUpdate = `*[]string`
				typeRes = `[]string`
				ruleReqCreate = `distinct|foreach|url|foreach|min-length:1`
				ruleReqUpdate = `distinct|foreach|url|foreach|min-length:1`
			}
		} else if garray.NewStrArrayFrom([]string{`list`, `arr`}).Contains(fieldSuffix) && (gstr.Pos(column[`Type`].String(), `json`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) { //list,arr等后缀
			if !column[`Null`].Bool() {
				isRequired = true
			}
			typeReqCreate = `*[]interface{}`
			typeReqUpdate = `*[]interface{}`
			typeRes = `[]interface{}`
			ruleReqCreate = `distinct`
			ruleReqUpdate = `distinct`
		} else if garray.NewStrArrayFrom([]string{`remark`, `desc`, `msg`, `message`, `intro`, `content`}).Contains(fieldSuffix) && (gstr.Pos(column[`Type`].String(), `varchar`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) { //remark,desc,msg,message,intro,content后缀
			typeReqCreate = `*string`
			typeReqUpdate = `*string`
			typeRes = `*string`
			if gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
				ruleReqCreate = `max-length:` + resultStr[1]
				ruleReqUpdate = `max-length:` + resultStr[1]
			}
		} else if gstr.Pos(column[`Type`].String(), `varchar`) != -1 { //varchar类型
			typeReqFilter = `string`
			typeReqCreate = `*string`
			typeReqUpdate = `*string`
			typeRes = `*string`
			ruleReqFilter = `max-length:` + resultStr[1]
			ruleReqCreate = `max-length:` + resultStr[1]
			ruleReqUpdate = `max-length:` + resultStr[1]
			if column[`Key`].String() == `UNI` && !column[`Null`].Bool() {
				isRequired = true
			}

			if garray.NewStrArrayFrom([]string{`name`}).Contains(fieldSuffix) { //name后缀
				if gstr.SubStr(gstr.CaseCamel(tpl.FieldPrimary), 0, -2)+`Name` == fieldCaseCamel {
					isRequired = true
				}
				ruleReqFilter += `|regex:^[\\p{L}\\p{M}\\p{N}_-]+$`
				ruleReqCreate += `|regex:^[\\p{L}\\p{M}\\p{N}_-]+$`
				ruleReqUpdate += `|regex:^[\\p{L}\\p{M}\\p{N}_-]+$`
			} else if garray.NewStrArrayFrom([]string{`code`}).Contains(fieldSuffix) { //code后缀
				ruleReqFilter += `|regex:^[\\p{L}\\p{M}\\p{N}_-]+$`
				ruleReqCreate += `|regex:^[\\p{L}\\p{M}\\p{N}_-]+$`
				ruleReqUpdate += `|regex:^[\\p{L}\\p{M}\\p{N}_-]+$`
			} else if garray.NewStrArrayFrom([]string{`phone`, `mobile`}).Contains(fieldSuffix) { //phone,mobile后缀
				ruleReqFilter += `|phone`
				ruleReqCreate += `|phone`
				ruleReqUpdate += `|phone`
			} else if garray.NewStrArrayFrom([]string{`url`, `link`}).Contains(fieldSuffix) { //url,link后缀
				ruleReqFilter += `|url`
				ruleReqCreate += `|url`
				ruleReqUpdate += `|url`
			} else if garray.NewStrArrayFrom([]string{`ip`}).Contains(fieldSuffix) { //IP后缀
				ruleReqFilter += `|ip`
				ruleReqCreate += `|ip`
				ruleReqUpdate += `|ip`
			}
		} else if gstr.Pos(column[`Type`].String(), `char`) != -1 { //char类型
			typeReqFilter = `string`
			typeReqCreate = `*string`
			typeReqUpdate = `*string`
			typeRes = `*string`
			ruleReqFilter = `max-length:` + resultStr[1]
			ruleReqCreate = `size:` + resultStr[1]
			ruleReqUpdate = `size:` + resultStr[1]
			if garray.NewStrArrayFrom([]string{`password`, `passwd`}).Contains(fieldSuffix) && column[`Type`].String() == `char(32)` { //password,passwd后缀
				typeReqFilter = ``
				typeRes = ``
				ruleReqFilter = ``
				isRequired = true
			} else if garray.NewStrArrayFrom([]string{`salt`}).Contains(fieldSuffix) && tpl.HandlePasswordMap[myGenThis.genHandlePasswordMapKey(field)].IsCoexist { //salt后缀
				continue
			} else {
				if column[`Key`].String() == `UNI` && !column[`Null`].Bool() {
					isRequired = true
				}
			}
		} else if gstr.Pos(column[`Type`].String(), `int`) != -1 && gstr.Pos(column[`Type`].String(), `point`) == -1 { //int等类型
			typeReqFilter = `*int`
			typeReqCreate = `*int`
			typeReqUpdate = `*int`
			typeRes = `*int`
			ruleReqFilter = ``
			ruleReqCreate = ``
			ruleReqUpdate = ``
			if gstr.Pos(column[`Type`].String(), `unsigned`) != -1 {
				typeReqFilter = `*uint`
				typeReqCreate = `*uint`
				typeReqUpdate = `*uint`
				typeRes = `*uint`
			}

			if field == `pid` { //pid
				if len(tpl.HandelLabelList) > 0 {
					apiResColumnAlloweFieldList += `P` + gstr.CaseCamel(tpl.HandelLabelList[0]) + ` *string ` + "`" + `json:"p` + gstr.CaseCamel(tpl.HandelLabelList[0]) + `,omitempty" dc:"父级"` + "`\n"
				}
			} else if field == `level` && tpl.HandlePid.IsCoexist { //level
				typeReqCreate = ``
				typeReqUpdate = ``
				ruleReqFilter += `min:1`
			} else if garray.NewStrArrayFrom([]string{`sort`, `weight`}).Contains(fieldSuffix) { //sort,weight等后缀
				typeReqFilter = ``
				ruleReqCreate += `between:0,100`
				ruleReqUpdate += `between:0,100`
			} else if garray.NewStrArrayFrom([]string{`id`}).Contains(fieldSuffix) { //id后缀
				if tpl.RelTableMap[field].TableRaw != `` && !tpl.RelTableMap[field].IsRedundRelNameField {
					relTable := tpl.RelTableMap[field]
					apiResColumnAlloweFieldList += gstr.CaseCamel(relTable.RelTableField) + relTable.RelSuffixCaseCamel + ` *string ` + "`" + `json:"` + relTable.RelTableField + relTable.RelSuffix + `,omitempty" dc:"` + relTable.RelTableFieldName + `"` + "`\n"
				}
				ruleReqFilter += `min:1`
				ruleReqCreate += `min:1`
				ruleReqUpdate += `min:1`
			} else if garray.NewStrArrayFrom([]string{`is`}).Contains(fieldPrefix) { //is_前缀
				ruleReqFilter += `in:0,1`
				ruleReqCreate += `in:0,1`
				ruleReqUpdate += `in:0,1`
			} else { //默认处理（int等类型）
				typeReqFilter = ``
			}
		} else if gstr.Pos(column[`Type`].String(), `decimal`) != -1 || gstr.Pos(column[`Type`].String(), `double`) != -1 || gstr.Pos(column[`Type`].String(), `float`) != -1 { //float类型
			typeReqFilter = `*float64`
			typeReqCreate = `*float64`
			typeReqUpdate = `*float64`
			typeRes = `*float64`
			ruleReqFilter = ``
			ruleReqCreate = ``
			ruleReqUpdate = ``
			if gstr.Pos(column[`Type`].String(), `unsigned`) != -1 {
				ruleReqFilter += `min:0`
				ruleReqCreate += `min:0`
				ruleReqUpdate += `min:0`
			}

			//默认处理（float类型）
			typeReqFilter = ``
		} else if gstr.Pos(column[`Type`].String(), `timestamp`) != -1 || gstr.Pos(column[`Type`].String(), `date`) != -1 { //timestamp或datetime或date类型
			typeReqFilter = ``
			typeReqCreate = `*gtime.Time`
			typeReqUpdate = `*gtime.Time`
			typeRes = `*gtime.Time`
			ruleReqFilter = `date-format:Y-m-d H:i:s`
			ruleReqCreate = `date-format:Y-m-d H:i:s`
			ruleReqUpdate = `date-format:Y-m-d H:i:s`
			if gstr.Pos(column[`Type`].String(), `date`) != -1 && gstr.Pos(column[`Type`].String(), `datetime`) == -1 {
				typeReqFilter = `*gtime.Time`
				typeRes = `*string`
				ruleReqFilter = `date-format:Y-m-d`
				ruleReqCreate = `date-format:Y-m-d`
				ruleReqUpdate = `date-format:Y-m-d`
			}
			if !column[`Null`].Bool() && column[`Default`].String() == `` {
				isRequired = true
			}

			if garray.NewStrArrayFrom([]string{`start`}).Contains(fieldPrefix) || garray.NewStrArrayFrom([]string{`end`}).Contains(fieldPrefix) { //start_前缀 //end_前缀
				typeReqFilter = `*gtime.Time`
			}
		} else if gstr.Pos(column[`Type`].String(), `json`) != -1 { //json类型
			if !column[`Null`].Bool() {
				isRequired = true
			}
			typeReqCreate = `*string`
			typeReqUpdate = `*string`
			typeRes = `*string`
			ruleReqCreate = `json`
			ruleReqUpdate = `json`
		} else if gstr.Pos(column[`Type`].String(), `text`) != -1 { //text类型
			typeReqCreate = `*string`
			typeReqUpdate = `*string`
			typeRes = `*string`
		} else { //默认处理
			typeReqFilter = `string`
			typeReqCreate = `*string`
			typeReqUpdate = `*string`
			typeRes = `*string`
		}

		if typeReqFilter != `` {
			apiReqFilterColumn += fieldCaseCamel + ` ` + typeReqFilter + ` ` + "`" + `json:"` + field + `,omitempty" v:"` + ruleReqFilter + `" dc:"` + comment + `"` + "`\n"
		}
		if typeReqCreate != `` {
			if isRequired {
				if ruleReqCreate == `` {
					ruleReqCreate = `required`
				} else {
					ruleReqCreate = `required|` + ruleReqCreate
				}
			}
			apiReqCreateColumn += fieldCaseCamel + ` ` + typeReqCreate + ` ` + "`" + `json:"` + field + `,omitempty" v:"` + ruleReqCreate + `" dc:"` + comment + `"` + "`\n"
		}
		if typeReqUpdate != `` {
			apiReqUpdateColumn += fieldCaseCamel + ` ` + typeReqUpdate + ` ` + "`" + `json:"` + field + `,omitempty" v:"` + ruleReqUpdate + `" dc:"` + comment + `"` + "`\n"
		}
		if typeRes != `` {
			apiResColumn += fieldCaseCamel + ` ` + typeRes + ` ` + "`" + `json:"` + field + `,omitempty" dc:"` + comment + `"` + "`\n"
		}
	}

	apiReqFilterColumn = gstr.SubStr(apiReqFilterColumn, 0, -len("\n"))
	apiReqCreateColumn = gstr.SubStr(apiReqCreateColumn, 0, -len("\n"))
	apiReqUpdateColumn = gstr.SubStr(apiReqUpdateColumn, 0, -len("\n"))
	apiResColumn = gstr.SubStr(apiResColumn, 0, -len("\n"))

	tplApi := `package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

`
	if option.IsList {
		tplApi += `
/*--------列表 开始--------*/
type ` + tpl.TableCaseCamel + `ListReq struct {
	g.Meta ` + "`" + `path:"/` + tpl.TableCaseKebab + `/list" method:"post" tags:"` + myGenThis.sceneName + `/` + option.CommonName + `" sm:"列表"` + "`" + `
	Filter ` + tpl.TableCaseCamel + `ListFilter ` + "`" + `json:"filter" dc:"过滤条件"` + "`" + `
	Field  []string        ` + "`" + `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"` + "`" + `
	Sort   string          ` + "`" + `json:"sort" default:"id DESC" dc:"排序"` + "`" + `
	Page   int             ` + "`" + `json:"page" v:"min:1" default:"1" dc:"页码"` + "`" + `
	Limit  int             ` + "`" + `json:"limit" v:"min:0" default:"10" dc:"每页数量。可传0取全部"` + "`" + `
}

type ` + tpl.TableCaseCamel + `ListFilter struct {
	Id             *uint       ` + "`" + `json:"id,omitempty" v:"min:1" dc:"ID"` + "`" + `
	IdArr          []uint      ` + "`" + `json:"idArr,omitempty" v:"distinct|foreach|min:1" dc:"ID数组"` + "`" + `
	ExcId          *uint       ` + "`" + `json:"excId,omitempty" v:"min:1" dc:"排除ID"` + "`" + `
	ExcIdArr       []uint      ` + "`" + `json:"excIdArr,omitempty" v:"distinct|foreach|min:1" dc:"排除ID数组"` + "`" + `
	` + apiReqFilterColumn + `
}

type ` + tpl.TableCaseCamel + `ListRes struct {`
		if option.IsCount {
			tplApi += `
	Count int         ` + "`" + `json:"count" dc:"总数"` + "`"
		}
		tplApi += `
	List  []` + tpl.TableCaseCamel + `ListItem ` + "`" + `json:"list" dc:"列表"` + "`" + `
}

type ` + tpl.TableCaseCamel + `ListItem struct {
	Id          *uint       ` + "`" + `json:"id,omitempty" dc:"ID"` + "`" + `
	` + apiResColumn + `
	` + apiResColumnAlloweFieldList + `
}

/*--------列表 结束--------*/

`
	}
	if option.IsInfo {
		tplApi += `/*--------详情 开始--------*/
type ` + tpl.TableCaseCamel + `InfoReq struct {
	g.Meta ` + "`" + `path:"/` + tpl.TableCaseKebab + `/info" method:"post" tags:"` + myGenThis.sceneName + `/` + option.CommonName + `" sm:"详情"` + "`" + `
	Id     uint     ` + "`" + `json:"id" v:"required|min:1" dc:"ID"` + "`" + `
	Field  []string ` + "`" + `json:"field" v:"distinct|foreach|min-length:1" dc:"查询字段，传值参考返回的字段名，默认返回全部字段。注意：如前端页面所需字段较少，建议传指定字段，可大幅减轻服务器及数据库压力"` + "`" + `
}

type ` + tpl.TableCaseCamel + `InfoRes struct {
	Info ` + tpl.TableCaseCamel + `Info ` + "`" + `json:"info" dc:"详情"` + "`" + `
}

type ` + tpl.TableCaseCamel + `Info struct {
	Id          *uint       ` + "`" + `json:"id,omitempty" dc:"ID"` + "`" + `
	` + apiResColumn + `
}

/*--------详情 结束--------*/

`
	}
	if option.IsCreate {
		tplApi += `/*--------新增 开始--------*/
type ` + tpl.TableCaseCamel + `CreateReq struct {
	g.Meta      ` + "`" + `path:"/` + tpl.TableCaseKebab + `/create" method:"post" tags:"` + myGenThis.sceneName + `/` + option.CommonName + `" sm:"新增"` + "`" + `
	` + apiReqCreateColumn + `
}

/*--------新增 结束--------*/

`
	}

	if option.IsUpdate {
		tplApi += `/*--------修改 开始--------*/
type ` + tpl.TableCaseCamel + `UpdateReq struct {
	g.Meta      ` + "`" + `path:"/` + tpl.TableCaseKebab + `/update" method:"post" tags:"` + myGenThis.sceneName + `/` + option.CommonName + `" sm:"修改"` + "`" + `
	IdArr       []uint  ` + "`" + `json:"idArr,omitempty" v:"required|distinct|foreach|min:1" dc:"ID数组"` + "`" + `
	` + apiReqUpdateColumn + `
}

/*--------修改 结束--------*/

`
	}

	if option.IsDelete {
		tplApi += `/*--------删除 开始--------*/
type ` + tpl.TableCaseCamel + `DeleteReq struct {
	g.Meta ` + "`" + `path:"/` + tpl.TableCaseKebab + `/del" method:"post" tags:"` + myGenThis.sceneName + `/` + option.CommonName + `" sm:"删除"` + "`" + `
	IdArr  []uint ` + "`" + `json:"idArr,omitempty" v:"required|distinct|foreach|min:1" dc:"ID数组"` + "`" + `
}

/*--------删除 结束--------*/
`
	}

	if option.IsList && tpl.HandlePid.PidField != `` {
		tplApi += `
/*--------列表（树状） 开始--------*/
type ` + tpl.TableCaseCamel + `TreeReq struct {
	g.Meta ` + "`" + `path:"/` + tpl.TableCaseKebab + `/tree" method:"post" tags:"` + myGenThis.sceneName + `/` + option.CommonName + `" sm:"列表（树状）"` + "`" + `
	Field  []string       ` + "`" + `json:"field" v:"foreach|min-length:1"` + "`" + `
	Filter ` + tpl.TableCaseCamel + `ListFilter ` + "`" + `json:"filter" dc:"过滤条件"` + "`" + `
}

type ` + tpl.TableCaseCamel + `TreeRes struct {
	Tree []` + tpl.TableCaseCamel + `TreeItem ` + "`" + `json:"tree" dc:"列表（树状）"` + "`" + `
}

type ` + tpl.TableCaseCamel + `TreeItem struct {
	Id       *uint       ` + "`" + `json:"id,omitempty" dc:"ID"` + "`" + `
	` + apiResColumn + `
	Children []` + tpl.TableCaseCamel + `TreeItem ` + "`" + `json:"children" dc:"子级列表"` + "`" + `
}

/*--------列表（树状） 结束--------*/
`
	}

	gfile.PutContents(saveFile, tplApi)
	utils.GoFileFmt(saveFile)
}

// controller模板生成
func (myGenThis *myGenHandler) genController() {
	option := myGenThis.option
	tpl := myGenThis.tpl

	saveFile := gfile.SelfDir() + `/internal/controller/` + option.SceneCode + `/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseSnake + `.go`
	if !option.IsCover && gfile.IsFile(saveFile) {
		return
	}

	controllerAlloweFieldList := "`id`, "
	controllerAlloweFieldInfo := "`id`, "
	controllerAlloweFieldTree := "`id`, "
	controllerAlloweFieldNoAuth := "`id`, "
	// controllerAlloweFieldDiff := `` // 可以不要。数据返回时，会根据API文件中的结构体做过滤
	daoImportOtherDao := ``
	if len(tpl.HandelLabelList) > 0 {
		controllerAlloweFieldList += "`label`, "
		controllerAlloweFieldInfo += "`label`, "
		controllerAlloweFieldTree += "`label`, "
		if tpl.HandlePid.PidField != `` {
			controllerAlloweFieldList += "`p" + gstr.CaseCamel(tpl.HandelLabelList[0]) + "`, "
			// controllerAlloweFieldInfo += "`p" + gstr.CaseCamel(tpl.HandelLabelList[0]) + "`, "
		}
		controllerAlloweFieldNoAuth += "`label`, "
		//TODO
		for _, v := range tpl.HandelLabelList {
			controllerAlloweFieldNoAuth += `dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.Columns().` + gstr.CaseCamel(v) + `, `
		}
		/* if tpl.LabelHandle.IsCoexist {
			controllerAlloweFieldNoAuth += `dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.Columns().Phone, ` + `dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.Columns().Account, `
		} else {
			controllerAlloweFieldNoAuth += `dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.Columns().` + gstr.CaseCamel(tpl.HandelLabelList[0]) + `, `
		} */
	}
	for _, column := range tpl.TableColumnList {
		field := column[`Field`].String()
		fieldCaseCamel := gstr.CaseCamel(field)
		fieldCaseSnake := gstr.CaseSnake(field)
		fieldCaseSnakeOfRemove := gstr.Split(fieldCaseSnake, `_of_`)[0]
		// fieldCaseCamelOfRemove := gstr.CaseCamel(fieldCaseSnakeOfRemove)
		fieldSplitArr := gstr.Split(fieldCaseSnakeOfRemove, `_`)
		// fieldPrefix := fieldSplitArr[0]
		fieldSuffix := fieldSplitArr[len(fieldSplitArr)-1]

		if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` { //主键
			if field != `id` {
				controllerAlloweFieldNoAuth += `dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.Columns().` + fieldCaseCamel + `, `
			}
		} else if gstr.Pos(column[`Type`].String(), `char`) != -1 { //char类型
			if garray.NewStrArrayFrom([]string{`password`, `passwd`}).Contains(fieldSuffix) && column[`Type`].String() == `char(32)` { //password,passwd后缀
				// controllerAlloweFieldDiff += `dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.Columns().` + fieldCaseCamel + `, `
			} else if garray.NewStrArrayFrom([]string{`salt`}).Contains(fieldSuffix) && tpl.HandlePasswordMap[myGenThis.genHandlePasswordMapKey(field)].IsCoexist { //salt后缀
				// controllerAlloweFieldDiff += `dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.Columns().` + fieldCaseCamel + `, `
			}
		} else if gstr.Pos(column[`Type`].String(), `int`) != -1 && gstr.Pos(column[`Type`].String(), `point`) == -1 { //int等类型
			if garray.NewStrArrayFrom([]string{`id`}).Contains(fieldSuffix) { //id后缀
				if tpl.RelTableMap[field].TableRaw != `` && !tpl.RelTableMap[field].IsRedundRelNameField {
					relTable := tpl.RelTableMap[field]
					// controllerAlloweFieldList += "`" + relTable.RelNameField + "`, "
					daoPath := `dao` + relTable.RelDaoDirCaseCamel + `.` + relTable.RelTableCaseCamel
					daoImportOtherDao += `
	dao` + relTable.RelDaoDirCaseCamel + ` "api/internal/dao/` + relTable.RelDaoDir + `"`
					controllerAlloweFieldList += daoPath + `.Columns().` + gstr.CaseCamel(relTable.RelTableField)
					if relTable.RelSuffix != `` {
						controllerAlloweFieldList += "+`" + relTable.RelSuffix + "`"
					}
					controllerAlloweFieldList += ", "
				}
			}
		}
	}
	controllerAlloweFieldList = gstr.SubStr(controllerAlloweFieldList, 0, -len(`, `))
	controllerAlloweFieldInfo = gstr.SubStr(controllerAlloweFieldInfo, 0, -len(`, `))
	controllerAlloweFieldTree = gstr.SubStr(controllerAlloweFieldTree, 0, -len(`, `))
	controllerAlloweFieldNoAuth = gstr.SubStr(controllerAlloweFieldNoAuth, 0, -len(`, `))
	// controllerAlloweFieldDiff = gstr.SubStr(controllerAlloweFieldDiff, 0, -len(`, `))

	tplController := `package controller

import (
	"api/api"
	api` + tpl.ModuleDirCaseCamel + ` "api/api/` + option.SceneCode + `/` + tpl.ModuleDirCaseKebab + `"
	dao` + tpl.ModuleDirCaseCamel + ` "api/internal/dao/` + tpl.ModuleDirCaseKebab + `"` + daoImportOtherDao + `
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
)

type ` + tpl.TableCaseCamel + ` struct{}

func New` + tpl.TableCaseCamel + `() *` + tpl.TableCaseCamel + ` {
	return &` + tpl.TableCaseCamel + `{}
}
`
	if option.IsList {
		tplController += `
// 列表
func (controllerThis *` + tpl.TableCaseCamel + `) List(ctx context.Context, req *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `ListReq) (res *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `ListRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.Map(req.Filter, gconv.MapOption{Deep: true, OmitEmpty: true})
	if filter == nil {
		filter = map[string]interface{}{}
	}
`
		tplController += `
	allowField := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.ColumnArr().Slice()
	allowField = append(allowField, ` + controllerAlloweFieldList + `)`
		/* if controllerAlloweFieldDiff != `` {
				tplController += `
		allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{` + controllerAlloweFieldDiff + `})).Slice() //移除敏感字段`
			} */
		tplController += `
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	/**--------参数处理 结束--------**/
`
		if option.IsAuthAction {
			actionCode := gstr.CaseCamelLower(myGenThis.tpl.LogicStructName) + `Look`
			actionName := option.CommonName + `-查看`
			myGenThis.genAction(myGenThis.sceneId, actionCode, actionName) // 数据库权限操作处理
			tplController += `
	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, ` + "`" + actionCode + "`" + `)
	if !isAuth {
		field = []string{` + controllerAlloweFieldNoAuth + `}
	}
	/**--------权限验证 结束--------**/
`
		}
		tplController += `
	daoModelThis := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.CtxDaoModel(ctx).Filters(filter)`
		if option.IsCount {
			tplController += `
	count, err := daoModelThis.CountPri()
	if err != nil {
		return
	}`
		}
		tplController += `
	list, err := daoModelThis.Fields(field).HookSelect().Order(req.Sort).Page(req.Page, req.Limit).ListPri()
	if err != nil {
		return
	}

	res = &api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `ListRes{`
		if option.IsCount {
			tplController += `Count: count, `
		}
		tplController += `List:  []api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `ListItem{}}
	list.Structs(&res.List)
	return
}
`
	}
	if option.IsInfo {
		tplController += `
// 详情
func (controllerThis *` + tpl.TableCaseCamel + `) Info(ctx context.Context, req *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `InfoReq) (res *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `InfoRes, err error) {
	/**--------参数处理 开始--------**/
	allowField := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.ColumnArr().Slice()
	allowField = append(allowField, ` + controllerAlloweFieldInfo + `)`
		/* if controllerAlloweFieldDiff != `` {
				tplController += `
		allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{` + controllerAlloweFieldDiff + `})).Slice() //移除敏感字段`
			} */
		tplController += `
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	filter := map[string]interface{}{` + "`id`" + `: req.Id}
	/**--------参数处理 结束--------**/
`
		if option.IsAuthAction {
			actionCode := gstr.CaseCamelLower(myGenThis.tpl.LogicStructName) + `Look`
			actionName := option.CommonName + `-查看`
			myGenThis.genAction(myGenThis.sceneId, actionCode, actionName) // 数据库权限操作处理
			tplController += `
	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, ` + "`" + actionCode + "`" + `)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/
`
		}
		tplController += `
	info, err := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.CtxDaoModel(ctx).Filters(filter).Fields(field).HookSelect().InfoPri()
	if err != nil {
		return
	}
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 29999998, ` + "``" + `)
		return
	}

	res = &api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `InfoRes{}
	info.Struct(&res.Info)
	return
}
`
	}
	if option.IsCreate {
		tplController += `
// 新增
func (controllerThis *` + tpl.TableCaseCamel + `) Create(ctx context.Context, req *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `CreateReq) (res *api.CommonCreateRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true})
	/**--------参数处理 结束--------**/
`
		if option.IsAuthAction {
			actionCode := gstr.CaseCamelLower(myGenThis.tpl.LogicStructName) + `Create`
			actionName := option.CommonName + `-新增`
			myGenThis.genAction(myGenThis.sceneId, actionCode, actionName) // 数据库权限操作处理
			tplController += `
	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, ` + "`" + actionCode + "`" + `)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/
`
		}
		tplController += `
	id, err := service.` + myGenThis.tpl.LogicStructName + `().Create(ctx, data)
	if err != nil {
		return
	}
	res = &api.CommonCreateRes{Id: id}
	return
}
`
	}

	if option.IsUpdate {
		tplController += `
// 修改
func (controllerThis *` + tpl.TableCaseCamel + `) Update(ctx context.Context, req *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `UpdateReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true})
	delete(data, ` + "`idArr`" + `)
	if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ` + "``" + `)
		return
	}
	filter := map[string]interface{}{` + "`id`" + `: req.IdArr}
	/**--------参数处理 结束--------**/
`
		if option.IsAuthAction {
			actionCode := gstr.CaseCamelLower(myGenThis.tpl.LogicStructName) + `Update`
			actionName := option.CommonName + `-编辑`
			myGenThis.genAction(myGenThis.sceneId, actionCode, actionName) // 数据库权限操作处理
			tplController += `
	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, ` + "`" + actionCode + "`" + `)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/
`
		}
		tplController += `
	_, err = service.` + myGenThis.tpl.LogicStructName + `().Update(ctx, filter, data)
	return
}
`
	}

	if option.IsDelete {
		tplController += `
// 删除
func (controllerThis *` + tpl.TableCaseCamel + `) Delete(ctx context.Context, req *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `DeleteReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	filter := map[string]interface{}{` + "`id`" + `: req.IdArr}
	/**--------参数处理 结束--------**/
`
		if option.IsAuthAction {
			actionCode := gstr.CaseCamelLower(myGenThis.tpl.LogicStructName) + `Delete`
			actionName := option.CommonName + `-删除`
			myGenThis.genAction(myGenThis.sceneId, actionCode, actionName) // 数据库权限操作处理
			tplController += `
	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, ` + "`" + actionCode + "`" + `)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/
`
		}
		tplController += `
	_, err = service.` + myGenThis.tpl.LogicStructName + `().Delete(ctx, filter)
	return
}
`
	}

	if option.IsList && tpl.HandlePid.PidField != `` {
		tplController += `
// 列表（树状）
func (controllerThis *` + tpl.TableCaseCamel + `) Tree(ctx context.Context, req *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `TreeReq) (res *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `TreeRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.Map(req.Filter, gconv.MapOption{Deep: true, OmitEmpty: true})
	if filter == nil {
		filter = map[string]interface{}{}
	}

	allowField := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.ColumnArr().Slice()
	allowField = append(allowField, ` + controllerAlloweFieldTree + `)`
		/* if controllerAlloweFieldDiff != `` {
				tplController += `
		allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{` + controllerAlloweFieldDiff + `})).Slice() //移除敏感字段`
			} */
		tplController += `
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	/**--------参数处理 结束--------**/
`
		if option.IsAuthAction {
			actionCode := gstr.CaseCamelLower(myGenThis.tpl.LogicStructName) + `Look`
			actionName := option.CommonName + `-查看`
			myGenThis.genAction(myGenThis.sceneId, actionCode, actionName) // 数据库权限操作处理
			tplController += `
	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, ` + "`" + actionCode + "`" + `)
	if !isAuth {
		field = []string{` + controllerAlloweFieldNoAuth + `}
	}
	/**--------权限验证 结束--------**/
`
		}
		tplController += `
	field = append(field, ` + "`tree`" + `)

	list, err :=dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.CtxDaoModel(ctx).Filters(filter).Fields(field).HookSelect().ListPri()
	if err != nil {
		return
	}
	tree := utils.Tree(list.List(), 0, dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.Columns().` + gstr.CaseCamel(tpl.FieldPrimary) + `, dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.Columns().` + gstr.CaseCamel(tpl.HandlePid.PidField) + `)

	res = &api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `TreeRes{}
	gconv.Structs(tree, &res.Tree)
	return
}
`
	}

	gfile.PutContents(saveFile, tplController)
	utils.GoFileFmt(saveFile)
}

// 后端路由生成
func (myGenThis *myGenHandler) genRouter() {
	option := myGenThis.option
	tpl := myGenThis.tpl

	saveFile := gfile.SelfDir() + `/internal/router/` + option.SceneCode + `.go`

	tplRouter := gfile.GetContents(saveFile)

	//控制器不存在时导入
	importControllerStr := `controller` + tpl.ModuleDirCaseCamel + ` "api/internal/controller/` + option.SceneCode + `/` + tpl.ModuleDirCaseKebab + `"`
	if gstr.Pos(tplRouter, importControllerStr) == -1 {
		tplRouter = gstr.Replace(tplRouter, `"api/internal/middleware"`, importControllerStr+`
	"api/internal/middleware"`, 1)
		//路由生成
		tplRouter = gstr.Replace(tplRouter, `/*--------后端路由自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/`, `group.Group(`+"`"+`/`+tpl.ModuleDirCaseKebab+"`"+`, func(group *ghttp.RouterGroup) {
				group.Bind(controller`+tpl.ModuleDirCaseCamel+`.New`+tpl.TableCaseCamel+`())
			})

			/*--------后端路由自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/`, 1)
		gfile.PutContents(saveFile, tplRouter)
	} else {
		//路由不存在时需生成
		if gstr.Pos(tplRouter, `group.Bind(controller`+tpl.ModuleDirCaseCamel+`.New`+tpl.TableCaseCamel+`())`) == -1 {
			//路由生成
			tplRouter = gstr.Replace(tplRouter, `group.Group(`+"`"+`/`+tpl.ModuleDirCaseKebab+"`"+`, func(group *ghttp.RouterGroup) {`, `group.Group(`+"`"+`/`+tpl.ModuleDirCaseKebab+"`"+`, func(group *ghttp.RouterGroup) {
				group.Bind(controller`+tpl.ModuleDirCaseCamel+`.New`+tpl.TableCaseCamel+`())`, 1)
			gfile.PutContents(saveFile, tplRouter)
		}
	}

	utils.GoFileFmt(saveFile)
}

// 视图模板Index生成
func (myGenThis *myGenHandler) genViewIndex() {
	option := myGenThis.option
	tpl := myGenThis.tpl

	saveFile := gfile.SelfDir() + `/../view/` + option.SceneCode + `/src/views/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/Index.vue`
	if !option.IsCover && gfile.IsFile(saveFile) {
		return
	}

	tplView := `<script setup lang="tsx">
import List from './List.vue'
import Query from './Query.vue'`
	if option.IsCreate || option.IsUpdate {
		tplView += `
import Save from './Save.vue'`
	}
	tplView += `

//搜索
const queryCommon = reactive({
    data: {},
})
provide('queryCommon', queryCommon)

//列表
const listCommon = reactive({
    ref: null as any,
})
provide('listCommon', listCommon)`
	if option.IsCreate || option.IsUpdate {
		tplView += `

//保存
const saveCommon = reactive({
    visible: false,
    title: '', //新增|编辑|复制
    data: {},
})
provide('saveCommon', saveCommon)`
	}
	tplView += `
</script>

<template>
    <el-container class="main-table-container">
        <el-header>
            <query />
        </el-header>

        <list :ref="(el: any) => listCommon.ref = el" />`
	if option.IsCreate || option.IsUpdate {
		tplView += `

        <!-- 加上v-if每次都重新生成组件。可防止不同操作之间的影响；新增操作数据的默认值也能写在save组件内 -->
        <save v-if="saveCommon.visible" />`
	}
	tplView += `
    </el-container>
</template>
`

	gfile.PutContents(saveFile, tplView)
}

// 视图模板List生成
func (myGenThis *myGenHandler) genViewList() {
	option := myGenThis.option
	tpl := myGenThis.tpl

	saveFile := gfile.SelfDir() + `/../view/` + option.SceneCode + `/src/views/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/List.vue`
	if !option.IsCover && gfile.IsFile(saveFile) {
		return
	}

	tableRowHeight := 50
	viewListColumn := ``
	for _, column := range tpl.TableColumnList {
		field := column[`Field`].String()
		fieldCaseCamel := gstr.CaseCamel(field)
		fieldCaseSnake := gstr.CaseSnake(field)
		fieldCaseSnakeOfRemove := gstr.Split(fieldCaseSnake, `_of_`)[0]
		fieldCaseCamelOfRemove := gstr.CaseCamel(fieldCaseSnakeOfRemove)
		fieldSplitArr := gstr.Split(fieldCaseSnakeOfRemove, `_`)
		fieldPrefix := fieldSplitArr[0]
		fieldSuffix := fieldSplitArr[len(fieldSplitArr)-1]

		dataKeyOfColumn := `dataKey: '` + field + `',`
		titleOfColumn := `title: t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `'),`
		keyOfColumn := `key: '` + field + `',`
		alignOfColumn := `align: 'center',`
		widthOfColumn := `width: 150,`
		sortableOfColumn := ``
		hiddenOfColumn := ``
		cellRendererOfColumn := ``

		if garray.NewStrArrayFrom([]string{`DeletedAt`, `DeleteAt`, `DeletedTime`, `DeleteTime`}).Contains(fieldCaseCamel) {
			continue
		} else if garray.NewStrArrayFrom([]string{`UpdatedAt`, `UpdateAt`, `UpdatedTime`, `UpdateTime`}).Contains(fieldCaseCamel) {
			titleOfColumn = `title: t('common.name.updatedAt'),`
			sortableOfColumn = `sortable: true,`
		} else if garray.NewStrArrayFrom([]string{`CreatedAt`, `CreateAt`, `CreatedTime`, `CreateTime`}).Contains(fieldCaseCamel) {
			titleOfColumn = `title: t('common.name.createdAt'),`
			sortableOfColumn = `sortable: true,`
		} else if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` { //主键
			continue
		} else if fieldCaseCamel == `IdPath` && (gstr.Pos(column[`Type`].String(), `varchar`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) && tpl.HandlePid.IsCoexist { //idPath|id_path
			hiddenOfColumn = `hidden: true,`
		} else if garray.NewStrArrayFrom([]string{`status`, `type`, `method`, `pos`, `position`, `gender`}).Contains(fieldSuffix) && ((gstr.Pos(column[`Type`].String(), `int`) != -1 && gstr.Pos(column[`Type`].String(), `point`) == -1) || gstr.Pos(column[`Type`].String(), `char`) != -1) { //status,type,method,pos,position,gender等后缀
			widthOfColumn = `width: 100,`
			cellRendererOfColumn = `cellRenderer: (props: any): any => {
                let tagType = tm('config.const.tagType') as string[]
                let obj = tm('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.status.` + field + `') as { value: any, label: string }[]
                let index = obj.findIndex((item) => { return item.value == props.rowData.` + field + ` })
                return <el-tag type={tagType[index % tagType.length]}>{obj[index]?.label}</el-tag>
            },`
		} else if (garray.NewStrArrayFrom([]string{`icon`, `cover`, `avatar`, `img`, `image`}).Contains(fieldSuffix) || gstr.SubStr(fieldCaseCamelOfRemove, -7) == `ImgList` || gstr.SubStr(fieldCaseCamelOfRemove, -6) == `ImgArr` || gstr.SubStr(fieldCaseCamelOfRemove, -9) == `ImageList` || gstr.SubStr(fieldCaseCamelOfRemove, -8) == `ImageArr`) && (gstr.Pos(column[`Type`].String(), `varchar`) != -1 || gstr.Pos(column[`Type`].String(), `json`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) { //icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀
			widthOfColumn = `width: 100,`
			cellRendererOfColumn = `cellRenderer: (props: any): any => {
                if (!props.rowData.` + field + `) {
                    return
                }`
			if gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
				cellRendererOfColumn += `
                const imageList = [props.rowData.` + field + `]`
			} else {
				cellRendererOfColumn += `
                let imageList: string[]
                if (Array.isArray(props.rowData.` + field + `)) {
                    imageList = props.rowData.` + field + `
                } else {
                    imageList = JSON.parse(props.rowData.` + field + `)
                }`
			}
			cellRendererOfColumn += `
                return [
                    <el-scrollbar wrap-style="display: flex; align-items: center;" view-style="margin: auto;">
                        {imageList.map((item) => {
                            //修改宽高时，可同时修改table属性row-height增加行高，则不会显示滚动条
                            return <el-image style="width: 45px;" src={item} lazy={true} hide-on-click-modal={true} preview-teleported={true} preview-src-list={imageList} />
                        })}
                    </el-scrollbar>
                ]
            },`
		} else if (garray.NewStrArrayFrom([]string{`video`}).Contains(fieldSuffix) || gstr.SubStr(fieldCaseCamelOfRemove, -9) == `VideoList` || gstr.SubStr(fieldCaseCamelOfRemove, -8) == `VideoArr`) && (gstr.Pos(column[`Type`].String(), `varchar`) != -1 || gstr.Pos(column[`Type`].String(), `json`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) { //video,video_list,videoList,video_arr,videoArr等后缀
			if tableRowHeight < 100 {
				tableRowHeight = 100
			}
			cellRendererOfColumn = `cellRenderer: (props: any): any => {
                if (!props.rowData.` + field + `) {
                    return
                }`
			if gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
				cellRendererOfColumn += `
                const videoList = [props.rowData.` + field + `]`
			} else {
				cellRendererOfColumn += `
                let videoList: string[]
                if (Array.isArray(props.rowData.` + field + `)) {
                    videoList = props.rowData.` + field + `
                } else {
                    videoList = JSON.parse(props.rowData.` + field + `)
                }`
			}
			cellRendererOfColumn += `
                return [
                    <el-scrollbar wrap-style="display: flex; align-items: center;" view-style="margin: auto;">
                        {videoList.map((item) => {
                            //修改宽高时，可同时修改table属性row-height增加行高，则不会显示滚动条
                            return <video style="width: 120px; height: 80px;" preload="none" controls={true} src={item} />
                        })}
                    </el-scrollbar>,
                ]
            },`
		} else if garray.NewStrArrayFrom([]string{`list`, `arr`}).Contains(fieldSuffix) && (gstr.Pos(column[`Type`].String(), `json`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) { //list,arr等后缀
			widthOfColumn = `width: 100,`
			cellRendererOfColumn = `cellRenderer: (props: any): any => {
                if (!props.rowData.` + field + `) {
                    return
                }`
			cellRendererOfColumn += `
                let arrList: any[]
                if (Array.isArray(props.rowData.` + field + `)) {
                    arrList = props.rowData.` + field + `
                } else {
                    arrList = JSON.parse(props.rowData.` + field + `)
                }
                let tagType = tm('config.const.tagType') as string[]
                return [
                    <el-scrollbar wrap-style="display: flex; align-items: center;" view-style="margin: auto;">
                        {arrList.map((item, index) => {
                            return [
                                <el-tag style="margin: auto 5px 5px auto;" type={tagType[index % tagType.length]}>
                                    {item}
                                </el-tag>,
                            ]
                        })}
                    </el-scrollbar>,
                ]
            },`
		} else if garray.NewStrArrayFrom([]string{`remark`, `desc`, `msg`, `message`, `intro`, `content`}).Contains(fieldSuffix) && (gstr.Pos(column[`Type`].String(), `varchar`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) { //remark,desc,msg,message,intro,content后缀
			hiddenOfColumn = `hidden: true,`
		} else if gstr.Pos(column[`Type`].String(), `varchar`) != -1 { //varchar类型
			if garray.NewStrArrayFrom([]string{`url`, `link`}).Contains(fieldSuffix) { //url,link后缀
				widthOfColumn = `width: 200,`
			}
		} else if gstr.Pos(column[`Type`].String(), `char`) != -1 { //char类型
			if garray.NewStrArrayFrom([]string{`password`, `passwd`}).Contains(fieldSuffix) && column[`Type`].String() == `char(32)` { //password,passwd后缀
				continue
			} else if garray.NewStrArrayFrom([]string{`salt`}).Contains(fieldSuffix) && tpl.HandlePasswordMap[myGenThis.genHandlePasswordMapKey(field)].IsCoexist { //salt后缀
				continue
			}
		} else if gstr.Pos(column[`Type`].String(), `int`) != -1 && gstr.Pos(column[`Type`].String(), `point`) == -1 { //int等类型
			if field == `pid` { //pid
				dataKeyOfColumn = `dataKey: 'p` + gstr.CaseCamel(tpl.HandelLabelList[0]) + `',`
			} else if field == `level` && tpl.HandlePid.IsCoexist { //level
				widthOfColumn = `width: 100,`
				sortableOfColumn = `sortable: true,`
			} else if garray.NewStrArrayFrom([]string{`sort`, `weight`}).Contains(fieldSuffix) { //sort,weight等后缀
				widthOfColumn = `width: 100,`
				sortableOfColumn = `sortable: true,`
				if option.IsUpdate {
					cellRendererOfColumn = `cellRenderer: (props: any): any => {
                if (props.rowData.edit` + gstr.CaseCamel(field) + `) {
                    let currentRef: any
                    let currentVal = props.rowData.` + field + `
                    return [
                        <el-input-number
                            ref={(el: any) => {
                                currentRef = el
                                el?.focus()
                            }}
                            model-value={currentVal}
                            placeholder={t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.tip.` + field + `')}
                            precision={0}
                            min={0}
                            max={100}
                            step={1}
                            step-strictly={true}
                            controls={false} //控制按钮会导致诸多问题。如：焦点丢失；` + field + `是0或100时，只一个按钮可点击
                            controls-position="right"
                            onChange={(val: number) => (currentVal = val)}
                            onBlur={() => {
                                props.rowData.edit` + gstr.CaseCamel(field) + ` = false
                                if ((currentVal || currentVal === 0) && currentVal != props.rowData.` + field + `) {
                                    handleUpdate({
                                        idArr: [props.rowData.id],
                                        ` + field + `: currentVal,
                                    })
                                        .then((res) => {
                                            props.rowData.` + field + ` = currentVal
                                        })
                                        .catch((error) => {})
                                }
                            }}
                            onKeydown={(event: any) => {
                                switch (event.keyCode) {
                                    // case 27:    //Esc键：Escape
                                    // case 32:    //空格键：" "
                                    case 13: //Enter键：Enter
                                        // props.rowData.edit` + gstr.CaseCamel(field) + ` = false    //也会触发onBlur事件
                                        currentRef?.blur()
                                        break
                                }
                            }}
                        />,
                    ]
                }
                return [
                    <div class="inline-edit" onClick={() => (props.rowData.edit` + gstr.CaseCamel(field) + ` = true)}>
                        {props.rowData.` + field + `}
                    </div>,
                ]
            },`
				}
			} else if garray.NewStrArrayFrom([]string{`id`}).Contains(fieldSuffix) { //id后缀
				if tpl.RelTableMap[field].TableRaw != `` && !tpl.RelTableMap[field].IsRedundRelNameField {
					dataKeyOfColumn = `dataKey: '` + tpl.RelTableMap[field].RelTableField + tpl.RelTableMap[field].RelSuffix + `',`
				}
			} else if garray.NewStrArrayFrom([]string{`is`}).Contains(fieldPrefix) { //is_前缀
				widthOfColumn = `width: 100,`
				cellRendererOfColumn = `cellRenderer: (props: any): any => {
                return [
                    <el-switch
                        model-value={props.rowData.` + field + `}
                        // disabled={true}
                        active-value={1}
                        inactive-value={0}
                        inline-prompt={true}
                        active-text={t('common.yes')}
                        inactive-text={t('common.no')}
                        style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success);"`
				if option.IsUpdate {
					cellRendererOfColumn += `
                        onChange={(val: number) => {
                            handleUpdate({
                                idArr: [props.rowData.id],
                                ` + field + `: val,
                            })
                                .then((res) => {
                                    props.rowData.` + field + ` = val
                                })
                                .catch((error) => {})
                        }}`
				}
				cellRendererOfColumn += `
                    />,
                ]
            },`
			}
		} else if gstr.Pos(column[`Type`].String(), `timestamp`) != -1 || gstr.Pos(column[`Type`].String(), `date`) != -1 { //timestamp或datetime或date类型
			sortableOfColumn = `sortable: true,`
			if gstr.Pos(column[`Type`].String(), `date`) != -1 && gstr.Pos(column[`Type`].String(), `datetime`) == -1 {
				widthOfColumn = `width: 100,`
			}
		} else if gstr.Pos(column[`Type`].String(), `json`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1 { //json类型 //text类型
			widthOfColumn = `width: 200,`
			hiddenOfColumn = `hidden: true,`
		}

		viewListColumn += `
        {
            ` + dataKeyOfColumn + `
            ` + titleOfColumn + `
            ` + keyOfColumn + `
            ` + alignOfColumn + `
            ` + widthOfColumn
		if sortableOfColumn != `` {
			viewListColumn += `
            ` + sortableOfColumn
		}
		if hiddenOfColumn != `` {
			viewListColumn += `
            ` + hiddenOfColumn
		}
		if cellRendererOfColumn != `` {
			viewListColumn += `
            ` + cellRendererOfColumn
		}
		viewListColumn += `
        },`
	}

	tplView := `<script setup lang="tsx">
const { t, tm } = useI18n()

const table = reactive({
    columns: [
        {
            dataKey: 'id',
            title: t('common.name.id'),
            key: 'id',
            align: 'center',
            width: 200,
            fixed: 'left',
            sortable: true,`
	if option.IsUpdate || option.IsDelete {
		tplView += `
            headerCellRenderer: () => {
                const allChecked = table.data.every((item: any) => item.checked)
                const someChecked = table.data.some((item: any) => item.checked)
                return [
                    //阻止冒泡
                    <div class="id-checkbox" onClick={(event: any) => event.stopPropagation()}>
                        <el-checkbox
                            model-value={table.data.length ? allChecked : false}
                            indeterminate={someChecked && !allChecked}
                            onChange={(val: boolean) => {
                                table.data.forEach((item: any) => {
                                    item.checked = val
                                })
                            }}
                        />
                    </div>,
                    <div>{t('common.name.id')}</div>,
                ]
            },
            cellRenderer: (props: any): any => {
                return [<el-checkbox class="id-checkbox" model-value={props.rowData.checked} onChange={(val: boolean) => (props.rowData.checked = val)} />, <div>{props.rowData.id}</div>]
            },`
	}
	tplView += `
        },` + viewListColumn
	if option.IsCreate || option.IsUpdate || option.IsDelete {
		tplView += `
        {
            title: t('common.name.action'),
            key: 'action',
            align: 'center',
            width: 250,
            fixed: 'right',
            cellRenderer: (props: any): any => {
                return [`
		if option.IsUpdate {
			tplView += `
                    <el-button type="primary" size="small" onClick={() => handleEditCopy(props.rowData.id)}>
                        <autoicon-ep-edit />
                        {t('common.edit')}
                    </el-button>,`
		}
		if option.IsDelete {
			tplView += `
                    <el-button type="danger" size="small" onClick={() => handleDelete(props.rowData.id)}>
                        <autoicon-ep-delete />
                        {t('common.delete')}
                    </el-button>,`
		}
		if option.IsCreate {
			tplView += `
                    <el-button type="warning" size="small" onClick={() => handleEditCopy(props.rowData.id, 'copy')}>
                        <autoicon-ep-document-copy />
                        {t('common.copy')}
                    </el-button>,`
		}
		tplView += `
                ]
            },
        },`
	}
	tplView += `
    ] as any,
    data: [],
    loading: false,
    sort: { key: 'id', order: 'desc' } as any,
    handleSort: (sort: any) => {
        table.sort.key = sort.key
        table.sort.order = sort.order
        getList()
    },
})`
	if option.IsCreate || option.IsUpdate {
		tplView += `

const saveCommon = inject('saveCommon') as { visible: boolean; title: string; data: { [propName: string]: any } }`
	}
	if option.IsCreate {
		tplView += `
//新增
const handleAdd = () => {
    saveCommon.data = {}
    saveCommon.title = t('common.add')
    saveCommon.visible = true
}`
	}
	if option.IsDelete {
		tplView += `
//批量删除
const handleBatchDelete = () => {
    const idArr: number[] = []
    table.data.forEach((item: any) => {
        if (item.checked) {
            idArr.push(item.id)
        }
    })
    if (idArr.length) {
        handleDelete(idArr)
    } else {
        ElMessage.error(t('common.tip.selectDelete'))
    }
}`
	}
	if option.IsCreate || option.IsUpdate {
		tplView += `
//编辑|复制
const handleEditCopy = (id: number, type: string = 'edit') => {
    request(t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/info', { id: id })
        .then((res) => {
            saveCommon.data = { ...res.data.info }
            switch (type) {
                case 'edit':
                    saveCommon.data.idArr = [saveCommon.data.id]
                    delete saveCommon.data.id
                    saveCommon.title = t('common.edit')
                    break
                case 'copy':
                    delete saveCommon.data.id
                    saveCommon.title = t('common.copy')
                    break
            }
            saveCommon.visible = true
        })
        .catch(() => {})
}`
	}
	if option.IsDelete {
		tplView += `
//删除
const handleDelete = (idArr: number[]) => {
    ElMessageBox.confirm('', {
        type: 'warning',
        title: t('common.tip.configDelete'),
        center: true,
        showClose: false,
    })
        .then(() => {
            request(t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/del', { idArr: idArr }, true)
                .then((res) => {
                    getList()
                })
                .catch(() => {})
        })
        .catch(() => {})
}`
	}
	if option.IsUpdate {
		tplView += `
//更新
const handleUpdate = async (param: { idArr: number[]; [propName: string]: any }) => {
    await request(t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/update', param, true)
}`
	}
	tplView += `

//分页
const settingStore = useSettingStore()
const pagination = reactive({
    total: 0,
    page: 1,
    size: settingStore.pagination.size,
    sizeList: settingStore.pagination.sizeList,
    layout: settingStore.pagination.layout,
    sizeChange: (val: number) => {
        getList()
    },
    pageChange: (val: number) => {
        getList()
    },
})

const queryCommon = inject('queryCommon') as { data: { [propName: string]: any } }
//列表
const getList = async (resetPage: boolean = false) => {
    if (resetPage) {
        pagination.page = 1
    }
    const param = {
        field: [],
        filter: removeEmptyOfObj(queryCommon.data, true, true),
        sort: table.sort.key + ' ' + table.sort.order,
        page: pagination.page,
        limit: pagination.size,
    }
    table.loading = true
    try {
        const res = await request(t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/list', param)
        table.data = res.data.list?.length ? res.data.list : []
        pagination.total = res.data.count
    } catch (error) {}
    table.loading = false
}
getList()

//暴露组件接口给父组件
defineExpose({
    getList,
})
</script>

<template>
    <el-row class="main-table-tool">
        <el-col :span="16">
            <el-space :size="10" style="height: 100%; margin-left: 10px">`
	if option.IsCreate {
		tplView += `
                <el-button type="primary" @click="handleAdd"> <autoicon-ep-edit-pen />{{ t('common.add') }} </el-button>`
	}
	if option.IsDelete {
		tplView += `
                <el-button type="danger" @click="handleBatchDelete"> <autoicon-ep-delete-filled />{{ t('common.batchDelete') }} </el-button>`
	}
	tplView += `
            </el-space>
        </el-col>
        <el-col :span="8" style="text-align: right">
            <el-space :size="10" style="height: 100%">
                <my-export-button i18nPrefix="` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `" :headerList="table.columns" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/list', param: { filter: queryCommon.data, sort: table.sort.key + ' ' + table.sort.order } }" />
                <el-dropdown max-height="300" :hide-on-click="false">
                    <el-button type="info" :circle="true">
                        <autoicon-ep-hide />
                    </el-button>
                    <template #dropdown>
                        <el-dropdown-menu>
                            <el-dropdown-item v-for="(item, index) in table.columns" :key="index">
                                <el-checkbox v-model="item.hidden">
                                    {{ item.title }}
                                </el-checkbox>
                            </el-dropdown-item>
                        </el-dropdown-menu>
                    </template>
                </el-dropdown>
            </el-space>
        </el-col>
    </el-row>

    <el-main>
        <el-auto-resizer>
            <template #default="{ height, width }">
                <el-table-v2 class="main-table" :columns="table.columns" :data="table.data" :sort-by="table.sort" @column-sort="table.handleSort" :width="width" :height="height" :fixed="true" :row-height="` + gconv.String(tableRowHeight) + `">
                    <template v-if="table.loading" #overlay>
                        <el-icon class="is-loading" color="var(--el-color-primary)" :size="25">
                            <autoicon-ep-loading />
                        </el-icon>
                    </template>
                </el-table-v2>
            </template>
        </el-auto-resizer>
    </el-main>

    <el-row class="main-table-pagination">
        <el-col :span="24">
            <el-pagination
                :total="pagination.total"
                v-model:currentPage="pagination.page"
                v-model:page-size="pagination.size"
                @size-change="pagination.sizeChange"
                @current-change="pagination.pageChange"
                :page-sizes="pagination.sizeList"
                :layout="pagination.layout"
                :background="true"
            />
        </el-col>
    </el-row>
</template>
`

	gfile.PutContents(saveFile, tplView)
}

// 视图模板Query生成
func (myGenThis *myGenHandler) genViewQuery() {
	option := myGenThis.option
	tpl := myGenThis.tpl

	saveFile := gfile.SelfDir() + `/../view/` + option.SceneCode + `/src/views/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/Query.vue`
	if !option.IsCover && gfile.IsFile(saveFile) {
		return
	}

	viewQueryDataInit := ``
	viewQueryField := ``
	for _, column := range tpl.TableColumnList {
		field := column[`Field`].String()
		fieldCaseCamel := gstr.CaseCamel(field)
		fieldCaseSnake := gstr.CaseSnake(field)
		fieldCaseSnakeOfRemove := gstr.Split(fieldCaseSnake, `_of_`)[0]
		fieldCaseCamelOfRemove := gstr.CaseCamel(fieldCaseSnakeOfRemove)
		fieldSplitArr := gstr.Split(fieldCaseSnakeOfRemove, `_`)
		fieldPrefix := fieldSplitArr[0]
		fieldSuffix := fieldSplitArr[len(fieldSplitArr)-1]
		resultStr, _ := gregex.MatchString(`.*\((\d*)\)`, column[`Type`].String())
		/* resultFloat, _ := gregex.MatchString(`.*\((\d*),(\d*)\)`, column[`Type`].String())
		if len(resultFloat) < 3 {
			resultFloat = []string{``, `10`, `2`}
		} */

		if garray.NewStrArrayFrom([]string{`DeletedAt`, `DeleteAt`, `DeletedTime`, `DeleteTime`}).Contains(fieldCaseCamel) {
		} else if garray.NewStrArrayFrom([]string{`UpdatedAt`, `UpdateAt`, `UpdatedTime`, `UpdateTime`}).Contains(fieldCaseCamel) {
		} else if garray.NewStrArrayFrom([]string{`CreatedAt`, `CreateAt`, `CreatedTime`, `CreateTime`}).Contains(fieldCaseCamel) {
			viewQueryDataInit += `
    timeRange: (() => {
        return undefined
        /* const date = new Date()
        return [
            new Date(date.getFullYear(), date.getMonth(), date.getDate(), 0, 0, 0),
            new Date(date.getFullYear(), date.getMonth(), date.getDate(), 23, 59, 59),
        ] */
    })(),
    timeRangeStart: computed(() => {
        if (queryCommon.data.timeRange?.length) {
            return dayjs(queryCommon.data.timeRange[0]).format('YYYY-MM-DD HH:mm:ss')
        }
        return ''
    }),
    timeRangeEnd: computed(() => {
        if (queryCommon.data.timeRange?.length) {
            return dayjs(queryCommon.data.timeRange[1]).format('YYYY-MM-DD HH:mm:ss')
        }
        return ''
    }),`
			viewQueryField += `
        <el-form-item prop="timeRange">
            <el-date-picker v-model="queryCommon.data.timeRange" type="datetimerange" range-separator="-" :default-time="[new Date(2000, 0, 1, 0, 0, 0), new Date(2000, 0, 1, 23, 59, 59)]" :start-placeholder="t('common.name.timeRangeStart')" :end-placeholder="t('common.name.timeRangeEnd')" />
        </el-form-item>`
		} else if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` { //主键
		} else if fieldCaseCamel == `IdPath` && (gstr.Pos(column[`Type`].String(), `varchar`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) && tpl.HandlePid.IsCoexist { //idPath|id_path
		} else if garray.NewStrArrayFrom([]string{`status`, `type`, `method`, `pos`, `position`, `gender`}).Contains(fieldSuffix) && ((gstr.Pos(column[`Type`].String(), `int`) != -1 && gstr.Pos(column[`Type`].String(), `point`) == -1) || gstr.Pos(column[`Type`].String(), `char`) != -1) { //status,type,method,pos,position,gender等后缀
			viewQueryField += `
        <el-form-item prop="` + field + `" style="width: 120px">
            <el-select-v2 v-model="queryCommon.data.` + field + `" :options="tm('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.status.` + field + `')" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" :clearable="true" />
        </el-form-item>`
		} else if (garray.NewStrArrayFrom([]string{`icon`, `cover`, `avatar`, `img`, `image`}).Contains(fieldSuffix) || gstr.SubStr(fieldCaseCamelOfRemove, -7) == `ImgList` || gstr.SubStr(fieldCaseCamelOfRemove, -6) == `ImgArr` || gstr.SubStr(fieldCaseCamelOfRemove, -9) == `ImageList` || gstr.SubStr(fieldCaseCamelOfRemove, -8) == `ImageArr` || garray.NewStrArrayFrom([]string{`video`}).Contains(fieldSuffix) || gstr.SubStr(fieldCaseCamelOfRemove, -9) == `VideoList` || gstr.SubStr(fieldCaseCamelOfRemove, -8) == `VideoArr`) && (gstr.Pos(column[`Type`].String(), `varchar`) != -1 || gstr.Pos(column[`Type`].String(), `json`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) { //icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀 //video,video_list,videoList,video_arr,videoArr等后缀
		} else if garray.NewStrArrayFrom([]string{`list`, `arr`}).Contains(fieldSuffix) && (gstr.Pos(column[`Type`].String(), `json`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) { //list,arr等后缀
		} else if garray.NewStrArrayFrom([]string{`remark`, `desc`, `msg`, `message`, `intro`, `content`}).Contains(fieldSuffix) && (gstr.Pos(column[`Type`].String(), `varchar`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) { //remark,desc,msg,message,intro,content后缀
		} else if gstr.Pos(column[`Type`].String(), `varchar`) != -1 { //varchar类型
			viewQueryField += `
        <el-form-item prop="` + field + `">
            <el-input v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" maxlength="` + resultStr[1] + `" :clearable="true" />
        </el-form-item>`
		} else if gstr.Pos(column[`Type`].String(), `char`) != -1 { //char类型
			if garray.NewStrArrayFrom([]string{`password`, `passwd`}).Contains(fieldSuffix) && column[`Type`].String() == `char(32)` { //password,passwd后缀
			} else if garray.NewStrArrayFrom([]string{`salt`}).Contains(fieldSuffix) && tpl.HandlePasswordMap[myGenThis.genHandlePasswordMapKey(field)].IsCoexist { //salt后缀
			} else {
				viewQueryField += `
        <el-form-item prop="` + field + `">
            <el-input v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" minlength="` + resultStr[1] + `" maxlength="` + resultStr[1] + `" :clearable="true" />
        </el-form-item>`
			}
		} else if gstr.Pos(column[`Type`].String(), `int`) != -1 && gstr.Pos(column[`Type`].String(), `point`) == -1 { //int等类型
			if field == `pid` { //pid
				viewQueryField += `
        <el-form-item prop="` + field + `">
            <my-cascader v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/tree' }" :defaultOptions="[{ id: 0, label: t('common.name.allTopLevel') }]" :props="{ checkStrictly: true, emitPath: false }" />
        </el-form-item>`
			} else if field == `level` && tpl.HandlePid.IsCoexist { //level
				viewQueryField += `
        <el-form-item prop="` + field + `">
            <el-input-number v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" :min="1" :controls="false" />
        </el-form-item>`
			} else if garray.NewStrArrayFrom([]string{`sort`, `weight`}).Contains(fieldSuffix) { //sort,weight等后缀
			} else if garray.NewStrArrayFrom([]string{`id`}).Contains(fieldSuffix) { //id后缀
				apiUrl := tpl.ModuleDirCaseKebab + `/` + gstr.CaseCamelLower(gstr.SubStr(field, 0, -2))
				if tpl.RelTableMap[field].TableRaw != `` {
					relTable := tpl.RelTableMap[field]
					apiUrl = relTable.RelDaoDirCaseCamelLower + `/` + relTable.RelTableCaseCamelLower
				}
				if tpl.RelTableMap[field].RelTableIsExistPidField {
					viewQueryField += `
        <el-form-item prop="` + field + `">
            <my-cascader v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + apiUrl + `/tree' }" :props="{ emitPath: false }" />
        </el-form-item>`
				} else {
					viewQueryField += `
        <el-form-item prop="` + field + `">
            <my-select v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + apiUrl + `/list' }" />
        </el-form-item>`
				}
			} else if garray.NewStrArrayFrom([]string{`is`}).Contains(fieldPrefix) { //is_前缀
				viewQueryField += `
        <el-form-item prop="` + field + `" style="width: 120px">
            <el-select-v2 v-model="queryCommon.data.` + field + `" :options="tm('common.status.whether')" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" :clearable="true" />
        </el-form-item>`
			} else { //默认处理（int等类型）
				/* if gstr.Pos(column[`Type`].String(), `unsigned`) != -1 {
				               viewQueryField += `
				   <el-form-item prop="` + field + `">
				       <el-input-number v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" :min="0" :controls="false" />
				   </el-form-item>`
				           } else {
				               viewQueryField += `
				   <el-form-item prop="` + field + `">
				       <el-input-number v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" :controls="false" />
				   </el-form-item>`
				           } */
			}
		} else if gstr.Pos(column[`Type`].String(), `decimal`) != -1 || gstr.Pos(column[`Type`].String(), `double`) != -1 || gstr.Pos(column[`Type`].String(), `float`) != -1 { //float类型
			/* if gstr.Pos(column[`Type`].String(), `unsigned`) != -1 {
			           viewQueryField += `
			   <el-form-item prop="` + field + `">
			       <el-input-number v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" :min="0" :precision="` + resultFloat[2] + `" :controls="false" />
			   </el-form-item>`
			       } else {
			           viewQueryField += `
			   <el-form-item prop="` + field + `">
			       <el-input-number v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" :precision="` + resultFloat[2] + `" :controls="false" />
			   </el-form-item>`
			       } */
		} else if gstr.Pos(column[`Type`].String(), `timestamp`) != -1 || gstr.Pos(column[`Type`].String(), `date`) != -1 { //timestamp或datetime或date类型
			typeDatePicker := ``
			formatDatePicker := `YYYY-MM-DD HH:mm:ss`
			defaultTimeDatePicker := ``
			if gstr.Pos(column[`Type`].String(), `date`) != -1 && gstr.Pos(column[`Type`].String(), `datetime`) == -1 {
				typeDatePicker = `date`
				formatDatePicker = `YYYY-MM-DD`
			}

			if garray.NewStrArrayFrom([]string{`start`}).Contains(fieldPrefix) { //start_前缀
				typeDatePicker = `datetime`
				if formatDatePicker == `YYYY-MM-DD HH:mm:ss` {
					defaultTimeDatePicker = ` :default-time="new Date(2000, 0, 1, 0, 0, 0)"`
				}
			} else if garray.NewStrArrayFrom([]string{`end`}).Contains(fieldPrefix) { //end_前缀
				typeDatePicker = `datetime`
				if formatDatePicker == `YYYY-MM-DD HH:mm:ss` {
					defaultTimeDatePicker = ` :default-time="new Date(2000, 0, 1, 23, 59, 59)"`
				}
			}

			if typeDatePicker != `` {
				viewQueryField += `
        <el-form-item prop="` + field + `">
            <el-date-picker v-model="queryCommon.data.` + field + `" type="` + typeDatePicker + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" format="` + formatDatePicker + `" value-format="` + formatDatePicker + `"` + defaultTimeDatePicker + ` />
        </el-form-item>`
			}
		} else if gstr.Pos(column[`Type`].String(), `json`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1 { //json类型 //text类型
		} else { //默认处理
			viewQueryField += `
        <el-form-item prop="` + field + `">
            <el-input v-model="queryCommon.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" :clearable="true" />
        </el-form-item>`
		}
	}

	tplView := `<script setup lang="tsx">
import dayjs from 'dayjs'

const { t, tm } = useI18n()

const queryCommon = inject('queryCommon') as { data: { [propName: string]: any } }
queryCommon.data = {
    ...queryCommon.data,` + viewQueryDataInit + `
}
const listCommon = inject('listCommon') as { ref: any }
const queryForm = reactive({
    ref: null as any,
    loading: false,
    submit: () => {
        queryForm.loading = true
        listCommon.ref.getList(true).finally(() => {
            queryForm.loading = false
        })
    },
    reset: () => {
        queryForm.ref.resetFields()
        //queryForm.submit()
    },
})
</script>

<template>
    <el-form class="query-form" :ref="(el: any) => queryForm.ref = el" :model="queryCommon.data" :inline="true" @keyup.enter="queryForm.submit">
        <el-form-item prop="id">
            <el-input-number v-model="queryCommon.data.id" :placeholder="t('common.name.id')" :min="1" :controls="false" />
        </el-form-item>` + viewQueryField + `
        <el-form-item>
            <el-button type="primary" @click="queryForm.submit" :loading="queryForm.loading"> <autoicon-ep-search />{{ t('common.query') }} </el-button>
            <el-button type="info" @click="queryForm.reset"> <autoicon-ep-circle-close />{{ t('common.reset') }} </el-button>
        </el-form-item>
    </el-form>
</template>
`

	gfile.PutContents(saveFile, tplView)
}

// 视图模板Save生成
func (myGenThis *myGenHandler) genViewSave() {
	option := myGenThis.option
	tpl := myGenThis.tpl

	if !(option.IsCreate || option.IsUpdate) {
		return
	}
	saveFile := gfile.SelfDir() + `/../view/` + option.SceneCode + `/src/views/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/Save.vue`
	if !option.IsCover && gfile.IsFile(saveFile) {
		return
	}

	viewSaveImport := ``
	viewSaveParamHandle := ``
	viewSaveDataInitBefore := ``
	viewSaveDataInitAfter := ``
	viewSaveRule := ``
	viewSaveField := ``
	viewFieldHandle := ``
	for _, column := range tpl.TableColumnList {
		field := column[`Field`].String()
		fieldCaseCamel := gstr.CaseCamel(field)
		fieldCaseSnake := gstr.CaseSnake(field)
		fieldCaseSnakeOfRemove := gstr.Split(fieldCaseSnake, `_of_`)[0]
		fieldCaseCamelOfRemove := gstr.CaseCamel(fieldCaseSnakeOfRemove)
		fieldSplitArr := gstr.Split(fieldCaseSnakeOfRemove, `_`)
		fieldPrefix := fieldSplitArr[0]
		fieldSuffix := fieldSplitArr[len(fieldSplitArr)-1]
		comment := gstr.Trim(gstr.ReplaceByArray(column[`Comment`].String(), g.SliceStr{
			"\n", ` `,
			"\r", ` `,
			`"`, `\"`,
		}))
		resultStr, _ := gregex.MatchString(`.*\((\d*)\)`, column[`Type`].String())
		resultFloat, _ := gregex.MatchString(`.*\((\d*),(\d*)\)`, column[`Type`].String())
		if len(resultFloat) < 3 {
			resultFloat = []string{``, `10`, `2`}
		}

		if garray.NewStrArrayFrom([]string{`DeletedAt`, `DeleteAt`, `DeletedTime`, `DeleteTime`}).Contains(fieldCaseCamel) {
		} else if garray.NewStrArrayFrom([]string{`UpdatedAt`, `UpdateAt`, `UpdatedTime`, `UpdateTime`}).Contains(fieldCaseCamel) {
		} else if garray.NewStrArrayFrom([]string{`CreatedAt`, `CreateAt`, `CreatedTime`, `CreateTime`}).Contains(fieldCaseCamel) {
		} else if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` { //主键
		} else if fieldCaseCamel == `IdPath` && (gstr.Pos(column[`Type`].String(), `varchar`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) && tpl.HandlePid.IsCoexist { //idPath|id_path
		} else if garray.NewStrArrayFrom([]string{`status`, `type`, `method`, `pos`, `position`, `gender`}).Contains(fieldSuffix) && ((gstr.Pos(column[`Type`].String(), `int`) != -1 && gstr.Pos(column[`Type`].String(), `point`) == -1) || gstr.Pos(column[`Type`].String(), `char`) != -1) { //status,type,method,pos,position,gender等后缀
			isStr := true
			if gstr.Pos(column[`Type`].String(), `int`) != -1 && gstr.Pos(column[`Type`].String(), `point`) == -1 {
				isStr = false
			}
			statusList := myGenThis.genStatusList(comment, isStr)
			defaultVal := column[`Default`].String()
			if defaultVal == `` {
				defaultVal = statusList[0][0]
			}
			if isStr {
				viewSaveDataInitBefore += `
        ` + field + `: '` + defaultVal + `',`
			} else {
				viewSaveDataInitBefore += `
        ` + field + `: ` + defaultVal + `,`
			}
			viewSaveRule += `
        ` + field + `: [
            { type: 'enum', enum: (tm('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.status.` + field + `') as any).map((item: any) => item.value), trigger: 'change', message: t('validation.select') },
        ],`
			viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" prop="` + field + `">`
			//超过5个状态用select组件，小于5个用radio组件
			if len(statusList) > 5 {
				viewSaveField += `
                    <el-select-v2 v-model="saveForm.data.` + field + `" :options="tm('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.status.` + field + `')" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" :clearable="false" />`
			} else {
				viewSaveField += `
                    <el-radio-group v-model="saveForm.data.` + field + `">
                        <el-radio v-for="(item, index) in (tm('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.status.` + field + `') as any)" :key="index" :label="item.value">
                            {{ item.label }}
                        </el-radio>
                    </el-radio-group>`
			}
			viewSaveField += `
                </el-form-item>`
		} else if (garray.NewStrArrayFrom([]string{`icon`, `cover`, `avatar`, `img`, `image`}).Contains(fieldSuffix) || gstr.SubStr(fieldCaseCamelOfRemove, -7) == `ImgList` || gstr.SubStr(fieldCaseCamelOfRemove, -6) == `ImgArr` || gstr.SubStr(fieldCaseCamelOfRemove, -9) == `ImageList` || gstr.SubStr(fieldCaseCamelOfRemove, -8) == `ImageArr`) && (gstr.Pos(column[`Type`].String(), `varchar`) != -1 || gstr.Pos(column[`Type`].String(), `json`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) { //icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀
			multipleStr := ``
			if gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
				viewSaveRule += `
        ` + field + `: [
            { type: 'string', max: ` + resultStr[1] + `, trigger: 'blur', message: t('validation.max.string', { max: ` + resultStr[1] + ` }) },
            { type: 'url', trigger: 'change', message: t('validation.upload') },
        ],`
			} else {
				multipleStr = ` :multiple="true"`
				requiredStr := ``
				if !column[`Null`].Bool() {
					requiredStr = ` required: true,`
				}
				viewSaveRule += `
        ` + field + `: [
            { type: 'array',` + requiredStr + ` trigger: 'change', message: t('validation.upload'), defaultField: { type: 'url', message: t('validation.url') } },
            // { type: 'array',` + requiredStr + ` max: 10, trigger: 'change', message: t('validation.max.upload', { max: 10 }), defaultField: { type: 'url', message: t('validation.url') } },
        ],`
			}
			viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" prop="` + field + `">
                    <my-upload v-model="saveForm.data.` + field + `" accept="image/*"` + multipleStr + ` />
                </el-form-item>`
		} else if (garray.NewStrArrayFrom([]string{`video`}).Contains(fieldSuffix) || gstr.SubStr(fieldCaseCamelOfRemove, -9) == `VideoList` || gstr.SubStr(fieldCaseCamelOfRemove, -8) == `VideoArr`) && (gstr.Pos(column[`Type`].String(), `varchar`) != -1 || gstr.Pos(column[`Type`].String(), `json`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) { //video,video_list,videoList,video_arr,videoArr等后缀
			multipleStr := ``
			if gstr.Pos(column[`Type`].String(), `varchar`) != -1 {
				viewSaveRule += `
        ` + field + `: [
            { type: 'string', max: ` + resultStr[1] + `, trigger: 'blur', message: t('validation.max.string', { max: ` + resultStr[1] + ` }) },
            { type: 'url', trigger: 'change', message: t('validation.upload') },
        ],`
			} else {
				multipleStr = ` :multiple="true"`
				requiredStr := ``
				if !column[`Null`].Bool() {
					requiredStr = ` required: true,`
				}
				viewSaveRule += `
        ` + field + `: [
            { type: 'array',` + requiredStr + ` trigger: 'change', message: t('validation.upload'), defaultField: { type: 'url', message: t('validation.url') } },
            // { type: 'array',` + requiredStr + ` max: 10, trigger: 'change', message: t('validation.max.upload', { max: 10 }), defaultField: { type: 'url', message: t('validation.url') } },
        ],`
			}
			viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" prop="` + field + `">
                    <my-upload v-model="saveForm.data.` + field + `" accept="video/*" :isImage="false"` + multipleStr + ` />
                </el-form-item>`
		} else if garray.NewStrArrayFrom([]string{`list`, `arr`}).Contains(fieldSuffix) && (gstr.Pos(column[`Type`].String(), `json`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) { //list,arr等后缀
			viewSaveDataInitBefore += `
        ` + field + `: [],`
			requiredStr := ``
			if !column[`Null`].Bool() {
				requiredStr = ` required: true,`
			}
			viewSaveRule += `
        ` + field + `: [
            { type: 'array',` + requiredStr + ` trigger: 'change', message: t('validation.required') },
            // { type: 'array',` + requiredStr + ` max: 10, trigger: 'change', message: t('validation.max.array', { max: 10 }), defaultField: { type: 'string', message: t('validation.input') } },
        ],`
			viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" prop="` + field + `">
                    <el-tag v-for="(item, index) in saveForm.data.` + field + `" :type="` + field + `Handle.tagType[index % ` + field + `Handle.tagType.length]" @close="` + field + `Handle.delValue(item)" :key="index" :closable="true" style="margin-right: 10px;">
                        {{ item }}
                    </el-tag>
                    <!-- <el-input-number v-if="` + field + `Handle.visible" :ref="(el: any) => ` + field + `Handle.ref = el" v-model="` + field + `Handle.value" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" @keyup.enter="` + field + `Handle.addValue" @blur="` + field + `Handle.addValue" size="small" style="width: 100px;" :controls="false" /> -->
                    <el-input v-if="` + field + `Handle.visible" :ref="(el: any) => ` + field + `Handle.ref = el" v-model="` + field + `Handle.value" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" @keyup.enter="` + field + `Handle.addValue" @blur="` + field + `Handle.addValue" size="small" style="width: 100px;" />
                    <el-button v-else type="primary" size="small" @click="` + field + `Handle.visibleChange">
                        <autoicon-ep-plus />{{ t('common.add') }}
                    </el-button>
                </el-form-item>`
			viewFieldHandle += `

const ` + field + `Handle = reactive({
    ref: null as any,
    visible: false,
    value: undefined,
    tagType: tm('config.const.tagType') as string[],
    visibleChange: () => {
        ` + field + `Handle.visible = true
        nextTick(() => {
            ` + field + `Handle.ref?.focus()
        })
    },
    addValue: () => {
        if (` + field + `Handle.value) {
            saveForm.data.` + field + `.push(` + field + `Handle.value)
        }
        ` + field + `Handle.visible = false
        ` + field + `Handle.value = undefined
    },
    delValue: (item: any) => {
        saveForm.data.` + field + `.splice(saveForm.data.` + field + `.indexOf(item), 1)
    },
})`
		} else if garray.NewStrArrayFrom([]string{`remark`, `desc`, `msg`, `message`, `intro`, `content`}).Contains(fieldSuffix) && (gstr.Pos(column[`Type`].String(), `varchar`) != -1 || gstr.Pos(column[`Type`].String(), `text`) != -1) { //remark,desc,msg,message,intro,content后缀
			if gstr.Pos(column[`Type`].String(), `text`) != -1 {
				viewSaveRule += `
        ` + field + `: [
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],`
				viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" prop="` + field + `">
                    <my-editor v-model="saveForm.data.` + field + `" />
                </el-form-item>`
			} else {
				viewSaveRule += `
        ` + field + `: [
            { type: 'string', max: ` + resultStr[1] + `, trigger: 'blur', message: t('validation.max.string', { max: ` + resultStr[1] + ` }) },
        ],`
				viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" prop="` + field + `">
                    <el-input v-model="saveForm.data.` + field + `" type="textarea" :autosize="{ minRows: 3 }" maxlength="` + resultStr[1] + `" :show-word-limit="true" />
                </el-form-item>`
			}
		} else if gstr.Pos(column[`Type`].String(), `varchar`) != -1 { //varchar类型
			ruleStr := ``
			requiredStr := ``
			viewSaveFieldTip := ` />`
			if column[`Key`].String() == `UNI` {
				if !column[`Null`].Bool() {
					requiredStr = ` required: true,`
				}
				viewSaveFieldTip = ` style="max-width: 250px" />
                    <label>
                        <el-alert :title="t('common.tip.notDuplicate')" type="info" :show-icon="true" :closable="false" />
                    </label>`
			}
			if garray.NewStrArrayFrom([]string{`name`}).Contains(fieldSuffix) { //name后缀
				if gstr.SubStr(gstr.CaseCamel(tpl.FieldPrimary), 0, -2)+`Name` == fieldCaseCamel {
					requiredStr = ` required: true,`
				}
				ruleStr += `
            { pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') },`
			} else if garray.NewStrArrayFrom([]string{`code`}).Contains(fieldSuffix) { //code后缀
				ruleStr += `
            { pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') },`
			} else if garray.NewStrArrayFrom([]string{`phone`, `mobile`}).Contains(fieldSuffix) { //phone,mobile后缀
				ruleStr += `
            { pattern: /^1[3-9]\d{9}$/, trigger: 'blur', message: t('validation.phone') },`
			} else if garray.NewStrArrayFrom([]string{`url`, `link`}).Contains(fieldSuffix) { //url,link后缀
				ruleStr += `
            { type: 'url', trigger: 'change', message: t('validation.url') },`
			}
			viewSaveRule += `
        ` + field + `: [
            { type: 'string',` + requiredStr + ` max: ` + resultStr[1] + `, trigger: 'blur', message: t('validation.max.string', { max: ` + resultStr[1] + ` }) },` + ruleStr + `
        ],`
			viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" prop="` + field + `">
                    <el-input v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" maxlength="` + resultStr[1] + `" :show-word-limit="true" :clearable="true"` + viewSaveFieldTip + `
                </el-form-item>`
		} else if gstr.Pos(column[`Type`].String(), `char`) != -1 { //char类型
			if garray.NewStrArrayFrom([]string{`password`, `passwd`}).Contains(fieldSuffix) && column[`Type`].String() == `char(32)` { //password,passwd后缀
				viewSaveImportTmp := `
import md5 from 'js-md5'`
				if gstr.Pos(viewSaveImport, viewSaveImportTmp) == -1 {
					viewSaveImport += viewSaveImportTmp
				}
				viewSaveParamHandle += `
            param.` + field + ` ? param.` + field + ` = md5(param.` + field + `) : delete param.` + field

				viewSaveRule += `
        ` + field + `: [
            { type: 'string', required: computed((): boolean => { return saveForm.data.idArr?.length ? false : true; }), min: 6, max: 20, trigger: 'blur', message: t('validation.between.string', { min: 6, max: 20 }) },
        ],`
				viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" prop="` + field + `">
                    <el-input v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" minlength="6" maxlength="20" :show-word-limit="true" :clearable="true" :show-password="true" style="max-width: 250px" />
                    <label v-if="saveForm.data.idArr?.length">
                        <el-alert :title="t('common.tip.notRequired')" type="info" :show-icon="true" :closable="false" />
                    </label>
                </el-form-item>`
			} else if garray.NewStrArrayFrom([]string{`salt`}).Contains(fieldSuffix) && tpl.HandlePasswordMap[myGenThis.genHandlePasswordMapKey(field)].IsCoexist { //salt后缀
			} else {
				ruleStr := ``
				requiredStr := ``
				viewSaveFieldTip := ` />`
				if column[`Key`].String() == `UNI` {
					if !column[`Null`].Bool() {
						requiredStr = ` required: true,`
					}
					viewSaveFieldTip = ` style="max-width: 250px" />
                    <label>
                        <el-alert :title="t('common.tip.notDuplicate')" type="info" :show-icon="true" :closable="false" />
                    </label>`
				}
				viewSaveRule += `
        ` + field + `: [
            { type: 'string',` + requiredStr + ` len: ` + resultStr[1] + `, trigger: 'blur', message: t('validation.size.string', { size: ` + resultStr[1] + ` }) },` + ruleStr + `
        ],`
				viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" prop="` + field + `">
                    <el-input v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" minlength="` + resultStr[1] + `" maxlength="` + resultStr[1] + `" :show-word-limit="true" :clearable="true"` + viewSaveFieldTip + `
                </el-form-item>`
			}
		} else if gstr.Pos(column[`Type`].String(), `int`) != -1 && gstr.Pos(column[`Type`].String(), `point`) == -1 { //int等类型
			if field == `pid` { //pid
				viewSaveParamHandle += `
            if (param.` + field + ` === undefined) {
                param.` + field + ` = 0
            }`
				viewSaveRule += `
        ` + field + `: [
            { type: 'integer', min: 0, trigger: 'change', message: t('validation.select') },
        ],`
				viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" prop="` + field + `">
                    <my-cascader v-model="saveForm.data.` + field + `" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/tree', param: { filter: { excIdArr: saveForm.data.idArr } } }" :props="{ checkStrictly: true, emitPath: false }" />
                </el-form-item>`
			} else if field == `level` && tpl.HandlePid.IsCoexist { //level
			} else if garray.NewStrArrayFrom([]string{`sort`, `weight`}).Contains(fieldSuffix) { //sort,weight等后缀
				defaultVal := column[`Default`].Int()
				if defaultVal != 0 {
					viewSaveDataInitBefore += `
        ` + field + `: ` + gconv.String(defaultVal) + `,`
				}
				viewSaveRule += `
        ` + field + `: [
            { type: 'integer', min: 0, max: 100, trigger: 'change', message: t('validation.between.number', { min: 0, max: 100 }) },
        ],`
				viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" prop="` + field + `">
                    <el-input-number v-model="saveForm.data.` + field + `" :precision="0" :min="0" :max="100" :step="1" :step-strictly="true" controls-position="right" :value-on-clear="` + gconv.String(defaultVal) + `" />
                    <label>
                        <el-alert :title="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.tip.` + field + `')" type="info" :show-icon="true" :closable="false" />
                    </label>
                </el-form-item>`
			} else if garray.NewStrArrayFrom([]string{`id`}).Contains(fieldSuffix) { //id后缀
				apiUrl := tpl.ModuleDirCaseKebab + `/` + gstr.CaseCamelLower(gstr.SubStr(field, 0, -2))
				if tpl.RelTableMap[field].TableRaw != `` {
					relTable := tpl.RelTableMap[field]
					apiUrl = relTable.RelDaoDirCaseCamelLower + `/` + relTable.RelTableCaseCamelLower
				}
				viewSaveParamHandle += `
            if (param.` + field + ` === undefined) {
                param.` + field + ` = 0
            }`
				viewSaveRule += `
        ` + field + `: [
            { type: 'integer', min: 1, trigger: 'change', message: t('validation.select') },
        ],`
				if tpl.RelTableMap[field].RelTableIsExistPidField {
					viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" prop="` + field + `">
                    <my-cascader v-model="saveForm.data.` + field + `" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + apiUrl + `/tree' }" :props="{ emitPath: false }" />
                </el-form-item>`
				} else {
					viewSaveDataInitAfter += `
        ` + field + `: saveCommon.data.` + field + ` ? saveCommon.data.` + field + ` : undefined,`
					viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" prop="` + field + `">
                    <my-select v-model="saveForm.data.` + field + `" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + apiUrl + `/list' }" />
                </el-form-item>`
				}
			} else if garray.NewStrArrayFrom([]string{`is`}).Contains(fieldPrefix) { //is_前缀
				defaultVal := column[`Default`].Int()
				if defaultVal != 0 {
					viewSaveDataInitBefore += `
        ` + field + `: ` + gconv.String(defaultVal) + `,`
				}
				viewSaveRule += `
        ` + field + `: [
            { type: 'enum', enum: (tm('common.status.whether') as any).map((item: any) => item.value), trigger: 'change', message: t('validation.select') },
        ],`
				viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" prop="` + field + `">
                    <el-switch v-model="saveForm.data.` + field + `" :active-value="1" :inactive-value="0" :inline-prompt="true" :active-text="t('common.yes')" :inactive-text="t('common.no')" style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success);" />
                </el-form-item>`
			} else { //默认处理（int等类型）
				if gstr.Pos(column[`Type`].String(), `unsigned`) != -1 {
					defaultVal := column[`Default`].Uint()
					if defaultVal != 0 {
						viewSaveDataInitBefore += `
        ` + field + `: ` + gconv.String(defaultVal) + `,`
					}
					viewSaveRule += `
        ` + field + `: [
            { type: 'integer', min: 0, trigger: 'change', message: t('validation.min.number', { min: 0 }) },
        ],`
					viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" prop="` + field + `">
                    <el-input-number v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" :min="0" :controls="false" :value-on-clear="` + gconv.String(defaultVal) + `" />
                </el-form-item>`
				} else {
					defaultVal := column[`Default`].Int()
					if defaultVal != 0 {
						viewSaveDataInitBefore += `
        ` + field + `: ` + gconv.String(defaultVal) + `,`
					}
					viewSaveRule += `
        ` + field + `: [
            { type: 'integer', trigger: 'change', message: t('validation.input') },
        ],`
					viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" prop="` + field + `">
                    <el-input-number v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" :controls="false" :value-on-clear="` + gconv.String(defaultVal) + `" />
                </el-form-item>`
				}
			}
		} else if gstr.Pos(column[`Type`].String(), `decimal`) != -1 || gstr.Pos(column[`Type`].String(), `double`) != -1 || gstr.Pos(column[`Type`].String(), `float`) != -1 { //float类型
			defaultVal := column[`Default`].Float64()
			if defaultVal != 0 {
				viewSaveDataInitBefore += `
        ` + field + `: ` + gconv.String(defaultVal) + `,`
			}
			if gstr.Pos(column[`Type`].String(), `unsigned`) != -1 {
				viewSaveRule += `
        ` + field + `: [
            { type: 'number'/* 'float' */, min: 0, trigger: 'change', message: t('validation.min.number', { min: 0 }) },    // 类型float值为0时验证不能通过
        ],`
				viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" prop="` + field + `">
                    <el-input-number v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" :min="0" :precision="` + resultFloat[2] + `" :controls="false" :value-on-clear="` + gconv.String(defaultVal) + `" />
                </el-form-item>`
			} else {
				viewSaveRule += `
        ` + field + `: [
            { type: 'number'/* 'float' */, trigger: 'change', message: t('validation.input') },    // 类型float值为0时验证不能通过
        ],`
				viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" prop="` + field + `">
                    <el-input-number v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" :precision="` + resultFloat[2] + `" :controls="false" :value-on-clear="` + gconv.String(defaultVal) + `" />
                </el-form-item>`
			}
		} else if gstr.Pos(column[`Type`].String(), `timestamp`) != -1 || gstr.Pos(column[`Type`].String(), `date`) != -1 { //timestamp或datetime或date类型
			typeDatePicker := `datetime`
			formatDatePicker := `YYYY-MM-DD HH:mm:ss`
			defaultTimeDatePicker := ``
			if gstr.Pos(column[`Type`].String(), `date`) != -1 && gstr.Pos(column[`Type`].String(), `datetime`) == -1 {
				typeDatePicker = `date`
				formatDatePicker = `YYYY-MM-DD`
			}
			requiredStr := ``
			if !column[`Null`].Bool() && column[`Default`].String() == `` {
				requiredStr = ` required: true,`
			}

			if garray.NewStrArrayFrom([]string{`start`}).Contains(fieldPrefix) && formatDatePicker == `YYYY-MM-DD HH:mm:ss` { //start_前缀
				defaultTimeDatePicker = ` :default-time="new Date(2000, 0, 1, 0, 0, 0)"`
			} else if garray.NewStrArrayFrom([]string{`end`}).Contains(fieldPrefix) && formatDatePicker == `YYYY-MM-DD HH:mm:ss` { //end_前缀
				defaultTimeDatePicker = ` :default-time="new Date(2000, 0, 1, 23, 59, 59)"`
			}

			viewSaveRule += `
        ` + field + `: [
            { type: 'string',` + requiredStr + ` trigger: 'change', message: t('validation.select') },
        ],`
			viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" prop="` + field + `">
                    <el-date-picker v-model="saveForm.data.` + field + `" type="` + typeDatePicker + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" format="` + formatDatePicker + `" value-format="` + formatDatePicker + `"` + defaultTimeDatePicker + ` />
                </el-form-item>`
		} else if gstr.Pos(column[`Type`].String(), `json`) != -1 { //json类型
			requiredStr := ``
			if !column[`Null`].Bool() {
				requiredStr = `
                required: true,`
			}
			viewSaveRule += `
        ` + field + `: [
            {
                type: 'object',` + requiredStr + `
                /* fields: {
                    xxxx: { type: 'string', required: true, message: 'xxxx' + t('validation.required') },
                    xxxx: { type: 'integer', required: true, min: 1, message: 'xxxx' + t('validation.min.number', { min: 1 }) },
                }, */
                transform(value: any) {
                    if (value === '' || value === null || value === undefined) {
                        return undefined
                    }
                    try {
                        return JSON.parse(value)
                    } catch (e) {
                        return value
                    }
                },
                trigger: 'blur',
                message: t('validation.json'),
            },
        ],`
			viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" prop="` + field + `">
                    <el-alert :title="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.tip.` + field + `')" type="info" :show-icon="true" :closable="false" />
                    <el-input v-model="saveForm.data.` + field + `" type="textarea" :autosize="{ minRows: 3 }" />
                </el-form-item>`
		} else if gstr.Pos(column[`Type`].String(), `text`) != -1 { //text类型
			viewSaveRule += `
        ` + field + `: [],`
			viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" prop="` + field + `">
                    <my-editor v-model="saveForm.data.` + field + `" />
                </el-form-item>`
		} else { //默认处理
			viewSaveRule += `
        ` + field + `: [],`
			viewSaveField += `
                <el-form-item :label="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" prop="` + field + `">
                    <el-input v-model="saveForm.data.` + field + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + field + `')" :clearable="true" />
                </el-form-item>`
		}
	}

	tplView := `<script setup lang="tsx">` + viewSaveImport + `
const { t, tm } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean; title: string; data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {` + viewSaveDataInitBefore + `
        ...saveCommon.data,` + viewSaveDataInitAfter + `
    } as { [propName: string]: any },
    rules: {` + viewSaveRule + `
    } as any,
    submit: () => {
        saveForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return false
            }
            saveForm.loading = true
            const param = removeEmptyOfObj(saveForm.data)` + viewSaveParamHandle + `
            try {
                if (param?.idArr?.length > 0) {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/update', param, true)
                } else {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/create', param, true)
                }
                listCommon.ref.getList(true)
                saveCommon.visible = false
            } catch (error) {}
            saveForm.loading = false
        })
    },
})

const saveDrawer = reactive({
    ref: null as any,
    size: useSettingStore().saveDrawer.size,
    beforeClose: (done: Function) => {
        if (useSettingStore().saveDrawer.isTipClose) {
            ElMessageBox.confirm('', {
                type: 'info',
                title: t('common.tip.configExit'),
                center: true,
                showClose: false,
            })
                .then(() => {
                    done()
                })
                .catch(() => {})
        } else {
            done()
        }
    },
    buttonClose: () => {
        //saveCommon.visible = false
        saveDrawer.ref.handleClose() //会触发beforeClose
    },
})` + viewFieldHandle + `
</script>

<template>
    <el-drawer class="save-drawer" :ref="(el: any) => saveDrawer.ref = el" v-model="saveCommon.visible" :title="saveCommon.title" :size="saveDrawer.size" :before-close="saveDrawer.beforeClose">
        <el-scrollbar>
            <el-form :ref="(el: any) => saveForm.ref = el" :model="saveForm.data" :rules="saveForm.rules" label-width="auto" :status-icon="true" :scroll-to-error="true">` + viewSaveField + `
            </el-form>
        </el-scrollbar>
        <template #footer>
            <el-button @click="saveDrawer.buttonClose">{{ t('common.cancel') }}</el-button>
            <el-button type="primary" @click="saveForm.submit" :loading="saveForm.loading">
                {{ t('common.save') }}
            </el-button>
        </template>
    </el-drawer>
</template>
`

	gfile.PutContents(saveFile, tplView)
}

// 视图模板I18n生成
func (myGenThis *myGenHandler) genViewI18n() {
	option := myGenThis.option
	tpl := myGenThis.tpl

	saveFile := gfile.SelfDir() + `/../view/` + option.SceneCode + `/src/i18n/language/zh-cn/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `.ts`
	if !option.IsCover && gfile.IsFile(saveFile) {
		return
	}

	viewI18nName := ``
	viewI18nStatus := ``
	viewI18nTip := ``
	for _, column := range tpl.TableColumnList {
		field := column[`Field`].String()
		fieldCaseCamel := gstr.CaseCamel(field)
		fieldCaseSnake := gstr.CaseSnake(field)
		fieldCaseSnakeOfRemove := gstr.Split(fieldCaseSnake, `_of_`)[0]
		// fieldCaseCamelOfRemove := gstr.CaseCamel(fieldCaseSnakeOfRemove)
		fieldSplitArr := gstr.Split(fieldCaseSnakeOfRemove, `_`)
		// fieldPrefix := fieldSplitArr[0]
		fieldSuffix := fieldSplitArr[len(fieldSplitArr)-1]
		tmp, _ := gregex.MatchString(`[^\n\r\.。:：\(（]*`, column[`Comment`].String())
		fieldName := gstr.Trim(tmp[0])
		comment := gstr.Trim(gstr.ReplaceByArray(column[`Comment`].String(), g.SliceStr{
			"\n", ` `,
			"\r", ` `,
			`"`, `\"`,
		}))
		tip := gstr.Replace(comment, fieldName, ``, 1)
		tmp, _ = gregex.MatchString(`\n\r\.。:：\(（`, column[`Comment`].String())
		if len(tmp) > 0 {
			gstr.TrimLeft(tip, tmp[0])
		}
		for _, v := range []string{"\n", "\r", `.`, `。`, `:`, `：`, `(`, `（`, `)`, `）`, ` `, `,`, `，`, `;`, `；`} {
			tip = gstr.Trim(tip, v)
		}
		tip = gstr.ReplaceByArray(tip, g.SliceStr{
			`\"`, `"`,
			`}`, `' + "{'}'}" + '`,
			`{"`, `' + "{'{'}" + '"`,
		})

		if garray.NewStrArrayFrom([]string{`DeletedAt`, `DeleteAt`, `DeletedTime`, `DeleteTime`}).Contains(fieldCaseCamel) {
			continue
		} else if garray.NewStrArrayFrom([]string{`UpdatedAt`, `UpdateAt`, `UpdatedTime`, `UpdateTime`}).Contains(fieldCaseCamel) {
			continue
		} else if garray.NewStrArrayFrom([]string{`CreatedAt`, `CreateAt`, `CreatedTime`, `CreateTime`}).Contains(fieldCaseCamel) {
			continue
		} else if column[`Key`].String() == `PRI` && column[`Extra`].String() == `auto_increment` { //主键
			continue
		} else if garray.NewStrArrayFrom([]string{`salt`}).Contains(fieldSuffix) && tpl.HandlePasswordMap[myGenThis.genHandlePasswordMapKey(field)].IsCoexist { //salt后缀
			continue
		} else if garray.NewStrArrayFrom([]string{`status`, `type`, `method`, `pos`, `position`, `gender`}).Contains(fieldSuffix) && ((gstr.Pos(column[`Type`].String(), `int`) != -1 && gstr.Pos(column[`Type`].String(), `point`) == -1) || gstr.Pos(column[`Type`].String(), `char`) != -1) { //status,type,method,pos,position,gender等后缀
			isStr := true
			if gstr.Pos(column[`Type`].String(), `int`) != -1 && gstr.Pos(column[`Type`].String(), `point`) == -1 {
				isStr = false
			}
			statusList := myGenThis.genStatusList(comment, isStr)
			viewI18nStatus += `
        ` + field + `: [`
			for _, status := range statusList {
				if isStr {
					viewI18nStatus += `
            { value: '` + status[0] + `', label: '` + status[1] + `' },`
				} else {
					viewI18nStatus += `
            { value: ` + status[0] + `, label: '` + status[1] + `' },`
				}
			}
			viewI18nStatus += `
        ],`
		} else if gstr.Pos(column[`Type`].String(), `int`) != -1 && gstr.Pos(column[`Type`].String(), `point`) == -1 { //int等类型
			if field == `pid` { //pid
				fieldName = `父级`
			} else if garray.NewStrArrayFrom([]string{`sort`, `weight`}).Contains(fieldSuffix) { //sort,weight等后缀
				viewI18nTip += `
        ` + field + `: '` + tip + `',`
			} else if garray.NewStrArrayFrom([]string{`id`}).Contains(fieldSuffix) { //id后缀
				if tpl.RelTableMap[field].TableRaw != `` && !tpl.RelTableMap[field].IsRedundRelNameField {
					fieldName = tpl.RelTableMap[field].RelTableFieldName
				}
			}
		} else if gstr.Pos(column[`Type`].String(), `json`) != -1 { //json类型
			viewI18nTip += `
        ` + field + `: '` + tip + `',`
		}

		viewI18nName += `
        ` + field + `: '` + fieldName + `',`
	}
	tplView := `export default {
    name: {` + viewI18nName + `
    },
    status: {` + viewI18nStatus + `
    },
    tip: {` + viewI18nTip + `
    },
}
`

	gfile.PutContents(saveFile, tplView)
}

// 前端路由生成
func (myGenThis *myGenHandler) genViewRouter() {
	option := myGenThis.option
	tpl := myGenThis.tpl

	saveFile := gfile.SelfDir() + `/../view/` + option.SceneCode + `/src/router/index.ts`

	tplViewRouter := gfile.GetContents(saveFile)

	path := `/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab
	replaceStr := `{
                path: '` + path + `',
                component: async () => {
                    const component = await import('@/views` + path + `/Index.vue')
                    component.default.name = '` + path + `'
                    return component
                },
                meta: { isAuth: true, keepAlive: true, componentName: '` + path + `' }
            },`

	if gstr.Pos(tplViewRouter, `'`+path+`'`) == -1 { //路由不存在时新增
		tplViewRouter = gstr.Replace(tplViewRouter, `/*--------前端路由自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/`, replaceStr+`
            /*--------前端路由自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/`, 1)
	} else { //路由已存在则替换
		tplViewRouter, _ = gregex.ReplaceString(`\{
                path: '`+path+`',[\s\S]*'`+path+`' \}
            \},`, replaceStr, tplViewRouter)
	}
	gfile.PutContents(saveFile, tplViewRouter)

	myGenThis.genMenu(myGenThis.sceneId, path, option.CommonName, tpl.TableCaseCamel) // 数据库权限菜单处理
}

// 自动生成操作权限
func (myGenThis *myGenHandler) genAction(sceneId uint, actionCode string, actionName string) {
	ctx := myGenThis.ctx

	actionName = gstr.Replace(actionName, `/`, `-`)

	idVar, _ := daoAuth.Action.CtxDaoModel(ctx).Filter(daoAuth.Action.Columns().ActionCode, actionCode).Value(daoAuth.Action.PrimaryKey())
	id := idVar.Int64()
	if id == 0 {
		id, _ = daoAuth.Action.CtxDaoModel(ctx).HookInsert(map[string]interface{}{
			daoAuth.Action.Columns().ActionCode: actionCode,
			daoAuth.Action.Columns().ActionName: actionName,
		}).InsertAndGetId()
	} else {
		daoAuth.Action.CtxDaoModel(ctx).Filter(daoAuth.Action.PrimaryKey(), id).HookUpdate(g.Map{daoAuth.Action.Columns().ActionName: actionName}).Update()
	}
	daoAuth.ActionRelToScene.CtxDaoModel(ctx).Data(map[string]interface{}{
		daoAuth.ActionRelToScene.Columns().ActionId: id,
		daoAuth.ActionRelToScene.Columns().SceneId:  sceneId,
	}).Save()
}

// 自动生成菜单
func (myGenThis *myGenHandler) genMenu(sceneId uint, menuUrl string, menuName string, menuNameOfEn string) {
	ctx := myGenThis.ctx

	menuNameArr := gstr.Split(menuName, `/`)

	var pid int64 = 0
	for _, v := range menuNameArr[:len(menuNameArr)-1] {
		pidVar, _ := daoAuth.Menu.CtxDaoModel(ctx).Filters(g.Map{
			daoAuth.Menu.Columns().SceneId:  sceneId,
			daoAuth.Menu.Columns().MenuName: v,
		}).Value(daoAuth.Menu.PrimaryKey())
		if pidVar.Uint() == 0 {
			pid, _ = daoAuth.Menu.CtxDaoModel(ctx).HookInsert(g.Map{
				daoAuth.Menu.Columns().SceneId:   sceneId,
				daoAuth.Menu.Columns().Pid:       pid,
				daoAuth.Menu.Columns().MenuName:  v,
				daoAuth.Menu.Columns().MenuIcon:  `autoicon-ep-link`,
				daoAuth.Menu.Columns().MenuUrl:   ``,
				daoAuth.Menu.Columns().ExtraData: `{"i18n": {"title": {"en": "", "zh-cn": "` + v + `"}}}`,
			}).InsertAndGetId()
		} else {
			pid = pidVar.Int64()
		}
	}

	menuName = menuNameArr[len(menuNameArr)-1]
	idVar, _ := daoAuth.Menu.CtxDaoModel(ctx).Filters(g.Map{
		daoAuth.Menu.Columns().SceneId: sceneId,
		daoAuth.Menu.Columns().MenuUrl: menuUrl,
	}).Value(daoAuth.Menu.PrimaryKey())
	id := idVar.Uint()
	if id == 0 {
		daoAuth.Menu.CtxDaoModel(ctx).HookInsert(g.Map{
			daoAuth.Menu.Columns().SceneId:   sceneId,
			daoAuth.Menu.Columns().Pid:       pid,
			daoAuth.Menu.Columns().MenuName:  menuName,
			daoAuth.Menu.Columns().MenuIcon:  `autoicon-ep-link`,
			daoAuth.Menu.Columns().MenuUrl:   menuUrl,
			daoAuth.Menu.Columns().ExtraData: `{"i18n": {"title": {"en": "` + menuNameOfEn + `", "zh-cn": "` + menuName + `"}}}`,
		}).Insert()
	} else {
		daoAuth.Menu.CtxDaoModel(ctx).Filter(daoAuth.Menu.PrimaryKey(), id).
			SetIdArr().
			HookUpdate(g.Map{
				daoAuth.Menu.Columns().MenuName:  menuName,
				daoAuth.Menu.Columns().Pid:       pid,
				daoAuth.Menu.Columns().ExtraData: `{"i18n": {"title": {"en": "` + menuNameOfEn + `", "zh-cn": "` + menuName + `"}}}`,
			}).Update()
	}
}

// 获取HandlePasswordMap的Key（以Password为主）
func (myGenThis *myGenHandler) command(title string, isOut bool, dir string, name string, arg ...string) {
	command := exec.Command(name, arg...)
	if dir != `` {
		command.Dir = dir
	}
	fmt.Println(title + ` 开始`)
	fmt.Println(`执行命令：` + command.String())
	stdout, _ := command.StdoutPipe()
	command.Start()
	if isOut {
		buf := make([]byte, 1024)
		for {
			n, err := stdout.Read(buf)
			if err != nil {
				break
			}
			if n > 0 {
				fmt.Print(string(buf[:n]))
			}
		}
	} else {
		fmt.Println(`请稍等，命令正在执行中...`)
	}
	command.Wait()
	fmt.Println(title + ` 结束`)
}

// status字段注释解析
func (myGenThis *myGenHandler) genStatusList(comment string, isStr bool) (statusList [][2]string) {
	var tmp [][]string
	if isStr {
		tmp, _ = gregex.MatchAllString(`([A-Za-z0-9]+)[-=:：]?([^\s,，;；)）]+)`, comment)
	} else {
		// tmp, _ = gregex.MatchAllString(`(-?\d+)[-=:：]?([^\d\s,，;；)）]+)`, comment)
		tmp, _ = gregex.MatchAllString(`(-?\d+)[-=:：]?([^\s,，;；)）]+)`, comment)
	}

	if len(tmp) == 0 {
		statusList = [][2]string{{`0`, `请设置表字段注释后，再生成代码`}}
		return
	}
	statusList = make([][2]string, len(tmp))
	for k, v := range tmp {
		statusList[k] = [2]string{v[1], v[2]}
	}
	return
}

// 获取HandlePasswordMap的Key（以Password为主）
func (myGenThis *myGenHandler) genHandlePasswordMapKey(passwordOrsalt string) (handlePasswordMapKey string) {
	passwordOrsalt = gstr.Replace(gstr.CaseCamel(passwordOrsalt), `Salt`, `Password`, 1) //替换salt
	passwordOrsalt = gstr.Replace(passwordOrsalt, `Passwd`, `Password`, 1)               //替换passwd
	handlePasswordMapKey = gstr.CaseCamelLower(passwordOrsalt)                           //默认：小驼峰
	if gstr.CaseCamelLower(passwordOrsalt) != passwordOrsalt {                           //判断字段是不是蛇形
		handlePasswordMapKey = gstr.CaseSnake(handlePasswordMapKey)
	}
	return
}

// 获取id后缀字段关联的表信息
func (myGenThis *myGenHandler) genRelTable(tpl myGenTpl, field string, fieldName string) relTableItem {
	fieldCaseSnake := gstr.CaseSnake(field)
	fieldCaseSnakeOfRemove := gstr.Split(fieldCaseSnake, `_of_`)[0]
	fieldCaseCamelOfRemove := gstr.CaseCamel(fieldCaseSnakeOfRemove)

	relTableCaseCamel := gstr.SubStr(fieldCaseCamelOfRemove, 0, -2)
	relTableItem := relTableItem{
		RelTableCaseSnake:      gstr.CaseSnake(relTableCaseCamel),
		RelTableCaseCamel:      relTableCaseCamel,
		RelTableCaseCamelLower: gstr.CaseCamelLower(relTableCaseCamel),
		RelTableFieldName:      fieldName,
	}
	fieldCaseSnakeArr := gstr.Split(fieldCaseSnake, `_of_`)
	if len(fieldCaseSnakeArr) > 1 {
		relTableItem.RelSuffixCaseSnake = `_of_` + gstr.Join(fieldCaseSnakeArr[1:], `_of_`)
		relTableItem.RelSuffixCaseCamel = gstr.CaseCamel(relTableItem.RelSuffixCaseSnake)
		relTableItem.RelSuffix = relTableItem.RelSuffixCaseCamel //默认：小驼峰
	}
	relTableItem.RelTableField = relTableItem.RelTableCaseCamelLower + `Name` //默认：小驼峰
	if gstr.CaseCamelLower(field) != field {                                  //判断字段是不是蛇形
		relTableItem.RelTableField = relTableItem.RelTableCaseSnake + `_name`
		if len(fieldCaseSnakeArr) > 1 {
			relTableItem.RelSuffix = relTableItem.RelSuffixCaseSnake
		}
	}
	if gstr.ToUpper(gstr.SubStr(relTableItem.RelTableFieldName, -2)) == `ID` {
		relTableItem.RelTableFieldName = gstr.SubStr(relTableItem.RelTableFieldName, 0, -2)
	}

	/*--------确定关联表 开始--------*/
	tableListSame := []string{} //关联表在同模块目录下，但表后缀一致
	tableSame := ``             //表名完全一致的表
	tableList := []string{}     //表后缀一致的表列表
	for _, v := range myGenThis.tableArr {
		if v == tpl.TableRaw { //自身跳过
			continue
		}
		if v == tpl.RemovePrefix+relTableItem.RelTableCaseSnake { //关联表在同模块目录下，且表名一致
			tableIndexList, _ := myGenThis.db.GetAll(myGenThis.ctx, `SHOW Index FROM `+v+` WHERE Key_name = 'PRIMARY'`)
			primaryKey := tableIndexList[0][`Column_name`].String()
			if len(tableIndexList) == 1 && (primaryKey == `id` || primaryKey == field) {
				relTableItem.TableRaw = v
				relTableItem.IsSameDir = true
				break
			}
		} else if gstr.Pos(v, `_`+tpl.RemovePrefix) == 0 && len(v) == gstr.PosR(v, `_`+relTableItem.RelTableCaseSnake)+len(`_`+relTableItem.RelTableCaseSnake) { //关联表在同模块目录下，但表后缀一致
			tableIndexList, _ := myGenThis.db.GetAll(myGenThis.ctx, `SHOW Index FROM `+v+` WHERE Key_name = 'PRIMARY'`)
			primaryKey := tableIndexList[0][`Column_name`].String()
			if len(tableIndexList) == 1 && (primaryKey == `id` || primaryKey == field) {
				tableListSame = append(tableListSame, v)
			}
		} else if gstr.Replace(v, myGenThis.option.RemovePrefixCommon, ``, 1) == relTableItem.RelTableCaseSnake { //表名完全一致
			tableIndexList, _ := myGenThis.db.GetAll(myGenThis.ctx, `SHOW Index FROM `+v+` WHERE Key_name = 'PRIMARY'`)
			primaryKey := tableIndexList[0][`Column_name`].String()
			if len(tableIndexList) == 1 && (primaryKey == `id` || primaryKey == field) {
				tableSame = v
			}
		} else if len(v) == gstr.PosR(v, `_`+relTableItem.RelTableCaseSnake)+len(`_`+relTableItem.RelTableCaseSnake) { //表后缀一致
			tableIndexList, _ := myGenThis.db.GetAll(myGenThis.ctx, `SHOW Index FROM `+v+` WHERE Key_name = 'PRIMARY'`)
			primaryKey := tableIndexList[0][`Column_name`].String()
			if len(tableIndexList) == 1 && (primaryKey == `id` || primaryKey == field) {
				tableList = append(tableList, v)
			}
		}
	}
	if relTableItem.TableRaw == `` {
		if len(tableListSame) > 0 {
			if len(tableListSame) == 1 {
				relTableItem.TableRaw = tableListSame[0]
				relTableItem.IsSameDir = true
			} else {
				// TODO
				count := 0 //关联表在同模块目录下，但表后缀一致
				tableSameDir := ``
				for _, v := range tableList {
					if gstr.Count(v, `_`) == gstr.Count(tpl.TableRaw, `_`) {
						count++
						tableSameDir = v
					}
				}
				if count == 1 { //当只存在一个表后缀一致的表时，直接使用该表
					relTableItem.TableRaw = tableSameDir
					relTableItem.IsSameDir = true
				}
			}
		} else if tableSame != `` {
			relTableItem.TableRaw = tableSame
		} else {
			if len(tableList) == 1 {
				relTableItem.TableRaw = tableList[0]
			} else {
				// TODO
				count := 0 //与当前模块同层的其它模块存在多少表后缀一致的表
				tableSameDir := ``
				for _, v := range tableList {
					if gstr.Count(v, `_`) == gstr.Count(tpl.TableRaw, `_`) {
						count++
						tableSameDir = v
					}
				}
				if count == 1 { //当只存在一个表后缀一致的表时，直接使用该表
					relTableItem.TableRaw = tableSameDir
				}
			}
		}
	}
	/*--------确定关联表 结束--------*/

	if relTableItem.TableRaw != `` {
		if relTableItem.IsSameDir {
			relTableItem.RemovePrefix = tpl.RemovePrefix
			relTableItem.RelDaoDir = tpl.ModuleDirCaseKebab
		} else {
			relTableItem.RemovePrefix = gstr.TrimRightStr(relTableItem.TableRaw, relTableItem.RelTableCaseSnake, 1)
			relDaoDir := gstr.Trim(relTableItem.RemovePrefix, `_`)
			relDaoDir = gstr.TrimLeftStr(relDaoDir, myGenThis.option.RemovePrefixCommon, 1)
			if relDaoDir == `` {
				relDaoDir = relTableItem.RelTableCaseSnake
			}
			if myGenThis.option.DbGroup != `default` {
				relDaoDir = gstr.CaseKebab(myGenThis.option.DbGroup) + `/` + gstr.CaseKebab(relDaoDir)
			}
			relTableItem.RelDaoDir = relDaoDir
		}
		// 判断dao文件是否存在，不存在则生成
		if !gfile.IsFile(gfile.SelfDir() + `/internal/dao/` + relTableItem.RelDaoDir + `/` + relTableItem.RelTableCaseSnake + `.go`) {
			myGenThis.command(`关联表（`+relTableItem.TableRaw+`）dao生成`, true, ``,
				`gf`, `gen`, `dao`,
				`--link`, myGenThis.dbLink,
				`--group`, myGenThis.option.DbGroup,
				`--removePrefix`, relTableItem.RemovePrefix,
				`--daoPath`, `dao/`+relTableItem.RelDaoDir,
				`--doPath`, `model/entity/`+relTableItem.RelDaoDir,
				`--entityPath`, `model/entity/`+relTableItem.RelDaoDir,
				`--tables`, relTableItem.TableRaw,
				`--tplDaoIndexPath`, `resource/gen/gen_dao_template_dao.txt`,
				`--tplDaoInternalPath`, `resource/gen/gen_dao_template_dao_internal.txt`,
				`--overwriteDao`, `false`)
		}
		//判断关联表是否存在pid字段
		relTableItem.TableColumnList, _ = myGenThis.db.GetAll(myGenThis.ctx, `SHOW FULL COLUMNS FROM `+relTableItem.TableRaw)
		for _, v := range relTableItem.TableColumnList {
			if v[`Field`].String() == `pid` {
				relTableItem.RelTableIsExistPidField = true
				break
			}
		}
	}

	relDaoDirArr := gstr.Split(relTableItem.RelDaoDir, `/`)
	relDaoDirCaseCamelArr := []string{}
	relDaoDirCaseCamelLowerArr := []string{}
	for _, v := range relDaoDirArr {
		relDaoDirCaseCamelArr = append(relDaoDirCaseCamelArr, gstr.CaseCamel(v))
		relDaoDirCaseCamelLowerArr = append(relDaoDirCaseCamelLowerArr, gstr.CaseCamelLower(v))
	}
	relTableItem.RelDaoDirCaseCamel = gstr.Join(relDaoDirCaseCamelArr, ``)
	relTableItem.RelDaoDirCaseCamelLower = gstr.Join(relDaoDirCaseCamelLowerArr, `/`)
	return relTableItem
}
