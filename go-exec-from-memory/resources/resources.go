package resources

import "github.com/GeertJohan/go.rice"

func Box() *rice.Box {
	conf := rice.Config{
		LocateOrder: []rice.LocateMethod{rice.LocateEmbedded, rice.LocateAppended, rice.LocateFS},
	}

	box, err := conf.FindBox("data")
	if err != nil {
		panic(err)
	}

	return box
}
