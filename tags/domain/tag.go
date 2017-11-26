package domain

type (
	TypeEnum uint8
	Name string
)

const (
	EmptyType TypeEnum = iota
	Label
	Area
	Contact
)

type tag struct {
	original_name Name
	name          Name
	t             TypeEnum
}

func (tag *tag) Rename(name Name) {
	tag.name = name
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

func (t tag) GetOriginalName() Name {
	return t.original_name
}

func (t tag) GetName() Name {
	return t.name
}

type Tag interface {
	Rename(name Name)
	ConvertToLabel()
	ConvertToArea()
	ConvertToContact()
	GetType() TypeEnum
	GetOriginalName() Name
	GetName() Name
}

func CreateLabel(name Name) Tag {
	return &tag{
		original_name: name,
		name:          name,
		t:             Label,
	}
}

func CreateArea(name Name) Tag {
	return &tag{
		original_name: name,
		name:          name,
		t:             Area,
	}
}

func CreateContact(name Name) Tag {
	return &tag{
		original_name: name,
		name:          name,
		t:             Contact,
	}
}
