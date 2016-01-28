package controllers
import (
	"github.com/astaxie/beego"
	"crypto/md5"
	"dcsolusion/models"
	"encoding/hex"
	"net/http"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController)Get(){
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	this.TplName = "login.html"
	this.Ctx.Request.ParseForm()
	username := this.Ctx.Request.Form.Get("username")
	password := this.Ctx.Request.Form.Get("password")

	md5Maker  := md5.New()
	md5Maker.Write([]byte(password))
	md5Password := hex.EncodeToString(md5Maker.Sum(nil))
	fmt.Println(md5Password)
	userInfo, err := models.GetAdminUser(username)
	if err == orm.ErrNoRows {
		this.Data["login_err"] = true
	}else if err != nil{
		this.Abort(strconv.Itoa(http.StatusServiceUnavailable))
	}else if userInfo.Password == md5Password {
		sess, _ := beego.GlobalSessions.SessionStart(this.Ctx.ResponseWriter, this.Ctx.Request)
		defer sess.SessionRelease(this.Ctx.ResponseWriter)
		sess.Set("userid", userInfo.Id)
		sess.Set("username", userInfo.Name)
		this.Redirect("/admin", http.StatusFound)
	}else{
		this.Data["login_err"] = true
	}
}
