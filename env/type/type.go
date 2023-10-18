package envType

func (c ChargerInfoListRes) GetStationAreaInfo() []StationAreaInfo {
	saiArr := []StationAreaInfo{}
	for _, val := range c.Cinfo {
		var sai StationAreaInfo
		sai.StationAreaId = "0"
		sai.Bid = val.Bid
		sai.Sid = val.Sid
		sai.Zcode = val.Zcode
		sai.StationAreaName = val.Name
		sai.Addr = val.Addr
		sai.Addrdtl = val.Addrdtl
		sai.Daddr = val.Daddr
		sai.Daddrdtl = val.Daddrdtl
		sai.Kind = val.Kind
		sai.Kinddtl = val.Kinddtl
		sai.Gps = val.Gps
		sai.UseTime = val.Usetime

		saiArr = append(saiArr, sai)
	}
	return saiArr
}

func (c ChargerInfoListAllRes) GetStationAreaInfo() []StationAreaInfo {
	saiArr := []StationAreaInfo{}
	for _, val := range c.Cinfo {
		var sai StationAreaInfo
		sai.StationAreaId = "0"
		sai.Bid = val.Bid
		sai.Sid = val.Sid
		sai.Zcode = val.Zcode
		sai.StationAreaName = val.Name
		sai.Addr = val.Addr
		sai.Addrdtl = val.Addrdtl
		sai.Daddr = val.Daddr
		sai.Daddrdtl = val.Daddrdtl
		sai.Kind = val.Kind
		sai.Kinddtl = val.Kinddtl
		sai.Gps = val.Gps
		sai.UseTime = val.Usetime

		saiArr = append(saiArr, sai)
	}
	return saiArr
}

func (c ChargerInfoListRes) GetChargerInfo() []ChargerInfo {
	return c.Cinfo
}

func (c ChargerInfoListAllRes) GetChargerInfo() []ChargerInfo {
	return c.Cinfo
}

func (c ChargerStatusListRes) GetChargerStatus() []ChargerInfo {
	return c.Cstat
}

func (c ChargerStatusListAllRes) GetChargerStatus() []ChargerInfo {
	return c.Cstat
}

func (c ChargerInfoMyListAllRes) GetStationAreaInfo() []StationAreaInfo {
	saiArr := []StationAreaInfo{}
	for _, val := range c.Cinfo {
		var sai StationAreaInfo
		sai.StationAreaId = "0"
		sai.Bid = val.Bid
		sai.Sid = val.Sid
		sai.Zcode = val.Zcode
		sai.StationAreaName = val.Name
		sai.Addr = val.Addr
		sai.Addrdtl = val.Addrdtl
		sai.Daddr = val.Daddr
		sai.Daddrdtl = val.Daddrdtl
		sai.Kind = val.Kind
		sai.Kinddtl = val.Kinddtl
		sai.Gps = val.Gps
		sai.UseTime = val.Usetime

		saiArr = append(saiArr, sai)
	}
	return saiArr
}

func (c ChargerInfoMyListAllRes) GetChargerInfoMy() []ChargerInfoMy {
	return c.Cinfo
}

func (c CardListRes) GetEnvCardInfo() []Card {
	return c.Card
}

func (c CardListAllRes) GetEnvCardInfo() []Card {
	return c.Card
}
