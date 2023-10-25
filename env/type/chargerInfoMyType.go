package envType

// 3. 충전기 관리
type ChargerInfoMyListAllReq struct {
	Bid    string `json:"bid" bson:"bid"`
	Bkey   string `json:"bkey" bson:"bkey"`
	Kind   string `json:"kind" bson:"kind"`
	Pageno string `json:"pageno,omitempty" bson:"pageno"`
	Rowcnt string `json:"rowcnt,omitempty" bson:"rowcnt"`
	Sid    string `json:"sid,omitempty" bson:"sid"`
	Cid    string `json:"cid,omitempty" bson:"cid"`
}

type ChargerInfoMyListAllRes struct {
	Result   string          `json:"result" bson:"result"`
	Rdate    string          `json:"rdate" bson:"rdate"`
	Totalcnt int             `json:"totalcnt" bson:"totalcnt"`
	Pageno   int             `json:"pageno" bson:"pageno"`
	Rowcnt   string          `json:"rowcnt" bson:"rowcnt"`
	Cinfo    []ChargerInfoMy `json:"cinfo" bson:"cinfo"`

	Errcode string `json:"errcode" bson:"errcode"`
	Errdtl  string `json:"errdtl" bson:"Errdtl"`
}

type ChargerInfoMy struct {
	Bid         string `json:"bid" bson:"bid"`
	Sid         string `json:"sid" bson:"sid"`
	Cid         string `json:"cid" bson:"cid"`
	Zcode       string `json:"zcode" bson:"zcode"`
	Name        string `json:"name" bson:"name"`
	Addr        string `json:"addr" bson:"addr"`
	Addrdtl     string `json:"addrdtl" bson:"addrdtl"`
	Daddr       string `json:"daddr" bson:"daddr"`
	Daddrdtl    string `json:"daddrdtl" bson:"daddrdtl"`
	Kind        string `json:"kind" bson:"kind"`
	Kinddtl     string `json:"kinddtl" bson:"kinddtl"`
	Gps         string `json:"gps" bson:"gps"`
	Usetime     string `json:"usetime" bson:"usetime"`
	Free        string `json:"free" bson:"free"`
	Freedtl     string `json:"freedtl" bson:"freedtl"`
	Bname       string `json:"bname" bson:"bname"`
	Bcall       string `json:"bcall" bson:"bcall"`
	Type        string `json:"type" bson:"type"`
	Reserv      string `json:"reserv" bson:"reserv"`
	Member      string `json:"member" bson:"member"`
	Pay         string `json:"pay" bson:"pay"`
	Fee         string `json:"fee" bson:"fee"`
	Cable       string `json:"cable" bson:"cable"`
	Status      string `json:"status" bson:"status"`
	Statdt      string `json:"statdt" bson:"statdt"`
	Note        string `json:"note" bson:"note"`
	Bmngid      string `json:"bmngid" bson:"bmngid"`
	Limityn     string `json:"limityn" bson:"limityn"`
	Limitdetail string `json:"limitdetail" bson:"limitdetail"`
	Delyn       string `json:"delyn" bson:"delyn"`
	Deldetail   string `json:"deldetail" bson:"deldetail"`
	LastTsdt    string `json:"last_tsdt" bson:"last_tsdt"`
	LastTedt    string `json:"last_tedt" bson:"last_tedt"`
	NowTsdt     string `json:"now_tsdt" bson:"now_tsdt"`
	Method      string `json:"method" bson:"method"`
	Output      string `json:"output" bson:"output"`
	Expireddt   string `json:"expireddt" bson:"expireddt"`
	Year        string `json:"year" bson:"year"`
	Stype       string `json:"stype" bson:"stype"`
	Maker       string `json:"maker" bson:"maker"`
	Subsid      string `json:"subsid" bson:"subsid"`
	Owner       string `json:"owner" bson:"owner"`
	Sign        string `json:"sign" bson:"sign"`
	Setup       string `json:"setup" bson:"setup"`
	Month       string `json:"month" bson:"month"`
	Subsidy     string `json:"subsidy" bson:"subsidy"`
	Motie       string `json:"motie" bson:"motie"`
	Regdate     string `json:"regdate" bson:"regdate"`
	Upddate     string `json:"upddate" bson:"upddate"`
}

type ChargerStatusUpdateReq struct {
	Bid   string          `json:"bid" bson:"bid"`
	Bkey  string          `json:"bkey" bson:"bkey"`
	Cstat []ChargerStatus `json:"cstat,omitempty" bson:"cstat"`
}

type ChargerStatus struct {
	Sid      string `json:"sid" bson:"sid"`
	Cid      string `json:"cid" bson:"cid"`
	Status   string `json:"status" bson:"status"`
	Statdt   string `json:"statdt" bson:"statdt"`
	LastTsdt string `json:"last_tsdt" bson:"last_tsdt"`
	LastTedt string `json:"last_tedt" bson:"last_tedt"`
	NowTsdt  string `json:"now_tsdt" bson:"now_tsdt"`
}

type ChargerStatusUpdateRes struct {
	Result   string                   `json:"result" bson:"result"`
	Rdate    string                   `json:"rdate" bson:"rdate"`
	Reqcnt   string                   `json:"reqcnt" bson:"reqcnt"`
	Updcnt   string                   `json:"updcnt" bson:"updcnt"`
	Limitcnt string                   `json:"limitcnt" bson:"limitcnt"`
	Errcnt   string                   `json:"errcnt" bson:"errcnt"`
	Errcode  string                   `json:"errcode" bson:"errcode"`
	Errdtl   string                   `json:"errdtl" bson:"Errdtl"`
	Errlist  []ChargerStatusUpdateErr `json:"errlist" bson:"errlist"`
}

type ChargerStatusUpdateErr struct {
	Sid    string `json:"sid" bson:"sid"`
	Cid    string `json:"cid" bson:"cid"`
	Errmsg string `json:"errmsg" bson:"errmsg"`
}
