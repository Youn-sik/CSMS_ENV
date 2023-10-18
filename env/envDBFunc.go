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
