package main

import "fmt"

type Subject struct {
	observers []Observer
	context   string
}

func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}

func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) Notify() {
	for _, o := range s.observers {
		o.Update()
	}
}

func (s *Subject) UpdateContext(context string) {
	s.context = context
	s.Notify()
}

type Observer interface {
	Update()
}

type Reader struct {
	bookName string
}

func (r *Reader) SetBookName(str string) {
	r.bookName = str
}

func (r Reader) Update() {
	fmt.Println(r.bookName)
}

func NewReader(bookName string) *Reader {
	return &Reader{
		bookName: bookName,
	}
}

func main() {

}
