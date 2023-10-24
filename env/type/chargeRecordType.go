package envType

// 5. 충전이력 관리
type TradeListReq struct {
	Bid  string `json:"bid" bson:"bid"`
	Bkey string `json:"bkey" bson:"bkey"`
	Kind string `json:"kind" bson:"kind"`
}

type TradeListRes struct {
	Result string  `json:"result" bson:"result"`
	Rdate  string  `json:"rdate" bson:"rdate"`
	Rowcnt string  `json:"rowcnt" bson:"rowcnt"`
	Trade  []Trade `json:"trade" bson:"trade"`

	Errcode string `json:"errcode" bson:"errcode"`
	Errdlt  string `json:"errdlt" bson:"errdlt"`
}

type Trade struct {
	Bid     string `json:"bid" bson:"bid"`
	No      string `json:"no" bson:"no"`
	Sid     string `json:"sid" bson:"sid"`
	Cid     string `json:"cid" bson:"cid"`
	Tbid    string `json:"tbid" bson:"tbid"`
	Tsdt    string `json:"tsdt" bson:"tsdt"`
	Tedt    string `json:"tedt" bson:"tedt"`
	Btid    string `json:"btid" bson:"btid"`
	Pow     string `json:"pow" bson:"pow"`
	Mon     string `json:"mon" bson:"mon"`
	Bprice  string `json:"bprice" bson:"bprice"`
	Tbprice string `json:"tbprice" bson:"tbprice"`
	Bmon    string `json:"bmon" bson:"bmon"`
	Regdate string `json:"regdate" bson:"regdate"`
	Tseq    string `json:"tseq" bson:"tseq"`
}

type TradeListAllReq struct {
	Bid    string `json:"bid" bson:"bid"`
	Bkey   string `json:"bkey" bson:"bkey"`
	Kind   string `json:"kind" bson:"kind"`
	Rbid   string `json:"rbid" bson:"rbid"`
	Start  string `json:"start" bson:"start"`
	End    string `json:"end" bson:"end"`
	Pageno string `json:"pageno" bson:"pageno"`
	Rowcnt string `json:"rowcnt" bson:"rowcnt"`
}

type TradeListAllRes struct {
	Result   string  `json:"result" bson:"result"`
	Rdate    string  `json:"rdate" bson:"rdate"`
	Totalcnt string  `json:"totalcnt" bson:"totalcnt"`
	Pageno   string  `json:"pageno" bson:"pageno"`
	Rowcnt   string  `json:"rowcnt" bson:"rowcnt"`
	Trade    []Trade `json:"trade" bson:"trade"`

	Errcode string `json:"errcode" bson:"errcode"`
	Errdlt  string `json:"errdlt" bson:"errdlt"`
}

type TradeRegisterReq struct {
	Bid   string          `json:"bid" bson:"bid"`
	Bkey  string          `json:"bkey" bson:"bkey"`
	Trade []TradeRegister `json:"trade" bson:"trade"`
}
type TradeRegister struct {
	No      string `json:"no" bson:"no"`
	Sid     string `json:"sid" bson:"sid"`
	Cid     string `json:"cid" bson:"cid"`
	Tbid    string `json:"tbid" bson:"tbid"`
	Tsdt    string `json:"tsdt" bson:"tsdt"`
	Tedt    string `json:"tedt" bson:"tedt"`
	Btid    string `json:"btid" bson:"btid"`
	Pow     string `json:"pow" bson:"pow"`
	Mon     string `json:"mon" bson:"mon"`
	Bprice  string `json:"bprice" bson:"bprice"`
	Tbprice string `json:"tbprice" bson:"tbprice"`
	Bmon    string `json:"bmon" bson:"bmon"`
}
type TradeRegisterRes struct {
	Result   string         `json:"result" bson:"result"`
	Rdate    string         `json:"rdate" bson:"rdate"`
	Reqcnt   string         `json:"reqcnt" bson:"reqcnt"`
	Inscnt   string         `json:"inscnt" bson:"inscnt"`
	Dupcnt   string         `json:"dupcnt" bson:"dupcnt"`
	Limitcnt string         `json:"limitcnt" bson:"limitcnt"`
	Errcnt   string         `json:"errcnt" bson:"errcnt"`
	Errcode  string         `json:"errcode" bson:"errcode"`
	Errdlt   string         `json:"errdlt" bson:"errdlt"`
	Errlist  []TradeListErr `json:"errlist" bson:"errlist"`
}

type TradeListErr struct {
	No     string `json:"no" bson:"no"`
	Sid    string `json:"sid" bson:"sid"`
	Cid    string `json:"cid" bson:"cid"`
	Tbid   string `json:"tbid" bson:"tbid"`
	Tsdt   string `json:"tsdt" bson:"tsdt"`
	Btid   string `json:"btid" bson:"btid"`
	Errmsg string `json:"errmsg" bson:"errmsg"`
}

type TradeExlistReq struct {
	Bid   string `json:"bid" bson:"bid"`
	Bkey  string `json:"bkey" bson:"bkey"`
	Kind  string `json:"kind" bson:"kind"`
	Start string `json:"start" bson:"start"`
	End   string `json:"end" bson:"end"`
}

type TradeExlistRes struct {
	Result string        `json:"result" bson:"result"`
	Rdate  string        `json:"rdate" bson:"rdate"`
	Rowcnt string        `json:"rowcnt" bson:"rowcnt"`
	Trade  []TradeExlist `json:"trade" bson:"trade"`

	Errcode string `json:"errcode" bson:"errcode"`
	Errdlt  string `json:"errdlt" bson:"errdlt"`
}
type TradeExlist struct {
	Bid     string `json:"bid" bson:"bid"`
	No      string `json:"no" bson:"no"`
	Sid     string `json:"sid" bson:"sid"`
	Cid     string `json:"cid" bson:"cid"`
	Tbid    string `json:"tbid" bson:"tbid"`
	Tsdt    string `json:"tsdt" bson:"tsdt"`
	Tedt    string `json:"tedt" bson:"tedt"`
	Btid    string `json:"btid" bson:"btid"`
	Pow     string `json:"pow" bson:"pow"`
	Mon     string `json:"mon" bson:"mon"`
	Bprice  string `json:"bprice" bson:"bprice"`
	Tbprice string `json:"tbprice" bson:"tbprice"`
	Bmon    string `json:"bmon" bson:"bmon"`
	Regdate string `json:"regdate" bson:"regdate"`
	Tseq    string `json:"tseq" bson:"tseq"`
	Excode  string `json:"excode" bson:"excode"`
}

type TradeStatReq struct {
	Bid   string `json:"bid" bson:"bid"`
	Bkey  string `json:"bkey" bson:"bkey"`
	Rbid  string `json:"rbid" bson:"rbid"`
	Start string `json:"start" bson:"start"`
	End   string `json:"end" bson:"end"`
}

type TradeStatRes struct {
	Result string      `json:"result" bson:"result"`
	Tstat  []TradeStat `json:"tstat" bson:"tstat"`

	Errcode string `json:"errcode" bson:"errcode"`
	Errdlt  string `json:"errdlt" bson:"errdlt"`
}

type TradeStat struct {
	Bid     string `json:"bid" bson:"bid"`
	Cnt     string `json:"cnt" bson:"cnt"`
	Powsum  string `json:"powsum" bson:"powsum"`
	Monsum  string `json:"monsum" bson:"monsum"`
	Bmonsum string `json:"bmonsum" bson:"bmonsum"`
	Excnt   string `json:"excnt" bson:"excnt"`
}
