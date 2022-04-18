package main

func main() {
	r := SetupRouter()
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
