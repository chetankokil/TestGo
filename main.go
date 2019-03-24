package main

func main() {
	a := App{}
	a.Initialize("sa", "", "TEST_DB")
	a.Run(":8080")
}
