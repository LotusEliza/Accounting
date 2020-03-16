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
	CategoryResponse struct {
		Items []*models.Category `json:"Items"`
	}
)

func CategoriesGet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err error
	)

	session := database.GetDefaultSession()
	var resp CategoryResponse
	_, err = session.
		Select("*").
		From(models.CategoriesTableName).
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

func CategorySet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err          error
		itemCategory *models.Category
	)

	defer req.Body.Close()

	itemCategory = new(models.Category)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("CategorySet read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemCategory)
	if err != nil {
		logger.StdErr().Errorf("CategorySet json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		InsertInto(models.CategoriesTableName).
		Columns(models.CategoryColumns...).
		Record(itemCategory).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("CategorySet create error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}

func CategoryRemove(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err          error
		itemCategory *models.Category
	)

	defer req.Body.Close()

	itemCategory = new(models.Category)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("CategoryRemove read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemCategory)
	if err != nil {
		logger.StdErr().Errorf("CategoryRemove json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		DeleteFrom(models.CategoriesTableName).
		Where("id=?",
			itemCategory.ID).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("CategoryRemove remove error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}

func CategoryUpdate(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err          error
		itemCategory *models.Category
	)

	defer req.Body.Close()

	itemCategory = new(models.Category)
	ba, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.StdErr().Errorf("CategoryUpdate read body error: %s", err)
		response.ErrorBadRequest(response.ERROR_NO_CONTENT, rw)
		return
	}

	err = json.Unmarshal(ba, itemCategory)
	if err != nil {
		logger.StdErr().Errorf("CategoryUpdate json error: %s", err)
		response.ErrorBadRequest(response.ERROR_REQUEST_DATA, rw)
		return
	}

	session := database.GetDefaultSession()
	_, err = session.
		Update(models.CategoriesTableName).
		Set("category_name", itemCategory.CategoryName).
		Set("category_description", itemCategory.CategoryDescription).
		Where("id=?", itemCategory.ID).
		Exec()
	if err != nil {
		logger.StdErr().Errorf("CategoryUpdate update error: %s", err)
		response.ErrorInternalServer(response.ERROR_DBERROR, rw)
		return
	}
	response.Error("", rw)
}
