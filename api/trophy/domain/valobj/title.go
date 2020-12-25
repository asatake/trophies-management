package valobj

type Title struct {
	Ja string
	En string
}

func (t Title) Empty() bool {
	if t.Ja == "" && t.En == "" {
		return false
	}
	return true
}
