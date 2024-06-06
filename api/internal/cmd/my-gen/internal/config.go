package internal

import "github.com/gogf/gf/v2/container/garray"

const ( //配置
	ConfigMaxLenOfStrFilter = 30  // 字段是TypeVarchar或TypeChar时，字段长度大于该值时，不生成过滤条件
	ConfigMaxLenOfStrHiddle = 120 // 字段是TypeVarchar或TypeChar时，字段长度大于等于该值时，前端列表显示设置with: 200, hidden: true
)

var (
	ConfigFieldNameArrCreated = []string{`CreatedAt`, `CreateAt`, `CreatedTime`, `CreateTime`} //创建时间字段命名（大驼峰比对）
	ConfigFieldNameArrUpdated = []string{`UpdatedAt`, `UpdateAt`, `UpdatedTime`, `UpdateTime`} //更新时间字段命名（大驼峰比对）
	ConfigFieldNameArrDeleted = []string{`DeletedAt`, `DeleteAt`, `DeletedTime`, `DeleteTime`} //软删除字段命名（大驼峰比对）

	// 当Id和Label找不到符合条件的字段时，排除这些字段后，默认剩余字段中的第一个和第二个
	ConfigIdAndLabelExcField = []any{
		TypeNameCreated,
		TypeNameUpdated,
		TypeNameDeleted,
		MyGenFieldArrOfTypeName{
			FieldTypeName: TypeNameIsPrefix,
			FieldArr:      garray.NewStrArrayFrom([]string{`is_stop`, `isStop`}),
		},
	}
	// 从上往下，将对应的字段放最后显示
	ConfigAfterField = []any{
		TypeNameCreated,
		TypeNameUpdated,
		TypeNameDeleted,
		MyGenFieldArrOfTypeName{
			FieldTypeName: TypeNameIsPrefix,
			FieldArr:      garray.NewStrArrayFrom([]string{`is_stop`, `isStop`}),
		},
		MyGenFieldArrOfTypeName{
			FieldType:     TypeText,
			FieldTypeName: TypeNameRemarkSuffix,
			FieldArr:      garray.NewStrArray(),
		},
		TypeNameIsPrefix,
	}
)

type MyGenFieldArrOfTypeName struct {
	FieldType     MyGenFieldType
	FieldTypeName MyGenFieldTypeName
	FieldArr      *garray.StrArray
}

type MyGenTableType = uint
type MyGenFieldTypePrimary = string
type MyGenFieldType = int
type MyGenFieldTypeName = string

const (
	TableTypeDefault    MyGenTableType = 0  //默认
	TableTypeExtendOne  MyGenTableType = 1  //扩展表（一对一）
	TableTypeExtendMany MyGenTableType = 2  //扩展表（一对多）
	TableTypeMiddleOne  MyGenTableType = 11 //中间表（一对一）
	TableTypeMiddleMany MyGenTableType = 12 //中间表（一对多）

	//用于结构体中，需从1开始，否则结构体会默认0
	TypeInt       MyGenFieldType = iota + 1 // `int等类型`
	TypeIntU                                // `int等类型（unsigned）`
	TypeFloat                               // `float等类型`
	TypeFloatU                              // `float等类型（unsigned）`
	TypeVarchar                             // `varchar类型`
	TypeChar                                // `char类型`
	TypeText                                // `text类型`
	TypeJson                                // `json类型`
	TypeDatetime                            // `datetime类型`
	TypeDate                                // `date类型`
	TypeTime                                // `time类型`
	TypeTimestamp                           // `timestamp类型`

	TypePrimary            MyGenFieldTypePrimary = `独立主键`
	TypePrimaryAutoInc     MyGenFieldTypePrimary = `独立主键（自增）`
	TypePrimaryMany        MyGenFieldTypePrimary = `联合主键`
	TypePrimaryManyAutoInc MyGenFieldTypePrimary = `联合主键（自增）`

	TypeNameDeleted        MyGenFieldTypeName = `软删除字段`
	TypeNameUpdated        MyGenFieldTypeName = `更新时间字段`
	TypeNameCreated        MyGenFieldTypeName = `创建时间字段`
	TypeNamePid            MyGenFieldTypeName = `命名：pid；	类型：int等类型；`
	TypeNameLevel          MyGenFieldTypeName = `命名：level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；`
	TypeNameIdPath         MyGenFieldTypeName = `命名：idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；`
	TypeNamePasswordSuffix MyGenFieldTypeName = `命名：password,passwd后缀；	类型：char(32)；`
	TypeNameSaltSuffix     MyGenFieldTypeName = `命名：salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；`
	TypeNameNameSuffix     MyGenFieldTypeName = `命名：name,title后缀；	类型：varchar；`
	TypeNameCodeSuffix     MyGenFieldTypeName = `命名：code后缀；	类型：varchar；`
	TypeNameAccountSuffix  MyGenFieldTypeName = `命名：account后缀；	类型：varchar；`
	TypeNamePhoneSuffix    MyGenFieldTypeName = `命名：phone,mobile后缀；	类型：varchar；`
	TypeNameEmailSuffix    MyGenFieldTypeName = `命名：email后缀；	类型：varchar；`
	TypeNameUrlSuffix      MyGenFieldTypeName = `命名：url,link后缀；	类型：varchar；`
	TypeNameIpSuffix       MyGenFieldTypeName = `命名：IP后缀；	类型：varchar；`
	TypeNameColorSuffix    MyGenFieldTypeName = `命名：color后缀；	类型：varchar；`
	TypeNameIdSuffix       MyGenFieldTypeName = `命名：id后缀；	类型：int等类型；`
	TypeNameSortSuffix     MyGenFieldTypeName = `命名：sort,num,number,weight等后缀；	类型：int等类型；`
	TypeNameNoSuffix       MyGenFieldTypeName = `命名：level,rank,no等后缀；	类型：int等类型；`
	TypeNameStatusSuffix   MyGenFieldTypeName = `命名：status,type,method,pos,position,gender,currency等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）`
	TypeNameIsPrefix       MyGenFieldTypeName = `命名：is_前缀；	类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）`
	TypeNameStartPrefix    MyGenFieldTypeName = `命名：start_前缀；	类型：datetime或date或timestamp或time；`
	TypeNameEndPrefix      MyGenFieldTypeName = `命名：end_前缀；	类型：datetime或date或timestamp或time；`
	TypeNameRemarkSuffix   MyGenFieldTypeName = `命名：remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器`
	TypeNameImageSuffix    MyGenFieldTypeName = `命名：icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text`
	TypeNameVideoSuffix    MyGenFieldTypeName = `命名：video,video_list,videoList,video_arr,videoArr等后缀；	类型：单视频varchar，多视频json或text`
	TypeNameFileSuffix     MyGenFieldTypeName = `命名：file,file_list,fileList,file_arr,fileArr等后缀；	类型：单文件varchar，多文件json或text`
	TypeNameArrSuffix      MyGenFieldTypeName = `命名：list,arr等后缀；	类型：json或text；`
)
