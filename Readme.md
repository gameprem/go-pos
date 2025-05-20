# go-pos

ระบบตัวอย่าง RESTful API สำหรับจัดการผู้ใช้ (User Management) ด้วย Go (Golang), [Fiber](https://gofiber.io/), JWT, และ PostgreSQL  
**มี Swagger UI สำหรับดูเอกสาร API**

## คุณสมบัติ

- สมัครสมาชิก / ล็อกอิน (JWT Auth)
- ดูข้อมูลโปรไฟล์ผู้ใช้ (ต้องล็อกอิน)
- ใช้ GORM เชื่อมต่อ PostgreSQL
- มี Swagger UI (`/docs`) สำหรับดูและทดสอบ API

## โครงสร้างโปรเจกต์

```
main.go                  // Entry point ของแอป
internal/
  database/              // การเชื่อมต่อและ migrate database
  handler/               // ฟังก์ชันจัดการ endpoint (auth, profile)
  middleware/            // JWT Middleware
  model/                 // โครงสร้างข้อมูล (User, UserInfo)
  router/                // กำหนด route หลัก
docs/
  swagger.json|yaml      // OpenAPI spec สำหรับ Swagger UI
docker-compose.yml       // สำหรับรัน PostgreSQL ด้วย Docker
go.mod, go.sum           // Go modules
```

## การติดตั้งและรัน

### 1. Clone และติดตั้ง dependencies

```bash
git clone https://github.com/gameprem/go-pos.git
cd go-pos
go mod tidy
```

### 2. รัน PostgreSQL ด้วย Docker

```bash
docker-compose up -d
```

- Database จะรันที่ `localhost:5432`
- ค่า default: user=`root`, password=`password`, db=`mydb`

### 3. รันแอป

```bash
go run main.go
```

- Server จะรันที่ `http://localhost:3030`
- Swagger UI: [http://localhost:3030/docs](http://localhost:3030/docs)

## เอกสาร API

- Swagger UI: [http://localhost:3030/docs](http://localhost:3030/docs)
- สเปค OpenAPI: `docs/swagger.json` หรือ `docs/swagger.yaml`

## License

Apache 2.0
