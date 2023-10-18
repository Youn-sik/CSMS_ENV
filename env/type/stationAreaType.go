package envType

type ChargerInfoListReq struct {
	Bdate string `json:"bdate" bson:"bdate"`
	Bkey  string `json:"bkey" bson:"bkey"`
	Bid   string `json:"bid" bson:"bid"`
	Cid   string `json:"cid" bson:"cid"`
	Sid   string `json:"sid" bson:"sid"`
	Rbid  string `json:"rbid" bson:"rbid"`
	Kind  string `json:"kind" bson:"kind"`
}

type ChargerInfoListRes struct {
	Result string        `json:"result" bson:"result"`
	Rdate  string        `json:"rdate" bson:"rdate"`
	Rowcnt string        `json:"rowcnt" bson:"rowcnt"`
	Cinfo  []ChargerInfo `json:"cinfo" bson:"cinfo"`

	Errcode string `json:"errcode" bson:"errcode"`
	Errdtl  string `json:"errdtl" bson:"errdtl"`
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
	Last_tsdt   string `json:"last_tsdt" bson:"last_tsdt"`
	Last_tedt   string `json:"last_tedt" bson:"last_tedt"`
	Now_tsdt    string `json:"now_tsdt" bson:"now_tsdt"`
	Method      string `json:"method" bson:"method"`
	Output      string `json:"output" bson:"output"`
}

type ChargerInfoListAllReq struct {
	Bid    string `json:"bid" bson:"bid"`
	Bkey   string `json:"bkey" bson:"bkey"`
	Kind   string `json:"kind" bson:"kind"`
	Rbid   string `json:"rbid" bson:"rbid"`
	Pageno string `json:"pageno" bson:"pageno"`
	Rowcnt string `json:"rowcnt" bson:"rowcnt"`
}

type ChargerInfoListAllRes struct {
	Result   string        `json:"result" bson:"result"`
	Rdate    string        `json:"rdate" bson:"rdate"`
	Totalcnt string        `json:"totalcnt" bson:"totalcnt"`
	Pageno   string        `json:"pageno" bson:"pageno"`
	Rowcnt   string        `json:"rowcnt" bson:"rowcnt"`
	Cinfo    []ChargerInfo `json:"cinfo" bson:"cinfo"`

	Errcode string `json:"errcode" bson:"errcode"`
	Errdtl  string `json:"errdtl" bson:"errdtl"`
}

type ChargerStatusListReq struct {
	Bid  string `json:"bid" bson:"bid"`
	Bkey string `json:"bkey" bson:"bkey"`
	Kind string `json:"kind" bson:"kind"`
	Rbid string `json:"rbid" bson:"rbid"`
	Sid  string `json:"sid" bson:"sid"`
	Cid  string `json:"cid" bson:"cid"`
}

type ChargerStatusListRes struct {
	Result string        `json:"result" bson:"result"`
	Rdate  string        `json:"rdate" bson:"rdate"`
	Rowcnt string        `json:"rowcnt" bson:"rowcnt"`
	Cstat  []ChargerInfo `json:"cstat" bson:"cstat"`

	Errcode string `json:"errcode" bson:"errcode"`
	Errdtl  string `json:"errdtl" bson:"errdtl"`
}

type ChargerStatusListAllReq struct {
	Bid    string `json:"bid" bson:"bid"`
	Bkey   string `json:"bkey" bson:"bkey"`
	Kind   string `json:"kind" bson:"kind"`
	Rbid   string `json:"rbid" bson:"rbid"`
	Pageno string `json:"pageno" bson:"pageno"`
	Rowcnt string `json:"rowcnt" bson:"rowcnt"`
}

type ChargerStatusListAllRes struct {
	Result   string        `json:"result" bson:"result"`
	Rdate    string        `json:"rdate" bson:"rdate"`
	Totalcnt string        `json:"totalcnt" bson:"totalcnt"`
	Pageno   string        `json:"pageno" bson:"pageno"`
	Rowcnt   string        `json:"rowcnt" bson:"rowcnt"`
	Cstat    []ChargerInfo `json:"cstat" bson:"cstat"`

	Errcode string `json:"errcode" bson:"errcode"`
	Errdtl  string `json:"errdtl" bson:"errdtl"`
}
