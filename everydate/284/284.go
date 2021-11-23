package main

func main() {

}

type Iterator struct {
}

func (this *Iterator) hasNext() bool {
	// Returns true if the iteration has more elements.
	return true
}

func (this *Iterator) next() int {
	// Returns the next element in the iteration.
	return 0
}

type PeekingIterator struct {
	iter *Iterator
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{iter: iter}
}

func (this *PeekingIterator) hasNext() bool {
	return this.iter.hasNext()
}

func (this *PeekingIterator) next() int {
	return this.iter.next()
}

func (this PeekingIterator) peek() int {
	return this.iter.next()
}
