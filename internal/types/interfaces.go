package types

type Thing interface {
	GetID() string
	GetPath() string
	GetTitle() string
}

type IGetThingFromID func(string) Thing
