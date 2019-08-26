package eventemitter

import	"log"
import	"testing"



func BenchmarkObject(b *testing.B) {
	ser := New(1000000)
	c1 := ser.On("ch1")
	var pnt int
	for i := 0; i < b.N; i++ {
		ser.Emit(`{
			"data": "data",
			"data 2": "data 2"
			}`, "ch1")

		if val, ok := <-c1; !ok {
			log.Println(" Error found on event ch1.")
		} else {
			pnt++
			log.Println(ok)
			log.Println("val :", val)
			log.Println(pnt)
		}
	}
}
