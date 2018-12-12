package main

import (
	"image"
	"image/color"
	"os"
	"path/filepath"
	"strings"
	
	"gocv.io/x/gocv"
)

var (
	PathToImages = "/images/form"
)

func detect(path string, info os.FileInfo) {
	//window := gocv.NewWindow("detected circles")
	//defer window.Close()
	
	img := gocv.IMRead(path, gocv.IMReadGrayScale)
	defer img.Close()
	
	gocv.MedianBlur(img, &img, 5)

	cimg := gocv.NewMat()
	defer cimg.Close()

	gocv.CvtColor(img, &cimg, gocv.ColorGrayToBGR)

	circles := gocv.NewMat()
	defer circles.Close()

	gocv.HoughCirclesWithParams(
		img,
		&circles,
		gocv.HoughGradient,
		1, // dp
		float64(img.Rows()/8), // minDist
		75, // param1
		20, // param2
		10, // minRadius
		0,  // maxRadius
	)

	blue := color.RGBA{0, 0, 255, 0}
	red := color.RGBA{255, 0, 0, 0}

	for i := 0; i < circles.Cols(); i++ {
		v := circles.GetVecfAt(0, i)
		// if circles are found
		if len(v) > 2 {
			x := int(v[0])
			y := int(v[1])
			r := int(v[2])

			gocv.Circle(&cimg, image.Pt(x, y), r, blue, 2)
			gocv.Circle(&cimg, image.Pt(x, y), 2, red, 3)
		}
	}

	//for {
	//	window.IMShow(cimg)
	//
	//	if window.WaitKey(10) >= 0 {
	//		break
	//	}
	//}
	
	gocv.IMWrite(strings.Replace(path, info.Name(), strings.Join([]string{"out", info.Name()}, "/"), 1), cimg)
}

func main() {

	allowedExt := []string{".jpg", ".png", ".JPG", ".PNG"}
	
    err := filepath.Walk(PathToImages, func(path string, info os.FileInfo, err error) error {
    	if err != nil {
			return err
		}
    	if info.IsDir() {
			return nil
		}
		for _, e := range allowedExt {
			if filepath.Ext(path) == e {
				detect(path, info)
        		return nil
			}
		}
        return nil
    })
    if err != nil {
        panic(err)
    }
}