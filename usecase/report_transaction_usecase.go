package usecase

import (
	"fmt"
	"time"
	"tmp-report-transactions/entities"
	"tmp-report-transactions/entities/err"
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
	reportTransactionRepo, _ := reporttransactionrepo.NewReportTransactionsRepo()

	// Search existing transaction..
	data, er := reportTransactionRepo.GetExistingTransaction(input.ReffID, input.Dest)
	if er != nil {
		if er == err.NoRecord {
			entities.PrintLog(fmt.Sprintf("Incoming new data: %s", input.ReffID))
		} else {
			return er
		}

	}

	// Compose payload
	payVoucher := entities.VoucherReversal{}
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

	if input.ReffID == data.ClientTRXID && input.Dest == data.Msisdn {
		// Existing transaction..
		entities.PrintLog(fmt.Sprintf("re-callback data: %s", data.ClientTRXID))
		er = reportTransactionRepo.Update(payVoucher)
		if er != nil {
			return er
		}
	} else {
		// New transaction..
		er = reportTransactionRepo.Save(payVoucher)
		if er != nil {
			return er
		}

	}

	return nil
}
