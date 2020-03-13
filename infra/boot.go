package infra

import (
	"github.com/tietang/props/kvs"
)

type BootApplication struct {
	conf         kvs.ConfigSource
	startContext StarterContext
}

func New(config kvs.ConfigSource) *BootApplication {
	b:=  &BootApplication{conf:config,startContext:StarterContext{}}
	b.startContext[KeyProps] = config
	return  b
}

func (b *BootApplication)Start()  {
	//1 初始化starter
	b.init()
	//2 安装startter
	b.setup()
	//3 启动starter
	b.start()
}

func (b *BootApplication)init()  {
	for _,starter := range StarterRegister.AllStarters() {
		starter.Init(b.startContext)
	}
}

func (b *BootApplication)setup()  {
	for _,starter := range StarterRegister.AllStarters() {
		starter.Setup(b.startContext)
	}
}

func (b *BootApplication)start(){
	for index,starter := range StarterRegister.AllStarters() {
		if starter.StartBlocking() {
			if index+1 == len(StarterRegister.AllStarters()) {
				starter.Start(b.startContext)
			}else { //使用携程来异步启动，防止阻塞后面的starter
				go 	starter.Start(b.startContext)
			}
		}else {
			starter.Start(b.startContext)
		}
	}
}