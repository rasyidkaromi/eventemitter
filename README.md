# EventEmitter
Go EventEmitter


Install
---------------
go get github.com/rasyidkaromi/eventemitter
        
        
Usage
---------------
```go
	package main

	import "fmt"
	import "github.com/rasyidkaromi/eventemitter"


	func main() {

          	Event := eventemitter.New(1)
        	Event1 := Event.On("event1")
        	Event2 := Event.On("event2")
          
        	Event.Emit("`{
			"data": "data",
			"data 2": "data 2"
		}`", "event1")
              
        	Event.Emit("test 2", "event2")
          
        	fmt.Println(<-Event1)
        	fmt.Println(<-Event2)


           // Menambahkan event dalam channel yang sama | "eventbaru" ke channel Event1.          
        	Event.EventPlus(Event1, "eventbaru")

            // Emit data baru di event "event2"
        	Event.Emit("test 3", "event2")

        	fmt.Println(<-Event2)
        	
            // membuang "eventbaru" di dalam channel event Event1
        	Event.EventMin(Event1, "eventbaru")
          
	}
```

	dual core intel p6300 
	goos: windows
	goarch: amd64
	BenchmarkObject-2
	 5000	    793015 ns/op	    3289 B/op	       4 allocs/op


	*) kelebihannya yaitu kita bisa membuat dan menerima banyak event dalam satu channel dan memilahnya
	




