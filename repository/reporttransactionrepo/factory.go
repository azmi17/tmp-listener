package reporttransactionrepo

import (
	"database/sql"
	"errors"
	"tmp-report-transactions/repository/databasefactory"
	"tmp-report-transactions/repository/databasefactory/drivers"
)

func NewReportTransactionsRepo() (ReportTransactions, error) {

	conn := databasefactory.PayVoucher.GetConnection()

	currentDriver := databasefactory.PayVoucher.GetDriverName()
	if currentDriver == drivers.MYSQL {
		return newReportTransactionsMysqlImpl(conn.(*sql.DB)), nil
	} else {
		return nil, errors.New("unimplemented database driver")
	}

}
