package grpc

import (
	envc "CSMS_ENV/env"
	envType "CSMS_ENV/env/type"
	pb "CSMS_ENV/grpc/adminService"
	"context"
	"io"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *GrpcServer) StatusNotification(ctx context.Context, in *pb.StatusNotificationInfo) (*pb.Result, error) {
	log.Printf("--- gRPC Input [StatusNotification] ---\n%v", in)

	cstat := envType.ChargerStatus{
		Sid:      in.Sid,
		Cid:      in.Cid,
		Status:   in.Status,
		Statdt:   in.Statdt,
		LastTsdt: in.LastTsdt,
		LastTedt: in.LastTedt,
		NowTsdt:  in.NowTsdt,
	}
	go envc.ChargerStatusUpdate([]envType.ChargerStatus{cstat})

	return &pb.Result{Result: true, ErrStr: ""}, status.New(codes.OK, "").Err()
}

func (r *GrpcServer) StatusNotificationStream(stream pb.AdminService_StatusNotificationStreamServer) error {
	log.Printf("--- gRPC Input [StatusNotificationStream] ---\n")

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			log.Println("Client Stream Is Done")
			return nil
		}
		if err != nil {
			log.Printf("Client Stream Error: %v", err)
			return err
		}
		log.Printf("--- gRPC Input Stream Data [StatusNotificationStream] ---\n%v", in)

		cstat := envType.ChargerStatus{
			Sid:      in.Sid,
			Cid:      in.Cid,
			Status:   in.Status,
			Statdt:   in.Statdt,
			LastTsdt: in.LastTsdt,
			LastTedt: in.LastTedt,
			NowTsdt:  in.NowTsdt,
		}
		go envc.ChargerStatusUpdate([]envType.ChargerStatus{cstat})

		result := pb.Result{Result: true, ErrStr: ""}
		err = stream.Send(&result)
		if err != nil {
			log.Printf("Server Stream Send Error: %v", err)
			return err
		}
	}
}

func (r *GrpcServer) StopTransaction(ctx context.Context, in *pb.StopTransactionInfo) (*pb.Result, error) {
	log.Printf("--- gRPC Input [StopTransaction] ---\n%v", in)

	cstr := envType.TradeRegister{
		No:      in.No,
		Sid:     in.Sid,
		Cid:     in.Cid,
		Tbid:    in.Tbid,
		Tsdt:    in.Tsdt,
		Tedt:    in.Tedt,
		Btid:    in.Btid,
		Pow:     in.Pow,
		Mon:     in.Mon,
		Bprice:  in.Bprice,
		Tbprice: in.Tbprice,
		Bmon:    in.Bmon,
	}
	csur := envType.UseRegister{
		Sid:     in.Sid,
		Cid:     in.Cid,
		Tbid:    in.Tbid,
		Tsdt:    in.Tsdt,
		Tedt:    in.Tedt,
		Pow:     in.Pow,
		Mon:     in.Mon,
		Rcvdate: in.Rcvdate,
	}
	go envc.TradeRegister([]envType.TradeRegister{cstr})
	go envc.UseRegister([]envType.UseRegister{csur})

	return &pb.Result{Result: true, ErrStr: ""}, status.New(codes.OK, "").Err()
}
