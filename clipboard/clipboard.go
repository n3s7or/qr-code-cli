package clipboard

// ClipboardDataTypes holds the types of clipboard data that your project will handle.
type ClipboardDataTypes struct {
    TextPlain string
    ImagePNG  string
    ImageJPEG string
    ImageBMP  string
}

// clipboardDataTypes is a single instance of ClipboardDataTypes that can be used throughout the package.
var clipboardDataTypes = ClipboardDataTypes{
    TextPlain: "text/plain",
    ImagePNG:  "image/png",
    ImageJPEG: "image/jpeg",
    ImageBMP:  "image/bmp",
}

// GetClipboardDataTypes returns the singleton instance of ClipboardDataTypes.
func GetClipboardDataTypes() ClipboardDataTypes {
    return clipboardDataTypes
}
