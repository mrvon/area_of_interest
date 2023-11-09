package xy_linked_list

import (
	"math/rand"
	"testing"
)

func assert(b bool) {
	if !b {
		panic("assert failed.")
	}
}

func _Test1(t *testing.T) {
	s := CreateScene(true)
	o1 := s.CreateObject(10, 10)
	s.Enter(o1)
	o2 := s.CreateObject(15, 15)
	s.Enter(o2)
	o3 := s.CreateObject(18, 18)
	s.Enter(o3)
	ns1 := s.NearSet(o1)
	assert(ns1[o2.ObjectId])
	assert(!ns1[o3.ObjectId])
	ns2 := s.NearSet(o2)
	assert(ns2[o1.ObjectId])
	assert(ns2[o3.ObjectId])
	ns3 := s.NearSet(o3)
	assert(!ns3[o1.ObjectId])
	assert(ns3[o2.ObjectId])
	s.Leave(o1)
	ns2 = s.NearSet(o2)
	assert(!ns2[o1.ObjectId])
	assert(ns2[o3.ObjectId])
	s.Move(o2, 20, 20)
	ns2 = s.NearSet(o2)
	assert(ns2[o3.ObjectId])
	s.Move(o2, 25, 25)
	ns2 = s.NearSet(o2)
	assert(!ns2[o3.ObjectId])
}

func Test2(t* testing.T) {
	s := CreateScene(true)
	o1 := s.CreateObject(1, 9)
	s.Enter(o1)
	o2 := s.CreateObject(2, 7)
	s.Enter(o2)
	o3 := s.CreateObject(3, 5)
	s.Enter(o3)
	s.Dump()
}

func Benchmark1(b *testing.B) {
	s := CreateScene(false)
	objects := make(map[int]*Object)
	for i := 0; i < 10000; i++ {
		x := rand.Int() % 100
		y := rand.Int() % 100
		object := s.CreateObject(x, y)
		objects[object.ObjectId] = object
	}

	for _, object := range objects {
		s.Enter(object)
	}

	// s.Dump()

	for _, object := range objects {
		x := rand.Int() % 30
		y := rand.Int() % 30
		s.Move(object, x, y)
	}

	for _, object := range objects {
		s.Leave(object)
	}
}
