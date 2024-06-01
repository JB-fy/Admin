package my_gen

import (
	"api/internal/cmd/my-gen/internal"
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type myGenTpl struct {
	DbHandler          internal.MyGenDbHandler  //数据库处理器
	FieldStyle         internal.MyGenFieldStyle //表字段命名风格
	Link               string                   //当前数据库连接配置（gf gen dao命令生成dao需要）
	TableArr           []string                 //当前数据库全部数据表（获取扩展表，中间表等需要）
	Group              string                   //数据库分组
	RemovePrefixCommon string                   //要删除的共有前缀
	RemovePrefixAlone  string                   //要删除的独有前缀
	RemovePrefix       string                   //要删除的前缀
	Table              string                   //表名（原始，包含前缀）
	TableCaseSnake     string                   //表名（蛇形，已去除前缀）
	TableCaseCamel     string                   //表名（大驼峰，已去除前缀）
	TableCaseKebab     string                   //表名（横线，已去除前缀）
	KeyList            []internal.MyGenKey      //索引列表
	FieldList          []myGenField             //字段列表
	FieldListOfDefault []myGenField             //除FieldListOfAfter外，表的其它字段数组。
	FieldListOfAfter   []myGenField             //最后处理的字段数组。且需按该数组顺序处理
	ModuleDirCaseCamel string                   //模块目录（大驼峰，/会被去除）
	ModuleDirCaseKebab string                   //模块目录（横线，/会被保留）
	LogicStructName    string                   //logic层结构体名称，也是权限操作前缀（大驼峰，由ModuleDirCaseCamel+TableCaseCamel组成。命名原因：gf gen service只支持logic单层目录，可能导致service层重名）
	I18nPath           string                   //前端多语言使用
	Handle             struct {                 //需特殊处理的字段
		Id struct { //主键列表（无主键时，默认为排除internal.ConfigIdAndLabelExcField过后的第一个字段）。联合主键有多字段，需按顺序存入
			List      []myGenField
			IsPrimary bool //是否主键
		}
		/*
			label列表。sql查询可设为别名label的字段（常用于前端my-select或my-cascader等组件，或用于关联表查询）。按以下优先级存入：
				表名去掉前缀 + Name > 主键去掉ID + Name > Name >
				表名去掉前缀 + Title > 主键去掉ID + Title > Title >
				表名去掉前缀 + Phone > 主键去掉ID + Phone > Phone >
				表名去掉前缀 + Email > 主键去掉ID + Email > Email >
				表名去掉前缀 + Account > 主键去掉ID + Account > Account >
				表名去掉前缀 + Nickname > 主键去掉ID + Nickname > Nickname >
				上面字段都没有时，默认为排除internal.ConfigIdAndLabelExcField过后的第二个字段
		*/
		LabelList   []string
		PasswordMap map[string]handlePassword //password|passwd,salt同时存在时，需特殊处理
		Pid         struct {                  //pid,level,idPath|id_path同时存在时，需特殊处理
			IsCoexist bool     //是否同时存在pid,level,idPath|id_path
			Pid       string   //父级字段
			Level     string   //层级字段
			IdPath    string   //层级路径字段
			Sort      []string //排序字段列表（当有排序字段时，树状列表对这些字段做正序排序）
		}
		RelIdMap            map[string]handleRelId //id后缀字段，需特殊处理
		ExtendTableOneList  []handleExtendMiddle   //扩展表（一对一）：表命名：主表名_xxxx，并存在与主表主键同名的字段，且字段设为不递增主键或唯一索引
		ExtendTableManyList []handleExtendMiddle   //扩展表（一对多）：表命名：主表名_xxxx，并存在与主表主键同名的字段，且字段设为普通索引
		MiddleTableOneList  []handleExtendMiddle   //中间表（一对一）：表命名使用_rel_to_或_rel_of_关联两表，不同模块两表必须全名，同模块第二个表可全名也可省略前缀。存在与两个关联表主键同名的字段，用_rel_to_做关联时，第一个表的关联字段做主键或唯一索引，用_rel_of_做关联时，第二个表的关联字段做主键或唯一索引。
		MiddleTableManyList []handleExtendMiddle   //中间表（一对多）：表命名使用_rel_to_或_rel_of_关联两表，不同模块两表必须全名，同模块第二个表可全名也可省略前缀。存在与两个关联表主键同名的字段，两关联字段做联合主键或联合唯一索引
		OtherRelTableList   []handleOtherRel       //其它关联表（不含扩展表和中间表）
	}
}

type myGenField struct {
	internal.MyGenField
	IsUnique             bool                           // 是否独立的唯一索引
	FieldType            internal.MyGenFieldType        // 字段类型（数据类型）
	FieldTypePrimary     internal.MyGenFieldTypePrimary // 字段类型（主键类型）
	FieldTypeName        internal.MyGenFieldTypeName    // 字段类型（命名类型）
	FieldCaseSnake       string                         // 字段（蛇形）
	FieldCaseCamel       string                         // 字段（大驼峰）
	FieldCaseSnakeRemove string                         // 字段（蛇形。去除_of_后）
	FieldCaseCamelRemove string                         // 字段（大驼峰。去除_of_后）
	FieldName            string                         // 字段名称。由注释解析出来，前端显示用。符号[\n\r.。:：(（]之前的部分或整个注释，将作为字段名称使用）
	FieldDesc            string                         // 字段说明。由注释解析出来，API文档用。符号[\n\r]换成` `，"增加转义换成\"
	FieldTip             string                         // 字段提示。由注释解析出来，前端提示用。
	StatusList           [][2]string                    // 状态列表。由注释解析出来，前端显示用。多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
	FieldLimitStr        string                         // 字符串字段限制。varchar表示最大长度；char表示长度；
	FieldLimitInt        internal.MyGenFieldLimitInt    // 整数字段限制
	FieldLimitFloat      internal.MyGenFieldLimitFloat  // 浮点数字段限制
	FieldShowLenMax      int                            // 显示长度。公式：汉字个数 + (其它字符个数 / 2)。前端el-select-v2等部分组件生成时，根据该值设置宽度
}

type handlePassword struct {
	IsCoexist      bool   //是否同时存在
	PasswordField  string //密码字段
	PasswordLength string //密码字段长度
	SaltField      string //密码盐字段
	SaltLength     string //密码盐字段长度
}

type handleRelId struct {
	tpl             myGenTpl
	FieldName       string //字段名称
	IsRedundName    bool   //是否冗余过关联表名称字段
	Suffix          string //关联表字段后缀（原始，大驼峰或蛇形）。字段含[_of_]时，_of_及之后的部分。示例：userIdOfSend对应OfSend；user_id_of_send对应_of_send
	SuffixCaseCamel string //关联表字段后缀（大驼峰）
	SuffixCaseSnake string //关联表字段后缀（蛇形）
}

type handleExtendMiddle struct {
	tplOfTop                 myGenTpl
	tpl                      myGenTpl
	TableType                internal.MyGenTableType //表类型。按该字段区分哪种功能表
	RelId                    string                  //关联字段（tplOfTop中的字段名）
	FieldVar                 string                  //字段变量名
	daoPath                  string
	daoTable                 string
	daoTableVar              string       //表变量名
	FieldList                []myGenField //字段数组。除了自增主键，RelId，创建时间，更新时间，软删除等字段外的其它字段，这些字段才会生成代码
	FieldListOfIdSuffix      []myGenField //FieldList中的id后缀字段数组
	FieldListOfOther         []myGenField //FieldList中除id后缀字段外的其它字段数组
	FieldColumnArr           []string
	FieldColumnArrOfIdSuffix []string
	FieldColumnArrOfOther    []string
}

type handleOtherRel struct {
	tplOfTop myGenTpl
	tpl      myGenTpl
	RelId    string //关联字段（tpl中的字段名）
	daoPath  string
}

// 创建模板参数
func createTpl(ctx context.Context, group, table, removePrefixCommon, removePrefixAlone string, isTop bool) (tpl myGenTpl) {
	tpl = myGenTpl{
		Group:              group,
		RemovePrefixCommon: removePrefixCommon,
		RemovePrefixAlone:  removePrefixAlone,
		RemovePrefix:       removePrefixCommon + removePrefixAlone,
		Table:              table,
	}
	tpl.DbHandler = internal.NewMyGenDbHandler(ctx, g.DB(tpl.Group).GetConfig().Type)
	tpl.Link = gconv.String(gconv.SliceMap(g.Cfg().MustGet(ctx, `database`).MapStrAny()[tpl.Group])[0][`link`])
	tpl.TableArr = tpl.DbHandler.GetTableArr(ctx, tpl.Group)
	tpl.KeyList = tpl.DbHandler.GetKeyList(ctx, tpl.Group, tpl.Table)
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
	tpl.ModuleDirCaseCamel = moduleDirCaseCamel
	tpl.I18nPath = gstr.Replace(moduleDirCaseKebab, `/`, `.`) + `.` + tpl.TableCaseKebab

	fieldListTmp := tpl.DbHandler.GetFieldList(ctx, tpl.Group, tpl.Table)
	tpl.FieldStyle = internal.GetFieldStyle(fieldListTmp)
	fieldList := make([]myGenField, len(fieldListTmp))
	for k, v := range fieldListTmp {
		fieldTmp := myGenField{MyGenField: v}
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

		fieldTmp.FieldShowLenMax = internal.GetShowLen(fieldTmp.FieldName)
		fieldTmp.FieldLimitStr = tpl.DbHandler.GetFieldLimitStr(ctx, v, tpl.Group, tpl.Table)
		fieldTmp.FieldLimitInt = tpl.DbHandler.GetFieldLimitInt(ctx, v, tpl.Group, tpl.Table)
		fieldTmp.FieldLimitFloat = tpl.DbHandler.GetFieldLimitFloat(ctx, v, tpl.Group, tpl.Table)

		/*--------确定字段数据类型 开始--------*/
		fieldTmp.FieldType = tpl.DbHandler.GetFieldType(ctx, v, tpl.Group, tpl.Table)
		/*--------确定字段数据类型 结束--------*/

		/*--------确定字段主键类型 开始--------*/
		for _, key := range tpl.KeyList {
			if !garray.NewStrArrayFrom(key.FieldArr).Contains(fieldTmp.FieldRaw) {
				continue
			}
			if key.IsUnique && len(key.FieldArr) == 1 {
				fieldTmp.IsUnique = true
			}
			if !key.IsPrimary {
				continue
			}
			if len(key.FieldArr) == 1 {
				fieldTmp.FieldTypePrimary = internal.TypePrimary
				if key.IsAutoInc {
					fieldTmp.FieldTypePrimary = internal.TypePrimaryAutoInc
				}
			} else {
				fieldTmp.FieldTypePrimary = internal.TypePrimaryMany
				if key.IsAutoInc {
					fieldTmp.FieldTypePrimary = internal.TypePrimaryManyAutoInc
				}
			}
		}
		/*--------确定字段主键类型 结束--------*/

		/*--------确定字段命名类型（部分命名类型需做二次确定） 开始--------*/
		fieldSplitArr := gstr.Split(fieldTmp.FieldCaseSnakeRemove, `_`)
		fieldPrefix := fieldSplitArr[0]
		fieldSuffix := fieldSplitArr[len(fieldSplitArr)-1]

		if garray.NewStrArrayFrom(internal.ConfigFieldNameArrDeleted).Contains(fieldTmp.FieldCaseCamel) {
			fieldTmp.FieldTypeName = internal.TypeNameDeleted
		} else if garray.NewStrArrayFrom(internal.ConfigFieldNameArrUpdated).Contains(fieldTmp.FieldCaseCamel) {
			fieldTmp.FieldTypeName = internal.TypeNameUpdated
		} else if garray.NewStrArrayFrom(internal.ConfigFieldNameArrCreated).Contains(fieldTmp.FieldCaseCamel) {
			fieldTmp.FieldTypeName = internal.TypeNameCreated
		} else if garray.NewIntArrayFrom([]int{internal.TypeVarchar, internal.TypeText}).Contains(fieldTmp.FieldType) && fieldTmp.FieldCaseCamel == `IdPath` { //idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效
			fieldTmp.FieldTypeName = internal.TypeNameIdPath

			tpl.Handle.Pid.IdPath = fieldTmp.FieldRaw
		} else if garray.NewIntArrayFrom([]int{internal.TypeInt, internal.TypeIntU, internal.TypeVarchar, internal.TypeChar}).Contains(fieldTmp.FieldType) && garray.NewStrArrayFrom([]string{`status`, `type`, `method`, `pos`, `position`, `gender`, `currency`}).Contains(fieldSuffix) { //status,type,method,pos,position,gender,currency等后缀
			fieldTmp.FieldTypeName = internal.TypeNameStatusSuffix

			isStr := false
			if garray.NewIntArrayFrom([]int{internal.TypeVarchar, internal.TypeChar}).Contains(fieldTmp.FieldType) {
				isStr = true
			}
			fieldTmp.StatusList = internal.GetStatusList(fieldTmp.FieldTip, isStr)

			for _, status := range fieldTmp.StatusList {
				showLen := internal.GetShowLen(status[1])
				if showLen > fieldTmp.FieldShowLenMax {
					fieldTmp.FieldShowLenMax = showLen
				}
			}
		} else if garray.NewIntArrayFrom([]int{internal.TypeVarchar, internal.TypeText, internal.TypeJson}).Contains(fieldTmp.FieldType) && (garray.NewStrArrayFrom([]string{`icon`, `cover`, `avatar`, `img`, `image`}).Contains(fieldSuffix) || gstr.SubStr(fieldTmp.FieldCaseCamelRemove, -7) == `ImgList` || gstr.SubStr(fieldTmp.FieldCaseCamelRemove, -6) == `ImgArr` || gstr.SubStr(fieldTmp.FieldCaseCamelRemove, -9) == `ImageList` || gstr.SubStr(fieldTmp.FieldCaseCamelRemove, -8) == `ImageArr`) { //icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀
			fieldTmp.FieldTypeName = internal.TypeNameImageSuffix
		} else if garray.NewIntArrayFrom([]int{internal.TypeVarchar, internal.TypeText, internal.TypeJson}).Contains(fieldTmp.FieldType) && (garray.NewStrArrayFrom([]string{`video`}).Contains(fieldSuffix) || gstr.SubStr(fieldTmp.FieldCaseCamelRemove, -9) == `VideoList` || gstr.SubStr(fieldTmp.FieldCaseCamelRemove, -8) == `VideoArr`) { //video,video_list,videoList,video_arr,videoArr等后缀
			fieldTmp.FieldTypeName = internal.TypeNameVideoSuffix
		} else if garray.NewIntArrayFrom([]int{internal.TypeVarchar, internal.TypeText, internal.TypeJson}).Contains(fieldTmp.FieldType) && (garray.NewStrArrayFrom([]string{`file`}).Contains(fieldSuffix) || gstr.SubStr(fieldTmp.FieldCaseCamelRemove, -9) == `FileList` || gstr.SubStr(fieldTmp.FieldCaseCamelRemove, -8) == `FileArr`) { //file,file_list,fileList,file_arr,fileArr等后缀
			fieldTmp.FieldTypeName = internal.TypeNameFileSuffix
		} else if garray.NewIntArrayFrom([]int{internal.TypeText, internal.TypeJson}).Contains(fieldTmp.FieldType) && garray.NewStrArrayFrom([]string{`list`, `arr`}).Contains(fieldSuffix) { //list,arr等后缀
			fieldTmp.FieldTypeName = internal.TypeNameArrSuffix
		} else if garray.NewIntArrayFrom([]int{internal.TypeVarchar, internal.TypeText}).Contains(fieldTmp.FieldType) && garray.NewStrArrayFrom([]string{`remark`, `desc`, `msg`, `message`, `intro`, `content`}).Contains(fieldSuffix) { //remark,desc,msg,message,intro,content后缀
			fieldTmp.FieldTypeName = internal.TypeNameRemarkSuffix
		} else if fieldTmp.FieldType == internal.TypeVarchar { //varchar类型
			if garray.NewStrArrayFrom([]string{`name`, `title`}).Contains(fieldSuffix) { //name,title后缀
				fieldTmp.FieldTypeName = internal.TypeNameNameSuffix
			} else if garray.NewStrArrayFrom([]string{`code`}).Contains(fieldSuffix) { //code后缀
				fieldTmp.FieldTypeName = internal.TypeNameCodeSuffix
			} else if garray.NewStrArrayFrom([]string{`account`}).Contains(fieldSuffix) { //account后缀
				fieldTmp.FieldTypeName = internal.TypeNameAccountSuffix
			} else if garray.NewStrArrayFrom([]string{`phone`, `mobile`}).Contains(fieldSuffix) { //phone,mobile后缀
				fieldTmp.FieldTypeName = internal.TypeNamePhoneSuffix
			} else if garray.NewStrArrayFrom([]string{`email`}).Contains(fieldSuffix) { //email后缀
				fieldTmp.FieldTypeName = internal.TypeNameEmailSuffix
			} else if garray.NewStrArrayFrom([]string{`url`, `link`}).Contains(fieldSuffix) { //url,link后缀
				fieldTmp.FieldTypeName = internal.TypeNameUrlSuffix
			} else if garray.NewStrArrayFrom([]string{`ip`}).Contains(fieldSuffix) { //IP后缀
				fieldTmp.FieldTypeName = internal.TypeNameIpSuffix
			}
		} else if fieldTmp.FieldType == internal.TypeChar { //char类型
			if garray.NewStrArrayFrom([]string{`password`, `passwd`}).Contains(fieldSuffix) && fieldTmp.FieldLimitStr == `32` { //password,passwd后缀
				fieldTmp.FieldTypeName = internal.TypeNamePasswordSuffix

				passwordMapKey := internal.GetHandlePasswordMapKey(fieldTmp.FieldRaw)
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
				fieldTmp.FieldTypeName = internal.TypeNameSaltSuffix

				passwordMapKey := internal.GetHandlePasswordMapKey(fieldTmp.FieldRaw)
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
		} else if garray.NewIntArrayFrom([]int{internal.TypeInt, internal.TypeIntU}).Contains(fieldTmp.FieldType) { //int等类型
			if fieldTmp.FieldRaw == `pid` { //pid
				fieldTmp.FieldTypeName = internal.TypeNamePid

				tpl.Handle.Pid.Pid = fieldTmp.FieldRaw
			} else if garray.NewStrArrayFrom([]string{`sort`, `num`, `number`, `weight`}).Contains(fieldSuffix) { //sort,num,number,weight等后缀
				fieldTmp.FieldTypeName = internal.TypeNameSortSuffix
				if fieldSuffix == `sort` {
					tpl.Handle.Pid.Sort = append(tpl.Handle.Pid.Sort, fieldTmp.FieldRaw)
				}
			} else if garray.NewStrArrayFrom([]string{`no`, `level`, `rank`}).Contains(fieldSuffix) { //no,level,rank等后缀
				fieldTmp.FieldTypeName = internal.TypeNameNoSuffix
				if fieldTmp.FieldRaw == `level` { //level，且pid,level,idPath|id_path同时存在时（才）有效。该命名类型需做二次确定
					fieldTmp.FieldTypeName = internal.TypeNameLevel

					tpl.Handle.Pid.Level = fieldTmp.FieldRaw
				}
			} else if garray.NewStrArrayFrom([]string{`id`}).Contains(fieldSuffix) { //id后缀
				primaryKeyArr := []string{gstr.TrimLeftStr(gstr.TrimLeftStr(tpl.Table, tpl.RemovePrefixCommon, 1), tpl.RemovePrefixAlone, 1) + `_id`}
				if primaryKeyArr[0] != `id` {
					primaryKeyArr = append(primaryKeyArr, `id`)
				}
				if !garray.NewStrArrayFrom(primaryKeyArr).Contains(fieldTmp.FieldCaseSnake) { // 本表id字段不算
					fieldTmp.FieldTypeName = internal.TypeNameIdSuffix

					fieldTmp.FieldLimitInt.Min = `1`
					handleRelIdObj := handleRelId{
						tpl:       tpl.getRelIdTpl(ctx, tpl, fieldTmp.FieldRaw),
						FieldName: fieldTmp.FieldName,
					}
					if gstr.ToUpper(gstr.SubStr(handleRelIdObj.FieldName, -2)) == `ID` {
						handleRelIdObj.FieldName = gstr.SubStr(handleRelIdObj.FieldName, 0, -2)
					}
					if pos := gstr.Pos(fieldTmp.FieldCaseSnake, `_of_`); pos != -1 {
						handleRelIdObj.SuffixCaseSnake = gstr.SubStr(fieldTmp.FieldCaseSnake, pos)
						handleRelIdObj.SuffixCaseCamel = gstr.CaseCamel(handleRelIdObj.SuffixCaseSnake)
						handleRelIdObj.Suffix = handleRelIdObj.SuffixCaseSnake
						if fieldTmp.FieldRaw != fieldTmp.FieldCaseSnake {
							handleRelIdObj.Suffix = handleRelIdObj.SuffixCaseCamel
						}
					}
					tpl.Handle.RelIdMap[fieldTmp.FieldRaw] = handleRelIdObj
				}
			} else if garray.NewStrArrayFrom([]string{`is`}).Contains(fieldPrefix) { //is_前缀
				fieldTmp.FieldTypeName = internal.TypeNameIsPrefix
				// TODO 可改成状态一样处理，同时需要修改前端开关组件属性设置（暂时不改）
			}
		} else if garray.NewIntArrayFrom([]int{internal.TypeDatetime, internal.TypeTimestamp, internal.TypeDate, internal.TypeTime}).Contains(fieldTmp.FieldType) { //类型：datetime或date或timestamp或time
			if garray.NewStrArrayFrom([]string{`start`}).Contains(fieldPrefix) { //start_前缀
				fieldTmp.FieldTypeName = internal.TypeNameStartPrefix
			} else if garray.NewStrArrayFrom([]string{`end`}).Contains(fieldPrefix) { //end_前缀
				fieldTmp.FieldTypeName = internal.TypeNameEndPrefix
			}
		}
		/*--------确定字段命名类型（部分命名类型需做二次确定） 结束--------*/

		fieldList[k] = fieldTmp
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
	if isTop {
		if tpl.Handle.Pid.Pid != `` && tpl.Handle.Pid.Level != `` && tpl.Handle.Pid.IdPath != `` {
			tpl.Handle.Pid.IsCoexist = true
		}
	}
	/*--------解析影响命名类型二次确认，且需特殊处理的字段 结束--------*/

	/*--------命名类型二次确认的字段 开始--------*/
	for k, v := range fieldList {
		switch v.FieldTypeName {
		case internal.TypeNameLevel: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
			if !tpl.Handle.Pid.IsCoexist {
				fieldList[k].FieldTypeName = internal.TypeNameNoSuffix
			} else {
				fieldList[k].FieldLimitInt.Min = `1`
			}
		case internal.TypeNameIdPath: // idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
			if !tpl.Handle.Pid.IsCoexist {
				fieldList[k].FieldTypeName = ``
			}
		case internal.TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
			passwordMapKey := internal.GetHandlePasswordMapKey(v.FieldRaw)
			if !tpl.Handle.PasswordMap[passwordMapKey].IsCoexist {
				fieldList[k].FieldTypeName = ``
			}
		}
	}
	/*--------命名类型二次确认的字段 结束--------*/

	/*--------需特殊处理的字段解析 开始--------*/
	//主键列表（无主键时，默认为排除internal.ConfigIdAndLabelExcField过后的第一个字段）。联合主键有多字段，需按顺序存入
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
	/*
		label列表。sql查询可设为别名label的字段（常用于前端my-select或my-cascader等组件，或用于关联表查询）。按以下优先级存入：
			表名去掉前缀 + Name > 主键去掉ID + Name > Name >
			表名去掉前缀 + Title > 主键去掉ID + Title > Title >
			表名去掉前缀 + Phone > 主键去掉ID + Phone > Phone >
			表名去掉前缀 + Email > 主键去掉ID + Email > Email >
			表名去掉前缀 + Account > 主键去掉ID + Account > Account >
			表名去掉前缀 + Nickname > 主键去掉ID + Nickname > Nickname >
			上面字段都没有时，默认为排除internal.ConfigIdAndLabelExcField过后的第二个字段
	*/
	labelList := []string{}
	for _, v := range []string{`Name`, `Title`, `Phone`, `Email`, `Account`, `Nickname`} {
		labelTmp := tpl.TableCaseCamel + v
		labelList = append(labelList, labelTmp)
		if len(tpl.Handle.Id.List) == 1 && tpl.Handle.Id.IsPrimary {
			fieldSplitArr := gstr.Split(tpl.Handle.Id.List[0].FieldCaseSnake, `_`)
			if fieldSplitArr[len(fieldSplitArr)-1] == `id` {
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
			if v == item.FieldCaseCamel && garray.NewIntArrayFrom([]int{internal.TypeVarchar, internal.TypeChar}).Contains(item.FieldType) {
				tpl.Handle.LabelList = append(tpl.Handle.LabelList, item.FieldRaw)
				break
			}
		}
	}

	if len(tpl.Handle.Id.List) == 0 || len(tpl.Handle.LabelList) == 0 {
		idAndLabelfieldList := []myGenField{}
		for _, v := range tpl.FieldList {
			isFind := false
			for _, afterField := range internal.ConfigIdAndLabelExcField {
				switch val := afterField.(type) {
				case internal.MyGenFieldTypeName:
					if val == v.FieldTypeName {
						isFind = true
					}
				case internal.MyGenFieldArrOfTypeName:
					if (val.FieldType == 0 || val.FieldType == v.FieldType) &&
						(val.FieldTypeName == `` || val.FieldTypeName == v.FieldTypeName) &&
						(val.FieldArr.IsEmpty() || val.FieldArr.Contains(v.FieldRaw)) {
						isFind = true
					}
				}
			}
			if !isFind {
				idAndLabelfieldList = append(idAndLabelfieldList, v)
			}
		}

		if len(tpl.Handle.Id.List) == 0 {
			if len(idAndLabelfieldList) > 0 {
				tpl.Handle.Id.List = append(tpl.Handle.Id.List, idAndLabelfieldList[0])
			} else {
				tpl.Handle.Id.List = append(tpl.Handle.Id.List, fieldList[0])
			}
		}
		if len(tpl.Handle.LabelList) == 0 {
			if len(idAndLabelfieldList) > 1 {
				tpl.Handle.LabelList = append(tpl.Handle.LabelList, idAndLabelfieldList[1].FieldRaw)
			} else {
				tpl.Handle.LabelList = append(tpl.Handle.LabelList, fieldList[1].FieldRaw)
			}
		}
	}

	//id后缀字段
	for k, v := range tpl.Handle.RelIdMap {
		if v.tpl.Table == `` {
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

	if isTop {
		tpl.Handle.ExtendTableOneList, tpl.Handle.ExtendTableManyList = tpl.getExtendTable(ctx, tpl) //扩展表
		tpl.Handle.MiddleTableOneList, tpl.Handle.MiddleTableManyList = tpl.getMiddleTable(ctx, tpl) //中间表
		tpl.Handle.OtherRelTableList = tpl.getOtherRel(ctx, tpl)                                     //其它关联当前表主键的表（不含当前表的扩展表和中间表）
	}
	/*--------需特殊处理的字段解析 结束--------*/

	tpl.FieldList = fieldList

	afterFieldArr := garray.NewStrArray()
	for _, afterField := range internal.ConfigAfterField {
		for _, v := range tpl.FieldList {
			isFind := false
			switch val := afterField.(type) {
			case internal.MyGenFieldTypeName:
				if val == v.FieldTypeName {
					isFind = true
				}
			case internal.MyGenFieldArrOfTypeName:
				if (val.FieldType == 0 || val.FieldType == v.FieldType) &&
					(val.FieldTypeName == `` || val.FieldTypeName == v.FieldTypeName) &&
					(val.FieldArr.IsEmpty() || val.FieldArr.Contains(v.FieldRaw)) {
					isFind = true
				}
			}
			if isFind && !afterFieldArr.Contains(v.FieldRaw) {
				tpl.FieldListOfAfter = append([]myGenField{v}, tpl.FieldListOfAfter...)
				afterFieldArr.Append(v.FieldRaw)
				break
			}
		}
	}
	for _, v := range tpl.FieldList {
		if !afterFieldArr.Contains(v.FieldRaw) {
			tpl.FieldListOfDefault = append(tpl.FieldListOfDefault, v)
		}
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
		`--doPath`, `model/do/` + myGenTplThis.ModuleDirCaseKebab,
		`--entityPath`, `model/entity/` + myGenTplThis.ModuleDirCaseKebab,
		`--tables`, myGenTplThis.Table,
		`--tplDaoIndexPath`, `resource/gen/gen_dao_template_dao.txt`,
		`--tplDaoInternalPath`, `resource/gen/gen_dao_template_dao_internal.txt`,
	}
	if isOverwriteDao {
		commandArg = append(commandArg, `--overwriteDao=true`)
	}
	internal.Command(`表（`+myGenTplThis.Table+`）dao生成`, true, ``, `gf`, commandArg...)
}

// 判断字段是否与表主键一致
func (myGenTplThis *myGenTpl) IsSamePrimary(tpl myGenTpl, field string) bool {
	primaryKeyArr := []string{tpl.Handle.Id.List[0].FieldCaseSnake}
	if primaryKeyArr[0] == `id` {
		primaryKeyArr = append(primaryKeyArr, gstr.TrimLeftStr(gstr.TrimLeftStr(tpl.Table, tpl.RemovePrefixCommon, 1), tpl.RemovePrefixAlone, 1)+`_id`)
	}
	return garray.NewStrArrayFrom(primaryKeyArr).Contains(gstr.CaseSnake(field))
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
		tableKeyList := tpl.DbHandler.GetKeyList(ctx, tpl.Group, table)
		for _, v := range tableKeyList {
			if v.IsPrimary && len(v.FieldArr) == 1 && garray.NewStrArrayFrom([]string{`id`, fieldCaseSnakeOfRemove}).Contains(gstr.CaseSnake(v.FieldArr[0])) {
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

		relTpl = createTpl(ctx, tpl.Group, table, removePrefixCommon, removePrefixAlone, false)
		relTpl.gfGenDao(false) //dao文件生成
	}
	return
}

// 创建扩展表和中间表模板参数
func (myGenTplThis *myGenTpl) createExtendMiddleTpl(tplOfTop myGenTpl, extendMiddleTpl myGenTpl, relId string) (handleExtendMiddleObj handleExtendMiddle) {
	extendMiddleTpl.gfGenDao(false) //dao文件生成

	handleExtendMiddleObj = handleExtendMiddle{
		tplOfTop:    tplOfTop,
		tpl:         extendMiddleTpl,
		RelId:       relId,
		FieldVar:    internal.GetStrByFieldStyle(tplOfTop.FieldStyle, extendMiddleTpl.TableCaseCamel),
		daoPath:     extendMiddleTpl.TableCaseCamel,
		daoTable:    extendMiddleTpl.TableCaseCamel + `.ParseDbTable(m.GetCtx())`,
		daoTableVar: `table` + extendMiddleTpl.TableCaseCamel,
	}
	if extendMiddleTpl.ModuleDirCaseKebab != tplOfTop.ModuleDirCaseKebab {
		handleExtendMiddleObj.FieldVar = internal.GetStrByFieldStyle(tplOfTop.FieldStyle, extendMiddleTpl.ModuleDirCaseCamel+extendMiddleTpl.TableCaseCamel)
		handleExtendMiddleObj.daoPath = `dao` + extendMiddleTpl.ModuleDirCaseCamel + `.` + extendMiddleTpl.TableCaseCamel
		handleExtendMiddleObj.daoTable = `dao` + extendMiddleTpl.ModuleDirCaseCamel + `.` + extendMiddleTpl.TableCaseCamel + `.ParseDbTable(m.GetCtx())`
		if extendMiddleTpl.ModuleDirCaseCamel != extendMiddleTpl.TableCaseCamel {
			handleExtendMiddleObj.daoTableVar = `table` + extendMiddleTpl.ModuleDirCaseCamel + extendMiddleTpl.TableCaseCamel
		}
	}

	fieldArrOfIgnore := []string{relId}
	if extendMiddleTpl.Handle.Id.IsPrimary && len(extendMiddleTpl.Handle.Id.List) == 1 && extendMiddleTpl.Handle.Id.List[0].FieldRaw != relId {
		fieldArrOfIgnore = append(fieldArrOfIgnore, extendMiddleTpl.Handle.Id.List[0].FieldRaw)
	}
	for _, v := range extendMiddleTpl.FieldList {
		if garray.NewStrArrayFrom(fieldArrOfIgnore).Contains(v.FieldRaw) || garray.NewStrArrayFrom([]string{internal.TypeNameDeleted, internal.TypeNameUpdated, internal.TypeNameCreated}).Contains(v.FieldTypeName) {
			continue
		}
		handleExtendMiddleObj.FieldList = append(handleExtendMiddleObj.FieldList, v)
		if v.FieldTypeName == internal.TypeNameIdSuffix {
			handleExtendMiddleObj.FieldListOfIdSuffix = append(handleExtendMiddleObj.FieldListOfIdSuffix, v)
		} else {
			handleExtendMiddleObj.FieldListOfOther = append(handleExtendMiddleObj.FieldListOfOther, v)
		}
	}
	for _, v := range handleExtendMiddleObj.FieldList {
		handleExtendMiddleObj.FieldColumnArr = append(handleExtendMiddleObj.FieldColumnArr, handleExtendMiddleObj.daoPath+`.Columns().`+v.FieldCaseCamel)
	}
	for _, v := range handleExtendMiddleObj.FieldListOfIdSuffix {
		handleExtendMiddleObj.FieldColumnArrOfIdSuffix = append(handleExtendMiddleObj.FieldColumnArrOfIdSuffix, handleExtendMiddleObj.daoPath+`.Columns().`+v.FieldCaseCamel)
	}
	for _, v := range handleExtendMiddleObj.FieldListOfOther {
		handleExtendMiddleObj.FieldColumnArrOfOther = append(handleExtendMiddleObj.FieldColumnArrOfOther, handleExtendMiddleObj.daoPath+`.Columns().`+v.FieldCaseCamel)
	}
	return
}

// 获取扩展表
func (myGenTplThis *myGenTpl) getExtendTable(ctx context.Context, tpl myGenTpl) (extendTableOneList []handleExtendMiddle, extendTableManyList []handleExtendMiddle) {
	if len(tpl.Handle.Id.List) > 1 || !tpl.Handle.Id.IsPrimary { //联合主键或无主键时，不获取扩展表
		return
	}

	removePrefixCommon := tpl.RemovePrefixCommon
	removePrefixAlone := tpl.RemovePrefixAlone
	if removePrefixAlone == `` {
		removePrefixAlone = gstr.TrimLeftStr(tpl.Table, removePrefixCommon, 1) + `_`
	}
	for _, v := range tpl.TableArr {
		if v == tpl.Table { //自身跳过
			continue
		}
		if gstr.Pos(v, `_rel_to_`) != -1 || gstr.Pos(v, `_rel_of_`) != -1 { //中间表跳过
			continue
		}
		if gstr.Pos(v, tpl.Table+`_`) != 0 { // 不符合扩展表命名（主表名_xxxx）的跳过
			continue
		}
		extendTpl := createTpl(ctx, tpl.Group, v, removePrefixCommon, removePrefixAlone, false)
		for _, key := range extendTpl.KeyList {
			if len(key.FieldArr) != 1 {
				continue
			}
			if !myGenTplThis.IsSamePrimary(tpl, key.FieldArr[0]) {
				continue
			}
			handleExtendMiddleObj := myGenTplThis.createExtendMiddleTpl(tpl, extendTpl, key.FieldArr[0])
			if len(handleExtendMiddleObj.FieldList) == 0 { //没有要处理的字段，估计表有问题，不处理
				continue
			}
			if key.IsPrimary { //主键
				if !key.IsAutoInc { //不自增
					handleExtendMiddleObj.TableType = internal.TableTypeExtendOne
				}
			} else {
				if key.IsUnique { //唯一索引
					handleExtendMiddleObj.TableType = internal.TableTypeExtendOne
				} else { //普通索引
					handleExtendMiddleObj.TableType = internal.TableTypeExtendMany
				}
			}

			switch handleExtendMiddleObj.TableType {
			case internal.TableTypeExtendOne:
				isExtendOne := true
				fmt.Println(color.HiYellowString(`因扩展表的命名方式要求，无法百分百确定扩展表，故需手动确认`))
				isExtendOneStr := gcmd.Scan(color.BlueString(`> 表(` + extendTpl.Table + `)疑似为扩展表(一对一)，请确认？默认(yes)：`))
			isExtendOneEnd:
				for {
					switch isExtendOneStr {
					case ``, `1`, `yes`:
						isExtendOne = true
						break isExtendOneEnd
					case `0`, `no`:
						isExtendOne = false
						break isExtendOneEnd
					default:
						isExtendOneStr = gcmd.Scan(color.RedString(`    输入错误，请重新输入，表(` + extendTpl.Table + `)疑似为扩展表(一对一)，请确认？默认(yes)：`))
					}
				}
				if isExtendOne {
					extendTableOneList = append(extendTableOneList, handleExtendMiddleObj)
				}
			case internal.TableTypeExtendMany:
				isExtendMany := true
				fmt.Println(color.HiYellowString(`因扩展表的命名方式要求，无法百分百确定扩展表，故需手动确认`))
				isExtendManyStr := gcmd.Scan(color.BlueString(`> 表(` + extendTpl.Table + `)疑似为扩展表(一对多)，请确认？默认(yes)：`))
			isExtendManyEnd:
				for {
					switch isExtendManyStr {
					case ``, `1`, `yes`:
						isExtendMany = true
						break isExtendManyEnd
					case `0`, `no`:
						isExtendMany = false
						break isExtendManyEnd
					default:
						isExtendManyStr = gcmd.Scan(color.RedString(`    输入错误，请重新输入，表(` + extendTpl.Table + `)疑似为扩展表(一对多)，请确认？默认(yes)：`))
					}
				}
				if isExtendMany {
					if len(handleExtendMiddleObj.FieldList) == 1 {
						handleExtendMiddleObj.FieldVar = internal.GetStrByFieldStyle(handleExtendMiddleObj.tplOfTop.FieldStyle, handleExtendMiddleObj.FieldList[0].FieldRaw, ``, `arr`)
					} else {
						handleExtendMiddleObj.FieldVar = internal.GetStrByFieldStyle(handleExtendMiddleObj.tplOfTop.FieldStyle, gstr.TrimRightStr(gstr.TrimLeftStr(gstr.CaseCamelLower(handleExtendMiddleObj.FieldVar), `relTo`, 1), `RelOf`, 1), ``, `list`)
					}
					extendTableManyList = append(extendTableManyList, handleExtendMiddleObj)
				}
			}
		}
	}
	return
}

// 获取中间表
func (myGenTplThis *myGenTpl) getMiddleTable(ctx context.Context, tpl myGenTpl) (middleTableOneList []handleExtendMiddle, middleTableManyList []handleExtendMiddle) {
	if len(tpl.Handle.Id.List) > 1 || !tpl.Handle.Id.IsPrimary { //联合主键或无主键时，不获取中间表
		return
	}

	removePrefixCommon := ``
	removePrefixAlone := ``
	for _, v := range tpl.TableArr {
		if v == tpl.Table { //自身跳过
			continue
		}
		if gstr.Pos(v, `_rel_to_`) == -1 && gstr.Pos(v, `_rel_of_`) == -1 { //不是中间表跳过
			continue
		}
		if gstr.Pos(v, `_rel_to_`) != -1 {
			if gstr.Pos(v, tpl.Table+`_rel_to_`) != 0 { //不符合中间表_rel_to_命名的跳过
				continue
			}
			removePrefixCommon = tpl.RemovePrefixCommon
			removePrefixAlone = tpl.RemovePrefixAlone
			if removePrefixAlone == `` {
				removePrefixAlone = gstr.TrimLeftStr(tpl.Table, removePrefixCommon, 1) + `_`
			}
		} else {
			if gstr.Pos(v, tpl.RemovePrefix) == 0 { //不符合中间表_rel_of_命名的跳过（同模块）
				if len(v) != gstr.Pos(v, `_rel_of_`+tpl.Table)+len(`_rel_of_`+tpl.Table) || len(v) != gstr.Pos(v, `_rel_of_`+gstr.Replace(tpl.Table, tpl.RemovePrefix, ``, 1))+len(`_rel_of_`+gstr.Replace(tpl.Table, tpl.RemovePrefix, ``, 1)) {
					continue
				}
				removePrefixCommon = tpl.RemovePrefixCommon
				removePrefixAlone = tpl.RemovePrefixAlone
				if removePrefixAlone == `` {
					removePrefixAlone = gstr.TrimLeftStr(tpl.Table, removePrefixCommon, 1) + `_`
				}
			} else { //不符合中间表_rel_of_命名的跳过（不同模块）
				if len(v) != gstr.Pos(v, `_rel_of_`+tpl.Table)+len(`_rel_of_`+tpl.Table) {
					continue
				}
				removePrefixCommon = tpl.RemovePrefixCommon
				if gstr.Pos(v, tpl.RemovePrefixCommon) != 0 {
					removePrefixCommon = ``
				}
				// 第一个分隔符之前的部分设置为removePrefixAlone
				tableRemove := gstr.TrimLeftStr(v, removePrefixCommon, 1)
				removePrefixAlone = gstr.SubStr(tableRemove, 0, gstr.Pos(tableRemove, `_`)+1)
			}
		}

		middleTpl := createTpl(ctx, tpl.Group, v, removePrefixCommon, removePrefixAlone, false)
		for _, key := range middleTpl.KeyList {
			if !key.IsUnique { // 必须唯一
				continue
			}
			keyField := ``
			for _, keyFieldTmp := range key.FieldArr {
				if myGenTplThis.IsSamePrimary(tpl, keyFieldTmp) {
					keyField = keyFieldTmp
					break
				}
			}
			if keyField == `` {
				continue
			}
			handleExtendMiddleObj := myGenTplThis.createExtendMiddleTpl(tpl, middleTpl, keyField)
			if len(handleExtendMiddleObj.FieldList) == 0 { //没有要处理的字段，估计表有问题，不处理
				continue
			}
			if len(handleExtendMiddleObj.FieldListOfIdSuffix) == 0 { //没有其它表的关联id字段，不是中间表
				continue
			}
			if len(key.FieldArr) == 1 {
				if key.IsPrimary { //主键
					if !key.IsAutoInc { //不自增
						handleExtendMiddleObj.TableType = internal.TableTypeMiddleOne
					}
				} else { //唯一索引
					handleExtendMiddleObj.TableType = internal.TableTypeMiddleOne
				}
			} else {
				isAllId := true
				for _, v := range key.FieldArr {
					vArr := gstr.Split(gstr.CaseSnake(v), `_`)
					if vArr[len(vArr)-1] != `id` {
						isAllId = false
					}
				}
				if isAllId { //联合主键 或 联合唯一索引
					handleExtendMiddleObj.TableType = internal.TableTypeMiddleMany
				}
			}

			switch handleExtendMiddleObj.TableType {
			case internal.TableTypeMiddleOne:
				middleTableOneList = append(middleTableOneList, handleExtendMiddleObj)
			case internal.TableTypeMiddleMany:
				if len(handleExtendMiddleObj.FieldList) == 1 {
					handleExtendMiddleObj.FieldVar = internal.GetStrByFieldStyle(handleExtendMiddleObj.tplOfTop.FieldStyle, handleExtendMiddleObj.FieldList[0].FieldRaw, ``, `arr`)
				} else {
					handleExtendMiddleObj.FieldVar = internal.GetStrByFieldStyle(handleExtendMiddleObj.tplOfTop.FieldStyle, gstr.TrimRightStr(gstr.TrimLeftStr(gstr.CaseCamelLower(handleExtendMiddleObj.FieldVar), `relTo`, 1), `RelOf`, 1), ``, `list`)
				}
				middleTableManyList = append(middleTableManyList, handleExtendMiddleObj)
			}
		}
	}
	return
}

// 其它关联表（不含扩展表和中间表）
func (myGenTplThis *myGenTpl) getOtherRel(ctx context.Context, tpl myGenTpl) (otherRelTableList []handleOtherRel) {
	if len(tpl.Handle.Id.List) > 1 || !tpl.Handle.Id.IsPrimary { //联合主键或无主键时，不获取其它关联表
		return
	}
	extendMiddleTableArr := garray.NewStrArray()
	for _, v := range tpl.Handle.ExtendTableOneList {
		extendMiddleTableArr.Append(v.tpl.Table)
	}
	for _, v := range tpl.Handle.ExtendTableManyList {
		extendMiddleTableArr.Append(v.tpl.Table)
	}
	for _, v := range tpl.Handle.MiddleTableOneList {
		extendMiddleTableArr.Append(v.tpl.Table)
	}
	for _, v := range tpl.Handle.MiddleTableManyList {
		extendMiddleTableArr.Append(v.tpl.Table)
	}

	extendMiddleTableArr.Append(tpl.Table) //追加自身
	for _, v := range tpl.TableArr {
		if extendMiddleTableArr.Contains(v) {
			continue
		}

		removePrefixCommon := tpl.RemovePrefixCommon
		if gstr.Pos(v, tpl.RemovePrefixCommon) != 0 {
			// continue //不是相同公共前缀跳过
			removePrefixCommon = ``
		}
		// 第一个分隔符之前的部分设置为removePrefixAlone
		tableRemove := gstr.TrimLeftStr(v, removePrefixCommon, 1)
		removePrefixAlone := gstr.SubStr(tableRemove, 0, gstr.Pos(tableRemove, `_`)+1)
		/* if removePrefixAlone != tpl.RemovePrefixAlone { //不是相同模块跳过
			continue
		} */

		otherRelTpl := createTpl(ctx, tpl.Group, v, removePrefixCommon, removePrefixAlone, false)
		for _, field := range otherRelTpl.FieldList {
			if myGenTplThis.IsSamePrimary(tpl, field.FieldCaseSnakeRemove) {
				handleOtherRelObj := handleOtherRel{
					tplOfTop: tpl,
					tpl:      otherRelTpl,
					RelId:    field.FieldRaw,
					daoPath:  otherRelTpl.TableCaseCamel,
				}
				if handleOtherRelObj.tpl.ModuleDirCaseKebab != handleOtherRelObj.tplOfTop.ModuleDirCaseKebab {
					handleOtherRelObj.daoPath = `dao` + handleOtherRelObj.tpl.ModuleDirCaseCamel + `.` + handleOtherRelObj.tpl.TableCaseCamel
				}
				otherRelTableList = append(otherRelTableList, handleOtherRelObj)
				break
			}
		}
	}
	return
}
