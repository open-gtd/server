package domain

type (
	TypeEnum uint8
	Name     string
)

const (
	Label   TypeEnum = 1
	Area    TypeEnum = 2
	Context TypeEnum = 3
)

type tag struct {
	n Name
	t TypeEnum
}

func (tag *tag) Rename(name Name) {
	tag.n = name
}

func (tag *tag) ConvertToLabel() {
	tag.t = Label
}

func (tag *tag) ConvertToArea() {
	tag.t = Area
}

func (tag *tag) ConvertToContext() {
	tag.t = Context
}

func (t *tag) GetType() TypeEnum {
	return t.t
}

type Tag interface {
	Rename(name Name)
	ConvertToLabel()
	ConvertToArea()
	ConvertToContext()
	GetType() TypeEnum
}

func CreateLabel(name Name) Tag {
	return &tag{
		n: name,
		t: Label,
	}
}

func CreateArea(name Name) Tag {
	return &tag{
		n: name,
		t: Area,
	}
}

func CreateContext(name Name) Tag {
	return &tag{
		n: name,
		t: Context,
	}
}
