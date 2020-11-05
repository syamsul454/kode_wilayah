package CheckWilayah

import (
	"strconv"

	"github.com/syamsul454/kode_wilayah/dataWilayah"
)

func Check(nik int64) (string, string, string) {
	t := strconv.FormatInt(int64(nik), 10)
	a := []rune(t)
	provinsi := string(a[0:2])
	kbtn := string(a[2:4])
	kcmtan := string(a[4:6])
	ConvertProvinsi, _ := strconv.ParseInt(provinsi, 10, 64)
	ConvertKecamatan, _ := strconv.ParseInt(provinsi+kbtn, 10, 64)
	propinsi := dataWilayah.Propinsi[provinsi]
	kabupaten := dataWilayah.Kbtn[ConvertProvinsi][kbtn]
	kecamatan := dataWilayah.Kecamatan[ConvertKecamatan][kcmtan]

	return propinsi, kabupaten, kecamatan
}
