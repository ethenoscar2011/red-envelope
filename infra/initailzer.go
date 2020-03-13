package infra

type Initializer interface {
	Init()
}

type InitializerRegister struct {
	Initializers   []Initializer
}

func (i *InitializerRegister)Register(s Initializer)  {
	i.Initializers=append(i.Initializers,s)
}

