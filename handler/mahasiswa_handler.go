package handler

import (
	"be_latihan/model"
	"be_latihan/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetAllMahasiswa godoc
// @Summary Ambil semua data mahasiswa
// @Description Mengambil seluruh data mahasiswa. Endpoint ini hanya dapat diakses oleh admin.
// @Tags Mahasiswa
// @Produce json
// @Security BearerAuth
// @Success 200 {object} model.MahasiswaListSuccessResponse "Berhasil mengambil data mahasiswa"
// @Failure 401 {object} model.UnauthorizedResponse "Token tidak dikirim atau tidak valid"
// @Failure 403 {object} model.ForbiddenResponse "User tidak memiliki akses admin"
// @Failure 500 {object} model.InternalServerErrorResponse "Gagal mengambil data mahasiswa"
// @Router /api/mahasiswa/ [get]
func GetAllMahasiswa(c *fiber.Ctx) error {
	data, err := repository.GetAllMahasiswa()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Message: "gagal mengambil data mahasiswa",
			Error:   err.Error(),
		})
	}

	return c.JSON(model.Response{
		Message: "berhasil mengambil data mahasiswa",
		Data:    data,
	})
}

// GetMahasiswaByNPM godoc
// @Summary Ambil detail mahasiswa berdasarkan NPM
// @Description Mengambil satu data mahasiswa berdasarkan NPM. Endpoint ini hanya dapat diakses oleh admin.
// @Tags Mahasiswa
// @Produce json
// @Security BearerAuth
// @Param npm path string true "NPM mahasiswa"
// @Success 200 {object} model.MahasiswaDetailSuccessResponse "Berhasil mengambil data mahasiswa"
// @Failure 400 {object} model.BadRequestResponse "NPM tidak valid"
// @Failure 401 {object} model.UnauthorizedResponse "Token tidak dikirim atau tidak valid"
// @Failure 403 {object} model.ForbiddenResponse "User tidak memiliki akses admin"
// @Failure 404 {object} model.NotFoundResponse "Data mahasiswa tidak ditemukan"
// @Failure 500 {object} model.InternalServerErrorResponse "Gagal mengambil data mahasiswa"
// @Router /api/mahasiswa/{npm} [get]
func GetMahasiswaByNPM(c *fiber.Ctx) error {
	npm, err := strconv.ParseInt(c.Params("npm"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Message: "npm tidak valid",
			Error:   err.Error(),
		})
	}

	mhs, err := repository.GetMahasiswaByNPM(npm)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(model.Response{
				Message: "data mahasiswa tidak ditemukan",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Message: "gagal mengambil data mahasiswa",
			Error:   err.Error(),
		})
	}

	return c.JSON(model.Response{
		Message: "berhasil mengambil data mahasiswa",
		Data:    mhs,
	})
}

// InsertMahasiswa godoc
// @Summary Tambah data mahasiswa
// @Description Menambahkan data mahasiswa baru. Endpoint ini hanya dapat diakses oleh admin.
// @Tags Mahasiswa
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body model.Mahasiswa true "Data mahasiswa baru"
// @Success 201 {object} model.MahasiswaCreatedSuccessResponse "Berhasil menambahkan data mahasiswa"
// @Failure 400 {object} model.BadRequestResponse "Payload tidak valid"
// @Failure 401 {object} model.UnauthorizedResponse "Token tidak dikirim atau tidak valid"
// @Failure 403 {object} model.ForbiddenResponse "User tidak memiliki akses admin"
// @Failure 500 {object} model.InternalServerErrorResponse "Gagal menambahkan data mahasiswa"
// @Router /api/mahasiswa/ [post]
func InsertMahasiswa(c *fiber.Ctx) error {
	var payload model.Mahasiswa
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Message: "payload tidak valid",
			Error:   err.Error(),
		})
	}

	data, err := repository.InsertMahasiswa(&payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Message: "gagal menambahkan data mahasiswa",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(model.Response{
		Message: "berhasil menambahkan data mahasiswa",
		Data:    data,
	})
}

// UpdateMahasiswa godoc
// @Summary Ubah data mahasiswa
// @Description Mengubah data mahasiswa berdasarkan NPM. Endpoint ini hanya dapat diakses oleh admin.
// @Tags Mahasiswa
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param npm path string true "NPM mahasiswa"
// @Param request body model.Mahasiswa true "Data mahasiswa yang diperbarui"
// @Success 200 {object} model.MahasiswaUpdatedSuccessResponse "Berhasil mengubah data mahasiswa"
// @Failure 400 {object} model.BadRequestResponse "NPM atau payload tidak valid"
// @Failure 401 {object} model.UnauthorizedResponse "Token tidak dikirim atau tidak valid"
// @Failure 403 {object} model.ForbiddenResponse "User tidak memiliki akses admin"
// @Failure 404 {object} model.NotFoundResponse "Data mahasiswa tidak ditemukan"
// @Failure 500 {object} model.InternalServerErrorResponse "Gagal mengubah data mahasiswa"
// @Router /api/mahasiswa/{npm} [put]
func UpdateMahasiswa(c *fiber.Ctx) error {
	npm, err := strconv.ParseInt(c.Params("npm"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Message: "npm tidak valid",
			Error:   err.Error(),
		})
	}

	var payload model.Mahasiswa
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Message: "payload tidak valid",
			Error:   err.Error(),
		})
	}

	data, err := repository.UpdateMahasiswa(npm, &payload)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(model.Response{
				Message: "data mahasiswa tidak ditemukan",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Message: "gagal mengubah data mahasiswa",
			Error:   err.Error(),
		})
	}

	return c.JSON(model.Response{
		Message: "berhasil mengubah data mahasiswa",
		Data:    data,
	})
}

// DeleteMahasiswa godoc
// @Summary Hapus data mahasiswa
// @Description Menghapus data mahasiswa berdasarkan NPM. Endpoint ini hanya dapat diakses oleh admin.
// @Tags Mahasiswa
// @Produce json
// @Security BearerAuth
// @Param npm path string true "NPM mahasiswa"
// @Success 200 {object} model.MahasiswaDeletedSuccessResponse "Berhasil menghapus data mahasiswa"
// @Failure 400 {object} model.BadRequestResponse "NPM tidak valid"
// @Failure 401 {object} model.UnauthorizedResponse "Token tidak dikirim atau tidak valid"
// @Failure 403 {object} model.ForbiddenResponse "User tidak memiliki akses admin"
// @Failure 500 {object} model.InternalServerErrorResponse "Gagal menghapus data mahasiswa"
// @Router /api/mahasiswa/{npm} [delete]
func DeleteMahasiswa(c *fiber.Ctx) error {
	npm, err := strconv.ParseInt(c.Params("npm"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Message: "npm tidak valid",
			Error:   err.Error(),
		})
	}

	if err := repository.DeleteMahasiswa(npm); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Message: "gagal menghapus data mahasiswa",
			Error:   err.Error(),
		})
	}

	return c.JSON(model.Response{
		Message: "berhasil menghapus data mahasiswa",
	})
}
