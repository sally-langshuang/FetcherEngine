package graph

type person struct {
	quiet int
	id int
	minQuiet int
	minQuietID int
	richer []*person
}

func (p *person) getMinQuiet() (int, int) {
	if p.minQuiet == -1 {
		p.minQuiet, p.minQuietID = p.quiet, p.id
		for i := range p.richer {
			if quietVal, id := p.richer[i].getMinQuiet(); quietVal < p.minQuiet {
				p.minQuiet = quietVal
				p.minQuietID = id
			}
		}
	}
	return p.minQuiet, p.minQuietID
}
func loudAndRich(richer [][]int, quiet []int) []int {
	persons := make([]*person, len(quiet))
	ans := make([]int, len(quiet))
	for i := range quiet {
		persons[i] = &person{quiet: quiet[i], id: i, richer: []*person{}, minQuiet: -1}
	}
	for i := range richer {
		x, y := persons[richer[i][0]], persons[richer[i][1]]
		y.richer = append(y.richer, x)
	}
	for i := range persons {
		_, ans[i] = persons[i].getMinQuiet()
	}
	return ans
}
