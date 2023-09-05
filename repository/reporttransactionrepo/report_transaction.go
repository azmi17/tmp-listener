package reporttransactionrepo

import (
	"tmp-report-transactions/entities"
)

type ReportTransactions interface {
	Save(input entities.VoucherReversal) error
	Update(input entities.VoucherReversal) error
	GetExistingTransaction(stan, destNumber string) (entities.VoucherReversal, error)
}
