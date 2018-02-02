package filter

type FilterManager struct {
	fc *FilterChain
}

func NewFilterManager(target *Target) *FilterManager {
	fc := NewFilterChain()
	fc.SetTarget(target)
	return &FilterManager{
		fc: fc,
	}
}

func (fm *FilterManager) SetFilter(filter Filter) {
	fm.fc.AddFilter(filter)
}

func (fm *FilterManager) FilterRequest(req string) {
	fm.fc.Execute(req)
}
