package database

type Contest_db interface {
	get(contest string) []Contest
	add(contest Contest)
	delete(contest string)
}
