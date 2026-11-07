package main

import (
	"context"
	"log"
	"time"

	pb "grpc-echo-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// اتصال به سرور روی پورت 50051 (همون آدرسی که سرور روش گوش میده)
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("خطا در اتصال به سرور: %v", err)
	}
	defer conn.Close()

	// ساخت یه کلاینت برای سرویس Echo با استفاده از اتصالی که بالا ساختیم
	client := pb.NewEchoServiceClient(conn)

	// یه context با timeout می‌سازیم (یعنی اگه تا 5 ثانیه جواب نگرفتیم، لغو کن)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// فراخوانی متد Echo روی سرور، با یه پیام نمونه
	response, err := client.Echo(ctx, &pb.EchoRequest{Message: "سلام سرور!"})
	if err != nil {
		log.Fatalf("خطا در فراخوانی Echo: %v", err)
	}

	log.Printf("پاسخ دریافت شد: %s", response.GetMessage())
}
