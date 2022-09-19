package main

type Command struct {
	TheShiz string
}

type Repository[T any] interface {
	GetCollection(name string) Collection[T]
}

type TheRepo[T any] struct {
	collections map[string]any
}

func NewRepo[T any]() Repository[T] {
	return TheRepo[T]{collections: make(map[string]any)}
}

func (t TheRepo[T]) GetCollection(name string) Collection[T] {
	collection, ok := t.collections[name]
	if ok {
		return collection.(Collection[T])
	}
	newCollection := &TheCollection[T]{elems: make(map[string]T)}
	t.collections[name] = Collection[T](newCollection)
	return newCollection
}

type Collection[T any] interface {
	Get(id string) T
	Put(id string, elem T)
}

type TheCollection[T any] struct {
	elems map[string]T
}

func (t *TheCollection[T]) Get(id string) T {
	return t.elems[id]
}

func (t *TheCollection[T]) Put(id string, elem T) {
	t.elems[id] = elem
}

func main() {
	//repo := NewRepo()
	//collection := Repository[Command](repo).GetCollection("ayo")
	//collection.Put("ayo", Command{TheShiz: "Noooowaayyy"})
	//fmt.Println(collection.Get("ayo"))
}
