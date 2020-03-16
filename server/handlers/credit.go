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
	CreditResponse struct {
		Items []*models.Credit `json:"Items"`
	}
)

type (
	ReportResponse struct {
		Items []*models.Report `json:"Items"`
	}
)

func CreditGet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err error
	)

	session := database.GetDefaultSession()
	var resp CreditResponse
	_, err = session.
		Select("*").
		From(models.CreditTableName).
		Load(&resp.Items)
	if err != nil {
		logger.StdOut().Warnf(
			"CreditGet Select error: %s string: %s",
			err)
		response.Error(response.ERROR_DBERROR, rw)
		return
	}
	response.JsonIntent(resp, rw)
}

func CreditSet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err        error
		itemCredit *models.Credit
	)

	defer req.Body.Close()

	itemCredit = new(models.Credit)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("CreditSet read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemCredit)
	if err != nil {
		logger.StdErr().Errorf("CreditSet json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		InsertInto(models.CreditTableName).
		Columns(models.CreditColumns...).
		Record(itemCredit).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("CreditSet create error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}

func ReportMonthGet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

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
		logger.StdErr().Errorf("ReportMonthGet read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemReport)
	if err != nil {
		logger.StdErr().Errorf("ReportMonthGet json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	//TODO CHANgE
	session := database.GetDefaultSession()
	var resp ReportResponse
	_, err = session.
		Select("credit.credit, supply.date_create, supply.comment").
		From(models.CreditTableName).
		LeftJoin(models.SupplyTableName, "credit.supply_id=supply.id").
		//LeftJoin(models.UtilTableName, "credit.bill_id=utility_bills.id").
		Where("supply.date_create BETWEEN ? and ?",
			itemReport.From, itemReport.To).
		Load(&resp.Items)
	if err != nil {
		logger.StdErr().Errorf("ReportMonthGet create error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.JsonIntent(resp, rw)
}

//session := database.GetDefaultSession()
//var resp ReportResponse
//_, err = session.
//Select("credit.credit, supply.date_create, supply.comment, utility_bills.date_create").
//From(models.CreditTableName).
//LeftJoin(models.SupplyTableName, "credit.supply_id=supply.id").
//LeftJoin(models.UtilTableName, "credit.bill_id=utility_bills.id").
//Where("supply.date_create BETWEEN ? and ? OR utility_bills.date_create BETWEEN ? and ?", itemReport.From, itemReport.To, itemReport.From, itemReport.To).
//Load(&resp.Items)
//if err != nil {
//logger.StdErr().Errorf("ReportMonthGet create error: %s", err)
//response.ErrorInternalServer(response.ERROR_DBERROR, rw)
//return
//}
//response.JsonIntent(resp, rw)

func CreditRemove(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err        error
		itemCredit *models.Credit
	)

	defer req.Body.Close()

	itemCredit = new(models.Credit)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("CreditRemove read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemCredit)
	if err != nil {
		logger.StdErr().Errorf("CreditRemove json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		DeleteFrom(models.CreditTableName).
		Where("supply_id=?",
			itemCredit.SupplyID).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("CreditRemove remove error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}

func CreditUpdate(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err        error
		itemCredit *models.Credit
	)

	defer req.Body.Close()

	itemCredit = new(models.Credit)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("CreditUpdate read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemCredit)
	if err != nil {
		logger.StdErr().Errorf("CreditUpdate json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		Update(models.CreditTableName).
		Set("credit", itemCredit.Credit).
		Where("supply_id=?", itemCredit.SupplyID).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("CreditUpdate update error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}
