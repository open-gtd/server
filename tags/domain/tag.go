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

type Tag struct {
	originalName Name
	name         Name
	t            TypeEnum
}

func (tag *Tag) Rename(name Name) {
	tag.name = name
}

func (tag *Tag) ConvertToLabel() {
	tag.t = Label
}

func (tag *Tag) ConvertToArea() {
	tag.t = Area
}

func (tag *Tag) ConvertToContact() {
	tag.t = Contact
}

func (tag Tag) GetType() TypeEnum {
	return tag.t
}

func (tag Tag) GetOriginalName() Name {
	return tag.originalName
}

func (tag Tag) GetName() Name {
	return tag.name
}

func CreateLabel(name Name) Tag {
	return Tag{
		originalName: name,
		name:         name,
		t:            Label,
	}
}

func CreateArea(name Name) Tag {
	return Tag{
		originalName: name,
		name:         name,
		t:            Area,
	}
}

func CreateContact(name Name) Tag {
	return Tag{
		originalName: name,
		name:         name,
		t:            Contact,
	}
}
