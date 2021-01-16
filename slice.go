package main

import (
	"fmt"
	"strconv"
)

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice = months[4:7]
	fmt.Println(slice)
	fmt.Println("Length : " + strconv.Itoa(len(slice)))
	fmt.Println("Capacity : " + strconv.Itoa(cap(slice)))

	// jika mengubah arraynya maka slicenya pun akan ikut berubah
	//months[5] = "Diubah"
	//fmt.Println(slice)

	// ketika mengubah slicenya maka arraynya pun akan ikut berubah
	//slice[0] = "Mei Lagi"
	//fmt.Println(months)

	// menambah element baru pada array pada posisi terakhir pada slice
	// jika append dilakukan pada slice yg sudah melebihi kapasitas
	// maka akan membuat array baru
	//slice1 := tambahElement(slice)

	var slice2 = months[11:]

	// ini akan membuat array baru karena slice yg diappend tidak mempunyai
	// kapasitas lagi untuk menampung
	slice3 := tambahElement(slice2, "Element Baru")
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	slice4 := months[8:]
	slice5 := tambahElement(slice4, "Tambah Array")
	fmt.Println(slice5)

	// membuat slice baru
	sliceBaru := make([]string, 2, 5)
	sliceBaru[0] = "Agung"
	sliceBaru[1] = "Nugraha"
	fmt.Println(sliceBaru)
}

func tambahElement(monthSlice []string, e string) []string  {
	return append(monthSlice, e)
}
