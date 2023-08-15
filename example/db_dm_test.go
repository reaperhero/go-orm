package example

import (
	"fmt"
	_ "gitee.com/chunanyong/dm"
	"github.com/reaperhero/go-orm/orm"
	"testing"
)
// dm tips
// 关联查询，必须要有on条件
// 查询时，值是字符串，必须要用单引号
// curd操作，表和字段需要用双引号, 字段别名和表别名 不需要加双引号
// 查询系统自带表或者字段，不用加引号





func init() {
	orm.RegisterModel(new(Student),new(Post),new(Profile),new(Tag))
	orm.RegisterDriver("dm", orm.DRDM)
	// ?logLevel=all

	err := orm.RegisterDataBase(alias, "dm", "dm://SYSDBA:SYSDBA001@172.16.104.165:5236")
	if err!=nil{
		panic(err)
	}
	orm.Debug = true
	orm.SetMaxIdleConns(alias, 10) // 设置数据库最大空闲连接数
	orm.SetMaxOpenConns(alias, 50) // 设置数据库最大连接数

}

var o orm.Ormer

const  alias = "default"

func TestName(t *testing.T)  {
	o = orm.NewOrmUsingDB(alias)
	runSyncDB()
	insert()
	builder()
}


func runSyncDB()  {
	err := orm.RunSyncdb("default", true, true)
	if err != nil {
		return
	}
}
func builder() {
	var maps []orm.ParamsList
	builder, _ := orm.NewQueryBuilder("dm")
	builder = builder.Select("deploy_product_list.*").
		From("deploy_cluster_smooth_upgrade_product_rel").
		LeftJoin("deploy_product_list").
		On(`"deploy_product_list"."id"="deploy_cluster_smooth_upgrade_product_rel"."pid"`).
		Where(`"pid" = ? AND "deploy_cluster_smooth_upgrade_product_rel"."is_deleted"=0`)
	list, err := o.Raw(builder.String()).SetArgs(1).ValuesList(&maps)
	if err != nil {
		panic(err)
	}
	fmt.Println(list)
}

func insert()  {
	o = orm.NewOrmUsingDB("default")
	student := &Student{
		Name:  "chenqiangjun",
		Email: "chenqiangjun@com",
	}
	profile := &Profile{
		Age:     10,
		Student: student,
	}
	insert, err := o.Insert(profile)
	if err != nil {
		panic(err)
	}
	fmt.Println(insert)

	student.Profile = profile
	insert, err = o.Insert(student)
	if err != nil {
		panic(err)
	}

	fmt.Println(insert)
}
