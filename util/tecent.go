package util

import (
	"fmt"
	"time"
)

func GetPushUrl(bizId int, liveId int, time time.Time) (string, string) {
	pushUrl := "rtmp://"
	bizIdStr := fmt.Sprintf("%d", bizId)
	liveIdStr := fmt.Sprintf("%d", liveId)
	streamId := Md5(liveIdStr)[0:7] + "_" + liveIdStr
	liveCode := bizIdStr + "_" + streamId
	timeStr := fmt.Sprintf("%x", time.Unix())
	key := "00ebfd9b01e033cf1bab72ab110b7715"
	txSecret := Md5(key + liveCode + timeStr)
	pushUrl = pushUrl + bizIdStr + ".livepush.myqcloud.com/live/" + liveCode + "?bizid=" + bizIdStr + "&txSecret=" + txSecret + "&txTime=" + timeStr + "&record=flv&record_interval=600"
	return pushUrl, streamId
}

func GetPlayUrl(bizId int, liveId int) string {
	liveUrl := "http://"
	bizIdStr := fmt.Sprintf("%d", bizId)
	liveIdStr := fmt.Sprintf("%d", liveId)
	streamId := Md5(liveIdStr)[0:7] + "_" + liveIdStr
	liveCode := bizIdStr + "_" + streamId
	liveUrl += bizIdStr + ".liveplay.myqcloud.com/live/" + liveCode + ".flv"
	return liveUrl
}

func VerifyCallBackSign(sign string, t string) bool {
	key := "00ebfd9b01e033cf1bab72ab110b7715"
	if sign == Md5(key+t) {
		return true
	}
	return false
}

func CloseStream(streamId string) bool {
	now := time.Now()
	endOfDay := GetEndOfThatDay(now).Unix()
	endOfDayStr := fmt.Sprintf("%d", endOfDay)
	key := "00ebfd9b01e033cf1bab72ab110b7715"
	sign := Md5(key + endOfDayStr)
	param := map[string]string{
		"cmd":                "1253302018",
		"interface":          "Live_Channel_SetStatus",
		"Param.s.channel_id": streamId,
		"Param.n.status":     "0",
		"t":                  endOfDayStr,
		"sign":               sign,
	}
	baseUrl := "http://fcgi.video.qcloud.com/common_access"
	HTTPGet(baseUrl, param)
	return true
}
