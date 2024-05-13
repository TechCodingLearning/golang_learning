package imageSyntax

import (
	"fmt"
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

/*
图像
image 包定义了 Image 接口：

package image

type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}
注意: Bounds 方法的返回值 Rectangle 实际上是一个 image.Rectangle，它在 image 包中声明。

（请参阅文档了解全部信息。）

color.Color 和 color.Model 类型也是接口，但是通常因为直接使用预定义的实现 image.RGBA 和 image.RGBAModel 而被忽视了。这些接口和类型由 image/color 包定义。
*/

type Image struct{}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{R: uint8(0), G: uint8(0), B: uint8(255), A: uint8(255)}
}

func TestImage() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())

	m1 := Image{}
	pic.ShowImage(m1)
}
