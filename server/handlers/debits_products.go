package handlers

import (
	"accounting_software/server/models"
	"accounting_software/server/tcsctx"
	"encoding/json"
	"io/ioutil"

	//"encoding/json"
	//"io/ioutil"

	//"encoding/json"
	"github.com/finalist736/gokit/database"
	"github.com/finalist736/gokit/logger"
	"github.com/finalist736/gokit/response"
	"github.com/gocraft/web"
	//"io/ioutil"
)

type (
	DebitsProductsResponse struct {
		Items []*models.DebitsProducts `json:"Items"`
	}
)

func DebitsProductsGet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err error
	)

	session := database.GetDefaultSession()
	var resp DebitsProductsResponse
	_, err = session.
		Select("*").
		From(models.DebitsProductsTableName).
		Load(&resp.Items)
	if err != nil {
		logger.StdOut().Warnf(
			"DebitsProductsGet Select error: %s string: %s",
			err)
		response.Error(response.ERROR_DBERROR, rw)
		return
	}
	response.JsonIntent(resp, rw)
}

func DebitsProductsSet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err                error
		itemDebitsProducts *models.DebitsProducts
	)

	defer req.Body.Close()

	itemDebitsProducts = new(models.DebitsProducts)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("DebitsProductsSet read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemDebitsProducts)
	if err != nil {
		logger.StdErr().Errorf("DebitsProductsSet json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		InsertInto(models.DebitsProductsTableName).
		Columns(models.DebitsProductsColumns...).
		Record(itemDebitsProducts).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("DebitsProductsSet create error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}
