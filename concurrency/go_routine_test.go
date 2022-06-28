package concurrency

import (
	"fmt"
	"testing"
	"time"
)

// go routines are the light-weighted thread. Its normal go function with prefix `go`.
// There is one main go routine which starts when we execute the program.
// when we add any new go routine, it will be added as child go routine to main go routine.
// If main go routines keeps on running and completes the execution no other go routines will get chance to be executed.

// As per below example, If line 25 is not there `hey there` wont be logged on console.
// With presence of line 25, we are freezing the execution of the main go routine which gives a chance to run anonymous go routine.

// How it execute above code.
// - on detection of go routine, it adds that go routine in stack(without executine it) once it get chance(on blocking main thread) it gets executed.
// TODO: add how go routine is lightweighted thread
func Test_GoRoutine(t *testing.T) {
	go func() {
		fmt.Println("hey there")
	}()

	time.Sleep(time.Second)
}
