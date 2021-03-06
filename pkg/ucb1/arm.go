package ucb1

type Arm struct {
	Count  int64
	Reward float64
}

func (a *Arm) AvgIncome() float64 {
	return a.Reward / float64(a.Count)
}

type Arms []Arm
