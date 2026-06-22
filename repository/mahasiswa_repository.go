package repository

import (
	"fmt"
	"be_latihan/config"
	"be_latihan/model"
)

// Ambil semua data mahasiswa
func GetAllMahasiswa() ([]model.Mahasiswa, error) {
	var data []model.Mahasiswa
	result := config.GetDB().Find(&data)
	return data, result.Error
}

// Insert mahasiswa baru
func InsertMahasiswa(mhs *model.Mahasiswa) (*model.Mahasiswa, error) {
	result := config.GetDB().Create(mhs)
	return mhs, result.Error
}

// Ambil satu data mahasiswa berdasarkan NPM
func GetMahasiswaByNPM(npm int64) (model.Mahasiswa, error) {
	var mhs model.Mahasiswa
	npmStr := fmt.Sprintf("%d", npm)
	result := config.GetDB().First(&mhs, "npm = ?", npmStr)
	return mhs, result.Error
}

// Update data mahasiswa berdasarkan NPM
func UpdateMahasiswa(npm int64, newData *model.Mahasiswa) (*model.Mahasiswa, error) {
	var mhs model.Mahasiswa

	db := config.GetDB()

	npmStr := fmt.Sprintf("%d", npm)
	if err := db.First(&mhs, "npm = ?", npmStr).Error; err != nil {
		return nil, err
	}

	if err := db.Model(&mhs).Updates(newData).Error; err != nil {
		return nil, err
	}

	return &mhs, nil
}

// Hapus data mahasiswa berdasarkan NPM
func DeleteMahasiswa(npm int64) error {
	npmStr := fmt.Sprintf("%d", npm)
	result := config.GetDB().Where("npm = ?", npmStr).Delete(&model.Mahasiswa{})
	return result.Error
}
