package todo

type Label struct {
	color string
	text  string
}

func (l *Label) AddTo(t Todo) {
	t.Labels = append(t.Labels, l)
}
