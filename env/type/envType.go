package envType

type EnvHttpRequestInterface interface {
	ChargerInfoListReq |
		ChargerInfoListAllReq |
		ChargerStatusListReq |
		ChargerStatusListAllReq |

		ChargerInfoMyListAllReq |
		ChargerStatusUpdateReq |

		CardListReq |
		CardListAllReq |
		CardUpdateReq |
		CardStatReq |

		TradeListReq |
		TradeListAllReq |
		TradeRegisterReq |
		TradeExlistReq |
		TradeStatReq |

		UseListReq |
		UseRegisterReq |
		UseDeleteReq |
		UseStatReq
}

type EnvHttpResponseInterface interface {
	ChargerInfoListRes |
		ChargerInfoListAllRes |
		ChargerStatusListRes |
		ChargerStatusListAllRes |

		ChargerInfoMyListAllRes |
		ChargerStatusUpdateRes |

		CardListRes |
		CardListAllRes |
		CardUpdateRes |
		CardStatRes |

		TradeListRes |
		TradeListAllRes |
		TradeRegisterRes |
		TradeExlistRes |
		TradeStatRes |

		UseListRes |
		UseRegisterRes |
		UseDeleteRes |
		UseStatRes
}

type ChargerInfoInterface interface {
	EnvHttpResponseInterface
	GetStationAreaInfo() []StationAreaInfo
	GetChargerInfo() []ChargerInfo
}

type ChargePointStatusInterface interface {
	EnvHttpResponseInterface
	GetChargerStatus() []ChargerInfo
}

type ChargerInfoMyInterface interface {
	EnvHttpResponseInterface
	GetStationAreaInfo() []StationAreaInfo
	GetChargerInfoMy() []ChargerInfoMy
}

type CardInfoInterface interface {
	EnvHttpResponseInterface
	GetEnvCardInfo() []Card
}

type StationAreaInfo struct {
	StationAreaId   string `json:"station_area_id"`
	Bid             string `json:"bid"`
	Sid             string `json:"sid"`
	Zcode           string `json:"zcode"`
	StationAreaName string `json:"station_area_name"`
	Addr            string `json:"addr"`
	Addrdtl         string `json:"addrdtl"`
	Daddr           string `json:"daddr"`
	Daddrdtl        string `json:"daddrdtl"`
	Kind            string `json:"kind"`
	Kinddtl         string `json:"kinddtl"`
	Gps             string `json:"gps"`
	UseTime         string `json:"use_time"`
}

type PageNoRowCnt struct {
	PageNo int
	RowCnt int
}
