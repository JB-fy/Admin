package my_gen

import (
	"slices"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

type myGenI18n struct {
	list []myGenI18nField
}

type myGenI18nField struct {
	key string
	val string
}

func (i18nThis *myGenI18n) Add(i18nField myGenI18nField) {
	if i18nField.key != `` {
		i18nThis.list = append(i18nThis.list, i18nField)
	}
}

func (i18nThis *myGenI18n) Merge(i18nOther myGenI18n) {
	i18nThis.list = append(i18nThis.list, i18nOther.list...)
}

func (i18nThis *myGenI18n) Unique() {
	var keyArr []string
	listTmp := []myGenI18nField{}
	for _, v := range i18nThis.list {
		if !slices.Contains(keyArr, v.key) {
			keyArr = append(keyArr, v.key)
			listTmp = append(listTmp, v)
		}
	}
	i18nThis.list = listTmp
}

// 自动生成I18n
func genI18n(i18n myGenI18n) {
	saveFile := gfile.SelfDir() + `/manifest/i18n/zh-cn/name.yaml`
	tplI18n := gfile.GetContents(saveFile)

	i18nAppend := []string{}
	for _, v := range i18n.list {
		if gstr.Pos(tplI18n, v.key) == -1 {
			i18nAppend = append(i18nAppend, v.key+`: "`+v.val+`"`)
		}
	}

	if len(i18nAppend) > 0 {
		tplI18n += gstr.Join(append([]string{``}, i18nAppend...), `
`)
		gfile.PutContents(saveFile, tplI18n)
	}
}
