package handlers

import (
	"accounting_software/server/models"
	"accounting_software/server/tcsctx"
	"encoding/json"
	"io/ioutil"

	//"encoding/json"
	"github.com/finalist736/gokit/database"
	"github.com/finalist736/gokit/logger"
	"github.com/finalist736/gokit/response"
	"github.com/gocraft/web"
	//"io/ioutil"
)

type (
	SupplyResponse struct {
		Items []*models.Supply `json:"Items"`
	}
)

type (
	SupplyResponseSelect struct {
		Items []*models.SelectSupply `json:"Items"`
	}
)

type (
	SelectDecommissioned struct {
		Items []*models.SelectDecommissioned `json:"Items"`
	}
)

type (
	IdSelect struct {
		Items []*models.IdSelect `json:"Items"`
	}
)

func SupplyGet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err error
	)

	session := database.GetDefaultSession()
	var resp SupplyResponseSelect
	_, err = session.
		Select("supply.id, supply.date_create, supply.buy_price, supply.sell_price,"+
			" supply.product_id, supply.supplier_id,"+
			" supply.surcharge, supply.sup_amount, supply.amount, supply.unit_id, supply.comment,"+
			" products.product_name, suppliers.company_name, units.unit_name").
		From(models.SupplyTableName).
		LeftJoin(models.SuppliersTableName, "supply.supplier_id=suppliers.id").
		LeftJoin(models.ProductsTableName, "supply.product_id=products.id").
		LeftJoin(models.UnitsTableName, "supply.unit_id=units.id").
		Load(&resp.Items)
	if err != nil {
		logger.StdOut().Warnf(
			"ProductsGet Select error: %s string: %s",
			err)
		response.Error(response.ERROR_DBERROR, rw)
		return
	}
	response.JsonIntent(resp, rw)
}

func SupplySet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err        error
		itemSupply *models.Supply
		//itemWarehouse *models.Warehouse
	)

	defer req.Body.Close()

	//itemWarehouse = new(models.Warehouse)

	itemSupply = new(models.Supply)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("SupplySet read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemSupply)
	if err != nil {
		logger.StdErr().Errorf("SupplySet json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		InsertInto(models.SupplyTableName).
		Columns(models.SupplyColumns...).
		Record(itemSupply).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("SupplySet create error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}

func SupplyRemove(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err        error
		itemSupply *models.Supply
	)

	defer req.Body.Close()

	itemSupply = new(models.Supply)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("SupplyRemove read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemSupply)
	if err != nil {
		logger.StdErr().Errorf("SupplyRemove json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		DeleteFrom(models.SupplyTableName).
		Where("id=?",
			itemSupply.ID).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("SupplyRemove remove error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}

func SupplyUpdate(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err       error
		itemSuply *models.Supply
	)

	defer req.Body.Close()

	itemSuply = new(models.Supply)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("SupplyUpdate read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemSuply)
	if err != nil {
		logger.StdErr().Errorf("SupplyUpdate json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		Update(models.SupplyTableName).
		Set("product_id", itemSuply.ProductID).
		Set("supplier_id", itemSuply.SupplierID).
		Set("buy_price", itemSuply.BuyPrice).
		Set("sell_price", itemSuply.SellPrice).
		Set("surcharge", itemSuply.Surcharge).
		Set("amount", itemSuply.Amount).
		Set("sup_amount", itemSuply.SupAmount).
		Set("unit_id", itemSuply.UnitID).
		Set("comment", itemSuply.Comment).
		Where("id=?", itemSuply.ID).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("SupplyUpdate update error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}

func SupplyUpdateAmount(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err       error
		itemSuply *models.Supply
	)

	defer req.Body.Close()

	itemSuply = new(models.Supply)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("SupplyUpdateAmount read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemSuply)
	if err != nil {
		logger.StdErr().Errorf("SupplyUpdateAmount json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		Update(models.SupplyTableName).
		Set("amount", itemSuply.Amount).
		Where("id=?", itemSuply.ID).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("SupplyUpdateAmount update error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}

func Decommissioned(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err                error
		itemDecommissioned *models.SelectDecommissioned
	)

	defer req.Body.Close()

	itemDecommissioned = new(models.SelectDecommissioned)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("ProductDecommissioned read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemDecommissioned)
	if err != nil {
		logger.StdErr().Errorf("ProductDecommissioned json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	tx, err := session.Begin()
	defer tx.RollbackUnlessCommitted()
	_, err = session.
		InsertInto(models.DecommissionedTableName).
		Columns(models.DecommissionedColumns...).
		Record(itemDecommissioned).
		Exec()
	_, err = session.
		Update(models.SupplyTableName).
		Set("amount", itemDecommissioned.Amount).
		Where("id=?", itemDecommissioned.ID).
		Exec()
	tx.Commit()
	if err != nil {
		logger.StdErr().Errorf("ProductDecommissioned error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}

func DecommissionedMonthGet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err        error
		itemReport *models.ReportDate
		//itemWarehouse *models.Warehouse
	)

	defer req.Body.Close()

	//itemWarehouse = new(models.Warehouse)

	itemReport = new(models.ReportDate)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("DecommissionedMonthGet read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemReport)
	if err != nil {
		logger.StdErr().Errorf("DecommissionedMonthGet json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	var resp SelectDecommissioned
	_, err = session.
		Select("decommissioned.*, products.product_name").
		From(models.DecommissionedTableName).
		LeftJoin(models.ProductsTableName, "decommissioned.product_id=products.id").
		Where("date_create BETWEEN ? and ?", itemReport.From, itemReport.To).
		Load(&resp.Items)
	if err != nil {
		logger.StdErr().Errorf("DecommissionedMonthGet create error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.JsonIntent(resp, rw)
}

func SupplyGetId(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err error
	)

	session := database.GetDefaultSession()
	var resp IdSelect
	_, err = session.
		Select("id").
		From(models.SupplyTableName).
		OrderDesc("id").
		Limit(1).
		Load(&resp.Items)
	if err != nil {
		logger.StdOut().Warnf(
			"CategoriesGet Select error: %s string: %s",
			err)
		response.Error(response.ERROR_DBERROR, rw)
		return
	}
	response.JsonIntent(resp, rw)
}

//func Decommissioned(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {
//
//	var (
//		err                error
//		itemDecommissioned *models.SelectDecommissioned
//	)
//
//	defer req.Body.Close()
//
//	itemDecommissioned = new(models.SelectDecommissioned)
//	ba, err := ioutil.ReadAll(req.Body)
//	if err != nil {
//		logger.StdErr().Errorf("ProductDecommissioned read body error: %s", err)
//		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
//		return
//	}
//
//	err = json.Unmarshal(ba, itemDecommissioned)
//	if err != nil {
//		logger.StdErr().Errorf("ProductDecommissioned json error: %s", err)
//		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
//		return
//	}
//
//	session := database.GetDefaultSession()
//	tx, err := session.Begin()
//	defer tx.RollbackUnlessCommitted()
//	_, err = session.
//		InsertInto(models.DecommissionedTableName).
//		Columns(models.DecommissionedColumns...).
//		Record(itemDecommissioned).
//		Exec()
//	_, err = session.
//		Update(models.SupplyTableName).
//		Set("amount", itemDecommissioned.Amount).
//		Where("id=?", itemDecommissioned.ProductID).
//		Exec()
//	tx.Commit()
//	if err != nil {
//		logger.StdErr().Errorf("ProductDecommissioned error: %s", err)
//		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
//		return
//	}
//	response.Error("", rw)
//}
