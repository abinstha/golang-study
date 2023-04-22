package main

import (
	"fmt"
)

//

// Unbuffered Channel
// - channel that needs a receiver as soon as message is emitted to the channel

// Buffered Channel
// - Give the size of the buffer
// - channel will block by default if its full
// - if we donot specify the size then it will only hold one

func main() {
	userch := make(chan string, 2)

	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	userch <- "abin" // blocking
	// }()

	userch <- "abin" // ["abin", ""]
	userch <- "stha" // ["abin", "stha"]
	userch <- "abin" // > is waiting for a consume of the channel

	user := <-userch

	fmt.Println(user)
}
