package manduco

import (
	"fmt"
	"time"
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=Optimizer
type Optimizer int

const (
	Ect Optimizer = iota
	Oxipng
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=Level
type Level int

const (
	Medium Level = iota
	Max
)

type OptimageReq struct {
	Image     []byte
	Optimizer Optimizer
	Level     Level
}

func (r *OptimageReq) String() string {
	return fmt.Sprintf("Image: %d bytes, Optimizer: %s, Level: %s", len(r.Image), r.Optimizer, r.Level)
}

type OptimageRes struct {
	OptimizedImage []byte
	TimeTaken      time.Duration
}
