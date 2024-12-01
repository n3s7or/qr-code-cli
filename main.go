package main

import (
    "fmt"

	"qr-code-cli/clipboard"
	"qr-code-cli/xclip"
)

func main() {

	fmt.Println("Reading from clipboard:")

	clipboardDataTypes := clipboard.GetClipboardDataTypes()
	text, err := xclip.PasteFromClipboard(clipboardDataTypes.ImageJPEG)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Clipboard content: ", text)

    // // Read image data from the clipboard as a string
    // // imgData, err := clipboard.ReadAll()
    // if err != nil {
    //     fmt.Println("Error reading from clipboard:", err)
    //     return
    // }

    // if imgData == "" {
    //     fmt.Println("Clipboard is empty.")
    //     return
    // } else {
	// 	fmt.Println("Clipboard: .", imgData)
	// }

    // // Convert the image data string to bytes
    // imgBytes := []byte(imgData)

    // // Decode the image from bytes
    // img, _, err := image.Decode(bytes.NewReader(imgBytes))
    // if err != nil {
    //     fmt.Println("Error decoding image:", err)
    //     return
    // }

    // // Create a binary bitmap from the decoded image
    // imgBm, err := gozxing.NewBinaryBitmapFromImage(img)
    // if err != nil {
    //     fmt.Println("Error creating binary bitmap:", err)
    //     return
    // }

    // // Initialize the QR code reader
    // reader := qrcode.NewQRCodeReader()

    // // Decode the QR code
    // result, err := reader.Decode(imgBm, nil)
    // if err != nil {
    //     fmt.Println("Error decoding QR code:", err)
    //     return
    // }

    // // Print the content of the QR code
    // fmt.Println("QR Code Content:", result.GetText())
}
