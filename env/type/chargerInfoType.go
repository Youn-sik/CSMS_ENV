package envType

// 2. 충전소 관리
type ChargerInfoListReq struct {
	Bid   string `json:"bid" bson:"bid"`
	Bkey  string `json:"bkey" bson:"bkey"`
	Kind  string `json:"kind" bson:"kind"`
	Bdate string `json:"bdate,omitempty" bson:"bdate"`
	Cid   string `json:"cid,omitempty" bson:"cid"`
	Sid   string `json:"sid,omitempty" bson:"sid"`
	Rbid  string `json:"rbid,omitempty" bson:"rbid"`
}

type ChargerInfoListRes struct {
	Result string        `json:"result" bson:"result"`
	Rdate  string        `json:"rdate" bson:"rdate"`
	Rowcnt string        `json:"rowcnt" bson:"rowcnt"`
	Cinfo  []ChargerInfo `json:"cinfo" bson:"cinfo"`

	Errcode string `json:"errcode" bson:"errcode"`
	Errdtl  string `json:"errdtl" bson:"Errdtl"`
}

type ChargerInfo struct {
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
}

type ChargerInfoListAllReq struct {
	Bid    string `json:"bid" bson:"bid"`
	Bkey   string `json:"bkey" bson:"bkey"`
	Kind   string `json:"kind" bson:"kind"`
	Rbid   string `json:"rbid,omitempty" bson:"rbid"`
	Pageno string `json:"pageno,omitempty" bson:"pageno"`
	Rowcnt string `json:"rowcnt,omitempty" bson:"rowcnt"`
}

type ChargerInfoListAllRes struct {
	Result   string        `json:"result" bson:"result"`
	Rdate    string        `json:"rdate" bson:"rdate"`
	Totalcnt int           `json:"totalcnt" bson:"totalcnt"`
	Pageno   int           `json:"pageno" bson:"pageno"`
	Rowcnt   string        `json:"rowcnt" bson:"rowcnt"`
	Cinfo    []ChargerInfo `json:"cinfo" bson:"cinfo"`

	Errcode string `json:"errcode" bson:"errcode"`
	Errdtl  string `json:"errdtl" bson:"Errdtl"`
}

type ChargerStatusListReq struct {
	Bid  string `json:"bid" bson:"bid"`
	Bkey string `json:"bkey" bson:"bkey"`
	Kind string `json:"kind" bson:"kind"`
	Rbid string `json:"rbid,omitempty" bson:"rbid"`
	Sid  string `json:"sid,omitempty" bson:"sid"`
	Cid  string `json:"cid,omitempty" bson:"cid"`
}

type ChargerStatusListRes struct {
	Result string        `json:"result" bson:"result"`
	Rdate  string        `json:"rdate" bson:"rdate"`
	Rowcnt string        `json:"rowcnt" bson:"rowcnt"`
	Cstat  []ChargerInfo `json:"cstat" bson:"cstat"`

	Errcode string `json:"errcode" bson:"errcode"`
	Errdtl  string `json:"errdtl" bson:"Errdtl"`
}

type ChargerStatusListAllReq struct {
	Bid    string `json:"bid" bson:"bid"`
	Bkey   string `json:"bkey" bson:"bkey"`
	Kind   string `json:"kind" bson:"kind"`
	Rbid   string `json:"rbid,omitempty" bson:"rbid"`
	Pageno string `json:"pageno,omitempty" bson:"pageno"`
	Rowcnt string `json:"rowcnt,omitempty" bson:"rowcnt"`
}

type ChargerStatusListAllRes struct {
	Result   string        `json:"result" bson:"result"`
	Rdate    string        `json:"rdate" bson:"rdate"`
	Totalcnt int           `json:"totalcnt" bson:"totalcnt"`
	Pageno   int           `json:"pageno" bson:"pageno"`
	Rowcnt   string        `json:"rowcnt" bson:"rowcnt"`
	Cstat    []ChargerInfo `json:"cstat" bson:"cstat"`

	Errcode string `json:"errcode" bson:"errcode"`
	Errdtl  string `json:"errdtl" bson:"Errdtl"`
}
