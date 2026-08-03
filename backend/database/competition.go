package database

type Competition struct {
	id   int
	Name string
}

func (c *Competition) GetID() int {
	return c.id
}
