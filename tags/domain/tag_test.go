package domain

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestCreateArea(t *testing.T) {
	name := Name("areaTagName")

	tag := CreateArea(name)
	assert.Equal(t, name, tag.GetName())
	assert.Equal(t, name, tag.GetOriginalName())
	assert.Equal(t, Area, tag.GetType())
}

func TestCreateLabel(t *testing.T) {
	name := Name("labelTagName")

	tag := CreateLabel(name)
	assert.Equal(t, name, tag.GetName())
	assert.Equal(t, name, tag.GetOriginalName())
	assert.Equal(t, Label, tag.GetType())
}

func TestCreateContact(t *testing.T) {
	name := Name("contactTagName")

	tag := CreateContact(name)
	assert.Equal(t, name, tag.GetName())
	assert.Equal(t, name, tag.GetOriginalName())
	assert.Equal(t, Contact, tag.GetType())
}

func TestRename(t *testing.T) {
	oldName := Name("tagName")
	newName := Name("newTagName")

	tag := CreateContact(oldName)
	tag.Rename(newName)

	assert.Equal(t, newName, tag.GetName())
	assert.Equal(t, oldName, tag.GetOriginalName())
}

func TestTagConvertToArea(t *testing.T) {
	oldName := Name("tagName")

	tag := CreateContact(oldName)
	tag.ConvertToArea()

	assert.Equal(t, Area, tag.GetType())
}

func TestTagConvertToContact(t *testing.T) {
	oldName := Name("tagName")

	tag := CreateArea(oldName)
	tag.ConvertToContact()

	assert.Equal(t, Contact, tag.GetType())
}

func TestTagConvertToLabel(t *testing.T) {
	oldName := Name("tagName")

	tag := CreateContact(oldName)
	tag.ConvertToLabel()

	assert.Equal(t, Label, tag.GetType())
}
