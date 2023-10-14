// Modified from here: https://go.dev/tour/concurrency/3
// https://stackoverflow.com/questions/22747711/what-is-the-use-of-buffered-channels-in-go
// https://gobyexample.com/channel-buffering
// https://www.sohamkamani.com/golang/channels/
// https://go.dev/tour/concurrency/5

/*

The official tour of go (link above) uses this example:
---
package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
---

It encourages users to add another value to the channel to see what happens:

---
package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
    ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
---

This results in the following error:

---
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
	/tmp/sandbox3343513707/prog.go:9 +0x58

Program exited.
---

I understand how to do this, but why to do this is a little less clear, so I asked
ChatGPT:

---

Buffering a channel in Go is a common practice that can be useful in various situations
to manage the communication between goroutines. Channels in Go are used for communication
and synchronization between concurrent goroutines. When you create a channel without
specifying a buffer size (unbuffered channel), it operates in a blocking mode, meaning that
the sender will block until there is a receiver ready to receive the data, and vice versa.

Buffering a channel means that you specify a size for the channel when you create it.
This size determines the number of values that can be sent into the channel without
blocking. There are several reasons why you might want to use a buffered channel in Go:

Asynchronous Communication: Buffered channels allow for asynchronous communication between
goroutines. The sender can send data to a buffered channel and continue its work without
waiting for the receiver to be ready.

Reducing Blocking: Unbuffered channels can cause goroutines to block if there's no
immediate receiver. Buffered channels can help avoid this blocking and improve overall
program efficiency.

Rate Limiting: You can use buffered channels to implement rate limiting by limiting the
number of items that can be sent through the channel at a time. This can help control the
rate of execution of concurrent tasks.

Decoupling Producers and Consumers: Buffered channels can decouple producers and consumers.
Producers can keep producing data and put it in the channel, while consumers can process
the data as they are ready, without both being tightly synchronized.

Handling Bursts: Buffered channels can handle bursts of data. If there is a temporary surge
in data production, buffered channels can absorb some of the load and prevent the system
from becoming overwhelmed.

Improved Responsiveness: When you need to ensure that certain events do not block your main
application flow (e.g., handling user requests), using buffered channels can improve the
responsiveness of your application.

However, it's important to choose an appropriate buffer size. If the buffer is too small,
you might still experience blocking, and if it's too large, you could use more memory than
necessary. The buffer size should be chosen based on the specific requirements and
characteristics of your application.

In summary, buffering a channel in Go can be a useful tool for managing communication
between goroutines and optimizing the flow of data in concurrent programs, but it should
be used thoughtfully and tailored to the needs of your application.
---

My take after reading this:

Essentially, instead of defaulting to blocking mode, where a channel will block until
there is a receiver ready to receive the data, you're saying "I'm going to send you extra
stuff up to a certain point". The certain point is the buffer size. I think of it as the
lobby of a resturant. You can tolerate a certain number of people in the lobby even if you
don't have a table for them. After a certain number of people, you have to start turning
people away. The buffer size is the number of people you can tolerate in the lobby.

*/

package main

import (
    "fmt"
)

// Limits can be set on the number of values that can be sent to a channel
// before blocking.
func main() {

    // This channel can hold up to 2 values before blocking
    c := make(chan int, 2)

    // add two values to the channel
    c <- 1
    c <- 2

    // select is used to block and wait for one of several operations to finish.
    // Because Go doesn't use traditional try/catch blocks, this is how you
    // handle errors with adding to a channel. ChatGPT gave me an example of
    // how to handle errors in this context, because usually you store an error
    // in a variable and then check if the variable is nil.
    select {
        // try to add a third value to the channel
        case c <- 3:
            fmt.Println("Added 3 to the channel succussfully")
        // if the channel is full, this will run
        default:
            fmt.Println("Can't add anymore, possibly because the buffer is full")
    }
    // NOTE: if you delete the c <- 1 or c <- 2 above, then the c <- 3 will work
    // because there is room in the buffer

}


/*

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)

	go func() {
		// Attempt to add data to the buffered channel
		select {
		case ch <- 42:
			// Successful addition
		default:
			// Failed to add due to a full channel
			fmt.Println("Error: Channel is full")
		}
	}()

	time.Sleep(1 * time.Second) // Wait for the goroutine to finish

	// Read data from the channel
	select {
	case data := <-ch:
		// Successfully received data
		fmt.Println("Received data:", data)
	default:
		// No data received, you can handle this as an error
		fmt.Println("Error: No data received")
	}
}

*/
