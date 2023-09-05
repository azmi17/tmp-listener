package entities

type VoucherReversal struct {
	CustID      string
	ServerTRXID string
	ClientTRXID string
	Product     string
	Msisdn      string
	Status      string
	Price       string
	Balance     string
	SN          string
	Message     string
	TransID     int
	Result      string
	VN          string
	TglTrans    string
	// TimeStamp    time.Time
}

type ReportTransactionRequest struct {
	Msgs     string `uri:"msgs"`
	Produk   string `uri:"produk"`
	Balance  string `uri:"balance"`
	Price    string `uri:"price"`
	ReffID   string `uri:"reffid"`
	SN       string `uri:"sn"`
	Dest     string `uri:"dest"`
	ServerID string `uri:"serverid"`
	Rescode  string `uri:"rescode"`
}

const (
	EmptyStringVal = ""
	IntZeroVal     = 0

	StrSuccesCodeVal = "0"
	StrFailedCodeVal = "11"
)
