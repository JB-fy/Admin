package api

type ReqCommonList struct {
	Field []string `p:"field"  v:"foreach|min-length:1"`
	Order []string `p:"order"  v:""`
	Page  uint     `p:"page"  v:"min:1"`
	Limit uint     `p:"limit"  v:""` //传0取全部
}

type ReqCommonListFilter struct {
	Id        *uint  `c:"id,omitempty" p:"id" v:"min:1"`
	IdArr     []uint `c:"idArr,omitempty" p:"idArr" v:"foreach|min:1"`
	ExcId     *uint  `c:"excId,omitempty" p:"excId" v:"min:1"`
	ExcIdArr  []uint `c:"excIdArr,omitempty" p:"excIdArr" v:"foreach|min:1"`
	StartTime string `c:"startTime,omitempty" p:"startTime" v:"date-format:Y-m-d H:i:s"`
	EndTime   string `c:"endTime,omitempty" p:"endTime" v:"date-format:Y-m-d H:i:s|after-equal:StartTime"`
}
