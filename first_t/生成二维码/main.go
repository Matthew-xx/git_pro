// 生成二维码 project main.go
package main

import (
	"image/color"
	"log"

	"github.com/skip2/go-qrcode"
)

func main() {
	qr, err := qrcode.New("http://c.biancheng.net/", qrcode.Medium)
	if err != nil {
		log.Fatal(err)
	} else {
		qr.BackgroundColor = color.RGBA{50, 205, 50, 255}
		qr.ForegroundColor = color.White
		qr.WriteFile(256, "./golang_qrcode.png")
	}
}
