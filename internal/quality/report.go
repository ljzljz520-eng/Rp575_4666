package quality

import "supplierhub/internal/model"

func PassedCount(items []model.QualityResult) int {
	n := 0
	for _, x := range items {
		if x.Passed {
			n++
		}
	}
	return n
}
func GradeSummary(items []model.QualityResult) map[string]int {
	m := map[string]int{}
	for _, x := range items {
		m[x.Grade]++
	}
	return m
}
