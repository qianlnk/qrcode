package qrcode

import (
	"testing"
)

func TestQRCode(t *testing.T) {
	qr := NewQRCode("https://github.com/qianlnk/qrcode", true)
	qr.Debug()
	qr.Output()
}
