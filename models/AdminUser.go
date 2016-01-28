package models
import (
	"github.com/astaxie/beego/orm"
)

type AdminUser struct {
	Id int
	Name string
	Password string
}

func GetAdminUser(username string)(adminUser *AdminUser, err error){
	o := orm.NewOrm()
	adminUser = &AdminUser{Name: username}
	err = o.Read(adminUser, "Name")
	return
}
