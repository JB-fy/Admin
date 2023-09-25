package allow

// 解决dao层互相引用造成的import cycle not allowed问题。示例如下
/*
1：在api/internal/dao/user/user.go中写入以下代码
import (
	"api/internal/dao/user/allow"
)

func init() {
	allow.RegisterUser(&User)
}

2：api/internal/dao/user目录导入过的dao层，如需导入api/internal/dao/user，则改成导入api/internal/dao/user/allow
import (
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
		Columns() internal.UserColumns
	}
)

var (
	User DaoUser
)

func RegisterUser(d DaoUser) {
	User = d
} */
