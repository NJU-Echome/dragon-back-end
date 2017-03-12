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
	var storesResult []models.Store
	var storeList []vo.Simple_store
	if len(tagIds) == 0 {
		storesResult = models.GetStoresWithNoTag(city, district)
	} else {
		storesResult = models.GetStoresByCondition(tagIds, city, district)
		// for i, value := range tagIds {
		// 	if i == 0 {
		// 		storesResult = models.GetStoresByCondition(value, city, district)
		// 	} else {
		// 		stores := models.GetStoresByCondition(value, city, district)
		// 		for j, store := range storesResult {
		// 			ok := 0
		// 			for _, storeTmp := range stores {
		// 				if store.Id == storeTmp.Id {
		// 					ok = 1
		// 					break
		// 				}
		// 			}
		// 			if ok == 0 {
		// 				// 删除该元素
		// 				storesResult = append(storesResult[:j], storesResult[j+1:]...)
		// 			}
		// 		}
		// 	}

		// }
	}
	util.PoListToVoList(storesResult, &storeList)
	this.RetData(storeList)
}
