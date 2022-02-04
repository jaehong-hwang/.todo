package todo

import "strings"

type Label struct {
	Text string
}

type Labels []*Label

func (l *Label) AddTo(t Todo) error {
	return t.AddLabel(l)
}

func (l *Labels) ToString() string {
	var strs []string
	for _, lb := range *l {
		strs = append(strs, lb.Text)
	}

	return strings.Join(strs, " / ")
}
