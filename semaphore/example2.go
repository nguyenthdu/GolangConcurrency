package semaphore

import "fmt"

type Empty1 interface{}
type semaphore chan Empty1

// acquire n resources
func (s semaphore) P(n int) {
	e := new(Empty)
	for i := 0; i < n; i++ {
		s <- e
	}
}

// release n resources
func (s semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}
func main() {
	var N int //them vao tranh loi
	sem := make(semaphore, N)
	fmt.Println(sem) //them vao tranh loi
}

/* mutexes */
func (s semaphore) Lock() {
	s.P(1)
}
func (s semaphore) Unlock() {
	s.V(1)
}

/* signal-wait */
func (s semaphore) Wait(n int) {
	s.P(n)
}
func (s semaphore) Signal() {
	var n int //them vao tranh loi
	s.V(n)
}
