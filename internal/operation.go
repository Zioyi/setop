package internal

func Intersect(setA, setB []string) []string {
	result := []string{}
	m := map[string]int{}

	for _, i := range setA {
		m[i]++
	}

	for _, i := range setB {
		if _, ok := m[i]; ok {
			result = append(result, i)
		}
	}

	return result
}

func Union(setA, setB []string) []string {
	result := []string{}
	m := map[string]int{}

	for _, i := range setA {
		m[i]++

	}

	for _, i := range setB {
		m[i]++
	}

	for k, _ := range m {
		result = append(result, k)
	}

	return result
}

func Subtract(setA, setB []string) []string {
	result := []string{}
	m := map[string]int{}

	for _, i := range setA {
		m[i]++
	}

	for _, i := range setB {
		if _, ok := m[i]; ok {
			delete(m, i)
		}
	}

	for k, _ := range m {
		result = append(result, k)
	}

	return result
}
