package envc

import (
	"log"
	"strconv"

	envType "CSMS_ENV/env/type"
	"CSMS_ENV/logger"
)

// 2.1 충전소 조회 (자세한 내용은 환경부 연동 문서 붙임 1 참고)
// 해당 요청은 '이전 요청 이후 충전소 변경 분 정보를 조회'합니다.
// 요청이 없을 때, '최대 12시간 이내 변경 정보'만 조회합니다.
func ChargerInfoList() {
	var req envType.ChargerInfoListReq
	req.Bid = bid
	req.Bkey = bkey
	req.Kind = "1" // 자사 제외
	req.Rbid = ""

	res, err := envHttpRequest[envType.ChargerInfoListReq, envType.ChargerInfoListRes]("/charger/info/list", req)
	if err != nil {
		// 실패 에러 처리
		logger.PrintErrorLogLevel4(err)
		return
	}
	if res.Result == "2" || res.Result == "1" {
		// (부분) 실패 에러 처리
		log.Println("--- Failed To Request To Env ---")
		log.Printf("req: %+v", req)
		log.Printf("res: %+v", res)
		return
	}

	// 성공 처리
	// updateStationArea[envType.ChargerInfoListRes](res)
	updateChargePoint[envType.ChargerInfoListRes](res)
}

// 2.2 충전소 전체 조회 (자세한 내용은 환경부 연동 문서 붙임 1 참고)
// 해당 요청은 '충전소 전체 정보'를 조회합니다. '페이지 조건을 통해 일정 건수로 조회'합니다.
// 최초 요청 시 페이지 및 전체 충전소 개수에 따라 재요청이 필요합니다.
func ChargerInfoListAll(p envType.PageNoRowCnt) {
	var req envType.ChargerInfoListAllReq
	req.Bid = bid
	req.Bkey = bkey
	req.Kind = "1" // 자사 제외
	req.Rbid = ""
	req.Pageno = strconv.Itoa(p.PageNo)
	req.Rowcnt = p.RowCnt

	res, err := envHttpRequest[envType.ChargerInfoListAllReq, envType.ChargerInfoListAllRes]("/charger/info/listall", req)
	if err != nil {
		// 실패 에러 처리
		logger.PrintErrorLogLevel4(err)
		return
	}
	if res.Result == "2" || res.Result == "1" {
		// (부분) 실패 에러 처리
		log.Println("--- Failed To Request To Env ---")
		log.Printf("req: %+v", req)
		log.Printf("res: %+v", res)
		return
	}

	// 성공 처리
	// setStationArea[envType.ChargerInfoListAllRes](res)
	setChargePoint[envType.ChargerInfoListAllRes](res)

	if (res.Totalcnt - p.RowCnt*p.PageNo) > 0 {
		p.PageNo++
		ChargerInfoListAll(p)
	}
}

// 2.3 충전기 상태 조회 (자세한 내용은 환경부 연동 문서 붙임 1 참고)
// 해당 요청은 '충전기 상태 정보'를 조회합니다. '마지막 요청 이후 변경된 상태 정보'만 조회합니다.
// 30분 이상 상태가 갱신되지 않은 경우 상태 미확인(9)로 조회됩니다.
// 이전 요청이 없을 때, '최대 10분 이내 변경 정보'만 조회합니다.
func ChargerStatusList() {
	var req envType.ChargerStatusListReq
	req.Bid = bid
	req.Bkey = bkey
	req.Kind = "1" // 자사 제외
	req.Rbid = ""

	res, err := envHttpRequest[envType.ChargerStatusListReq, envType.ChargerStatusListRes]("/charger/status/list", req)
	if err != nil {
		// 실패 에러 처리
		logger.PrintErrorLogLevel4(err)
		return
	}
	if res.Result == "2" || res.Result == "1" {
		// (부분) 실패 에러 처리
		log.Println("--- Failed To Request To Env ---")
		log.Printf("req: %+v", req)
		log.Printf("res: %+v", res)
		return
	}

	// 성공 처리
	updateChargePointStatus[envType.ChargerStatusListRes](res)
}

// 2.4 충전기 상태 전체 조회 (자세한 내용은 환경부 연동 문서 붙임 1 참고)
// 해당 요청은 '충전기 상태 전체 정보'를 조회합니다. '페이지 조건을 통해 일정 건수로 조회'합니다.
// 1분 이내 충전기 상태가 조회됩니다.
// 최초 요청 시 페이지 및 전체 충전기 개수에 따라 재요청이 필요합니다.
func ChargerStatusListAll(p envType.PageNoRowCnt) {
	var req envType.ChargerStatusListAllReq
	req.Bid = bid
	req.Bkey = bkey
	req.Kind = "1" // 자사 제외
	req.Rbid = ""
	req.Pageno = strconv.Itoa(p.PageNo)
	req.Rowcnt = p.RowCnt

	res, err := envHttpRequest[envType.ChargerStatusListAllReq, envType.ChargerStatusListAllRes]("/charger/status/listall", req)
	if err != nil {
		// 실패 에러 처리
		logger.PrintErrorLogLevel4(err)
		return
	}
	if res.Result == "2" || res.Result == "1" {
		// (부분) 실패 에러 처리
		log.Println("--- Failed To Request To Env ---")
		log.Printf("req: %+v", req)
		log.Printf("res: %+v", res)
		return
	}

	// 성공 처리
	updateChargePointStatus[envType.ChargerStatusListAllRes](res)

	if (res.Totalcnt - p.RowCnt*p.PageNo) > 0 {
		p.PageNo++
		ChargerStatusListAll(p)
	}
}

// 3.1 충전기 관리 정보 조회 (자세한 내용은 환경부 연동 문서 붙임 1 참고)
// 해당 요청은 '사업자가 관리하는 충전기 정보'를 조회합니다.
// 응답 건수가 5,000 건 초과인 경우 조회가 불가능합니다.
func ChargerInfoMyList(p envType.PageNoRowCnt) {
	var req envType.ChargerInfoMyListAllReq
	req.Bid = bid
	req.Bkey = bkey
	req.Kind = "1" // 자사 전체 조회
	req.Pageno = strconv.Itoa(p.PageNo)
	req.Rowcnt = p.RowCnt

	res, err := envHttpRequest[envType.ChargerInfoMyListAllReq, envType.ChargerInfoMyListAllRes]("/charger/info/mylist", req)
	if err != nil {
		// 실패 에러 처리
		logger.PrintErrorLogLevel4(err)
		return
	}
	if res.Result == "2" || res.Result == "1" {
		// (부분) 실패 에러 처리
		log.Println("--- Failed To Request To Env ---")
		log.Printf("req: %+v", req)
		log.Printf("res: %+v", res)
		return
	}

	// 성공 처리
	// setStationAreaMy[envType.ChargerInfoMyListAllRes](res)
	setChargePointMy[envType.ChargerInfoMyListAllRes](res)

	if (res.Totalcnt - p.RowCnt*p.PageNo) > 0 {
		p.PageNo++
		ChargerInfoMyList(p)
	}
}

// 3.2 충전기 상태 갱신 (자세한 내용은 환경부 연동 문서 붙임 1 참고)
// 해당 요청은 '사업자가 관리하는 충전기 상태를 갱신'합니다.
// 운영기관(자사)의 충전기 상태가 변경되는 경우, 1분 이내에 갱신합니다.
// 갱신 이후 30분이 지난 경우, '상태미확인' 상태로 서비스에 제공됩니다.
func ChargerStatusUpdate(cstat []envType.ChargerStatus) {
	var req envType.ChargerStatusUpdateReq
	req.Bid = bid
	req.Bkey = bkey
	req.Cstat = cstat

	res, err := envHttpRequest[envType.ChargerStatusUpdateReq, envType.ChargerStatusUpdateRes]("/charger/status/update", req)
	if err != nil {
		// 실패 에러 처리
		logger.PrintErrorLogLevel4(err)
		return
	}
	if res.Result == "2" || res.Result == "1" {
		// (부분) 실패 에러 처리
		log.Println("--- Failed To Request To Env ---")
		log.Printf("req: %+v", req)
		log.Printf("res: %+v", res)
		return
	}

	// 성공 처리
}

// 4.1 회원카드 조회 (자세한 내용은 환경부 연동 문서 붙임 1 참고)
// 해당 요청은 '공동사용 협약한 사용자간 회원카드'를 조회합니다.
// 타 사업자가 보유한 회원카드의 등록 결과를 조회합니다.
// kind = 1 은 '이전 요청 이후 회원카드 변경 분 정보'를 조회합니다.
// 요청이 없으면, 24시간 이내 변경 정보만 조회합니다.
func CardList() {
	var req envType.CardListReq
	req.Bid = bid
	req.Bkey = bkey
	req.Kind = "1" // 자사 제외
	req.Rbid = ""

	res, err := envHttpRequest[envType.CardListReq, envType.CardListRes]("/card/list", req)
	if err != nil {
		// 실패 에러 처리
		logger.PrintErrorLogLevel4(err)
		return
	}
	if res.Result == "2" || res.Result == "1" {
		// (부분) 실패 에러 처리
		log.Println("--- Failed To Request To Env ---")
		log.Printf("req: %+v", req)
		log.Printf("res: %+v", res)
		return
	}

	// 성공 처리
	updateEnvCard[envType.CardListRes](res)
}

// 4.2 회원카드 전체 조회 (자세한 내용은 환경부 연동 문서 붙임 1 참고)
// 해당 요청은 '회원카드 정체 정보'를 조회합니다. '페이지 조건을 통해 일정 건수로 조회'합니다.
// 최초 요청 시 페이지 및 전체 충전소 개수에 따라 재요청이 필요합니다.
func CardListAll(p envType.PageNoRowCnt) {
	var req envType.CardListAllReq
	req.Bid = bid
	req.Bkey = bkey
	req.Kind = "1" // 자사 제외
	req.Rbid = ""
	req.Pageno = strconv.Itoa(p.PageNo)
	req.Rowcnt = p.RowCnt

	res, err := envHttpRequest[envType.CardListAllReq, envType.CardListAllRes]("/card/listall", req)
	if err != nil {
		// 실패 에러 처리
		logger.PrintErrorLogLevel4(err)
		return
	}
	if res.Result == "2" || res.Result == "1" {
		// (부분) 실패 에러 처리
		log.Println("--- Failed To Request To Env ---")
		log.Printf("req: %+v", req)
		log.Printf("res: %+v", res)
		return
	}

	// 성공 처리
	setEnvCard[envType.CardListAllRes](res)

	if (res.Totalcnt - p.RowCnt*p.PageNo) > 0 {
		p.PageNo++
		CardListAll(p)
	}
}

// 4.3 회원카드 갱신 (자세한 내용은 환경부 연동 문서 붙임 1 참고)
// 해당 요청은 '회원카드 상태가 변경되면 갱신'합니다.
// 정지 회원이 다른 사업자 충전기에서 인증이 될 수도 있습니다.
// 회원카드 변경 정보를 10분 이내에 갱신하고, 갱신 누락을 대비하여 하루 한 번 전체 갱신을 권고합니다.
// 한 번에 2000건 까지 갱신 가능합니다.
func CardUpdate(cardNoStop []envType.CardNoStop) {
	var req envType.CardUpdateReq
	req.Bid = bid
	req.Bkey = bkey
	req.Card = cardNoStop

	res, err := envHttpRequest[envType.CardUpdateReq, envType.CardUpdateRes]("/card/update", req)
	if err != nil {
		// 실패 에러 처리
		logger.PrintErrorLogLevel4(err)
		return
	}
	if res.Result == "2" || res.Result == "1" {
		// (부분) 실패 에러 처리
		log.Println("--- Failed To Request To Env ---")
		log.Printf("req: %+v", req)
		log.Printf("res: %+v", res)
		return
	}

	// 성공 처리
}

// 4.4 회원카드 통계 조회 (자세한 내용은 환경부 연동 문서 붙임 1 참고)
// 해당 요청은 '회원카드 사용자의 상태(정지)별 건수'를 조회합니다.
func CardStat() {
	var req envType.CardStatReq
	req.Bid = bid
	req.Bkey = bkey
	req.Rbid = ""

	res, err := envHttpRequest[envType.CardStatReq, envType.CardStatRes]("/card/stat", req)
	if err != nil {
		// 실패 에러 처리
		logger.PrintErrorLogLevel4(err)
		return
	}
	if res.Result == "2" || res.Result == "1" {
		// (부분) 실패 에러 처리
		log.Println("--- Failed To Request To Env ---")
		log.Printf("req: %+v", req)
		log.Printf("res: %+v", res)
		return
	}

	// 성공 처리
}

// 5.1 충전이력 조회 (자세한 내용은 환경부 연동 문서 붙임 1 참고)
// 해당 요청은 '로밍에 등록 된 시간 기준'으로 조회합니다.
// 마지막 요청 이후 충전이력 등록 정보를 조회합니다.
// 이전 요청이 없으면, 24시간 이내 등록된 정보가지만 조회합니다.
func TradeList() {
	var req envType.TradeListReq
	req.Bid = bid
	req.Bkey = bkey
	req.Kind = "1" // 자사 제외

	res, err := envHttpRequest[envType.TradeListReq, envType.TradeListRes]("/trade/list", req)
	if err != nil {
		// 실패 에러 처리
		logger.PrintErrorLogLevel4(err)
		return
	}
	if res.Result == "2" || res.Result == "1" {
		// (부분) 실패 에러 처리
		log.Println("--- Failed To Request To Env ---")
		log.Printf("req: %+v", req)
		log.Printf("res: %+v", res)
		return
	}

	// 성공 처리
}

// 5.2 충전이력 전체 조회 (자세한 내용은 환경부 연동 문서 붙임 1 참고)
// 해당 요청은 '서버세 등록 된 시간 기준'으로 조회합니다.
// 건수가 많을 경우 페이지 단위로 조회 가능합니다.
// 최초 요청 시 페이지 및 전체 충전소 개수에 따라 재요청이 필요합니다.
func TradeListAll(p envType.PageNoRowCnt) {
	var req envType.TradeListAllReq
	req.Bid = bid
	req.Bkey = bkey
	req.Kind = "1" // 자사 제외
	req.Rbid = ""
	req.Start = "2023100100000000"
	req.End = "2023101800000000"
	req.Pageno = strconv.Itoa(p.PageNo)
	req.Rowcnt = p.RowCnt

	res, err := envHttpRequest[envType.TradeListAllReq, envType.TradeListAllRes]("/trade/listall", req)
	if err != nil {
		// 실패 에러 처리
		logger.PrintErrorLogLevel4(err)
		return
	}
	if res.Result == "2" || res.Result == "1" {
		// (부분) 실패 에러 처리
		log.Println("--- Failed To Request To Env ---")
		log.Printf("req: %+v", req)
		log.Printf("res: %+v", res)
		return
	}

	// 성공 처리
}

// 5.3 충전이력 등록 (자세한 내용은 환경부 연동 문서 붙임 1 참고)
// 해당 요청은 '회원카드 인증을 통해 발생한 충전이력'을 등록합니다.
// 한 번에 2000건 까지 등록 가능합니다.
// 잘못 전송 된 데이터는 별도 관리자 화면에서 데이터를 수정하거나, 삭제 후 재전송 합니다.
func TradeRegister(tradeRegisterInfo []envType.TradeRegister) {
	var req envType.TradeRegisterReq
	req.Bid = bid
	req.Bkey = bkey
	req.Trade = tradeRegisterInfo

	res, err := envHttpRequest[envType.TradeRegisterReq, envType.TradeRegisterRes]("/trade/regi", req)
	if err != nil {
		// 실패 에러 처리
		logger.PrintErrorLogLevel4(err)
		return
	}
	if res.Result == "2" || res.Result == "1" {
		// (부분) 실패 에러 처리
		log.Println("--- Failed To Request To Env ---")
		log.Printf("req: %+v", req)
		log.Printf("res: %+v", res)
		return
	}

	// 성공 처리
}

// 5.4 충전이력 제외 조회 (자세한 내용은 환경부 연동 문서 붙임 1 참고)
// 해당 요청은 '등록한 사업자만 조회' 가능합니다.
// 등록한 충전이력 중 이상 데이터로 검출될 경우, 정산 제외 처리됩니다.
// 제외된 데이터를 확인하여 잘못된 데이터인 경우, 관리사이트에서 데이터 삭제 후 다시 등록합니다.
// 응답 건수가 5000 건 초과인 경우 조회가 불가능합니다.
func TradeExList() {
	var req envType.TradeExlistReq
	req.Bid = bid
	req.Bkey = bkey
	req.Kind = "1"

	res, err := envHttpRequest[envType.TradeExlistReq, envType.TradeExlistRes]("/trade/exlist", req)
	if err != nil {
		// 실패 에러 처리
		logger.PrintErrorLogLevel4(err)
		return
	}
	if res.Result == "2" || res.Result == "1" {
		// (부분) 실패 에러 처리
		log.Println("--- Failed To Request To Env ---")
		log.Printf("req: %+v", req)
		log.Printf("res: %+v", res)
		return
	}

	// 성공 처리
}

// 5.5 충전이력 통계 조회 (자세한 내용은 환경부 연동 문서 붙임 1 참고)
// 해당 요청은 '충전 이력과 조회 대상 충전 이력에 대한 데이터의 건수과 충전량, 충전 금액의 합산 금액'을 조회합니다.
// 기관별, 결제 종류별, 충전량, 충전금액 합계가 조회됩니다.
func TradeStat() {
	var req envType.TradeStatReq
	req.Bid = bid
	req.Bkey = bkey
	req.Rbid = ""
	req.Start = ""
	req.End = ""

	res, err := envHttpRequest[envType.TradeStatReq, envType.TradeStatRes]("/trade/stat", req)
	if err != nil {
		// 실패 에러 처리
		logger.PrintErrorLogLevel4(err)
		return
	}
	if res.Result == "2" || res.Result == "1" {
		// (부분) 실패 에러 처리
		log.Println("--- Failed To Request To Env ---")
		log.Printf("req: %+v", req)
		log.Printf("res: %+v", res)
		return
	}

	// 성공 처리
}

// 6.1 이용정보 조회 (자세한 내용은 환경부 문서 붙임 1 참고)
// 해당 요청은 '자사 충전 이용 정보'를 조회합니다.
// 페이지 조건을 통해 일정 건수로 조회합니다.
// 조회 등록 실시 순으로 조회됩니다.
func UseList(p envType.PageNoRowCnt) {
	var req envType.UseListReq
	req.Bid = bid
	req.Bkey = bkey
	req.Kind = "1"
	req.Tbase = ""
	req.Start = ""
	req.End = ""
	req.Bdate = ""
	req.Pageno = strconv.Itoa(p.PageNo)
	req.Rowcnt = p.RowCnt

	res, err := envHttpRequest[envType.UseListReq, envType.UseListRes]("/use/list", req)
	if err != nil {
		// 실패 에러 처리
		logger.PrintErrorLogLevel4(err)
		return
	}
	if res.Result == "2" || res.Result == "1" {
		// (부분) 실패 에러 처리
		log.Println("--- Failed To Request To Env ---")
		log.Printf("req: %+v", req)
		log.Printf("res: %+v", res)
		return
	}

	// 성공 처리
}

// 6.2 이용정보 등록 (자세한 내용은 환경부 문서 붙임 1 참고)
// 해당 요청은 '자사 충전 이용 정보'를 등록합니다.
// 중복 등록 되지 않으며 (sid, cid, tsdt 가 일치하는 경우) 변경이 필요한 경우 삭제 후 재등록 합니다.
func UseRegister(useRegisterInfo []envType.UseRegister) {
	var req envType.UseRegisterReq
	req.Bid = bid
	req.Bkey = bkey
	req.Use = useRegisterInfo

	res, err := envHttpRequest[envType.UseRegisterReq, envType.UseRegisterRes]("/use/regi", req)
	if err != nil {
		// 실패 에러 처리
		logger.PrintErrorLogLevel4(err)
		return
	}
	if res.Result == "2" || res.Result == "1" {
		// (부분) 실패 에러 처리
		log.Println("--- Failed To Request To Env ---")
		log.Printf("req: %+v", req)
		log.Printf("res: %+v", res)
		return
	}

	// 성공 처리
}

// 6.3 이용정보 삭제 (자세한 내용은 환경부 문서 붙임 1 참고)
// 해당 요청은 '자사 충전 이용 정보'를 삭제합니다.
// 충전소, 충전기 ID, 충전시작일시 및 종료일시가 일치하는 정보를 삭제합니다.
func UseDelete() {
	useDeleteInfo := []envType.UseDelete{}
	var req envType.UseDeleteReq
	req.Bid = bid
	req.Bkey = bkey
	req.Use = useDeleteInfo

	res, err := envHttpRequest[envType.UseDeleteReq, envType.UseDeleteRes]("/use/del", req)
	if err != nil {
		// 실패 에러 처리
		logger.PrintErrorLogLevel4(err)
		return
	}
	if res.Result == "2" || res.Result == "1" {
		// (부분) 실패 에러 처리
		log.Println("--- Failed To Request To Env ---")
		log.Printf("req: %+v", req)
		log.Printf("res: %+v", res)
		return
	}

	// 성공 처리
}

// 6.4 이용정보 통계 조회 (자세한 내용은 환경부 문서 붙임 1 참고)
// 해당 요청은 '자사 충전기 이용 정보 건수와 충전량, 충전 금액 합산 금액'을 조회합니다.
func UseStat() {
	var req envType.UseStatReq
	req.Bid = bid
	req.Bkey = bkey
	req.Tbase = ""
	req.Start = ""
	req.End = ""

	res, err := envHttpRequest[envType.UseStatReq, envType.UseStatRes]("/use/stat", req)
	if err != nil {
		// 실패 에러 처리
		logger.PrintErrorLogLevel4(err)
		return
	}
	if res.Result == "2" || res.Result == "1" {
		// (부분) 실패 에러 처리
		log.Println("--- Failed To Request To Env ---")
		log.Printf("req: %+v", req)
		log.Printf("res: %+v", res)
		return
	}

	// 성공 처리
}
