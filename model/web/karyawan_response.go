package web

type KaryawanResponse struct {
	Id           int    `json:"id"`
	Nama         string `json:"nama"`
	Nip          string `json:"nip"`
	TempatLahir  string `json:"tempat_lahir"`
	TanggalLahir string `json:"tanggal_lahir"`
	Umur         int    `json:"umur"`
	Alamat       string `json:"alamat"`
	Agama        string `json:"agama"`
	JenisKelamin string `json:"jenis_kelamin"`
	NoHandphone  string `json:"no_handphone"`
	Email        string `json:"email"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	DeletedAt    string `json:"deleted_at"`
}
