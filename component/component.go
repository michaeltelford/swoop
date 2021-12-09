package component

type (
	IComponent interface {
		Name() string
		Content() string
	}
)
