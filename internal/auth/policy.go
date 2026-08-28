package auth

import "supplierhub/internal/model"

func CanRead(p model.Permission, resource string) bool {
	switch resource {
	case "inbound":
		return p.Inbound
	case "quality":
		return p.Quality
	case "settlement":
		return p.Settlement
	default:
		return false
	}
}
