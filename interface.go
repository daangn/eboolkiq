package eboolkiq

type DB interface {
	jobDB
	queueDB
}
