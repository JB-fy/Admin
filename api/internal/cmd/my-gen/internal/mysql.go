package internal

import (
	"context"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
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

func (dbHandler mysql) GetFieldLimitStr(ctx context.Context, field MyGenField, group, table string) (fieldLimitStr string) {
	fieldLimitStrTmp, _ := gregex.MatchString(`.*\((\d*)\)`, field.FieldTypeRaw)
	if len(fieldLimitStrTmp) > 1 {
		fieldLimitStr = fieldLimitStrTmp[1]
	}
	return
}

func (dbHandler mysql) GetFieldLimitInt(ctx context.Context, field MyGenField, group, table string) (fieldLimitInt int) {
	fieldLimitInt = 4
	if gstr.Pos(field.FieldTypeRaw, `tinyint`) != -1 || gstr.Pos(field.FieldTypeRaw, `smallint`) != -1 {
		fieldLimitInt = 2
	} else if gstr.Pos(field.FieldTypeRaw, `bigint`) != -1 {
		fieldLimitInt = 8
	}
	return
}

func (dbHandler mysql) GetFieldLimitFloat(ctx context.Context, field MyGenField, group, table string) (fieldLimitFloat [2]string) {
	fieldLimitFloatTmp, _ := gregex.MatchString(`.*\((\d*),(\d*)\)`, field.FieldTypeRaw)
	if len(fieldLimitFloatTmp) < 3 {
		fieldLimitFloatTmp = []string{``, `10`, `2`}
	}
	fieldLimitFloat = [2]string{fieldLimitFloatTmp[1], fieldLimitFloatTmp[2]}
	return
}

func (dbHandler mysql) GetFieldType(ctx context.Context, field MyGenField, group, table string) (fieldType MyGenFieldType) {
	if gstr.Pos(field.FieldTypeRaw, `int`) != -1 && gstr.Pos(field.FieldTypeRaw, `point`) == -1 { //int等类型
		fieldType = TypeInt
		if gstr.Pos(field.FieldTypeRaw, `unsigned`) != -1 || field.IsAutoInc || garray.NewStrArrayFrom([]string{`pid`, `level`}).Contains(field.FieldRaw) {
			fieldType = TypeIntU
		} else {
			fieldCaseSnake := gstr.CaseSnake(field.FieldRaw)
			fieldCaseSnakeRemove := gstr.Split(fieldCaseSnake, `_of_`)[0]
			fieldSplitArr := gstr.Split(fieldCaseSnakeRemove, `_`)
			fieldSuffix := fieldSplitArr[len(fieldSplitArr)-1]
			if fieldSuffix == `id` {
				fieldType = TypeIntU
			}
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
