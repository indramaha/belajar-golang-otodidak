package main

import "fmt"

func main() {
	var ujian = 80
	var absensi = 80

	var lulusUjian = ujian > 80
	var lulusAbsensi = absensi > 80
	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	var hasil = lulusAbsensi && lulusUjian
	fmt.Println(hasil)

	fmt.Println(ujian > 75 && absensi > 70)
}
