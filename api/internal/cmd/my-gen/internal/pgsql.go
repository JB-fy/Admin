package internal

import (
	"context"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
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
			Default:      v[`column_default`].Interface(),
			Comment:      v[`column_comment`].String(),
		}
		if !v[`column_default`].IsNil() {
			defaultStr := v[`column_default`].String()
			if gstr.Pos(defaultStr, `nextval`) == 0 {
				field.IsAutoInc = true
				// field.Default = 0
			} else {
				if gstr.Pos(defaultStr, `::`) != -1 {
					switch gstr.Split(defaultStr, `::`)[0] {
					case `''`:
						field.Default = ``
					case `NULL`:
						field.Default = nil
					}
				}
			}
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
		if !gvar.New(v.Default).IsNil() {
			defaultStr := gconv.String(v.Default)
			if gstr.Pos(defaultStr, `nextval`) == 0 {
				field.IsAutoInc = true
				// field.Default = 0
			} else {
				if gstr.Pos(defaultStr, `::`) != -1 {
					switch gstr.Split(defaultStr, `::`)[0] {
					case `''`:
						field.Default = ``
					case `NULL`:
						field.Default = nil
					}
				}
			}
		}
		fieldList[v.Index] = field
	}
	return
}

// TODO
func (dbHandler pgsql) GetKeyList(ctx context.Context, group, table string) (keyList []MyGenKey) {
	keyListTmp, _ := g.DB(group).GetAll(ctx, `SELECT pg_index.* FROM pg_index JOIN pg_class ON pg_index.indrelid = pg_class.OID WHERE pg_class.relname = '`+table+`'`)
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
