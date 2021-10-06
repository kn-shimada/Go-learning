package main

func main() {
	const a = 1 //型省略が可能、大概は指定しない
	const (
		b = 2
		c = 3
	)

	println("a =", a, "\nb = ", b, "\nc =", c)
}
