package helpers

var people = []Person{
	{Id: 1, Nama: "Jane", Alamat: "Johannesburg", Pekerjaan: "Bartender", Alasan: "Switch Career"},
	{Id: 2, Nama: "Doe", Alamat: "Medan", Pekerjaan: "Pelajar", Alasan: "Menambah Ilmu"},
	{Id: 3, Nama: "William", Alamat: "Detroit", Pekerjaan: "Guru", Alasan: "Menambah Ilmu"},
	{Id: 4, Nama: "Dafoe", Alamat: "Lahore", Pekerjaan: "Mahasiswa", Alasan: "Menambah referensi"},
}

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
