package models

import (
	"fmt"

	"gocv.io/x/gocv"
)

type Camera struct {
	Name     string
	Port     string
	Ip       string
	Protocol string
	FileName string
	Capturer *gocv.VideoCapture
}

func (c *Camera) MJPEGCapture() error {
	img := gocv.NewMat()
	defer img.Close()
	window := gocv.NewWindow(c.Ip + ":" + c.Port)
	for {
		if ok := c.Capturer.Read(&img); !ok {
			return fmt.Errorf("Can't read camera")
		}
		window.IMShow(img)
		window.WaitKey(2)
		//buf, _ := gocv.IMEncode(".jpg", img)
		//stream.UpdateJPEG(buf.GetBytes())
		//buf.Close()
	}
}
