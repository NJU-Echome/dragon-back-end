package base

import (
	"github.com/NJU-Echome/dragon-back-end/err"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) RetError(e *err.Error, moreInfo string) {
	if mode := beego.AppConfig.String("runmode"); mode == "prod" {
		e.DevInfo = ""
	}
	//more info
	e.MoreInfo = moreInfo

	errorResult := make(map[string]err.Error)
	errorResult["error"] = *e

	c.Data["json"] = errorResult
	c.Ctx.ResponseWriter.WriteHeader(e.Status)
	c.ServeJSON()
	c.Abort("")
}

func (c *BaseController) RetData(data interface{}) {
	result := map[string]interface{}{
		"data": data,
	}
	c.Data["json"] = result
	c.ServeJSON()
	c.Abort("")
}
