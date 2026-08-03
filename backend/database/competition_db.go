package database

type CompetitionDatabase interface {
	List() ([]*Competition, error)
	GetById(id int) (*Competition, error)
	GetByName(name string) (*Competition, error)

	Add(competition *Competition) error
	Update(competition *Competition) error
	Delete(id int) error
}
