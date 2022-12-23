package set

type void struct{}
type Set map[byte]void

func (s Set) AddValue(char byte) {
	var voidElem void
	s[char] = voidElem
}

func (s Set) ValueExists(char byte) bool {
	if _, ok := s[char]; ok {
		return true
	} else {
		return false
	}
}

func (s Set) FillFromString(line string) {
	for i := 0; i < len(line); i += 1 {
		s.AddValue(line[i])
	}
}
