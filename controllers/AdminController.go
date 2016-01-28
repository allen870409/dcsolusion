package controllers
import (
	"github.com/astaxie/beego"
	"net/http"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) Prepare() {
	sess, _ := beego.GlobalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
	sess_userid := sess.Get("userid")
	sess_username := sess.Get("username")
	if sess_userid == nil {
		this.Ctx.Redirect(http.StatusFound, "/login")
		return
	}
	this.Data["username"] = sess_username
}

func (this *AdminController)Get(){
	this.TplName = "admin.html"
}
