package internal

import (
	"context"

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
	keyList = make([]MyGenKey, len(keyListTmp))
	fieldList := dbHandler.GetFieldList(ctx, group, table)
	fieldArrMap := map[string][]string{}
	for k, v := range keyListTmp {
		key := MyGenKey{
			Name:     v[`Key_name`].String(),
			Index:    v[`Seq_in_index`].Uint(),
			Field:    v[`Column_name`].String(),
			IsUnique: !v[`Non_unique`].Bool(),
		}
		if key.Name == `PRIMARY` {
			key.IsPrimary = true
			for _, field := range fieldList {
				if key.Field == field.FieldRaw && field.IsAutoInc {
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
	return
}
