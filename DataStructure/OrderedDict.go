package main

import (
	"container/list"
	"fmt"
)

type Element struct {
	Key, Value interface{}

	element *list.Element
}

func newElement(e *list.Element) *Element {
	if e == nil {
		return nil
	}

	element := e.Value.(*OrderedMapElement)

	return &Element{
		element: e,
		Key:     element.key,
		Value:   element.val,
	}
}

// Next returns the next element, or nil if it finished.
func (e *Element) Next() *Element {
	return newElement(e.element.Next())
}

type OrderedMapElement struct {
	key, val interface{}
}
type OrderedMap struct {
	kv map[interface{}]*list.Element
	ll *list.List
}

func NewOrderedMap() *OrderedMap {
	return &OrderedMap{kv: make(map[interface{}]*list.Element),
		ll: list.New(),
	}
}

func (mp *OrderedMap) Get(key interface{}) (interface{}, bool) {
	value, ok := mp.kv[key]
	if ok {
		return value.Value.(*OrderedMapElement).val, true

	}

	return nil, false

}
func (mp *OrderedMap) Set(key, val interface{}) bool {
	_, exist := mp.kv[key]
	if !exist {
		//Key does not exist , insert into list first
		ele := mp.ll.PushBack(&OrderedMapElement{key, val})
		mp.kv[key] = ele
	} else {
		mp.Delete(key)
		mp.Set(key, val)
	}

	return !exist
}
func (mp *OrderedMap) Delete(key interface{}) bool {
	val, ok := mp.kv[key]
	if ok {
		mp.ll.Remove(val)
		delete(mp.kv, key)
	}

	return ok
}

func (m *OrderedMap) Front() *Element {
	front := m.ll.Front()
	if front == nil {
		return nil
	}

	element := front.Value.(*OrderedMapElement)

	return &Element{
		element: front,
		Key:     element.key,
		Value:   element.val,
	}
}
func main() {
	m := NewOrderedMap()

	m.Set("foo", "bar")
	m.Set("qux", 1.23)
	m.Set(123, true)
	m.Delete("qux")
	m.Set("foo", "baro")

	for el := m.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Key, el.Value)
	}

}
