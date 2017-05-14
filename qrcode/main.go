package main

import (
	"fmt"
	"os"

	"github.com/qianlnk/qrcode"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage:\n qrcode <string>\n")
	}

	if len(os.Args[1]) > 60 {
		fmt.Println("\033[31mERR: max context length is 60.\033[0m")
		return
	}

	qr := qrcode.NewQRCode(os.Args[1], false)
	qr.Output()
}
