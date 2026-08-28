package workflow

func StepNames() []string {
	return []string{"authenticate supplier", "check permission", "load inbound records", "load quality results", "load settlement bills", "write audit log", "export report"}
}
