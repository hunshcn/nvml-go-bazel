package main

import (
	"fmt"

	"github.com/NVIDIA/go-nvml/pkg/nvml"
)

func main() {
	fmt.Println(nvml.Init())
}
