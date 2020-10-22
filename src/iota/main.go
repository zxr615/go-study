package main

func main() {
	const (
		x = iota
		y = iota
		z = iota
	)

	const i = iota
	const j = iota

	const (
		e, f, g = iota, iota, iota
	)


	println ("x = ", x)
	println("y = ", y)
	println("z = ", z)
	println("i = ", i)
	println("j = ", j)
	println("e = ", e)
	println("f = ", f)
	println("g = ", g)
}
