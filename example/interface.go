package example

import (
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type ract struct {
	width, height float64
}

type circle struct {
	r float64
}

func (r *ract) area() float64 {
	return r.width * r.height
}

func (r *ract) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c *circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *circle) perim() float64 {
	return 2 * math.Pi * c.r
}

func CacalAreaPerim() {
	var g geometry
	g = &ract{width: 2, height: 3}
	g.area()
	g.perim()

	g = &circle{r: 2}
	g.area()
	g.perim()
}
