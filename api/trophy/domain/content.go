package domain

import "github.com/asatake/trophies-management/api/trophy/domain/valobj"

type Content struct {
	Id           int
	Title        valobj.Title
	Achievements []Achievement
}

func (c Content) Completed() bool {
	completed := true
	for _, a := range c.Achievements {
		completed = completed && a.Completed()
	}
	return completed
}

// return (rate, completed-count, all)
func (c Content) Progress() (float64, int, int) {
	completed := 0
	for _, a := range c.Achievements {
		if a.Completed() {
			completed++
		}
	}

	return float64(completed) / float64(len(c.Achievements)), completed, len(c.Achievements)
}
