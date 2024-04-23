package internal

import "github.com/gogf/gf/v2/text/gstr"

type MyGenFieldStyle = string

const (
	FieldStyleCaseSnake      MyGenFieldStyle = `CaseSnake`      //蛇形
	FieldStyleCaseCamelLower MyGenFieldStyle = `CaseCamelLower` //小驼峰
	FieldStyleCaseCamel      MyGenFieldStyle = `CaseCamel`      //大驼峰
	FieldStyleCaseKebab      MyGenFieldStyle = `CaseKebab`      //大驼峰
)

func GetFieldStyle(fieldList []MyGenField) MyGenFieldStyle {
	for _, v := range fieldList {
		if gstr.CaseSnake(v.FieldRaw) != v.FieldRaw {
			return FieldStyleCaseCamelLower
		}
	}
	return FieldStyleCaseSnake
}

// 将字符串转换成对应的风格后返回
func GetStrByFieldStyle(fieldStyle MyGenFieldStyle, strRaw string, fixArr ...string) (strNew string) {
	strNew = gstr.TrimStr(gstr.CaseSnake(strRaw), `_`)
	prefix := ``
	suffix := ``
	if len(fixArr) == 1 {
		prefix = fixArr[0]
	} else if len(fixArr) > 1 {
		prefix = fixArr[0]
		suffix = fixArr[1]
	}
	if prefix != `` {
		strNew = gstr.TrimStr(gstr.CaseSnake(prefix), `_`) + `_` + strNew
	}
	if suffix != `` {
		strNew = strNew + `_` + gstr.TrimStr(gstr.CaseSnake(suffix), `_`)
	}
	switch fieldStyle {
	// case FieldStyleCaseSnake:
	case FieldStyleCaseCamelLower:
		strNew = gstr.CaseCamelLower(strNew)
	case FieldStyleCaseCamel:
		strNew = gstr.CaseCamel(strNew)
	case FieldStyleCaseKebab:
		strNew = gstr.CaseKebab(strNew)
	}
	return
}
