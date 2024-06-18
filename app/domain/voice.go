package domain

type Voice struct {
	Id      int
	Name    string
	Content string
}

func (voice Voice) Build(id int, name string, content string) Voice {
	return Voice{
		Id:      id,
		Name:    name,
		Content: content,
	}
}
