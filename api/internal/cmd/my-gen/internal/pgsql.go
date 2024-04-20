package internal

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type pgsql struct {
	common
}

func (dbHandler pgsql) GetFieldList(ctx context.Context, group, table string) (fieldList []MyGenField) {
	/* fieldListTmp, _ := g.DB(group).GetAll(ctx, `SELECT *, col_description ( 'tes' :: REGCLASS, ordinal_position ) AS column_comment  FROM information_schema.COLUMNS WHERE TABLE_NAME = '`+table+`'`)
	fieldList = make([]MyGenField, len(fieldListTmp))
	for _, v := range fieldListTmp {
		field := MyGenField{
			FieldRaw:     v[`column_name`].String(),
			FieldTypeRaw: v[`data_type`].String(),
			IsNull:       v[`is_nullable`].Bool(),
			Default:      v[`column_default`].String(),
			Comment:      v[`column_comment`].String(),
		}
		if v[`Extra`].String() == `auto_increment` {
			field.IsAutoInc = true
		}
		fieldList[v[`ordinal_position`].Int()-1] = field
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
func (dbHandler pgsql) GetKeyList(ctx context.Context, group, table string) (keyList []MyGenKey) {
	keyListTmp, _ := g.DB(group).GetAll(ctx, `SELECT * FROM pg_index JOIN pg_class ON pg_index.indrelid = pg_class.OID WHERE pg_class.relname = '`+table+`'`)
	keyList = make([]MyGenKey, len(keyListTmp))
	fieldList := dbHandler.GetFieldList(ctx, group, table)
	fieldArrMap := map[string][]string{}
	for k, v := range keyListTmp {
		key := MyGenKey{
			Name:     v[`relname`].String(),
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
