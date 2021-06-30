# Đồ án ocg

Web ecommerce using Golang and vuejs

Đồ án gồm 2 phần frontend và backend

Phần 1. Thư mục Backend
  1. Cấu trúc thư mục
```
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
```

  1. Vào thư mục backend: ``` cd > backend ```
  2. Lấy mã trong thư mục sql gồm 2 file OnlineShop.sql và AddData.sql bỏ vào mysql để chạy tạo cơ sở dữ liệu gồm các bảng và data
  
   ![image](https://user-images.githubusercontent.com/59026656/124004782-8fa9ed00-da02-11eb-97b8-85417e26cc07.png)

  4. Chạy lệnh ``` go get all ``` để tải xuống các package phụ thuộc trong file go.mod
  5. Chạy lệnh ``` go run . ``` để khởi chạy server. Server sẽ chạy trên cổng 3000 
 

Phần 2. Thư mục Frontend
  1. Vào thư mục frontend: ``` cd > frontend ```
      ```
      ├── public
      ├── src
      │   ├── assets
      │   ├── components
      │   ├── router
      │   ├── store
      │   └── views
      └── tests
        └── unit
      ```

  2. Chạy lệnh ``` yarn ínstall ``` để download các dependency cho project
  3. Chạy lệnh ``` yarn serve ``` để khởi chạy project
  4. Truy cập http://localhost:8080/ để xem giao diện
![Screenshot (1)](https://user-images.githubusercontent.com/59026656/124007104-3d1e0000-da05-11eb-85be-b5fa1f7cf4fa.png)
![Screenshot (2)](https://user-images.githubusercontent.com/59026656/124007205-5a52ce80-da05-11eb-8f53-650026099d0d.png)
![Screenshot (3)](https://user-images.githubusercontent.com/59026656/124007289-70f92580-da05-11eb-9f3e-a17798ac888c.png)
![Screenshot (4)](https://user-images.githubusercontent.com/59026656/124007362-81110500-da05-11eb-8829-9649121e6b4a.png)
![Screenshot (5)](https://user-images.githubusercontent.com/59026656/124007655-c2a1b000-da05-11eb-8721-6b20b6717865.png)



