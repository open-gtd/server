package domain

type (
	TypeEnum uint8
	Name     string
)

const (
	EmptyType TypeEnum = iota
	Label
	Area
	Contact
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

func (tag *tag) ConvertToContact() {
	tag.t = Contact
}

func (t tag) GetType() TypeEnum {
	return t.t
}

func (t tag) GetName() Name {
	return t.n
}

type Tag interface {
	Rename(name Name)
	ConvertToLabel()
	ConvertToArea()
	ConvertToContact()
	GetType() TypeEnum
	GetName() Name
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

func CreateContact(name Name) Tag {
	return &tag{
		n: name,
		t: Contact,
	}
}
