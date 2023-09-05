package reporttransactionrepo

import (
	"database/sql"
	"errors"
	"fmt"
	"tmp-report-transactions/entities"
	"tmp-report-transactions/entities/err"
)

func newReportTransactionsMysqlImpl(payVoucherDB *sql.DB) ReportTransactions {
	return &reportTransactionsMysqlImpl{payVoucherDB: payVoucherDB}
}

type reportTransactionsMysqlImpl struct {
	payVoucherDB *sql.DB
}

func (report *reportTransactionsMysqlImpl) Save(input entities.VoucherReversal) error {
	stmt, er := report.payVoucherDB.Prepare(`INSERT INTO voucher_reversal(
		cust_id,
		server_trxid,
		client_trxid,
		product,
		msisdn,
		status,
		price,
		balance,
		sn,
		msg,
		result,
		vn,
		tgl_trans
	) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)`)
	if er != nil {
		return errors.New(fmt.Sprint("error while prepare save tmp report transaction : ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	if _, er := stmt.Exec(
		input.CustID,
		input.ServerTRXID,
		input.ClientTRXID,
		input.Product,
		input.Msisdn,
		input.Status,
		input.Price,
		input.Balance,
		input.SN,
		input.Message,
		input.Result,
		input.VN,
		input.TglTrans,
	); er != nil {
		return errors.New(fmt.Sprint("error while create save tmp report transaction : ", er.Error()))
	} else {

		return er
	}
}

func (report *reportTransactionsMysqlImpl) Update(input entities.VoucherReversal) (er error) {
	stmt, er := report.payVoucherDB.Prepare(`UPDATE voucher_reversal SET 
		server_trxid = ?,
		client_trxid = ?,
		product = ?,
		msisdn = ?,
		status = ?,
		price = ?,
		balance = ?,
		sn = ?,
		msg = ?,
		result = ?,
		tgl_trans  = ?
		WHERE client_trxid = ? AND msisdn = ?`)
	if er != nil {
		return errors.New(fmt.Sprint("error while prepare update tmp report transaction : ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	if _, er := stmt.Exec(
		input.ServerTRXID,
		input.ClientTRXID,
		input.Product,
		input.Msisdn,
		input.Status,
		input.Price,
		input.Balance,
		input.SN,
		input.Message,
		input.Result,
		input.TglTrans,
		input.ClientTRXID,
		input.Msisdn); er != nil {
		return errors.New(fmt.Sprint("error while update tmp report transaction : ", er.Error()))
	}

	return nil
}

func (report *reportTransactionsMysqlImpl) GetExistingTransaction(stan, destNumber string) (data entities.VoucherReversal, er error) {

	row := report.payVoucherDB.QueryRow(`SELECT
	client_trxid,
	msisdn
	FROM voucher_reversal WHERE client_trxid = ? AND msisdn = ?`, stan, destNumber)
	er = row.Scan(
		&data.ClientTRXID,
		&data.Msisdn,
	)
	if er != nil {
		if er == sql.ErrNoRows {
			return data, err.NoRecord
		} else {
			return data, errors.New(fmt.Sprint("error while get existing transaction: ", er.Error()))
		}
	}
	return
}
