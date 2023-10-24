package envType

// 4. 회원카드 관리
type Card struct {
	Bid     string `json:"bid" bson:"bid"`
	No      string `json:"no" bson:"no"`
	Stop    string `json:"stop" bson:"stop"`
	Regdate string `json:"regdate" bson:"regdate"`
	Upddate string `json:"upddate" bson:"upddate"`
}

type CardListReq struct {
	Bid    string `json:"bid" bson:"bid"`
	Bkey   string `json:"bkey" bson:"bkey"`
	Kind   string `json:"kind" bson:"kind"`
	Bdate  string `json:"bdate" bson:"bdate"`
	Rbid   string `json:"rbid" bson:"rbid"`
	Cardno string `json:"cardno" bson:"sid"`
}

type CardListRes struct {
	Result string `json:"result" bson:"result"`
	Rdate  string `json:"rdate" bson:"rdate"`
	Rowcnt string `json:"rowcnt" bson:"rowcnt"`
	Card   []Card `json:"card" bson:"card"`

	Errcode string `json:"errcode" bson:"errcode"`
	Errdlt  string `json:"errdlt" bson:"Errdlt"`
}

type CardListAllReq struct {
	Bid    string `json:"bid" bson:"bid"`
	Bkey   string `json:"bkey" bson:"bkey"`
	Kind   string `json:"kind" bson:"kind"`
	Rbid   string `json:"rbid" bson:"rbid"`
	Pageno string `json:"pageno" bson:"pageno"`
	Rowcnt string `json:"rowcnt" bson:"rowcnt"`
}

type CardListAllRes struct {
	Result   string `json:"result" bson:"result"`
	Rdate    string `json:"rdate" bson:"rdate"`
	Totalcnt string `json:"totalcnt" bson:"totalcnt"`
	Pageno   string `json:"pageno" bson:"pageno"`
	Rowcnt   string `json:"rowcnt" bson:"rowcnt"`
	Card     []Card `json:"card" bson:"card"`

	Errcode string `json:"errcode" bson:"errcode"`
	Errdlt  string `json:"errdlt" bson:"Errdlt"`
}

type CardUpdateReq struct {
	Bid  string       `json:"bid" bson:"bid"`
	Bkey string       `json:"bkey" bson:"bkey"`
	Card []CardNoStop `json:"card" bson:"card"`
}

type CardNoStop struct {
	No   string `json:"no" bson:"no"`
	Stop string `json:"stop" bson:"stop"`
}

type CardUpdateRes struct {
	Result   string          `json:"result" bson:"result"`
	Rdate    string          `json:"rdate" bson:"rdate"`
	Reqcnt   string          `json:"reqcnt" bson:"reqcnt"`
	Updcnt   string          `json:"updcnt" bson:"updcnt"`
	Dupcnt   string          `json:"dupcnt" bson:"dupcnt"`
	Limitcnt string          `json:"limitcnt" bson:"limitcnt"`
	Errcnt   string          `json:"errcnt" bson:"errcnt"`
	Errcode  string          `json:"errcode" bson:"errcode"`
	Errdlt   string          `json:"errdlt" bson:"Errdlt"`
	Errlist  []CardUpdateErr `json:"errlist" bson:"errlist"`
}

type CardUpdateErr struct {
	No     string `json:"no" bson:"no"`
	Errmsg string `json:"errmsg" bson:"errmsg"`
}

type CardStatReq struct {
	Bid  string `json:"bid" bson:"bid"`
	Bkey string `json:"bkey" bson:"bkey"`
	Rbid string `json:"rbid" bson:"rbid"`
}

type CardStatRes struct {
	Result string     `json:"result" bson:"result"`
	Cstat  []CardStat `json:"cstat" bson:"cstat"`

	Errcode string `json:"errcode" bson:"errcode"`
	Errdlt  string `json:"errdlt" bson:"Errdlt"`
}

type CardStat struct {
	Bid  string `json:"bid" bson:"bid"`
	Stop string `json:"stop" bson:"stop"`
	Cnt  string `json:"cnt" bson:"cnt"`
}
