# system-design-lab-fe

React frontend cho [System Design Lab](../README.md).

## Yêu cầu

- Node.js 18+
- npm hoặc pnpm
- Backend đang chạy tại `http://localhost:8080`

## Chạy nhanh

```bash
npm install
npm run dev
```

App chạy tại `http://localhost:5173`.

Proxy `/api` → `http://localhost:8080` đã được cấu hình sẵn, không cần thêm gì.

## Biến môi trường

Tạo file `.env.local` nếu cần override URL mặc định:

```env
VITE_API_BASE_URL=/api/v1
```

Mặc định đã trỏ đúng qua Vite proxy, chỉ cần đổi khi deploy production.

## Lệnh

```bash
# Dev server (hot reload)
npm run dev

# Build production
npm run build

# Preview production build
npm run preview

# Type check
npm run tsc
```

## Các trang

| Path | Mô tả |
|------|-------|
| `/` | Landing page |
| `/quests` | Danh sách scenario |
| `/quests/:id/begin` | Bắt đầu quest (nhập username) |
| `/quests/:id/leaderboard` | Leaderboard của scenario |
| `/play/:sessionId` | Gameplay — chọn và submit |
| `/play/:sessionId/summary` | Tổng kết sau khi hoàn thành |
| `/profile/progress` | Tiến độ cá nhân |
