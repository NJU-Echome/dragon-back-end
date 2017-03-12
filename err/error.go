package err

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type Error struct {
	Status   int    `json:"status"`
	Code     int    `json:"code"`
	Message  string `json:"message"`
	DevInfo  string `json:"dev_info"`
	MoreInfo string `json:"more_info"`
}

var (
	Err404       = &Error{404, 404, "page not found", "page not found", ""}
	ErrInputData = &Error{200, 10001, "数据输入错误", "客户端参数错误", ""}
	ErrDatabase  = &Error{200, 10002, "服务器错误", "数据库操作错误", ""}
	ErrDupUser   = &Error{200, 10003, "用户信息已存在", "数据库记录重复", ""}
	ErrNoUser    = &Error{200, 10004, "用户信息不存在", "数据库记录不存在", ""}
	// ErrPass         = &Error{400, 10005, "用户信息不存在或密码不正确", "密码不正确", ""}
	ErrNoUserPass   = &Error{400, 10006, "用户信息不存在或密码不正确", "数据库记录不存在或密码不正确", ""}
	ErrNoUserChange = &Error{400, 10007, "用户信息不存在或数据未改变", "数据库记录不存在或数据未改变", ""}
	ErrInvalidUser  = &Error{400, 10008, "用户信息不正确", "Session信息不正确", ""}
	ErrOpenFile     = &Error{500, 10009, "服务器错误", "打开文件出错", ""}
	ErrWriteFile    = &Error{500, 10010, "服务器错误", "写文件出错", ""}
	ErrSystem       = &Error{500, 10011, "服务器错误", "操作系统错误", ""}
	ErrExpired      = &Error{400, 10012, "登录已过期", "验证token过期", ""}
	ErrPermission   = &Error{403, 10013, "没有权限", "没有操作权限", ""}
	ErrSendMessage  = &Error{200, 10014, "短信发送失败", "超过一天发送限制或短信服务出现问题", ""}
	ErrVerifyCode   = &Error{200, 10015, "验证失败", "验证码错误", ""}
	ErrUnauthorized = &Error{401, 10016, "需要登录", "未进行身份认证", ""}
	ErrGetCharge    = &Error{200, 10017, "获取charge失败", "输入charge_id错误", ""}
	ErrNil          = &Error{200, 10018, "获取数据失败", "空指针", ""}
)

func RetFilterError(ctx *context.Context, e *Error, moreInfo string) {
	if mode := beego.AppConfig.String("runmode"); mode == "prod" {
		e.DevInfo = ""
	}
	//more info
	e.MoreInfo = moreInfo

	ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	ctx.ResponseWriter.WriteHeader(e.Status)

	errorResult := make(map[string]Error)
	errorResult["error"] = *e

	hasIndent := true
	runmode := beego.AppConfig.String("runmode")
	if runmode == beego.PROD {
		hasIndent = false
	}
	ctx.Output.JSON(errorResult, hasIndent, false)
}
