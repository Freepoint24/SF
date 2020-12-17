package example

type geometry interface {
	area() float64
}
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

//структура rect  реализует интерфейс geometry
func (r rect) area() float64 {
	width := 6

	return r.width * r.height

}

// структура circle реализует только то, что сам хочет.
func (c circle) diameter() float64 {
	return c.radius * 2
}

func main() {

	height := 5

}
