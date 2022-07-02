package task

import (
	"bytes"
	"fmt"
	"image/png"
	"os"

	"github.com/nfnt/resize"
)

func img2txt(imgPath string, size uint, txts []string, rowend string, output string) {
	file, err := os.Open(imgPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var width = size
	var height = (size * uint(img.Bounds().Dy())) / (uint(img.Bounds().Dx()))
	height = height * 6 / 10
	newimg := resize.Resize(width, height, img, resize.Lanczos3)
	dx := newimg.Bounds().Dx()
	dy := newimg.Bounds().Dy()

	textBuffer := bytes.Buffer{}

	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			colorRgb := newimg.At(x, y)
			r, g, b, _ := colorRgb.RGBA()

			avg := uint8((r + g + b) / 3 >> 8)
			num := avg / uint8(256/len(txts))
			textBuffer.WriteString(txts[num])
			fmt.Print(txts[num])
		}

		textBuffer.WriteString(rowend)
		fmt.Print(rowend)
	}

	f, err := os.Create(output + ".txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()

	f.WriteString(textBuffer.String())
}
