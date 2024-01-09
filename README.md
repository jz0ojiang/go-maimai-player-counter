# Go-maimai-player-counter

API service for maimai-player-counter

# Data Source

- [全国音游地图](https://map.bemanicn.com/)【已授权使用数据】

- [舞萌DX - 店铺分布](https://wc.wahlap.net/maidx/location/index.html)【迁移至全国音游地图】

- [Administrative-divisions-of-China](https://github.com/modood/Administrative-divisions-of-China)

# Configuration

- `host`: host of the service, default: `:8080`

- `database`: database configuration

    - `sqlite`: path of sqlite3 database file, default: `data/pcdata.sqlite3`

    - `leveldb`: path of leveldb database directory, default: `data/counter`

- `captcha`: choose captcha service, default: `hCaptcha`

    - `hCaptcha`: use hCaptcha service

    - `turnstile/Turnstile`: use turnstile service

- `hcaptcha_secret`: secret key of hcaptcha

- `turnstile_secret`: secret key of cloudflare turnstile

- `totp_secret`: secret key of totp, need to generate by yourself

# API Doc

[API Doc](/api_doc.md)