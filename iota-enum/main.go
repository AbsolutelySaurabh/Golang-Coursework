package main

func main(){

	const(
		x = iota  // x == 0
		y = iota  // y == 1
		z = iota  // z == 2
		w  // If there is no expression after the constants name, it uses the last expression,
		//so it's saying w = iota implicitly. Therefore w == 3, and y and z both can omit "= iota" as well.
	)

	const v = iota // once iota meets keyword `const`, it resets to `0`, so v = 0.

	const (
		e, f, g = iota, iota, iota // e=0,f=0,g=0 values of iota are same in one line.
	)
}