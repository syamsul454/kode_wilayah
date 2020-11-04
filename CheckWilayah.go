package CheckWilayah

import (
	"fmt"
	"strconv"

	"github.com/syamsul454/kode_wilayah/dataWilayah"
)

func Check(nik int64) {
	t := strconv.FormatInt(int64(nik), 10)
	a := []rune(t)
	provinsi := string(a[0:2])
	kabupaten := string(a[2:4])
	kecamatan := string(a[4:6])
	ConvertProvinsi, _ := strconv.ParseInt(provinsi, 10, 64)
	ConvertKecamatan, _ := strconv.ParseInt(provinsi+kabupaten, 10, 64)
	fmt.Println("provinsi", dataWilayah.Propinsi[provinsi])
	fmt.Println("kabupaten", dataWilayah.Kbtn[ConvertProvinsi][kabupaten])
	fmt.Println("kecamatan", dataWilayah.Kecamatan[ConvertKecamatan][kecamatan])
	fmt.Println("nomor nik", t)
	return
}
