package main

type list struct {
	buf  [1000]byte
	next *list
}

func main() {
	var l *list
	for {
		l = &list{next: l}
	}
}
