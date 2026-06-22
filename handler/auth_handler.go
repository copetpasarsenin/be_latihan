package handler

import (
	"be_latihan/config/middleware"
	"be_latihan/model"
	"be_latihan/pkg/password"
	"be_latihan/repository"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Register godoc
// @Summary Registrasi user baru
// @Description Membuat akun user baru untuk mendapatkan akses ke API.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body model.AuthRequest true "Data registrasi user"
// @Success 201 {object} model.RegisterSuccessResponse "Registrasi berhasil"
// @Failure 400 {object} model.BadRequestResponse "Payload tidak valid"
// @Failure 409 {object} model.ConflictResponse "Username sudah digunakan atau data tidak valid"
// @Failure 500 {object} model.InternalServerErrorResponse "Gagal membuat hash password"
// @Router /register [post]
func Register(c *fiber.Ctx) error {
	var payload model.AuthRequest
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Message: "payload tidak valid",
			Error:   err.Error(),
		})
	}

	payload.Username = strings.TrimSpace(payload.Username)
	payload.Role = strings.TrimSpace(payload.Role)
	if payload.Role == "" {
		payload.Role = "admin"
	}

	if payload.Username == "" || payload.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Message: "username dan password wajib diisi",
		})
	}

	hashedPassword, err := password.HashPassword(payload.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Message: "gagal membuat hash password",
			Error:   err.Error(),
		})
	}

	user := model.User{
		Username: payload.Username,
		Password: hashedPassword,
		Role:     payload.Role,
	}

	data, err := repository.InsertUser(&user)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(model.Response{
			Message: "username sudah digunakan atau data tidak valid",
			Error:   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(model.Response{
		Message: "register berhasil",
		Data: model.AuthUserResponse{
			ID:       data.ID,
			Username: data.Username,
			Role:     data.Role,
		},
	})
}

// Login godoc
// @Summary Login user
// @Description Melakukan autentikasi user dan mengembalikan token JWT untuk akses endpoint yang dilindungi.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body model.AuthRequest true "Data login user"
// @Success 200 {object} model.LoginSuccessResponse "Login berhasil"
// @Failure 400 {object} model.BadRequestResponse "Payload tidak valid"
// @Failure 401 {object} model.UnauthorizedResponse "Username atau password salah"
// @Failure 500 {object} model.InternalServerErrorResponse "Gagal memproses login"
// @Router /login [post]
func Login(c *fiber.Ctx) error {
	var payload model.AuthRequest
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Message: "payload tidak valid",
			Error:   err.Error(),
		})
	}

	user, err := repository.FindUserByUsername(strings.TrimSpace(payload.Username))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusUnauthorized).JSON(model.Response{
				Message: "username atau password salah",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Message: "gagal mencari user",
			Error:   err.Error(),
		})
	}

	if !password.CheckPasswordHash(payload.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Message: "username atau password salah",
		})
	}

	token, err := middleware.GenerateJWT(user, 2*time.Hour)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Message: "gagal membuat token",
			Error:   err.Error(),
		})
	}

	return c.JSON(model.Response{
		Message: "login berhasil",
		Data: model.LoginResponse{
			Token: token,
			User: model.AuthUserResponse{
				ID:       user.ID,
				Username: user.Username,
				Role:     user.Role,
			},
		},
	})
}

// ChangePassword godoc
// @Summary Ubah password user
// @Description Mengubah password user yang sedang login menggunakan token JWT.
// @Tags Auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body model.ChangePasswordRequest true "Data perubahan password"
// @Success 200 {object} model.ChangePasswordSuccessResponse "Password berhasil diubah"
// @Failure 400 {object} model.BadRequestResponse "Payload atau konfirmasi password tidak valid"
// @Failure 401 {object} model.UnauthorizedResponse "Token tidak valid atau password lama salah"
// @Failure 500 {object} model.InternalServerErrorResponse "Gagal mengubah password"
// @Router /change-password [put]
func ChangePassword(c *fiber.Ctx) error {
	username, ok := c.Locals("username").(string)
	if !ok || username == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Message: "user belum login",
		})
	}

	var payload model.ChangePasswordRequest
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Message: "payload tidak valid",
			Error:   err.Error(),
		})
	}

	if payload.OldPassword == "" || payload.NewPassword == "" || payload.ConfirmPassword == "" {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Message: "password lama, password baru, dan konfirmasi wajib diisi",
		})
	}

	if payload.NewPassword != payload.ConfirmPassword {
		return c.Status(fiber.StatusBadRequest).JSON(model.Response{
			Message: "konfirmasi password baru tidak sama",
		})
	}

	user, err := repository.FindUserByUsername(username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Message: "gagal mencari user",
			Error:   err.Error(),
		})
	}

	if !password.CheckPasswordHash(payload.OldPassword, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(model.Response{
			Message: "password lama salah",
		})
	}

	hashedPassword, err := password.HashPassword(payload.NewPassword)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Message: "gagal membuat hash password",
			Error:   err.Error(),
		})
	}

	if err := repository.UpdateUserPassword(username, hashedPassword); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.Response{
			Message: "gagal mengubah password",
			Error:   err.Error(),
		})
	}

	return c.JSON(model.Response{
		Message: "password berhasil diubah",
	})
}
