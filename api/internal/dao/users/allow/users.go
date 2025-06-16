// 解决dao层互相引用造成的import cycle not allowed问题。步骤如下：
/*
步骤一：在api/internal/dao/users/users.go中写入以下代码
import "api/internal/dao/users/allow"

func init() {
	allow.RegisterUsers(&Users)
}

步骤二：api/internal/dao/下的其它目录，如需导入api/internal/dao/users，则改成导入api/internal/dao/users/allow
import daoUsers "api/internal/dao/users/allow"
*/
package allow

import (
	"api/internal/dao"
	"api/internal/dao/users/internal"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
)

type DaoUsers interface {
	dao.DaoInterface
	Columns() *internal.UsersColumns
	CacheGetInfo(ctx context.Context, id uint) (info gdb.Record, err error)
}

var Users DaoUsers

func RegisterUsers(d DaoUsers) {
	Users = d
}
