1. go mod init `name_project`.
2. create folder config, controllers, model, routes, database dan helper.
3. create file `main.go`.
4. install package :

   - gorm = `$ $ go get -u gorm.io/gorm `
   - postgres driver = `$ go get -u gorm.io/driver/postgres`
   - viper = `$ go get github.com/spf13/viper`
   - gorilla mux = `$ go get -u github.com/gorilla/mux`
   - logrus = `$ go get github.com/sirupsen/logrus`
   - air = `$ air init` (jika golang air sudah terinstall di laptop)
   - godotenv = `$ go get github.com/joho/godotenv `
   - validator package = `$ go get github.com/go-playground/validator/v10     `

5. buat config file di folder config
6. buat database file di folder database
7. buat file index, authorRoutes, bookRoutes di folder routes
8. buat file authorController, bookController di folder controllers
9. buat file author, book di folder models
