package envType

type EnvHttpRequestInterface interface {
	ChargerInfoListReq |
		ChargerInfoListAllReq |
		ChargerStatusListReq |
		ChargerStatusListAllReq |

		ChargerInfoMyListReq |
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

		ChargerInfoMyListRes |
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

type StationAreaInfoInterface interface {
	EnvHttpResponseInterface
	GetStationAreaInfo() []StationAreaInfo
	GetChargerInfo() []ChargerInfo
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
