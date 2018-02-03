package locator

type ServiceLocator struct {
	cache *Cache
}

func NewServiceLocator() *ServiceLocator {
	return &ServiceLocator{
		cache: new(Cache),
	}
}

func (sl *ServiceLocator) GetService(jndi string) Service {
	s := sl.cache.GetServices(jndi)
	if s != nil {
		return s
	}

	ctx := new(InitialContext)
	s = ctx.LookUp(jndi).(Service)
	sl.cache.AddService(s)
	return s
}
