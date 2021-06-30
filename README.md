# Đồ án ocg

Web ecommerce using Golang and vuejs

Đồ án gồm 2 phần frontend và backend

Phần 1. Thư mục backend
  1. Cấu trúc thư mục

      ├── Config.md
      ├── GORM.md
      ├── ReadMe.md
      ├── app.go
      ├── config
      │   └── config.go
      ├── controller
      │   ├── CategoryController.go
      │   ├── ManufactureController.go
      │   └── ProductController.go
      ├── dev.yml
      ├── go.mod
      ├── go.sum
      ├── images
      │   ├── category.jpg
      │   ├── connect_db.jpg
      │   ├── country.jpg
      │   └── mysql_extension.jpg
      ├── model
      │   ├── Category.go
      │   ├── City.go
      │   ├── Country.go
      │   ├── Manufacturer.go
      │   ├── Product.go
      │   ├── ProductMedia.go
      │   └── ProductProperty.go
      ├── public
      │   └── img
      │       ├── acer-1.jpg
      │       ├── acer-2.jpg
      │       ├── iphone-12-1.jpg
      │       ├── iphone-x-1.jpg
      │       ├── iphone-xr-1.jpg
      │       ├── iphone-xr-2.jpg
      │       ├── iphone-xr-3.jpg
      │       ├── nokia_150_1.jpg
      │       ├── nokia_150_2.jpg
      │       ├── nokia_150_3.jpg
      │       ├── nokia_210_1.jpg
      │       ├── pc-1.jpg
      │       ├── pc-2.jpg
      │       ├── samsung-galaxy-a32-1.jpg
      │       ├── samsung-galaxy-m51-1.jpg
      │       ├── tai-nghe-bluetooth-awei-a920bs-1.jpg
      │       └── xiaomi-redmi-note-10-1.jpg
      ├── repo
      │   ├── CategoryRepo.go
      │   ├── CountryRepo.go
      │   ├── ManufactureRepo.go
      │   ├── ProductRepo.go
      │   └── Repo.go
      ├── routes
      │   └── ConfigRouter.go
      ├── sql
      │   ├── AddData.sql
      │   ├── DropTable.sql
      │   └── OnlineShop.sql
      └── tmp
          └── main.exe


  1. Vào thư mục backend: cd > backend
  2. Chạy lệnh go run . để chạy server
 

Phần 2. Thư mục backend

 
 1. Vào thư mục frontend: cd > frontend
  2. Chạy lệnh yarn ínstall
  3. Chạy lệnh yarn serve để khởi chạy project
  4. Truy cập http://localhost:8080/ để xem giao diện
