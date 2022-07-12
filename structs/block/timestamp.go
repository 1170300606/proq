package block

type Tiemstamp struct {
	Time int `json:"time"`
}

func NewTiemstamp(time int) Tiemstamp {
	ti := Tiemstamp{
		Time: time,
	}
	return ti
}

func (time *Tiemstamp) Islower(t2 Tiemstamp) bool {
	return (time.Time <= t2.Time)
}

func (time *Tiemstamp) Ishigher(t2 Tiemstamp) bool {
	return (time.Time >= t2.Time)
}

func (time *Tiemstamp) Equals(t2 Tiemstamp) bool {
	return (time.Time == t2.Time)
}
