package main

import (
	"fmt"
	"os"
	"reflect"
	"sesi-tiga/helpers"
)

// type Person struct {
// 	Id        int
// 	Nama      string
// 	Alamat    string
// 	Pekerjaan string
// 	Alasan    string
// }

// var people = []Person{
// 	{Id: 1, Nama: "Jane", Alamat: "Johannesburg", Pekerjaan: "Bartender", Alasan: "Switch Career"},
// 	{Id: 2, Nama: "Doe", Alamat: "Medan", Pekerjaan: "Pelajar", Alasan: "Menambah Ilmu"},
// 	{Id: 3, Nama: "William", Alamat: "Detroit", Pekerjaan: "Guru", Alasan: "Menambah Ilmu"},
// 	{Id: 4, Nama: "Dafoe", Alamat: "Lahore", Pekerjaan: "Mahasiswa", Alasan: "Menambah referensi"},
// }

func main() {
	var person *helpers.Person

	if len(os.Args) < 2 {
		fmt.Println("Tolong Masukan Nama atau Nomor Absen!")
		return
	}

	if reflect.TypeOf(os.Args[1]).Kind() == reflect.String {
		person = helpers.FindPersonByName(os.Args[1])
	} else {
		person = helpers.FindPersonById(os.Args[1])
	}

	if person == nil {
		fmt.Println("Data dengan nama/absen tsb tidak tersedia")
	} else {
		fmt.Printf("Id: %+v\n", *&person.Id)
		fmt.Printf("Nama: %+v\n", *&person.Nama)
		fmt.Printf("Alamat: %+v\n", *&person.Alamat)
		fmt.Printf("Pekerjaan: %+v\n", *&person.Pekerjaan)
		fmt.Printf("Alasan: %+v\n", *&person.Alasan)
	}
}

// func findPersonById(id interface{}) *Person {
// 	for _, p := range people {
// 		if p.Id == id {
// 			return &p
// 		}
// 	}
// 	return nil
// }

// func findPersonByName(name interface{}) *Person {
// 	for _, p := range people {
// 		if p.Nama == name {
// 			return &p
// 		}
// 	}
// 	return nil
// }
