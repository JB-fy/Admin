package internal

import (
	"context"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
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
