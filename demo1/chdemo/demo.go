package chdemo

var ch chan bool = make(chan bool)

func Provider() {
	ch <- true
}

func Consumer() {
	<- ch
}
