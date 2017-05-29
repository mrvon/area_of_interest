package xy_linked_list

import (
	"fmt"
	"math/rand"
	"testing"
)

func Test1(t *testing.T) {
	s := create_scene()
	objects := make(map[int]*Object)
	for i := 0; i < 1000; i++ {
		x := rand.Int() % 100
		y := rand.Int() % 100
		object := s.create_object(x, y)
		objects[object.object_id] = object
	}

	fmt.Println("------------------------------------- STEP 1")

	for _, object := range objects {
		s.enter(object)
	}

	fmt.Println("------------------------------------- STEP 2")

	for id := range objects {
		x := rand.Int() % 30
		y := rand.Int() % 30
		s.move(id, x, y)
	}

	fmt.Println("------------------------------------- STEP 3")

	for id := range objects {
		s.leave(id)
	}
}
