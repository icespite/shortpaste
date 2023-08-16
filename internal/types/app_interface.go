package types

type AppInf interface {
	ShouldLink307Redirect() bool
	GetFileDB() FileDB
	GetDataDB() DataDB
}
