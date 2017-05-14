# QRCode

a tool to generate qrcode and print it on console.

## Install

```shell
go get -u github.com/qianlnk/qrcode/..
```

## Usage

* [ ] cmd

```shell
qrcode 'https://github.com/qianlnk/qrcode'
```

* [ ] package

```golang
package main

import (
	"github.com/qianlnk/qrcode"
)

func main() {
	qr := qrcode.NewQRCode("https://github.com/qianlnk/qrcode", false)
	qr.Output()
}

```

* [ ] result

![](qrcode.gif)

