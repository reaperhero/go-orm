package example


import (
	"database/sql"
	"fmt"
	_ "gitee.com/chunanyong/dm"
	"github.com/jmoiron/sqlx"
	"testing"
	"time"
)

const (
	driverName = "dm"
)

type Agent struct {
	Id             string     `db:"ID" orm:"column(id);pk;auto"`
	SidecarId      string     `db:"sidecar_id" orm:"column(sidecar_id)"`
	Type           int        `db:"type" orm:"column(type)"`
	Name           string     `db:"name" orm:"column(name)"`
	Version        string     `db:"version" orm:"column(version)"`
	IsUninstalled  int        `db:"is_uninstalled" orm:"column(is_uninstalled)"`
	DeployDate     *time.Time `db:"deploy_date" orm:"column(deploy_date)"`
	AutoDeployment int        `db:"auto_deployment" orm:"column(auto_deployment)"`
	LastUpdateDate *time.Time `db:"last_update_date" orm:"column(last_update_date)"`
	AutoUpdated    int        `db:"auto_updated" orm:"column(auto_updated)"`
}

func TestDMExec(t *testing.T)  {
	open, err := sql.Open(driverName, "dm://SYSDBA:SYSDBA001@172.16.104.165:5236")
	if err!=nil{
		fmt.Println(err)
		return
	}
	sqlxDB := sqlx.NewDb(open, driverName) // returns *sqlx.DB
	var info Agent
	err = sqlxDB.Get(&info, `select * from "AGENT_LIST" limit 1;`)


	if err!=nil{
		fmt.Println(err)
	}

	re ,err := sqlxDB.Exec(`insert into AGENT_LIST values("06e1b337-1ba0-462b-b6a8-db4015e97487","2d277758-1c7a-4780-afa8-4e9f418bbafe",0,"tengine,0,null,0,null;,0)`)

	if err!=nil{
		fmt.Println(err)
		id, err := re.LastInsertId()
		if err != nil {
			return
		}
		fmt.Println(err,id)
	}

	err = sqlxDB.Select(&info, `select * from AGENT_LIST limit 1;`)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(info)
}


