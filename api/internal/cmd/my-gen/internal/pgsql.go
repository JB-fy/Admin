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

func (dbHandler pgsql) GetKeyList(ctx context.Context, group, table string) (keyList []MyGenKey) {
	indrelid, _ := g.DB(group).GetValue(ctx, `SELECT oid FROM pg_class WHERE relname = '`+table+`'`)
	keyListTmp, _ := g.DB(group).GetAll(ctx, `SELECT * FROM pg_index WHERE indrelid = '`+indrelid.String()+`'`)
	fieldList := dbHandler.GetFieldList(ctx, group, table)
	for _, v := range keyListTmp {
		if !v[`indisvalid`].Bool() {
			continue
		}
		key := MyGenKey{
			IsPrimary: v[`indisprimary`].Bool(),
			IsUnique:  v[`indisunique`].Bool(),
		}
		// g.DB(group).GetValue(ctx, `SELECT indkey FROM pg_index WHERE indexrelid = `+v[`indexrelid`].String())
		fieldIndex := v[`indkey`].Int() //TODO indkey返回值有BUG。联合索引应该返回4 5，但这里却是0。
		if fieldIndex != 0 {
			key.FieldArr = append(key.FieldArr, fieldList[fieldIndex-1].FieldRaw)
			if fieldList[fieldIndex-1].IsAutoInc {
				key.IsAutoInc = true
			}
		}
		keyList = append(keyList, key)
	}
	return
}
