package service

func equal[T comparable](a, b *T) bool {
	if a == b {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	return *a == *b
}
