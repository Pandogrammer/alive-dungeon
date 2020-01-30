package clock

type Clock struct {
	Actual int
}

func (t Clock) Advance() Clock {
	t.Actual = t.Actual + 1
	return t
}