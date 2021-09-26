package slice

type String []string

func (s String) Contains(lookup string) bool {
	return s.IndexOf(lookup) >= 0
}

func (s String) IndexOf(lookup string) int {
	for i, v := range s {
		if v == lookup {
			return i
		}
	}

	return -1
}
