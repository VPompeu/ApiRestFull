package main

func main() {
	a := App{}
	a.Initialize(
		"postgres",
		"3236",
		"postgres")

	a.Run(":8010")
}