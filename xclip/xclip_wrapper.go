package xclip

import (
    "bytes"
    "fmt"
    "os/exec"

    "qr-code-cli/clipboard"
)


// func CopyToClipboard(data string) error {
//     cmd := exec.Command("xclip", "-selection", "clipboard", "-t", "image/plain")
//     cmd.Stdin = bytes.NewBufferString(data)
//     err := cmd.Run()
//     if err != nil {
//         return fmt.Errorf("failed to copy to clipboard: %v", err)
//     }
//     return nil
// }


func PasteFromClipboard(target string) ([]byte, error) {
    dataTypes := clipboard.GetClipboardDataTypes()
    if target != dataTypes.TextPlain &&
       target != dataTypes.ImagePNG &&
       target != dataTypes.ImageJPEG &&
       target != dataTypes.ImageBMP {
        return []byte(""), fmt.Errorf("invalid clipboard data type: %s", target)
    }
    
    cmd := exec.Command("xclip", "-selection", "clipboard", "-t", target, "-o")
    var out bytes.Buffer
    var errOut bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = &errOut

    err := cmd.Run()
    if err != nil {
        return []byte(""), fmt.Errorf("failed to paste from clipboard: %v, error output: %s", err, errOut.String())
    }
    return out.Bytes(), nil
}


// func _GetTargets() (string, error) {
//     cmd := exec.Command("xclip", "-selection", "clipboard", "-t", "TARGETS", "-o")
//     var out bytes.Buffer
//     cmd.Stdout = &out
//     err := cmd.Run()
//     if err != nil {
//         return "", fmt.Errorf("failed to paste from clipboard: %v", err)
//     }
//     return out.String(), nil
// }

