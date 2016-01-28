package controllers
import (
	"github.com/astaxie/beego"
)

type LogOutController struct {
	beego.Controller
}

func (this *LogOutController)Get(){
	this.TplName = "login.html"
	sess, _ := beego.GlobalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
	defer sess.SessionRelease(this.Ctx.ResponseWriter)
	sess.Delete("userid")
	sess.Delete("username")
}
