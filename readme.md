# Saku Dompet API

## Author 
Condro Wiyono (github.com/condrowiyono)

## Desc
Proyek menggabungkan semua yang biasa berada di saku dompet, seperti berbagai macam kartu, boarding pass, e-tiket, kupon, tiket, dll. Terinspirasi dari Apple Wallet.

## Stack
GoLang dan MySQL dengan ORM

## Local
- Clone the repo
- Install Dep
```
go get github.com/joho/godotenv
go get github.com/julienschmidt/httprouter
go get github.com/jinzhu/gorm
go get github.com/go-sql-driver/mysql
```
- Copy env and create one
```sh
cp sample.env .env
```
- Run by either ```go build``` or ```go run api/main.go```
- See on browser ```localhost:3000```

## Endpoint
- GET /healthz
- GET /debits

## Progress
- [x] Arsitektur Dasar
- [ ] Go Dep
- SakuDompet
- [x] Kartu Debit
- [ ] Kartu Kredit
- [ ] OVO
- [ ] GOPay
- [ ] Dana
- [ ] Tiket KAI
- [ ] many more to come

## Thank You
