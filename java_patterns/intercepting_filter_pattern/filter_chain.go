package filter

type FilterChain struct {
	filters []Filter
	target  *Target
}

func NewFilterChain() *FilterChain {
	return &FilterChain{
		filters: make([]Filter, 0),
		target:  new(Target),
	}
}

func (fc *FilterChain) AddFilter(filter Filter) {
	fc.filters = append(fc.filters, filter)
}

func (fc *FilterChain) Execute(req string) {
	for _, f := range fc.filters {
		f.Execute(req)
	}
	fc.target.Execute(req)
}

func (fc *FilterChain) SetTarget(target *Target) {
	fc.target = target
}
