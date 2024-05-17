/*
后台常用生成示例：./main myGen -sceneCode=platform -dbGroup=default -dbTable=auth_test -removePrefixCommon= -removePrefixAlone=auth_ -commonName=权限管理/测试 -isList=1 -isCount=1 -isInfo=1 -isCreate=1 -isUpdate=1 -isDelete=1 -isApi=1 -isAuthAction=1 -isView=1 -isResetLogic=0
APP常用生成示例：./main myGen -sceneCode=app -dbGroup=xxxx -dbTable=user -removePrefixCommon= -removePrefixAlone= -commonName=用户 -isList=1 -isCount=0 -isInfo=1 -isCreate=0 -isUpdate=0 -isDelete=0 -isApi=1 -isAuthAction=0 -isView=0 -isResetLogic=0

强烈建议搭配Git使用

表名统一使用蛇形命名。不同功能表按以下表规则命名
	主表：正常命名即可。参考以下示例
		platform_admin
		user
		good
		good_category
	扩展表（一对一）：表命名：主表名_xxxx，并存在与主表（主键 或 表名去掉前缀 + ID）同名的id后缀字段，且字段设为：非递增主键 或 唯一索引
	扩展表（一对多）：表命名：主表名_xxxx，并存在与主表（主键 或 表名去掉前缀 + ID）同名的id后缀字段，且字段设为：普通索引
		参考以下示例
			user_config		说明：存放user主表用户的配置信息
			good_content	说明：存放good主表商品的详情
			good_image		说明：存放good主表商品的图片
	中间表（一对一）：表命名：主表名_rel_to_xxxx 或 xxxx_rel_of_主表名，同模块时，后面部分可省略独有前缀，并存在至少2个与关联表（主键 或 表名去掉前缀 + ID）同名的id后缀字段。主表的关联字段设为：非递增主键 或 唯一索引
	中间表（一对多）：表命名：主表名_rel_to_xxxx 或 xxxx_rel_of_主表名，同模块时，后面部分可省略独有前缀，并存在至少2个与关联表（主键 或 表名去掉前缀 + ID）同名的id后缀字段。所有表的关联字段设为：联合主键 或 联合唯一索引
	关于扩展表和中间表的区别说明，特别是扩展表（一对一）和中间表（一对一）很容易让人误解：
		扩展表各字段功能独立，故当存在除主表id字段外的其它id后缀字段时，这些id后缀字段在更新时，都可设为0，不会删除与主表id对应的记录。且扩展表记录一般只在主表做删除时，才会删除
		中间表其它非id后缀字段，功能都是依赖于id后缀字段存在的，故当除主表id字段外的其它id后缀字段在更新时，如果都设为0，会删除与主表id对应的记录

表字段名统一使用小驼峰或蛇形命名（建议：小驼峰）
	尽量根据表名设置以下两个字段（作用1：常用于前端部分组件，如my-select或my-cascader等组件；作用2：用于关联表查询）
		xxId主键字段。示例：good表命名goodId, good_category表命名categoryId
			注意：考虑兼容旧数据库，主键可命名为id（id命名只允许用于独立主键，表其它字段禁用；表是联合主键则全表禁用）
		xxName或xxTitle字段。示例：good表命名goodName, article表命名articleTitle
			注意：如果不存在xxName或xxTitle字段，按以下优先级使用
				表名去掉前缀 + Name > 主键去掉ID + Name > Name >
				表名去掉前缀 + Title > 主键去掉ID + Title > Title >
				表名去掉前缀 + Phone > 主键去掉ID + Phone > Phone >
				表名去掉前缀 + Email > 主键去掉ID + Email > Email >
				表名去掉前缀 + Account > 主键去掉ID + Account > Account >
				表名去掉前缀 + Nickname > 主键去掉ID + Nickname > Nickname >
				上面字段都没有时，默认第二个字段

	字段都必须有注释。以下符号[\n\r.。:：(（]之前的部分或整个注释，将作为字段名称使用

	字段按以下规则命名时，会做特殊处理，其它情况根据字段类型做默认处理
		限制命名：
			ID				命名：id；		只允许用于独立主键，表其它字段禁用；表是联合主键则全表禁用
			Label			命名：label；	全表禁用

		固定命名：
			父级			命名：pid；	类型：int等类型；
			层级			命名：level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
			层级路径		命名：idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
				建议直接使用text，当level层级大时，不用考虑字符长度问题。
				当level层级不大时，可使用varchar，但必须设置足够的字段长度，否则会丢失路径后面的部分字符。

		常用命名(字段含[_of_]时，会忽略[_of_]及其之后的部分)：
			密码			命名：password,passwd后缀；								类型：char(32)；
			密码盐 			命名：salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
			名称			命名：name,title后缀；									类型：varchar；
			标识			命名：code后缀；										类型：varchar；
			账号			命名：account后缀；										类型：varchar；
			手机			命名：phone,mobile后缀；								类型：varchar；
			邮箱			命名：email后缀；										类型：varchar；
			链接			命名：url,link后缀；									类型：varchar；
			IP				命名：IP后缀；											类型：varchar；
			关联ID			命名：id后缀；											类型：int等类型；
			排序|数量|权重	命名：sort,num,number,weight等后缀；					类型：int等类型；
			编号|层级|等级	命名：no,level,rank等后缀；								类型：int等类型；
			状态|类型		命名：status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
			是否			命名：is_前缀；											类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）
			开始时间		命名：start_前缀；										类型：datetime或date或timestamp或time；
			结束时间		命名：end_前缀；										类型：datetime或date或timestamp或time；
			(富)文本		命名：remark,desc,msg,message,intro,content后缀；		类型：varchar或text；	前端对应组件：varchar文本输入框，text富文本编辑器
			图片			命名：icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text
			视频			命名：video,video_list,videoList,video_arr,videoArr等后缀；	类型：单视频varchar，多视频json或text
			文件			命名：file,file_list,fileList,file_arr,fileArr等后缀；	类型：单文件varchar，多文件json或text
			数组			命名：list,arr等后缀；									类型：json或text；
*/

package my_gen

import (
	"api/internal/cmd/my-gen/internal"
	daoAuth "api/internal/dao/auth"
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// 命令参数解析后的数据
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
	/*
		是否重置logic层。一般情况下不建议重置，原因：logic层生成基本不会有任何变化，且常会在该层手写一些逻辑验证和自定义方法。只建议在以下两种情况下重置：
			1、logic层生成模板发生重大变化，即对gen_logic.go中生成的模板代码做修改
			2、表新增或删除了对logic层生成代码有影响的字段。目前有影响的字段只有命名为pid的字段，该字段会生成逻辑验证代码
	*/
	IsResetLogic bool       `json:"isResetLogic"`
	SceneInfo    gdb.Record //场景信息
}

// 生成代码
func Run(ctx context.Context, parser *gcmd.Parser) {
	option := createOption(ctx, parser)
	tpl := createTpl(ctx, option.DbGroup, option.DbTable, option.RemovePrefixCommon, option.RemovePrefixAlone, true)

	genDao(tpl)           // dao模板生成
	genLogic(option, tpl) // logic模板生成

	if option.IsApi {
		genApi(option, tpl)         // api模板生成
		genController(option, tpl)  // controller模板生成
		genRouter(option, tpl)      // 后端路由生成
		genAction(ctx, option, tpl) // 操作权限生成
	}

	if option.IsView {
		genViewIndex(option, tpl)  // 视图模板Index生成
		genViewList(option, tpl)   // 视图模板List生成
		genViewQuery(option, tpl)  // 视图模板Query生成
		genViewSave(option, tpl)   // 视图模板Save生成
		genViewI18n(option, tpl)   // 视图模板I18n生成
		genViewRouter(option, tpl) // 前端路由生成
		genMenu(ctx, option, tpl)  // 菜单生成

		internal.Command(`前端代码格式化`, false, gfile.SelfDir()+`/../view/`+option.SceneCode, `npm`, `run`, `format`) // 前端代码格式化
	}
}

// 创建命令选项
func createOption(ctx context.Context, parser *gcmd.Parser) (option myGenOption) {
	optionMap := parser.GetOptAll()
	gconv.Struct(optionMap, &option)

	// 命令执行前提示搭配Git使用
	gcmd.Scan(
		color.HiYellowString(`重要提示：强烈建议搭配Git使用，防止代码覆盖风险。`)+"\n",
		color.HiYellowString(`    Git库未创建，请按`)+color.HiRedString(`[Ctrl + C]`)+color.HiYellowString(`中断执行`)+"\n",
		color.HiYellowString(`    Git库已创建或忽略风险，请按`)+color.HiGreenString(`[Enter]`)+color.HiYellowString(`继续执行`)+"\n",
	)

	// 场景标识
	if option.SceneCode == `` {
		option.SceneCode = gcmd.Scan(color.BlueString(`> 请输入场景标识：`))
	}
	for {
		if option.SceneCode != `` {
			option.SceneInfo, _ = daoAuth.Scene.CtxDaoModel(ctx).Filter(daoAuth.Scene.Columns().SceneCode, option.SceneCode).One()
			if !option.SceneInfo.IsEmpty() {
				break
			}
		}
		option.SceneCode = gcmd.Scan(color.RedString(`    场景标识不存在，请重新输入：`))
	}
	// db分组
	if option.DbGroup == `` {
		option.DbGroup = gcmd.Scan(color.BlueString(`> 请输入db分组，默认(default)：`))
		if option.DbGroup == `` {
			option.DbGroup = `default`
		}
	}
	for {
		err := g.Try(ctx, func(ctx context.Context) {
			g.DB(option.DbGroup)
		})
		if err == nil {
			break
		}
		option.DbGroup = gcmd.Scan(color.RedString(`    db分组不存在，请重新输入，默认(default)：`))
		if option.DbGroup == `` {
			option.DbGroup = `default`
		}
	}
	// db表
	tableArr, _ := g.DB(option.DbGroup).Tables(ctx)
	if option.DbTable == `` {
		option.DbTable = gcmd.Scan(color.BlueString(`> 请输入db表：`))
	}
	for {
		if option.DbTable != `` && garray.NewStrArrayFrom(tableArr).Contains(option.DbTable) {
			break
		}
		option.DbTable = gcmd.Scan(color.RedString(`    db表不存在，请重新输入：`))
	}
	// 要删除的共有前缀
	if _, ok := optionMap[`removePrefixCommon`]; !ok {
		option.RemovePrefixCommon = gcmd.Scan(color.BlueString(`> 请输入要删除的共有前缀，默认(空)：`))
	}
	for {
		if option.RemovePrefixCommon == `` || gstr.Pos(option.DbTable, option.RemovePrefixCommon) == 0 {
			break
		}
		option.RemovePrefixCommon = gcmd.Scan(color.RedString(`    要删除的共有前缀不存在，请重新输入，默认(空)：`))
	}
	// 要删除的独有前缀
	if _, ok := optionMap[`removePrefixAlone`]; !ok {
		option.RemovePrefixAlone = gcmd.Scan(color.BlueString(`> 请输入要删除的独有前缀，默认(空)：`))
	}
	for {
		if option.RemovePrefixAlone == `` || gstr.Pos(option.DbTable, option.RemovePrefixCommon+option.RemovePrefixAlone) == 0 {
			break
		}
		option.RemovePrefixAlone = gcmd.Scan(color.RedString(`    要删除的独有前缀不存在，请重新输入，默认(空)：`))
	}
	// 公共名称，将同时在swagger文档Tag标签，权限菜单和权限操作中使用。示例：场景
	for {
		if option.CommonName != `` {
			break
		}
		option.CommonName = gcmd.Scan(color.BlueString(`> 请输入公共名称，将同时在swagger文档Tag标签，权限菜单和权限操作中使用：`))
	}
noAllRestart:
	// 是否生成列表接口
	isList, ok := optionMap[`isList`]
	if !ok {
		isList = gcmd.Scan(color.BlueString(`> 是否生成列表接口，默认(yes)：`))
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
			isList = gcmd.Scan(color.RedString(`    输入错误，请重新输入，是否生成列表接口，默认(yes)：`))
		}
	}
	// 列表接口是否返回总数
	isCount, ok := optionMap[`isCount`]
	if !ok {
		isCount = gcmd.Scan(color.BlueString(`> 列表接口是否返回总数，默认(yes)：`))
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
			isCount = gcmd.Scan(color.RedString(`    输入错误，请重新输入，列表接口是否返回总数，默认(yes)：`))
		}
	}
	// 是否生成详情接口
	isInfo, ok := optionMap[`isInfo`]
	if !ok {
		isInfo = gcmd.Scan(color.BlueString(`> 是否生成详情接口，默认(yes)：`))
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
			isInfo = gcmd.Scan(color.RedString(`    输入错误，请重新输入，是否生成详情接口，默认(yes)：`))
		}
	}
	// 是否生成创建接口
	isCreate, ok := optionMap[`isCreate`]
	if !ok {
		isCreate = gcmd.Scan(color.BlueString(`> 是否生成创建接口，默认(yes)：`))
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
			isCreate = gcmd.Scan(color.RedString(`    输入错误，请重新输入，是否生成创建接口，默认(yes)：`))
		}
	}
	// 是否生成更新接口
	isUpdate, ok := optionMap[`isUpdate`]
	if !ok {
		isUpdate = gcmd.Scan(color.BlueString(`> 是否生成更新接口，默认(yes)：`))
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
			isUpdate = gcmd.Scan(color.RedString(`    输入错误，请重新输入，是否生成更新接口，默认(yes)：`))
		}
	}
	// 是否生成删除接口
	isDelete, ok := optionMap[`isDelete`]
	if !ok {
		isDelete = gcmd.Scan(color.BlueString(`> 是否生成删除接口，默认(yes)：`))
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
			isDelete = gcmd.Scan(color.RedString(`    输入错误，请重新输入，是否生成删除接口，默认(yes)：`))
		}
	}
	if !(option.IsList || option.IsInfo || option.IsCreate || option.IsUpdate || option.IsDelete) {
		fmt.Println(`请重新选择生成哪些接口，不能全是no！`)
		goto noAllRestart
	}
	// 是否生成后端接口文件
	isApi, ok := optionMap[`isApi`]
	if !ok {
		isApi = gcmd.Scan(color.BlueString(`> 是否生成后端接口文件，默认(yes)：`))
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
			isApi = gcmd.Scan(color.RedString(`    输入错误，请重新输入，是否生成后端接口文件，默认(yes)：`))
		}
	}
	if option.IsApi {
		// 是否判断操作权限，如是，则同时会生成操作权限
		isAuthAction, ok := optionMap[`isAuthAction`]
		if !ok {
			isAuthAction = gcmd.Scan(color.BlueString(`> 是否判断操作权限，如是，则同时会生成操作权限，默认(yes)：`))
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
				isAuthAction = gcmd.Scan(color.RedString(`    输入错误，请重新输入，是否判断操作权限，如是，则同时会生成操作权限，默认(yes)：`))
			}
		}
	}
	// 是否生成前端视图文件
	isView, ok := optionMap[`isView`]
	if !ok {
		isView = gcmd.Scan(color.BlueString(`> 是否生成前端视图文件，默认(yes)：`))
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
			isView = gcmd.Scan(color.RedString(`    输入错误，请重新输入，是否生成前端视图文件，默认(yes)：`))
		}
	}
	/*
		是否重置logic层。一般情况下不建议重置，原因：logic层生成基本不会有任何变化，且常会在该层手写一些逻辑验证和自定义方法。只建议在以下两种情况下重置：
			1、logic层生成模板发生重大变化，即对gen_logic.go中生成的模板代码做修改
			2、表新增或删除了对logic层生成代码有影响的字段。目前有影响的字段只有命名为pid的字段，该字段会生成逻辑验证代码
	*/
	isResetLogic, ok := optionMap[`isResetLogic`]
	if !ok {
		isResetLogic = gcmd.Scan(
			color.HiYellowString(`提示：是否重置logic层，一般情况下不建议重置，原因：logic层生成基本不会有任何变化，且常会在该层手写一些逻辑验证和自定义方法。只建议在以下两种情况下重置：`)+"\n",
			color.HiYellowString(`    1、logic层生成模板发生重大变化，即对gen_logic.go中生成的模板代码做修改`)+"\n",
			color.HiYellowString(`    2、表新增或删除了对logic层生成代码有影响的字段。目前有影响的字段只有命名为pid的字段，该字段会生成逻辑验证代码`)+"\n",
			color.BlueString(`> 是否重置logic层，默认(no)：`),
		)
	}
isResetLogicEnd:
	for {
		switch isResetLogic {
		case `1`, `yes`:
			option.IsResetLogic = true
			break isResetLogicEnd
		case ``, `0`, `no`:
			option.IsResetLogic = false
			break isResetLogicEnd
		default:
			isResetLogic = gcmd.Scan(color.RedString(`    输入错误，请重新输入，是否重置logic层，默认(no)：`))
		}
	}
	return
}
