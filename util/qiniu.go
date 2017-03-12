package util

import (
	"github.com/astaxie/beego"
	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
)

func GetResourceUrl(domain string, key string, isPrivate bool) string {
	if isPrivate {
		conf.ACCESS_KEY = beego.AppConfig.String("qiniu_access_key")
		conf.SECRET_KEY = beego.AppConfig.String("qiniu_secret_key")
		return downloadUrl(domain, key)
	} else {
		return "http://" + domain + "/" + key
	}
}

func downloadUrl(domain string, key string) string {
	// 调用MakeBaseUrl()方法将domain,key处理成http://domain/key的形式
	baseUrl := kodo.MakeBaseUrl(domain, key)
	policy := kodo.GetPolicy{}
	// 生成一个client对象
	c := kodo.New(0, nil)
	// 调用MakePrivateUrl方法返回url
	return c.MakePrivateUrl(baseUrl, &policy)
}
