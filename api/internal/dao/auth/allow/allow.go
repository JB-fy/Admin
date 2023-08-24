package allow

// 解决dao层互相引用造成的import cycle not allowed问题
/* //下面方法写入dao层文件中，其中一方使用import "api/internal/dao/auth/allow"即可
func init() {
	dao.RegisterAuthScene(&Scene)
} */
import (
	"api/internal/dao"
	"api/internal/dao/auth/internal"
)

type (
	DaoScene interface {
		dao.DaoInterface
		Columns() internal.SceneColumns
	}
	DaoRoleRelOfPlatformAdmin interface {
		dao.DaoInterface
		Columns() internal.RoleRelOfPlatformAdminColumns
	}
)

var (
	Scene                  DaoScene
	RoleRelOfPlatformAdmin DaoRoleRelOfPlatformAdmin
)

func RegisterScene(d DaoScene) {
	Scene = d
}

func RegisterRoleRelOfPlatformAdmin(d DaoRoleRelOfPlatformAdmin) {
	RoleRelOfPlatformAdmin = d
}
