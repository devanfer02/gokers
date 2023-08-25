package configs

var FacultyMap 	map[string]map[string]string

func InitMajors() {
	FacultyMap = make(map[string]map[string]string)

	FacultyMap["FH"] = map[string]string{
		"code": 	"501",
		"Ilmu Hukum": "01",
	}

	FacultyMap["FEB"] = map[string]string{
		"code": 	"502",
		"Akuntansi":                        "01",
		"Ekonomi Pembangunan":              "02",
		"Manajemen":                        "03",
		"Ekonomi Islam":                    "04",
		"Kewirausahaan":                   "05",
		"Ekonomi, Keuangan dan Perbankan": "06",
	}

	FacultyMap["FIA"] = map[string]string{
		"code": 	"503",
		"Ilmu Administrasi Negara/Publik":   "01",
		"Ilmu Administrasi Niaga/Bisnis":    "02",
		"Perpajakan":                        "03",
		"Ilmu Perpustakaan":                 "04",
		"Pariwisata":                        "05",
		"Administrasi Pendidikan":           "06",
	}
	
	
	FacultyMap["FP"] = map[string]string{
		"code": 	"504",
		"Agribisnis":           "01",
		"Agroekoteknologi":     "02",
		"Agribisnis (PSDKU Kediri)":       "03",
		"Agroekoteknologi (PSDKU Kediri)": "04",
		"Kehutanan":            "05",
	}

	FacultyMap["FAPET"] = map[string]string{
		"code":	"505",
		"Peternakan":           "01",
		"Peternakan (PSDKU Kediri)":      "02",
	}

	FacultyMap["FT"] = map[string]string{
		"code":	"506",
		"Teknik Sipil":                   "01",
		"Teknik Mesin":                   "02",
		"Teknik Pengairan":               "03",
		"Teknik Elektro":                 "04",
		"Teknik Industri":                "05",
		"Teknik Kimia":                   "06",
		"Perencanaan Wilayah dan Kota":   "07",
		"Arsitektur":                     "08",
	}

	FacultyMap["FK"] = map[string]string{
		"code":	"507",
		"Kedokteran":          "01",
		"Farmasi":             "02",
		"Kebidanan":           "03",
	}

	FacultyMap["FPIK"] = map[string]string{
		"code":	"508",
		"Agrobisnis Perikanan":         "01",
		"Budidaya Perairan":           "02",
		"Ilmu Kelautan":               "03",
		"Manajemen Sumberdaya Perairan": "04",
		"Pemanfaatan Sumberdaya Perikanan": "05",
		"Teknologi Hasil Perikanan":   "06",
		"Akuakultur (PSDKU Kediri)":   "07",
		"Sosial Ekonomi Perikanan (PSDKU Kediri)": "08",
	}

	FacultyMap["FMIPA"] = map[string]string{
		"code":	"509",
		"Biologi":                 "01",
		"Fisika":                  "02",
		"Kimia":                   "03",
		"Matematika":              "04",
		"Statistika":              "05",
		"Teknik Geofisika":        "06",
		"Instrumentasi":           "07",
		"Ilmu Aktuaria":           "08",
		"Sains Data":			   "09",
	}

	FacultyMap["FTP"] = map[string]string{
		"code":	"510",
		"Ilmu dan Teknologi Pangan":  "01",
		"Teknologi Industri Pertanian": "02",
		"Teknik Lingkungan":          "03",
		"Bioteknologi":               "04",
		"Teknologi Bioproses":        "05",
		"Teknik Pertanian dan Biosistem": "06",
	}

	FacultyMap["FISIP"] = map[string]string{
		"code":	"511",
		"Hubungan Internasional": "01",
		"Ilmu Komunikasi":       "02",
		"Ilmu Politik":          "03",
		"Psikologi":             "04",
		"Sosiologi":             "05",
		"Ilmu Pemerintahan":     "06",
	}

	FacultyMap["FIB"] = map[string]string{
		"code":	"512",
		"Sastra Inggris":               "01",
		"Sastra Jepang":               "02",
		"Sastra Cina":                "03",
		"Bahasa dan Sastra Perancis": "04",
		"Pendidikan Bahasa dan Sastra Indonesia": "05",
		"Pendidikan Bahasa Inggris":  "06",
		"Pendidikan Bahasa Jepang":   "07",
		"Antropologi (Sosial)":       "08",
		"Seni Rupa Murni":            "09",
	}

	FacultyMap["FKH"] = map[string]string{
		"code":	"513",
		"Pendidikan Dokter Hewan": "01",
	}

	FacultyMap["VOKASI"] = map[string]string{
		"code": "514",
		"Manajemen Perhotelan": "01",
		"Desain Grafis": "02",
		"Administrasi Bisnis": "03",
		"Keuangan dan Perbankan": "04",
		"Teknologi Informasi": "05",
	}

	FacultyMap["FILKOM"] = map[string]string{
		"code":	"515",
		"Teknik Informatika":           "01",
		"Teknik Komputer":              "02",
		"Sistem Informasi":             "03",
		"Teknologi Informasi":          "04",
		"Pendidikan Teknologi Informasi": "05",
	}

	FacultyMap["FKG"] = map[string]string{
		"code":	"516",
		"Kedokteran Gigi": "01",
	}

	FacultyMap["FIKES"] = map[string]string{
		"code":	"517",
		"Ilmu Keperawatan": "01",
		"Ilmu Gizi":        "02",
	}
}