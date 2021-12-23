package main

import (
	"testing"
)

func TestObserver(t *testing.T) {
		 reader1 := NewReader("Ghost")
		 reader2 := NewReader("Mount")
		 subj := NewSubject()
		 subj.Attach(reader1)
		 subj.Attach(reader2)

		 subj.Notify()
}