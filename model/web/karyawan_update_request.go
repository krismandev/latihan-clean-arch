package web

type KaryawanUpdateRequest struct {
	Nama         string `validate:"required" json:"nama"`
	Nip          string `validate:"required" json:"nip"`
	TempatLahir  string `validate:"required" json:"tempat_lahir"`
	TanggalLahir string `validate:"required" json:"tanggal_lahir"`
	Umur         int    `validate:"required" json:"umur"`
	Alamat       string `json:"alamat"`
	Agama        string `validate:"required" json:"agama"`
	JenisKelamin string `validate:"required" json:"jenis_kelamin"`
	NoHandphone  string `json:"no_handphone"`
	Email        string `validate:"required" json:"email"`
}
