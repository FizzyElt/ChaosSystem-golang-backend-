package chaos

import (
	"math"
)

const count int = 3000


//座標
type Coordinate struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

//座標列表 3000筆
type List struct {
	list []*Coordinate
}

func NewList() *List {
	return &List{
		list: nil,
	}
}
func (l *List) TinkerBell() bool {
	xn := -0.72
	yn := -0.64
	var nextX float64 = 0
	var nextY float64 = 0
	for i := 0; i < count; i++ {
		coordinate := &Coordinate{
			X: math.Trunc(xn*1000) / 1000,
			Y: math.Trunc(yn*1000) / 1000,
		}
		l.list = append(l.list, coordinate)
		nextX = (xn * xn) - (yn * yn) + (0.9 * xn) + (-0.6013 * yn)
		nextY = (2 * xn * yn) + (2 * xn) + (0.5 * yn)
		xn = nextX
		yn = nextY
	}
	return true
}
func (l *List) Duffing() bool {
	xn := -0.1
	yn := -0.1
	var nextX float64 = 0
	var nextY float64 = 0
	for i := 0; i < count; i++ {
		coordinate := &Coordinate{
			X: math.Trunc(xn*1000) / 1000,
			Y: math.Trunc(yn*1000) / 1000,
		}
		l.list = append(l.list, coordinate)
		nextX = yn
		nextY = -(0.2 * xn) + (2.75 * yn) - (yn * yn * yn)
		xn = nextX
		yn = nextY
	}
	return true
}
func (l *List) Henon() bool {
	xn := -0.72
	yn := -0.64
	var nextX float64 = 0
	var nextY float64 = 0
	for i := 0; i < count; i++ {
		coordinate := &Coordinate{
			X: math.Trunc(xn*1000) / 1000,
			Y: math.Trunc(yn*1000) / 1000,
		}
		l.list = append(l.list, coordinate)
		nextX = 1 - (1.4 * xn * xn) + yn
		nextY = 0.3 * xn
		xn = nextX
		yn = nextY
	}
	return true
}


func (l *List) GetData() []*Coordinate {
	return l.list
}
