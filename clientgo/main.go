package main

func main() {
	a := App{}
	a.Initialize("guest", "guest", "test")
	a.Run(":3000")
}
