package todo

// Filters is list filter of collection
type Filters struct {
	WithDone bool
	Status   []string
}

// GetList with current filter
func (f *Filters) GetList(t Todos) Todos {
	todos := Todos{}
	for _, todo := range t {
		isValidate := true
		if len(f.Status) > 0 {
			for _, s := range f.Status {
				if s == todo.Status {
					isValidate = true
					break
				}
			}
		}

		if isValidate {
			todos = append(todos, todo)
		}
	}
	return todos
}
