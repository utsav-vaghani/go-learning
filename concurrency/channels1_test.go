package concurrency

import (
	"fmt"
	"testing"
)

// channels are orchestration for the go routines.
// channels are used to share memory by communicating.
// channels are not internally like queue but it behaves like it.
// Two types of channels: 1) unbuffered channels, 2) buffered channels

// 1. unbuffered channels
// - this type of channels are powerful channels as its giving gurantee of deliverying data or signal.
// - the sender and the receiver are blocked without each other. when both are ready channel will be unblocked and data will be transfered.
// - trade off is high latency, as channel will be blocked until both are not ready.
// - use this channel with extra care, if our main go routine is blocked you are out of business.

// 2. buffered channel
// - special type of channel which provides buffer to store the data.
// - this channel can be useful when one is looking for high latency.
// - With buffered channel, sender doesnt blocked for receiver to be ready to receive the data.
// - If the buffer can have more data sender can send data over channel and walk away.
// - If receiver is receiving data from the empty channel the it will be blocked.

// close channel
// - when channel is closed, sender can not send data over the channel and receiver can not receive it.
// - If the receiver is blocked due to empty buffered channel then receiver will be unblocked without the data.
// - Ideally only sender should close the channel, sender is telling that from now onwards no data will be published on this channel.

// Initialization of the Channel
// - for now,  we have only one way to initialize channel through only make function.
// - instance of the channel is like pointer to the large data-structure.
// - channel can be 1) read-write, 2) read only, 3) write only

func Test_unBufferedChannel(t *testing.T) {
	ch := make(chan string)

	go func() {
		fmt.Println(<-ch)
	}()

	ch <- "hey there"
}

func Test_bufferedChannel(t *testing.T) {
	ch := make(chan string, 2)

	go func() {
		ch <- "hey there"
		close(ch) // sender is closing channel.
	}()

	fmt.Println(<-ch)
	fmt.Println(<-ch) // trying to read second time but as channel is closed, wont be able to read but read operation will be performed gracefully.
}

func Test_readOnlyChannel(t *testing.T) {
	ch := make(chan string)

	go func(ch <-chan string) {
		fmt.Println(<-ch)
		// ch <- "hey there"  -- compile time error: cannot send data on receive only channel
	}(ch)

	ch <- "hey there"
}

func Test_writeOnlyChannel(t *testing.T) {
	ch := make(chan string)

	go func(ch chan<- string) {
		ch <- "hey there"
	}(ch)

	
	fmt.Println(<-ch)
	// ch <- "hey there"  -- compile time error: cannot receive data on send only channel
}
