# uid

Package for generate unique uid.

### install

```bash
go get github.com/go-jedi/uid
```

### example

init uid:

```bash
Initialize uid:

uid.NewUID(uid.Option{
  Chars: []rune("0123456789ABCDEFHKLMNPQRSTUVWXYZabcdefghkmnpqrstuvwxyz"),
  Cnt:   10, # length uid
})
```

using uid:
```bash
Example use uid:

# Chars = 0123456789ABCDEFHKLMNPQRSTUVWXYZabcdefghkmnpqrstuvwxyz
# Cnt = 10
uid.Generate() # ZUPgdh1TEX

# Chars = 0123456789ABCDEFHKLMNPQRSTUVWXYZabcdefghkmnpqrstuvwxyz
# Cnt = 15:
uid.Generate() # 7E5Z5AfaFxcACp8

# Chars = 0123456789ABCDEFHKLMNPQRSTUVWXYZabcdefghkmnpqrstuvwxyz
# Cnt = 20:
uid.Generate() # pRSHn1ZQyZaVpKYqedxZ
```