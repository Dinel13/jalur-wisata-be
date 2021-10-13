sql: no rows in result set karena kita hanya cek jika err != nil
tapi juga harus cek apakah err == empty row

gunakan queryRow instead of execContext to insert into db
if you want return the full resutl

passing value from params url harus ewat context
jadi params di ambil dari context 
lalu di essert jadi httpROuter.params
baru getName


nano /lib/systemd/system/goapp.service

[Unit]
Description=simple go application

[Service]
Type=simple
Restart=always
RestartSec=5s
ExecStart=/path/to/binary/file

[Install]
WantedBy=multi-user.target

service goapp start


eror di file upload karena pada saat di server os.getwd() meretutn
/ bukan lokasi project ini

tidak bisa akses file static karena di cofigurasi nginx
harus ada /image yang mengarah ke lokasi file
location /images {
   alias /var/www/project/assets 
}
alias digunakan karena location link dan file berbeda



