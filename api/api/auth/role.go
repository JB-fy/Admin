package api

import (
	apiCommon "api/api"
)

type RoleListReq struct {
	apiCommon.CommonListReq
	Filter RoleListFilterReq `p:"filter"`
}

/* type Role struct {
    RoleId     uint        `json:"roleId"     `// 权限角色ID
    SceneId    uint        `json:"sceneId"    `// 权限场景ID
    TableId    uint        `json:"tableId"    `// 关联表ID（0表示平台创建，其他值根据authSceneId对应不同表，表示是哪个表内哪个机构或个人创建）
    RoleName   string      `json:"roleName"   `// 名称
    IsStop     uint        `json:"isStop"     `// 是否停用：0否 1是
    UpdateTime *gtime.Time `json:"updateTime" `// 更新时间
    CreateTime *gtime.Time `json:"createTime" `// 创建时间
} */

type RoleListFilterReq struct {
	apiCommon.CommonListFilterReq `c:",omitempty"`
	RoleId                        *uint  `c:"roleId,omitempty" p:"roleId" v:"min:1"`
	RoleName                      string `c:"roleName,omitempty" p:"roleName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	IsStop                        *uint  `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
}

type RoleInfoReq struct {
	apiCommon.CommonInfoReq
}

type RoleCreateReq struct {
	RoleName *string `c:"roleName,omitempty" p:"roleName" v:"required|length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	IsStop   *uint   `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
}

type RoleUpdateReq struct {
	apiCommon.CommonUpdateDeleteIdArrReq `c:",omitempty"`
	RoleName                             *string `c:"roleName,omitempty" p:"roleName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	IsStop                               *uint   `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
}

type RoleDeleteReq struct {
	apiCommon.CommonUpdateDeleteIdArrReq
}
