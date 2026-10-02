package avatar

type Module struct{}

func Init() (*Module, error) {
	return &Module{}, nil
}
