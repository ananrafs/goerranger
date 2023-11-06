package goerranger

type Engine interface {
	Hit(func())
}

type MegaZord interface {
	Engine
	GetDisposer() Disposer
}
