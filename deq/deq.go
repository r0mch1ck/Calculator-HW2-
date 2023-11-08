package deq

type deque struct {
	head   *elemDeque
	Tail   *elemDeque
	length int
}

type elemDeque struct {
	leftElem  *elemDeque
	rightElem *elemDeque
	Value     int
}

func ZeroDeque() *deque {
	tmp := deque{nil, nil, 0}
	return &tmp
}

func IsDequeEmpty(deque *deque) bool {
	return deque.length == 0
}

func ValueOfElemDeque(deque *deque, numOfElem int) *elemDeque {
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

func (deque *deque) appendLeft(value int) {
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

func (deque *deque) AppendRight(value int) {
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

func (deque *deque) popLeft() {
	deque.head = deque.head.rightElem
	deque.head.leftElem = nil
	deque.length -= 1
}

func (deque *deque) PopRight() {
	deque.Tail = deque.Tail.leftElem
	deque.Tail.rightElem = nil
	deque.length -= 1
}
