package internal

import (
	"api/internal/utils/cql/model"
	"context"
	"fmt"
	"strings"

	"github.com/gocql/gocql"
	"github.com/gogf/gf/v2/frame/g"
)

func InitSession(ctx context.Context, config *model.Config) (session *gocql.Session, err error) {
	/*--------数据库配置 开始--------*/
	cluster := gocql.NewCluster(config.Hosts...)
	if config.UserName != `` && config.Password != `` {
		cluster.Authenticator = gocql.PasswordAuthenticator{
			Username: config.UserName,
			Password: config.Password,
			// AllowedAuthenticators: []string{},
		}
	}
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.DCAwareRoundRobinPolicy(config.DcName))
	if config.ProtoVersion != nil { //默认0。驱动程序将自动尝试发现集群支持的最高协议版本
		cluster.ProtoVersion = *config.ProtoVersion
	}
	if config.Consistency != nil { //一致性级别。默认gocql.Quorum，至少3个节点同意
		cluster.Consistency = gocql.Consistency(*config.Consistency)
	}
	if config.NumConns != nil { //每个节点的连接数
		cluster.NumConns = *config.NumConns
	}
	if config.Debug { //日志记录
		observer := &model.Observer{Config: config, Log: g.Log(`cql`)}
		cluster.QueryObserver = observer
		cluster.BatchObserver = observer
	}
	/*--------数据库配置 结束--------*/

	cluster.Keyspace = config.Keyspace
	session, err = cluster.CreateSession()
	if err == nil {
		/* //修改键空间配置（不建议）。键空间配置一般不修改，而修改则必须在服务器上执行命令进行修复，使数据重新分布：nodetool repair -full 键空间名
		replType := `SimpleStrategy`
		replDcFormat := `'%s': %d`
		replDcList := []string{fmt.Sprintf(replDcFormat, `replication_factor`, config.DcList[0].ReplNum)}
		if len(config.DcList) > 1 || config.DcList[0].DcName != `` {
			replType = `NetworkTopologyStrategy`
			replDcList = []string{}
			for _, v := range config.DcList {
				replDcList = append(replDcList, fmt.Sprintf(replDcFormat, v.DcName, v.ReplNum))
			}
		}
		session.ExecStmt(fmt.Sprintf(`ALTER KEYSPACE %s WITH REPLICATION = {'class': '%s', %s}`, config.Keyspace, replType, strings.Join(replDcList, `, `))) */
		return
	}

	// 键空间不存在也会报错，故尝试创建键空间，重连数据库
	if config.Keyspace == `` || len(config.DcList) == 0 { //这两个未配置时，无法创建键空间，直接返回报错
		return
	}

	// 创建无键空间的连接（创建键空间用）
	cluster.Keyspace = ``
	cluster.PoolConfig.HostSelectionPolicy = gocql.DCAwareRoundRobinPolicy(config.DcName)
	sessionTmp, err := cluster.CreateSession()
	if err != nil {
		return
	}
	defer sessionTmp.Close()

	// 创建键空间
	replType := `SimpleStrategy`
	replDcFormat := `'%s': %d`
	replDcList := []string{fmt.Sprintf(replDcFormat, `replication_factor`, config.DcList[0].ReplNum)}
	if len(config.DcList) > 1 || config.DcList[0].DcName != `` {
		replType = `NetworkTopologyStrategy`
		replDcList = []string{}
		for _, v := range config.DcList {
			replDcList = append(replDcList, fmt.Sprintf(replDcFormat, v.DcName, v.ReplNum))
		}
	}
	err = sessionTmp.Query(fmt.Sprintf(`CREATE KEYSPACE IF NOT EXISTS %s WITH REPLICATION = {'class': '%s', %s}`, config.Keyspace, replType, strings.Join(replDcList, `, `))).Exec()
	if err != nil {
		return
	}

	// 重连数据库
	cluster.Keyspace = config.Keyspace
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.DCAwareRoundRobinPolicy(config.DcName))
	session, err = cluster.CreateSession()
	return
}
