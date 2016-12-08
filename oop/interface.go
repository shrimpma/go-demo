package main

type Exchanger interface {
	Exchange()
}

type StringPaire struct {
	first, second string
}

func (pair *StringPaire) Exchang() {
	pair.first, pair.second = pair.second, pair.first
}

type Point [2]int

func (point *Point) Exchange() {
	point[0], point[1] = point[1], point[0]
}
