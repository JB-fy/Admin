package api

import "github.com/gogf/gf/v2/frame/g"

type ReqOrderBy struct {
	Key   string `p:"key" v:"min:1"`
	Order string `p:"order" v:""` //传0取全部
}

type ReqCommonList struct {
	Field []string `p:"field" v:"foreach|min-length:1"`
	Order []string `p:"order"`
	Page  uint     `p:"page" v:"min:1"`
	Limit uint     `p:"limit"` //可传0取全部
}

type ReqCommonListFilter struct {
	Id       *uint  `c:"id,omitempty" p:"id" v:"min:1"`
	IdArr    []uint `c:"idArr,omitempty" p:"idArr" v:"foreach|min:1"`
	ExcId    *uint  `c:"excId,omitempty" p:"excId" v:"min:1"`
	ExcIdArr []uint `c:"excIdArr,omitempty" p:"excIdArr" v:"foreach|min:1"`
	//StartTime *gtime.Time `c:"startTime,omitempty" p:"startTime" v:"date-format:Y-m-d H:i:s"`
	StartTime string `c:"startTime,omitempty" p:"startTime" v:"date-format:Y-m-d H:i:s"`
	EndTime   string `c:"endTime,omitempty" p:"endTime" v:"date-format:Y-m-d H:i:s|after-equal:StartTime"`
}

type ReqTest struct {
	g.Meta `path:"/test" tags:"Test" method:"get" summary:"测试"`
}

type ResTest struct {
	g.Meta   `mime:"text/html" example:"string"`
	UserName string `p:"username"  v:"required|length:4,30#请输入账号|账号长度为:{min}到:{max}位"`
}
