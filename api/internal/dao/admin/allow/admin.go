// 解决dao层互相引用造成的import cycle not allowed问题。步骤如下：
/*
步骤一：在api/internal/dao/admin/admin.go中写入以下代码
import "api/internal/dao/admin/allow"

func init() {
	allow.RegisterAdmin(&Admin)
}

步骤二：api/internal/dao/下的其它目录，如需导入api/internal/dao/admin，则改成导入api/internal/dao/admin/allow
import daoAdmin "api/internal/dao/admin/allow"
*/
package allow

import (
	"api/internal/dao"
	"api/internal/dao/admin/internal"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
)

type DaoAdmin interface {
	dao.DaoInterface
	Columns() *internal.AdminColumns
	CacheGetInfo(ctx context.Context, id uint) (info gdb.Record, err error)
	JoinLoginName(orgId uint, isSuper uint8, loginName string) string
	GetLoginName(loginName string) string
}

var Admin DaoAdmin

func RegisterAdmin(d DaoAdmin) {
	Admin = d
}
