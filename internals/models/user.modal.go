package models

type pendidikan struct {
	namaSekolah string
	jurusan     string
}

type dataDiri struct {
	nama              string
	foto              string
	email             string
	umur              int
	nomorTelepon      string
	statusMenikah     bool
	riwayatPendidikan []pendidikan
}

func TampilinData() dataDiri {
	profil := dataDiri{
		nama:          "M. Hanif Irfan",
		foto:          "hanif.jpg",
		email:         "mhirfan173@gmail.com",
		umur:          28,
		nomorTelepon:  "0812345678910",
		statusMenikah: false,
		riwayatPendidikan: []pendidikan{
			{namaSekolah: "UNP",
				jurusan: "Kimia"},
			{namaSekolah: "SMA Al-Washliyah 3 Medan", jurusan: "IPA"},
		},
	}
	return profil
}
