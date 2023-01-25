# Account-Service-App-Project
Latar Belakang
Account Service App adalah aplikasi CLI yang digunakan untuk management sebuah account.

## Jalankan di lokal Anda !!
- Clone repo ini di lokal komputer anda:
```
git clone https://github.com/TeguhPutra16/Account-Service-App-Project.git
```
- Ubah nama .env example menjadi .env
- Dalam .env ubah `YourMYSQLPassword` jadi password MySQL anda dan `DATABASENAME` jadi nama database yang anda punya contoh:
```
export DB_CONNECTION="root:qwert123@tcp(127.0.0.1:3306)/account_service_app?parseTime=true"
```
- Buat table secara manual di MySQL sesuai kode yang ada di file account_service_app_project.sql
- sebelum menjalankan program jangan lupa untuk menjalankan `source .env`.
- setelah itu `go run main.go` untuk menjalankannya.
