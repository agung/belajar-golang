package pointer

import (
	"fmt"
	"testing"
)

type Mahasiswa struct {
	ID, Nama, Alamat string
}

//pointer di struct method
func (m Mahasiswa) SetNama(nama string) {
	m.Nama = nama
}

// pointer di function
func ChangeNameMahasiswa(m *Mahasiswa) {
	m.Nama = "Mahasiswa Abadi"
}

func TestStruct(t *testing.T) {
	mahasiswa := Mahasiswa{
		ID:     "15030043",
		Nama:   "Agung Nugraha",
		Alamat: "Sleman, Yogyakarta",
	}
	fmt.Println(mahasiswa)
	//mahasiswa.SetNama("Nugraha Aja")
	//fmt.Println(mahasiswa)

	//function pointer
	ChangeNameMahasiswa(&mahasiswa)
	//var m *Mahasiswa = &mahasiswa
	//ChangeNameMahasiswa(m)
	fmt.Println(mahasiswa)
}
