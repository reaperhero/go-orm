package example

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/reaperhero/go-orm/orm"
	"testing"
)


//func init() {
//	// register model
//	orm.RegisterModel(new(Student),new(Post),new(Profile),new(Tag))
//
//	// set default database
//	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/dtagent?charset=utf8")
//}

func TestMysql(t *testing.T) {
	err := orm.RunSyncdb("default", true, true)
	if err != nil {
		return
	}
	//db := orm.NewOrmUsingDB("mysqlOrm")
}
