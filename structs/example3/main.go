package main

import "fmt"

type Box struct{
	length, width, height int
	color Color
}

type Color byte
type BoxList []Box

const(
	RED = iota
	GREEN 		// refreeing the prop of groups, it'll auto take the above declaration and referring
				// to iota property it'll be = 1 value as iota auto-increments
	YELLOW
)

func (b Box) volume()int{
	return b.length*b.width*b.height
}

func (c Color) String() string{
	colors := []string{"RED", "GREEN", "YELLOW"}
	return colors[c]
}

func main(){

	boxes := BoxList{
		Box{10, 11, 12, RED},
		Box{13, 14, 15, GREEN},
		Box{16, 17, 18, YELLOW},
	}

	fmt.Println("We have %d boxes in our set.", len(boxes))
	fmt.Println("The volume of the first one is: ",boxes[0].volume());
	fmt.Println("The color of the last box is: ", boxes[len(boxes)-1].color)

}