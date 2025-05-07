package main

import (
	"log"
	"teamotp/internal/models"

	"github.com/beego/beego/v2/client/orm"
)

func init_db(path string) {
	orm.RegisterModel(new(models.OTPCode))
	orm.RegisterDataBase("default", "sqlite3", path)

	if err := orm.RunSyncdb("default", false, true); err != nil {
		log.Fatal(err)
	}
}
