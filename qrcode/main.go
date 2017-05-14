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

	qr := qrcode.NewQRCode(os.Args[1], false)
	qr.Output()
}
