sql: no rows in result set karena kita hanya cek jika err != nil
tapi juga harus cek apakah err == empty row

gunakan queryRow instead of execContext to insert into db
if you want return the full resutl

