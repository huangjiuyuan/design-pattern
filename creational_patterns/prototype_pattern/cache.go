package shape

type ShapeCache struct {
	ShapeMap map[string]Shape
}

// GetShape returns a clone of the shape with the specified ID.
func (sc *ShapeCache) GetShape(id string) Shape {
	s := sc.ShapeMap[id]
	return s.Clone()
}

// LoadCache initializes the ShapeCache.
func (sc *ShapeCache) LoadCache() {
	// Initialize the map in the cache.
	if sc.ShapeMap == nil {
		sc.ShapeMap = make(map[string]Shape)
	}

	r := &rect{
		shape:  shape{shapeType: "rect"},
		width:  20,
		height: 15,
	}
	r.SetID("1")
	sc.ShapeMap[r.GetID()] = r

	s := &square{
		shape: shape{shapeType: "square"},
		width: 15,
	}
	s.SetID("2")
	sc.ShapeMap[s.GetID()] = s

	c := &circle{
		shape:  shape{shapeType: "circle"},
		radius: 10,
	}
	c.SetID("3")
	sc.ShapeMap[c.GetID()] = c
}
