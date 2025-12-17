package server

import (
	"fmt"
)

const (
	colorReset  = "\033[0m"
	colorCyan   = "\033[36m"
	colorPurple = "\033[35m"
	colorGreen  = "\033[32m"
)

func PrintBanner(version string) {
	fmt.Println(colorCyan + `
███████╗ ██████╗██╗  ██╗ ██████╗ 
██╔════╝██╔════╝██║  ██║██╔═══██╗
█████╗  ██║     ███████║██║   ██║
██╔══╝  ██║     ██╔══██║██║   ██║
███████╗╚██████╗██║  ██║╚██████╔╝
╚══════╝ ╚═════╝╚═╝  ╚═╝ ╚═════╝ 
` + colorReset)

	fmt.Println(colorPurple + "EchoDB — In-Memory Data Store" + colorReset)
	fmt.Printf(colorGreen+"Version: %s\n"+colorReset, version)
	fmt.Println(colorGreen + "Status : Starting up..." + colorReset)
	fmt.Println()
}
