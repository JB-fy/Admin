package api

type SortReq struct {
	Key   string `p:"key" v:"min-length:1"`
	Order string `p:"order" v:"in:asc,desc,ASC,DESC"`
}

type CommonListReq struct {
	Field []string `p:"field" v:"foreach|min-length:1"`
	Sort  SortReq  `p:"sort"`
	Page  uint     `p:"page" v:"min:1"`
	Limit uint     `p:"limit"` //可传0取全部
}

type CommonListFilterReq struct {
	Id        *uint  `c:"id,omitempty" p:"id" v:"min:1"`
	IdArr     []uint `c:"idArr,omitempty" p:"idArr" v:"foreach|min:1"`
	ExcId     *uint  `c:"excId,omitempty" p:"excId" v:"min:1"`
	ExcIdArr  []uint `c:"excIdArr,omitempty" p:"excIdArr" v:"foreach|min:1"`
	StartTime string `c:"startTime,omitempty" p:"startTime" v:"date-format:Y-m-d H:i:s"` //不建议用*gtime.Time类型。传空字符串时，gconv.Map转换会报错
	EndTime   string `c:"endTime,omitempty" p:"endTime" v:"date-format:Y-m-d H:i:s|after-equal:StartTime"`
}

type CommonInfoReq struct {
	Id    uint     `p:"id" v:"required|integer|min:1"`
	Field []string `p:"field" v:"foreach|min-length:1"`
}

type CommonUpdateDeleteIdArrReq struct {
	IdArr []uint `c:"idArr,omitempty" p:"idArr" v:"required|foreach|min:1"`
}
