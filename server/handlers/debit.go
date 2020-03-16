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
	DebitResponse struct {
		Items []*models.Debit `json:"Items"`
	}
)

type (
	DebitLastId struct {
		Item []*models.LastId `json:"Items"`
	}
)

type (
	ReportResponseDebit struct {
		Items []*models.ReportDebit `json:"Items"`
	}
)

func ReportDebitMonthGet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

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
		logger.StdErr().Errorf("ReportDebitMonthGet read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemReport)
	if err != nil {
		logger.StdErr().Errorf("ReportDebitMonthGet json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	var resp ReportResponseDebit
	_, err = session.
		Select("*").
		From(models.DebitTableName).
		Where("date_create BETWEEN ? and ?", itemReport.From, itemReport.To).
		Load(&resp.Items)
	if err != nil {
		logger.StdErr().Errorf("ReportDebitMonthGet create error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.JsonIntent(resp, rw)
}

func DebitGetLastId(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err error
	)

	session := database.GetDefaultSession()
	var resp DebitLastId
	_, err = session.
		Select("*").
		From(models.DebitTableName).
		OrderDesc("id").
		Limit(1).
		Load(&resp.Item)
	if err != nil {
		logger.StdOut().Warnf(
			"DebitGetLastId Select error: %s string: %s",
			err)
		response.Error(response.ERROR_DBERROR, rw)
		return
	}
	response.JsonIntent(resp, rw)
}

func DebitGet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err error
	)

	session := database.GetDefaultSession()
	var resp DebitResponse
	_, err = session.
		Select("*").
		From(models.DebitTableName).
		Load(&resp.Items)
	if err != nil {
		logger.StdOut().Warnf(
			"DebitGet Select error: %s string: %s",
			err)
		response.Error(response.ERROR_DBERROR, rw)
		return
	}
	response.JsonIntent(resp, rw)
}

func DebitSet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err       error
		itemDebit *models.Debit
	)

	defer req.Body.Close()

	itemDebit = new(models.Debit)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("DebitSet read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemDebit)
	if err != nil {
		logger.StdErr().Errorf("DebitSet json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		InsertInto(models.DebitTableName).
		Columns(models.DebitColumns...).
		Record(itemDebit).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("DebitSet create error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}

	response.Error("", rw)
}

//
//func ReportMonthGet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {
//
//	var (
//		err        error
//		itemReport *models.ReportDate
//		//itemWarehouse *models.Warehouse
//	)
//
//	defer req.Body.Close()
//
//	//itemWarehouse = new(models.Warehouse)
//
//	itemReport = new(models.ReportDate)
//	ba, err := ioutil.ReadAll(req.Body)
//	if err != nil {
//		logger.StdErr().Errorf("ReportMonthGet read body error: %s", err)
//		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
//		return
//	}
//
//	err = json.Unmarshal(ba, itemReport)
//	if err != nil {
//		logger.StdErr().Errorf("ReportMonthGet json error: %s", err)
//		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
//		return
//	}
//
//	session := database.GetDefaultSession()
//	var resp ReportResponse
//	_, err = session.
//		Select("credit.credit, supply.date_create, supply.comment").
//		From(models.DebitTableName).
//		LeftJoin(models.SupplyTableName, "credit.supply_id=supply.id").
//		Where("supply.date_create BETWEEN ? and ?", itemReport.From, itemReport.To).
//		Load(&resp.Items)
//	if err != nil {
//		logger.StdErr().Errorf("ReportMonthGet create error: %s", err)
//		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
//		return
//	}
//	response.JsonIntent(resp, rw)
//}
//
//func CreditRemove(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {
//
//	var (
//		err        error
//		itemDebit *models.Credit
//	)
//
//	defer req.Body.Close()
//
//	itemDebit = new(models.Credit)
//	ba, err := ioutil.ReadAll(req.Body)
//	if err != nil {
//		logger.StdErr().Errorf("CreditRemove read body error: %s", err)
//		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
//		return
//	}
//
//	err = json.Unmarshal(ba, itemDebit)
//	if err != nil {
//		logger.StdErr().Errorf("CreditRemove json error: %s", err)
//		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
//		return
//	}
//
//	session := database.GetDefaultSession()
//	_, err = session.
//		DeleteFrom(models.DebitTableName).
//		Where("supply_id=?",
//			itemDebit.SupplyID).
//		Exec()
//	if err != nil {
//		logger.StdErr().Errorf("CreditRemove remove error: %s", err)
//		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
//		return
//	}
//	response.Error("", rw)
//}
//
//func CreditUpdate(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {
//
//	var (
//		err        error
//		itemDebit *models.Credit
//	)
//
//	defer req.Body.Close()
//
//	itemDebit = new(models.Credit)
//	ba, err := ioutil.ReadAll(req.Body)
//	if err != nil {
//		logger.StdErr().Errorf("CreditUpdate read body error: %s", err)
//		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
//		return
//	}
//
//	err = json.Unmarshal(ba, itemDebit)
//	if err != nil {
//		logger.StdErr().Errorf("CreditUpdate json error: %s", err)
//		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
//		return
//	}
//
//	session := database.GetDefaultSession()
//	_, err = session.
//		Update(models.DebitTableName).
//		Set("credit", itemDebit.Credit).
//		Where("supply_id=?", itemDebit.SupplyID).
//		Exec()
//	if err != nil {
//		logger.StdErr().Errorf("CreditUpdate update error: %s", err)
//		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
//		return
//	}
//	response.Error("", rw)
//}
