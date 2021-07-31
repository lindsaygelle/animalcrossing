package animal

type Animals interface {
	ForEach(func(string, Animal))
	Get(string) (Animal, bool)
	Len() int
}

type animals map[string]animal

func (v animals) ForEach(fn func(string, Animal)) {
	for k, v := range v {
		fn(k, v)
	}
}

func (v animals) Get(k string) (Animal, bool) {
	animal, ok := v[k]
	return animal, ok
}

func (v animals) Len() int {
	return len(v)
}

var (
	_ Animals = animals{}
)
