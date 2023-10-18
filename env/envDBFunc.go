package envc

import (
	envType "CSMS_ENV/env/type"
	"CSMS_ENV/logger"
)

/*
	tx, err := tidbClientEnv.Begin()
	if err != nil {
		logger.PrintErrorLogLevel4(err)
		return false, "Tx Initializing Failed"
	}
	defer tx.Rollback()

	tx.Commit()
	return true, ""
*/

func setStationArea[TR envType.ChargerInfoInterface](sa TR) (bool, string) {
	tx, err := tidbClientEnv.Begin()
	if err != nil {
		logger.PrintErrorLogLevel4(err)
		return false, "Tx Initializing Failed"
	}
	defer tx.Rollback()

	stationArea := sa.GetStationAreaInfo()
	for _, val := range stationArea {
		_, err = tx.Exec("insert into station_area (bid, sid, zcode, station_area_name, addr, addrdtl, daddr, daddrdtl, kind, kinddtl, gps, usetime) "+
			"values(?,?,?,?,?,?,?,?,?,?,?,?)", val.Bid, val.Sid, val.Zcode, val.StationAreaName, val.Addr, val.Addrdtl, val.Daddr, val.Daddrdtl,
			val.Kind, val.Kinddtl, val.Gps, val.UseTime)
		if err != nil {
			logger.PrintErrorLogLevel4(err)
			return false, "Failed To Insert Data"
		}
	}

	tx.Commit()
	return true, ""
}

func updateStationArea[TR envType.ChargerInfoInterface](sa TR) (bool, string) {
	tx, err := tidbClientEnv.Begin()
	if err != nil {
		logger.PrintErrorLogLevel4(err)
		return false, "Tx Initializing Failed"
	}
	defer tx.Rollback()

	stationArea := sa.GetStationAreaInfo()
	for _, val := range stationArea {
		_, err = tx.Exec("delete from station_area where bid = ? and sid = ?", val.Bid, val.Sid)
		if err != nil {
			logger.PrintErrorLogLevel4(err)
			return false, "Failed To Delete Data"
		}
		_, err = tx.Exec("insert into station_area (bid, sid, zcode, station_area_name, addr, addrdtl, daddr, daddrdtl, kind, kinddtl, gps, usetime) "+
			"values(?,?,?,?,?,?,?,?,?,?,?,?) ", val.Bid, val.Sid, val.Zcode, val.StationAreaName, val.Addr, val.Addrdtl, val.Daddr, val.Daddrdtl,
			val.Kind, val.Kinddtl, val.Gps, val.UseTime)
		if err != nil {
			logger.PrintErrorLogLevel4(err)
			return false, "Failed To Insert Data"
		}
		// 또는 아래 코드
		/*
			_, err = tx.Exec("update station_area "+
				" set zcode = ?, station_area_name = ?, addr = ?, addrdtl = ?, daddr = ?, daddrdtl = ?, kind = ?, kinddtl = ?, gps = ?, usetime = ? "+
				"where bid = ? and sid = ?",
				val.Zcode, val.StationAreaName, val.Addr, val.Addrdtl, val.Daddr, val.Daddrdtl, val.Kind, val.Kinddtl, val.Gps, val.UseTime,
				val.Bid, val.Sid)
			if err != nil {
				logger.PrintErrorLogLevel4(err)
				return false, "Failed To Update Data"
			}
		*/
	}

	tx.Commit()
	return true, ""
}

func setChargePoint[TR envType.ChargerInfoInterface](cp TR) (bool, string) {
	tx, err := tidbClientEnv.Begin()
	if err != nil {
		logger.PrintErrorLogLevel4(err)
		return false, "Tx Initializing Failed"
	}
	defer tx.Rollback()

	chargePoint := cp.GetChargerInfo()
	for _, val := range chargePoint {
		_, err = tx.Exec("insert into charge_point (bid, sid, cid, zcode, station_area_name, addr, addrdtl, daddr, daddrdtl, kind, kinddtl, gps, usetime, "+
			"free, freedtl, bname, bcall, type, reserv, member, pay, fee, cable, status, statdt, note, "+
			"bmngid, limityn, limitdetail, delyn, deldetail, last_tsdt, last_tedt, now_tsdt, method, output) "+
			"values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ",
			val.Bid, val.Sid, val.Cid, val.Zcode, val.Name, val.Addr, val.Addrdtl, val.Daddr, val.Daddrdtl, val.Kind, val.Kinddtl, val.Gps, val.Usetime,
			val.Free, val.Freedtl, val.Bname, val.Bcall, val.Type, val.Reserv, val.Member, val.Pay, val.Fee, val.Cable, val.Status, val.Statdt, val.Note,
			val.Bmngid, val.Limityn, val.Limitdetail, val.Delyn, val.Deldetail, val.LastTsdt, val.LastTedt, val.NowTsdt, val.Method, val.Output)
		if err != nil {
			logger.PrintErrorLogLevel4(err)
			return false, "Failed To Insert Data"
		}
	}

	tx.Commit()
	return true, ""
}

func updateChargePoint[TR envType.ChargerInfoInterface](cp TR) (bool, string) {
	tx, err := tidbClientEnv.Begin()
	if err != nil {
		logger.PrintErrorLogLevel4(err)
		return false, "Tx Initializing Failed"
	}
	defer tx.Rollback()

	chargePoint := cp.GetChargerInfo()
	for _, val := range chargePoint {
		_, err = tx.Exec("delete from charge_point where bid = ? and cid = ?", val.Bid, val.Cid)
		if err != nil {
			logger.PrintErrorLogLevel4(err)
			return false, "Failed To Delete Data"
		}
		_, err = tx.Exec("insert into charge_point (bid, sid, cid, zcode, station_area_name, addr, addrdtl, daddr, daddrdtl, kind, kinddtl, gps, usetime, "+
			"free, freedtl, bname, bcall, type, reserv, member, pay, fee, cable, status, statdt, note, "+
			"bmngid, limityn, limitdetail, delyn, deldetail, last_tsdt, last_tedt, now_tsdt, method, output) "+
			"values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
			val.Bid, val.Sid, val.Cid, val.Zcode, val.Name, val.Addr, val.Addrdtl, val.Daddr, val.Daddrdtl, val.Kind, val.Kinddtl, val.Gps, val.Usetime,
			val.Free, val.Freedtl, val.Bname, val.Bcall, val.Type, val.Reserv, val.Member, val.Pay, val.Fee, val.Cable, val.Status, val.Statdt, val.Note,
			val.Bmngid, val.Limityn, val.Limitdetail, val.Delyn, val.Deldetail, val.LastTsdt, val.LastTedt, val.NowTsdt, val.Method, val.Output)
		if err != nil {
			logger.PrintErrorLogLevel4(err)
			return false, "Failed To Insert Data"
		}
		// 또는 아래 코드
		/*
			_, err = tx.Exec("update charge_point set sid = ?, zcode = ?, station_area_name = ?, addr = ?, addrdtl = ?, "+
				"daddr = ?, daddrdtl = ?, kind = ?, kinddtl = ?, gps = ?, usetime = ?, "+
				"free = ?, freedtl = ?, bname = ?, bcall = ?, type = ?, reserv = ?, member = ?, pay = ?, fee = ?, cable = ?, status = ?, statdt = ?, note = ?, "+
				"bmngid = ?, limityn = ?, limitdetail = ?, delyn = ?, deldetail = ?, last_tsdt = ?, last_tedt = ?, now_tsdt = ?, method = ?, output = ?) "+
				"where bid = ? and cid = ?",
				val.Sid, val.Zcode, val.Name, val.Addr, val.Addrdtl,
				val.Daddr, val.Daddrdtl, val.Kind, val.Kinddtl, val.Gps, val.Usetime,
				val.Free, val.Freedtl, val.Bname, val.Bcall, val.Type, val.Reserv, val.Member, val.Pay, val.Fee, val.Cable, val.Status, val.Statdt, val.Note,
				val.Bmngid, val.Limityn, val.Limitdetail, val.Delyn, val.Deldetail, val.LastTsdt, val.LastTedt, val.NowTsdt, val.Method, val.Output, val.Bid, val.Cid)
			if err != nil {
				logger.PrintErrorLogLevel4(err)
				return false, "Failed To Update Data"
			}
		*/
	}

	tx.Commit()
	return true, ""
}

func updateChargePointStatus[TR envType.ChargePointStatusInterface](cpst TR) (bool, string) {
	tx, err := tidbClientEnv.Begin()
	if err != nil {
		logger.PrintErrorLogLevel4(err)
		return false, "Tx Initializing Failed"
	}
	defer tx.Rollback()

	chargePoint := cpst.GetChargerStatus()
	for _, val := range chargePoint {
		_, err = tx.Exec("update charge_point set status = ?, statdt = ?, last_tsdt = ?, last_tedt = ?, now_tsdt = ? "+
			"where bid = ? and sid = ? and cid = ?",
			val.Status, val.Statdt, val.LastTsdt, val.LastTedt, val.NowTsdt, val.Bid, val.Sid, val.Cid)
		if err != nil {
			logger.PrintErrorLogLevel4(err)
			return false, "Failed To Update Data"
		}
	}

	tx.Commit()
	return true, ""
}

func setStationAreaMy[TR envType.ChargerInfoMyInterface](sa TR) (bool, string) {
	tx, err := tidbClientEnv.Begin()
	if err != nil {
		logger.PrintErrorLogLevel4(err)
		return false, "Tx Initializing Failed"
	}
	defer tx.Rollback()

	stationArea := sa.GetStationAreaInfo()
	for _, val := range stationArea {
		_, err = tx.Exec("insert into station_area (bid, sid, zcode, station_area_name, addr, addrdtl, daddr, daddrdtl, kind, kinddtl, gps, usetime) "+
			"values(?,?,?,?,?,?,?,?,?,?,?,?)", val.Bid, val.Sid, val.Zcode, val.StationAreaName, val.Addr, val.Addrdtl, val.Daddr, val.Daddrdtl,
			val.Kind, val.Kinddtl, val.Gps, val.UseTime)
		if err != nil {
			logger.PrintErrorLogLevel4(err)
			return false, "Failed To Insert Data"
		}
	}

	tx.Commit()
	return true, ""
}

func setChargePointMy[TR envType.ChargerInfoMyInterface](cp TR) (bool, string) {
	tx, err := tidbClientEnv.Begin()
	if err != nil {
		logger.PrintErrorLogLevel4(err)
		return false, "Tx Initializing Failed"
	}
	defer tx.Rollback()

	chargePoint := cp.GetChargerInfoMy()
	for _, val := range chargePoint {
		_, err = tx.Exec("insert into charge_point (bid, sid, cid, zcode, station_area_name, addr, addrdtl, daddr, daddrdtl, kind, kinddtl, gps, usetime, "+
			"free, freedtl, bname, bcall, type, reserv, member, pay, fee, cable, status, statdt, note, "+
			"bmngid, limityn, limitdetail, delyn, deldetail, last_tsdt, last_tedt, now_tsdt, method, output, "+
			"expireddt, year, stype, maker, subsid, owner, sign, setup, month, subsidy, motie, regdate, upddate) "+
			"values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?) ",
			val.Bid, val.Sid, val.Cid, val.Zcode, val.Name, val.Addr, val.Addrdtl, val.Daddr, val.Daddrdtl, val.Kind, val.Kinddtl, val.Gps, val.Usetime,
			val.Free, val.Freedtl, val.Bname, val.Bcall, val.Type, val.Reserv, val.Member, val.Pay, val.Fee, val.Cable, val.Status, val.Statdt, val.Note,
			val.Bmngid, val.Limityn, val.Limitdetail, val.Delyn, val.Deldetail, val.LastTsdt, val.LastTedt, val.NowTsdt, val.Method, val.Output,
			val.Expireddt, val.Year, val.Stype, val.Maker, val.Subsid, val.Owner, val.Sign, val.Setup, val.Month, val.Subsidy, val.Motie, val.Regdate, val.Upddate)
		if err != nil {
			logger.PrintErrorLogLevel4(err)
			return false, "Failed To Insert Data"
		}
	}

	tx.Commit()
	return true, ""
}

func setEnvCard[TR envType.CardInfoInterface](ci TR) (bool, string) {
	tx, err := tidbClientEnv.Begin()
	if err != nil {
		logger.PrintErrorLogLevel4(err)
		return false, "Tx Initializing Failed"
	}
	defer tx.Rollback()

	card := ci.GetEnvCardInfo()
	for _, val := range card {
		_, err = tx.Exec("insert into env_card(bid, no, stop, regdate, upddate) values(?,?,?,?,?)",
			val.Bid, val.No, val.Stop, val.Regdate, val.Upddate)
		if err != nil {
			logger.PrintErrorLogLevel4(err)
			return false, "Failed To Insert Data"
		}
	}

	tx.Commit()
	return true, ""
}

func updateEnvCard[TR envType.CardInfoInterface](ci TR) (bool, string) {
	tx, err := tidbClientEnv.Begin()
	if err != nil {
		logger.PrintErrorLogLevel4(err)
		return false, "Tx Initializing Failed"
	}
	defer tx.Rollback()

	card := ci.GetEnvCardInfo()
	for _, val := range card {
		_, err = tx.Exec("delete from env_card where bid = ? and no = ?", val.Bid, val.No)
		if err != nil {
			logger.PrintErrorLogLevel4(err)
			return false, "Failed To Delete Data"
		}
		_, err = tx.Exec("insert into env_card(bid, no, stop, regdate, upddate) values(?,?,?,?,?)",
			val.Bid, val.No, val.Stop, val.Regdate, val.Upddate)
		if err != nil {
			logger.PrintErrorLogLevel4(err)
			return false, "Failed To Insert Data"
		}
		// 또는 아래 코드
		/*
			_, err = tx.Exec("update env_card set stop = ?, regdate = ?, upddate = ? where bid = ? and no = ?",
				val.Stop, val.Regdate, val.Upddate, val.Bid, val.No)
			if err != nil {
				logger.PrintErrorLogLevel4(err)
				return false, "Failed To Update Data"
			}
		*/
	}

	tx.Commit()
	return true, ""
}
