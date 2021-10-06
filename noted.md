sql: no rows in result set karena kita hanya cek jika err != nil
tapi juga harus cek apakah err == empty row

gunakan queryRow instead of execContext to insert into db
if you want return the full resutl

passing value from params url harus ewat context
jadi params di ambil dari context 
lalu di essert jadi httpROuter.params
baru getName