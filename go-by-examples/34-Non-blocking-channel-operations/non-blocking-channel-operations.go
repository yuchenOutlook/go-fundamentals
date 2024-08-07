// package main

// import "fmt"

// func main() {
// 	// whether the message has a buffer or not will affect the behavior of the select statement.
// 	// messages := make(chan string, 1)
// 	messages := make(chan string)
// 	signals := make(chan bool)

// 	select {
// 	case msg := <-messages:
// 		fmt.Println("received message", msg)
// 	default:
// 		fmt.Println("no message received")
// 	}

// 	/*
// 	A non-blocking send works similarly.
// 	Here msg cannot be sent to the messages channel,
// 	because the channel has no buffer and there is no receiver.
// 	Therefore the default case is selected.
// 	*/
// 	msg := "hi"
// 	select {
// 	case messages <- msg:
// 		fmt.Println("sent message", msg)
// 	default:
// 		fmt.Println("no message sent")
// 	}

// 	select {
// 	case msg := <-messages:
// 		fmt.Println("received message", msg)
// 	case sig := <-signals:
// 		fmt.Println("received signal", sig)
// 	default:
// 		fmt.Println("no activity")
// 	}

// }

/*
Version 2: If you want the message to be sent to the channel, you can make the channel buffered.
*/
package main

import "fmt"

func main() {
    messages := make(chan string, 1)  // Make a buffered channel
    signals := make(chan bool)

    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    default:
        fmt.Println("no message received")
    }

    msg := "sample message"
    select {
    case messages <- msg:
        fmt.Println("sent message", msg)
    default:
        fmt.Println("no message sent")
    }

    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    case sig := <-signals:
        fmt.Println("received signal", sig)
    default:
        fmt.Println("no activity")
    }
}