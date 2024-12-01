package main

import (
	"bytes"
    "fmt"
	"image"
	_ "image/jpeg" // For JPEG decoding
    _ "image/png"  // For PNG decoding
    _ "golang.org/x/image/bmp" // For BMP decoding

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"

	"qr-code-cli/clipboard"
	"qr-code-cli/xclip"
)

func main() {

	fmt.Println("Reading from clipboard:")

	clipboardDataTypes := clipboard.GetClipboardDataTypes()
	imgBytes, err := xclip.PasteFromClipboard(clipboardDataTypes.ImageBMP)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

    // Decode the image from bytes
    img, _, err := image.Decode(bytes.NewReader(imgBytes))
    if err != nil {
        fmt.Println("Error decoding image:", err)
        return
    }

    // Create a binary bitmap from the decoded image
    imgBm, err := gozxing.NewBinaryBitmapFromImage(img)
    if err != nil {
        fmt.Println("Error creating binary bitmap:", err)
        return
    }

    // Initialize the QR code reader
    reader := qrcode.NewQRCodeReader()

    // Decode the QR code
    result, err := reader.Decode(imgBm, nil)
    if err != nil {
        fmt.Println("Error decoding QR code:", err)
        return
    }

    // Print the content of the QR code
    fmt.Println("QR Code Content:", result.GetText())
}
