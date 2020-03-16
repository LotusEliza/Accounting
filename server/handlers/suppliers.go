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
	SupplierResponse struct {
		Items []*models.Supplier `json:"Items"`
	}
)

func SuppliersGet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err error
	)

	session := database.GetDefaultSession()
	var resp SupplierResponse
	_, err = session.
		Select("*").
		From(models.SuppliersTableName).
		Load(&resp.Items)
	if err != nil {
		logger.StdOut().Warnf(
			"SuppliersGet Select error: %s string: %s",
			err)
		response.Error(response.ERROR_DBERROR, rw)
		return
	}
	response.JsonIntent(resp, rw)
}

func SupplierSet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err          error
		itemSupplier *models.Supplier
	)

	defer req.Body.Close()

	itemSupplier = new(models.Supplier)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("SupplierSet read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemSupplier)
	if err != nil {
		logger.StdErr().Errorf("SupplierSet json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		InsertInto(models.SuppliersTableName).
		Columns(models.SupplierColumns...).
		Record(itemSupplier).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("SupplierSet create error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}

func SupplierRemove(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err          error
		itemSupplier *models.Supplier
	)

	defer req.Body.Close()

	itemSupplier = new(models.Supplier)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("SupplierRemove read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemSupplier)
	if err != nil {
		logger.StdErr().Errorf("SupplierRemove json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		DeleteFrom(models.SuppliersTableName).
		Where("id=?",
			itemSupplier.ID).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("SupplierRemove remove error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}

func SupplierUpdate(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err          error
		itemSupplier *models.Supplier
	)

	defer req.Body.Close()

	itemSupplier = new(models.Supplier)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("SupplierUpdate read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemSupplier)
	if err != nil {
		logger.StdErr().Errorf("SupplierUpdate json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		Update(models.SuppliersTableName).
		Set("company_name", itemSupplier.CompanyName).
		Set("contact_name", itemSupplier.ContactName).
		Set("contact_title", itemSupplier.ContactTitle).
		Set("address", itemSupplier.Address).
		Set("city", itemSupplier.City).
		Set("phone", itemSupplier.Phone).
		Set("email", itemSupplier.Email).
		Where("id=?", itemSupplier.ID).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("SupplierUpdate update error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}
