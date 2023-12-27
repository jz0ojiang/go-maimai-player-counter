# Go-maimai-player-counter

API service for maimai-player-counter

# Data Source

- [舞萌DX - 店铺分布](https://wc.wahlap.net/maidx/location/index.html)

- [Administrative-divisions-of-China](https://github.com/modood/Administrative-divisions-of-China)

# Configuration

- `host`: host of the service, default: `:8080`

- `database`: database configuration

    - `sqlite`: path of sqlite3 database file, default: `data/pcdata.sqlite3`

    - `leveldb`: path of leveldb database directory, default: `data/counter`

- `hcaptcha_secret`: secret key of hcaptcha

- `totp_secret`: secret key of totp, need to generate by yourself

# API Doc

[API Doc](/api_doc.md)