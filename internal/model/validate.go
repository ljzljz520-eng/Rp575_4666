package model

func (s Supplier) Valid() bool { return s.ID != "" && s.Name != "" && s.Active }
func (p Permission) Any() bool { return p.Inbound || p.Quality || p.Settlement }
