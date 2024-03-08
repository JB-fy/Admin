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
后台常用生成示例：./main myGen -sceneCode=platform -dbGroup=default -dbTable=auth_test -removePrefixCommon= -removePrefixAlone=auth_ -commonName=权限管理/测试 -isList=1 -isCount=1 -isInfo=1 -isCreate=1 -isUpdate=1 -isDelete=1 -isApi=1 -isAuthAction=1 -isView=1 -isCover=0
APP常用生成示例：./main myGen -sceneCode=app -dbGroup=xxxx -dbTable=user -removePrefixCommon= -removePrefixAlone= -commonName=用户 -isList=1 -isCount=0 -isInfo=1 -isCreate=0 -isUpdate=0 -isDelete=0 -isApi=1 -isAuthAction=0 -isView=0 -isCover=0

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

	尽量根据表名设置以下两个字段（作用1：常用于前端部分组件，如my-select或my-cascader等组件；作用2：用于关联表查询）
		xxId主键字段。示例：good表命名goodId, good_category表命名categoryId
		xxName或xxTitle字段。示例：good表命名goodName, article表命名articleTitle
			如果不存在xxName或xxTitle字段，按以下优先级默认一个
				表名去掉前缀 + Name > 主键去掉ID + Name > Name >
				表名去掉前缀 + Title > 主键去掉ID + Title > Title >
				表名去掉前缀 + Phone > 主键去掉ID + Phone > Phone >
				表名去掉前缀 + Email > 主键去掉ID + Email > Email >
				表名去掉前缀 + Account > 主键去掉ID + Account > Account >
				表名去掉前缀 + Nickname > 主键去掉ID + Nickname > Nickname

	字段都必须有注释。以下符号[\n\r.。:：(（]之前的部分或整个注释，将作为字段名称使用

	字段按以下规则命名时，会做特殊处理，其它情况根据字段类型做默认处理
		固定命名：
			父级		命名：pid；	类型：int等类型；
			层级		命名：level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
			层级路径	命名：idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
			排序		命名：sort，且pid,level,idPath|id_path,sort同时存在时（才）有效；	类型：int等类型；

		常用命名(字段含[_of_]时，会忽略[_of_]及其之后的部分)：
			密码		命名：password,passwd后缀；		类型：char(32)；
			加密盐 		命名：salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
			名称		命名：name,title后缀；			类型：varchar；
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
	myGenHandlerObj.tpl = myGenHandlerObj.createTpl(myGenHandlerObj.option.DbTable, myGenHandlerObj.option.RemovePrefixCommon, myGenHandlerObj.option.RemovePrefixAlone)

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
	CommonName         string `json:"commonName"`         //公共名称，将同时在swagger文档Tag标签，权限菜单和权限操作中使用。示例：用户，权限管理/测试
	IsList             bool   `json:"isList" `            //是否生成列表接口(0和no为false，1和yes为true)
	IsCount            bool   `json:"isCount" `           //列表接口是否返回总数
	IsInfo             bool   `json:"isInfo" `            //是否生成详情接口
	IsCreate           bool   `json:"isCreate"`           //是否生成创建接口
	IsUpdate           bool   `json:"isUpdate"`           //是否生成更新接口
	IsDelete           bool   `json:"isDelete"`           //是否生成删除接口
	IsApi              bool   `json:"isApi"`              //是否生成后端接口文件
	IsAuthAction       bool   `json:"isAuthAction"`       //是否判断操作权限，如是，则同时会生成操作权限
	IsView             bool   `json:"isView"`             //是否生成前端视图文件
	IsCover            bool   `json:"isCover"`            //是否覆盖原文件(设置为true时，建议与git一起使用，防止代码覆盖风险)
}

type myGenTpl struct {
	RemovePrefixCommon        string       //要删除的共有前缀
	RemovePrefixAlone         string       //要删除的独有前缀
	RemovePrefix              string       //要删除的前缀
	TableRaw                  string       //表名（原始，包含前缀）
	TableCaseSnake            string       //表名（蛇形，已去除前缀）
	TableCaseCamel            string       //表名（大驼峰，已去除前缀）
	TableCaseKebab            string       //表名（横线，已去除前缀）
	FieldListRaw              gdb.Result   //字段列表（原始）。SHOW FULL COLUMNS FROM xxTable的查询数据
	FieldList                 []myGenField //字段列表
	ModuleDirCaseCamel        string       //模块目录（大驼峰，/会被去除）
	ModuleDirCaseKebab        string       //模块目录（横线，/会被保留）
	ModuleDirCaseKebabReplace string       //模块目录（横线，/被替换成.）
	LogicStructName           string       //logic层结构体名称，也是权限操作前缀（大驼峰，由ModuleDirCaseCamel+TableCaseCamel组成。命名原因：gf gen service只支持logic单层目录，可能导致service层重名）
	FieldPrimary              string       //主键字段
	Handle                    struct {     //该属性记录需做特殊处理字段
		/*
			label列表。sql查询可设为别名label的字段（常用于前端my-select或my-cascader等组件，或用于关联表查询）。按以下优先级存入：
				表名去掉前缀 + Name > 主键去掉ID + Name > Name >
				表名去掉前缀 + Title > 主键去掉ID + Title > Title >
				表名去掉前缀 + Phone > 主键去掉ID + Phone > Phone >
				表名去掉前缀 + Email > 主键去掉ID + Email > Email >
				表名去掉前缀 + Account > 主键去掉ID + Account > Account >
				表名去掉前缀 + Nickname > 主键去掉ID + Nickname > Nickname
		*/
		LabelList   []string
		PasswordMap map[string]handlePassword //password|passwd,salt同时存在时，需特殊处理
		Pid         struct {                  //pid,level,idPath|id_path同时存在时，需特殊处理
			IsCoexist bool   //是否同时存在pid,level,idPath|id_path
			Pid       string //父级字段
			Level     string //层级字段
			IdPath    string //层级路径字段
			Sort      string //排序字段
		}
		RelIdMap map[string]handleRelId //TODO id后缀字段关联表信息
	}
}

type myGenFieldType = uint
type myGenFieldTypeName = string

const (
	//用于结构体中，需从1开始，否则结构体会默认0，即Int
	TypeInt       myGenFieldType = iota + 1 // `int等类型`
	TypeIntU                                // `int等类型（unsigned）`
	TypeFloat                               // `float等类型`
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
	TypeNamePid            myGenFieldTypeName = `命名：pid；	类型：int等类型；`
	TypeNameLevel          myGenFieldTypeName = `命名：level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；`
	TypeNameIdPath         myGenFieldTypeName = `命名：idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；`
	TypeNameSort           myGenFieldTypeName = `命名：sort，且pid,level,idPath|id_path,sort同时存在时（才）有效；	类型：int等类型；`
	TypeNamePasswordSuffix myGenFieldTypeName = `命名：password,passwd后缀；		类型：char(32)；`
	TypeNameSaltSuffix     myGenFieldTypeName = `命名：salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；`
	TypeNameNameSuffix     myGenFieldTypeName = `命名：name,title后缀；	类型：varchar；`
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

type handleRelId struct {
	tpl          myGenTpl
	FieldName    string //字段名称
	IsRedundName bool   //是否冗余过关联表名称字段
	Suffix       string //关联表字段后缀（原始，大驼峰或蛇形）。字段含[_of_]时，_of_及之后的部分。示例：userIdOfSend对应OfSend；user_id_of_send对应_of_send
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
func (myGenThis *myGenHandler) createTpl(table, removePrefixCommon string, removePrefixAlone string) (tpl myGenTpl) {
	ctx := myGenThis.ctx
	db := myGenThis.db
	option := myGenThis.option

	tpl = myGenTpl{
		RemovePrefixCommon: removePrefixCommon,
		RemovePrefixAlone:  removePrefixAlone,
		RemovePrefix:       removePrefixCommon + removePrefixAlone,
		TableRaw:           table,
	}
	tpl.Handle.PasswordMap = map[string]handlePassword{}
	tpl.Handle.RelIdMap = map[string]handleRelId{}
	tpl.FieldListRaw, _ = db.GetAll(ctx, `SHOW FULL COLUMNS FROM `+table)
	tpl.TableCaseSnake = gstr.CaseSnake(gstr.Replace(tpl.TableRaw, tpl.RemovePrefix, ``, 1))
	tpl.TableCaseCamel = gstr.CaseCamel(tpl.TableCaseSnake)
	tpl.TableCaseKebab = gstr.CaseKebab(tpl.TableCaseSnake)

	logicStructName := gstr.TrimLeftStr(table, tpl.RemovePrefixCommon, 1)
	moduleDirCaseCamel := gstr.CaseCamel(logicStructName)
	moduleDirCaseKebab := gstr.CaseKebab(logicStructName)
	if tpl.RemovePrefixAlone != `` {
		moduleDirCaseCamel = gstr.CaseCamel(tpl.RemovePrefixAlone)
		moduleDirCaseKebab = gstr.CaseKebab(gstr.Trim(tpl.RemovePrefixAlone, `_`))
	}
	if option.DbGroup != `default` {
		logicStructName = option.DbGroup + `_` + logicStructName
		moduleDirCaseCamel = gstr.CaseCamel(option.DbGroup) + moduleDirCaseCamel
		moduleDirCaseKebab = gstr.CaseKebab(option.DbGroup) + `/` + moduleDirCaseKebab
	}
	tpl.LogicStructName = gstr.CaseCamel(logicStructName)
	tpl.ModuleDirCaseKebab = moduleDirCaseKebab
	tpl.ModuleDirCaseKebabReplace = gstr.Replace(moduleDirCaseKebab, `/`, `.`)
	tpl.ModuleDirCaseCamel = moduleDirCaseCamel

	fieldList := make([]myGenField, len(tpl.FieldListRaw))
	for k, v := range tpl.FieldListRaw {
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
			if gstr.Pos(fieldTmp.FieldTypeRaw, `datetime`) != -1 {
				fieldTmp.FieldType = TypeDatetime
			} else if gstr.Pos(fieldTmp.FieldTypeRaw, `date`) != -1 {
				fieldTmp.FieldType = TypeDate
			}
		}
		/*--------确定字段数据类型 结束--------*/

		/*--------确定字段命名类型（部分命名类型需做二次确定） 开始--------*/
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

				tpl.FieldPrimary = fieldTmp.FieldRaw
			}
		} else if garray.NewFrom([]interface{}{TypeVarchar, TypeText}).Contains(fieldTmp.FieldType) && fieldTmp.FieldCaseCamel == `IdPath` { //idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效
			fieldTmp.FieldTypeName = TypeNameIdPath

			tpl.Handle.Pid.IdPath = fieldTmp.FieldRaw
		} else if garray.NewFrom([]interface{}{TypeInt, TypeIntU, TypeVarchar, TypeChar}).Contains(fieldTmp.FieldType) && garray.NewStrArrayFrom([]string{`status`, `type`, `method`, `pos`, `position`, `gender`}).Contains(fieldSuffix) { //status,type,method,pos,position,gender等后缀
			fieldTmp.FieldTypeName = TypeNameStatusSuffix

			isStr := false
			if garray.NewFrom([]interface{}{TypeVarchar, TypeChar}).Contains(fieldTmp.FieldType) {
				isStr = true
			}
			fieldTmp.StatusList = myGenThis.getStatusList(fieldTmp.FieldDesc, isStr)
		} else if garray.NewFrom([]interface{}{TypeVarchar, TypeText, TypeJson}).Contains(fieldTmp.FieldType) && (garray.NewStrArrayFrom([]string{`icon`, `cover`, `avatar`, `img`, `image`}).Contains(fieldSuffix) || gstr.SubStr(fieldTmp.FieldCaseCamelRemove, -7) == `ImgList` || gstr.SubStr(fieldTmp.FieldCaseCamelRemove, -6) == `ImgArr` || gstr.SubStr(fieldTmp.FieldCaseCamelRemove, -9) == `ImageList` || gstr.SubStr(fieldTmp.FieldCaseCamelRemove, -8) == `ImageArr`) { //icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀
			fieldTmp.FieldTypeName = TypeNameImageSuffix
		} else if garray.NewFrom([]interface{}{TypeVarchar, TypeText, TypeJson}).Contains(fieldTmp.FieldType) && (garray.NewStrArrayFrom([]string{`video`}).Contains(fieldSuffix) || gstr.SubStr(fieldTmp.FieldCaseCamelRemove, -9) == `VideoList` || gstr.SubStr(fieldTmp.FieldCaseCamelRemove, -8) == `VideoArr`) { //video,video_list,videoList,video_arr,videoArr等后缀
			fieldTmp.FieldTypeName = TypeNameVideoSuffix
		} else if garray.NewFrom([]interface{}{TypeText, TypeJson}).Contains(fieldTmp.FieldType) && garray.NewStrArrayFrom([]string{`list`, `arr`}).Contains(fieldSuffix) { //list,arr等后缀
			fieldTmp.FieldTypeName = TypeNameArrSuffix
		} else if garray.NewFrom([]interface{}{TypeVarchar, TypeText}).Contains(fieldTmp.FieldType) && garray.NewStrArrayFrom([]string{`remark`, `desc`, `msg`, `message`, `intro`, `content`}).Contains(fieldSuffix) { //remark,desc,msg,message,intro,content后缀
			fieldTmp.FieldTypeName = TypeNameRemarkSuffix
		} else if fieldTmp.FieldType == TypeVarchar { //varchar类型
			if garray.NewStrArrayFrom([]string{`name`, `title`}).Contains(fieldSuffix) { //name,title后缀
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

				passwordMapKey := myGenThis.getHandlePasswordMapKey(fieldTmp.FieldRaw)
				handlePasswordObj, ok := tpl.Handle.PasswordMap[passwordMapKey]
				if ok {
					handlePasswordObj.PasswordField = fieldTmp.FieldRaw
					handlePasswordObj.PasswordLength = fieldTmp.FieldLimitStr
				} else {
					handlePasswordObj = handlePassword{
						PasswordField:  fieldTmp.FieldRaw,
						PasswordLength: fieldTmp.FieldLimitStr,
					}
				}
				tpl.Handle.PasswordMap[passwordMapKey] = handlePasswordObj
			} else if garray.NewStrArrayFrom([]string{`salt`}).Contains(fieldSuffix) { //salt后缀，且对应的password,passwd后缀存在时（才）有效。该命名类型需做二次确定
				fieldTmp.FieldTypeName = TypeNameSaltSuffix

				passwordMapKey := myGenThis.getHandlePasswordMapKey(fieldTmp.FieldRaw)
				handlePasswordObj, ok := tpl.Handle.PasswordMap[passwordMapKey]
				if ok {
					handlePasswordObj.SaltField = fieldTmp.FieldRaw
					handlePasswordObj.SaltLength = fieldTmp.FieldLimitStr
				} else {
					handlePasswordObj = handlePassword{
						SaltField:  fieldTmp.FieldRaw,
						SaltLength: fieldTmp.FieldLimitStr,
					}
				}
				tpl.Handle.PasswordMap[passwordMapKey] = handlePasswordObj
			}
		} else if garray.NewFrom([]interface{}{TypeInt, TypeIntU}).Contains(fieldTmp.FieldType) { //int等类型
			if fieldTmp.FieldRaw == `pid` { //pid
				fieldTmp.FieldTypeName = TypeNamePid

				tpl.Handle.Pid.Pid = fieldTmp.FieldRaw
			} else if fieldTmp.FieldRaw == `level` { //level，且pid,level,idPath|id_path同时存在时（才）有效。该命名类型需做二次确定
				fieldTmp.FieldTypeName = TypeNameLevel

				tpl.Handle.Pid.Level = fieldTmp.FieldRaw
			} else if garray.NewStrArrayFrom([]string{`sort`, `weight`}).Contains(fieldSuffix) { //sort,weight等后缀。该命名类型需做二次确定
				fieldTmp.FieldTypeName = TypeNameSortSuffix
				if fieldTmp.FieldRaw == `sort` { //sort，且pid,level,idPath|id_path,sort同时存在时（才）有效。该命名类型需做二次确定
					fieldTmp.FieldTypeName = TypeNameSort

					tpl.Handle.Pid.Sort = fieldTmp.FieldRaw
				}
			} else if garray.NewStrArrayFrom([]string{`id`}).Contains(fieldSuffix) { //id后缀
				fieldTmp.FieldTypeName = TypeNameIdSuffix
				handleRelIdObj := handleRelId{
					tpl:       myGenThis.getRelIdTpl(tpl, fieldTmp.FieldRaw),
					FieldName: fieldTmp.FieldName,
				}
				if gstr.ToUpper(gstr.SubStr(handleRelIdObj.FieldName, -2)) == `ID` {
					handleRelIdObj.FieldName = gstr.SubStr(handleRelIdObj.FieldName, 0, -2)
				}
				if pos := gstr.Pos(fieldTmp.FieldCaseSnake, `_of_`); pos != -1 {
					handleRelIdObj.Suffix = gstr.SubStr(fieldTmp.FieldCaseSnake, pos)
					if fieldTmp.FieldRaw != fieldTmp.FieldCaseSnake {
						handleRelIdObj.Suffix = gstr.CaseCamel(handleRelIdObj.Suffix)
					}
				}
				tpl.Handle.RelIdMap[fieldTmp.FieldRaw] = handleRelIdObj
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
		/*--------确定字段命名类型（部分命名类型需做二次确定） 结束--------*/

		fieldList[k] = fieldTmp
	}

	/*--------需做特殊处理字段解析 开始--------*/
	/*
		label列表。sql查询可设为别名label的字段（常用于前端my-select或my-cascader等组件，或用于关联表查询）。按以下优先级存入：
			表名去掉前缀 + Name > 主键去掉ID + Name > Name >
			表名去掉前缀 + Title > 主键去掉ID + Title > Title >
			表名去掉前缀 + Phone > 主键去掉ID + Phone > Phone >
			表名去掉前缀 + Email > 主键去掉ID + Email > Email >
			表名去掉前缀 + Account > 主键去掉ID + Account > Account >
			表名去掉前缀 + Nickname > 主键去掉ID + Nickname > Nickname
	*/
	labelList := []string{}
	for _, v := range []string{`Name`, `Title`, `Phone`, `Email`, `Account`, `Nickname`} {
		labelTmp := tpl.TableCaseCamel + v
		labelList = append(labelList, labelTmp)
		labelTmp1 := gstr.SubStr(gstr.CaseCamel(tpl.FieldPrimary), 0, -2) + v
		if labelTmp1 != labelTmp && labelTmp1 != v {
			labelList = append(labelList, labelTmp1)
		}
		labelList = append(labelList, v)
	}
	tpl.Handle.LabelList = []string{}
	for _, v := range labelList {
		for _, item := range fieldList {
			if v == item.FieldCaseCamel {
				tpl.Handle.LabelList = append(tpl.Handle.LabelList, item.FieldRaw)
				break
			}
		}
	}

	for k, v := range tpl.Handle.RelIdMap {
		if len(v.tpl.Handle.LabelList) > 0 {
			for _, item := range fieldList {
				if item.FieldRaw == v.tpl.Handle.LabelList[0]+v.Suffix {
					v.IsRedundName = true
					tpl.Handle.RelIdMap[k] = v
					break
				}
			}
		}
	}

	for k, v := range tpl.Handle.PasswordMap {
		if v.PasswordField != `` && v.SaltField != `` {
			v.IsCoexist = true
			tpl.Handle.PasswordMap[k] = v
		}
	}

	if tpl.Handle.Pid.Pid != `` && tpl.Handle.Pid.Level != `` && tpl.Handle.Pid.IdPath != `` {
		tpl.Handle.Pid.IsCoexist = true
	}
	/*--------需做特殊处理字段解析 结束--------*/

	/*--------部分命名类型需要二次确认 开始--------*/
	for k, v := range fieldList {
		switch v.FieldTypeName {
		case TypeNameLevel, TypeNameIdPath: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；	// idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
			if !tpl.Handle.Pid.IsCoexist {
				fieldList[k].FieldTypeName = ``
			}
		case TypeNameSort: // sort，且pid,level,idPath|id_path,sort同时存在时（才）有效；	类型：int等类型；
			if !tpl.Handle.Pid.IsCoexist {
				fieldList[k].FieldTypeName = TypeNameSortSuffix
			}
		case TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
			passwordMapKey := myGenThis.getHandlePasswordMapKey(v.FieldRaw)
			if !tpl.Handle.PasswordMap[passwordMapKey].IsCoexist {
				fieldList[k].FieldTypeName = ``
			}
		}
	}
	/*--------部分命名类型需要二次确认 结束--------*/
	tpl.FieldList = fieldList
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
		`--removePrefix`, myGenThis.tpl.RemovePrefix,
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

	type dao struct {
		importDao string
		insert    struct {
			point       string
			parse       string
			parseBefore string
			hook        string
		}
		update struct {
			point      string
			parse      string
			hookBefore string
			hookAfter  string
		}
		field struct {
			point string
			parse string
			hook  string
		}
		filter struct {
			parse string
		}
		order struct {
			parse string
		}
		join struct {
			parse string
		}
	}
	daoObj := dao{}

	labelListLen := len(tpl.Handle.LabelList)
	if labelListLen > 0 {
		fieldParseStr := `
			case ` + "`label`" + `:
				m = m.Fields(daoModel.DbTable + ` + "`.`" + ` + daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.LabelList[0]) + ` + ` + "` AS `" + ` + v)`
		filterParseStr := `
			case ` + "`label`" + `:
				m = m.WhereLike(daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.LabelList[0]) + `, ` + "`%`" + `+gconv.String(v)+` + "`%`" + `)`
		if labelListLen > 1 {
			fieldParseStrTmp := "` + daoModel.DbTable + `.` + daoThis.Columns()." + gstr.CaseCamel(tpl.Handle.LabelList[labelListLen-1]) + " + `"
			parseFilterStr := "WhereOrLike(daoModel.DbTable+`.`+daoThis.Columns()." + gstr.CaseCamel(tpl.Handle.LabelList[labelListLen-1]) + ", `%`+gconv.String(v)+`%`)"
			for i := labelListLen - 2; i >= 0; i-- {
				fieldParseStrTmp = "IF(` + daoModel.DbTable + `.` + daoThis.Columns()." + gstr.CaseCamel(tpl.Handle.LabelList[i]) + " + `, ` + daoModel.DbTable + `.` + daoThis.Columns()." + gstr.CaseCamel(tpl.Handle.LabelList[i]) + " + `, " + fieldParseStrTmp + ")"
				if i == 0 {
					parseFilterStr = "WhereLike(daoModel.DbTable+`.`+daoThis.Columns()." + gstr.CaseCamel(tpl.Handle.LabelList[i]) + ", `%`+gconv.String(v)+`%`)." + parseFilterStr
				} else {
					parseFilterStr = "WhereOrLike(daoModel.DbTable+`.`+daoThis.Columns()." + gstr.CaseCamel(tpl.Handle.LabelList[i]) + ", `%`+gconv.String(v)+`%`)." + parseFilterStr
				}
			}
			fieldParseStr = `
			case ` + "`label`" + `:
				m = m.Fields(` + "`" + fieldParseStrTmp + " AS ` + v)"
			filterParseStr = `
			case ` + "`label`" + `:
				m = m.Where(m.Builder().` + parseFilterStr + `)`
		}
		if gstr.Pos(tplDao, fieldParseStr) == -1 {
			daoObj.field.parse += fieldParseStr
		}
		if gstr.Pos(tplDao, filterParseStr) == -1 {
			daoObj.filter.parse += filterParseStr
		}
	}

	for _, v := range tpl.FieldList {
		/*--------根据字段命名类型处理 开始--------*/
		switch v.FieldTypeName {
		case TypeNameDeleted: // 软删除字段
		case TypeNameUpdated: // 更新时间字段
		case TypeNameCreated: // 创建时间字段
			filterParseStr := `
			case ` + "`timeRangeStart`" + `:
				m = m.WhereGTE(daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + v.FieldCaseCamel + `, v)
			case ` + "`timeRangeEnd`" + `:
				m = m.WhereLTE(daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + v.FieldCaseCamel + `, v)`
			if gstr.Pos(tplDao, filterParseStr) == -1 {
				daoObj.filter.parse += filterParseStr
			}
		case TypeNamePri: // 主键
		case TypeNamePriAutoInc: // 主键（自增）
		case TypeNamePid: // pid；	类型：int等类型；
			if len(tpl.Handle.LabelList) > 0 {
				fieldParseStr := `
			case ` + "`p" + gstr.CaseCamel(tpl.Handle.LabelList[0]) + "`" + `:
				tableP := ` + "`p_`" + ` + daoModel.DbTable
				m = m.Fields(tableP + ` + "`.`" + ` + daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.LabelList[0]) + ` + ` + "` AS `" + ` + v)
				m = m.Handler(daoThis.ParseJoin(tableP, daoModel))`
				if gstr.Pos(tplDao, fieldParseStr) == -1 {
					daoObj.field.parse += fieldParseStr
				}
			}
			fieldParseStr := `
			case ` + "`tree`" + `:
				m = m.Fields(daoModel.DbTable + ` + "`.`" + ` + daoThis.PrimaryKey())
				m = m.Fields(daoModel.DbTable + ` + "`.`" + ` + daoThis.Columns().` + v.FieldCaseCamel + `)
				m = m.Handler(daoThis.ParseOrder([]string{` + "`tree`" + `}, daoModel))`
			if gstr.Pos(tplDao, fieldParseStr) == -1 {
				daoObj.field.parse += fieldParseStr
			}
			orderParseStr := `
			case ` + "`tree`" + `:
				m = m.OrderAsc(daoModel.DbTable + ` + "`.`" + ` + daoThis.Columns().` + v.FieldCaseCamel + `)`
			if tpl.Handle.Pid.Sort != `` {
				orderParseStr += `
				m = m.OrderAsc(daoModel.DbTable + ` + "`.`" + ` + daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Sort) + `)`
			}
			orderParseStr += `
				m = m.OrderAsc(daoModel.DbTable + ` + "`.`" + ` + daoThis.PrimaryKey())`
			if gstr.Pos(tplDao, orderParseStr) == -1 {
				daoObj.order.parse += orderParseStr
			}
			joinParseStr := `
		case ` + "`p_`" + ` + daoModel.DbTable:
			m = m.LeftJoin(daoModel.DbTable+` + "` AS `" + `+joinTable, joinTable+` + "`.`" + `+daoThis.PrimaryKey()+` + "` = `" + `+daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + v.FieldCaseCamel + `)`
			if gstr.Pos(tplDao, joinParseStr) == -1 {
				daoObj.join.parse += joinParseStr
			}

			if tpl.Handle.Pid.IsCoexist {
				insertParseBeforeStr := `
		if _, ok := insert[daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Pid) + `]; !ok {
			insert[daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Pid) + `] = 0
		}`
				if gstr.Pos(tplDao, insertParseBeforeStr) == -1 {
					daoObj.insert.parseBefore += insertParseBeforeStr
				}
				insertParseStr := `
			case daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Pid) + `:
				insertData[k] = v
				if gconv.Uint(v) > 0 {
					pInfo, _ := daoThis.CtxDaoModel(m.GetCtx()).Filter(daoThis.PrimaryKey(), v).One()
					daoModel.AfterInsert[` + "`pIdPath`" + `] = pInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.IdPath) + `].String()
					daoModel.AfterInsert[` + "`pLevel`" + `] = pInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Level) + `].Uint()
				} else {
					daoModel.AfterInsert[` + "`pIdPath`" + `] = ` + "`0`" + `
					daoModel.AfterInsert[` + "`pLevel`" + `] = 0
				}`
				if gstr.Pos(tplDao, insertParseStr) == -1 {
					daoObj.insert.parse += insertParseStr
				}
				insertHookStr := `

			updateSelfData := map[string]interface{}{}
			for k, v := range daoModel.AfterInsert {
				switch k {
				case ` + "`pIdPath`" + `:
					updateSelfData[daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.IdPath) + `] = gconv.String(v) + ` + "`-`" + ` + gconv.String(id)
				case ` + "`pLevel`" + `:
					updateSelfData[daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Level) + `] = gconv.Uint(v) + 1
				}
			}
			if len(updateSelfData) > 0 {
				daoModel.CloneNew().Filter(daoThis.PrimaryKey(), id).HookUpdate(updateSelfData).Update()
			}`
				if gstr.Pos(tplDao, insertHookStr) == -1 {
					daoObj.insert.hook += insertHookStr
				}
				updateParseStr := `
			case daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Pid) + `:
				updateData[daoModel.DbTable+` + "`.`" + `+k] = v
				pIdPath := ` + "`0`" + `
				var pLevel uint = 0
				if gconv.Uint(v) > 0 {
					pInfo, _ := daoThis.CtxDaoModel(m.GetCtx()).Filter(daoThis.PrimaryKey(), v).One()
					pIdPath = pInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.IdPath) + `].String()
					pLevel = pInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Level) + `].Uint()
				}
				updateData[daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.IdPath) + `] = gdb.Raw(` + "`CONCAT('`" + ` + pIdPath + ` + "`-', `" + ` + daoThis.PrimaryKey() + ` + "`)`" + `)
				updateData[daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Level) + `] = pLevel + 1
				//更新所有子孙级的idPath和level
				updateChildIdPathAndLevelList := []map[string]interface{}{}
				oldList, _ := daoThis.CtxDaoModel(m.GetCtx()).Filter(daoThis.PrimaryKey(), daoModel.IdArr).All()
				for _, oldInfo := range oldList {
					if gconv.Uint(v) != oldInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Pid) + `].Uint() {
						updateChildIdPathAndLevelList = append(updateChildIdPathAndLevelList, map[string]interface{}{
							` + "`pIdPathOfOld`" + `: oldInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.IdPath) + `],
							` + "`pIdPathOfNew`" + `: pIdPath + ` + "`-`" + ` + oldInfo[daoThis.PrimaryKey()].String(),
							` + "`pLevelOfOld`" + `:  oldInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Level) + `],
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
				updateData[daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.IdPath) + `] = gdb.Raw(` + "`REPLACE(`" + ` + daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.IdPath) + ` + ` + "`, '`" + ` + pIdPathOfOld + ` + "`', '`" + ` + pIdPathOfNew + ` + "`')`" + `)
			case ` + "`childLevel`" + `: //更新所有子孙级的level。参数：map[string]interface{}{` + "`pLevelOfOld`" + `: ` + "`父级Level（旧）`" + `, ` + "`pLevelOfNew`" + `: ` + "`父级Level（新）`" + `}
				val := gconv.Map(v)
				pLevelOfOld := gconv.Uint(val[` + "`pLevelOfOld`" + `])
				pLevelOfNew := gconv.Uint(val[` + "`pLevelOfNew`" + `])
				updateData[daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Level) + `] = gdb.Raw(daoModel.DbTable + ` + "`.`" + ` + daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Level) + ` + ` + "` + `" + ` + gconv.String(pLevelOfNew-pLevelOfOld))
				if pLevelOfNew < pLevelOfOld {
					updateData[daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Level) + `] = gdb.Raw(daoModel.DbTable + ` + "`.`" + ` + daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Level) + ` + ` + "` - `" + ` + gconv.String(pLevelOfOld-pLevelOfNew))
				}`
				if gstr.Pos(tplDao, updateParseStr) == -1 {
					daoObj.update.parse += updateParseStr
				}
				updateHookAfterStr := `

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
				if gstr.Pos(tplDao, updateHookAfterStr) == -1 {
					daoObj.update.hookAfter += updateHookAfterStr
				}
				filterParseStr := `
			case ` + "`pIdPathOfOld`" + `: //父级IdPath（旧）
				m = m.WhereLike(daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.IdPath) + `, gconv.String(v)+` + "`-%`" + `)`
				if gstr.Pos(tplDao, filterParseStr) == -1 {
					daoObj.filter.parse += filterParseStr
				}
			}
		case TypeNameLevel: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
			orderParseStr := `
			case daoThis.Columns().` + v.FieldCaseCamel + `:
				m = m.Order(daoModel.DbTable + ` + "`.`" + ` + v)
				m = m.OrderDesc(daoModel.DbTable + ` + "`.`" + ` + daoThis.PrimaryKey())` //追加主键倒序。mysql排序字段有重复值时，分页会导致同一条数据可能在不同页都出现
			if gstr.Pos(tplDao, orderParseStr) == -1 {
				daoObj.order.parse += orderParseStr
			}
		case TypeNameIdPath: // idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
			continue
		case TypeNamePasswordSuffix: // password,passwd后缀；		类型：char(32)；
			insertParseStr := `
			case daoThis.Columns().` + v.FieldCaseCamel + `:
				password := gconv.String(v)
				if len(password) != 32 {
					password = gmd5.MustEncrypt(password)
				}`
			updateParseStr := `
			case daoThis.Columns().` + v.FieldCaseCamel + `:
				password := gconv.String(v)
				if len(password) != 32 {
					password = gmd5.MustEncrypt(password)
				}`
			passwordMapKey := myGenThis.getHandlePasswordMapKey(v.FieldRaw)
			if tpl.Handle.PasswordMap[passwordMapKey].IsCoexist {
				insertParseStr += `
				salt := grand.S(` + tpl.Handle.PasswordMap[passwordMapKey].SaltLength + `)
				insertData[daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.PasswordMap[passwordMapKey].SaltField) + `] = salt
				password = gmd5.MustEncrypt(password + salt)`
				updateParseStr += `
				salt := grand.S(` + tpl.Handle.PasswordMap[passwordMapKey].SaltLength + `)
				updateData[daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.PasswordMap[passwordMapKey].SaltField) + `] = salt
				password = gmd5.MustEncrypt(password + salt)`
			}
			insertParseStr += `
				insertData[k] = password`
			updateParseStr += `
				updateData[daoModel.DbTable+` + "`.`" + `+k] = password`
			if gstr.Pos(tplDao, insertParseStr) == -1 {
				daoObj.insert.parse += insertParseStr
			}
			if gstr.Pos(tplDao, updateParseStr) == -1 {
				daoObj.update.parse += updateParseStr
			}
			continue
		case TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
			continue
		case TypeNameNameSuffix: // name,title后缀；	类型：varchar；
			filterParseStr := `
			case daoThis.Columns().` + v.FieldCaseCamel + `:
				m = m.WhereLike(daoModel.DbTable+` + "`.`" + `+k, ` + "`%`" + `+gconv.String(v)+` + "`%`" + `)`
			if gstr.Pos(tplDao, filterParseStr) == -1 {
				daoObj.filter.parse += filterParseStr
			}
		case TypeNameCodeSuffix: // code后缀；	类型：varchar；
		case TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
		case TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
		case TypeNameIpSuffix: // IP后缀；	类型：varchar；
		case TypeNameIdSuffix: // id后缀；	类型：int等类型；
			if tpl.Handle.RelIdMap[v.FieldRaw].tpl.TableRaw != `` {
				relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
				daoPath := relIdObj.tpl.TableCaseCamel
				if relIdObj.tpl.RemovePrefixAlone != tpl.RemovePrefixAlone {
					daoPath = `dao` + relIdObj.tpl.ModuleDirCaseCamel + `.` + relIdObj.tpl.TableCaseCamel
					importDaoStr := `
	dao` + relIdObj.tpl.ModuleDirCaseCamel + ` "api/internal/dao/` + relIdObj.tpl.ModuleDirCaseKebab + `"`
					if gstr.Pos(tplDao, importDaoStr) == -1 {
						daoObj.importDao += importDaoStr
					}
				}
				if !tpl.Handle.RelIdMap[v.FieldRaw].IsRedundName {
					fieldParseStr := `//因前端页面已用该字段名显示，故不存在时改成` + "`" + relIdObj.tpl.Handle.LabelList[0] + relIdObj.Suffix + "`" + `（控制器也要改）。同时下面Fields方法改成m = m.Fields(table` + relIdObj.tpl.TableCaseCamel + gstr.CaseCamel(relIdObj.Suffix) + ` + ` + "`.`" + ` + ` + daoPath + `.Columns().Xxxx + ` + "` AS `" + ` + v)`
					if gstr.Pos(tplDao, fieldParseStr) == -1 {
						if relIdObj.Suffix != `` {
							fieldParseStr = `
			case ` + daoPath + `.Columns().` + gstr.CaseCamel(relIdObj.tpl.Handle.LabelList[0]) + " + `" + relIdObj.Suffix + "`: " + fieldParseStr + `
				table` + relIdObj.tpl.TableCaseCamel + gstr.CaseCamel(relIdObj.Suffix) + ` := ` + daoPath + `.ParseDbTable(m.GetCtx()) + ` + "`" + gstr.CaseSnake(relIdObj.Suffix) + "`" + `
				m = m.Fields(table` + relIdObj.tpl.TableCaseCamel + gstr.CaseCamel(relIdObj.Suffix) + ` + ` + "`.`" + ` + ` + daoPath + `.Columns().` + gstr.CaseCamel(relIdObj.tpl.Handle.LabelList[0]) + ` + ` + "` AS `" + ` + v)
				m = m.Handler(daoThis.ParseJoin(table` + relIdObj.tpl.TableCaseCamel + gstr.CaseCamel(relIdObj.Suffix) + `, daoModel))`
						} else {
							fieldParseStr = `
			case ` + daoPath + `.Columns().` + gstr.CaseCamel(relIdObj.tpl.Handle.LabelList[0]) + `: ` + fieldParseStr + `
				table` + relIdObj.tpl.TableCaseCamel + ` := ` + daoPath + `.ParseDbTable(m.GetCtx())
				m = m.Fields(table` + relIdObj.tpl.TableCaseCamel + ` + ` + "`.`" + ` + v)
				m = m.Handler(daoThis.ParseJoin(table` + relIdObj.tpl.TableCaseCamel + `, daoModel))`
						}
						daoObj.field.parse += fieldParseStr
					}
				}
				joinParseStr := `
		case ` + daoPath + `.ParseDbTable(m.GetCtx()):
			m = m.LeftJoin(joinTable, joinTable+` + "`.`" + `+` + daoPath + `.PrimaryKey()+` + "` = `" + `+daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + v.FieldCaseCamel + `)`
				if relIdObj.Suffix != `` {
					joinParseStr = `
		case ` + daoPath + `.ParseDbTable(m.GetCtx()) + ` + "`" + gstr.CaseSnake(relIdObj.Suffix) + "`" + `:
			m = m.LeftJoin(` + daoPath + `.ParseDbTable(m.GetCtx())+` + "` AS `" + `+joinTable, joinTable+` + "`.`" + `+` + daoPath + `.PrimaryKey()+` + "` = `" + `+daoModel.DbTable+` + "`.`" + `+daoThis.Columns().` + v.FieldCaseCamel + `)`
				}
				if gstr.Pos(tplDao, joinParseStr) == -1 {
					daoObj.join.parse += joinParseStr
				}
			}
		case TypeNameSortSuffix, TypeNameSort: // sort,weight等后缀；	类型：int等类型； // sort，且pid,level,idPath|id_path,sort同时存在时（才）有效；	类型：int等类型；
			orderParseStr := `
			case daoThis.Columns().` + v.FieldCaseCamel + `:
				m = m.Order(daoModel.DbTable + ` + "`.`" + ` + v)
				m = m.OrderDesc(daoModel.DbTable + ` + "`.`" + ` + daoThis.PrimaryKey())` //追加主键倒序。mysql排序字段有重复值时，分页会导致同一条数据可能在不同页都出现
			if gstr.Pos(tplDao, orderParseStr) == -1 {
				daoObj.order.parse += orderParseStr
			}
		case TypeNameStatusSuffix: // status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
		case TypeNameIsPrefix: // is_前缀；		类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）
		case TypeNameStartPrefix: // start_前缀；	类型：timestamp或datetime或date；
			filterParseStr := `
			case daoThis.Columns().` + v.FieldCaseCamel + `:
				m = m.WhereLTE(daoModel.DbTable+` + "`.`" + `+k, v)`
			if !v.IsNull && gconv.String(v.Default) == `` {
				filterParseStr = `
			case daoThis.Columns().` + v.FieldCaseCamel + `:
				m = m.Where(m.Builder().WhereLTE(daoModel.DbTable+` + "`.`" + `+k, v).WhereOrNull(daoModel.DbTable + ` + "`.`" + ` + k))`
			}
			if gstr.Pos(tplDao, filterParseStr) == -1 {
				daoObj.filter.parse += filterParseStr
			}
		case TypeNameEndPrefix: // end_前缀；	类型：timestamp或datetime或date；
			filterParseStr := `
			case daoThis.Columns().` + v.FieldCaseCamel + `:
				m = m.WhereGTE(daoModel.DbTable+` + "`.`" + `+k, v)`
			if !v.IsNull && gconv.String(v.Default) == `` {
				filterParseStr = `
			case daoThis.Columns().` + v.FieldCaseCamel + `:
				m = m.Where(m.Builder().WhereGTE(daoModel.DbTable+` + "`.`" + `+k, v).WhereOrNull(daoModel.DbTable + ` + "`.`" + ` + k))`
			}
			if gstr.Pos(tplDao, filterParseStr) == -1 {
				daoObj.filter.parse += filterParseStr
			}
		case TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
		case TypeNameImageSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text
		case TypeNameVideoSuffix: // video,video_list,videoList,video_arr,videoArr等后缀；		类型：单视频varchar，多视频json或text
		case TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
		}
		/*--------根据字段命名类型处理 结束--------*/

		/*--------根据字段数据类型处理（注意：这里是字段命名类型处理的后续操作，改动需考虑兼容） 开始--------*/
		switch v.FieldType {
		case TypeInt: // `int等类型`
		case TypeIntU: // `int等类型（unsigned）`
		case TypeFloat: // `float等类型`
		case TypeFloatU: // `float等类型（unsigned）`
		case TypeVarchar, TypeChar: // `varchar类型`	// `char类型`
			if v.IndexRaw == `UNI` && v.IsNull {
				insertParseStr := `
			case daoThis.Columns().` + v.FieldCaseCamel + `:
				insertData[k] = v
				if gconv.String(v) == ` + "``" + ` {
					insertData[k] = nil
				}`
				if gstr.Pos(tplDao, insertParseStr) == -1 {
					daoObj.insert.parse += insertParseStr
				}
				updateParseStr := `
			case daoThis.Columns().` + v.FieldCaseCamel + `:
				updateData[daoModel.DbTable+` + "`.`" + `+k] = v
				if gconv.String(v) == ` + "``" + ` {
					updateData[daoModel.DbTable+` + "`.`" + `+k] = nil
				}`
				if gstr.Pos(tplDao, updateParseStr) == -1 {
					daoObj.update.parse += updateParseStr
				}
			}
		case TypeText: // `text类型`
		case TypeJson: // `json类型`
			if v.IsNull {
				insertParseStr := `
			case daoThis.Columns().` + v.FieldCaseCamel + `:
				insertData[k] = v
				if gconv.String(v) == ` + "``" + ` {
					insertData[k] = nil
				}`
				if gstr.Pos(tplDao, insertParseStr) == -1 {
					daoObj.insert.parse += insertParseStr
				}
				updateParseStr := `
			case daoThis.Columns().` + v.FieldCaseCamel + `:
				updateData[daoModel.DbTable+` + "`.`" + `+k] = gvar.New(v)
				if gconv.String(v) == ` + "``" + ` {
					updateData[daoModel.DbTable+` + "`.`" + `+k] = nil
				}`
				if gstr.Pos(tplDao, updateParseStr) == -1 {
					daoObj.update.parse += updateParseStr
				}
			}
		case TypeTimestamp: // `timestamp类型`
		case TypeDatetime: // `datetime类型`
		case TypeDate: // `date类型`
			orderParseStr := `
			case daoThis.Columns().` + v.FieldCaseCamel + `:
				m = m.Order(daoModel.DbTable + ` + "`.`" + ` + v)
				m = m.OrderDesc(daoModel.DbTable + ` + "`.`" + ` + daoThis.PrimaryKey())` //追加主键倒序。mysql排序字段有重复值时，分页会导致同一条数据可能在不同页都出现
			if gstr.Pos(tplDao, orderParseStr) == -1 {
				daoObj.order.parse += orderParseStr
			}
		}
		/*--------根据字段数据类型处理（注意：这里是字段命名类型处理的后续操作，改动需考虑兼容） 结束--------*/
	}

	if daoObj.insert.parseBefore != `` {
		pointOfInsertParseBefore := `
		insertData := map[string]interface{}{}`
		tplDao = gstr.Replace(tplDao, pointOfInsertParseBefore, daoObj.insert.parseBefore+pointOfInsertParseBefore, 1)
	}
	if daoObj.insert.parse != `` {
		pointOfInsertParse := `case ` + "`id`" + `:
				insertData[daoThis.PrimaryKey()] = v`
		tplDao = gstr.Replace(tplDao, pointOfInsertParse, pointOfInsertParse+daoObj.insert.parse, 1)
	}
	if daoObj.insert.hook != `` {
		pointOfInsertHook := `// id, _ := result.LastInsertId()

			/* for k, v := range daoModel.AfterInsert {
				switch k {
				case ` + "`xxxx`" + `:
					daoModel.CloneNew().Filter(daoThis.PrimaryKey(), id).HookUpdate(g.Map{k: v}).Update()
				}
			} */`
		tplDao = gstr.Replace(tplDao, pointOfInsertHook, `id, _ := result.LastInsertId()`+daoObj.insert.hook, 1)
	}
	if daoObj.update.parse != `` {
		pointOfUpdateParse := `case ` + "`id`" + `:
				updateData[daoModel.DbTable+` + "`.`" + `+daoThis.PrimaryKey()] = v`
		tplDao = gstr.Replace(tplDao, pointOfUpdateParse, pointOfUpdateParse+daoObj.update.parse, 1)
	}
	if daoObj.update.hookBefore != `` || daoObj.update.hookAfter != `` {
		pointOfUpdateHook := `

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
		if daoObj.update.hookBefore != `` {
			tplDao = gstr.Replace(tplDao, pointOfUpdateHook, daoObj.update.hookBefore+pointOfUpdateHook, 1)
		}
		if daoObj.update.hookAfter != `` {
			tplDao = gstr.Replace(tplDao, pointOfUpdateHook, `

			row, _ := result.RowsAffected()
			if row == 0 {
				return
			}`+daoObj.update.hookAfter, 1)
		}
	}
	if daoObj.field.parse != `` {
		pointOfFieldParse := `case ` + "`id`" + `:
				m = m.Fields(daoModel.DbTable + ` + "`.`" + ` + daoThis.PrimaryKey() + ` + "` AS `" + ` + v)`
		tplDao = gstr.Replace(tplDao, pointOfFieldParse, pointOfFieldParse+daoObj.field.parse, 1)
	}
	if daoObj.field.hook != `` {
		pointOfFieldHook := `
					default:
						record[v] = gvar.New(nil)`
		tplDao = gstr.Replace(tplDao, pointOfFieldHook, daoObj.field.hook+pointOfFieldHook, 1)
	}
	if daoObj.filter.parse != `` {
		pointOfFilterParse := `case ` + "`id`, `idArr`" + `:
				m = m.Where(daoModel.DbTable+` + "`.`" + `+daoThis.PrimaryKey(), v)`
		tplDao = gstr.Replace(tplDao, pointOfFilterParse, pointOfFilterParse+daoObj.filter.parse, 1)
	}
	if daoObj.order.parse != `` {
		pointOfOrderParse := `case ` + "`id`" + `:
				m = m.Order(daoModel.DbTable + ` + "`.`" + ` + gstr.Replace(v, k, daoThis.PrimaryKey(), 1))`
		tplDao = gstr.Replace(tplDao, pointOfOrderParse, pointOfOrderParse+daoObj.order.parse, 1)
	}
	if daoObj.join.parse != `` {
		pointOfJoinParse := `
		/* case Xxxx.ParseDbTable(m.GetCtx()):
		m = m.LeftJoin(joinTable, joinTable+` + "`.`" + `+Xxxx.Columns().XxxxId+` + "` = `" + `+daoModel.DbTable+` + "`.`" + `+daoThis.PrimaryKey())
		// m = m.LeftJoin(Xxxx.ParseDbTable(m.GetCtx())+` + "` AS `" + `+joinTable, joinTable+` + "`.`" + `+Xxxx.Columns().XxxxId+` + "` = `" + `+daoModel.DbTable+` + "`.`" + `+daoThis.PrimaryKey()) */`
		tplDao = gstr.Replace(tplDao, pointOfJoinParse, pointOfJoinParse+daoObj.join.parse, 1)
	}
	if daoObj.importDao != `` {
		pointOfImportDao := `"api/internal/dao/` + tpl.ModuleDirCaseKebab + `/internal"`
		tplDao = gstr.Replace(tplDao, pointOfImportDao, pointOfImportDao+daoObj.importDao, 1)
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
	if tpl.Handle.Pid.Pid != `` {
		tplLogic += `
	if _, ok := data[daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Pid) + `]; ok {
		pid := gconv.Uint(data[daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Pid) + `])
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
	if tpl.Handle.Pid.Pid != `` {
		tplLogic += `
	if _, ok := data[daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Pid) + `]; ok {`
		if tpl.Handle.Pid.IsCoexist {
			tplLogic += `
		pid := gconv.Uint(data[daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Pid) + `])
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
				if pid != oldInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Pid) + `].Uint() {
					if garray.NewStrArrayFrom(gstr.Split(pInfo[daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.IdPath) + `].String(), ` + "`-`" + `)).Contains(oldInfo[daoThis.PrimaryKey()].String()) { //父级不能是自身的子孙级
						err = utils.NewErrorCode(ctx, 29999995, ` + "``" + `)
						return
					}
				}
			}
		}`
		} else {
			tplLogic += `
		pid := gconv.Uint(data[daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Pid) + `])
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
	if tpl.Handle.Pid.Pid != `` {
		tplLogic += `
	count, _ := daoModelThis.CloneNew().Filter(daoThis.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Pid) + `, daoModelThis.IdArr).Count()
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

	type api struct {
		filter   []string
		create   []string
		update   []string
		res      []string
		resOfAdd []string
	}
	apiObj := api{
		filter:   []string{},
		create:   []string{},
		update:   []string{},
		res:      []string{},
		resOfAdd: []string{},
	}
	if len(tpl.Handle.LabelList) > 0 {
		apiObj.filter = append(apiObj.filter, `Label string `+"`"+`json:"label,omitempty" v:"max-length:30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"标签。常用于前端组件"`+"`")
		apiObj.res = append(apiObj.res, `Label *string `+"`"+`json:"label,omitempty" dc:"标签。常用于前端组件"`+"`")
	}

	type apiItem struct {
		isSkip bool //字段命名类型处理后，不再需要根据字段数据类型对filter,create,update,res四个属性再做处理时，设置为true
		filter bool
		create bool
		update bool
		res    bool

		isSkipType bool //字段命名类型处理后，不再需要根据字段数据类型对filterType,createType,updateType,resType四个属性再做处理时，设置为true
		filterType string
		createType string
		updateType string
		resType    string

		filterRule []string
		createRule []string
		updateRule []string
		isRequired bool //用于方便将required规则放首位
	}

	for _, v := range tpl.FieldList {
		apiItemObj := apiItem{
			filterRule: []string{},
			createRule: []string{},
			updateRule: []string{},
		}

		/*--------根据字段命名类型处理 开始--------*/
		switch v.FieldTypeName {
		case TypeNameDeleted: // 软删除字段
			continue
		case TypeNameUpdated: // 更新时间字段
			apiItemObj.isSkip = true
			apiItemObj.res = true
		case TypeNameCreated: // 创建时间字段
			apiObj.filter = append(apiObj.filter,
				`TimeRangeStart *gtime.Time `+"`"+`json:"timeRangeStart,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"`+"`",
				`TimeRangeEnd   *gtime.Time `+"`"+`json:"timeRangeEnd,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"`+"`",
			)
			apiItemObj.isSkip = true
			apiItemObj.res = true
		case TypeNamePri: // 主键
		case TypeNamePriAutoInc: // 主键（自增）
			if v.FieldRaw == `id` {
				continue
			}
			apiItemObj.isSkip = true
			apiItemObj.filter = true
			apiItemObj.res = true

			apiItemObj.filterRule = append(apiItemObj.filterRule, `min:1`)
		case TypeNamePid: // pid；	类型：int等类型；
			apiItemObj.isSkip = true
			apiItemObj.filter = true
			apiItemObj.create = true
			apiItemObj.update = true
			apiItemObj.res = true

			if len(tpl.Handle.LabelList) > 0 {
				apiObj.resOfAdd = append(apiObj.resOfAdd, `P`+gstr.CaseCamel(tpl.Handle.LabelList[0])+` *string `+"`"+`json:"p`+gstr.CaseCamel(tpl.Handle.LabelList[0])+`,omitempty" dc:"父级"`+"`")
			}
		case TypeNameLevel: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
			apiItemObj.isSkip = true
			apiItemObj.filter = true
			apiItemObj.res = true

			apiItemObj.filterRule = append(apiItemObj.filterRule, `min:1`)
		case TypeNameIdPath: // idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
			apiItemObj.isSkip = true
			apiItemObj.res = true
		case TypeNamePasswordSuffix: // password,passwd后缀；		类型：char(32)；
			apiItemObj.isSkip = true
			apiItemObj.create = true
			apiItemObj.update = true

			apiItemObj.isRequired = true
		case TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
			continue
		case TypeNameNameSuffix: // name,title后缀；	类型：varchar；
			// 去掉该验证规则。有时会用到特殊符号
			// apiItemObj.filterRule = append(apiItemObj.filterRule, `regex:^[\\p{L}\\p{M}\\p{N}_-]+$`)
			// apiItemObj.createRule = append(apiItemObj.createRule, `regex:^[\\p{L}\\p{M}\\p{N}_-]+$`)
			// apiItemObj.updateRule = append(apiItemObj.updateRule, `regex:^[\\p{L}\\p{M}\\p{N}_-]+$`)
			if len(tpl.Handle.LabelList) > 0 && gstr.CaseCamel(tpl.Handle.LabelList[0]) == v.FieldCaseCamel {
				apiItemObj.isRequired = true
			}
		case TypeNameCodeSuffix: // code后缀；	类型：varchar；
			apiItemObj.filterRule = append(apiItemObj.filterRule, `regex:^[\\p{L}\\p{M}\\p{N}_-]+$`)
			apiItemObj.createRule = append(apiItemObj.createRule, `regex:^[\\p{L}\\p{M}\\p{N}_-]+$`)
			apiItemObj.updateRule = append(apiItemObj.updateRule, `regex:^[\\p{L}\\p{M}\\p{N}_-]+$`)
		case TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
			apiItemObj.filterRule = append(apiItemObj.filterRule, `phone`)
			apiItemObj.createRule = append(apiItemObj.createRule, `phone`)
			apiItemObj.updateRule = append(apiItemObj.updateRule, `phone`)
		case TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
			apiItemObj.filterRule = append(apiItemObj.filterRule, `url`)
			apiItemObj.createRule = append(apiItemObj.createRule, `url`)
			apiItemObj.updateRule = append(apiItemObj.updateRule, `url`)
		case TypeNameIpSuffix: // IP后缀；	类型：varchar；
			apiItemObj.filterRule = append(apiItemObj.filterRule, `ip`)
			apiItemObj.createRule = append(apiItemObj.createRule, `ip`)
			apiItemObj.updateRule = append(apiItemObj.updateRule, `ip`)
		case TypeNameIdSuffix: // id后缀；	类型：int等类型；
			apiItemObj.isSkip = true
			apiItemObj.filter = true
			apiItemObj.create = true
			apiItemObj.update = true
			apiItemObj.res = true

			apiItemObj.filterRule = append(apiItemObj.filterRule, `min:1`)
			apiItemObj.createRule = append(apiItemObj.createRule, `min:1`)
			apiItemObj.updateRule = append(apiItemObj.updateRule, `min:1`)

			if tpl.Handle.RelIdMap[v.FieldRaw].tpl.TableRaw != `` && !tpl.Handle.RelIdMap[v.FieldRaw].IsRedundName {
				relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
				apiObj.resOfAdd = append(apiObj.resOfAdd, gstr.CaseCamel(relIdObj.tpl.Handle.LabelList[0])+gstr.CaseCamel(relIdObj.Suffix)+` *string `+"`"+`json:"`+relIdObj.tpl.Handle.LabelList[0]+relIdObj.Suffix+`,omitempty" dc:"`+relIdObj.FieldName+`"`+"`")
			}
		case TypeNameSortSuffix, TypeNameSort: // sort,weight等后缀；	类型：int等类型； // sort，且pid,level,idPath|id_path,sort同时存在时（才）有效；	类型：int等类型；
			apiItemObj.createRule = append(apiItemObj.createRule, `between:0,100`)
			apiItemObj.updateRule = append(apiItemObj.updateRule, `between:0,100`)
		case TypeNameStatusSuffix: // status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
			apiItemObj.isSkip = true
			apiItemObj.filter = true
			apiItemObj.create = true
			apiItemObj.update = true
			apiItemObj.res = true

			statusArr := make([]string, len(v.StatusList))
			for index, item := range v.StatusList {
				statusArr[index] = item[0]
			}
			statusStr := gstr.Join(statusArr, `,`)
			apiItemObj.filterRule = append(apiItemObj.filterRule, `in:`+statusStr)
			apiItemObj.createRule = append(apiItemObj.createRule, `in:`+statusStr)
			apiItemObj.updateRule = append(apiItemObj.updateRule, `in:`+statusStr)
		case TypeNameIsPrefix: // is_前缀；		类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）
			apiItemObj.isSkip = true
			apiItemObj.filter = true
			apiItemObj.create = true
			apiItemObj.update = true
			apiItemObj.res = true

			/* TODO 可以和状态一样处理。可能也得改前端开关组件属性
			statusArr := make([]string, len(v.StatusList))
			for index, item := range v.StatusList {
				statusArr[index] = item[0]
			}
			statusStr := gstr.Join(statusArr, `,`)
			apiItemObj.filterRule = append(apiItemObj.filterRule, `in:`+statusStr)
			apiItemObj.createRule = append(apiItemObj.createRule, `in:`+statusStr)
			apiItemObj.updateRule = append(apiItemObj.updateRule, `in:`+statusStr) */
			apiItemObj.filterRule = append(apiItemObj.filterRule, `in:0,1`)
			apiItemObj.createRule = append(apiItemObj.createRule, `in:0,1`)
			apiItemObj.updateRule = append(apiItemObj.updateRule, `in:0,1`)
		case TypeNameStartPrefix: // start_前缀；	类型：timestamp或datetime或date；
			apiItemObj.isSkip = true
			apiItemObj.filter = true
			apiItemObj.create = true
			apiItemObj.update = true
			apiItemObj.res = true
		case TypeNameEndPrefix: // end_前缀；	类型：timestamp或datetime或date；
			apiItemObj.isSkip = true
			apiItemObj.filter = true
			apiItemObj.create = true
			apiItemObj.update = true
			apiItemObj.res = true
		case TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
			apiItemObj.isSkip = true
			apiItemObj.create = true
			apiItemObj.update = true
			apiItemObj.res = true
		case TypeNameImageSuffix, TypeNameVideoSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text	// video,video_list,videoList,video_arr,videoArr等后缀；		类型：单视频varchar，多视频json或text
			apiItemObj.isSkip = true
			apiItemObj.create = true
			apiItemObj.update = true
			apiItemObj.res = true
			if v.FieldType == TypeVarchar {
				apiItemObj.createRule = append(apiItemObj.createRule, `url`)
				apiItemObj.updateRule = append(apiItemObj.updateRule, `url`)
			} else {
				apiItemObj.isSkipType = true
				apiItemObj.createType = `*[]string`
				apiItemObj.updateType = `*[]string`
				apiItemObj.resType = `[]string`

				apiItemObj.createRule = append(apiItemObj.createRule, `distinct`, `foreach`, `url`, `foreach`, `min-length:1`)
				apiItemObj.updateRule = append(apiItemObj.updateRule, `distinct`, `foreach`, `url`, `foreach`, `min-length:1`)
				if !v.IsNull {
					apiItemObj.isRequired = true
				}
			}
		case TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
			apiItemObj.isSkip = true
			apiItemObj.create = true
			apiItemObj.update = true
			apiItemObj.res = true

			apiItemObj.isSkipType = true
			apiItemObj.createType = `*[]interface{}`
			apiItemObj.updateType = `*[]interface{}`
			apiItemObj.resType = `[]interface{}`

			apiItemObj.createRule = append(apiItemObj.createRule, `distinct`)
			apiItemObj.updateRule = append(apiItemObj.updateRule, `distinct`)
			if !v.IsNull {
				apiItemObj.isRequired = true
			}
		}
		/*--------根据字段命名类型处理 结束--------*/

		/*--------根据字段数据类型处理（注意：这里是字段命名类型处理的后续操作，改动需考虑兼容） 开始--------*/
		switch v.FieldType {
		case TypeInt: // `int等类型`
			if !apiItemObj.isSkip {
				// apiItemObj.filter = true
				apiItemObj.create = true
				apiItemObj.update = true
				apiItemObj.res = true
			}
			if !apiItemObj.isSkipType {
				apiItemObj.filterType = `*int`
				apiItemObj.createType = `*int`
				apiItemObj.updateType = `*int`
				apiItemObj.resType = `*int`
			}
		case TypeIntU: // `int等类型（unsigned）`
			if !apiItemObj.isSkip {
				// apiItemObj.filter = true
				apiItemObj.create = true
				apiItemObj.update = true
				apiItemObj.res = true
			}
			if !apiItemObj.isSkipType {
				apiItemObj.filterType = `*uint`
				apiItemObj.createType = `*uint`
				apiItemObj.updateType = `*uint`
				apiItemObj.resType = `*uint`
			}
		case TypeFloat: // `float等类型`
			if !apiItemObj.isSkip {
				// apiItemObj.filter = true
				apiItemObj.create = true
				apiItemObj.update = true
				apiItemObj.res = true
			}
			if !apiItemObj.isSkipType {
				apiItemObj.filterType = `*float64`
				apiItemObj.createType = `*float64`
				apiItemObj.updateType = `*float64`
				apiItemObj.resType = `*float64`
			}
		case TypeFloatU: // // `float等类型（unsigned）`
			if !apiItemObj.isSkip {
				// apiItemObj.filter = true
				apiItemObj.create = true
				apiItemObj.update = true
				apiItemObj.res = true
			}
			if !apiItemObj.isSkipType {
				apiItemObj.filterType = `*float64`
				apiItemObj.createType = `*float64`
				apiItemObj.updateType = `*float64`
				apiItemObj.resType = `*float64`
			}

			apiItemObj.filterRule = append([]string{`min:0`}, apiItemObj.filterRule...)
			apiItemObj.createRule = append([]string{`min:0`}, apiItemObj.createRule...)
			apiItemObj.updateRule = append([]string{`min:0`}, apiItemObj.updateRule...)
		case TypeVarchar: // `varchar类型`
			if !apiItemObj.isSkip {
				apiItemObj.filter = true
				apiItemObj.create = true
				apiItemObj.update = true
				apiItemObj.res = true
			}
			if !apiItemObj.isSkipType {
				apiItemObj.filterType = `string`
				apiItemObj.createType = `*string`
				apiItemObj.updateType = `*string`
				apiItemObj.resType = `*string`
			}

			apiItemObj.filterRule = append([]string{`max-length:` + v.FieldLimitStr}, apiItemObj.filterRule...)
			apiItemObj.createRule = append([]string{`max-length:` + v.FieldLimitStr}, apiItemObj.createRule...)
			apiItemObj.updateRule = append([]string{`max-length:` + v.FieldLimitStr}, apiItemObj.updateRule...)
			if v.IndexRaw == `UNI` && !v.IsNull {
				apiItemObj.isRequired = true
			}
		case TypeChar: // `char类型`
			if !apiItemObj.isSkip {
				apiItemObj.filter = true
				apiItemObj.create = true
				apiItemObj.update = true
				apiItemObj.res = true
			}
			if !apiItemObj.isSkipType {
				apiItemObj.filterType = `string`
				apiItemObj.createType = `*string`
				apiItemObj.updateType = `*string`
				apiItemObj.resType = `*string`
			}

			apiItemObj.filterRule = append([]string{`max-length:` + v.FieldLimitStr}, apiItemObj.filterRule...)
			apiItemObj.createRule = append([]string{`size:` + v.FieldLimitStr}, apiItemObj.createRule...)
			apiItemObj.updateRule = append([]string{`size:` + v.FieldLimitStr}, apiItemObj.updateRule...)
			if v.IndexRaw == `UNI` && !v.IsNull {
				apiItemObj.isRequired = true
			}
		case TypeText: // `text类型`
			if !apiItemObj.isSkip {
				// apiItemObj.filter = true
				apiItemObj.create = true
				apiItemObj.update = true
				apiItemObj.res = true
			}
			if !apiItemObj.isSkipType {
				apiItemObj.filterType = `string`
				apiItemObj.createType = `*string`
				apiItemObj.updateType = `*string`
				apiItemObj.resType = `*string`
			}
		case TypeJson: // `json类型`
			if !apiItemObj.isSkip {
				// apiItemObj.filter = true
				apiItemObj.create = true
				apiItemObj.update = true
				apiItemObj.res = true
			}
			if !apiItemObj.isSkipType {
				apiItemObj.filterType = `string`
				apiItemObj.createType = `*string`
				apiItemObj.updateType = `*string`
				apiItemObj.resType = `*string`
			}

			apiItemObj.filterRule = append([]string{`json`}, apiItemObj.filterRule...)
			apiItemObj.createRule = append([]string{`json`}, apiItemObj.createRule...)
			apiItemObj.updateRule = append([]string{`json`}, apiItemObj.updateRule...)
			if !v.IsNull {
				apiItemObj.isRequired = true
			}
		case TypeTimestamp, TypeDatetime: // `timestamp类型` // `datetime类型`
			if !apiItemObj.isSkip {
				// apiItemObj.filter = true
				apiItemObj.create = true
				apiItemObj.update = true
				apiItemObj.res = true
			}
			if !apiItemObj.isSkipType {
				apiItemObj.filterType = `*gtime.Time`
				apiItemObj.createType = `*gtime.Time`
				apiItemObj.updateType = `*gtime.Time`
				apiItemObj.resType = `*gtime.Time`
			}

			apiItemObj.filterRule = append([]string{`date-format:Y-m-d H:i:s`}, apiItemObj.filterRule...)
			apiItemObj.createRule = append([]string{`date-format:Y-m-d H:i:s`}, apiItemObj.createRule...)
			apiItemObj.updateRule = append([]string{`date-format:Y-m-d H:i:s`}, apiItemObj.updateRule...)
			if !v.IsNull && gconv.String(v.Default) == `` {
				apiItemObj.isRequired = true
			}
		case TypeDate: // `date类型`
			if !apiItemObj.isSkip {
				apiItemObj.filter = true
				apiItemObj.create = true
				apiItemObj.update = true
				apiItemObj.res = true
			}
			if !apiItemObj.isSkipType {
				apiItemObj.filterType = `*gtime.Time`
				apiItemObj.createType = `*gtime.Time`
				apiItemObj.updateType = `*gtime.Time`
				apiItemObj.resType = `*string`
			}

			apiItemObj.filterRule = append([]string{`date-format:Y-m-d`}, apiItemObj.filterRule...)
			apiItemObj.createRule = append([]string{`date-format:Y-m-d`}, apiItemObj.createRule...)
			apiItemObj.updateRule = append([]string{`date-format:Y-m-d`}, apiItemObj.updateRule...)
			if !v.IsNull && gconv.String(v.Default) == `` {
				apiItemObj.isRequired = true
			}
		default:
			if !apiItemObj.isSkip {
				apiItemObj.filter = true
				apiItemObj.create = true
				apiItemObj.update = true
				apiItemObj.res = true
			}
			if !apiItemObj.isSkipType {
				apiItemObj.filterType = `string`
				apiItemObj.createType = `*string`
				apiItemObj.updateType = `*string`
				apiItemObj.resType = `*string`
			}
		}
		/*--------根据字段数据类型处理（注意：这里是字段命名类型处理的后续操作，改动需考虑兼容） 结束--------*/

		if apiItemObj.filter {
			apiObj.filter = append(apiObj.filter, v.FieldCaseCamel+` `+apiItemObj.filterType+` `+"`"+`json:"`+v.FieldRaw+`,omitempty" v:"`+gstr.Join(apiItemObj.filterRule, `|`)+`" dc:"`+v.FieldDesc+`"`+"`")
		}
		if apiItemObj.create {
			if apiItemObj.isRequired {
				apiItemObj.createRule = append([]string{`required`}, apiItemObj.createRule...)
			}
			apiObj.create = append(apiObj.create, v.FieldCaseCamel+` `+apiItemObj.createType+` `+"`"+`json:"`+v.FieldRaw+`,omitempty" v:"`+gstr.Join(apiItemObj.createRule, `|`)+`" dc:"`+v.FieldDesc+`"`+"`")
		}
		if apiItemObj.update {
			apiObj.update = append(apiObj.update, v.FieldCaseCamel+` `+apiItemObj.updateType+` `+"`"+`json:"`+v.FieldRaw+`,omitempty" v:"`+gstr.Join(apiItemObj.updateRule, `|`)+`" dc:"`+v.FieldDesc+`"`+"`")
		}
		if apiItemObj.res {
			apiObj.res = append(apiObj.res, v.FieldCaseCamel+` `+apiItemObj.resType+` `+"`"+`json:"`+v.FieldRaw+`,omitempty" dc:"`+v.FieldDesc+`"`+"`")
		}
	}

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
	ExcIdArr       []uint      ` + "`" + `json:"excIdArr,omitempty" v:"distinct|foreach|min:1" dc:"排除ID数组"` + "`" + gstr.Join(append([]string{``}, apiObj.filter...), `
	`) + `
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
	Id          *uint       ` + "`" + `json:"id,omitempty" dc:"ID"` + "`" + gstr.Join(append([]string{``}, apiObj.res...), `
	`) + gstr.Join(append([]string{``}, apiObj.resOfAdd...), `
	`) + `
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
	Id          *uint       ` + "`" + `json:"id,omitempty" dc:"ID"` + "`" + gstr.Join(append([]string{``}, apiObj.res...), `
	`) + `
}

/*--------详情 结束--------*/

`
	}
	if option.IsCreate {
		tplApi += `/*--------新增 开始--------*/
type ` + tpl.TableCaseCamel + `CreateReq struct {
	g.Meta      ` + "`" + `path:"/` + tpl.TableCaseKebab + `/create" method:"post" tags:"` + myGenThis.sceneName + `/` + option.CommonName + `" sm:"新增"` + "`" + gstr.Join(append([]string{``}, apiObj.create...), `
	`) + `
}

/*--------新增 结束--------*/

`
	}

	if option.IsUpdate {
		tplApi += `/*--------修改 开始--------*/
type ` + tpl.TableCaseCamel + `UpdateReq struct {
	g.Meta      ` + "`" + `path:"/` + tpl.TableCaseKebab + `/update" method:"post" tags:"` + myGenThis.sceneName + `/` + option.CommonName + `" sm:"修改"` + "`" + `
	IdArr       []uint  ` + "`" + `json:"idArr,omitempty" v:"required|distinct|foreach|min:1" dc:"ID数组"` + "`" + gstr.Join(append([]string{``}, apiObj.update...), `
	`) + `
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

	if option.IsList && tpl.Handle.Pid.Pid != `` {
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
	Id       *uint       ` + "`" + `json:"id,omitempty" dc:"ID"` + "`" + gstr.Join(append([]string{``}, apiObj.res...), `
	`) + `
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

	type controller struct {
		importDao string
		list      []string
		info      []string
		tree      []string
		noAuth    []string
		// diff      []string // 可以不要。数据返回时，会根据API文件中的结构体做过滤
	}
	controllerObj := controller{
		list:   []string{"`id`"},
		info:   []string{"`id`"},
		tree:   []string{"`id`"},
		noAuth: []string{"`id`"},
	}

	if len(tpl.Handle.LabelList) > 0 {
		controllerObj.list = append(controllerObj.list, "`label`")
		controllerObj.info = append(controllerObj.info, "`label`")
		controllerObj.tree = append(controllerObj.tree, "`label`")
		if tpl.Handle.Pid.Pid != `` {
			controllerObj.list = append(controllerObj.list, "`p"+gstr.CaseCamel(tpl.Handle.LabelList[0])+"`")
			// controllerObj.info = append(controllerObj.info, "`p"+gstr.CaseCamel(tpl.Handle.LabelList[0])+"`")
		}
		controllerObj.noAuth = append(controllerObj.noAuth, "`label`")
		if tpl.FieldPrimary != `` && tpl.FieldPrimary != `id` {
			controllerObj.noAuth = append(controllerObj.noAuth, `dao`+tpl.ModuleDirCaseCamel+`.`+tpl.TableCaseCamel+`.Columns().`+gstr.CaseCamel(tpl.FieldPrimary))
		}
		controllerObj.noAuth = append(controllerObj.noAuth, `dao`+tpl.ModuleDirCaseCamel+`.`+tpl.TableCaseCamel+`.Columns().`+gstr.CaseCamel(tpl.Handle.LabelList[0]))
		/* for _, v := range tpl.Handle.LabelList {
			controllerObj.noAuth = append(controllerObj.noAuth, `dao`+tpl.ModuleDirCaseCamel+`.`+tpl.TableCaseCamel+`.Columns().`+gstr.CaseCamel(v))
		} */
	}

	for _, v := range tpl.FieldList {
		/*--------根据字段命名类型处理 开始--------*/
		switch v.FieldTypeName {
		case TypeNameDeleted: // 软删除字段
		case TypeNameUpdated: // 更新时间字段
		case TypeNameCreated: // 创建时间字段
		case TypeNamePri: // 主键
		case TypeNamePriAutoInc: // 主键（自增）
		case TypeNamePid: // pid；	类型：int等类型；
		case TypeNameLevel: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
		case TypeNameIdPath: // idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
		case TypeNamePasswordSuffix: // password,passwd后缀；		类型：char(32)；
			// controllerObj.diff = append(controllerObj.diff, `dao`+tpl.ModuleDirCaseCamel+`.`+tpl.TableCaseCamel+`.Columns().`+v.FieldCaseCamel)
		case TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
			// controllerObj.diff = append(controllerObj.diff, `dao`+tpl.ModuleDirCaseCamel+`.`+tpl.TableCaseCamel+`.Columns().`+v.FieldCaseCamel)
		case TypeNameNameSuffix: // name,title后缀；	类型：varchar；
		case TypeNameCodeSuffix: // code后缀；	类型：varchar；
		case TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
		case TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
		case TypeNameIpSuffix: // IP后缀；	类型：varchar；
		case TypeNameIdSuffix: // id后缀；	类型：int等类型；
			relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
			if relIdObj.tpl.TableRaw != `` && !relIdObj.IsRedundName {
				daoPath := `dao` + relIdObj.tpl.ModuleDirCaseCamel + `.` + relIdObj.tpl.TableCaseCamel
				controllerObj.importDao += `
dao` + relIdObj.tpl.ModuleDirCaseCamel + ` "api/internal/dao/` + relIdObj.tpl.ModuleDirCaseKebab + `"`
				fieldTmp := daoPath + `.Columns().` + gstr.CaseCamel(relIdObj.tpl.Handle.LabelList[0])
				if relIdObj.Suffix != `` {
					fieldTmp += "+`" + relIdObj.Suffix + "`"
				}
				controllerObj.list = append(controllerObj.list, fieldTmp)
			}
		case TypeNameSortSuffix, TypeNameSort: // sort,weight等后缀；	类型：int等类型； // sort，且pid,level,idPath|id_path,sort同时存在时（才）有效；	类型：int等类型；
		case TypeNameStatusSuffix: // status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
		case TypeNameIsPrefix: // is_前缀；		类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）
		case TypeNameStartPrefix: // start_前缀；	类型：timestamp或datetime或date；
		case TypeNameEndPrefix: // end_前缀；	类型：timestamp或datetime或date；
		case TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
		case TypeNameImageSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text
		case TypeNameVideoSuffix: // video,video_list,videoList,video_arr,videoArr等后缀；		类型：单视频varchar，多视频json或text
		case TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
		}
		/*--------根据字段命名类型处理 结束--------*/
	}

	tplController := `package controller

import (
	"api/api"
	api` + tpl.ModuleDirCaseCamel + ` "api/api/` + option.SceneCode + `/` + tpl.ModuleDirCaseKebab + `"
	dao` + tpl.ModuleDirCaseCamel + ` "api/internal/dao/` + tpl.ModuleDirCaseKebab + `"` + controllerObj.importDao + `
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
	allowField = append(allowField` + gstr.Join(append([]string{``}, controllerObj.list...), `, `) + `)`
		/* if len(controllerObj.diff) > 0 {
			tplController += `
		allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{` + gstr.Join(controllerObj.diff, `, `) + `})).Slice() //移除敏感字段`
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
		field = []string{` + gstr.Join(controllerObj.noAuth, `, `) + `}
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
	allowField = append(allowField` + gstr.Join(append([]string{``}, controllerObj.info...), `, `) + `)`
		/* if len(controllerObj.diff) > 0 {
			tplController += `
		allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{` + gstr.Join(controllerObj.diff, `, `) + `})).Slice() //移除敏感字段`
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

	if option.IsList && tpl.Handle.Pid.Pid != `` {
		tplController += `
// 列表（树状）
func (controllerThis *` + tpl.TableCaseCamel + `) Tree(ctx context.Context, req *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `TreeReq) (res *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `TreeRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.Map(req.Filter, gconv.MapOption{Deep: true, OmitEmpty: true})
	if filter == nil {
		filter = map[string]interface{}{}
	}

	allowField := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.ColumnArr().Slice()
	allowField = append(allowField` + gstr.Join(append([]string{``}, controllerObj.tree...), `, `) + `)`
		/* if len(controllerObj.diff) > 0 {
			tplController += `
		allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{` + gstr.Join(controllerObj.diff, `, `) + `})).Slice() //移除敏感字段`
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
		field = []string{` + gstr.Join(controllerObj.noAuth, `, `) + `}
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
	tree := utils.Tree(list.List(), 0, dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.Columns().` + gstr.CaseCamel(tpl.FieldPrimary) + `, dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Pid) + `)

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

	type viewList struct {
		rowHeight uint
		columns   []string
	}
	viewListObj := viewList{
		rowHeight: 50,
		columns:   []string{},
	}
	type columnAttr struct {
		dataKey      string
		title        string
		key          string
		align        string
		width        string
		sortable     string
		hidden       string
		cellRenderer string
	}
	for _, v := range tpl.FieldList {
		columnAttrObj := columnAttr{
			dataKey: `'` + v.FieldRaw + `'`,
			title:   `t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')`,
			key:     `'` + v.FieldRaw + `'`,
			align:   `'center'`,
			width:   `150`,
		}

		/*--------根据字段命名类型处理 开始--------*/
		switch v.FieldTypeName {
		case TypeNameDeleted: // 软删除字段
			continue
		case TypeNameUpdated: // 更新时间字段
			columnAttrObj.title = `t('common.name.updatedAt')`
		case TypeNameCreated: // 创建时间字段
			columnAttrObj.title = `t('common.name.createdAt')`
		case TypeNamePri: // 主键
		case TypeNamePriAutoInc: // 主键（自增）
			continue
		case TypeNamePid: // pid；	类型：int等类型；
			if len(tpl.Handle.LabelList) > 0 {
				columnAttrObj.dataKey = `'p` + gstr.CaseCamel(tpl.Handle.LabelList[0]) + `'`
			}
		case TypeNameLevel: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
			columnAttrObj.sortable = `true`
		case TypeNameIdPath: // idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
			columnAttrObj.hidden = `true`
		case TypeNamePasswordSuffix: // password,passwd后缀；		类型：char(32)；
			continue
		case TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
			continue
		case TypeNameNameSuffix: // name,title后缀；	类型：varchar；
		case TypeNameCodeSuffix: // code后缀；	类型：varchar；
		case TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
		case TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
		case TypeNameIpSuffix: // IP后缀；	类型：varchar；
		case TypeNameIdSuffix: // id后缀；	类型：int等类型；
			if tpl.Handle.RelIdMap[v.FieldRaw].tpl.TableRaw != `` && !tpl.Handle.RelIdMap[v.FieldRaw].IsRedundName {
				columnAttrObj.dataKey = `'` + tpl.Handle.RelIdMap[v.FieldRaw].tpl.Handle.LabelList[0] + tpl.Handle.RelIdMap[v.FieldRaw].Suffix + `'`
			}
		case TypeNameSortSuffix, TypeNameSort: // sort,weight等后缀；	类型：int等类型； // sort，且pid,level,idPath|id_path,sort同时存在时（才）有效；	类型：int等类型；
			columnAttrObj.sortable = `true`
			if option.IsUpdate {
				columnAttrObj.cellRenderer = `(props: any): any => {
                if (props.rowData.edit` + gstr.CaseCamel(v.FieldRaw) + `) {
                    let currentRef: any
                    let currentVal = props.rowData.` + v.FieldRaw + `
                    return [
                        <el-input-number
                            ref={(el: any) => {
                                currentRef = el
                                el?.focus()
                            }}
                            model-value={currentVal}
                            placeholder={t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.tip.` + v.FieldRaw + `')}
                            precision={0}
                            min={0}
                            max={100}
                            step={1}
                            step-strictly={true}
                            controls={false} //控制按钮会导致诸多问题。如：焦点丢失；` + v.FieldRaw + `是0或100时，只一个按钮可点击
                            controls-position="right"
                            onChange={(val: number) => (currentVal = val)}
                            onBlur={() => {
                                props.rowData.edit` + gstr.CaseCamel(v.FieldRaw) + ` = false
                                if ((currentVal || currentVal === 0) && currentVal != props.rowData.` + v.FieldRaw + `) {
                                    handleUpdate({
                                        idArr: [props.rowData.id],
                                        ` + v.FieldRaw + `: currentVal,
                                    })
                                        .then((res) => {
                                            props.rowData.` + v.FieldRaw + ` = currentVal
                                        })
                                        .catch((error) => {})
                                }
                            }}
                            onKeydown={(event: any) => {
                                switch (event.keyCode) {
                                    // case 27:    //Esc键：Escape
                                    // case 32:    //空格键：" "
                                    case 13: //Enter键：Enter
                                        // props.rowData.edit` + gstr.CaseCamel(v.FieldRaw) + ` = false    //也会触发onBlur事件
                                        currentRef?.blur()
                                        break
                                }
                            }}
                        />,
                    ]
                }
                return [
                    <div class="inline-edit" onClick={() => (props.rowData.edit` + gstr.CaseCamel(v.FieldRaw) + ` = true)}>
                        {props.rowData.` + v.FieldRaw + `}
                    </div>,
                ]
            }`
			}
		case TypeNameStatusSuffix: // status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
			columnAttrObj.cellRenderer = `(props: any): any => {
                let tagType = tm('config.const.tagType') as string[]
                let obj = tm('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.status.` + v.FieldRaw + `') as { value: any, label: string }[]
                let index = obj.findIndex((item) => { return item.value == props.rowData.` + v.FieldRaw + ` })
                return <el-tag type={tagType[index % tagType.length]}>{obj[index]?.label}</el-tag>
            }`
		case TypeNameIsPrefix: // is_前缀；		类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）
			cellRendererTmp := `disabled={true}`
			if option.IsUpdate {
				cellRendererTmp = `onChange={(val: number) => {
                            handleUpdate({
                                idArr: [props.rowData.id],
                                ` + v.FieldRaw + `: val,
                            })
                                .then((res) => {
                                    props.rowData.` + v.FieldRaw + ` = val
                                })
                                .catch((error) => {})
                        }}`
			}
			columnAttrObj.cellRenderer = `(props: any): any => {
                return [
                    <el-switch
                        model-value={props.rowData.` + v.FieldRaw + `}
                        active-value={1}
                        inactive-value={0}
                        inline-prompt={true}
                        active-text={t('common.yes')}
                        inactive-text={t('common.no')}
                        style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success);"
                        ` + cellRendererTmp + `
                    />,
                ]
            }`
		case TypeNameStartPrefix: // start_前缀；	类型：timestamp或datetime或date；
		case TypeNameEndPrefix: // end_前缀；	类型：timestamp或datetime或date；
		case TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
		case TypeNameImageSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text
			columnAttrObj.width = `100`
			cellRendererTmp := `
                const imageList = [props.rowData.` + v.FieldRaw + `]`
			if v.FieldType != TypeVarchar {
				cellRendererTmp = `
                let imageList: string[]
                if (Array.isArray(props.rowData.` + v.FieldRaw + `)) {
                    imageList = props.rowData.` + v.FieldRaw + `
                } else {
                    imageList = JSON.parse(props.rowData.` + v.FieldRaw + `)
                }`
			}
			columnAttrObj.cellRenderer = `(props: any): any => {
                if (!props.rowData.` + v.FieldRaw + `) {
                    return
                }` + cellRendererTmp + `
                return [
                    <el-scrollbar wrap-style="display: flex; align-items: center;" view-style="margin: auto;">
                        {imageList.map((item) => {
                            //修改宽高时，可同时修改table属性row-height增加行高，则不会显示滚动条
                            return <el-image style="width: 45px;" src={item} lazy={true} hide-on-click-modal={true} preview-teleported={true} preview-src-list={imageList} />
                        })}
                    </el-scrollbar>
                ]
            }`
			goto skipFieldTypeOfViewList
		case TypeNameVideoSuffix: // video,video_list,videoList,video_arr,videoArr等后缀；		类型：单视频varchar，多视频json或text
			if viewListObj.rowHeight < 100 {
				viewListObj.rowHeight = 100
			}
			cellRendererTmp := `
                const videoList = [props.rowData.` + v.FieldRaw + `]`
			if v.FieldType != TypeVarchar {
				cellRendererTmp = `
                let videoList: string[]
                if (Array.isArray(props.rowData.` + v.FieldRaw + `)) {
                    videoList = props.rowData.` + v.FieldRaw + `
                } else {
                    videoList = JSON.parse(props.rowData.` + v.FieldRaw + `)
                }`
			}
			columnAttrObj.cellRenderer = `(props: any): any => {
                if (!props.rowData.` + v.FieldRaw + `) {
                    return
                }` + cellRendererTmp + `
                return [
                    <el-scrollbar wrap-style="display: flex; align-items: center;" view-style="margin: auto;">
                        {videoList.map((item) => {
                            //修改宽高时，可同时修改table属性row-height增加行高，则不会显示滚动条
                            return <video style="width: 120px; height: 80px;" preload="none" controls={true} src={item} />
                        })}
                    </el-scrollbar>,
                ]
            }`
			goto skipFieldTypeOfViewList
		case TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
			columnAttrObj.width = `100`
			columnAttrObj.cellRenderer = `(props: any): any => {
                if (!props.rowData.` + v.FieldRaw + `) {
                    return
                }
                let arrList: any[]
                if (Array.isArray(props.rowData.` + v.FieldRaw + `)) {
                    arrList = props.rowData.` + v.FieldRaw + `
                } else {
                    arrList = JSON.parse(props.rowData.` + v.FieldRaw + `)
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
            }`
			goto skipFieldTypeOfViewList
		}
		/*--------根据字段命名类型处理 结束--------*/

		/*--------根据字段数据类型处理（注意：这里是字段命名类型处理的后续操作，改动需考虑兼容） 开始--------*/
		switch v.FieldType {
		case TypeInt, TypeIntU: // `int等类型` // `int等类型（unsigned）`
			if gstr.Pos(v.FieldTypeRaw, `tinyint`) != -1 || gstr.Pos(v.FieldTypeRaw, `smallint`) != -1 {
				columnAttrObj.width = `100`
			}
		case TypeFloat: // `float等类型`
		case TypeFloatU: // `float等类型（unsigned）`
		case TypeVarchar, TypeChar: // `varchar类型` // `char类型`
			if gconv.Uint(v.FieldLimitStr) >= 120 {
				columnAttrObj.width = `200`
				columnAttrObj.hidden = `true`
			}
		case TypeText, TypeJson: // `text类型` // `json类型`
			columnAttrObj.width = `200`
			columnAttrObj.hidden = `true`
		case TypeTimestamp, TypeDatetime: // `timestamp类型` // `datetime类型`
			columnAttrObj.sortable = `true`
		case TypeDate: // `date类型`
			columnAttrObj.width = `100`
			columnAttrObj.sortable = `true`
		}
		/*--------根据字段数据类型处理（注意：这里是字段命名类型处理的后续操作，改动需考虑兼容） 结束--------*/

	skipFieldTypeOfViewList: //跳过字段数据类型处理标签
		columnAttrStr := []string{
			`dataKey: ` + columnAttrObj.dataKey + `,`,
			`title: ` + columnAttrObj.title + `,`,
			`key: ` + columnAttrObj.key + `,`,
			`align: ` + columnAttrObj.align + `,`,
			`width: ` + columnAttrObj.width + `,`,
		}
		if columnAttrObj.sortable != `` {
			columnAttrStr = append(columnAttrStr, `sortable: `+columnAttrObj.sortable+`,`)
		}
		if columnAttrObj.hidden != `` {
			columnAttrStr = append(columnAttrStr, `hidden: `+columnAttrObj.hidden+`,`)
		}
		if columnAttrObj.cellRenderer != `` {
			columnAttrStr = append(columnAttrStr, `cellRenderer: `+columnAttrObj.cellRenderer+`,`)
		}
		viewListObj.columns = append(viewListObj.columns, `{`+gstr.Join(append([]string{``}, columnAttrStr...), `
            `)+`
        },`)
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
        },` + gstr.Join(append([]string{``}, viewListObj.columns...), `
        `)
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
                <el-table-v2 class="main-table" :columns="table.columns" :data="table.data" :sort-by="table.sort" @column-sort="table.handleSort" :width="width" :height="height" :fixed="true" :row-height="` + gconv.String(viewListObj.rowHeight) + `">
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

	type viewQuery struct {
		dataInit []string
		form     []string
	}
	viewQueryObj := viewQuery{
		dataInit: []string{},
		form:     []string{},
	}

	type viewQueryItem struct {
		elDatePicker struct {
			defaultTime string
		}
	}
	for _, v := range tpl.FieldList {
		viewQueryItemObj := viewQueryItem{}
		/*--------根据字段命名类型处理 开始--------*/
		switch v.FieldTypeName {
		case TypeNameDeleted: // 软删除字段
			continue
		case TypeNameUpdated: // 更新时间字段
			continue
		case TypeNameCreated: // 创建时间字段
			viewQueryObj.dataInit = append(viewQueryObj.dataInit,
				`timeRange: (() => {
        return undefined
        /* const date = new Date()
        return [
            new Date(date.getFullYear(), date.getMonth(), date.getDate(), 0, 0, 0),
            new Date(date.getFullYear(), date.getMonth(), date.getDate(), 23, 59, 59),
        ] */
    })(),`,
				`timeRangeStart: computed(() => {
        if (queryCommon.data.timeRange?.length) {
            return dayjs(queryCommon.data.timeRange[0]).format('YYYY-MM-DD HH:mm:ss')
        }
        return ''
    }),`,
				`timeRangeEnd: computed(() => {
        if (queryCommon.data.timeRange?.length) {
            return dayjs(queryCommon.data.timeRange[1]).format('YYYY-MM-DD HH:mm:ss')
        }
        return ''
    }),`,
			)
			viewQueryObj.form = append(viewQueryObj.form, `<el-form-item prop="timeRange">
            <el-date-picker v-model="queryCommon.data.timeRange" type="datetimerange" range-separator="-" :default-time="[new Date(2000, 0, 1, 0, 0, 0), new Date(2000, 0, 1, 23, 59, 59)]" :start-placeholder="t('common.name.timeRangeStart')" :end-placeholder="t('common.name.timeRangeEnd')" />
        </el-form-item>`)
			continue
		case TypeNamePri: // 主键
		case TypeNamePriAutoInc: // 主键（自增）
			continue
		case TypeNamePid: // pid；	类型：int等类型；
			viewQueryObj.form = append(viewQueryObj.form, `<el-form-item prop="`+v.FieldRaw+`">
            <my-cascader v-model="queryCommon.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/`+tpl.ModuleDirCaseKebab+`/`+tpl.TableCaseKebab+`/tree' }" :defaultOptions="[{ id: 0, label: t('common.name.allTopLevel') }]" :props="{ checkStrictly: true, emitPath: false }" />
        </el-form-item>`)
		case TypeNameLevel: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
			viewQueryObj.form = append(viewQueryObj.form, `<el-form-item prop="`+v.FieldRaw+`">
            <el-input-number v-model="queryCommon.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" :min="1" :controls="false" />
        </el-form-item>`)
		case TypeNameIdPath: // idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
			continue
		case TypeNamePasswordSuffix: // password,passwd后缀；		类型：char(32)；
			continue
		case TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
			continue
		case TypeNameNameSuffix: // name,title后缀；	类型：varchar；
		case TypeNameCodeSuffix: // code后缀；	类型：varchar；
		case TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
		case TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
		case TypeNameIpSuffix: // IP后缀；	类型：varchar；
		case TypeNameIdSuffix: // id后缀；	类型：int等类型；
			apiUrl := tpl.ModuleDirCaseKebab + `/` + gstr.CaseKebab(gstr.SubStr(v.FieldCaseCamelRemove, 0, -2))
			if tpl.Handle.RelIdMap[v.FieldRaw].tpl.TableRaw != `` {
				relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
				apiUrl = relIdObj.tpl.ModuleDirCaseKebab + `/` + relIdObj.tpl.TableCaseKebab
			}
			if tpl.Handle.RelIdMap[v.FieldRaw].tpl.Handle.Pid.Pid != `` {
				viewQueryObj.form = append(viewQueryObj.form, `<el-form-item prop="`+v.FieldRaw+`">
            <my-cascader v-model="queryCommon.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/`+apiUrl+`/tree' }" :props="{ emitPath: false }" />
        </el-form-item>`)
			} else {
				viewQueryObj.form = append(viewQueryObj.form, `<el-form-item prop="`+v.FieldRaw+`">
            <my-select v-model="queryCommon.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/`+apiUrl+`/list' }" />
        </el-form-item>`)
			}
		case TypeNameSortSuffix, TypeNameSort: // sort,weight等后缀；	类型：int等类型； // sort，且pid,level,idPath|id_path,sort同时存在时（才）有效；	类型：int等类型；
			continue
		case TypeNameStatusSuffix: // status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
			viewQueryObj.form = append(viewQueryObj.form, `<el-form-item prop="`+v.FieldRaw+`" style="width: 120px">
            <el-select-v2 v-model="queryCommon.data.`+v.FieldRaw+`" :options="tm('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.status.`+v.FieldRaw+`')" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" :clearable="true" />
        </el-form-item>`)
			continue
		case TypeNameIsPrefix: // is_前缀；		类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）
			viewQueryObj.form = append(viewQueryObj.form, `<el-form-item prop="`+v.FieldRaw+`" style="width: 120px">
            <el-select-v2 v-model="queryCommon.data.`+v.FieldRaw+`" :options="tm('common.status.whether')" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" :clearable="true" />
        </el-form-item>`)
		case TypeNameStartPrefix: // start_前缀；	类型：timestamp或datetime或date；
		case TypeNameEndPrefix: // end_前缀；	类型：timestamp或datetime或date；
			if v.FieldType != TypeDate {
				viewQueryItemObj.elDatePicker.defaultTime = ` :default-time="new Date(2000, 0, 1, 23, 59, 59)"`
			}
		case TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
			continue
		case TypeNameImageSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text
			continue
		case TypeNameVideoSuffix: // video,video_list,videoList,video_arr,videoArr等后缀；		类型：单视频varchar，多视频json或text
			continue
		case TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
			continue
		}
		/*--------根据字段命名类型处理 结束--------*/

		/*--------根据字段数据类型处理（注意：这里是字段命名类型处理的后续操作，改动需考虑兼容） 开始--------*/
		switch v.FieldType {
		case TypeInt: // `int等类型`
			/* viewQueryObj.form += `
			<el-form-item prop="` + v.FieldRaw + `">
				<el-input-number v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" :controls="false" />
			</el-form-item>` */
		case TypeIntU: // `int等类型（unsigned）`
			/* viewQueryObj.form += `
			<el-form-item prop="` + v.FieldRaw + `">
				<el-input-number v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" :min="0" :controls="false" />
			</el-form-item>` */
		case TypeFloat: // `float等类型`
			/* viewQueryObj.form += `
			<el-form-item prop="` + v.FieldRaw + `">
				<el-input-number v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" :precision="` + v.FieldLimitFloat[1] + `" :controls="false" />
			</el-form-item>` */
		case TypeFloatU: // `float等类型（unsigned）`
			/* viewQueryObj.form += `
			<el-form-item prop="` + v.FieldRaw + `">
				<el-input-number v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" :min="0" :precision="` + v.FieldLimitFloat[1] + `" :controls="false" />
			</el-form-item>` */
		case TypeVarchar: // `varchar类型`
			viewQueryObj.form = append(viewQueryObj.form, `<el-form-item prop="`+v.FieldRaw+`">
            <el-input v-model="queryCommon.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" maxlength="`+v.FieldLimitStr+`" :clearable="true" />
        </el-form-item>`)
		case TypeChar: // `char类型`
			viewQueryObj.form = append(viewQueryObj.form, `<el-form-item prop="`+v.FieldRaw+`">
            <el-input v-model="queryCommon.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" minlength="`+v.FieldLimitStr+`" maxlength="`+v.FieldLimitStr+`" :clearable="true" />
        </el-form-item>`)
		case TypeText: // `text类型`
		case TypeJson: // `json类型`
		case TypeTimestamp, TypeDatetime: // `timestamp类型` // `datetime类型`
			viewQueryObj.form = append(viewQueryObj.form, `<el-form-item prop="`+v.FieldRaw+`">
            <el-date-picker v-model="queryCommon.data.`+v.FieldRaw+`" type="datetime" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" format="YYYY-MM-DD HH:mm:ss" value-format="YYYY-MM-DD HH:mm:ss"`+viewQueryItemObj.elDatePicker.defaultTime+` />
        </el-form-item>`)
		case TypeDate: // `date类型`
			viewQueryObj.form = append(viewQueryObj.form, `<el-form-item prop="`+v.FieldRaw+`">
            <el-date-picker v-model="queryCommon.data.`+v.FieldRaw+`" type="date" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" format="YYYY-MM-DD" value-format="YYYY-MM-DD" />
        </el-form-item>`)
		default:
			viewQueryObj.form = append(viewQueryObj.form, `<el-form-item prop="`+v.FieldRaw+`">
            <el-input v-model="queryCommon.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" :clearable="true" />
        </el-form-item>`)
		}
		/*--------根据字段数据类型处理（注意：这里是字段命名类型处理的后续操作，改动需考虑兼容） 结束--------*/
	}

	tplView := `<script setup lang="tsx">
import dayjs from 'dayjs'

const { t, tm } = useI18n()

const queryCommon = inject('queryCommon') as { data: { [propName: string]: any } }
queryCommon.data = {
    ...queryCommon.data,` + gstr.Join(append([]string{``}, viewQueryObj.dataInit...), `
    `) + `
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
        </el-form-item>` + gstr.Join(append([]string{``}, viewQueryObj.form...), `
        `) + `
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

	type viewSave struct {
		importModule   []string
		paramHandle    []string
		dataInitBefore []string
		dataInitAfter  []string
		rule           []string
		form           []string
		formHandle     []string
	}
	viewSaveObj := viewSave{
		importModule:   []string{},
		paramHandle:    []string{},
		dataInitBefore: []string{},
		dataInitAfter:  []string{},
		rule:           []string{},
		form:           []string{},
		formHandle:     []string{},
	}

	type viewSaveItem struct {
		rule         []string
		required     string
		elDatePicker struct {
			defaultTime string
		}
	}
	for _, v := range tpl.FieldList {
		viewSaveItemObj := viewSaveItem{
			rule: []string{},
		}

		/*--------根据字段命名类型处理 开始--------*/
		switch v.FieldTypeName {
		case TypeNameDeleted: // 软删除字段
			continue
		case TypeNameUpdated: // 更新时间字段
			continue
		case TypeNameCreated: // 创建时间字段
			continue
		case TypeNamePri: // 主键
		case TypeNamePriAutoInc: // 主键（自增）
			continue
		case TypeNamePid: // pid；	类型：int等类型；
			viewSaveObj.paramHandle = append(viewSaveObj.paramHandle, `param.`+v.FieldRaw+` === undefined ? param.`+v.FieldRaw+` = 0 : null`)
			viewSaveObj.rule = append(viewSaveObj.rule, v.FieldRaw+`: [
            { type: 'integer', min: 0, trigger: 'change', message: t('validation.select') },
        ],`)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <my-cascader v-model="saveForm.data.`+v.FieldRaw+`" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/`+tpl.ModuleDirCaseKebab+`/`+tpl.TableCaseKebab+`/tree', param: { filter: { excIdArr: saveForm.data.idArr } } }" :props="{ checkStrictly: true, emitPath: false }" />
                </el-form-item>`)
			continue
		case TypeNameLevel: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
			continue
		case TypeNameIdPath: // idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
			continue
		case TypeNamePasswordSuffix: // password,passwd后缀；		类型：char(32)；
			if !garray.NewStrArrayFrom(viewSaveObj.importModule).Contains(`import md5 from 'js-md5'`) {
				viewSaveObj.importModule = append(viewSaveObj.importModule, `import md5 from 'js-md5'`)
			}
			viewSaveObj.paramHandle = append(viewSaveObj.paramHandle, `param.`+v.FieldRaw+` ? param.`+v.FieldRaw+` = md5(param.`+v.FieldRaw+`) : delete param.`+v.FieldRaw)
			viewSaveObj.rule = append(viewSaveObj.rule, v.FieldRaw+`: [
            { type: 'string', required: computed((): boolean => { return saveForm.data.idArr?.length ? false : true; }), min: 6, max: 20, trigger: 'blur', message: t('validation.between.string', { min: 6, max: 20 }) },
        ],`)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-input v-model="saveForm.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" minlength="6" maxlength="20" :show-word-limit="true" :clearable="true" :show-password="true" style="max-width: 250px" />
                    <label v-if="saveForm.data.idArr?.length">
                        <el-alert :title="t('common.tip.notRequired')" type="info" :show-icon="true" :closable="false" />
                    </label>
                </el-form-item>`)
			continue
		case TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
			continue
		case TypeNameNameSuffix: // name,title后缀；	类型：varchar；
			if len(tpl.Handle.LabelList) > 0 && gstr.CaseCamel(tpl.Handle.LabelList[0]) == v.FieldCaseCamel {
				viewSaveItemObj.required = ` required: true,`
			}
			// 去掉该验证规则。有时会用到特殊符号
			// viewSaveItemObj.rule = append(viewSaveItemObj.rule, `{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') },`)
		case TypeNameCodeSuffix: // code后缀；	类型：varchar；
			viewSaveItemObj.rule = append(viewSaveItemObj.rule, `{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') },`)
		case TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
			viewSaveItemObj.rule = append(viewSaveItemObj.rule, `{ pattern: /^1[3-9]\d{9}$/, trigger: 'blur', message: t('validation.phone') },`)
		case TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
			viewSaveItemObj.rule = append(viewSaveItemObj.rule, `{ type: 'url', trigger: 'change', message: t('validation.url') },`)
		case TypeNameIpSuffix: // IP后缀；	类型：varchar；
		case TypeNameIdSuffix: // id后缀；	类型：int等类型；
			apiUrl := tpl.ModuleDirCaseKebab + `/` + gstr.CaseKebab(gstr.SubStr(v.FieldCaseCamelRemove, 0, -2))
			if tpl.Handle.RelIdMap[v.FieldRaw].tpl.TableRaw != `` {
				relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
				apiUrl = relIdObj.tpl.ModuleDirCaseKebab + `/` + relIdObj.tpl.TableCaseKebab
			}
			viewSaveObj.paramHandle = append(viewSaveObj.paramHandle, `param.`+v.FieldRaw+` === undefined ? param.`+v.FieldRaw+` = 0 : null`)
			viewSaveObj.rule = append(viewSaveObj.rule, v.FieldRaw+`: [
            { type: 'integer', /* required: true, */ min: 1, trigger: 'change', message: t('validation.select') },
        ],`)
			if tpl.Handle.RelIdMap[v.FieldRaw].tpl.Handle.Pid.Pid != `` {
				viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <my-cascader v-model="saveForm.data.`+v.FieldRaw+`" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/`+apiUrl+`/tree' }" :props="{ emitPath: false }" />
                </el-form-item>`)
			} else {
				viewSaveObj.dataInitAfter = append(viewSaveObj.dataInitAfter, v.FieldRaw+`: saveCommon.data.`+v.FieldRaw+` ? saveCommon.data.`+v.FieldRaw+` : undefined,`)
				viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <my-select v-model="saveForm.data.`+v.FieldRaw+`" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/`+apiUrl+`/list' }" />
                </el-form-item>`)
			}
			continue
		case TypeNameSortSuffix, TypeNameSort: // sort,weight等后缀；	类型：int等类型； // sort，且pid,level,idPath|id_path,sort同时存在时（才）有效；	类型：int等类型；
			defaultVal := gconv.Int(v.Default)
			if defaultVal != 0 {
				viewSaveObj.dataInitBefore = append(viewSaveObj.dataInitBefore, v.FieldRaw+`: `+gconv.String(defaultVal)+`,`)
			}
			viewSaveObj.rule = append(viewSaveObj.rule, v.FieldRaw+`: [
            { type: 'integer', min: 0, max: 100, trigger: 'change', message: t('validation.between.number', { min: 0, max: 100 }) },
        ],`)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-input-number v-model="saveForm.data.`+v.FieldRaw+`" :precision="0" :min="0" :max="100" :step="1" :step-strictly="true" controls-position="right" :value-on-clear="`+gconv.String(defaultVal)+`" />
                    <label>
                        <el-alert :title="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.tip.`+v.FieldRaw+`')" type="info" :show-icon="true" :closable="false" />
                    </label>
                </el-form-item>`)
			continue
		case TypeNameStatusSuffix: // status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
			defaultVal := gconv.String(v.Default)
			if defaultVal == `` {
				defaultVal = v.StatusList[0][0]
			}
			if garray.NewFrom([]interface{}{TypeVarchar, TypeChar}).Contains(v.FieldType) {
				viewSaveObj.dataInitBefore = append(viewSaveObj.dataInitBefore, v.FieldRaw+`: '`+defaultVal+`',`)
			} else {
				viewSaveObj.dataInitBefore = append(viewSaveObj.dataInitBefore, v.FieldRaw+`: `+defaultVal+`,`)
			}

			viewSaveObj.rule = append(viewSaveObj.rule, v.FieldRaw+`: [
            { type: 'enum', enum: (tm('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.status.`+v.FieldRaw+`') as any).map((item: any) => item.value), trigger: 'change', message: t('validation.select') },
        ],`)

			//超过5个状态用select组件，小于5个用radio组件
			if len(v.StatusList) > 5 {
				viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-select-v2 v-model="saveForm.data.`+v.FieldRaw+`" :options="tm('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.status.`+v.FieldRaw+`')" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" :clearable="false" />
                </el-form-item>`)
			} else {
				viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-radio-group v-model="saveForm.data.`+v.FieldRaw+`">
                        <el-radio v-for="(item, index) in (tm('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.status.`+v.FieldRaw+`') as any)" :key="index" :label="item.value">
                            {{ item.label }}
                        </el-radio>
                    </el-radio-group>
                </el-form-item>`)
			}
			continue
		case TypeNameIsPrefix: // is_前缀；		类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）
			defaultVal := gconv.Int(v.Default)
			if defaultVal != 0 {
				viewSaveObj.dataInitBefore = append(viewSaveObj.dataInitBefore, v.FieldRaw+`: `+gconv.String(defaultVal)+`,`)
			}
			viewSaveObj.rule = append(viewSaveObj.rule, v.FieldRaw+`: [
            { type: 'enum', enum: (tm('common.status.whether') as any).map((item: any) => item.value), trigger: 'change', message: t('validation.select') },
        ],`)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-switch v-model="saveForm.data.`+v.FieldRaw+`" :active-value="1" :inactive-value="0" :inline-prompt="true" :active-text="t('common.yes')" :inactive-text="t('common.no')" style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success);" />
                </el-form-item>`)
			continue
		case TypeNameStartPrefix: // start_前缀；	类型：timestamp或datetime或date；
		case TypeNameEndPrefix: // end_前缀；	类型：timestamp或datetime或date；
			if v.FieldType != TypeDate {
				viewSaveItemObj.elDatePicker.defaultTime = ` :default-time="new Date(2000, 0, 1, 23, 59, 59)"`
			}
		case TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
			if v.FieldType == TypeVarchar {
				viewSaveObj.rule = append(viewSaveObj.rule, v.FieldRaw+`: [
            { type: 'string', max: `+v.FieldLimitStr+`, trigger: 'blur', message: t('validation.max.string', { max: `+v.FieldLimitStr+` }) },
        ],`)
				viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-input v-model="saveForm.data.`+v.FieldRaw+`" type="textarea" :autosize="{ minRows: 3 }" maxlength="`+v.FieldLimitStr+`" :show-word-limit="true" />
                </el-form-item>`)
				continue
			}
		case TypeNameImageSuffix, TypeNameVideoSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text // video,video_list,videoList,video_arr,videoArr等后缀；		类型：单视频varchar，多视频json或text
			multipleStr := ``
			if v.FieldType == TypeVarchar {
				viewSaveObj.rule = append(viewSaveObj.rule, v.FieldRaw+`: [
            { type: 'string', max: `+v.FieldLimitStr+`, trigger: 'blur', message: t('validation.max.string', { max: `+v.FieldLimitStr+` }) },
            { type: 'url', trigger: 'change', message: t('validation.upload') },
        ],`)
			} else {
				multipleStr = ` :multiple="true"`
				requiredStr := ``
				if !v.IsNull {
					requiredStr = ` required: true,`
				}
				viewSaveObj.rule = append(viewSaveObj.rule, v.FieldRaw+`: [
            { type: 'array',`+requiredStr+` trigger: 'change', message: t('validation.upload'), defaultField: { type: 'url', message: t('validation.url') } },
            // { type: 'array',`+requiredStr+` max: 10, trigger: 'change', message: t('validation.max.upload', { max: 10 }), defaultField: { type: 'url', message: t('validation.url') } },
        ],`)
			}

			acceptStr := ` accept="image/*"`
			if v.FieldTypeName == TypeNameVideoSuffix {
				acceptStr = ` accept="video/*" :isImage="false"`
			}
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <my-upload v-model="saveForm.data.`+v.FieldRaw+`"`+multipleStr+acceptStr+` />
                </el-form-item>`)
			continue
		case TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
			viewSaveObj.dataInitBefore = append(viewSaveObj.dataInitBefore, v.FieldRaw+`: [],`)
			requiredStr := ``
			if !v.IsNull {
				requiredStr = ` required: true,`
			}
			viewSaveObj.rule = append(viewSaveObj.rule, v.FieldRaw+`: [
            { type: 'array',`+requiredStr+` trigger: 'change', message: t('validation.required') },
            // { type: 'array',`+requiredStr+` max: 10, trigger: 'change', message: t('validation.max.array', { max: 10 }), defaultField: { type: 'string', message: t('validation.input') } },
        ],`)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-tag v-for="(item, index) in saveForm.data.`+v.FieldRaw+`" :type="`+v.FieldRaw+`Handle.tagType[index % `+v.FieldRaw+`Handle.tagType.length]" @close="`+v.FieldRaw+`Handle.delValue(item)" :key="index" :closable="true" style="margin-right: 10px;">
                        {{ item }}
                    </el-tag>
                    <!-- <el-input-number v-if="`+v.FieldRaw+`Handle.visible" :ref="(el: any) => `+v.FieldRaw+`Handle.ref = el" v-model="`+v.FieldRaw+`Handle.value" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" @keyup.enter="`+v.FieldRaw+`Handle.addValue" @blur="`+v.FieldRaw+`Handle.addValue" size="small" style="width: 100px;" :controls="false" /> -->
                    <el-input v-if="`+v.FieldRaw+`Handle.visible" :ref="(el: any) => `+v.FieldRaw+`Handle.ref = el" v-model="`+v.FieldRaw+`Handle.value" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" @keyup.enter="`+v.FieldRaw+`Handle.addValue" @blur="`+v.FieldRaw+`Handle.addValue" size="small" style="width: 100px;" />
                    <el-button v-else type="primary" size="small" @click="`+v.FieldRaw+`Handle.visibleChange">
                        <autoicon-ep-plus />{{ t('common.add') }}
                    </el-button>
                </el-form-item>`)
			viewSaveObj.formHandle = append(viewSaveObj.formHandle, `const `+v.FieldRaw+`Handle = reactive({
    ref: null as any,
    visible: false,
    value: undefined,
    tagType: tm('config.const.tagType') as string[],
    visibleChange: () => {
        `+v.FieldRaw+`Handle.visible = true
        nextTick(() => {
            `+v.FieldRaw+`Handle.ref?.focus()
        })
    },
    addValue: () => {
        if (`+v.FieldRaw+`Handle.value) {
            saveForm.data.`+v.FieldRaw+`.push(`+v.FieldRaw+`Handle.value)
        }
        `+v.FieldRaw+`Handle.visible = false
        `+v.FieldRaw+`Handle.value = undefined
    },
    delValue: (item: any) => {
        saveForm.data.`+v.FieldRaw+`.splice(saveForm.data.`+v.FieldRaw+`.indexOf(item), 1)
    },
})`)
			continue
		}
		/*--------根据字段命名类型处理 结束--------*/

		/*--------根据字段数据类型处理（注意：这里是字段命名类型处理的后续操作，改动需考虑兼容） 开始--------*/
		switch v.FieldType {
		case TypeInt: // `int等类型`
			defaultVal := gconv.Int(v.Default)
			if defaultVal != 0 {
				viewSaveObj.dataInitBefore = append(viewSaveObj.dataInitBefore, v.FieldRaw+`: `+gconv.String(defaultVal)+`,`)
			}
			viewSaveItemObj.rule = append(viewSaveItemObj.rule, `{ type: 'integer', trigger: 'change', message: t('validation.input') },`)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-input-number v-model="saveForm.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" :controls="false" :value-on-clear="`+gconv.String(defaultVal)+`" />
                </el-form-item>`)
		case TypeIntU: // `int等类型（unsigned）`
			defaultVal := gconv.Uint(v.Default)
			if defaultVal != 0 {
				viewSaveObj.dataInitBefore = append(viewSaveObj.dataInitBefore, v.FieldRaw+`: `+gconv.String(defaultVal)+`,`)
			}
			viewSaveItemObj.rule = append(viewSaveItemObj.rule, `{ type: 'integer', min: 0, trigger: 'change', message: t('validation.min.number', { min: 0 }) },`)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-input-number v-model="saveForm.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" :min="0" :controls="false" :value-on-clear="`+gconv.String(defaultVal)+`" />
                </el-form-item>`)
		case TypeFloat: // `float等类型`
			defaultVal := gconv.Float64(v.Default)
			if defaultVal != 0 {
				viewSaveObj.dataInitBefore = append(viewSaveObj.dataInitBefore, v.FieldRaw+`: `+gconv.String(defaultVal)+`,`)
			}
			viewSaveItemObj.rule = append(viewSaveItemObj.rule, `{ type: 'number'/* 'float' */, trigger: 'change', message: t('validation.input') },    // 类型float值为0时验证不能通过`)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-input-number v-model="saveForm.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" :precision="`+v.FieldLimitFloat[1]+`" :controls="false" :value-on-clear="`+gconv.String(defaultVal)+`" />
                </el-form-item>`)
		case TypeFloatU: // `float等类型（unsigned）`
			defaultVal := gconv.Float64(v.Default)
			if defaultVal != 0 {
				viewSaveObj.dataInitBefore = append(viewSaveObj.dataInitBefore, v.FieldRaw+`: `+gconv.String(defaultVal)+`,`)
			}
			viewSaveItemObj.rule = append(viewSaveItemObj.rule, `{ type: 'number'/* 'float' */, min: 0, trigger: 'change', message: t('validation.min.number', { min: 0 }) },    // 类型float值为0时验证不能通过`)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-input-number v-model="saveForm.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" :min="0" :precision="`+v.FieldLimitFloat[1]+`" :controls="false" :value-on-clear="`+gconv.String(defaultVal)+`" />
                </el-form-item>`)
		case TypeVarchar: // `varchar类型`
			if v.IndexRaw == `UNI` && !v.IsNull {
				viewSaveItemObj.required = ` required: true,`
			}
			viewSaveItemObj.rule = append([]string{`{ type: 'string',` + viewSaveItemObj.required + ` max: ` + v.FieldLimitStr + `, trigger: 'blur', message: t('validation.max.string', { max: ` + v.FieldLimitStr + ` }) },`}, viewSaveItemObj.rule...)

			if v.IndexRaw == `UNI` {
				viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-input v-model="saveForm.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" maxlength="`+v.FieldLimitStr+`" :show-word-limit="true" :clearable="true" style="max-width: 250px" />
                    <label>
                        <el-alert :title="t('common.tip.notDuplicate')" type="info" :show-icon="true" :closable="false" />
                    </label>
                </el-form-item>`)
			} else {
				viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-input v-model="saveForm.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" maxlength="`+v.FieldLimitStr+`" :show-word-limit="true" :clearable="true" />
                </el-form-item>`)
			}
		case TypeChar: // `char类型`
			if v.IndexRaw == `UNI` && !v.IsNull {
				viewSaveItemObj.required = ` required: true,`
			}
			viewSaveItemObj.rule = append([]string{`{ type: 'string',` + viewSaveItemObj.required + ` len: ` + v.FieldLimitStr + `, trigger: 'blur', message: t('validation.size.string', { size: ` + v.FieldLimitStr + ` }) },`}, viewSaveItemObj.rule...)

			if v.IndexRaw == `UNI` {
				viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-input v-model="saveForm.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" minlength="`+v.FieldLimitStr+`" maxlength="`+v.FieldLimitStr+`" :show-word-limit="true" :clearable="true" />
                </el-form-item>`)
			} else {
				viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-input v-model="saveForm.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" minlength="`+v.FieldLimitStr+`" maxlength="`+v.FieldLimitStr+`" :show-word-limit="true" :clearable="true" style="max-width: 250px" />
                    <label>
                        <el-alert :title="t('common.tip.notDuplicate')" type="info" :show-icon="true" :closable="false" />
                    </label>
                </el-form-item>`)
			}
		case TypeText: // `text类型`
			viewSaveItemObj.rule = append([]string{`{ type: 'string', trigger: 'blur', message: t('validation.input') },`}, viewSaveItemObj.rule...)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <my-editor v-model="saveForm.data.`+v.FieldRaw+`" />
                </el-form-item>`)
		case TypeJson: // `json类型`
			if !v.IsNull {
				viewSaveItemObj.required = `
                required: true,`
			}
			viewSaveItemObj.rule = append(viewSaveItemObj.rule, `{
                type: 'object',`+viewSaveItemObj.required+`
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
            },`)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-alert :title="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.tip.`+v.FieldRaw+`')" type="info" :show-icon="true" :closable="false" />
                    <el-input v-model="saveForm.data.`+v.FieldRaw+`" type="textarea" :autosize="{ minRows: 3 }" />
                </el-form-item>`)
		case TypeTimestamp, TypeDatetime: // `timestamp类型` // `datetime类型`
			if !v.IsNull && gconv.String(v.Default) == `` {
				viewSaveItemObj.required = ` required: true,`
			}
			viewSaveItemObj.rule = append([]string{`{ type: 'string',` + viewSaveItemObj.required + ` trigger: 'change', message: t('validation.select') },`}, viewSaveItemObj.rule...)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-date-picker v-model="saveForm.data.`+v.FieldRaw+`" type="datetime" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" format="YYYY-MM-DD HH:mm:ss" value-format="YYYY-MM-DD HH:mm:ss"`+viewSaveItemObj.elDatePicker.defaultTime+` />
                </el-form-item>`)
		case TypeDate: // `date类型`
			if !v.IsNull && gconv.String(v.Default) == `` {
				viewSaveItemObj.required = ` required: true,`
			}
			viewSaveItemObj.rule = append([]string{`{ type: 'string',` + viewSaveItemObj.required + ` trigger: 'change', message: t('validation.select') },`}, viewSaveItemObj.rule...)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-date-picker v-model="saveForm.data.`+v.FieldRaw+`" type="date" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" format="YYYY-MM-DD" value-format="YYYY-MM-DD" />
                </el-form-item>`)
		default:
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-input v-model="saveForm.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" :clearable="true" />
                </el-form-item>`)
		}
		/*--------根据字段数据类型处理（注意：这里是字段命名类型处理的后续操作，改动需考虑兼容） 结束--------*/

		if len(viewSaveItemObj.rule) > 0 {
			viewSaveObj.rule = append(viewSaveObj.rule, v.FieldRaw+`: [`+gstr.Join(append([]string{``}, viewSaveItemObj.rule...), `
			`)+`
        ],`)
		} else {
			viewSaveObj.rule = append(viewSaveObj.rule, v.FieldRaw+`: [],`)
		}
	}

	tplView := `<script setup lang="tsx">` + gstr.Join(append([]string{``}, viewSaveObj.importModule...), `
`) + `
const { t, tm } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean; title: string; data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {` + gstr.Join(append([]string{``}, viewSaveObj.dataInitBefore...), `
        `) + `
        ...saveCommon.data,` + gstr.Join(append([]string{``}, viewSaveObj.dataInitAfter...), `
        `) + `
    } as { [propName: string]: any },
    rules: {` + gstr.Join(append([]string{``}, viewSaveObj.rule...), `
        `) + `
    } as any,
    submit: () => {
        saveForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return false
            }
            saveForm.loading = true
            const param = removeEmptyOfObj(saveForm.data)` + gstr.Join(append([]string{``}, viewSaveObj.paramHandle...), `
            `) + `
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
})` + gstr.Join(append([]string{``}, viewSaveObj.formHandle...), `

`) + `
</script>

<template>
    <el-drawer class="save-drawer" :ref="(el: any) => saveDrawer.ref = el" v-model="saveCommon.visible" :title="saveCommon.title" :size="saveDrawer.size" :before-close="saveDrawer.beforeClose">
        <el-scrollbar>
            <el-form :ref="(el: any) => saveForm.ref = el" :model="saveForm.data" :rules="saveForm.rules" label-width="auto" :status-icon="true" :scroll-to-error="true">` + gstr.Join(append([]string{``}, viewSaveObj.form...), `
                `) + `
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

	type viewI18n struct {
		name   []string
		status []string
		tip    []string
	}
	viewI18nObj := viewI18n{
		name:   []string{},
		status: []string{},
		tip:    []string{},
	}
	for _, v := range tpl.FieldList {
		/*--------根据字段命名类型处理 开始--------*/
		switch v.FieldTypeName {
		case TypeNameDeleted: // 软删除字段
			continue
		case TypeNameUpdated: // 更新时间字段
			continue
		case TypeNameCreated: // 创建时间字段
			continue
		case TypeNamePri: // 主键
		case TypeNamePriAutoInc: // 主键（自增）
			continue
		case TypeNamePid: // pid；	类型：int等类型；
			viewI18nObj.name = append(viewI18nObj.name, v.FieldRaw+`: '父级',`)
			continue
		case TypeNameLevel: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
		case TypeNameIdPath: // idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
		case TypeNamePasswordSuffix: // password,passwd后缀；		类型：char(32)；
		case TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
			continue
		case TypeNameNameSuffix: // name,title后缀；	类型：varchar；
		case TypeNameCodeSuffix: // code后缀；	类型：varchar；
		case TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
		case TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
		case TypeNameIpSuffix: // IP后缀；	类型：varchar；
		case TypeNameIdSuffix: // id后缀；	类型：int等类型；
			relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
			if relIdObj.tpl.TableRaw != `` && !relIdObj.IsRedundName {
				viewI18nObj.name = append(viewI18nObj.name, v.FieldRaw+`: '`+relIdObj.FieldName+`',`)
				continue
			}
		case TypeNameSortSuffix, TypeNameSort: // sort,weight等后缀；	类型：int等类型； // sort，且pid,level,idPath|id_path,sort同时存在时（才）有效；	类型：int等类型；
			viewI18nObj.tip = append(viewI18nObj.tip, v.FieldRaw+`: '`+v.FieldTip+`',`)
		case TypeNameStatusSuffix: // status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
			statusItem := []string{}
			if garray.NewFrom([]interface{}{TypeVarchar, TypeChar}).Contains(v.FieldType) {
				for _, status := range v.StatusList {
					statusItem = append(statusItem, `{ value: '`+status[0]+`', label: '`+status[1]+`' },`)
				}
			} else {
				for _, status := range v.StatusList {
					statusItem = append(statusItem, `{ value: `+status[0]+`, label: '`+status[1]+`' },`)
				}
			}
			viewI18nObj.status = append(viewI18nObj.status, v.FieldRaw+`: [`+gstr.Join(append([]string{``}, statusItem...), `
            `)+`
        ],`)
		case TypeNameIsPrefix: // is_前缀；		类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）
		case TypeNameStartPrefix: // start_前缀；	类型：timestamp或datetime或date；
		case TypeNameEndPrefix: // end_前缀；	类型：timestamp或datetime或date；
		case TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
		case TypeNameImageSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text
		case TypeNameVideoSuffix: // video,video_list,videoList,video_arr,videoArr等后缀；		类型：单视频varchar，多视频json或text
		case TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
		}
		/*--------根据字段命名类型处理 结束--------*/

		/*--------根据字段数据类型处理（注意：这里是字段命名类型处理的后续操作，改动需考虑兼容） 开始--------*/
		switch v.FieldType {
		case TypeInt: // `int等类型`
		case TypeIntU: // `int等类型（unsigned）`
		case TypeFloat: // `float等类型`
		case TypeFloatU: // `float等类型（unsigned）`
		case TypeVarchar: // `varchar类型`
		case TypeChar: // `char类型`
		case TypeText: // `text类型`
		case TypeJson: // `json类型`
			viewI18nObj.tip = append(viewI18nObj.tip, v.FieldRaw+`: '`+v.FieldTip+`',`)
		case TypeTimestamp: // `timestamp类型`
		case TypeDatetime: // `datetime类型`
		case TypeDate: // `date类型`
		default:
		}
		/*--------根据字段数据类型处理（注意：这里是字段命名类型处理的后续操作，改动需考虑兼容） 结束--------*/
		viewI18nObj.name = append(viewI18nObj.name, v.FieldRaw+`: '`+v.FieldName+`',`)
	}

	tplView := `export default {
    name: {` + gstr.Join(append([]string{``}, viewI18nObj.name...), `
        `) + `
    },
    status: {` + gstr.Join(append([]string{``}, viewI18nObj.status...), `
        `) + `
    },
    tip: {` + gstr.Join(append([]string{``}, viewI18nObj.tip...), `
        `) + `
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

// 执行命令
func (myGenThis *myGenHandler) command(title string, isOut bool, dir string, name string, arg ...string) {
	command := exec.Command(name, arg...)
	if dir != `` {
		command.Dir = dir
	}
	fmt.Println()
	fmt.Println(`================` + title + ` 开始================`)
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
	fmt.Println(`================` + title + ` 结束================`)
}

// status字段注释解析
func (myGenThis *myGenHandler) getStatusList(comment string, isStr bool) (statusList [][2]string) {
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

// 获取Handle.PasswordMap的Key（以Password为主）
func (myGenThis *myGenHandler) getHandlePasswordMapKey(passwordOrsalt string) (passwordMapKey string) {
	passwordOrsalt = gstr.Replace(gstr.CaseCamel(passwordOrsalt), `Salt`, `Password`, 1) //替换salt
	passwordOrsalt = gstr.Replace(passwordOrsalt, `Passwd`, `Password`, 1)               //替换passwd
	passwordMapKey = gstr.CaseCamelLower(passwordOrsalt)                                 //默认：小驼峰
	if gstr.CaseCamelLower(passwordOrsalt) != passwordOrsalt {                           //判断字段是不是蛇形
		passwordMapKey = gstr.CaseSnake(passwordMapKey)
	}
	return
}

// 获取id后缀字段关联的表信息
func (myGenThis *myGenHandler) getRelIdTpl(tpl myGenTpl, field string) (relTpl myGenTpl) {
	fieldCaseSnake := gstr.CaseSnake(field)
	fieldCaseSnakeOfRemove := gstr.Split(fieldCaseSnake, `_of_`)[0]
	tableSuffix := gstr.TrimRightStr(fieldCaseSnakeOfRemove, `_id`, 1)
	/*--------确定关联表 开始--------*/
	// 按以下优先级确定关联表
	type mayBe struct {
		table1 string   // 同模块，全部前缀 + 表后缀一致。规则：tpl.RemovePrefix + tableSuffix
		table2 []string // 同模块，全部前缀 + 任意字符_ + 表后缀一致。规则：tpl.RemovePrefix + xx_ + tableSuffix。同时存在多个放弃匹配
		table3 string   // 不同模块，公共前缀 + 表后缀一致。规则：tpl.RemovePrefixCommon + tableSuffix
		table4 string   // 不同模块，表后缀一致。规则：tableSuffix
		table5 []string // 不同模块，任意字符_ + 表后缀一致。规则：xx_ + tableSuffix。同时存在多个放弃匹配
	}
	mayBeObj := mayBe{
		table2: []string{},
		table5: []string{},
	}
	isSamePrimaryFunc := func(table string) bool {
		tableIndexList, _ := myGenThis.db.GetAll(myGenThis.ctx, `SHOW Index FROM `+table+` WHERE Key_name = 'PRIMARY'`)
		return len(tableIndexList) == 1 && garray.NewStrArrayFrom([]string{`id`, fieldCaseSnakeOfRemove}).Contains(gstr.CaseSnake(tableIndexList[0][`Column_name`].String()))
	}
	for _, v := range myGenThis.tableArr {
		if v == tpl.TableRaw { //自身跳过
			continue
		}
		if v == tpl.RemovePrefix+tableSuffix { //关联表在同模块目录下，且表名一致
			if isSamePrimaryFunc(v) {
				mayBeObj.table1 = v
				break
			}
		} else if gstr.Pos(v, tpl.RemovePrefix) == 0 && len(v) == gstr.PosR(v, `_`+tableSuffix)+len(`_`+tableSuffix) { //关联表在同模块目录下，但表后缀一致
			if isSamePrimaryFunc(v) {
				mayBeObj.table2 = append(mayBeObj.table2, v)
			}
		} else if v == tpl.RemovePrefixCommon+tableSuffix { //公共前缀+表名完全一致
			if isSamePrimaryFunc(v) {
				mayBeObj.table3 = v
			}
		} else if v == tableSuffix { //表名完全一致
			if isSamePrimaryFunc(v) {
				mayBeObj.table4 = v
			}
		} else if len(v) == gstr.PosR(v, `_`+tableSuffix)+len(`_`+tableSuffix) { //表后缀一致
			if isSamePrimaryFunc(v) {
				mayBeObj.table5 = append(mayBeObj.table5, v)
			}
		}
	}

	table := mayBeObj.table1
	if table == `` {
		if len(mayBeObj.table2) > 0 {
			if len(mayBeObj.table2) == 1 {
				table = mayBeObj.table2[0]
			}
		} else {
			if mayBeObj.table3 != `` {
				table = mayBeObj.table3
			} else if mayBeObj.table4 != `` {
				table = mayBeObj.table4
			} else if len(mayBeObj.table5) > 0 && len(mayBeObj.table5) == 1 {
				table = mayBeObj.table5[0]
			}
		}
	}
	/*--------确定关联表 结束--------*/

	removePrefixCommon := ``
	removePrefixAlone := ``
	if table != `` {
		if gstr.Pos(table, tpl.RemovePrefixCommon) == 0 {
			removePrefixCommon = tpl.RemovePrefixCommon
		}
		if gstr.Pos(table, tpl.RemovePrefix) == 0 {
			removePrefixAlone = tpl.RemovePrefixAlone
		}
		if removePrefixAlone == `` {
			// 当去掉公共前缀后，还存在分隔符`_`时，第一个分隔符之前的部分设置为removePrefixAlone
			tableRemove := gstr.TrimLeftStr(table, removePrefixCommon, 1)
			if gstr.Pos(tableRemove, `_`) != -1 {
				removePrefixAlone = gstr.Split(tableRemove, `_`)[0] + `_`
			}
			if pos := gstr.Pos(tableRemove, `_`); pos != -1 {
				removePrefixAlone = gstr.SubStr(tableRemove, 0, pos+1)
			}
		}

		relTpl = myGenThis.createTpl(table, removePrefixCommon, removePrefixAlone)

		// 判断dao文件是否存在，不存在则生成
		if !gfile.IsFile(gfile.SelfDir() + `/internal/dao/` + relTpl.ModuleDirCaseKebab + `/` + relTpl.TableCaseSnake + `.go`) {
			myGenThis.command(`关联表（`+relTpl.TableRaw+`）dao生成`, true, ``,
				`gf`, `gen`, `dao`,
				`--link`, myGenThis.dbLink,
				`--group`, myGenThis.option.DbGroup,
				`--removePrefix`, relTpl.RemovePrefix,
				`--daoPath`, `dao/`+relTpl.ModuleDirCaseKebab,
				`--doPath`, `model/entity/`+relTpl.ModuleDirCaseKebab,
				`--entityPath`, `model/entity/`+relTpl.ModuleDirCaseKebab,
				`--tables`, relTpl.TableRaw,
				`--tplDaoIndexPath`, `resource/gen/gen_dao_template_dao.txt`,
				`--tplDaoInternalPath`, `resource/gen/gen_dao_template_dao_internal.txt`,
				`--overwriteDao`, `false`)
		}
	}
	return
	// TODO
	/* type relTableItem struct {
		IsSameDir               bool       //关联表dao层是否与当前生成dao层在相同目录下
		RelTableField           string     //关联表字段
		RelTableFieldName       string     //关联表字段名称
		IsRedundRelNameField    bool       //当前表是否冗余关联表字段
		RelSuffix               string     //关联表字段后缀（原始，大驼峰或蛇形）。字段含[_of_]时，_of_及之后的部分。示例：userIdOfSend对应OfSend；user_id_of_send对应_of_send
		RelSuffixCaseCamel      string     //关联表字段后缀（大驼峰）。字段含[_of_]时，_of_及其之后的部分。示例：userIdOfSend和user_id_of_send都对应OfSend
		RelSuffixCaseSnake      string     //关联表字段后缀（蛇形）。字段含[_of_]时，_of_及其之后的部分。示例：userIdOfSend和user_id_of_send都对应_of_send
		RelTableIsExistPidField bool       //关联表是否pid字段。前端Query和Save视图组件则使用my-cascader组件，否则使用my-select组件
	} */
}
