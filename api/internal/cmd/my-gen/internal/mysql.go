package internal

import (
	"context"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type mysql struct {
	common
}

func (dbHandler mysql) GetFieldList(ctx context.Context, group, table string) (fieldList []MyGenField) {
	/* fieldListTmp, _ := g.DB(group).GetAll(ctx, `SHOW FULL COLUMNS FROM `+table)
	fieldList = make([]MyGenField, len(fieldListTmp))
	for k, v := range fieldListTmp {
		field := MyGenField{
			FieldRaw:     v[`Field`].String(),
			FieldTypeRaw: v[`Type`].String(),
			IsNull:       v[`Null`].Bool(),
			Default:      v[`Default`].Interface(),
			Comment:      v[`Comment`].String(),
		}
		if v[`Extra`].String() == `auto_increment` {
			field.IsAutoInc = true
		}
		fieldList[k] = field
	} */
	fieldListTmp, _ := g.DB(group).TableFields(ctx, table)
	fieldList = make([]MyGenField, len(fieldListTmp))
	for _, v := range fieldListTmp {
		field := MyGenField{
			FieldRaw:     v.Name,
			FieldTypeRaw: v.Type,
			IsNull:       v.Null,
			Default:      v.Default,
			Comment:      v.Comment,
		}
		if v.Extra == `auto_increment` {
			field.IsAutoInc = true
		}
		fieldList[v.Index] = field
	}
	return
}

func (dbHandler mysql) GetKeyList(ctx context.Context, group, table string) (keyList []MyGenKey) {
	keyListTmp, _ := g.DB(group).GetAll(ctx, `SHOW KEYS FROM `+table)
	fieldList := dbHandler.GetFieldList(ctx, group, table)
	keyNameList := []string{}
	fieldArrMap := map[string][]string{}
	for _, v := range keyListTmp {
		keyName := v[`Key_name`].String()
		if !garray.NewStrArrayFrom(keyNameList).Contains(keyName) {
			keyNameList = append(keyNameList, keyName)
		}
		fieldArrMap[keyName] = append(fieldArrMap[keyName], v[`Column_name`].String())
	}
	for _, keyName := range keyNameList {
		key := MyGenKey{}
		key.FieldArr = append(key.FieldArr, fieldArrMap[keyName]...)
		if keyName == `PRIMARY` {
			key.IsPrimary = true
		}
		for _, v := range keyListTmp {
			if keyName == v[`Key_name`].String() {
				key.IsUnique = !v[`Non_unique`].Bool()
				break
			}
		}
		if len(key.FieldArr) == 1 {
			for _, field := range fieldList {
				if key.FieldArr[0] == field.FieldRaw && field.IsAutoInc {
					key.IsAutoInc = true
					break
				}
			}
		}
		keyList = append(keyList, key)
	}
	return
}

func (dbHandler mysql) GetFieldType(ctx context.Context, field MyGenField, group, table string) (fieldType MyGenFieldType) {
	if gstr.Pos(field.FieldTypeRaw, `int`) != -1 && gstr.Pos(field.FieldTypeRaw, `point`) == -1 { //int等类型
		fieldType = TypeInt
		if gstr.Pos(field.FieldTypeRaw, `unsigned`) != -1 {
			fieldType = TypeIntU
		}
	} else if gstr.Pos(field.FieldTypeRaw, `decimal`) != -1 || gstr.Pos(field.FieldTypeRaw, `float`) != -1 || gstr.Pos(field.FieldTypeRaw, `double`) != -1 { //float类型
		fieldType = TypeFloat
		if gstr.Pos(field.FieldTypeRaw, `unsigned`) != -1 {
			fieldType = TypeFloatU
		}
	} else if gstr.Pos(field.FieldTypeRaw, `varchar`) != -1 { //varchar类型
		fieldType = TypeVarchar
	} else if gstr.Pos(field.FieldTypeRaw, `char`) != -1 { //char类型
		fieldType = TypeChar
	} else if gstr.Pos(field.FieldTypeRaw, `text`) != -1 { //text类型
		fieldType = TypeText
	} else if gstr.Pos(field.FieldTypeRaw, `json`) != -1 { //json类型
		fieldType = TypeJson
	} else if gstr.Pos(field.FieldTypeRaw, `datetime`) != -1 { //datetime类型
		fieldType = TypeDatetime
	} else if gstr.Pos(field.FieldTypeRaw, `date`) != -1 { //date类型
		fieldType = TypeDate
	} else if gstr.Pos(field.FieldTypeRaw, `timestamp`) != -1 { //timestamp类型
		fieldType = TypeTimestamp
	} else if gstr.Pos(field.FieldTypeRaw, `time`) != -1 { //time类型
		fieldType = TypeTime
	}
	return
}

func (dbHandler mysql) GetFieldLimitStr(ctx context.Context, field MyGenField, group, table string) (fieldLimitStr string) {
	fieldLimitStrTmp, _ := gregex.MatchString(`.*\((\d*)\)`, field.FieldTypeRaw)
	if len(fieldLimitStrTmp) > 1 {
		fieldLimitStr = fieldLimitStrTmp[1]
	}
	return
}

func (dbHandler mysql) GetFieldLimitInt(ctx context.Context, field MyGenField, group, table string) (fieldLimitInt MyGenFieldLimitInt) {
	fieldLimitInt.Size = 4
	if gstr.Pos(field.FieldTypeRaw, `tinyint`) != -1 || gstr.Pos(field.FieldTypeRaw, `smallint`) != -1 {
		fieldLimitInt.Size = 2
	} else if gstr.Pos(field.FieldTypeRaw, `bigint`) != -1 {
		fieldLimitInt.Size = 8
	}
	switch fieldLimitInt.Size {
	case 2:
		if gstr.Pos(field.FieldTypeRaw, `tinyint`) != -1 {
			fieldLimitInt.Min = `-128`
			fieldLimitInt.Max = `127`
			if gstr.Pos(field.FieldTypeRaw, `unsigned`) != -1 {
				fieldLimitInt.Min = `0`
				fieldLimitInt.Max = `255`
			}
		} else {
			fieldLimitInt.Min = `-32768`
			fieldLimitInt.Max = `32767`
			if gstr.Pos(field.FieldTypeRaw, `unsigned`) != -1 {
				fieldLimitInt.Min = `0`
				fieldLimitInt.Max = `65535`
			}
		}
	case 4:
		fieldLimitInt.Min = `-8388608`
		fieldLimitInt.Max = `8388607`
		if gstr.Pos(field.FieldTypeRaw, `unsigned`) != -1 {
			fieldLimitInt.Min = `0`
			fieldLimitInt.Max = `16777215`
		}
	case 8:
		fieldLimitInt.Min = `-9223372036854775808`
		fieldLimitInt.Max = `9223372036854775807`
		if gstr.Pos(field.FieldTypeRaw, `unsigned`) != -1 {
			fieldLimitInt.Min = `0`
			fieldLimitInt.Max = `18446744073709551615`
		}
	}
	if field.IsAutoInc {
		fieldLimitInt.Min = `1`
	}
	return
}

func (dbHandler mysql) GetFieldLimitFloat(ctx context.Context, field MyGenField, group, table string) (fieldLimitFloat MyGenFieldLimitFloat) {
	fieldLimitFloatTmp, _ := gregex.MatchString(`.*\((\d*),(\d*)\)`, field.FieldTypeRaw)
	if len(fieldLimitFloatTmp) < 3 {
		fieldLimitFloatTmp = []string{``, `10`, `2`}
	}
	fieldLimitFloat.Size = gconv.Int(fieldLimitFloatTmp[1])
	fieldLimitFloat.Precision = gconv.Int(fieldLimitFloatTmp[2])
	if gstr.Pos(field.FieldTypeRaw, `decimal`) != -1 /* || gstr.Pos(field.FieldTypeRaw, `float`) != -1 || gstr.Pos(field.FieldTypeRaw, `double`) != -1 */ {
		fieldLimitFloat.Max = gstr.Repeat(`9`, fieldLimitFloat.Size-fieldLimitFloat.Precision) + `.` + gstr.Repeat(`9`, fieldLimitFloat.Precision)
		fieldLimitFloat.Min = `-` + fieldLimitFloat.Max
	}
	if gstr.Pos(field.FieldTypeRaw, `unsigned`) != -1 {
		fieldLimitFloat.Min = `0`
	}
	return
}

func (dbHandler mysql) GetFuncFieldFormat(dbFuncCode MyGenDbFuncCode, field string) (fieldFormat string) {
	// 默认以Mysql为主，所以Mysql直接返回field
	fieldFormat = field //默认值
	return
}
