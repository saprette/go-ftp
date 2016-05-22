package common

type Pair struct {
	first  interface{}
	second interface{}
}

func (p *Pair) First() interface{} {
	return p.first
}

func (p *Pair) Second() interface{} {
	return p.second
}

func (p *Pair) SetFirst(first interface{}) {
	p.first = first
}

func (p *Pair) SetSecond(second interface{}) {
	p.second = second
}

func NewPair(first interface{}, second interface{}) *Pair {
	return &Pair{first, second}
}
