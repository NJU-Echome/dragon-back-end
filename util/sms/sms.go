package sms

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/NJU-Echome/dragon-back-end/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"math/rand"
	"sort"
	"strings"
	"time"
)

//SendSmsCode 发送短信验证码
func SendSmsCode(phonenunmber string) (bool, string) {
	limit, _ := beego.AppConfig.Int("sms_daily_limit")
	todayCount := models.GetTodaySendCount(phonenunmber)
	if int(todayCount) >= limit {
		return false, "超过当日发送次数限制"
	}
	sms_url := beego.AppConfig.String("sms_url")
	sms_type := beego.AppConfig.String("sms_type")
	sms_free_sign_name := beego.AppConfig.String("sms_free_sign_name")
	sms_template_code := beego.AppConfig.String("sms_template_code")
	sms_AppKey := beego.AppConfig.String("sms_AppKey")
	sms_AppSecret := beego.AppConfig.String("sms_AppSecret")
	sms_expire_time, _ := beego.AppConfig.Int("sms_expire_time")
	code := GenCode(6)
	sms_param := "{code:'" + code + "',product:'course+',valid_time:'" + fmt.Sprintf("%d", sms_expire_time) + "'}"
	req := httplib.Post(sms_url)

	m := map[string]string{
		"app_key":                     sms_AppKey,
		"timestamp":                   time.Now().Format("2006-01-02 15:04:05"),
		"v":                           "2.0",
		"method":                      "alibaba.aliqin.fc.sms.num.send",
		"partner_id":                  "top-apitools",
		"format":                      "json",
		"sms_type":                    sms_type,
		"rec_num":                     phonenunmber,
		"sms_free_sign_name":          sms_free_sign_name,
		"sms_template_code":           sms_template_code,
		"force_sensitive_param_fuzzy": "true",
		"sign_method":                 "md5",
		"sms_param":                   sms_param,
	}

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	singString := sms_AppSecret
	for _, k := range keys {
		req.Param(k, m[k])
		singString += k + m[k]
	}
	singString += sms_AppSecret

	signByte := md5.Sum([]byte(singString))
	sign := strings.ToUpper(fmt.Sprintf("%x", signByte))
	req.Param("sign", sign)

	result, _ := req.String()
	var f interface{}
	json.Unmarshal([]byte(result), &f)
	a := f.(map[string]interface{})
	if _, ok := a["error_response"]; ok {
		return false, "短信服务出现问题"
	}
	//发送成功
	models.AddRecord(phonenunmber)
	models.AddSms(phonenunmber, code, sms_expire_time*60)
	return true, ""
}

// VerifyCode 验证
func VerifyCode(phonenumber string, code string) (bool, string) {
	return models.Verify(phonenumber, code)
}

// GenCode  生成随机字符串
func GenCode(len int) string {
	source := []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	code := ""
	for index := 0; index < len; index++ {
		randID := rand.Intn(9)
		code += string(source[randID])
	}
	return code
}
