package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"

	"github.com/lucas-deangelis/manduco"
)

const (
	port = 1234
	ip   = "localhost"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: client <image.png>")
	}

	since("args check")

	address := fmt.Sprintf("%s:%d", ip, port)
	fmt.Printf("Connecting to %s with tcp\n", address)
	client, err := rpc.Dial("tcp", address)
	if err != nil {
		log.Fatalf("failed to dial %s with tcp: %v", address, err)
	}
	defer client.Close()

	since("dial")

	imageData, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("failed to read file %s: %v", os.Args[1], err)
	}

	since("read file")

	beforeLen := len(imageData)

	req := &manduco.OptimageReq{
		Image:     imageData,
		Optimizer: manduco.Ect,
		Level:     manduco.Medium,
	}

	var res manduco.OptimageRes

	since("before call")

	err = client.Call("Serv.Optimize", req, &res)
	if err != nil {
		log.Fatal("optimize error:", err)
	}

	since("after call")
	fmt.Printf("Time taken: %s\n", res.TimeTaken)

	fmt.Printf("Before: %d bytes\n", beforeLen)
	fmt.Printf("After: %d bytes\n", len(res.OptimizedImage))
	fmt.Printf("Saved: %.2f%%\n", 100.0-float64(len(res.OptimizedImage))/float64(beforeLen)*100.0)

	newName := strings.Replace(os.Args[1], ".png", "-optimized.png", 1)

	since("before write file")

	err = os.WriteFile(newName, res.OptimizedImage, 0o644)
	if err != nil {
		fmt.Printf("failed to write file: %v", err)
		os.Exit(1)
	}

	since("after write file")
}
