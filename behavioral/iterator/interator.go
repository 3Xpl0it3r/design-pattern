package iterator

type Book struct {
	name string
}
type BookGroup struct {
	books []Book
}

func(b *BookGroup) Add(newb Book){
	if b == nil {
		return }
	b.books = append(b.books, newb)
}

func(b  *BookGroup)CreateInterator()*BookIterator{
	if b == nil {return nil}
	return &BookIterator{
		g:     b,
		index: 0,
	}
}

//
type Iterator interface {
	First()interface{}
	Next()interface{}
}

type BookIterator struct {
	g *BookGroup
	index int
}

func(b *BookIterator)First()interface{}{
	if b == nil {return nil}
	if len(b.g.books) > 0 {
		b.index = 0
		return b.g.books[b.index]
	}
	return nil
}

func(b *BookIterator)Next()interface{}{
	if b == nil {
		return nil
	}
	if len(b.g.books) > b.index+1 {
		b.index++
		return b.g.books[b.index]
	}
	return nil
}