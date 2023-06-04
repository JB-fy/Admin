package api

import (
	apiCommon "api/api"
)

type MenuListReq struct {
	apiCommon.CommonListReq `c:",omitempty"`
	Filter                  MenuListFilterReq `p:"filter"`
}

/* type Menu struct {
    MenuId     uint        `json:"menuId"     `// 权限菜单ID
    SceneId    uint        `json:"sceneId"    `// 权限场景ID（只能是auth_scene表中sceneType为0的菜单类型场景）
    Pid        uint        `json:"pid"        `// 父ID
    MenuName   string      `json:"menuName"   `// 名称
    MenuIcon   string      `json:"menuIcon"   `// 图标
    MenuUrl    string      `json:"menuUrl"    `// 链接
    Level      uint        `json:"level"      `// 层级
    PidPath    string      `json:"pidPath"    `// 层级路径
    ExtraData  string      `json:"extraData"  `// 额外数据。（json格式：{"i18n（国际化设置）": {"title": {"语言标识":"标题",...}}）
    Sort       uint        `json:"sort"       `// 排序值（从小到大排序，默认50，范围0-100）
    IsStop     uint        `json:"isStop"     `// 是否停用：0否 1是
    UpdateTime *gtime.Time `json:"updateTime" `// 更新时间
    CreateTime *gtime.Time `json:"createTime" `// 创建时间
} */

type MenuListFilterReq struct {
	apiCommon.CommonListFilterReq `c:",omitempty"`
	// 下面根据自己需求修改
	IsStop   *uint  `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
	MenuId   *uint  `c:"menuId,omitempty" p:"menuId" v:"min:1"`
	MenuName string `c:"menuName,omitempty" p:"menuName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
}

type MenuInfoReq struct {
	apiCommon.CommonInfoReq `c:",omitempty"`
}

type MenuCreateReq struct {
	SceneId   *uint   `c:"sceneId,omitempty" p:"sceneId" v:"required|min:1"`
	Pid       *uint   `c:"sceneId,omitempty" p:"sceneId" v:"min:0"`
	MenuName  *string `c:"menuName,omitempty" p:"menuName" v:"required|length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	MenuIcon  *string `c:"menuIcon,omitempty" p:"menuIcon" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	MenuUrl   *string `c:"menuUrl,omitempty" p:"menuUrl" v:"length:1,120"`
	ExtraData *string `c:"extraData,omitempty" p:"extraData" v:"json"`
	Sort      *uint   `c:"sort,omitempty" p:"sort" v:"between:0,100"`
	IsStop    *uint   `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
}

type MenuUpdateReq struct {
	apiCommon.CommonUpdateDeleteIdArrReq `c:",omitempty"`
	SceneId                              *uint   `c:"sceneId,omitempty" p:"sceneId" v:"min:1"`
	Pid                                  *uint   `c:"sceneId,omitempty" p:"sceneId" v:"min:0"`
	MenuName                             *string `c:"menuName,omitempty" p:"menuName" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	MenuIcon                             *string `c:"menuIcon,omitempty" p:"menuIcon" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	MenuUrl                              *string `c:"menuUrl,omitempty" p:"menuUrl" v:"length:1,120"`
	ExtraData                            *string `c:"extraData,omitempty" p:"extraData" v:"json"`
	Sort                                 *uint   `c:"sort,omitempty" p:"sort" v:"between:0,100"`
	IsStop                               *uint   `c:"isStop,omitempty" p:"isStop" v:"in:0,1"`
}

type MenuDeleteReq struct {
	apiCommon.CommonUpdateDeleteIdArrReq `c:",omitempty"`
}
