package api

import "github.com/gogf/gf/v2/os/gtime"

type SortReq struct {
	Key   string `p:"key" v:"min-length:1"`
	Order string `p:"order" v:"in:asc,desc,ASC,DESC"`
}

type CommonListReq struct {
	Field []string `p:"field" v:"distinct|foreach|min-length:1"`
	Sort  SortReq  `p:"sort"`
	Page  int      `p:"page" v:"integer|min:1"`
	Limit *int     `p:"limit" v:"integer|min:0"` //可传0取全部
}

type CommonListFilterReq struct {
	Id        *uint       `c:"id,omitempty" p:"id" v:"integer|min:1"`
	IdArr     []uint      `c:"idArr,omitempty" p:"idArr" v:"distinct|foreach|integer|foreach|min:1"`
	ExcId     *uint       `c:"excId,omitempty" p:"excId" v:"integer|min:1"`
	ExcIdArr  []uint      `c:"excIdArr,omitempty" p:"excIdArr" v:"distinct|foreach|integer|foreach|min:1"`
	StartTime *gtime.Time `c:"startTime,omitempty" p:"startTime" v:"date-format:Y-m-d H:i:s"`
	EndTime   *gtime.Time `c:"endTime,omitempty" p:"endTime" v:"date-format:Y-m-d H:i:s|after-equal:StartTime"`
	Name      string      `c:"name,omitempty" p:"name" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
}

type CommonInfoReq struct {
	Id    uint     `p:"id" v:"required|integer|min:1"`
	Field []string `p:"field" v:"distinct|foreach|min-length:1"`
}

type CommonUpdateDeleteIdArrReq struct {
	IdArr []uint `c:"idArr,omitempty" p:"idArr" v:"required|distinct|foreach|integer|foreach|min:1"`
}
