package core

var store map[string]*Obj

type Obj struct {
	Value any
}

func init() {
	store = make(map[string]*Obj)
}

func NewObj(value any) *Obj {
	return &Obj{
		Value: value,
	}
}

func Put(k string, obj *Obj) {
	store[k] = obj
}
