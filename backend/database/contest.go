package database

type Contest struct {
	Id           int
	Name         string
	Competitions []int
	IsActive     bool
}
