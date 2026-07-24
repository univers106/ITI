package database

type Raw_db interface {
	get(contest string, competition string, student int) ([]Raw, error)
	getByCompetition(contest string, competition string) ([]Raw, error)
	getByStudent(contest string, student int) ([]Raw, error)

	add(contest string, raw Raw) error
	delete(contest string, raw Raw) error
}
