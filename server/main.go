package main

import (
	"context"
	"log"
	"net"

	pb "grpc-echo-service/proto"

	"google.golang.org/grpc"
)

// این استراکت، سرور ماست. باید متد Echo رو که توی proto تعریف کردیم پیاده‌سازی کنه
type server struct {
	pb.UnimplementedEchoServiceServer
}

// پیاده‌سازی متد Echo - همون چیزی که توی echo.proto تعریف کردیم
func (s *server) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Printf("پیام دریافت شد: %s", req.GetMessage())

	// یه پیام ساده برمی‌گردونیم که نشون بده سرور پاسخ داده
	response := &pb.EchoResponse{
		Message: "سرور گفت: " + req.GetMessage(),
	}

	return response, nil
}

func main() {
	// روی پورت 50051 گوش می‌کنیم (پورت رایج برای gRPC)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("خطا در گوش دادن به پورت: %v", err)
	}

	// ساخت یه سرور gRPC جدید
	grpcServer := grpc.NewServer()

	// ثبت سرویس Echo‌مون روی این سرور
	pb.RegisterEchoServiceServer(grpcServer, &server{})

	log.Println("سرور gRPC روی پورت 50051 در حال اجراست...")

	// شروع به سرویس‌دهی
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("خطا در اجرای سرور: %v", err)
	}
}
