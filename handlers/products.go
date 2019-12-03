package handlers

import (
	"accounting_software/models"
	"accounting_software/tcsctx"
	"encoding/json"
	"github.com/finalist736/gokit/database"
	"github.com/finalist736/gokit/logger"
	"github.com/finalist736/gokit/response"
	"github.com/gocraft/web"
	"io/ioutil"
)

type (
	ProductResponse struct {
		Items []*models.Product `json:"Items"`
	}
)

type (
	ProductResponseSelect struct {
		Items []*models.SelectProduct `json:"Items"`
	}
)

func ProductsGet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err error
	)

	session := database.GetDefaultSession()
	//SELECT a.id, a.name, a.num, b.date, b.roll FROM a INNER JOIN b ON a.id=b.id;
	var resp ProductResponseSelect
	_, err = session.
		Select("products.id, products.vendor_code, products.product_name, products.supplier_id, suppliers.company_name,"+
			" products.category_id, categories.category_name, products.buy_price, products.sell_price,"+
			" products.amount, products.date_create").
		From(models.ProductsTableName).
		LeftJoin(models.CategoriesTableName, "products.category_id=categories.id").
		LeftJoin(models.SuppliersTableName, "products.supplier_id=suppliers.id").
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

func ProductSet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err         error
		itemProduct *models.Product
	)

	defer req.Body.Close()

	itemProduct = new(models.Product)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("ProductSet read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemProduct)
	if err != nil {
		logger.StdErr().Errorf("ProductSet json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		InsertInto(models.ProductsTableName).
		Columns(models.ProductColumns...).
		Record(itemProduct).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("ProductSet create error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}

func ProductRemove(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err         error
		itemProduct *models.Product
	)

	defer req.Body.Close()

	itemProduct = new(models.Product)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("ProductRemove read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemProduct)
	if err != nil {
		logger.StdErr().Errorf("ProductRemove json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		DeleteFrom(models.ProductsTableName).
		Where("id=?",
			itemProduct.ID).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("ProductRemove remove error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}

func ProductUpdate(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err         error
		itemProduct *models.Product
	)

	defer req.Body.Close()

	itemProduct = new(models.Product)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("ProductUpdate read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemProduct)
	if err != nil {
		logger.StdErr().Errorf("ProductUpdate json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		Update(models.ProductsTableName).
		Set("product_name", itemProduct.ProductName).
		Set("supplier_id", itemProduct.SupplierID).
		Set("category_id", itemProduct.CategoryID).
		Set("buy_price", itemProduct.BuyPrice).
		Set("sell_price", itemProduct.SellPrice).
		Set("amount", itemProduct.Amount).
		Set("vendor_code", itemProduct.VendorCode).
		Where("id=?", itemProduct.ID).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("ProductUpdate update error: %s", err)
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
		Update(models.ProductsTableName).
		Set("amount", itemDecommissioned.Amount).
		Where("id=?", itemDecommissioned.ProductID).
		Exec()
	tx.Commit()
	if err != nil {
		logger.StdErr().Errorf("ProductDecommissioned error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}
