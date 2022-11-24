package slice

type Predicate[T any] func(item T, i int) bool

func Some[T any](s []T, p Predicate[T]) bool {
	for i, item := range s {
		if p(item, i) {
			return true
		}
	}
	return false
}

func Chunk[T any](slice []T, chunkSize int) [][]T {
	var chunks [][]T

	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

func Contains[T comparable](s []T, e T) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Unique[T comparable](s []T) []T {
	inResult := make(map[T]bool)
	var result []T
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}

func FindIndex[T any](s []T, p Predicate[T]) int {
	for i, item := range s {
		if p(item, i) {
			return i
		}
	}
	return -1
}

func Map[T1, T2 any](input []T1, f func(T1) T2) (output []T2) {
	output = make([]T2, 0, len(input))
	for _, v := range input {
		output = append(output, f(v))
	}
	return output
}

func Filter[T any](slice []T, p Predicate[T]) []T {
	var n []T
	for i, e := range slice {
		if p(e, i) {
			n = append(n, e)
		}
	}
	return n
}
