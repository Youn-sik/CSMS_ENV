package main

import (
	"log"
	"net"

	"github.com/robfig/cron"
	"google.golang.org/grpc"
)

const (
	port = ":9080"
)

func main() {
	cr := cron.New()
	// 2.1 -> 1시간 마다
	// 2.2 -> 최초 1회
	// 2.3 -> 10분 마다
	// 2.4 -> 최초 1회
	// 3.1 -> 이벤트 기반 (충전기를 환경부에 등록 요청 한 후)
	// 3.2 -> 10분 마다 (가장 최근 상태를 수신) + 이벤트 기반 (어드민 서비스에 statusNotification)
	// 4.1 -> 10분 마다
	// 4.2 -> 24시간 마다
	// 4.3 -> 24시간 마다 + 이벤트 기반 (사용자를 환경부에 등록 요청 한 후)
	// 4.4 -> (아직 미필요)
	// 5.1 -> (아직 미필요)
	// 5.2 -> (아직 미필요)
	// 5.3 -> 이벤트 기반 (충전이 종료 된 경우)
	// 5.4 -> (아직 미필요)
	// 5.5 -> (아직 미필요)
	// 6.1 -> (아직 미필요) -> 충전내역을 저장하는 경우, 이벤트 기반 (이력을 환경부에 등록 요청 한 후)
	// 6.2 -> 이벤트 기반 (충전이 종료 된 경우)
	// 6.3 -> (아직 미필요)
	// 6.4 -> (아직 미필요)
	cr.Start()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed To Listen: %v", err)
	}
	s := grpc.NewServer()
	// pb.RegisterProductInfoServer(s, &server{})
	log.Printf("Starting gRPC Listener On Port " + port)

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed To Serve: %v", err)
	}
}
