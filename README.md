# go_harians

... make:table Merchants=phone?type:uint64,name_toko?type:*string,address,nik?type:uint64

Merchants -> Nama Tabel
phone,name_toko,address,nik -> field tabel

- penulisan nama field menggunakan format snake_case
- setiap nama field diakhiri dengan tanda tanya "?"
- setiap field dipisahkan dengan koma ",".
- type defauld field string


# type
  ID           uint
  Name         string
  Email        *string
  Age          uint8
  Birthday     *time.Time
  MemberNumber sql.NullString
  ActivatedAt  sql.NullTime
  CreatedAt    time.Time
  UpdatedAt    time.Time
