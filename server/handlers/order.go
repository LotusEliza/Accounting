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
	//"github.com/gocraft/dbr"
)

type (
	OrderResponse struct {
		Items []*models.Order `json:"Items"`
	}
)

type (
	ResponseCustomer struct {
		Items []*models.Customer `json:"Items"`
	}
)

type (
	ResponseOrder struct {
		Items []*models.RespOrder `json:"Items"`
	}
)

//type (
//	RespOrder struct {
//		Items []*models.RespOrder `json:"Items"`
//	}
//)

func OrdersGet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err error
	)

	session := database.GetDefaultSession()
	var resp ResponseOrder
	_, err = session.
		Select("orders.amount, orders.date_create, orders.id, products.product_name,"+
			" customers.customer_name, customers.phone").
		From(models.OrdersTableName).
		LeftJoin(models.CustomersTableName, "customers.id=orders.customer_id").
		LeftJoin(models.ProductsTableName, "products.id=orders.product_id").
		Load(&resp.Items)
	if err != nil {
		logger.StdOut().Warnf(
			"OrdersGet Select error: %s string: %s",
			err)
		response.Error(response.ERROR_DBERROR, rw)
		return
	}
	response.JsonIntent(resp, rw)
}

func OrderSet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err       error
		itemOrder *models.RespOrder
	)

	defer req.Body.Close()

	itemOrder = new(models.RespOrder)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("OrderSet read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemOrder)
	if err != nil {
		logger.StdErr().Errorf("OrderSet json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		InsertInto(models.OrdersTableName).
		Columns(models.OrderColumns...).
		Record(itemOrder).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("OrderSet create error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)

}

func CustomerIDGet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err          error
		itemCustomer *models.Customer
		//itemWarehouse *models.Warehouse
	)

	defer req.Body.Close()

	//itemWarehouse = new(models.Warehouse)

	itemCustomer = new(models.Customer)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("CustomerIDGet read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemCustomer)
	if err != nil {
		logger.StdErr().Errorf("CustomerIDGet json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	var resp models.Customer
	_, err = session.
		Select("id").
		From(models.CustomersTableName).
		Where("phone=?", itemCustomer.Phone).
		Load(&resp)
	if err != nil {
		logger.StdErr().Errorf("CustomerIDGet create error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.JsonIntent(resp, rw)
}

func CustomerGetChatID(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err          error
		itemCustomer *models.Customer
		//itemWarehouse *models.Warehouse
	)

	defer req.Body.Close()

	//itemWarehouse = new(models.Warehouse)

	itemCustomer = new(models.Customer)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("CustomerGetChatID read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemCustomer)
	if err != nil {
		logger.StdErr().Errorf("CustomerGetChatID json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	var resp models.Customer
	_, err = session.
		Select("*").
		From(models.CustomersTableName).
		Where("chat_id=?", itemCustomer.ChatID).
		Load(&resp)
	if err != nil {
		logger.StdErr().Errorf("CustomerGetChatID create error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.JsonIntent(resp, rw)
}

func CustomersGetLastId(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err error
	)

	session := database.GetDefaultSession()
	var resp models.LastCustomerId
	_, err = session.
		Select("*").
		From(models.CustomersTableName).
		OrderDesc("id").
		Limit(1).
		Load(&resp.ID)
	if err != nil {
		logger.StdOut().Warnf(
			"CustomersGetLastId Select error: %s string: %s",
			err)
		response.Error(response.ERROR_DBERROR, rw)
		return
	}
	response.JsonIntent(resp, rw)
}

func CustomerSet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err          error
		itemCustomer *models.Customer
	)

	defer req.Body.Close()

	itemCustomer = new(models.Customer)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("OrderSet read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemCustomer)
	if err != nil {
		logger.StdErr().Errorf("OrderSet json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		InsertInto(models.CustomersTableName).
		Columns(models.CustomerColumns...).
		Record(itemCustomer).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("OrderSet create error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}

func OrderRemove(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err       error
		itemOrder *models.Order
	)

	defer req.Body.Close()

	itemOrder = new(models.Order)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("OrderRemove read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemOrder)
	if err != nil {
		logger.StdErr().Errorf("OrderRemove json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		DeleteFrom(models.OrdersTableName).
		Where("customer_id=?",
			itemOrder.ID).
		Limit(1).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("OrderRemove remove error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}

func OrderLastGet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err       error
		itemOrder *models.Order
	)

	defer req.Body.Close()

	itemOrder = new(models.Order)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("OrderLastGet read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemOrder)
	if err != nil {
		logger.StdErr().Errorf("OrderLastGet json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	var resp *models.RespOrder
	_, err = session.
		Select("orders.amount, products.product_name").
		From(models.OrdersTableName).
		LeftJoin(models.ProductsTableName, "products.id=orders.product_id").
		Where("customer_id=?",
			itemOrder.CustomerID).
		Limit(1).
		Load(&resp)
	if err != nil {
		logger.StdOut().Warnf(
			"OrderLastGet Select error: %s string: %s",
			err)
		response.Error(response.ERROR_DBERROR, rw)
		return
	}
	response.JsonIntent(resp, rw)
}
