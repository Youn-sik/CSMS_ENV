package main

import (
	"log"
	"net"

	envc "CSMS_ENV/env"
	envType "CSMS_ENV/env/type"
	rpc "CSMS_ENV/grpc"
	pb "CSMS_ENV/grpc/adminService"

	"github.com/robfig/cron"
	"google.golang.org/grpc"
)

var pageNoRowCnt = envType.PageNoRowCnt{
	PageNo: 1,
	RowCnt: 1000,
}
var port = ":9080"

func main() {
	cr := cron.New()

	// 2.1 -> 1시간 마다
	cr.AddFunc("@every 1h", func() {
		go envc.ChargerInfoList()
	})

	// 2.2 -> 최초 1회
	go envc.ChargerInfoListAll(pageNoRowCnt)

	// 2.3 -> 10분 마다
	cr.AddFunc("@every 10m", func() {
		go envc.ChargerStatusList()
	})

	// 2.4 -> 최초 1회
	go envc.ChargerStatusListAll(pageNoRowCnt)

	// 3.1 -> 24시간 마다
	cr.AddFunc("@every 24h", func() {
		go envc.ChargerInfoMyList(pageNoRowCnt)
	})

	// 3.2 -> 10분 마다 (가장 최근 상태를 수신) + 이벤트 기반 (chargePoint의 StatusNotification)
	cr.AddFunc("@every 10m", func() {
		cstat := []envType.ChargerStatus{} // 모든 충전기의 상태 값 : []envType.ChargerStatus{} -> GetChargePointStatus(from. gRPC)
		go envc.ChargerStatusUpdate(cstat)
	})

	// 4.1 -> 10분 마다
	cr.AddFunc("@every 10m", func() {
		go envc.CardList()
	})

	// 4.2 -> 24시간 마다
	cr.AddFunc("@every 24h", func() {
		go envc.CardListAll(pageNoRowCnt)
	})

	// 4.3 -> 24시간 마다
	cr.AddFunc("@every 24h", func() {
		cardNoStop := []envType.CardNoStop{} // 모든 사용자의 카드 정보 값 : []envType.CardNoStop{} -> GetEnvCardStatus(from. gRPC)
		go envc.CardUpdate(cardNoStop)
	})

	// 4.4 -> (아직 미필요)

	// 5.1 -> (아직 미필요)

	// 5.2 -> (아직 미필요)

	// 5.3 -> 이벤트 기반 (충전이 종료 된 경우)

	// 5.4 -> (아직 미필요)

	// 5.5 -> (아직 미필요)

	// 6.1 -> (아직 미필요) -> 충전내역을 저장하는 경우, 이벤트 기반 (이력을 환경부에 등록 요청 한 후),

	// 6.2 -> 이벤트 기반 (충전이 종료 된 경우)

	// 6.3 -> (아직 미필요)

	// 6.4 -> (아직 미필요)

	cr.Start()

	/*
		이벤트 정리: gRPC 을 통해 call 받은 후 처리

		3.2 (chargePoint의 StatusNotification) : o
		5.3 (충전이 종료 된 경우) : o
		6.2 (충전이 종료 된 경우) : o

		// 6.1 (이력을 환경부에 등록 요청 한 후) :
	*/

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed To Listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAdminServiceServer(s, &rpc.GrpcServer{})
	log.Printf("Starting gRPC Listener On Port " + port)

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed To Serve: %v", err)
	}
}
