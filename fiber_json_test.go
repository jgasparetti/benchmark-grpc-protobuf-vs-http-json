package benchmarks

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	fiberjson "benchmark-grpc-protobuf-vs-http-json/fiber-json"
)

func init() {
	go fiberjson.Start()
	time.Sleep(time.Second)
}

func BenchmarkFiberJSON(b *testing.B) {
	client := &http.Client{}

	b.ResetTimer()
	b.StartTimer()

	for n := 0; n < b.N; n++ {
		doPostFiber(client, b)
	}
}

func doPostFiber(client *http.Client, b *testing.B) {
	u := &fiberjson.User{
		Email:    "foo@bar.com",
		Name:     "Bench",
		Password: "bench",
		Other:    "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.", Field1: 1,
		Field2: 2.3,
		Field3: []string{"a", "b", "c"},
		Field4: []int64{0, 1, 2, 3},
		Field5: []float32{0, 1, 2, 3, 4},
	}
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(u)
	if err != nil {
		b.Fatalf("encoding error: %v", err)
	}

	resp, err := client.Post("http://127.0.0.1:60002/", "application/json", buf)
	if err != nil {
		b.Fatalf("http request failed: %v", err)
	}

	defer resp.Body.Close()

	// We need to parse response to have a fair comparison as gRPC does it
	var target fiberjson.Response
	decodeErr := json.NewDecoder(resp.Body).Decode(&target)
	if decodeErr != nil {
		b.Fatalf("unable to decode json: %v", decodeErr)
	}

	if target.Code != 200 || target.User == nil || target.User.ID != "1000000" {
		b.Fatalf("http response is wrong: %v", resp)
	}
}
