package repo

import (
	"github.com/TechMaster/golang/15GoMySQL/model"
)

func initManufacturer() {
	apple := model.Manufacturer{Name: "Apple", CountryCode: "US"}
	xiaomi := model.Manufacturer{Name: "Xiaomi", CountryCode: "CN"}
	vsmart := model.Manufacturer{Name: "Vsmart", CountryCode: "VN"}
	samsung := model.Manufacturer{Name: "Samsung", CountryCode: "KR"}
	nokia := model.Manufacturer{Name: "Nokia", CountryCode: "US"}

	Db.Create(&apple)
	Db.Create(&xiaomi)
	Db.Create(&vsmart)
	Db.Create(&samsung)
	Db.Create(&nokia)
}
