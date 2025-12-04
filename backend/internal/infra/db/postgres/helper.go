package postgres

func in[K comparable, V any](m map[K]V, k K) bool {
	_, ok := m[k]
	return ok
}
