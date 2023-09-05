package reporttransactionrepo

import (
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"
	"tmp-report-transactions/entities"

	"github.com/vjeantet/jodaTime"
)

func GetConnection() *sql.DB {
	dataSource := "root:uSS10nl1n3@tcp(192.169.253.56:3306)/pay_voucher?parseTime=true"
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func TestSaveTransaction(t *testing.T) {
	db := GetConnection()
	repo := newReportTransactionsMysqlImpl(db)

	payVoucher := entities.VoucherReversal{}
	payVoucher.ServerTRXID = ""
	payVoucher.ClientTRXID = "770887091021"
	payVoucher.Product = ""
	payVoucher.Msisdn = "085249242792"
	payVoucher.Status = ""
	payVoucher.Price = ""
	payVoucher.Balance = ""
	payVoucher.SN = ""
	payVoucher.Message = ""
	payVoucher.Result = ""
	payVoucher.TglTrans = jodaTime.Format("YYMMdd", time.Now())

	err := repo.Update(payVoucher)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("re-callback success")
}

func TestGetExistingTransaction(t *testing.T) {
	db := GetConnection()
	repo := newReportTransactionsMysqlImpl(db)

	data, err := repo.GetExistingTransaction("770887091021", "085249242792")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("dest:", data.Msisdn, "stan:", data.ClientTRXID)
}
