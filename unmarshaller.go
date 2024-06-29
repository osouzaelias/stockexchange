package main

type UnMarshaller struct{}

func (u UnMarshaller) Consume(lower, upper int64) {
	for ; lower <= upper; lower++ {
		// Transformar dados crus em eventos utilizáveis
		// Neste caso, os eventos já estão no formato correto
	}
}
