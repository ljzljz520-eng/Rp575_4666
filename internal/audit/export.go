package audit

import (
	"encoding/json"
	"supplierhub/internal/model"
)

func Encode(items []model.AuditLog) ([]byte, error) { return json.Marshal(items) }
