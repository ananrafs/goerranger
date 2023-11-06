package goerranger

func Init(engine func(Options) MegaZord, opt Options) (Engine, Disposer) {
	_engine := engine(opt)

	return _engine, _engine.GetDisposer()
}
