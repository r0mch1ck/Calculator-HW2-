package deqchar

type DequeChar struct {
	head   *elemDequeChar
	Tail   *elemDequeChar
	length int
}

type elemDequeChar struct {
	leftElem  *elemDequeChar
	rightElem *elemDequeChar
	Value     uint8
}

func ZeroDequeChar() *DequeChar {
	tmp := DequeChar{nil, nil, 0}
	return &tmp
}

func IsDequeCharEmpty(DequeChar *DequeChar) bool {
	return DequeChar.length == 0
}

func ValueOfElemDequeChar(DequeChar *DequeChar, numOfElem int) *elemDequeChar {
	if numOfElem == 0 {
		return DequeChar.head
	} else if numOfElem >= DequeChar.length {
		panic("Num of elem is out of range")
	} else {
		tmp := DequeChar.head.rightElem
		for i := 0; i != numOfElem; i++ {
			tmp = tmp.rightElem
		}
		return tmp
	}
}

func (DequeChar *DequeChar) appendLeft(value uint8) {
	if IsDequeCharEmpty(DequeChar) {
		tmpElem := elemDequeChar{nil, nil, value}
		DequeChar.head = &tmpElem
		DequeChar.Tail = &tmpElem
		DequeChar.length += 1
	} else {
		tmpElem := elemDequeChar{nil, DequeChar.head, value}
		DequeChar.head.leftElem = &tmpElem
		DequeChar.head = &tmpElem
		DequeChar.length += 1
	}
}

func (DequeChar *DequeChar) AppendRight(value uint8) {
	if IsDequeCharEmpty(DequeChar) {
		tmpElem := elemDequeChar{nil, nil, value}
		DequeChar.head = &tmpElem
		DequeChar.Tail = &tmpElem
		DequeChar.length += 1
	} else {
		tmpElem := elemDequeChar{DequeChar.Tail, nil, value}
		DequeChar.Tail.rightElem = &tmpElem
		DequeChar.Tail = &tmpElem
		DequeChar.length += 1
	}
}

func (DequeChar *DequeChar) popLeft() {
	DequeChar.head = DequeChar.head.rightElem
	DequeChar.head.leftElem = nil
	DequeChar.length -= 1
}

func (DequeChar *DequeChar) PopRight() {
	DequeChar.Tail = DequeChar.Tail.leftElem
	if DequeChar.Tail != nil {
		DequeChar.Tail.rightElem = nil
	}
	DequeChar.length -= 1
}
