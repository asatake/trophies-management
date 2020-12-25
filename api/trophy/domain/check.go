package domain

type Check map[string]bool

func (c Check) Check(key string) {
	c[key] = true
}

func (c Check) UnCheck(key string) {
	c[key] = false
}
