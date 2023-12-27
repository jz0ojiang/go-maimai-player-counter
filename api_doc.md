# go-maimai-player-counter

github.com/jz0ojiang/go-maimai-player-counter

API 文档

<!--- If we have only one group/collection, then no need for the "ungrouped" heading -->

## Variables

| Key | Value          | Type   |
| --- | -------------- | ------ |
| url | localhost:8080 | string |

## Endpoints

- [Province（省）](#province)
  1. [getProvinceList（获取省份列表）](#1-getprovincelist)
     - [成功示例](#i-example-request-)
- [City（市）](#city)
  1. [getCityList（获取某省城市列表）](#1-getcitylist)
     - [成功示例（fullcode）](#i-example-request-fullcode)
     - [成功示例（code）](#ii-example-request-code)
     - [失败示例（Not Found）](#iii-example-request-not-found)
  1. [getCity（获取城市）](#2-getcity)
     - [成功示例（code）](#i-example-request-code)
     - [失败示例（Not Found）](#ii-example-request-not-found)
- [Arcade（机厅）](#arcade)
  1. [getArcade（获取机厅详情）](#1-getarcade)
     - [成功示例](#i-example-request--1)
     - [失败示例](#ii-example-request-)
  1. [getArcadeList（获取机厅列表）](#2-getarcadelist)
     - [成功示例](#i-example-request--2)
  1. [getArcadeList/city（获取某城市机厅列表）](#3-getarcadelistcity)
     - [成功示例](#i-example-request--3)
     - [失败示例](#ii-example-request--1)
  1. [getArcadeList/province（获取某省份机厅列表）](#4-getarcadelistprovince)
     - [成功示例](#i-example-request--4)
     - [失败示例](#ii-example-request--2)
- [Count（人数）](#count)
  1. [getCount（获取机厅人数）](#1-getcount)
     - [成功示例](#i-example-request--5)
     - [失败示例](#ii-example-request--3)
  1. [logCount（记录机厅人数）](#2-logcount)
     - [成功示例](#i-example-request--6)
     - [失败示例（token 错误）](#ii-example-request-token)
     - [失败示例（type=0）](#iii-example-request-type0)
     - [失败示例（数量错误）](#iv-example-request-)
- [Admin（管理员用）](#admin)
  1. [createCustomArcade（创建机厅）](#1-createcustomarcade)
     - [示例](#i-example-request--7)
  1. [deleteArcade（删除机厅）](#2-deletearcade)
  1. [updateArcade（从华立更新机厅）](#3-updatearcade)
  1. [generateToken（生成 token）](#4-generatetoken-token)

---

## Province（省）

### 1. getProvinceList（获取省份列表）

**_Endpoint:_**

```bash
Method: GET
Type:
URL: {{url}}/getProvinceList
```

**_More example Requests/Responses:_**

#### I. Example Request: 成功示例

**_Body: None_**

#### I. Example Response: 成功示例

```js
{
    "code": 0,
    "data": [
        {
            "code": 11,
            "name": "北京市",
            "full_code": "110000000000"
        }
        ...
    ],
    "message": "success"
}
```

**_Status Code:_** 200

<br>

## City（市）

### 1. getCityList（获取某省城市列表）

**_Endpoint:_**

```bash
Method: GET
Type:
URL: {{url}}/getCityList/:provinceCode
```

**_URL variables:_**

| Key          | Value | Description |
| ------------ | ----- | ----------- |
| provinceCode | 11    |             |

**_More example Requests/Responses:_**

#### I. Example Request: 成功示例（fullcode）

**_Query:_**

| Key          | Value        | Description |
| ------------ | ------------ | ----------- |
| provinceCode | 110000000000 |             |

**_Body: None_**

#### I. Example Response: 成功示例（fullcode）

```js
{
    "code": 0,
    "data": [
        {
            "code": 1101,
            "name": "市辖区",
            "province_code": 11,
            "full_code": "110100000000"
        }
    ],
    "message": "success"
}
```

**_Status Code:_** 200

<br>

#### II. Example Request: 成功示例（code）

**_Query:_**

| Key          | Value | Description |
| ------------ | ----- | ----------- |
| provinceCode | 11    |             |

**_Body: None_**

#### II. Example Response: 成功示例（code）

```js
{
    "code": 0,
    "data": [
        {
            "code": 1101,
            "name": "市辖区",
            "province_code": 11,
            "full_code": "110100000000"
        }
    ],
    "message": "success"
}
```

**_Status Code:_** 200

<br>

#### III. Example Request: 失败示例（Not Found）

**_Query:_**

| Key          | Value | Description |
| ------------ | ----- | ----------- |
| provinceCode | 0     |             |

**_Body: None_**

#### III. Example Response: 失败示例（Not Found）

```js
{
    "code": -1,
    "data": [],
    "message": "no city found"
}
```

**_Status Code:_** 404

<br>

### 2. getCity（获取城市）

**_Endpoint:_**

```bash
Method: GET
Type:
URL: {{url}}/getCity/:cityCode
```

**_URL variables:_**

| Key      | Value | Description |
| -------- | ----- | ----------- |
| cityCode | 1401  |             |

**_More example Requests/Responses:_**

#### I. Example Request: 成功示例（code）

**_Query:_**

| Key      | Value | Description |
| -------- | ----- | ----------- |
| cityCode | 1401  |             |

**_Body: None_**

#### I. Example Response: 成功示例（code）

```js
{
    "code": 0,
    "data": {
        "code": 1401,
        "name": "太原市",
        "province_code": 14,
        "full_code": "140100000000"
    },
    "message": "success"
}
```

**_Status Code:_** 200

<br>

#### II. Example Request: 失败示例（Not Found）

**_Query:_**

| Key      | Value | Description |
| -------- | ----- | ----------- |
| cityCode | 0     |             |

**_Body: None_**

#### II. Example Response: 失败示例（Not Found）

```js
{
    "code": -1,
    "data": {},
    "message": "no city found"
}
```

**_Status Code:_** 404

<br>

## Arcade（机厅）

getArcade/:arcadeID 的机厅 ID 与华立公布的机厅列表的机厅 ID 一致

如果你想获取华立更新的机厅，直接进行一次请求即可，将自动添加至数据库中

### 1. getArcade（获取机厅详情）

**_Endpoint:_**

```bash
Method: GET
Type:
URL: {{url}}/getArcade/:arcadeID
```

**_URL variables:_**

| Key      | Value | Description |
| -------- | ----- | ----------- |
| arcadeID | 0     |             |

**_More example Requests/Responses:_**

#### I. Example Request: 成功示例

**_Query:_**

| Key      | Value | Description |
| -------- | ----- | ----------- |
| arcadeID | 1001  |             |

**_Body: None_**

#### I. Example Response: 成功示例

```js
{
    "code": 0,
    "data": {
        "arcade_id": 1001,
        "arcade_name": "环游嘉年华天河店",
        "machine_count": 4,
        "address": "广东省广州市天河区天河路208号天河城6楼",
        "province": {
            "code": 44,
            "name": "广东省",
            "full_code": "440000000000"
        },
        "city": {
            "code": 4401,
            "name": "广州市",
            "province_code": 44,
            "full_code": "440100000000"
        }
    }
}
```

**_Status Code:_** 200

<br>

#### II. Example Request: 失败示例

**_Query:_**

| Key      | Value | Description |
| -------- | ----- | ----------- |
| arcadeID | 0     |             |

**_Body: None_**

#### II. Example Response: 失败示例

```js
{
    "code": -1,
    "message": "arcade not found"
}
```

**_Status Code:_** 500

<br>

### 2. getArcadeList（获取机厅列表）

**_Endpoint:_**

```bash
Method: GET
Type:
URL: {{url}}/getArcadeList
```

**_More example Requests/Responses:_**

#### I. Example Request: 成功示例

**_Body: None_**

#### I. Example Response: 成功示例

```js
{
    "code": 0,
    "data": [
        {
            "arcade_id": 1001,
            "arcade_name": "环游嘉年华天河店",
            "machine_count": 4,
            "address": "广东省广州市天河区天河路208号天河城6楼",
            "province": {
                "code": 44,
                "name": "广东省",
                "full_code": "440000000000"
            },
            "city": {
                "code": 4401,
                "name": "广州市",
                "province_code": 44,
                "full_code": "440100000000"
            }
        },
        {
            "arcade_id": 1002,
            "arcade_name": "环游嘉年华番禺易发店",
            "machine_count": 3,
            "address": "广东省广州市番禺区市桥街易发商业街新大新百货5楼",
            "province": {
                "code": 44,
                "name": "广东省",
                "full_code": "440000000000"
            },
            "city": {
                "code": 4401,
                "name": "广州市",
                "province_code": 44,
                "full_code": "440100000000"
            }
        },
        ...(1300+)
    ]
}
```

**_Status Code:_** 200

<br>

### 3. getArcadeList/city（获取某城市机厅列表）

**_Endpoint:_**

```bash
Method: GET
Type:
URL: {{url}}/getArcadeList/city/:cityCode
```

**_URL variables:_**

| Key      | Value | Description |
| -------- | ----- | ----------- |
| cityCode | 1101  |             |

**_More example Requests/Responses:_**

#### I. Example Request: 成功示例

**_Query:_**

| Key      | Value | Description |
| -------- | ----- | ----------- |
| cityCode | 1101  |             |

**_Body: None_**

#### I. Example Response: 成功示例

```js
{
    "code": 0,
    "data": [
        {
            "arcade_id": 1011,
            "arcade_name": "北京乐酷电玩",
            "machine_count": 2,
            "address": "北京市海淀区远大路一号世纪金源购物中心5楼",
            "province": {
                "code": 11,
                "name": "北京市",
                "full_code": "110000000000"
            },
            "city": {
                "code": 1101,
                "name": "市辖区",
                "province_code": 11,
                "full_code": "110100000000"
            }
        },
        {
            "arcade_id": 1018,
            "arcade_name": "风云再起北京西单店",
            "machine_count": 2,
            "address": "北京市西城区堂子胡同9号新一代商城六层",
            "province": {
                "code": 11,
                "name": "北京市",
                "full_code": "110000000000"
            },
            "city": {
                "code": 1101,
                "name": "市辖区",
                "province_code": 11,
                "full_code": "110100000000"
            }
        },
        ...
    ]
}
```

**_Status Code:_** 200

<br>

#### II. Example Request: 失败示例

**_Query:_**

| Key      | Value | Description |
| -------- | ----- | ----------- |
| cityCode | 11011 |             |

**_Body: None_**

#### II. Example Response: 失败示例

```js
{
    "code": -1,
    "data": [],
    "message": "no arcade found"
}
```

**_Status Code:_** 404

<br>

### 4. getArcadeList/province（获取某省份机厅列表）

**_Endpoint:_**

```bash
Method: GET
Type:
URL: {{url}}/getArcadeList/province/:provinceCode
```

**_URL variables:_**

| Key          | Value | Description |
| ------------ | ----- | ----------- |
| provinceCode | 0     |             |

**_More example Requests/Responses:_**

#### I. Example Request: 成功示例

**_Query:_**

| Key          | Value | Description |
| ------------ | ----- | ----------- |
| provinceCode | 11    |             |

**_Body: None_**

#### I. Example Response: 成功示例

```js
{
    "code": 0,
    "data": [
        {
            "arcade_id": 1011,
            "arcade_name": "北京乐酷电玩",
            "machine_count": 2,
            "address": "北京市海淀区远大路一号世纪金源购物中心5楼",
            "province": {
                "code": 11,
                "name": "北京市",
                "full_code": "110000000000"
            },
            "city": {
                "code": 1101,
                "name": "市辖区",
                "province_code": 11,
                "full_code": "110100000000"
            }
        },
        {
            "arcade_id": 1018,
            "arcade_name": "风云再起北京西单店",
            "machine_count": 2,
            "address": "北京市西城区堂子胡同9号新一代商城六层",
            "province": {
                "code": 11,
                "name": "北京市",
                "full_code": "110000000000"
            },
            "city": {
                "code": 1101,
                "name": "市辖区",
                "province_code": 11,
                "full_code": "110100000000"
            }
        },
        ...
    ]
}
```

**_Status Code:_** 200

<br>

#### II. Example Request: 失败示例

**_Query:_**

| Key          | Value | Description |
| ------------ | ----- | ----------- |
| provinceCode | 0     |             |

**_Body: None_**

#### II. Example Response: 失败示例

```js
{
    "code": -1,
    "data": [],
    "message": "no arcade found"
}
```

**_Status Code:_** 404

<br>

## Count（人数）

### 1. getCount（获取机厅人数）

**_Endpoint:_**

```bash
Method: GET
Type:
URL: {{url}}/getCount/:arcadeID
```

**_URL variables:_**

| Key      | Value | Description |
| -------- | ----- | ----------- |
| arcadeID | 1001  |             |

**_More example Requests/Responses:_**

#### I. Example Request: 成功示例

**_Query:_**

| Key      | Value | Description |
| -------- | ----- | ----------- |
| arcadeID | 1002  |             |

**_Body: None_**

#### I. Example Response: 成功示例

```js
{
    "code": 0,
    "data": [
        {
            "arcade_id": 1002,
            "count": 0,
            "update_timestamp": 1703606400000,
            "type": 0
        }
    ],
    "message": "success"
}
```

**_Status Code:_** 200

<br>

#### II. Example Request: 失败示例

**_Query:_**

| Key      | Value | Description |
| -------- | ----- | ----------- |
| arcadeID | 0     |             |

**_Body: None_**

#### II. Example Response: 失败示例

```js
{
    "code": -1,
    "message": "arcade not found"
}
```

**_Status Code:_** 500

<br>

### 2. logCount（记录机厅人数）

**_Endpoint:_**

```bash
Method: POST
Type: RAW
URL: {{url}}/logCount
```

**_Body:_**

```js
{
    "arcade_id": 1001,
    "count": 0,
    "type": 2,
    "token": "test"
}
```

**_More example Requests/Responses:_**

#### I. Example Request: 成功示例

**_Body:_**

```js
{
    "arcade_id": 1001,
    "count": 1,
    "type": 2,
    "token": "test"
}
```

#### I. Example Response: 成功示例

```js
{
    "code": 0
}
```

**_Status Code:_** 200

<br>

#### II. Example Request: 失败示例（token 错误）

**_Body:_**

```js
{
    "arcade_id": 1001,
    "count": 1,
    "type": 2,
    "token": "test2"
}
```

#### II. Example Response: 失败示例（token 错误）

```js
{
    "code": -1,
    "message": "invalid token"
}
```

**_Status Code:_** 400

<br>

#### III. Example Request: 失败示例（type=0）

**_Body:_**

```js
{
    "arcade_id": 1001,
    "count": 1,
    "type": 0,
    "token": "test"
}
```

#### III. Example Response: 失败示例（type=0）

```js
{
    "code": -1,
    "message": "I'm a teapot"
}
```

**_Status Code:_** 418

<br>

#### IV. Example Request: 失败示例（数量错误）

**_Body:_**

```js
{
    "arcade_id": 1001,
    "count": -1,
    "type": 2,
    "token": "test"
}
```

#### IV. Example Response: 失败示例（数量错误）

```js
{
    "code": -1,
    "message": "invalid count"
}
```

**_Status Code:_** 400

<br>

## Admin（管理员用）

这里的 token 需要使用 TOTP 生成动态口令

### 1. createCustomArcade（创建机厅）

**_Endpoint:_**

```bash
Method: POST
Type: RAW
URL: {{url}}/createCustomArcade
```

**_Body:_**

```js
{
    "arcade_name": "机厅名称",
    "address": "机厅地址",
    "machine_count": 1, // 机台数量
    "province_code": 11, // 省份 ID
    "city_code": 1101, // 城市ID
    "token": "123456" // TOTP 动态密码
}
```

**_More example Requests/Responses:_**

#### I. Example Request: 示例

**_Body:_**

```js
{
    "arcade_name": "机厅名称",
    "address": "机厅地址",
    "machine_count": 1, // 机台数量
    "province_code": 11, // 省份 ID
    "city_code": 1101, // 城市ID
    "token": "123456" // TOTP 动态密码
}
```

#### I. Example Response: 示例

```js
{
    "code": -1,
    "message": "invalid token"
}
```

**_Status Code:_** 400

<br>

### 2. deleteArcade（删除机厅）

**_Endpoint:_**

```bash
Method: POST
Type: RAW
URL: {{url}}/deleteArcade
```

**_Body:_**

```js
{
    "arcade_id": 1101, // 机厅 ID
    "token": "123456" // TOTP 动态密码
}
```

### 3. updateArcade（从华立更新机厅）

**_Endpoint:_**

```bash
Method: POST
Type: RAW
URL: {{url}}/updateArcade
```

**_Body:_**

```js
{
    "token": "123456" // TOTP 动态密码
}
```

### 4. generateToken（生成 token）

**_Endpoint:_**

```bash
Method: POST
Type: RAW
URL: {{url}}/generateToken
```

**_Body:_**

```js
{
    "token": "123456", // TOTP 动态密码
    "remark": "备注"
}
```

---

[Back to top](#go-maimai-player-counter)

> Generated by [docgen](https://github.com/thedevsaddam/docgen)
