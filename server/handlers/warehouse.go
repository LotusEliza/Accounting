package handlers

import (
	"accounting_software/server/models"
	"accounting_software/server/tcsctx"
	"encoding/json"
	"github.com/finalist736/gokit/database"
	"github.com/finalist736/gokit/logger"
	"github.com/finalist736/gokit/response"
	"github.com/gocraft/web"
	"io/ioutil"
)

type (
	WarehouseResponse struct {
		Items []*models.Warehouse `json:"Items"`
	}
)

type (
	WarehouseResponseSelect struct {
		Items []*models.SelectWarehouse `json:"Items"`
	}
)

func WarehouseGet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err error
	)

	session := database.GetDefaultSession()
	//SELECT a.id, a.name, a.num, b.date, b.roll FROM a INNER JOIN b ON a.id=b.id;
	var resp WarehouseResponseSelect
	_, err = session.
		Select("warehouse.id,"+
			"warehouse.amount, warehouse.comment,"+
			"supply.supplier_id, supply.date_create, supply.buy_price, supply.unit,"+
			"supply.surcharge, products.product_name, suppliers.company_name").
		From(models.WarehouseTableName).
		LeftJoin(models.SupplyTableName, "warehouse.supply_id=supply.id").
		LeftJoin(models.ProductsTableName, "warehouse.product_id=products.id").
		LeftJoin(models.SuppliersTableName, "supply.supplier_id=suppliers.id").
		Load(&resp.Items)
	if err != nil {
		logger.StdOut().Warnf(
			"ProductsGet Select error: %s string: %s",
			err)
		response.Error(response.ERROR_DBERROR, rw)
		return
	}
	response.JsonIntent(resp, rw)

	//session := database.GetDefaultSession()
	////SELECT a.id, a.name, a.num, b.date, b.roll FROM a INNER JOIN b ON a.id=b.id;
	//var resp WarehouseResponse
	//_, err = session.
	//	Select("*").
	//	From(models.WarehouseTableName).
	//	Load(&resp.Items)
	//if err != nil {
	//	logger.StdOut().Warnf(
	//		"WarehouseGet Select error: %s string: %s",
	//		err)
	//	response.Error(response.ERROR_DBERROR, rw)
	//	return
	//}
	//response.JsonIntent(resp, rw)
}

func WarehouseSet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err           error
		itemWarehouse *models.Warehouse
	)

	defer req.Body.Close()

	itemWarehouse = new(models.Warehouse)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("WarehouseSet read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemWarehouse)
	if err != nil {
		logger.StdErr().Errorf("WarehouseSet json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		InsertInto(models.WarehouseTableName).
		Columns(models.WarehouseColumns...).
		Record(itemWarehouse).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("WarehouseSet create error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}

func WarehouseRemove(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err           error
		itemWarehouse *models.Warehouse
	)

	defer req.Body.Close()

	itemWarehouse = new(models.Warehouse)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("WarehouseRemove read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemWarehouse)
	if err != nil {
		logger.StdErr().Errorf("WarehouseRemove json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		DeleteFrom(models.WarehouseTableName).
		Where("id=?",
			itemWarehouse.ID).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("WarehouseRemove remove error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}

//func ProductUpdate(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {
//
//	var (
//		err         error
//		itemProduct *models.Product
//	)
//
//	defer req.Body.Close()
//
//	itemProduct = new(models.Product)
//	ba, err := ioutil.ReadAll(req.Body)
//	if err != nil {
//		logger.StdErr().Errorf("ProductUpdate read body error: %s", err)
//		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
//		return
//	}
//
//	err = json.Unmarshal(ba, itemProduct)
//	if err != nil {
//		logger.StdErr().Errorf("ProductUpdate json error: %s", err)
//		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
//		return
//	}
//
//	session := database.GetDefaultSession()
//	_, err = session.
//		Update(models.WarehouseTableName).
//		Set("product_name", itemProduct.ProductName).
//		Set("category_id", itemProduct.CategoryID).
//		Set("sell_price", itemProduct.SellPrice).
//		Set("amount", itemProduct.Amount).
//		Set("vendor_code", itemProduct.VendorCode).
//		Where("id=?", itemProduct.ID).
//		Exec()
//	if err != nil {
//		logger.StdErr().Errorf("ProductUpdate update error: %s", err)
//		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
//		return
//	}
//	response.Error("", rw)
//}
//
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
//		Update(models.WarehouseTableName).
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
