package models

import "gorm.io/gorm"

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&Role{}, &User{}, &PanicEvent{}, &AmbulanceLocation{})
	if err != nil {
		return err
	}

	roles := []Role{
		{Name: "admin"},
		{Name: "petugas"},
		{Name: "pasien"},
	}

	for _, r := range roles {
		var count int64
		db.Model(&Role{}).Where("name = ?", r.Name).Count(&count)
		if count == 0 {
			db.Create(&r)
		}
	}

	return nil
}
