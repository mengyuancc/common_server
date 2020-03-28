/**
 * 获取数据库连接
 */
package datasource

import (
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"

	"iris_server/conf"
)

var (
	masterEngine *xorm.Engine
	slaveEngine  *xorm.Engine
	lock         sync.Mutex
)

// 主库，单例
func InstanceMaster() *xorm.Engine {
	if masterEngine != nil {
		return masterEngine
	}
	lock.Lock()
	defer lock.Unlock()
	if masterEngine != nil {
		return masterEngine
	}
	c := conf.MasterDbConfig
	driveSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		c.User, c.Pwd, c.Host, c.Port, c.DbName)
	masterEngine, err := xorm.NewEngine(conf.DriverName, driveSource)
	if err != nil {
		log.Fatal("Master DB Connect Fail,", err)
		return nil
	}
	// 开启Debug模式
	masterEngine.ShowSQL(true)
	// 设置时区
	masterEngine.SetTZLocation(conf.SysTimeLocation)
	// 性能优化的时候才考虑，加上本机的SQL缓存
	//cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	//engine.SetDefaultCacher(cacher)

	return masterEngine
}

// 从库，单例
func InstanceSlave() *xorm.Engine {
	if slaveEngine != nil {
		return slaveEngine
	}
	lock.Lock()
	defer lock.Unlock()

	if slaveEngine != nil {
		return slaveEngine
	}
	c := conf.SlaveDbConfig
	slaveEngine, err := xorm.NewEngine(conf.DriverName,
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
			c.User, c.Pwd, c.Host, c.Port, c.DbName))
	if err != nil {
		log.Fatal("Salve DB Connect Fail", err)
		return nil
	}
	slaveEngine.SetTZLocation(conf.SysTimeLocation)

	return slaveEngine
}
