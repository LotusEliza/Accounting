package handlers

import (
	"accounting_software/server/models"
	"accounting_software/server/tcsctx"
	//"encoding/json"
	"github.com/finalist736/gokit/database"
	"github.com/finalist736/gokit/logger"
	"github.com/finalist736/gokit/response"
	"github.com/gocraft/web"
	//"io/ioutil"
)

type (
	UnitResponse struct {
		Items []*models.Unit `json:"Items"`
	}
)

func UnitsGet(ctx *tcsctx.Ctx, rw web.ResponseWriter, req *web.Request) {

	var (
		err error
	)

	session := database.GetDefaultSession()
	var resp UnitResponse
	_, err = session.
		Select("*").
		From(models.UnitsTableName).
		Load(&resp.Items)
	if err != nil {
		logger.StdOut().Warnf(
			"UnitsGet Select error: %s string: %s",
			err)
		response.Error(response.ERROR_DBERROR, rw)
		return
	}
	response.JsonIntent(resp, rw)
}
