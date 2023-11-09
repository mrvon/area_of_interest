package xy_linked_list

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test1(t *testing.T) {
	s := CreateScene()
	objects := make(map[int]*Object)
	for i := 0; i < 1000; i++ {
		x := rand.Int() % 100
		y := rand.Int() % 100
		object := s.CreateObject(x, y)
		objects[object.ObjectId] = object
	}


	fmt.Println("------------------------------------- STEP 1")

	for _, object := range objects {
		s.Enter(object)
	}

	// s.Dump()

	fmt.Println("------------------------------------- STEP 2")

	for _, object := range objects {
		x := rand.Int() % 30
		y := rand.Int() % 30
		s.Move(object, x, y)
	}

	fmt.Println("------------------------------------- STEP 3")

	for _, object := range objects {
		s.Leave(object)
	}
}
