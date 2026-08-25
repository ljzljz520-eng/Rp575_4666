package audit

import (
	"encoding/json"
	"fmt"
	"supplierhub/internal/model"
	"supplierhub/internal/store"
)

type Logger struct{ s *store.Store }

func New(s *store.Store) *Logger { return &Logger{s: s} }
func (l *Logger) Record(actor, action, resource string) error {
	id := fmt.Sprintf("%s-%s-%s", actor, action, resource)
	return l.s.Put("audits", id, model.AuditLog{ID: id, Actor: actor, Action: action, Resource: resource, At: "deterministic"})
}
func (l *Logger) List() ([]model.AuditLog, error) {
	raw, e := l.s.List("audits")
	if e != nil {
		return nil, e
	}
	out := make([]model.AuditLog, 0, len(raw))
	for _, r := range raw {
		var x model.AuditLog
		if e := json.Unmarshal(r, &x); e == nil {
			out = append(out, x)
		}
	}
	return out, nil
}
