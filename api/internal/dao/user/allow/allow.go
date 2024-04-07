package allow

// 解决dao层互相引用造成的import cycle not allowed问题。示例如下
/*
步骤一：在api/internal/dao/user/user.go中写入以下代码
import (
	"api/internal/dao/user/allow"
)

func init() {
	allow.RegisterUser(&User)
}

步骤二：api/internal/dao/下的其它目录，如需导入api/internal/dao/user，则改成导入api/internal/dao/user/allow
import (
	// daoUser "api/internal/dao/user"
	daoUser "api/internal/dao/user/allow"
)
*/

/* import (
	"api/internal/dao"
	"api/internal/dao/user/internal"
)

type (
	DaoUser interface {
		dao.DaoInterface
		Columns() *internal.UserColumns
	}
)

var (
	User DaoUser
)

func RegisterUser(d DaoUser) {
	User = d
} */
