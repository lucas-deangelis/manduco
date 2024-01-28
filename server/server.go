package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
	"os/exec"
	"time"

	"github.com/lucas-deangelis/manduco"
)

const ip = "1234"

type Serv int

func (s *Serv) Optimize(req *manduco.OptimageReq, res *manduco.OptimageRes) error {
	since("before call")
	log.Printf("Received request: %s", req)

	tmpFile, err := os.CreateTemp(".", "img-*.png")
	if err != nil {
		return err
	}

	since("create temp file")

	_, err = tmpFile.Write(req.Image)
	if err != nil {
		return err
	}

	since("write temp file")

	err = tmpFile.Close()
	if err != nil {
		return err
	}

	since("close temp file")

	var cmd *exec.Cmd
	if req.Optimizer == manduco.Ect {
		if req.Level == manduco.Medium {
			log.Printf("ect medium %s", tmpFile.Name())
			cmd = exec.Command("ect", "--strict", "-3", tmpFile.Name())
		} else if req.Level == manduco.Max {
			cmd = exec.Command("ect", "--strict", "-9", tmpFile.Name())
		} else {
			return errors.New("invalid level")
		}
	} else if req.Optimizer == manduco.Oxipng {
		if req.Level == manduco.Medium {
			cmd = exec.Command("oxipng", "-o2", tmpFile.Name())
		} else if req.Level == manduco.Max {
			cmd = exec.Command("oxipng", "-o", "max", tmpFile.Name())
		} else {
			return errors.New("invalid level")
		}
	} else {
		return errors.New("invalid optimizer")
	}

	since("create command")

	timeOpti := time.Now()
	if err := cmd.Run(); err != nil {
		log.Printf("failed to optimize image: %v", err)
		return fmt.Errorf("failed to optimize image: %v", err)
	}
	res.TimeTaken = time.Since(timeOpti)

	since("run command")

	optimizedImage, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		return fmt.Errorf("failed to read optimized image: %v", err)
	}

	since("read optimized image")

	res.OptimizedImage = optimizedImage

	err = os.Remove(tmpFile.Name())
	if err != nil {
		return fmt.Errorf("failed to remove temporary file: %v", err)
	}

	since("remove temp file")

	return nil
}

func main() {
	serv := new(Serv)
	rpc.Register(serv)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", ip))
	if err != nil {
		log.Fatal("Listen error:", err)
	}

	log.Printf("Server listening on port %s\n", ip)

	rpc.Accept(listener)
}
