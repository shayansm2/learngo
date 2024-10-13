package main

func main() {
	port := 1234
	server := Server{port}
	server.Start() // add mutex or whatsoever, so I can call it after it is initiated

	//client := InitHttpClient(port)
	//client.SamplePost()
}
