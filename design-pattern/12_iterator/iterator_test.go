package iterator

func ExampleIterator() {
	var aggregate Aggregate
	aggregate = NewNumbers(1, 10)
	IteratorPrint(aggregate.Iterator())
}
