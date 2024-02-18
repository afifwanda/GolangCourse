package helpers

func FindPersonById(id interface{}) *Person {
	for _, p := range people {
		if p.Id == id {
			return &p
		}
	}
	return nil
}

func FindPersonByName(name interface{}) *Person {
	for _, p := range people {
		if p.Nama == name {
			return &p
		}
	}
	return nil
}
