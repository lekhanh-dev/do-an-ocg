package repo

import (
	"github.com/TechMaster/golang/15GoMySQL/model"
)

func initCategory() {
	// categories := map[string]([]string){
	// 	"điện thoại, máy tính bảng": []string{"smart phone", "điện thoại phổ thông", "máy tính bảng", "máy đọc sách"},
	// 	"điện tử, điện lạnh": []string{"tivi", "dàn âm thanh", "tủ lạnh - tủ mát", "máy điều hoà", "phụ kiện điện lạnh", "máy giặt", "hút bụi", "lọc không khí", "rủa bát"},
	// 	"máy tính, laptop": []string{"desktop", "server", "laptop", "phụ kiện máy tính"},
	// 	"camera, camcoder": []string{"máy ảnh", "máy ảnh kỹ thuật số", "máy quay phim", "camera giám sá", "camera hành trình", "balo", "ống kính"},
	// 	"đồ bếp":           []string{"lò vi sóng", "máy say sinh tố", "máy đánh trứng", "bếp từ", "bếp hồng ngoại", "bếp ga"},
	// 	"âm thanh":         []string{"ampli", "microphone", "karaoke", "loa", "tai nghe"},
	// }
	categories := map[string]([]string){
		"Điện thoại": []string{"Smart phone", "Điện thoại phổ thông"},
		"Máy tính":   []string{"PC", "Laptop"},
		"Phụ kiện":   []string{"Sạc", "Tai nghe", "Ốp"},
	}

	for key, subcategories := range categories {
		category := model.Category{Name: key, ParentID: 0}
		result := Db.Create(&category)
		if result.Error != nil {
			panic(result.Error)
		}

		for _, subcat := range subcategories {
			subCategory := model.Category{Name: subcat, ParentID: category.ID}
			result := Db.Create(&subCategory)
			if result.Error != nil {
				panic(result.Error)
			}
		}

	}
}

func GetCategoryById(id int) model.Category {
	var category model.Category
	Db.Where("ID = ?", id).First(&category)

	return category
}

func GetCategoriesById(id int) []model.Category {
	var categories []model.Category
	Db.Where("parent_id = ?", id).Find(&categories)

	return categories
}
