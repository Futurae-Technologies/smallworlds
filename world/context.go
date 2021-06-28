package world

// Context type is feature vector.
type Context map[string]float64

// NewContext creates an empty context.
func NewContext() Context {
	return Context{}
}

// Len returns the size of the vector.
func (c Context) Len() int {
	return len(c)
}

// Empty is true if there are no features.
func (c Context) Empty() bool {
	return c.Len() == 0
}

// Keys returns all features' names.
func (c Context) Keys() []string {
	keys := make([]string, 0, 0)
	for key := range c {
		keys = append(keys, key)
	}

	return keys
}

// RightJoin adds new features from that, and overwrites
// feature values already in c with that's values.
func (c Context) RightJoin(that Context) Context {
	for _, key := range that.Keys() {
		c[key] = that[key]
	}

	return c
}

// LeftJoin adds only the features from that which
// which are not found in c.
func (c Context) LeftJoin(that Context) Context {
	for _, key := range that.Keys() {
		_, ok := c[key]
		if !ok {
			c[key] = that[key]
		}
	}

	return c
}
