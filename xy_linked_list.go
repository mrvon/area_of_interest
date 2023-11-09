package xy_linked_list

import (
	"fmt"
)

type Object struct {
	ObjectId int
	X        int
	Y        int
}

type ListNode struct {
	ObjectId int
	X        int
	Y        int
	XNext    *ListNode
	YNext    *ListNode
}

type Scene struct {
	XList       *ListNode
	YList       *ListNode
	MaxObjectId int
	ObjectMap   map[int]*Object
}

const (
	VisualRange = 5
)

func CreateScene() *Scene {
	return &Scene{
		XList:       &ListNode{},
		YList:       &ListNode{},
		MaxObjectId: 0,
		ObjectMap:   make(map[int]*Object),
	}
}

func (s *Scene) CreateObject(x int, y int) *Object {
	s.MaxObjectId++
	object := &Object{
		ObjectId: s.MaxObjectId,
		X:        x,
		Y:        y,
	}
	return object
}

func (s *Scene) NearSet(object *Object) map[int]bool {
	// X list
	xSet := make(map[int]bool)
	xNode := s.XList
	for xNode.XNext != nil {
		next := xNode.XNext
		if next.X <= object.X {
			if object.X-next.X <= VisualRange {
				xSet[next.ObjectId] = true
			}
		} else {
			if next.X-object.X <= VisualRange {
				xSet[next.ObjectId] = true
			} else {
				// Needn't go on
				break
			}
		}
		xNode = next
	}
	// Y list
	ySet := make(map[int]bool)
	yNode := s.YList
	for yNode.YNext != nil {
		next := yNode.YNext
		if next.Y <= object.Y {
			if object.Y-next.Y <= VisualRange {
				ySet[next.ObjectId] = true
			}
		} else {
			if next.Y-object.Y <= VisualRange {
				ySet[next.ObjectId] = true
			} else {
				// Needn't go on
				break
			}
		}
		yNode = next
	}
	n_set := make(map[int]bool)
	for id := range xSet {
		if id != object.ObjectId && ySet[id] {
			n_set[id] = true
		}
	}
	return n_set
}

func (s *Scene) SendEnterMessage(watcher *Object, object *Object) {
	fmt.Printf(
		"\tWatcher[%d](%d,%d) <- Object[%d](%d,%d) Enter\n",
		watcher.ObjectId, watcher.X, watcher.Y, object.ObjectId, object.X, object.Y,
	)
}

func (s *Scene) SendLeaveMessage(watcher *Object, object *Object) {
	fmt.Printf(
		"\tWatcher[%d](%d,%d) <- Object[%d](%d,%d) Leave\n",
		watcher.ObjectId, watcher.X, watcher.Y, object.ObjectId, object.X, object.Y,
	)
}

func (s *Scene) SendMoveMessage(watcher *Object, object *Object, old_x int, old_y int) {
	fmt.Printf(
		"\tWatcher[%d](%d,%d) <- Object[%d](%d,%d) Move to (%d,%d) \n",
		watcher.ObjectId, watcher.X, watcher.Y, object.ObjectId, old_x, old_y, object.X, object.Y,
	)
}

func (s *Scene) Enter(object *Object) {
	if s.ObjectMap[object.ObjectId] != nil {
		return
	}
	s.ObjectMap[object.ObjectId] = object
	fmt.Printf("Object[%d](%d,%d) Enter\n", object.ObjectId, object.X, object.Y)
	s.rawEnter(object)
	nearSet := s.NearSet(object)
	for id := range nearSet {
		s.SendEnterMessage(s.ObjectMap[id], object)
	}
}

func (s *Scene) rawEnter(object *Object) {
	newNode := &ListNode{
		ObjectId: object.ObjectId,
		X:        object.X,
		Y:        object.Y,
	}
	// X list
	xNode := s.XList
	for xNode.XNext != nil {
		next := xNode.XNext
		if object.X <= next.X {
			break
		} else {
			xNode = next
		}
	}
	newNode.XNext = xNode.XNext
	xNode.XNext = newNode
	// Y list
	yNode := s.YList
	for yNode.YNext != nil {
		next := yNode.YNext
		if object.Y <= next.Y {
			break
		} else {
			yNode = next
		}
	}
	newNode.YNext = yNode.YNext
	yNode.YNext = newNode
}

func (s *Scene) Leave(object *Object) {
	if s.ObjectMap[object.ObjectId] == nil {
		return
	}
	fmt.Printf("Object[%d](%d,%d) Leave\n", object.ObjectId, object.X, object.Y)
	nearSet := s.NearSet(object)
	for id := range nearSet {
		s.SendLeaveMessage(s.ObjectMap[id], object)
	}
	s.rawLeave(object)
	delete(s.ObjectMap, object.ObjectId)
}

func (s *Scene) rawLeave(object *Object) {
	// X list
	xNode := s.XList
	for xNode.XNext != nil {
		next := xNode.XNext
		if object.ObjectId == next.ObjectId {
			xNode.XNext = next.XNext
			break
		} else {
			xNode = next
		}
	}
	// Y list
	yNode := s.YList
	for yNode.YNext != nil {
		next := yNode.YNext
		if object.ObjectId == next.ObjectId {
			yNode.YNext = next.YNext
			break
		} else {
			yNode = next
		}
	}
}

func (s *Scene) Move(object *Object, new_x int, new_y int) {
	if s.ObjectMap[object.ObjectId] == nil {
		return
	}
	old_x := object.X
	old_y := object.Y
	fmt.Printf(
		"Object[%d](%d,%d) Move to (%d,%d)\n",
		object.ObjectId, old_x, old_y, new_x, new_y,
	)
	nearSetBefore := s.NearSet(object)
	s.rawLeave(object)
	object.X = new_x
	object.Y = new_y
	s.rawEnter(object)
	nearSetAfter := s.NearSet(object)
	for id := range nearSetBefore {
		if nearSetAfter[id] {
			s.SendMoveMessage(s.ObjectMap[id], object, old_x, old_y)
		} else {
			s.SendLeaveMessage(s.ObjectMap[id], object)
		}
	}
	for id := range nearSetAfter {
		if !nearSetBefore[id] {
			s.SendEnterMessage(s.ObjectMap[id], object)
		}
	}
}

func (s *Scene) Dump() {
	fmt.Println("-------- X --------")
	xNode := s.XList
	for xNode.XNext != nil {
		next := xNode.XNext
		fmt.Printf("[%d](%d,%d)\n", next.ObjectId, next.X, next.Y)
		xNode = next
	}
	fmt.Println("-------- Y --------")
	yNode := s.YList
	for yNode.YNext != nil {
		next := yNode.YNext
		fmt.Printf("[%d](%d,%d)\n", next.ObjectId, next.X, next.Y)
		yNode = next
	}
	fmt.Print("\n")
}
