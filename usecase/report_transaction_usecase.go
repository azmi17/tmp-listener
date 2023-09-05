package usecase

import (
	"time"
	"tmp-report-transactions/entities"
	"tmp-report-transactions/repository/reporttransactionrepo"

	"github.com/vjeantet/jodaTime"
)

type ReportTransactionsUsecase interface {
	ReportTransactionListener(input entities.ReportTransactionRequest) error
}

type reportTransactionsUsecase struct{}

func NewReportTransactionUsecase() ReportTransactionsUsecase {
	return &reportTransactionsUsecase{}
}

func (report *reportTransactionsUsecase) ReportTransactionListener(input entities.ReportTransactionRequest) (er error) {

	payVoucher := entities.VoucherReversal{}

	// empty string values
	payVoucher.CustID = entities.EmptyStringVal
	payVoucher.VN = entities.EmptyStringVal

	// compose payload
	payVoucher.ServerTRXID = input.ServerID
	payVoucher.ClientTRXID = input.ReffID
	payVoucher.Product = input.Produk
	payVoucher.Msisdn = input.Dest
	payVoucher.Price = input.Price
	payVoucher.Balance = input.Balance
	payVoucher.SN = input.SN
	payVoucher.Message = input.Msgs
	payVoucher.TglTrans = jodaTime.Format("YYMMdd", time.Now())
	if input.Rescode == "4" {
		payVoucher.Status = entities.StrSuccesCodeVal
		payVoucher.Result = entities.StrSuccesCodeVal
	} else {
		payVoucher.Status = entities.StrFailedCodeVal
		payVoucher.Result = entities.StrFailedCodeVal
	}

	// save
	reportTransactionRepo, _ := reporttransactionrepo.NewReportTransactionsRepo()
	er = reportTransactionRepo.Save(payVoucher)
	if er != nil {
		return er
	}

	return nil
}
