package envType

type UseListReq struct {
	Bid    string `json:"bid" bson:"bid"`
	Bkey   string `json:"bkey" bson:"bkey"`
	Kind   string `json:"kind" bson:"kind"`
	Tbase  string `json:"tbase" bson:"tbase"`
	Start  string `json:"start" bson:"start"`
	End    string `json:"end" bson:"end"`
	Bdate  string `json:"bdate" bson:"bdate"`
	Pageno string `json:"pageno" bson:"pageno"`
	Rowcnt string `json:"rowcnt" bson:"rowcnt"`
}

type UseListRes struct {
	Result   string `json:"result" bson:"result"`
	Rcnt     string `json:"rcnt" bson:"rcnt"`
	Rdate    string `json:"rdate" bson:"rdate"`
	Totalcnt string `json:"totalcnt" bson:"totalcnt"`
	Pageno   string `json:"pageno" bson:"pageno"`
	Rowcnt   string `json:"rowcnt" bson:"rowcnt"`
	Use      []Use  `json:"use" bson:"use"`

	Errcode string `json:"errcode" bson:"errcode"`
	Errdtl  string `json:"errdtl" bson:"errdtl"`
}

type Use struct {
	Sid     string `json:"sid" bson:"sid"`
	Cid     string `json:"cid" bson:"cid"`
	Tsdt    string `json:"tsdt" bson:"tsdt"`
	Tedt    string `json:"tedt" bson:"tedt"`
	Pow     string `json:"pow" bson:"pow"`
	Mon     string `json:"mon" bson:"mon"`
	Regdate string `json:"regdate" bson:"regdate"`
}

type UseRegisterReq struct {
	Bid  string        `json:"bid" bson:"bid"`
	Bkey string        `json:"bkey" bson:"bkey"`
	Use  []UseRegister `json:"use" bson:"use"`
}

type UseRegister struct {
	Sid     string `json:"sid" bson:"sid"`
	Cid     string `json:"cid" bson:"cid"`
	Tbid    string `json:"tbid" bson:"tbid"`
	Tsdt    string `json:"tsdt" bson:"tsdt"`
	Tedt    string `json:"tedt" bson:"tedt"`
	Pow     string `json:"pow" bson:"pow"`
	Mon     string `json:"mon" bson:"mon"`
	Rcvdate string `json:"rcvdate" bson:"rcvdate"`
}

type UseRegisterRes struct {
	Result   string           `json:"result" bson:"result"`
	Rdate    string           `json:"rdate" bson:"rdate"`
	Reqcnt   string           `json:"reqcnt" bson:"reqcnt"`
	Inscnt   string           `json:"inscnt" bson:"inscnt"`
	Dupcnt   string           `json:"dupcnt" bson:"dupcnt"`
	Limitcnt string           `json:"limitcnt" bson:"limitcnt"`
	Errcnt   string           `json:"errcnt" bson:"errcnt"`
	Errcode  string           `json:"errcode" bson:"errcode"`
	Errdtl   string           `json:"errdtl" bson:"errdtl"`
	Errlist  []UseRegisterErr `json:"errlist" bson:"errlist"`
}

type UseRegisterErr struct {
	Sid    string `json:"sid" bson:"sid"`
	Cid    string `json:"cid" bson:"cid"`
	Tsdt   string `json:"tsdt" bson:"tsdt"`
	Errmsg string `json:"errmsg" bson:"errmsg"`
}

type UseDeleteReq struct {
	Bid  string      `json:"bid" bson:"bid"`
	Bkey string      `json:"bkey" bson:"bkey"`
	Use  []UseDelete `json:"use" bson:"use"`
}
type UseDelete struct {
	Sid  string `json:"sid" bson:"sid"`
	Cid  string `json:"cid" bson:"cid"`
	Tsdt string `json:"tsdt" bson:"tsdt"`
	Tedt string `json:"tedt" bson:"tedt"`
}
type UseDeleteRes struct {
	Result  string         `json:"result" bson:"result"`
	Rdate   string         `json:"rdate" bson:"rdate"`
	Reqcnt  string         `json:"reqcnt" bson:"reqcnt"`
	Delcnt  string         `json:"delcnt" bson:"delcnt"`
	Errcnt  string         `json:"errcnt" bson:"errcnt"`
	Errcode string         `json:"errcode" bson:"errcode"`
	Errdtl  string         `json:"errdtl" bson:"errdtl"`
	Errlist []UseDeleteErr `json:"errlist" bson:"errlist"`
}

type UseDeleteErr struct {
	Sid  string `json:"sid" bson:"sid"`
	Cid  string `json:"cid" bson:"cid"`
	Tsdt string `json:"tsdt" bson:"tsdt"`
	Tedt string `json:"tedt" bson:"tedt"`
}

type UseStatReq struct {
	Bid   string `json:"bid" bson:"bid"`
	Bkey  string `json:"bkey" bson:"bkey"`
	Tbase string `json:"tbase" bson:"tbase"`
	Start string `json:"start" bson:"start"`
	End   string `json:"end" bson:"end"`
}

type UseStatRes struct {
	Result string    `json:"result" bson:"result"`
	Rcnt   string    `json:"rcnt" bson:"rcnt"`
	Rdate  string    `json:"rdate" bson:"rdate"`
	Tstat  []UseStat `json:"tstat" bson:"tstat"`

	Errcode string `json:"errcode" bson:"errcode"`
	Errdtl  string `json:"errdtl" bson:"errdtl"`
}

type UseStat struct {
	Cnt    string `json:"cnt" bson:"cnt"`
	Powsum string `json:"powsum" bson:"powsum"`
	Monsum string `json:"monsum" bson:"monsum"`
}
