package board

type Cell rune

func (c *Cell) setValue(v rune) bool {
	if *c == 'X' || *c == 'O' {
		return false
	}
	*c = Cell(v)
	return true
}
