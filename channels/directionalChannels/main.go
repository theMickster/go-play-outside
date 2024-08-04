package main

func main() {
	myChannel := make(chan string, 1)

	hello(myChannel, "Hello world from a directional channel")

}

func hello(sendOnlyChannel chan<- string, message string) {
	sendOnlyChannel <- message
}

func receive(receiveOnlyChannel <-chan string) {

}
