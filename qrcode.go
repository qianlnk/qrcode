package qrcode

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	qrcode "github.com/skip2/go-qrcode"
)

const (
	QR_CODE_SIZE        = 256
	SHRINK_QR_CODE_SIZE = 35
	MARGIN              = 29
)

type QRCode struct {
	url             string
	img             image.Image
	genImg          bool
	points          [QR_CODE_SIZE][QR_CODE_SIZE]int
	tmpShrinkPoints [QR_CODE_SIZE][SHRINK_QR_CODE_SIZE]int
	shrinkPoints    [SHRINK_QR_CODE_SIZE][SHRINK_QR_CODE_SIZE]int
}

//NewQRCode 返回二维码结构
func NewQRCode(uri string, genImg bool) *QRCode {
	qr := &QRCode{
		url:    uri,
		genImg: genImg,
	}

	qr.genQRCode()
	qr.binarization()
	qr.shrink()

	return qr
}

//genQRCode 生成二维码
func (qr *QRCode) genQRCode() error {
	code, err := qrcode.Encode(qr.url, qrcode.Medium, QR_CODE_SIZE)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(code)
	img, err := png.Decode(buf)
	if err != nil {
		return err
	}

	qr.img = img

	if qr.genImg {
		newPng, _ := os.Create("qrcode.png")
		defer newPng.Close()
		png.Encode(newPng, img)
	}

	return nil
}

//binarization 二维码图片二值化 0－1
func (qr *QRCode) binarization() {
	gray := image.NewGray(image.Rect(0, 0, QR_CODE_SIZE, QR_CODE_SIZE))
	for x := 0; x < QR_CODE_SIZE; x++ {
		for y := 0; y < QR_CODE_SIZE; y++ {
			r32, g32, b32, _ := qr.img.At(x, y).RGBA()
			r, g, b := int(r32>>8), int(g32>>8), int(b32>>8)
			if (r+g+b)/3 > 180 {
				qr.points[y][x] = 0
				gray.Set(x, y, color.Gray{uint8(255)})
			} else {
				qr.points[y][x] = 1
				gray.Set(x, y, color.Gray{uint8(0)})
			}
		}
	}

	if qr.genImg {
		newPng, _ := os.Create("qrcode.binarization.png")
		defer newPng.Close()
		png.Encode(newPng, gray)
	}
}

//shrink 缩小二值化数组
func (qr *QRCode) shrink() {
	for x := 0; x < QR_CODE_SIZE; x++ {
		cal := 1
		for y := MARGIN + 1; y < QR_CODE_SIZE-MARGIN; y += 6 {
			qr.tmpShrinkPoints[x][cal] = qr.points[x][y]
			cal++
		}
	}

	for y := 1; y < SHRINK_QR_CODE_SIZE-1; y++ {
		row := 1
		for x := MARGIN + 1; x < QR_CODE_SIZE-MARGIN; x += 6 {
			qr.shrinkPoints[row][y] = qr.tmpShrinkPoints[x][y]
			row++
		}
	}
}

//Output 控制台输出二维码
func (qr *QRCode) Output() {
	for x := 0; x < SHRINK_QR_CODE_SIZE; x++ {
		for y := 0; y < SHRINK_QR_CODE_SIZE; y++ {
			if qr.shrinkPoints[x][y] == 1 {
				fmt.Print("\033[40;37m  \033[0m")
			} else {
				fmt.Print("\033[47;30m  \033[0m")
			}
		}
		fmt.Println()
	}
}

//Debug 调试二维码二值化及缩小过程
func (qr *QRCode) Debug() {
	src, _ := os.Create("src.txt")
	for i := 0; i < len(qr.points); i++ {
		var line string
		for j := 0; j < len(qr.points[i]); j++ {
			if qr.points[i][j] == 1 {
				line += "1"
			} else {
				line += "0"
			}
		}
		line += "\n"
		src.WriteString(line)
	}
	src.Close()

	tmp, _ := os.Create("tmp.txt")
	for i := 0; i < len(qr.tmpShrinkPoints); i++ {
		var line string
		for j := 0; j < len(qr.tmpShrinkPoints[i]); j++ {
			if qr.tmpShrinkPoints[i][j] == 1 {
				line += "1"
			} else {
				line += "0"
			}
		}
		line += "\n"
		tmp.WriteString(line)
	}
	tmp.Close()

	dst, _ := os.Create("dst.txt")
	for i := 0; i < len(qr.shrinkPoints); i++ {
		var line string
		for j := 0; j < len(qr.shrinkPoints[i]); j++ {
			if qr.shrinkPoints[i][j] == 1 {
				line += "1"
			} else {
				line += "0"
			}
		}
		line += "\n"
		dst.WriteString(line)
	}
	dst.Close()
}
