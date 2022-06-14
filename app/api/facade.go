package api

type Publisher interface {
	Public() interface{}
}

func Public(i interface{}) interface{} {

	if p, ok := i.(Publisher); ok {
		return p.Public()
	}

	return i
}
