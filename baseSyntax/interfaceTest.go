package baseSyntax

import "fmt"

type Shape interface {
	Area() float64
	GetInfo() Info
}

type Info interface {
	Description() string
}

type Rectangle struct {
	Width  float64
	Height float64
}

type RectangleInfo struct {
	DescriptionText string
	Msg             string
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) GetInfo() Info {
	return RectangleInfo{DescriptionText: "This is a rectangle", Msg: "message"}
}

func (ri RectangleInfo) Description() string {
	return ri.DescriptionText
}

type Caller interface {
	Call() (Info, error)
}

//
//type PreOrderCaller struct {
//}
//
//func (caller *PreOrderCaller) Call() (Info, error) {
//	fmt.Println("aaaa")
//	return RectangleInfo{DescriptionText: "xxx"}, nil
//}
//
//type ServerProxy struct {
//	preOrderCaller Caller
//}
//
//func NewServerProxy() ServerProxy {
//	return ServerProxy{
//		preOrderCaller: &PreOrderCaller{},
//	}
//}

func TestInterface() {
	rectangle := Rectangle{Width: 10, Height: 5}
	shape := rectangle

	fmt.Println(shape.Area())

	info := shape.GetInfo()
	rectInfo, _ := info.(RectangleInfo)
	fmt.Println(rectInfo.Description())
	fmt.Println(rectInfo.Msg)
}
