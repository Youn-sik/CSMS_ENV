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

func setStationArea[TR envType.StationAreaInfoInterface](sa TR) (bool, string) {
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

func updateStationArea[TR envType.StationAreaInfoInterface](sa TR) (bool, string) {
	tx, err := tidbClientEnv.Begin()
	if err != nil {
		logger.PrintErrorLogLevel4(err)
		return false, "Tx Initializing Failed"
	}
	defer tx.Rollback()

	tx.Commit()
	return true, ""
}
