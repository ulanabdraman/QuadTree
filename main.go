package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
)

type pos struct {
	X, Y float64
}

type Quad struct {
	points   []pos
	Quads    []Quad
	Position pos
	H        float64
	W        float64
}

var stop int

func build(MainQuad *Quad) {
	for i := 1; i <= 4; i++ {
		var SubQuad Quad
		SubQuad.H = MainQuad.H / 2
		SubQuad.W = MainQuad.W / 2
		switch i {
		case 1:
			SubQuad.Position = pos{(MainQuad.Position.X*2 + MainQuad.W) / 2, (MainQuad.Position.Y*2 + MainQuad.H) / 2}
			//log.Println(SubQuad.Position)
			//log.Println(SubQuad.H, SubQuad.W)
			MainQuad.Quads = append(MainQuad.Quads, SubQuad)
		case 2:
			SubQuad.Position = pos{(MainQuad.Position.X*2 - MainQuad.W) / 2, (MainQuad.Position.Y*2 + MainQuad.H) / 2}
			//log.Println(SubQuad.Position)
			MainQuad.Quads = append(MainQuad.Quads, SubQuad)
		case 3:
			SubQuad.Position = pos{(MainQuad.Position.X*2 - MainQuad.W) / 2, (MainQuad.Position.Y*2 - MainQuad.H) / 2}
			//log.Println(SubQuad.Position)
			MainQuad.Quads = append(MainQuad.Quads, SubQuad)
		case 4:
			SubQuad.Position = pos{(MainQuad.Position.X*2 + MainQuad.W) / 2, (MainQuad.Position.Y*2 - MainQuad.H) / 2}
			//log.Println(SubQuad.Position)
			MainQuad.Quads = append(MainQuad.Quads, SubQuad)
		}
	}
	for _, point := range MainQuad.points {
		if point.X > MainQuad.Position.X && point.Y > MainQuad.Position.Y {
			MainQuad.Quads[0].points = append(MainQuad.Quads[0].points, point)
			//log.Println(point, " into 0")
		}
		if point.X < MainQuad.Position.X && point.Y > MainQuad.Position.Y {
			MainQuad.Quads[1].points = append(MainQuad.Quads[1].points, point)
			//log.Println(point, " into 1")
		}
		if point.X < MainQuad.Position.X && point.Y < MainQuad.Position.Y {
			MainQuad.Quads[2].points = append(MainQuad.Quads[2].points, point)
			//log.Println(point, " into 2")
		}
		if point.X > MainQuad.Position.X && point.Y < MainQuad.Position.Y {
			MainQuad.Quads[3].points = append(MainQuad.Quads[3].points, point)
			//log.Println(point, " into 3")
		}
	}

	for i, _ := range MainQuad.Quads {
		log.Println(MainQuad.Quads[i], i)
		if len(MainQuad.Quads[i].points) > 1 {
			build(&MainQuad.Quads[i])
		}

	}
}

func findNear(car pos, MainQuad Quad) []pos {
	var found []pos
	if len(MainQuad.points) == 1 {
		//fmt.Println("MainQuad poins len is 1")
		return MainQuad.points
	}
	if car.X > MainQuad.Position.X && car.Y > MainQuad.Position.Y {
		if len(MainQuad.Quads[0].points) == 0 {
			//fmt.Println("MainQuad poins len is 0")
			return MainQuad.points
		}
		found = findNear(car, MainQuad.Quads[0])
	}
	if car.X < MainQuad.Position.X && car.Y > MainQuad.Position.Y {
		if len(MainQuad.Quads[1].points) == 0 {
			//fmt.Println("MainQuad poins len is 0")
			return MainQuad.points
		}
		found = findNear(car, MainQuad.Quads[1])
	}
	if car.X < MainQuad.Position.X && car.Y < MainQuad.Position.Y {
		if len(MainQuad.Quads[2].points) == 0 {
			//fmt.Println("MainQuad poins len is 0")
			return MainQuad.points
		}
		found = findNear(car, MainQuad.Quads[2])
	}
	if car.X > MainQuad.Position.X && car.Y < MainQuad.Position.Y {
		if len(MainQuad.Quads[3].points) == 0 {
			//fmt.Println("MainQuad poins len is 0")
			return MainQuad.points
		}
		found = findNear(car, MainQuad.Quads[3])
	}
	return found
}

func showQuad(MainQuad Quad, o int) {
	for i := 0; i < o; i++ {
		fmt.Printf(" ")
	}
	fmt.Println(len(MainQuad.points), len(MainQuad.Quads), MainQuad.Position)
	for i, _ := range MainQuad.Quads {
		showQuad(MainQuad.Quads[i], o+1)
	}
}

func main() {
	var MainQuad Quad
	MainQuad.Position = pos{0, 0}
	MainQuad.H = 50
	MainQuad.W = 50
	for i := 0; i < 20; i++ {
		var point pos
		point.X = 50 - rand.Float64()*100
		point.Y = 50 - rand.Float64()*100
		MainQuad.points = append(MainQuad.points, point)
	}
	build(&MainQuad)
	showQuad(MainQuad, 0)
	car := pos{50 - rand.Float64()*100, 50 - rand.Float64()*100}
	fmt.Println(car)
	Poi := findNear(car, MainQuad)
	//fmt.Println(Poi)
	var minDis float64
	minDis = math.Inf(1)
	var point pos
	//var BrutePoint pos
	for i, _ := range Poi {
		dis := math.Sqrt(math.Pow(Poi[i].X-car.X, 2) + math.Pow(Poi[i].Y-car.Y, 2))
		if minDis > dis {
			minDis = dis
			point = Poi[i]
		}
	}
	fmt.Println(point)
	for i, _ := range MainQuad.points {
		dis := math.Sqrt(math.Pow(MainQuad.points[i].X-car.X, 2) + math.Pow(MainQuad.points[i].Y-car.Y, 2))
		if dis < minDis {
			fmt.Println("Fail")
		}
	}
}
