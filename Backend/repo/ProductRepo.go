package repo

import (
	"github.com/TechMaster/golang/15GoMySQL/model"
)

func initProduct() {
	var apple, samsung, xiaomi, nokia model.Manufacturer
	Db.Where("name = ?", "Samsung").First(&samsung)
	Db.Where("name = ?", "Apple").First(&apple)
	Db.Where("name = ?", "Xiaomi").First(&xiaomi)
	Db.Where("name = ?", "Nokia").First(&nokia)

	var phone_Cat, phukien_cat, phone_normal model.Category
	Db.Where("name LIKE ?", "Smart phone%").First(&phone_Cat)
	Db.Where("name LIKE ?", "Tai nghe%").First(&phukien_cat)
	Db.Where("name LIKE ?", "Điện thoại phổ thông%").First(&phone_normal)

	nokia_150 := model.Product{
		Name:         "Điện thoại Nokia 150",
		Description:  "Nokia 150 (2020) là phiên bản tiếp theo của Nokia 150 (2017) đã rất thành công trước đó. Được cải tiến nhiều nhất về thiết kế, nâng cấp bộ nhớ danh bạ đến 800 số và có bàn phím T9 lớn hơn dễ sử dụng.",
		Price:        660000,
		Madein:       "US",
		Manufacturer: &nokia,
		Category:     &phone_normal,
	}
	Db.Create(&nokia_150)

	nokia_210 := model.Product{
		Name:         "Điện thoại Nokia 210",
		Description:  "Mang đến MWC 2019 ngoài những chiếc smartphone hấp dẫn Nokia còn trình làng chiếc điện thoại cục gạch Nokia 210 dành cho những người thích sử dụng thiết bị nghe, gọi đơn giản.",
		Price:        790000,
		Madein:       "US",
		Manufacturer: &nokia,
		Category:     &phone_normal,
	}
	Db.Create(&nokia_210)

	iPhone_XR := model.Product{
		Name:         "iPhone XR",
		Description:  "Là chiếc điện thoại iPhone có mức giá dễ chịu, phù hợp với nhiều khách hàng hơn, iPhone Xr vẫn được ưu ái trang bị chip Apple A12 mạnh mẽ, màn hình tai thỏ cùng khả năng kháng nước kháng bụi.",
		Price:        25000000,
		Madein:       "US",
		Manufacturer: &apple,
		Category:     &phone_Cat,
	}
	Db.Create(&iPhone_XR)

	iPhone_12Pro := model.Product{
		Name:         "Điện thoại iPhone 12 Pro Max 128GB",
		Description:  "iPhone 12 Pro Max sở hữu diện mạo mới mới với khung viền được vát thẳng và mạnh mẽ như trên iPad Pro 2020, chấm dứt hơn 6 năm với kiểu thiết kế bo cong trên iPhone 6 được ra mắt lần đầu tiên vào năm 2014.",
		Price:        30990000,
		Madein:       "US",
		Manufacturer: &apple,
		Category:     &phone_Cat,
	}
	Db.Create(&iPhone_12Pro)

	iphone_x := model.Product{
		Name:         "Điện thoại iPhone X 64GB",
		Description:  "IPhone X 64 GB là cụm từ được rất nhiều người mong chờ muốn biết và tìm kiếm trên Google bởi đây là chiếc điện thoại mà Apple kỉ niệm 10 năm iPhone đầu tiên được bán ra",
		Price:        20000000,
		Madein:       "US",
		Manufacturer: &apple,
		Category:     &phone_Cat,
	}
	Db.Create(&iphone_x)

	samSung1 := model.Product{
		Name:         "Điện thoại Samsung Galaxy M51",
		Description:  "Galaxy M51 thuộc thế hệ Galaxy M đến từ Samsung và được nằm trong phân khúc giá tầm trung. Máy được nâng cấp và cải tiến với camera góc siêu rộng, dung lượng pin siêu khủng cùng vẻ ngoài sang trọng và thời thượng.",
		Price:        8490000,
		Madein:       "KR",
		Manufacturer: &samsung,
		Category:     &phone_Cat,
	}
	Db.Create(&samSung1)

	samSung2 := model.Product{
		Name:         "Điện thoại Samsung Galaxy A32 ",
		Description:  "Samsung Galaxy A32 4G là chiếc điện thoại thuộc phân khúc tầm trung nhưng sở hữu nhiều ưu điểm vượt trội về màn hình lớn sắc nét, bộ bốn camera 64 MP cùng vi xử lý hiệu năng cao và được bán ra với mức giá vô cùng tốt.",
		Price:        8490000,
		Madein:       "KR",
		Manufacturer: &samsung,
		Category:     &phone_Cat,
	}
	Db.Create(&samSung2)

	xiaomi_1 := model.Product{
		Name:         "Điện thoại Xiaomi Redmi Note 10 5G 8GB",
		Description:  "Redmi Note 10 5G một trong những mẫu điện thoại thuộc series Redmi Note 10 của Xiaomi, không chỉ sở hữu hiệu năng mạnh mẽ đáp ứng tốt nhu cầu chơi game, đây còn là chiếc điện thoại có hỗ trợ mạng 5G cho tốc độ kết nối nhanh chóng.",
		Price:        5490000,
		Madein:       "CN",
		Manufacturer: &xiaomi,
		Category:     &phone_Cat,
	}
	Db.Create(&xiaomi_1)

	tainghe_Awei_A920BS := model.Product{
		Name:         "Tai nghe Bluetooth Awei A920BS",
		Description:  "Tai nghe chống ổn chủ động thế hệ 3 của Sony, hiệu suất giảm tạp âm lên đến 95%",
		Price:        2299000,
		Madein:       "US",
		Manufacturer: &samsung,
		Category:     &phukien_cat,
	}
	Db.Create(&tainghe_Awei_A920BS)

}

/*
type Product struct {
	ID             uint
	Name           string
	Description    string
	Price          uint
	Madein         string
	Country        *Country `gorm:"foreignKey:Madein"`
	ManufacturerID uint
	Manufacturer   *Manufacturer
	CategoryID     uint
	Category       *Category
}
*/
