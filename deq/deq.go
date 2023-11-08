package deq

type Deque struct {
	head   *elemDeque
	Tail   *elemDeque
	length int
}

type elemDeque struct {
	leftElem  *elemDeque
	rightElem *elemDeque
	Value     int
}

func ZeroDeque() *Deque {
	tmp := Deque{nil, nil, 0}
	return &tmp
}

func IsDequeEmpty(deque *Deque) bool {
	return deque.length == 0
}

func ValueOfElemDeque(deque *Deque, numOfElem int) *elemDeque {
	if numOfElem == 0 {
		return deque.head
	} else if numOfElem >= deque.length {
		panic("Num of elem is out of range")
	} else {
		tmp := deque.head.rightElem
		for i := 0; i != numOfElem; i++ {
			tmp = tmp.rightElem
		}
		return tmp
	}
}

func (deque *Deque) appendLeft(value int) {
	if IsDequeEmpty(deque) {
		tmpElem := elemDeque{nil, nil, value}
		deque.head = &tmpElem
		deque.Tail = &tmpElem
		deque.length += 1
	} else {
		tmpElem := elemDeque{nil, deque.head, value}
		deque.head.leftElem = &tmpElem
		deque.head = &tmpElem
		deque.length += 1
	}
}

func (deque *Deque) AppendRight(value int) {
	if IsDequeEmpty(deque) {
		tmpElem := elemDeque{nil, nil, value}
		deque.head = &tmpElem
		deque.Tail = &tmpElem
		deque.length += 1
	} else {
		tmpElem := elemDeque{deque.Tail, nil, value}
		deque.Tail.rightElem = &tmpElem
		deque.Tail = &tmpElem
		deque.length += 1
	}
}

func (deque *Deque) popLeft() {
	deque.head = deque.head.rightElem
	deque.head.leftElem = nil
	deque.length -= 1
}

func (deque *Deque) PopRight() {
	deque.Tail = deque.Tail.leftElem
	deque.Tail.rightElem = nil
	deque.length -= 1
}
