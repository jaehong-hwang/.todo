package cmd

func save() error {
	content, err := collection.GetTodosJSONString()
	if err != nil {
		return err
	}

	return todoFile.FillContent(content)
}
