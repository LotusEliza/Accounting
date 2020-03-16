package handlers

import (
	"encoding/json"
	//"encoding/json"
	"accounting_software/server/models"
	"accounting_software/server/tcsctx"
	"github.com/finalist736/gokit/database"
	"github.com/finalist736/gokit/logger"
	"github.com/finalist736/gokit/response"
	"github.com/gocraft/web"
	"io/ioutil"
	//"io/ioutil"
)

type (
	CashBoxResponse struct {
		Items []*models.CashBox `json:"Items"`
	}
)

func CashBoxGet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err error
	)

	session := database.GetDefaultSession()
	var resp CashBoxResponse
	_, err = session.
		Select("*").
		From(models.CashBoxTableName).
		Load(&resp.Items)
	if err != nil {
		logger.StdOut().Warnf(
			"CashBoxGet Select error: %s string: %s",
			err)
		response.Error(response.ERROR_DBERROR, rw)
		return
	}
	response.JsonIntent(resp, rw)
}

func CashBoxSet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err         error
		itemCashBox *models.CashBox
	)

	defer req.Body.Close()

	itemCashBox = new(models.CashBox)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("CashBoxSet read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemCashBox)
	if err != nil {
		logger.StdErr().Errorf("CashBoxSet json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		InsertInto(models.CashBoxTableName).
		Columns(models.CashBoxColumns...).
		Record(itemCashBox).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("CashBoxSet create error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}

////SELECT DISTINCT Country FROM Customers;
//
func CashBoxRemove(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err         error
		itemCashBox *models.Supplier
	)

	defer req.Body.Close()

	itemCashBox = new(models.Supplier)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("CashBoxRemove read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemCashBox)
	if err != nil {
		logger.StdErr().Errorf("CashBoxRemove json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		DeleteFrom(models.CashBoxTableName).
		Where("id=?",
			itemCashBox.ID).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("CashBoxRemove remove error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}

//func SupplierUpdate(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {
//
//	var (
//		err           error
//		itemCashBox *models.Supplier
//	)
//
//	defer req.Body.Close()
//
//	itemCashBox = new(models.Supplier)
//	ba, err := ioutil.ReadAll(req.Body)
//	if err != nil {
//		logger.StdErr().Errorf("SupplierUpdate read body error: %s", err)
//		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
//		return
//	}
//
//	err = json.Unmarshal(ba, itemCashBox)
//	if err != nil {
//		logger.StdErr().Errorf("SupplierUpdate json error: %s", err)
//		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
//		return
//	}
//
//	session := database.GetDefaultSession()
//	_, err = session.
//		Update(models.CashBoxTableName).
//		Set("company_name", itemCashBox.CompanyName).
//		Set("contact_name", itemCashBox.ContactName).
//		Set("contact_title", itemCashBox.ContactTitle).
//		Set("address", itemCashBox.Address).
//		Set("city", itemCashBox.City).
//		Set("phone", itemCashBox.Phone).
//		Set("email", itemCashBox.Email).
//		Where("id=?", itemCashBox.ID).
//		Exec()
//	if err != nil {
//		logger.StdErr().Errorf("SupplierUpdate update error: %s", err)
//		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
//		return
//	}
//	response.Error("", rw)
//}
