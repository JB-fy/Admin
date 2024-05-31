package my_gen

import (
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

type myGenI18n struct {
	list [][2]string
}

type myGenI18nField struct {
	item [2]string
}

func (i18nThis *myGenI18n) Add(i18nField myGenI18nField) {
	if i18nField.item[0] != `` {
		i18nThis.list = append(i18nThis.list, i18nField.item)
	}
}

func (i18nThis *myGenI18n) Merge(i18nOther myGenI18n) {
	i18nThis.list = append(i18nThis.list, i18nOther.list...)
}

func (i18nThis *myGenI18n) Unique() {
	keyArr := garray.NewStrArray()
	listTmp := [][2]string{}
	for _, v := range i18nThis.list {
		if !keyArr.Contains(v[0]) {
			keyArr.Append(v[0])
			listTmp = append(listTmp, v)
		}
	}
	i18nThis.list = listTmp
}

// 自动生成I18n
func genI18n(i18n myGenI18n) {
	saveFile := gfile.SelfDir() + `/resource/i18n/zh-cn/name.yaml`
	tplI18n := gfile.GetContents(saveFile)

	i18nAppend := []string{}
	for _, item := range i18n.list {
		if gstr.Pos(tplI18n, item[0]) == -1 {
			i18nAppend = append(i18nAppend, item[0]+`: "`+item[1]+`"`)
		}
	}

	if len(i18nAppend) > 0 {
		tplI18n += gstr.Join(append([]string{``}, i18nAppend...), `
`)
		gfile.PutContents(saveFile, tplI18n)
	}
}
