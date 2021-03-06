package renderer

import (
	"errors"
	"image"
  "fmt"
  "bytes"

	cv "gocv.io/x/gocv"
)

func init() {
	RendererMap["ascii-256"] = ascii256
}

func ascii256(img cv.Mat, size image.Point) (string, error) {
  var buffer bytes.Buffer
	cv.Resize(img, &img, size, 0, 0, 1)

	imgPtr := img.DataPtrUint8()

	if img.Cols()*img.Rows()*3 != len(imgPtr) {
		return "", errors.New("Color RGB image is only supported")
	}

	for i := 0; i < img.Rows(); i += 2 {
		for j := 0; j < img.Cols()*3; j += 3 {
			fmt.Fprintf(&buffer, "\033[38;2;%d;%d;%dm█\033[39m",
				imgPtr[i*img.Cols()*3+j+2],
				imgPtr[i*img.Cols()*3+j+1],
				imgPtr[i*img.Cols()*3+j])
		}
		if i != img.Rows()-2 {
			fmt.Fprintf(&buffer, "\033[K\n")
		}
  }
	fmt.Fprintf(&buffer, "\033[J")
	return buffer.String(), nil
}
