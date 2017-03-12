package util

import (
	"time"
	"strconv"
	"github.com/pingplusplus/pingpp-go/pingpp"
	"github.com/pingplusplus/pingpp-go/pingpp/charge"
	"encoding/json"
	"fmt"
	"log"
	"github.com/pingplusplus/pingpp-go/pingpp/refund"
)

func init() {
	// LogLevel 是 Go SDK 提供的 debug 开关
	pingpp.LogLevel = 2
	//设置 API Key
	//pingpp.Key = "sk_live_DiX5WTmDer5G8mz5uLDinvvH"
	pingpp.Key = "sk_test_OWrPOCPm94q5j5qXLO5e1OO4"
	//获取 SDK 版本
	fmt.Println("Go SDK Version:", pingpp.Version())
	//设置错误信息语言，默认是中文
	pingpp.AcceptLanguage = "zh-CN"

	// 直接读取获取直接配置
	//	privateKey, err := ioutil.ReadFile("your_rsa_private_key.pem")
	//	if err != nil {
	//		fmt.Errorf("read failure: %v", err)
	//	}
	//	pingpp.AccountPrivateKey = string(privateKey)

	//设置商户的私钥 记得在Ping++上配置公钥
}

func GenerateOrderno() string{
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

func Pay(channel string, order_id int, order_type string, orderno string, amount uint64) *pingpp.Charge{
	var subject string=""
	var body = "body"
	switch order_type {
		case "COURSE":
			subject = "课程买断"
		case "ASK":
			subject = "提问"
	    	case "PEROID":
			subject = "课时购买"
	    	case "DOC_PACK":
			subject = "资料打包"
	    	case "PEEK":
			subject = "回答抢先看"
	    	case "DOC":
			subject = "资料购买"
	    	default:
			subject = "未知订单"
	}

	//针对metadata字段，可以在每一个 charge 对象中加入订单的一些详情
	metadata := make(map[string]interface{})
	metadata["order_id"] = order_id
	metadata["order_type"] = order_type
	//extra 参数根据渠道不同有区别，下面注释的是一部分的示例
	extra := make(map[string]interface{})
	extra["success_url"] = "http://127.0.0.1:3000/pay?order_id="+strconv.Itoa(order_id)
	params := &pingpp.ChargeParams{
		Order_no:  orderno,
		App:       pingpp.App{Id: "app_jnz9COjH08O4vLaD"},
		Amount:    amount,
		Channel:   channel,
		Currency:  "cny",
		Client_ip: "127.0.0.1",
		Subject:   subject,
		Body:      body,
		Extra:     extra,
		Metadata:  metadata,
	}
	//返回的第一个参数是 charge 对象，你需要将其转换成 json 给客户端，或者客户端接收后转换。
	ch, err := charge.New(params)
	if err != nil {
		errs, _ := json.Marshal(err)
		fmt.Println(string(errs))
		//log.Fatal(err)
		return ch
	}
	chstring, _ := json.Marshal(ch)
	log.Printf("%v\n", string(chstring))
	return ch
}

func GetCharge(ch_id string) *pingpp.Charge{
	ch, err := charge.Get(ch_id)
	if err != nil {
		log.Fatal(err)
	}
	chstring, _ := json.Marshal(ch)
	log.Printf("%v\n", string(chstring))
	return ch
}

func Refund(ch_id string, description string) *pingpp.Refund{
	params := &pingpp.RefundParams{
		Description: description,
	}
	re, err := refund.New(ch_id, params) //ch_id 是已付款的订单号

	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
		return nil
	}
	restring, _ := json.Marshal(re)
	log.Printf("%v\n", string(restring))
	return re
}

func GetRefund(ch_id string, refund_id string) *pingpp.Refund{
	re, err := refund.Get(ch_id, refund_id)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	restring, _ := json.Marshal(re)

	log.Printf("%v\n", string(restring))
	return re
}




