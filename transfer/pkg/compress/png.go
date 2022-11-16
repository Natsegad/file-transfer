package compress

import (
	"bytes"
	"fmt"
	"github.com/yusukebe/go-pngquant"
	"image/png"
	"io"
	"os"
)

func CompressPng(path string) {

	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error compress img %s %v \n", path, err)
		return
	}
	defer file.Close()
	decode, err := png.Decode(file)
	if err != nil {
		fmt.Printf("Error Decode png img %s %v \n", path, err)
		return
	}

	pngComp, err := pngquant.Compress(decode, "1")
	if err != nil {
		fmt.Printf("Error pngquant.Compress png img %s %v \n", path, err)
		return
	}

	out := new(bytes.Buffer)
	err = png.Encode(out, pngComp)
	if err != nil {
		fmt.Printf("Error Encode png img %s %v \n", path, err)
		return
	}

	fileLp, err := os.Create(path + "_lp")
	if err != nil {
		fmt.Printf("Error Create %s %v \n", path, err)
		return
	}
	defer fileLp.Close()

	io.Copy(fileLp, out)
}
