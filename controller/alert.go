package gtools

import (
	"fmt"
	"github.com/uvite/gvmdesk/configs"
	"github.com/uvite/gvmdesk/internal"
	"github.com/uvite/gvmdesk/util"
)

// 创建警报
func (a *App) AddAlertItem(item internal.AlertItem) *util.Resp {
	_, err := a.Db.InsertOne(&item)
	if err != nil {
		a.Log.Error(fmt.Sprintf(configs.AddAlertItemErr, item.Title, err.Error()))
		return util.Error(err.Error())
	}
	return a.GetAlertList()
}

// 获取全部
func (a *App) GetAlertList() *util.Resp {
	itemList := make([]internal.AlertItem, 0)

	err := a.Db.Desc("date").Find(&itemList)

	if err != nil {
		a.Log.Error(fmt.Sprintf(configs.GetAlertListErr, err.Error()))
		return util.Error(err.Error())
	}

	resultMap := make(map[string]interface{}, 0)
	resultMap["list"] = itemList

	return util.Success(resultMap)
}

func (a *App) DelAlertItem(item internal.AlertItem) *util.Resp {
	_, err := a.Db.Delete(&item)
	if err != nil {
		a.Log.Error(fmt.Sprintf(configs.DelAlertItemErr, item.Title, err.Error()))
		return util.Error(err.Error())
	}
	return a.GetAlertList()
}

func (a *App) GetAlertItem(id int) *util.Resp {
	item := internal.AlertItem{}
	has, err := a.Db.ID(id).Get(&item)
	if err != nil {
		a.Log.Error(fmt.Sprintf(configs.DelAlertItemErr, item.Title, err.Error()))
		return util.Error(err.Error())
	}
	if has {
		return util.Success(item)
	}
	return nil

}

func (a *App) UpdateAlertItem(item internal.AlertItem) *util.Resp {
	_, err := a.Db.ID(item.Id).Update(&item)
	if err != nil {
		return util.Error(err.Error())
	}
	return a.GetAlertList()
}

func (a *App) DelAlertItemById(item internal.AlertItem) *util.Resp {
	_, err := a.Db.ID(item.Id).Delete(&item)
	if err != nil {
		return util.Error(err.Error())
	}
	return a.GetAlertList()
}

func (a *App) SetAlertStatus(id int, status bool) *util.Resp {
	item := internal.AlertItem{}
	has, err := a.Db.ID(id).Get(&item)
	if err != nil {
		a.Log.Error(fmt.Sprintf(configs.DelAlertItemErr, item.Title, err.Error()))
		return util.Error(err.Error())
	}
	if has {
		a.AddSymbolInterval(item.Symbol, item.Interval)
		go a.RunTestFile(item)
		return util.Success(item)
	}
	return nil
}
