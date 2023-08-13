package example

import (
	"fmt"
	_ "gitee.com/chunanyong/dm"
	"github.com/reaperhero/go-orm/orm"
	"testing"
)

func init() {
	orm.RegisterModel(new(Student),new(Post),new(Profile),new(Tag))
	orm.RegisterDriver("dm", orm.DRDM)
	// ?logLevel=all

	orm.RegisterDataBase("default", "dm", "dm://SYSDBA:SYSDBA001@172.16.104.165:5236")
	orm.Debug = true
}

var o orm.Ormer

func TestName(t *testing.T)  {
	err := orm.RunSyncdb("default", true, true)
	if err != nil {
		return
	}
	
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

func builder() {
	qb, _ := orm.NewQueryBuilder("dm")
	qb.Select("t1.id", "t2.name").From("deploy_host01 t2").LeftJoin("deploy_host02 t1").On(`t1."instance_id" = t2."id"`).And(`"t1.instance_id = t2.id"`).
		Where(`"t1.id" = "12345678"`).And(`"t2"."is_deleted" = 0`).OrderBy("id").Desc().Limit(10).Offset(20)
	fmt.Println(qb.String())
	// SELECT "t1"."id", "t2"."name" FROM "deploy_host01" as t2 LEFT JOIN "deploy_host02" as t1 ON t1."instance_id" = t2."id" AND "t1.instance_id = t2.id" WHERE "t1.id" = "12345678" AND "t2"."is_deleted" = 0 ORDER BY "id" DESC LIMIT 10 OFFSET 20
}
