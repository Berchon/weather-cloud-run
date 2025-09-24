package dependencies

type Handlers struct {
	NameHandler interface{} //handler.NameHandler
}

func BuildDependencies() *Handlers {
	return &Handlers{}
}
