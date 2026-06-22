package model

type UnauthorizedResponse struct {
	Message string `json:"message" example:"authorization header wajib diisi"`
}

type ForbiddenResponse struct {
	Message string `json:"message" example:"user tidak memiliki akses untuk fitur ini"`
}

type BadRequestResponse struct {
	Message string `json:"message" example:"payload tidak valid"`
	Error   string `json:"error,omitempty" example:"detail error validasi"`
}

type NotFoundResponse struct {
	Message string `json:"message" example:"data mahasiswa tidak ditemukan"`
}

type ConflictResponse struct {
	Message string `json:"message" example:"username sudah digunakan atau data tidak valid"`
	Error   string `json:"error,omitempty" example:"duplicate key value violates unique constraint"`
}

type InternalServerErrorResponse struct {
	Message string `json:"message" example:"terjadi kesalahan pada server"`
	Error   string `json:"error,omitempty" example:"detail error server"`
}

type RegisterSuccessResponse struct {
	Message string           `json:"message" example:"register berhasil"`
	Data    AuthUserResponse `json:"data"`
}

type LoginSuccessResponse struct {
	Message string        `json:"message" example:"login berhasil"`
	Data    LoginResponse `json:"data"`
}

type ChangePasswordSuccessResponse struct {
	Message string `json:"message" example:"password berhasil diubah"`
}

type MahasiswaListSuccessResponse struct {
	Message string      `json:"message" example:"berhasil mengambil data mahasiswa"`
	Data    []Mahasiswa `json:"data"`
}

type MahasiswaDetailSuccessResponse struct {
	Message string    `json:"message" example:"berhasil mengambil data mahasiswa"`
	Data    Mahasiswa `json:"data"`
}

type MahasiswaCreatedSuccessResponse struct {
	Message string    `json:"message" example:"berhasil menambahkan data mahasiswa"`
	Data    Mahasiswa `json:"data"`
}

type MahasiswaUpdatedSuccessResponse struct {
	Message string    `json:"message" example:"berhasil mengubah data mahasiswa"`
	Data    Mahasiswa `json:"data"`
}

type MahasiswaDeletedSuccessResponse struct {
	Message string `json:"message" example:"berhasil menghapus data mahasiswa"`
}
