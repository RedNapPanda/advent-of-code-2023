package deque

import (
	"fmt"
	"testing"
)

func TestDeque_Resize(t *testing.T) {
	d := NewDeque[string](4)

	d.PushBack("A")
	d.PushBack("B")
	d.PushBack("C")
	d.PushBack("D")
	fmt.Printf("Cap %d\n", d.Cap())
	fmt.Printf("%+v\n", d)
	d.PushBack("E")
	fmt.Printf("Cap %d\n", d.Cap())
	fmt.Printf("%+v\n", d)
}

func TestDeque_PopFront(t *testing.T) {
	d := NewDeque[string](4)

	d.PushFront("A")
	fmt.Printf("Len %d\n", d.Len())
	fmt.Printf("Cap %d\n", d.Cap())
	fmt.Printf("%+v\n", d)
	d.PushFront("B")
	fmt.Printf("Len %d\n", d.Len())
	fmt.Printf("Cap %d\n", d.Cap())
	fmt.Printf("%+v\n", d)

	fmt.Printf("%+v\n", d.PopFront())
	fmt.Printf("Len %d\n", d.Len())
	fmt.Printf("Cap %d\n", d.Cap())
	fmt.Printf("%+v\n", d)
}

func TestDeque_PopBack(t *testing.T) {
	d := NewDeque[string](4)

	d.PushBack("A")
	fmt.Printf("Len %d\n", d.Len())
	fmt.Printf("Cap %d\n", d.Cap())
	fmt.Printf("%+v\n", d)
	d.PushBack("B")
	fmt.Printf("Len %d\n", d.Len())
	fmt.Printf("Cap %d\n", d.Cap())
	fmt.Printf("%+v\n", d)

	fmt.Printf("%+v\n", d.PopBack())
	fmt.Printf("Len %d\n", d.Len())
	fmt.Printf("Cap %d\n", d.Cap())
	fmt.Printf("%+v\n", d)
}
