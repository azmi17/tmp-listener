package handler

import (
	"net/http"
	"tmp-report-transactions/delivery/handler/httpio"
	"tmp-report-transactions/helper"

	"github.com/gin-gonic/gin"
)

func ApplicationInformation(ctx *gin.Context) {

	httpio := httpio.NewRequestIO(ctx)

	httpio.Recv()

	appInfo := map[string]interface{}{
		"App Name":        helper.AppName,
		"App Description": helper.AppDescription,
		"App Version":     helper.AppVersion,
		"App Last Build:": helper.LastBuild,
	}

	httpio.Response(http.StatusOK, appInfo)
}
