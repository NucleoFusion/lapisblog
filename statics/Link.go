package statics

type Link map[string]string

func NewLink(name string, link string) Link {
	return Link{name: link}
}
