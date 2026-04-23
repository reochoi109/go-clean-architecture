# Go Clean Architecture Template

Go로 클린 아키텍처를 적용한 백엔드 API 템플릿입니다. (Gin + RDBMS)

## Features

- Clean Architecture 레이어 분리 (`domain` ↔ `usecase` ↔ `repository/handler`)
- 환경변수 기반 설정 로딩 (`.env`, `godotenv`)
- Gin 기반 REST API + 미들웨어 예시
- `database/sql` 기반 MySQL 연결

## Requirements

- Go `1.25+` (see `go.mod`)
- MySQL (예: 8.x)

## Project Structure

```Plaintext
├── cmd/             # 프로그램의 시작점 (Entry Point)
├── internal/        # 프로젝트 내부 비즈니스 로직 (외부 접근 제한)
│   ├── handler/     # HTTP/API 처리 (Gin)
│   ├── usecase/     # 핵심 비즈니스 로직
│   ├── repository/  # 데이터 저장소 구현체 (MySQL)
│   ├── router/      # 라우팅 및 미들웨어
│   └── infra/       # DB 연결 및 외부 설정
├── domain/          # 기술에 의존하지 않는 핵심 인터페이스/모델
└── pkg/             # 외부에서 재사용 가능한 공통 유틸리티
```

## Getting Started

### 1) 환경변수 설정

```bash
cp .env.example .env
```

필수 환경변수:

- `DB_TYPE` (default: `mysql`)
- `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`

### 2) 의존성 설치 및 서버 실행

```bash
# 1. install dependency
go mod download

# 2. run server
go run ./cmd/main.go
```

기본 포트는 `PORT` 환경변수(기본값 `8080`)를 사용합니다.

## API

- `GET /v1/ping` → `pong`
- `GET /v1/author/:id` → Author + Articles
  - Header: `Authorization: secret`

예시:

```bash
curl -s http://localhost:8080/v1/ping
curl -s -H 'Authorization: secret' http://localhost:8080/v1/author/1
```

## Stack

- Web Framework: Gin
- Database: MySQL (database/sql, go-sql-driver/mysql)

## Contributing

- 이슈/PR 환영합니다.
