package common

import (
	"github.com/NJU-Echome/dragon-back-end/controllers/base"
	"github.com/NJU-Echome/dragon-back-end/err"
	"github.com/NJU-Echome/dragon-back-end/models"
	"github.com/NJU-Echome/dragon-back-end/util"
	"github.com/NJU-Echome/dragon-back-end/vo"
	"strconv"
	"strings"
)

type StoreController struct {
	base.BaseController
}

// @router / [get]
func (this *StoreController) Get() {
	city := this.GetString("city")
	if city == "" {
		this.RetError(err.ErrInputData, "city")
	}
	district := this.GetString("district")
	tagIdsStr := this.GetString("tagIds")
	tagIdStrs := strings.Split(tagIdsStr, ",")
	var tagIds []int
	for _, value := range tagIdStrs {
		id, err := strconv.Atoi(value)
		if err != nil {
			continue
		}
		tagIds = append(tagIds, id)
	}
	data := make(map[string]interface{})
	data["city"] = city
	data["district"] = district
	data["tags"] = tagIds
	var storeList []vo.Simple_store
	stores := models.GetStoresByCondition(4, city, district)
	util.PoListToVoList(stores, &storeList)
	data["stores"] = storeList
	this.RetData(data)
}
