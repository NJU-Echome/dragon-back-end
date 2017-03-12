package filters

import (
	"github.com/NJU-Echome/dragon-back-end/err"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

var FilterLogin = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("userId").(int)
	if !ok {
		// err.RetFilterError(ctx, err.Err404)
		err.RetFilterError(ctx, err.ErrUnauthorized, "未登录")
	}
}
var FilterTeacher = func(ctx *context.Context) {
	value, ok := ctx.Input.Session("type").(int)
	if !ok {
		// err.RetFilterError(ctx, err.Err404)
		err.RetFilterError(ctx, err.ErrUnauthorized, "未登录")
	}
	if value == 0 {
		err.RetFilterError(ctx, err.ErrUnauthorized, "非教师身份")
	}
}

func init() {
	//过滤器
	beego.InsertFilter("/v1/common/user/changePassword", beego.BeforeExec, FilterLogin)
	beego.InsertFilter("/v1/common/user/updateInfo", beego.BeforeExec, FilterLogin)
	beego.InsertFilter("/v1/common/order/generateOrder", beego.BeforeExec, FilterLogin)
	beego.InsertFilter("/v1/common/order/getOrder", beego.BeforeExec, FilterLogin)
	beego.InsertFilter("/v1/teacher/live/detail", beego.BeforeExec, FilterTeacher)
	beego.InsertFilter("/v1/teacher/live/finish", beego.BeforeExec, FilterTeacher)
	beego.InsertFilter("/v1/common/course/freeAskNum", beego.BeforeExec, FilterLogin)
}
