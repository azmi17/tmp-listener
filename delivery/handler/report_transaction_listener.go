package handler

import (
	"net/http"
	"tmp-report-transactions/delivery/handler/httpio"
	"tmp-report-transactions/entities"
	"tmp-report-transactions/usecase"

	"github.com/gin-gonic/gin"
)

func ReportTransactionsListener(ctx *gin.Context) {
	httpio := httpio.NewRequestIO(ctx)
	httpio.Recv()

	payload := entities.ReportTransactionRequest{}
	httpio.BindUri(&payload)

	payload.Msgs = ctx.Query("msgs")
	payload.Produk = ctx.Query("produk")
	payload.Balance = ctx.Query("balance")
	payload.Price = ctx.Query("price")
	payload.ReffID = ctx.Query("reffid")
	payload.SN = ctx.Query("sn")
	payload.Dest = ctx.Query("dest")
	payload.ServerID = ctx.Query("serverid")
	payload.Rescode = ctx.Query("rescode")

	usecase := usecase.NewReportTransactionUsecase()
	er := usecase.ReportTransactionListener(payload)
	if er != nil {
		entities.PrintLog(er.Error())
		entities.PrintError(er.Error())
		httpio.ResponseString(http.StatusInternalServerError, "internal service error")
		return

	} else {
		httpio.Response(http.StatusOK, "OK")
	}
}
