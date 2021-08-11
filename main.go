package main

import "example.com/go-demo-1/schedule"

func main() {
	schedule := schedule.Schedule{Name: "test.pdf", Url: "https://tcpdf.org/files/examples/example_012.pdf"}
	schedule.Process()
}
