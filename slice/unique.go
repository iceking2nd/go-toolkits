package slice

func Unique[S ~[]E, E any](s S) S {
	t := make(map[any]struct{})
	for _, v := range s {
		t[v] = struct{}{}
	}
	var n []E
	for k, _ := range t {
		n = append(n, k.(E))
	}
	return n
}
