# eventemitter
Go EventEmitter


Install
---------------
go get github.com/rasyidkaromi/eventemitter
        
        
Usage
---------------

	package main

	import "fmt"
	import "github.com/rasyidkaromi/eventemitter"


	func main() {

          	Event := eventemitter.New(1)
        	Event1 := Event.On("event1")
        	Event2 := Event.On("event1")
          
        	Event.Emit("`{
			"data": "data",
			"data 2": "data 2"
		}`", "Event1")
              
        	Event.Emit("test 2", "Event2")
          
        	fmt.Println(<-Event1)
        	fmt.Println(<-Event2)


           // Menambahkan event dalam channel yang sama "event baru" ke event1.          
        	Event.EventPlus(Event1, "event baru")

            // Publish new content in Event2
        	Event.Emit("test 3", "Event2")

        	fmt.Println(<-Event2)
        	
            // membuang event Event2 di dalam channel event Event1
        	Event.EventMin(Event1, "Event2")
          
	}


	dual core intel p6300 
	goos: windows
	goarch: amd64
	BenchmarkObject-2
	 5000	    793015 ns/op	    3289 B/op	       4 allocs/op






