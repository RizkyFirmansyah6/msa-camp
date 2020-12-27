package queue

type Queue interface {
	Push(key interface{})
	Pop() interface{}
	Contains(key interface{}) bool
	Len() int
	Keys() []interface{}
}

// Implements Queue interface
type queueImpl struct {
	push     func(key interface{})
	pop      func() interface{}
	contains func(key interface{}) bool
	len      func() int
	keys     func() []interface{}
}

func (i queueImpl) Push(key interface{}) {
	i.push(key)
}
func (i queueImpl) Pop() interface{} {
	return i.pop()
}
func (i queueImpl) Contains(key interface{}) bool {
	return i.contains(key)
}
func (i queueImpl) Len() int {
	return i.len()
}
func (i queueImpl) Keys() []interface{} {
	return i.keys()
}

func New(size int) Queue {
	var elements = make([]interface{}, size)
	return queueImpl{
		push: func(x interface{}) {
			for i := 0; i < size; i++ {
				if elements[i] == nil {
					elements[i] = x
					break
				}
			}
		},
		pop: func() interface{} {
			temp := elements[0]
			elements = elements[1:]
			return temp
		},
		contains: func(key interface{}) bool {
			b := false
			for i := 0; i < size; i++ {
				if elements[i] == key {
					b = true
					break
				}
			}
			return b
		},
		len: func() int { return len(elements) },
		keys: func() []interface{} {
			var tmp = make([]interface{}, 0)
			for i := 0; i < size; i++ {
				if elements[i] != nil {
					tmp = append(tmp, elements[i])
				}
			}
			return tmp
		},
	}
}
