// 解决dao层互相引用造成的import cycle not allowed问题。步骤如下：
/*
步骤一：在api/internal/dao/org/admin.go中写入以下代码
import "api/internal/dao/org/allow"

func init() {
	allow.RegisterAdmin(&Admin)
}

步骤二：api/internal/dao/下的其它目录，如需导入api/internal/dao/org，则改成导入api/internal/dao/org/allow
import daoAdmin "api/internal/dao/org/allow"
*/
package allow

import (
	"api/internal/dao"
	"api/internal/dao/org/internal"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
)

type DaoAdmin interface {
	dao.DaoInterface
	Columns() *internal.AdminColumns
	JoinLoginName(orgId uint, loginName string) string
	GetLoginName(loginName string) string
	CacheGetInfo(ctx context.Context, id uint) (info gdb.Record, err error)
}

var Admin DaoAdmin

func RegisterAdmin(d DaoAdmin) {
	Admin = d
}
