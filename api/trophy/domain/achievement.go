package domain

import "github.com/asatake/trophies-management/api/trophy/domain/valobj"

type Achievement struct {
	Id     int
	Title  valobj.Title
	Checks Check
}

func (a Achievement) Completed() bool {
	completed := true
	for _, v := range a.Checks {
		completed = completed && v
	}
	return completed
}
