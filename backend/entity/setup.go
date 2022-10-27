package entity

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(&ScholarAdmin{},
		&ScholarStatus{},
		&ScholarType{},
		&Scholarship{},
	)
	db = database

	admin1 := ScholarAdmin{
		Name:  "Anucha Kleawthong",
		Email: "anucha@gmail.com",
	}
	db.Model(&ScholarAdmin{}).Create(&admin1)
	admin2 := ScholarAdmin{
		Name:  "Amin naja",
		Email: "hahahaha@gmail.com",
	}
	db.Model(&ScholarAdmin{}).Create(&admin2)

	status1 := ScholarStatus{
		StatusName: "ยังเปิดรับอยู่",
	}
	db.Model(&ScholarStatus{}).Create(&status1)
	status2 := ScholarStatus{
		StatusName: "ปิดแล้วค้าบ",
	}
	db.Model(&ScholarStatus{}).Create(&status2)
	status3 := ScholarStatus{
		StatusName: "อัตเดตอยู่จ้า",
	}
	db.Model(&ScholarStatus{}).Create(&status3)

	type1 := ScholarType{
		TypeName: "ทุนให้เปล่า",
	}
	db.Model(&ScholarType{}).Create(&type1)
	type2 := ScholarType{
		TypeName: "ทุนต่อเนื่อง",
	}
	db.Model(&ScholarType{}).Create(&type2)
	type3 := ScholarType{
		TypeName: "ทุนต่างประเทศ",
	}
	db.Model(&ScholarType{}).Create(&type3)

	// Scholar1
	db.Model(&Scholarship{}).Create(&Scholarship{
		ScholarName:     "ทุนเด็กเรียนดี",
		ScholarAdminID:  &admin1.ID,
		ScholarStatusID: &status1.ID,
		ScholarTypeID:   &type1.ID,
		ScholarDetail:   "ทุนที่จะให้เด็กเรียนดีสูงสุดถึง 20000 บาท",
	})
}
