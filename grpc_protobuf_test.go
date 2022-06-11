package benchmarks

import (
	"testing"
	"time"

	grpcprotobuf "benchmark-grpc-protobuf-vs-http-json/grpc-protobuf"

	"benchmark-grpc-protobuf-vs-http-json/grpc-protobuf/generated/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func init() {
	go grpcprotobuf.Start()
	time.Sleep(time.Second)
}

func BenchmarkGRPCProtobuf(b *testing.B) {
	conn, err := grpc.Dial("127.0.0.1:60000", grpc.WithInsecure())
	if err != nil {
		b.Fatalf("grpc connection failed: %v", err)
	}
	client := proto.NewApiClient(conn)

	b.ResetTimer()
	b.StartTimer()

	defer conn.Close()

	for n := 0; n < b.N; n++ {
		doGRPC(client, b)
	}
}

func doGRPC(client proto.ApiClient, b *testing.B) {
	resp, err := client.CreateUser(context.Background(), &proto.User{
		Email:    "foo@bar.com",
		Name:     "Bench",
		Password: "bench",
		Other:    "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
		Field1:   1,
		Field2:   2.3,
		Field3:   []string{"a", "b", "c"},
		Field4:   []int64{0, 1, 2, 3},
		Field5:   []float32{0, 1, 2, 3, 4},
	})

	if err != nil {
		b.Fatalf("grpc request failed: %v", err)
	}

	if resp == nil || resp.Code != 200 || resp.User == nil || resp.User.Id != "1000000" {
		b.Fatalf("grpc response is wrong: %v", resp)
	}
}
