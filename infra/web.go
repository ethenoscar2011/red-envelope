package infra

var initailizerRegister *InitializerRegister =new(InitializerRegister)

func RegiserApi(s Initializer)  {
	initailizerRegister.Register(s)
}

func GetApiInitializers() []Initializer  {
	return initailizerRegister.Initializers
}

type InitailizerStarter struct {
	BaseStarter
}

func (w *InitailizerStarter)Setup(ctx StarterContext)  {
	for _, v := range GetApiInitializers() {
		v.Init()
	}

}