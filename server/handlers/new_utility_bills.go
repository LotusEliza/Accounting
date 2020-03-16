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
	NewUtilResponse struct {
		Items []*models.NewUtil `json:"Items"`
	}
)

type (
	UtilNamesResponse struct {
		Items []*models.UtilName `json:"Items"`
	}
)

type (
	ReportUtilBillResponse struct {
		Items []*models.ReportUtilBill `json:"Items"`
	}
)

//type (
//	ReportUtilResponse struct {
//		Items []*models.ReportUtil `json:"Items"`
//	}
//)

func NewUtilityGet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err error
	)

	session := database.GetDefaultSession()
	var resp NewUtilResponse
	_, err = session.
		Select("*").
		From(models.NewUtilTableName).
		Load(&resp.Items)
	if err != nil {
		logger.StdOut().Warnf(
			"UtilityGet Select error: %s string: %s",
			err)
		response.Error(response.ERROR_DBERROR, rw)
		return
	}
	response.JsonIntent(resp, rw)
}

func NewBillNamesGet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err error
	)

	session := database.GetDefaultSession()
	var resp UtilNamesResponse
	_, err = session.
		Select("bill_name").Distinct().
		From(models.NewUtilTableName).
		Load(&resp.Items)
	if err != nil {
		logger.StdOut().Warnf(
			"UtilityGet Select error: %s string: %s",
			err)
		response.Error(response.ERROR_DBERROR, rw)
		return
	}
	response.JsonIntent(resp, rw)
}

func NewUtilitySet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err      error
		itemUtil *models.NewUtil
	)

	defer req.Body.Close()

	itemUtil = new(models.NewUtil)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("NewUtilitySet read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemUtil)
	if err != nil {
		logger.StdErr().Errorf("NewUtilitySet json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		InsertInto(models.NewUtilTableName).
		Columns(models.NewUtilityColumns...).
		Record(itemUtil).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("NewUtilitySet create error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)

	//session := database.GetDefaultSession()
	//tx, err := session.Begin()
	//defer tx.RollbackUnlessCommitted()
	//_, err = session.
	//	InsertInto(models.UtilTableName).
	//	Columns(models.UtilityColumns...).
	//	Record(itemUtil).
	//	Exec()
	//_, err = session.
	//	InsertInto(models.CreditTableName).
	//	Columns(models.CreditColumns...).
	//	Record(itemUtil).
	//	Exec()
	//tx.Commit()
	//if err != nil {
	//	logger.StdErr().Errorf("UtilitySet error: %s", err)
	//	response.ErrorInternalServer(response.ERROR_DBERROR, rw)
	//	return
	//}
	//response.Error("", rw)
}

//
func NewBillRemove(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err      error
		itemBill *models.NewUtil
	)

	defer req.Body.Close()

	itemBill = new(models.NewUtil)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("NewBillRemove read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemBill)
	if err != nil {
		logger.StdErr().Errorf("NewBillRemove json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		DeleteFrom(models.NewUtilTableName).
		Where("id=?",
			itemBill.ID).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("NewBillRemove remove error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}

func UtilBillsMonthGet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

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
		logger.StdErr().Errorf("UtilBillsMonthGet read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemReport)
	if err != nil {
		logger.StdErr().Errorf("UtilBillsMonthGet json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	var resp ReportUtilBillResponse
	_, err = session.
		Select("*").
		From(models.NewUtilTableName).
		Where("new_utility_bills.date_create BETWEEN ? and ?", itemReport.From, itemReport.To).
		Load(&resp.Items)
	if err != nil {
		logger.StdErr().Errorf("UtilBillsMonthGet create error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.JsonIntent(resp, rw)
}

//
//func ReportUtilMonthGet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {
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
//		logger.StdErr().Errorf("ReportUtilMonthGet read body error: %s", err)
//		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
//		return
//	}
//
//	err = json.Unmarshal(ba, itemReport)
//	if err != nil {
//		logger.StdErr().Errorf("ReportUtilMonthGet json error: %s", err)
//		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
//		return
//	}
//
//	session := database.GetDefaultSession()
//	var resp ReportUtilResponse
//	_, err = session.
//		Select("credit.credit, utility_bills.date_create").
//		From(models.CreditTableName).
//		LeftJoin(models.UtilTableName, "credit.bill_id=utility_bills.id").
//		Where("utility_bills.date_create BETWEEN ? and ?", itemReport.From, itemReport.To).
//		Load(&resp.Items)
//	if err != nil {
//		logger.StdErr().Errorf("ReportUtilMonthGet create error: %s", err)
//		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
//		return
//	}
//	response.JsonIntent(resp, rw)
//}
