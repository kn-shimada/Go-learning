package main

func main() {
	var a int

	var b int = 1

	a = b - 1

	var c = 2

	d := 3

	var (
		e int = 4
		f int = 5
	)

	name, age := "ken", 14
	name, age = "shimada", 15

	println("a =", a, "\nb =", b, "\nc =", c, "\nd =", d, "\ne =", e, "\nf =", f, "\nname =", name, "\nage = ", age)
}
