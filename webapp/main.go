package main

func main() {
	a := App{}
	a.Initialize("root", "admin123", "test")

	a.Run(":80")
}
