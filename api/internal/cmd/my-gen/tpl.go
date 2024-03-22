package my_gen

import (
	"context"
	"math"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type myGenTpl struct {
	Link                      string                     //当前数据库连接配置（gf gen dao命令生成dao需要）
	TableArr                  []string                   //当前数据库全部数据表（获取扩展表，中间表等需要）
	Group                     string                     //数据库分组
	RemovePrefixCommon        string                     //要删除的共有前缀
	RemovePrefixAlone         string                     //要删除的独有前缀
	RemovePrefix              string                     //要删除的前缀
	TableType                 myGenTableType             //表类型。按该字段区分哪种功能表
	Table                     string                     //表名（原始，包含前缀）
	TableCaseSnake            string                     //表名（蛇形，已去除前缀）
	TableCaseCamel            string                     //表名（大驼峰，已去除前缀）
	TableCaseKebab            string                     //表名（横线，已去除前缀）
	KeyList                   []myGenKey                 //索引列表
	FieldListRaw              map[string]*gdb.TableField //字段列表（原始）
	FieldList                 []myGenField               //字段列表
	ModuleDirCaseCamel        string                     //模块目录（大驼峰，/会被去除）
	ModuleDirCaseKebab        string                     //模块目录（横线，/会被保留）
	ModuleDirCaseKebabReplace string                     //模块目录（横线，/被替换成.）
	LogicStructName           string                     //logic层结构体名称，也是权限操作前缀（大驼峰，由ModuleDirCaseCamel+TableCaseCamel组成。命名原因：gf gen service只支持logic单层目录，可能导致service层重名）
	Handle                    struct {                   //需特殊处理的字段
		Id struct { //主键列表（无主键时，默认第一个字段）。联合主键有多字段，需按顺序存入
			IsPrimary bool //是否主键
			List      []myGenField
			Type      myGenFieldType
		}
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
		RelIdMap            map[string]handleRelId //id后缀字段，需特殊处理
		ExtendTableOneList  []myGenTpl             //扩展表（一对一）：表命名：主表名_xxxx，并存在与主表主键同名的字段，且字段设为不递增主键或唯一索引
		ExtendTableManyList []myGenTpl             //扩展表（一对多）：表命名：主表名_xxxx，并存在与主表主键同名的字段，且字段设为普通索引
		MiddleTableOneList  []myGenTpl             //中间表（一对一）：表命名使用_rel_to_或_rel_of_关联两表，不同模块两表必须全名，同模块第二个表可全名也可省略前缀。存在与两个关联表主键同名的字段，用_rel_to_做关联时，第一个表的关联字段做主键或唯一索引，用_rel_of_做关联时，第二个表的关联字段做主键或唯一索引。
		MiddleTableManyList []myGenTpl             //中间表（一对多）：表命名使用_rel_to_或_rel_of_关联两表，不同模块两表必须全名，同模块第二个表可全名也可省略前缀。存在与两个关联表主键同名的字段，两关联字段做联合主键或联合唯一索引
	}
}

type myGenTableType = uint
type myGenFieldType = int
type myGenFieldTypeName = string

const (
	TableTypeGen    myGenTableType = 0  //生成表
	TableTypeExtend myGenTableType = 1  //扩展表
	TableTypeMiddle myGenTableType = 2  //中间表
	TableTypeRelId  myGenTableType = 10 //id后缀关联表

	//用于结构体中，需从1开始，否则结构体会默认0
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

	TypeNamePri            myGenFieldTypeName = `主键（非联合）`
	TypeNamePriAutoInc     myGenFieldTypeName = `自增主键（非联合）`
	TypeNameDeleted        myGenFieldTypeName = `软删除字段`
	TypeNameUpdated        myGenFieldTypeName = `更新时间字段`
	TypeNameCreated        myGenFieldTypeName = `创建时间字段`
	TypeNamePid            myGenFieldTypeName = `命名：pid；	类型：int等类型；`
	TypeNameLevel          myGenFieldTypeName = `命名：level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；`
	TypeNameIdPath         myGenFieldTypeName = `命名：idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；`
	TypeNameSort           myGenFieldTypeName = `命名：sort，且pid,level,idPath|id_path,sort同时存在时（才）有效；	类型：int等类型；`
	TypeNamePasswordSuffix myGenFieldTypeName = `命名：password,passwd后缀；		类型：char(32)；`
	TypeNameSaltSuffix     myGenFieldTypeName = `命名：salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；`
	TypeNameNameSuffix     myGenFieldTypeName = `命名：name,title后缀；	类型：varchar；`
	TypeNameCodeSuffix     myGenFieldTypeName = `命名：code后缀；	类型：varchar；`
	TypeNameAccountSuffix  myGenFieldTypeName = `命名：account后缀；	类型：varchar；`
	TypeNamePhoneSuffix    myGenFieldTypeName = `命名：phone,mobile后缀；	类型：varchar；`
	TypeNameEmailSuffix    myGenFieldTypeName = `命名：email后缀；	类型：varchar；`
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
	FieldRaw             string             // 字段（原始）
	FieldCaseSnake       string             // 字段（蛇形）
	FieldCaseCamel       string             // 字段（大驼峰）
	FieldCaseSnakeRemove string             // 字段（蛇形。去除_of_后）
	FieldCaseCamelRemove string             // 字段（大驼峰。去除_of_后）
	FieldTypeRaw         string             // 字段类型（原始）
	FieldType            myGenFieldType     // 字段类型（数据类型）
	FieldTypeName        myGenFieldTypeName // 字段类型（命名类型）
	IsUnique             bool               // 是否独立的唯一索引
	IsNull               bool               // 字段是否可为NULL
	Default              interface{}        // 默认值
	Comment              string             // 注释（原始）。
	FieldName            string             // 字段名称。由注释解析出来，前端显示用。符号[\n\r.。:：(（]之前的部分或整个注释，将作为字段名称使用）
	FieldDesc            string             // 字段说明。由注释解析出来，API文档用。符号[\n\r]换成` `，"增加转义换成\"
	FieldTip             string             // 字段提示。由注释解析出来，前端提示用。
	StatusList           [][2]string        // 状态列表。由注释解析出来，前端显示用。多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
	FieldLimitStr        string             // 字符串字段限制。varchar表示最大长度；char表示长度；
	FieldLimitFloat      [2]string          // 浮点数字段限制。第1个表示整数位，第2个表示小数位
	FieldShowLenMax      int                // 显示长度。公式：汉字个数 + (其它字符个数 / 2)。前端el-select-v2等部分组件生成时，根据该值设置宽度
}

type myGenKey struct {
	Name      string   // 索引名称。主键：PRIMARY；其它：定义
	Index     uint     // 索引顺序。从1开始，单索引都是1，联合索引按字段数量顺序递增
	Field     string   // 字段（原始）
	FieldArr  []string // 字段列表。联合索引有多字段，需按顺序存入
	IsPrimary bool     // 是否主键
	IsUnique  bool     // 是否唯一
	IsAutoInc bool     // 是否自增
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

// 创建模板参数
func createTpl(ctx context.Context, group, table, removePrefixCommon, removePrefixAlone string, tableTypeOpt ...myGenTableType) (tpl myGenTpl) {
	tpl = myGenTpl{
		Group:              group,
		RemovePrefixCommon: removePrefixCommon,
		RemovePrefixAlone:  removePrefixAlone,
		RemovePrefix:       removePrefixCommon + removePrefixAlone,
		Table:              table,
	}
	if len(tableTypeOpt) > 0 {
		tpl.TableType = tableTypeOpt[0]
	}
	tpl.Link = gconv.String(gconv.SliceMap(g.Cfg().MustGet(ctx, `database`).MapStrAny()[tpl.Group])[0][`link`])
	tpl.TableArr = tpl.getTable(ctx, tpl.Group)
	tpl.KeyList = tpl.getTableKey(ctx, tpl.Group, tpl.Table)
	tpl.FieldListRaw = tpl.getTableField(ctx, tpl.Group, tpl.Table)
	tpl.TableCaseSnake = gstr.CaseSnake(gstr.Replace(tpl.Table, tpl.RemovePrefix, ``, 1))
	tpl.TableCaseCamel = gstr.CaseCamel(tpl.TableCaseSnake)
	tpl.TableCaseKebab = gstr.CaseKebab(tpl.TableCaseSnake)
	tpl.Handle.PasswordMap = map[string]handlePassword{}
	tpl.Handle.RelIdMap = map[string]handleRelId{}
	logicStructName := gstr.TrimLeftStr(tpl.Table, tpl.RemovePrefixCommon, 1)
	moduleDirCaseCamel := gstr.CaseCamel(logicStructName)
	moduleDirCaseKebab := gstr.CaseKebab(logicStructName)
	if tpl.RemovePrefixAlone != `` {
		moduleDirCaseCamel = gstr.CaseCamel(tpl.RemovePrefixAlone)
		moduleDirCaseKebab = gstr.CaseKebab(gstr.Trim(tpl.RemovePrefixAlone, `_`))
	}
	if tpl.Group != `default` {
		logicStructName = tpl.Group + `_` + logicStructName
		moduleDirCaseCamel = gstr.CaseCamel(tpl.Group) + moduleDirCaseCamel
		moduleDirCaseKebab = gstr.CaseKebab(tpl.Group) + `/` + moduleDirCaseKebab
	}
	tpl.LogicStructName = gstr.CaseCamel(logicStructName)
	tpl.ModuleDirCaseKebab = moduleDirCaseKebab
	tpl.ModuleDirCaseKebabReplace = gstr.Replace(moduleDirCaseKebab, `/`, `.`)
	tpl.ModuleDirCaseCamel = moduleDirCaseCamel

	fieldList := make([]myGenField, len(tpl.FieldListRaw))
	for _, v := range tpl.FieldListRaw {
		fieldTmp := myGenField{
			FieldRaw:     v.Name,
			FieldTypeRaw: v.Type,
			IsNull:       v.Null,
			Default:      v.Default,
			Comment:      v.Comment,
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
		for _, v := range []string{`.`, `。`, `:`, `：`, ` `, `,`, `，`, `;`, `；`} {
			tmpFieldTip = gstr.Trim(tmpFieldTip, v)
		}
		if gstr.Pos(tmpFieldTip, `(`) == 0 {
			tmpFieldTip = gstr.TrimRightStr(gstr.TrimLeftStr(tmpFieldTip, `(`, 1), `)`, 1)
		}
		if gstr.Pos(tmpFieldTip, `（`) == 0 {
			tmpFieldTip = gstr.TrimRightStr(gstr.TrimLeftStr(tmpFieldTip, `（`, 1), `）`, 1)
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

		fieldTmp.FieldShowLenMax = tpl.getShowLen(fieldTmp.FieldName)

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
		//先确定主键（非联合主键）
		for _, key := range tpl.KeyList {
			if fieldTmp.FieldRaw != key.Field {
				continue
			}
			if key.IsUnique && len(key.FieldArr) == 1 {
				fieldTmp.IsUnique = true
			}
			if key.IsPrimary && len(key.FieldArr) == 1 {
				fieldTmp.FieldTypeName = TypeNamePri
				if key.IsAutoInc {
					fieldTmp.FieldTypeName = TypeNamePriAutoInc
				}
			}
		}
		if garray.NewStrArrayFrom([]string{TypeNamePriAutoInc, TypeNamePri}).Contains(fieldTmp.FieldTypeName) {
		} else if garray.NewStrArrayFrom([]string{`DeletedAt`, `DeleteAt`, `DeletedTime`, `DeleteTime`}).Contains(fieldTmp.FieldCaseCamel) {
			fieldTmp.FieldTypeName = TypeNameDeleted
		} else if garray.NewStrArrayFrom([]string{`UpdatedAt`, `UpdateAt`, `UpdatedTime`, `UpdateTime`}).Contains(fieldTmp.FieldCaseCamel) {
			fieldTmp.FieldTypeName = TypeNameUpdated
		} else if garray.NewStrArrayFrom([]string{`CreatedAt`, `CreateAt`, `CreatedTime`, `CreateTime`}).Contains(fieldTmp.FieldCaseCamel) {
			fieldTmp.FieldTypeName = TypeNameCreated
		} else if garray.NewIntArrayFrom([]int{TypeVarchar, TypeText}).Contains(fieldTmp.FieldType) && fieldTmp.FieldCaseCamel == `IdPath` { //idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效
			fieldTmp.FieldTypeName = TypeNameIdPath

			tpl.Handle.Pid.IdPath = fieldTmp.FieldRaw
		} else if garray.NewIntArrayFrom([]int{TypeInt, TypeIntU, TypeVarchar, TypeChar}).Contains(fieldTmp.FieldType) && garray.NewStrArrayFrom([]string{`status`, `type`, `method`, `pos`, `position`, `gender`}).Contains(fieldSuffix) { //status,type,method,pos,position,gender等后缀
			fieldTmp.FieldTypeName = TypeNameStatusSuffix

			isStr := false
			if garray.NewIntArrayFrom([]int{TypeVarchar, TypeChar}).Contains(fieldTmp.FieldType) {
				isStr = true
			}
			fieldTmp.StatusList = tpl.getStatusList(fieldTmp.FieldTip, isStr)

			for _, status := range fieldTmp.StatusList {
				showLen := tpl.getShowLen(status[1])
				if showLen > fieldTmp.FieldShowLenMax {
					fieldTmp.FieldShowLenMax = showLen
				}
			}
		} else if garray.NewIntArrayFrom([]int{TypeVarchar, TypeText, TypeJson}).Contains(fieldTmp.FieldType) && (garray.NewStrArrayFrom([]string{`icon`, `cover`, `avatar`, `img`, `image`}).Contains(fieldSuffix) || gstr.SubStr(fieldTmp.FieldCaseCamelRemove, -7) == `ImgList` || gstr.SubStr(fieldTmp.FieldCaseCamelRemove, -6) == `ImgArr` || gstr.SubStr(fieldTmp.FieldCaseCamelRemove, -9) == `ImageList` || gstr.SubStr(fieldTmp.FieldCaseCamelRemove, -8) == `ImageArr`) { //icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀
			fieldTmp.FieldTypeName = TypeNameImageSuffix
		} else if garray.NewIntArrayFrom([]int{TypeVarchar, TypeText, TypeJson}).Contains(fieldTmp.FieldType) && (garray.NewStrArrayFrom([]string{`video`}).Contains(fieldSuffix) || gstr.SubStr(fieldTmp.FieldCaseCamelRemove, -9) == `VideoList` || gstr.SubStr(fieldTmp.FieldCaseCamelRemove, -8) == `VideoArr`) { //video,video_list,videoList,video_arr,videoArr等后缀
			fieldTmp.FieldTypeName = TypeNameVideoSuffix
		} else if garray.NewIntArrayFrom([]int{TypeText, TypeJson}).Contains(fieldTmp.FieldType) && garray.NewStrArrayFrom([]string{`list`, `arr`}).Contains(fieldSuffix) { //list,arr等后缀
			fieldTmp.FieldTypeName = TypeNameArrSuffix
		} else if garray.NewIntArrayFrom([]int{TypeVarchar, TypeText}).Contains(fieldTmp.FieldType) && garray.NewStrArrayFrom([]string{`remark`, `desc`, `msg`, `message`, `intro`, `content`}).Contains(fieldSuffix) { //remark,desc,msg,message,intro,content后缀
			fieldTmp.FieldTypeName = TypeNameRemarkSuffix
		} else if fieldTmp.FieldType == TypeVarchar { //varchar类型
			if garray.NewStrArrayFrom([]string{`name`, `title`}).Contains(fieldSuffix) { //name,title后缀
				fieldTmp.FieldTypeName = TypeNameNameSuffix
			} else if garray.NewStrArrayFrom([]string{`code`}).Contains(fieldSuffix) { //code后缀
				fieldTmp.FieldTypeName = TypeNameCodeSuffix
			} else if garray.NewStrArrayFrom([]string{`account`}).Contains(fieldSuffix) { //account后缀
				fieldTmp.FieldTypeName = TypeNameAccountSuffix
			} else if garray.NewStrArrayFrom([]string{`phone`, `mobile`}).Contains(fieldSuffix) { //phone,mobile后缀
				fieldTmp.FieldTypeName = TypeNamePhoneSuffix
			} else if garray.NewStrArrayFrom([]string{`email`}).Contains(fieldSuffix) { //email后缀
				fieldTmp.FieldTypeName = TypeNameEmailSuffix
			} else if garray.NewStrArrayFrom([]string{`url`, `link`}).Contains(fieldSuffix) { //url,link后缀
				fieldTmp.FieldTypeName = TypeNameUrlSuffix
			} else if garray.NewStrArrayFrom([]string{`ip`}).Contains(fieldSuffix) { //IP后缀
				fieldTmp.FieldTypeName = TypeNameIpSuffix
			}
		} else if fieldTmp.FieldType == TypeChar { //char类型
			if garray.NewStrArrayFrom([]string{`password`, `passwd`}).Contains(fieldSuffix) && fieldTmp.FieldTypeRaw == `char(32)` { //password,passwd后缀
				fieldTmp.FieldTypeName = TypeNamePasswordSuffix

				passwordMapKey := tpl.getHandlePasswordMapKey(fieldTmp.FieldRaw)
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

				passwordMapKey := tpl.getHandlePasswordMapKey(fieldTmp.FieldRaw)
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
		} else if garray.NewIntArrayFrom([]int{TypeInt, TypeIntU}).Contains(fieldTmp.FieldType) { //int等类型
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
					tpl:       tpl.getRelIdTpl(ctx, tpl, fieldTmp.FieldRaw),
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
				// TODO 可改成状态一样处理，同时需要修改前端开关组件属性设置（暂时不改）
			}
		} else if garray.NewIntArrayFrom([]int{TypeTimestamp, TypeDatetime, TypeDate}).Contains(fieldTmp.FieldType) { //timestamp或datetime或date类型
			if garray.NewStrArrayFrom([]string{`start`}).Contains(fieldPrefix) { //start_前缀
				fieldTmp.FieldTypeName = TypeNameStartPrefix
			} else if garray.NewStrArrayFrom([]string{`end`}).Contains(fieldPrefix) { //end_前缀
				fieldTmp.FieldTypeName = TypeNameEndPrefix
			}
		}
		/*--------确定字段命名类型（部分命名类型需做二次确定） 结束--------*/

		fieldList[v.Index] = fieldTmp
	}

	/*--------解析影响命名类型二次确认，且需特殊处理的字段 开始--------*/
	//password|passwd,salt同时存在时，需特殊处理
	for k, v := range tpl.Handle.PasswordMap {
		if v.PasswordField != `` && v.SaltField != `` {
			v.IsCoexist = true
			tpl.Handle.PasswordMap[k] = v
		}
	}

	//pid,level,idPath|id_path同时存在时，需特殊处理
	if garray.NewFrom([]interface{}{TableTypeGen}).Contains(tpl.TableType) {
		if tpl.Handle.Pid.Pid != `` && tpl.Handle.Pid.Level != `` && tpl.Handle.Pid.IdPath != `` {
			tpl.Handle.Pid.IsCoexist = true
		}
	}
	/*--------解析影响命名类型二次确认，且需特殊处理的字段 结束--------*/

	/*--------命名类型二次确认的字段 开始--------*/
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
			passwordMapKey := tpl.getHandlePasswordMapKey(v.FieldRaw)
			if !tpl.Handle.PasswordMap[passwordMapKey].IsCoexist {
				fieldList[k].FieldTypeName = ``
			}
		}
	}
	/*--------命名类型二次确认的字段 结束--------*/

	/*--------需特殊处理的字段解析 开始--------*/
	//主键列表（无主键时，默认第一个字段）。联合主键有多字段，需按顺序存入
	for _, key := range tpl.KeyList {
		if !key.IsPrimary {
			continue
		}
		tpl.Handle.Id.IsPrimary = true
		for _, field := range key.FieldArr {
			for _, v := range fieldList {
				if v.FieldRaw == field {
					tpl.Handle.Id.List = append(tpl.Handle.Id.List, v)
					break
				}
			}
		}
		break
	}
	if len(tpl.Handle.Id.List) == 0 {
		tpl.Handle.Id.List = append(tpl.Handle.Id.List, fieldList[0])
	}
	if len(tpl.Handle.Id.List) == 1 {
		tpl.Handle.Id.Type = tpl.Handle.Id.List[0].FieldType
	}
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
		if len(tpl.Handle.Id.List) == 1 && tpl.Handle.Id.IsPrimary {
			fieldSplitArr := gstr.Split(tpl.Handle.Id.List[0].FieldCaseSnake, `_`)
			fieldSuffix := fieldSplitArr[len(fieldSplitArr)-1]
			if fieldSuffix == `id` {
				labelTmp1 := gstr.SubStr(tpl.Handle.Id.List[0].FieldCaseCamel, 0, -2) + v
				if labelTmp1 != labelTmp && labelTmp1 != v {
					labelList = append(labelList, labelTmp1)
				}
			}
		}
		labelList = append(labelList, v)
	}
	for _, v := range labelList {
		for _, item := range fieldList {
			if v == item.FieldCaseCamel && garray.NewIntArrayFrom([]int{TypeVarchar, TypeChar}).Contains(item.FieldType) {
				tpl.Handle.LabelList = append(tpl.Handle.LabelList, item.FieldRaw)
				break
			}
		}
	}

	//id后缀字段
	for k, v := range tpl.Handle.RelIdMap {
		if len(v.tpl.Handle.LabelList) == 0 {
			continue
		}
		for _, item := range fieldList {
			if item.FieldRaw == v.tpl.Handle.LabelList[0]+v.Suffix {
				v.IsRedundName = true
				tpl.Handle.RelIdMap[k] = v
				break
			}
		}
	}

	//扩展表
	if garray.NewFrom([]interface{}{TableTypeGen}).Contains(tpl.TableType) {
		tpl.Handle.ExtendTableOneList, tpl.Handle.ExtendTableManyList = tpl.getExtendTable(ctx, tpl)
	}
	/*--------需特殊处理的字段解析 结束--------*/

	tpl.FieldList = fieldList
	return
}

// 获取表
func (myGenTplThis *myGenTpl) getTable(ctx context.Context, group string) (tableArr []string) {
	tableArr, _ = g.DB(group).Tables(ctx) //框架只带的。应该兼容多种数据库吧
	return
}

// 获取表字段
func (myGenTplThis *myGenTpl) getTableField(ctx context.Context, group, table string) (fieldList map[string]*gdb.TableField) {
	// fieldList, _ = g.DB(group).GetAll(ctx, `SHOW FULL COLUMNS FROM `+table)	// 除mysql外，其它数据库不兼容该sql语句
	fieldList, _ = g.DB(group).TableFields(ctx, table) //框架只带的。应该兼容多种数据库吧
	return
}

// TODO 获取表索引（目前只支持mysql解析，故其它sql数据库，不支持扩展表和中间表关联生成）
func (myGenTplThis *myGenTpl) getTableKey(ctx context.Context, group, table string) (keyList []myGenKey) {
	switch g.DB(group).GetConfig().Type {
	case `mysql`:
		keyListTmp, _ := g.DB(group).GetAll(ctx, `SHOW KEYS FROM `+table)
		keyList = make([]myGenKey, len(keyListTmp))
		fieldList := myGenTplThis.getTableField(ctx, group, table)
		fieldArrMap := map[string][]string{}
		for k, v := range keyListTmp {
			key := myGenKey{
				Name:     v[`Key_name`].String(),
				Index:    v[`Seq_in_index`].Uint(),
				Field:    v[`Column_name`].String(),
				IsUnique: !v[`Non_unique`].Bool(),
			}
			if key.Name == `PRIMARY` {
				key.IsPrimary = true
				for _, field := range fieldList {
					if key.Field == field.Name && field.Extra == `auto_increment` {
						key.IsAutoInc = true
						break
					}
				}
			}
			keyList[k] = key

			fieldArrMap[key.Name] = append(fieldArrMap[key.Name], key.Field)
		}
		for k, v := range keyList {
			v.FieldArr = fieldArrMap[v.Name]
			keyList[k] = v
		}
	case `sqlite`:
	case `mssql`:
	case `pgsql`:
	case `oracle`:
	}
	return
}

// 执行gf gen dao命令生成dao文件
func (myGenTplThis *myGenTpl) gfGenDao(isOverwriteDao bool) {
	commandArg := []string{
		`gen`, `dao`,
		`--link`, myGenTplThis.Link,
		`--group`, myGenTplThis.Group,
		`--removePrefix`, myGenTplThis.RemovePrefix,
		`--daoPath`, `dao/` + myGenTplThis.ModuleDirCaseKebab,
		`--doPath`, `model/entity/` + myGenTplThis.ModuleDirCaseKebab,
		`--entityPath`, `model/entity/` + myGenTplThis.ModuleDirCaseKebab,
		`--tables`, myGenTplThis.Table,
		`--tplDaoIndexPath`, `resource/gen/gen_dao_template_dao.txt`,
		`--tplDaoInternalPath`, `resource/gen/gen_dao_template_dao_internal.txt`,
	}
	if isOverwriteDao {
		commandArg = append(commandArg, `--overwriteDao=true`)
	}
	command(`表（`+myGenTplThis.Table+`）dao生成`, true, ``, `gf`, commandArg...)
}

// status字段注释解析
func (myGenTplThis *myGenTpl) getStatusList(comment string, isStr bool) (statusList [][2]string) {
	var tmp [][]string
	if isStr {
		tmp, _ = gregex.MatchAllString(`([A-Za-z0-9]+)[-=:：]?([^\s,，;；]+)`, comment)
	} else {
		// tmp, _ = gregex.MatchAllString(`(-?\d+)[-=:：]?([^\d\s,，;；]+)`, comment)
		tmp, _ = gregex.MatchAllString(`(-?\d+)[-=:：]?([^\s,，;；]+)`, comment)
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

// 获取显示长度。汉字个数 + (其它字符个数 / 2) 后的值
func (myGenTplThis *myGenTpl) getShowLen(str string) int {
	len := len(str)
	lenRune := gstr.LenRune(str)
	countHan := (len - lenRune) / 2
	countOther := gconv.Int(math.Ceil(float64(len-countHan*3) / 2))
	return countHan + countOther
}

// 获取Handle.PasswordMap的Key（以Password为主）
func (myGenTplThis *myGenTpl) getHandlePasswordMapKey(passwordOrsalt string) (passwordMapKey string) {
	passwordOrsalt = gstr.Replace(gstr.CaseCamel(passwordOrsalt), `Salt`, `Password`, 1) //替换salt
	passwordOrsalt = gstr.Replace(passwordOrsalt, `Passwd`, `Password`, 1)               //替换passwd
	passwordMapKey = gstr.CaseCamelLower(passwordOrsalt)                                 //默认：小驼峰
	if gstr.CaseCamelLower(passwordOrsalt) != passwordOrsalt {                           //判断字段是不是蛇形
		passwordMapKey = gstr.CaseSnake(passwordMapKey)
	}
	return
}

// 获取id后缀字段关联的表信息
func (myGenTplThis *myGenTpl) getRelIdTpl(ctx context.Context, tpl myGenTpl, field string) (relTpl myGenTpl) {
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
	removePrefixAloneTmp := tpl.RemovePrefixAlone //moduleDir
	if removePrefixAloneTmp == `` {               //同模块当主表是user,good等无下划线时，找同模块关联表时，表前缀为：当前主表 + `_`
		removePrefixAloneTmp = gstr.TrimLeftStr(tpl.Table, tpl.RemovePrefixCommon, 1) + `_`
	}
	isSamePrimaryFunc := func(table string) bool {
		tableKeyList := tpl.getTableKey(ctx, tpl.Group, table)
		for _, v := range tableKeyList {
			if v.IsPrimary && len(v.FieldArr) == 1 && garray.NewStrArrayFrom([]string{`id`, fieldCaseSnakeOfRemove}).Contains(gstr.CaseSnake(v.Field)) {
				return true
			}
		}
		return false
	}
	for _, v := range tpl.TableArr {
		if v == tpl.Table { //自身跳过
			continue
		}
		if gstr.Pos(v, `_rel_to_`) != -1 || gstr.Pos(v, `_rel_of_`) != -1 { //中间表跳过
			continue
		}
		if v == tpl.RemovePrefixCommon+removePrefixAloneTmp+tableSuffix { //关联表在同模块目录下，且表名一致
			if isSamePrimaryFunc(v) {
				mayBeObj.table1 = v
				break
			}
		} else if gstr.Pos(v, tpl.RemovePrefixCommon+removePrefixAloneTmp) == 0 && len(v) == gstr.PosR(v, `_`+tableSuffix)+len(`_`+tableSuffix) { //关联表在同模块目录下，但表后缀一致
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
			if pos := gstr.Pos(tableRemove, `_`); pos != -1 {
				removePrefixAlone = gstr.SubStr(tableRemove, 0, pos+1)
			}
		}

		relTpl = createTpl(ctx, tpl.Group, table, removePrefixCommon, removePrefixAlone, TableTypeRelId)
		relTpl.gfGenDao(false) //dao文件生成
	}
	return
}

// TODO 获取扩展表
func (myGenTplThis *myGenTpl) getExtendTable(ctx context.Context, tpl myGenTpl) (extendTableOneList []myGenTpl, extendTableManyList []myGenTpl) {
	if len(tpl.Handle.Id.List) > 1 || !tpl.Handle.Id.IsPrimary { //联合主键或无主键时，不获取扩展表
		return
	}
	removePrefixCommon := tpl.RemovePrefixCommon
	removePrefixAlone := tpl.RemovePrefixAlone
	if removePrefixAlone == `` {
		removePrefixAlone = gstr.TrimLeftStr(tpl.Table, removePrefixCommon, 1) + `_`
	}

	fieldPrimaryArr := []string{tpl.Handle.Id.List[0].FieldCaseSnake}
	if fieldPrimaryArr[0] == `id` {
		fieldPrimaryArr = append(fieldPrimaryArr, gstr.TrimLeftStr(gstr.TrimLeftStr(tpl.Table, tpl.RemovePrefixCommon, 1), tpl.RemovePrefixAlone, 1)+`_id`)
	}

	for _, v := range tpl.TableArr {
		if v == tpl.Table { //自身跳过
			continue
		}
		if gstr.Pos(v, `_rel_to_`) != -1 || gstr.Pos(v, `_rel_of_`) != -1 { //中间表跳过
			continue
		}
		if gstr.Pos(v, tpl.Table+`_`) != 0 { // 不符合表命名（主表名_xxxx）的跳过
			continue
		}
		extendTpl := createTpl(ctx, tpl.Group, v, removePrefixCommon, removePrefixAlone, TableTypeExtend)
		for _, key := range extendTpl.KeyList {
			if len(key.FieldArr) == 1 && garray.NewStrArrayFrom(fieldPrimaryArr).Contains(gstr.CaseSnake(key.Field)) {
				if key.IsPrimary {
					if !key.IsAutoInc { //非自增主键
						extendTpl.gfGenDao(false) //dao文件生成
						extendTableOneList = append(extendTableOneList, extendTpl)
					}
				} else if key.IsUnique {
					extendTpl.gfGenDao(false) //dao文件生成
					extendTableOneList = append(extendTableOneList, extendTpl)
				} else {
					extendTpl.gfGenDao(false) //dao文件生成
					extendTableManyList = append(tpl.Handle.ExtendTableOneList, extendTpl)
				}
			}
		}
	}
	return
}
