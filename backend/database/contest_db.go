package database

type ContestDatabase interface {
	Get(contestName string) (Contest, error)
	List() ([]Contest, error)
	ActiveContests() ([]Contest, error)

	Edit(contest string, contestData Contest) error

	Add(contest Contest) error
	Delete(contest string) error
}
