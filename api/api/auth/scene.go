package api

type ReqList struct {
	Field    []string `p:"field"  v:"required|length:4,30"`
	Filter   []string `p:"field"  v:"required|length:4,30"`
	UserName string   `p:"username"  v:"required|length:4,30#请输入账号|账号长度为:{min}到:{max}位"`
}
