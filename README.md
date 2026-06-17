# System Design Lab

Nền tảng học system design thông qua việc đưa ra quyết định thực tế. Thay vì vẽ diagram, bạn phải chọn — và mỗi lựa chọn đều có hậu quả.

## Vấn đề

Dev junior không biết bắt đầu từ đâu khi thiết kế hệ thống. Các công cụ hiện tại chỉ cho vẽ tự do, không có feedback, không giải thích tại sao đúng hay sai.

## Giải pháp

System Design Lab đưa bạn vào các kịch bản thực tế (Chat App, URL Shortener, E-commerce...) và ép bạn đi theo một flow có cấu trúc:

```
Câu hỏi → Lựa chọn → Feedback + Metrics → Câu hỏi tiếp theo
```

Mỗi quyết định ảnh hưởng trực tiếp đến chỉ số hệ thống. Sai → thấy system degraded. Đúng → thấy metrics cải thiện.

## Tính năng

**Gameplay**
- Scenario-based: mỗi bài là một hệ thống thực tế với 5–7 bước quyết định
- Metrics thay đổi real-time theo từng lựa chọn (latency, cost, scalability)
- Feedback ngay lập tức sau mỗi bước
- Giải thích 3 cấp độ: Beginner / Intermediate / Advanced
- Gợi ý (hint) cho những ai cần thêm context

**Session**
- Chế độ Normal: học theo nhịp của mình
- Chế độ Interview: có giới hạn thời gian, simulate buổi phỏng vấn
- Abandon và Retry quest bất kỳ lúc nào

**Progress & Leaderboard**
- Theo dõi tiến độ qua từng scenario
- Leaderboard top 10 cho mỗi scenario
- Lịch sử điểm số và số lần hoàn thành

## Tech Stack

| Layer | Stack |
|-------|-------|
| Backend | Go, PostgreSQL, sqlc, pgx/v5 |
| Frontend | React, TypeScript, Vite, TanStack Query |

## Cấu trúc repo

```
system-design-lab/
├── system-design-lab-api/   # Go backend
└── system-design-lab-fe/    # React frontend
```

Xem hướng dẫn chạy trong từng thư mục:
- [Backend →](./system-design-lab-api/README.md)
- [Frontend →](./system-design-lab-fe/README.md)

## Screenshots

> Coming soon

## Roadmap

- [ ] AI Coach — hỏi ngược user, giải thích sâu hơn theo từng tình huống
- [ ] Visual Architecture — sơ đồ kiến trúc cập nhật theo từng quyết định
- [ ] Multiplayer — học nhóm, so sánh quyết định
- [ ] Interview Mode nâng cao — timer, pressure scenarios
