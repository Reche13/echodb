package server

import (
	"fmt"

	"github.com/reche13/echodb/internal/info"
)

const (
	colorReset  = "\033[0m"
	colorCyan   = "\033[36m"
	colorPurple = "\033[35m"
	colorGreen  = "\033[32m"
)

func PrintBanner() {
	fmt.Println(colorCyan + `
███████╗ ██████╗██╗  ██╗ ██████╗ 
██╔════╝██╔════╝██║  ██║██╔═══██╗
█████╗  ██║     ███████║██║   ██║
██╔══╝  ██║     ██╔══██║██║   ██║
███████╗╚██████╗██║  ██║╚██████╔╝
╚══════╝ ╚═════╝╚═╝  ╚═╝ ╚═════╝ 
` + colorReset)

	fmt.Println(colorPurple + "EchoDB — In-Memory Data Store" + colorReset)
	fmt.Printf(colorGreen+"Version: %s\n"+colorReset, info.Version)
	fmt.Println(colorGreen + "Status : Starting up..." + colorReset)
	fmt.Println()
}
