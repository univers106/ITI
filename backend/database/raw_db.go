package database

type RawDatabase interface {
	Get(contest string, competition string, student int) ([]Raw, error)
	GetByCompetition(contest string, competition string) ([]Raw, error)
	GetByStudent(contest string, student int) ([]Raw, error)

	Add(contest string, raw Raw) error
	Delete(contest string, raw Raw) error
}
