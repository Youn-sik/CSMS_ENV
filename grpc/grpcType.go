package grpc

import (
	pb "CSMS_ENV/grpc/adminService"
)

type GrpcServer struct {
	pb.UnimplementedAdminServiceServer
}
