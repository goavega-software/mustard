package core

import (
	"encoding/base64"
	"fmt"

	qrcode "github.com/skip2/go-qrcode"
)

/*
QRAsDataURI returns back the QRCode for a Url as base64 encoded string
*/
func QRAsDataURI(url string) string {
	png, _ := qrcode.Encode(url, qrcode.Medium, 75)
	return fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(png))
}
