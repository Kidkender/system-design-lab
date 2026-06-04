-- =============================================================================
-- Seed: Chat App Scenario (Complete)
-- Scenario : ca000000-0000-0000-0000-000000000000
-- Steps    : 5 steps
-- Safe to re-run: ON CONFLICT (id) DO NOTHING on every insert
-- =============================================================================

-- -----------------------------------------------------------------------------
-- Scenario
-- -----------------------------------------------------------------------------
INSERT INTO scenarios (id, title, description, difficulty)
VALUES (
    'ca000000-0000-0000-0000-000000000000',
    'Chat App',
    'Design a real-time chat application that can scale to 100k concurrent users and handle 1 million messages per day.',
    'medium'
) ON CONFLICT (id) DO NOTHING;

-- -----------------------------------------------------------------------------
-- Steps (all 5, inserted before FK update)
-- -----------------------------------------------------------------------------
INSERT INTO steps (id, scenario_id, question, context, order_index) VALUES
(
    'ca000000-0000-0000-0000-000000000001',
    'ca000000-0000-0000-0000-000000000000',
    'You need to plan your Chat App to support 100k concurrent users. What is your first move?',
    'You are the sole engineer. You have 3 months to launch. Estimate the scale before building.',
    1
),
(
    'ca000000-0000-0000-0000-000000000002',
    'ca000000-0000-0000-0000-000000000000',
    'Users send 1 million messages per day. How do you store messages reliably without overloading the database?',
    'Each message is ~1KB. Direct writes to Postgres could saturate connection pools under load.',
    2
),
(
    'ca000000-0000-0000-0000-000000000003',
    'ca000000-0000-0000-0000-000000000000',
    'Message history is heavily read — users scroll back through chat. How do you handle read traffic?',
    'Reads outnumber writes 10:1. Your Postgres read IOPS are already at 70%.',
    3
),
(
    'ca000000-0000-0000-0000-000000000004',
    'ca000000-0000-0000-0000-000000000000',
    'Users expect new messages to appear instantly without refreshing. How do you deliver real-time messages?',
    'You need to push messages to connected clients. 100k persistent connections.',
    4
),
(
    'ca000000-0000-0000-0000-000000000005',
    'ca000000-0000-0000-0000-000000000000',
    'The app has grown to 500k users and a single chat server is struggling. How do you scale?',
    'CPU on your chat server is at 90%. You have budget. Vertical vs horizontal?',
    5
)
ON CONFLICT (id) DO NOTHING;

-- Set start_step_id
UPDATE scenarios
SET start_step_id = 'ca000000-0000-0000-0000-000000000001'
WHERE id = 'ca000000-0000-0000-0000-000000000000';

-- -----------------------------------------------------------------------------
-- Choices (correct = 1, wrong = 2 per step)
-- Pattern: ca000000-0000-0000-<step>-00000000000<1|2>
-- -----------------------------------------------------------------------------
INSERT INTO choices (id, step_id, label, next_step_id, is_correct) VALUES
-- Step 1
(
    'ca000000-0000-0000-0001-000000000001',
    'ca000000-0000-0000-0000-000000000001',
    'Estimate load: 100k concurrent users × avg message rate → choose DB, queue, and infra size accordingly',
    'ca000000-0000-0000-0000-000000000002',
    true
),
(
    'ca000000-0000-0000-0001-000000000002',
    'ca000000-0000-0000-0000-000000000001',
    'Skip planning and start building — we will scale when we need to',
    'ca000000-0000-0000-0000-000000000002',
    false
),
-- Step 2
(
    'ca000000-0000-0000-0002-000000000001',
    'ca000000-0000-0000-0000-000000000002',
    'Use a message queue (Kafka / RabbitMQ) — producers push messages async, consumers write to DB',
    'ca000000-0000-0000-0000-000000000003',
    true
),
(
    'ca000000-0000-0000-0002-000000000002',
    'ca000000-0000-0000-0000-000000000002',
    'Write each message directly to the database on every API request',
    'ca000000-0000-0000-0000-000000000003',
    false
),
-- Step 3
(
    'ca000000-0000-0000-0003-000000000001',
    'ca000000-0000-0000-0000-000000000003',
    'Cache recent messages in Redis — serve most reads from memory, fall back to DB only on cache miss',
    'ca000000-0000-0000-0000-000000000004',
    true
),
(
    'ca000000-0000-0000-0003-000000000002',
    'ca000000-0000-0000-0000-000000000003',
    'Query the database directly on every request for message history',
    'ca000000-0000-0000-0000-000000000004',
    false
),
-- Step 4
(
    'ca000000-0000-0000-0004-000000000001',
    'ca000000-0000-0000-0000-000000000004',
    'Use WebSockets — establish a persistent bidirectional connection and push messages to clients in real-time',
    'ca000000-0000-0000-0000-000000000005',
    true
),
(
    'ca000000-0000-0000-0004-000000000002',
    'ca000000-0000-0000-0000-000000000004',
    'Use HTTP polling — clients ask the server for new messages every 2 seconds',
    'ca000000-0000-0000-0000-000000000005',
    false
),
-- Step 5 (last — next_step_id = NULL)
(
    'ca000000-0000-0000-0005-000000000001',
    'ca000000-0000-0000-0000-000000000005',
    'Horizontal scaling — add more chat server instances behind a load balancer, use Redis pub/sub to sync messages between nodes',
    NULL,
    true
),
(
    'ca000000-0000-0000-0005-000000000002',
    'ca000000-0000-0000-0000-000000000005',
    'Upgrade the existing server to a bigger machine (vertical scaling)',
    NULL,
    false
)
ON CONFLICT (id) DO NOTHING;

-- -----------------------------------------------------------------------------
-- Impacts
-- latency (ms), cost (1-10), scalability (1-10)
-- -----------------------------------------------------------------------------
INSERT INTO impacts (id, choice_id, metric, type, value) VALUES
-- Step 1 correct: planning cuts future latency, no cost yet
('ca000000-0001-0001-0001-000000000001', 'ca000000-0000-0000-0001-000000000001', 'latency',     'add', -5),
('ca000000-0001-0001-0002-000000000001', 'ca000000-0000-0000-0001-000000000001', 'scalability', 'add',  2),
-- Step 1 wrong: skip planning hurts scalability
('ca000000-0001-0002-0001-000000000002', 'ca000000-0000-0000-0001-000000000002', 'cost',        'add',  1),
('ca000000-0001-0002-0002-000000000002', 'ca000000-0000-0000-0001-000000000002', 'scalability', 'add', -2),
-- Step 2 correct: queue absorbs write burst
('ca000000-0002-0001-0001-000000000001', 'ca000000-0000-0000-0002-000000000001', 'latency',     'add', -20),
('ca000000-0002-0001-0002-000000000001', 'ca000000-0000-0000-0002-000000000001', 'scalability', 'add',  3),
('ca000000-0002-0001-0003-000000000001', 'ca000000-0000-0000-0002-000000000001', 'cost',        'add',  1),
-- Step 2 wrong: direct DB write spikes latency
('ca000000-0002-0002-0001-000000000002', 'ca000000-0000-0000-0002-000000000002', 'latency',     'add',  40),
('ca000000-0002-0002-0002-000000000002', 'ca000000-0000-0000-0002-000000000002', 'scalability', 'add', -2),
-- Step 3 correct: Redis cache slashes read latency
('ca000000-0003-0001-0001-000000000001', 'ca000000-0000-0000-0003-000000000001', 'latency',     'add', -30),
('ca000000-0003-0001-0002-000000000001', 'ca000000-0000-0000-0003-000000000001', 'scalability', 'add',  2),
('ca000000-0003-0001-0003-000000000001', 'ca000000-0000-0000-0003-000000000001', 'cost',        'add',  1),
-- Step 3 wrong: DB reads saturate IOPS
('ca000000-0003-0002-0001-000000000002', 'ca000000-0000-0000-0003-000000000002', 'latency',     'add',  50),
('ca000000-0003-0002-0002-000000000002', 'ca000000-0000-0000-0003-000000000002', 'scalability', 'add', -2),
-- Step 4 correct: WebSocket = low-latency push
('ca000000-0004-0001-0001-000000000001', 'ca000000-0000-0000-0004-000000000001', 'latency',     'add', -15),
('ca000000-0004-0001-0002-000000000001', 'ca000000-0000-0000-0004-000000000001', 'scalability', 'add',  2),
('ca000000-0004-0001-0003-000000000001', 'ca000000-0000-0000-0004-000000000001', 'cost',        'add',  1),
-- Step 4 wrong: polling floods server
('ca000000-0004-0002-0001-000000000002', 'ca000000-0000-0000-0004-000000000002', 'latency',     'add',  30),
('ca000000-0004-0002-0002-000000000002', 'ca000000-0000-0000-0004-000000000002', 'cost',        'add',  2),
-- Step 5 correct: horizontal scale efficient
('ca000000-0005-0001-0001-000000000001', 'ca000000-0000-0000-0005-000000000001', 'scalability', 'add',  3),
('ca000000-0005-0001-0002-000000000001', 'ca000000-0000-0000-0005-000000000001', 'cost',        'add',  2),
-- Step 5 wrong: vertical scaling hits ceiling
('ca000000-0005-0002-0001-000000000002', 'ca000000-0000-0000-0005-000000000002', 'latency',     'add',  15),
('ca000000-0005-0002-0002-000000000002', 'ca000000-0000-0000-0005-000000000002', 'scalability', 'add', -1)
ON CONFLICT (id) DO NOTHING;

-- -----------------------------------------------------------------------------
-- Explanations (3 levels per choice)
-- -----------------------------------------------------------------------------
-- UUID scheme: ea000000-SSSS-CCCC-LLLL-000000000001
--   S = step (0001-0005), C = choice (0001|0002), L = level (0001=beginner,0002=intermediate,0003=advanced)
INSERT INTO explanations (id, choice_id, level, content) VALUES
-- Step 1 correct
('ea000000-0001-0001-0001-000000000001', 'ca000000-0000-0000-0001-000000000001', 'beginner',
 'Giống như trước khi xây nhà, bạn phải tính xem bao nhiêu người sẽ sống ở đó. Biết trước quy mô giúp bạn chọn đúng "vật liệu xây dựng" cho hệ thống.'),
('ea000000-0001-0001-0002-000000000001', 'ca000000-0000-0000-0001-000000000001', 'intermediate',
 'Load estimation giúp chọn đúng tech stack: 100k concurrent users → cần async processing, caching layer, và horizontal scaling ngay từ đầu thay vì refactor sau.'),
('ea000000-0001-0001-0003-000000000001', 'ca000000-0000-0000-0001-000000000001', 'advanced',
 '100k concurrent × 1 msg/min ≈ 1,700 msg/s. Với message size 1KB, write bandwidth ~1.7MB/s. Con số này quyết định partition count của Kafka, replica count của DB, và memory footprint của Redis.'),
-- Step 1 wrong
('ea000000-0001-0002-0001-000000000001', 'ca000000-0000-0000-0001-000000000002', 'beginner',
 'Bắt đầu xây nhà mà không tính số người ở là sai. Hệ thống cũng vậy — không estimate sẽ dẫn đến sập dưới load thực.'),
('ea000000-0001-0002-0002-000000000001', 'ca000000-0000-0000-0001-000000000002', 'intermediate',
 '"Scale when needed" nghe có vẻ linh hoạt nhưng thực tế là tech debt lớn. Refactor một hệ thống đang chạy dưới production load nguy hiểm hơn build đúng từ đầu.'),
('ea000000-0001-0002-0003-000000000001', 'ca000000-0000-0000-0001-000000000002', 'advanced',
 'Premature optimization là xấu nhưng premature under-planning còn nguy hiểm hơn. Không có load estimate, bạn không set được resource limits đúng → OOM kills và cascading failures dưới production load.'),
-- Step 2 correct
('ea000000-0002-0001-0001-000000000001', 'ca000000-0000-0000-0002-000000000001', 'beginner',
 'Queue giống như xếp hàng tại quầy — thay vì tất cả khách hàng ùa vào cùng lúc, họ xếp hàng và được phục vụ lần lượt. DB không bị quá tải.'),
('ea000000-0002-0001-0002-000000000001', 'ca000000-0000-0000-0002-000000000001', 'intermediate',
 'Message queue tách producer (API server) khỏi consumer (DB writer). API trả về 200 ngay lập tức, queue đảm bảo message được ghi vào DB async. Throughput tăng, P99 latency giảm.'),
('ea000000-0002-0001-0003-000000000001', 'ca000000-0000-0000-0002-000000000001', 'advanced',
 'Kafka với replication factor 3 đảm bảo at-least-once delivery. Consumer group cho horizontal scale DB writers. Retention policy cho phép replay khi consumer fails. Throughput Kafka ~1M msg/s trên commodity hardware.'),
-- Step 2 wrong
('ea000000-0002-0002-0001-000000000001', 'ca000000-0000-0000-0002-000000000002', 'beginner',
 'Mỗi lần gửi tin nhắn là một lần "gõ cửa" database. 1 triệu tin nhắn/ngày = 1 triệu lần gõ cửa. Database sẽ không chịu được.'),
('ea000000-0002-0002-0002-000000000001', 'ca000000-0000-0000-0002-000000000002', 'intermediate',
 'Với 100k concurrent users, direct DB write dễ saturate connection pool (thường giới hạn 100-200 connections). Query queue buildup → timeout → cascade failure.'),
('ea000000-0002-0002-0003-000000000001', 'ca000000-0000-0000-0002-000000000002', 'advanced',
 'Write amplification + connection pool exhaustion: với 100k concurrent users và mỗi write mất 5ms, cần 500k connections/giây — không DB nào chịu được. P99 latency tăng exponentially do queue buildup tại DB layer.'),
-- Step 3 correct
('ea000000-0003-0001-0001-000000000001', 'ca000000-0000-0000-0003-000000000001', 'beginner',
 'Redis giống như bộ nhớ ngắn hạn — đọc từ đây nhanh hơn 100x so với đọc từ database. Tin nhắn gần đây nhất được lưu trong Redis, chỉ tin nhắn cũ mới cần tra cứu DB.'),
('ea000000-0003-0001-0002-000000000001', 'ca000000-0000-0000-0003-000000000001', 'intermediate',
 'Cache-aside pattern: read từ Redis trước, nếu miss thì query DB và lưu kết quả vào Redis. Hot conversations chiếm 90% reads (Pareto principle) — cache hit rate thường đạt 95%+ trong thực tế.'),
('ea000000-0003-0001-0003-000000000001', 'ca000000-0000-0000-0003-000000000001', 'advanced',
 'Redis throughput >1M ops/s với sub-millisecond latency. Cache hot conversations với LRU eviction. Invalidation strategy: append-on-write với max list size cap, tránh thundering herd khi cache expires.'),
-- Step 3 wrong
('ea000000-0003-0002-0001-000000000001', 'ca000000-0000-0000-0003-000000000002', 'beginner',
 'Mỗi lần scroll lên xem tin nhắn cũ là một lần "hỏi" database. 100k người scroll cùng lúc = database trả lời 100k câu hỏi một lúc. Chậm và tốn kém.'),
('ea000000-0003-0002-0002-000000000001', 'ca000000-0000-0000-0003-000000000002', 'intermediate',
 'Reads outnumber writes 10:1 trong chat apps. Direct DB reads saturate IOPS và buffer pool. Read replicas giúp nhưng thêm replication lag và complexity. Cache là giải pháp đúng.'),
('ea000000-0003-0002-0003-000000000001', 'ca000000-0000-0000-0003-000000000002', 'advanced',
 'DB reads at scale: mỗi query với ORDER BY + LIMIT trên conversation_id cần index scan. Với 100M messages và 100k concurrent reads, DB IOPS bị saturate. Read replica giúp nhưng thêm replication lag.'),
-- Step 4 correct
('ea000000-0004-0001-0001-000000000001', 'ca000000-0000-0000-0004-000000000001', 'beginner',
 'WebSocket giống như một đường dây điện thoại luôn bật — server có thể "gọi" cho client bất cứ lúc nào có tin nhắn mới, không cần client hỏi trước.'),
('ea000000-0004-0001-0002-000000000001', 'ca000000-0000-0000-0004-000000000001', 'intermediate',
 'WebSocket duy trì persistent TCP connection, loại bỏ HTTP overhead cho mỗi message. Server push thay vì client pull — latency giảm từ 2000ms (polling interval) xuống <50ms.'),
('ea000000-0004-0001-0003-000000000001', 'ca000000-0000-0000-0004-000000000001', 'advanced',
 'WebSocket với 100k connections: dùng event-driven server (Go goroutines, Node.js) thay vì thread-per-connection. Sticky session hoặc Redis pub/sub để route messages qua multiple chat server instances.'),
-- Step 4 wrong
('ea000000-0004-0002-0001-000000000001', 'ca000000-0000-0000-0004-000000000002', 'beginner',
 'Polling giống như cứ 2 giây lại hỏi "có tin nhắn mới không?" — dù không có gì mới. Vừa chậm, vừa lãng phí tài nguyên server.'),
('ea000000-0004-0002-0002-000000000001', 'ca000000-0000-0000-0004-000000000002', 'intermediate',
 '100k clients × 1 request/2s = 50,000 HTTP requests/giây chỉ để check tin nhắn mới. Phần lớn là empty response — pure overhead. CPU và bandwidth bị lãng phí.'),
('ea000000-0004-0002-0003-000000000001', 'ca000000-0000-0000-0004-000000000002', 'advanced',
 'Polling tạo thundering herd problem khi tất cả clients sync interval. Long polling cải thiện nhưng vẫn có HTTP handshake overhead mỗi response. WebSocket là giải pháp đúng cho bidirectional real-time communication.'),
-- Step 5 correct
('ea000000-0005-0001-0001-000000000001', 'ca000000-0000-0000-0005-000000000001', 'beginner',
 'Thay vì mua một xe tải lớn hơn (vertical), bạn thuê thêm nhiều xe tải nhỏ (horizontal). Nếu một xe hỏng, các xe khác vẫn chạy. Rẻ hơn và an toàn hơn.'),
('ea000000-0005-0001-0002-000000000001', 'ca000000-0000-0000-0005-000000000001', 'intermediate',
 'Horizontal scaling với load balancer phân tán traffic. Redis pub/sub làm message bus giữa các instances. Mỗi instance handle ~50k connections thay vì 500k trên một máy.'),
('ea000000-0005-0001-0003-000000000001', 'ca000000-0000-0000-0005-000000000001', 'advanced',
 'Consistent hashing để route WebSocket connections đến đúng server instance. Circuit breaker đảm bảo traffic không route đến unhealthy instances. Cost-effective: n × small instances rẻ hơn 1 × large instance (economies of commodity scale).'),
-- Step 5 wrong
('ea000000-0005-0002-0001-000000000001', 'ca000000-0000-0000-0005-000000000002', 'beginner',
 'Mua máy mạnh hơn có giới hạn — không máy nào xử lý được vô hạn người dùng. Và nếu máy đó hỏng, toàn bộ hệ thống sập.'),
('ea000000-0005-0002-0002-000000000001', 'ca000000-0000-0000-0005-000000000002', 'intermediate',
 'Vertical scaling có ceiling cứng (max RAM/CPU của hardware). Downtime khi upgrade. Single point of failure. Không autoscale được. Cloud pricing: reserved large instance đắt hơn nhiều small instances.'),
('ea000000-0005-0002-0003-000000000001', 'ca000000-0000-0000-0005-000000000002', 'advanced',
 'Amdahl''s Law: speedup từ vertical scaling bị giới hạn bởi phần serial của code. Moore''s Law chậm lại — CPU frequency không tăng nhiều. Cloud pricing: 1 × 96-core đắt hơn 8 × 12-core và kém flexible hơn.')
ON CONFLICT (id) DO NOTHING;

-- -----------------------------------------------------------------------------
-- Consequences (flags set bởi choices)
-- -----------------------------------------------------------------------------
INSERT INTO consequences (id, choice_id, type, keys, value) VALUES
('ac000000-0002-0001-0000-000000000001', 'ca000000-0000-0000-0002-000000000001', 'flag', 'has_queue',     true),
('ac000000-0003-0001-0000-000000000001', 'ca000000-0000-0000-0003-000000000001', 'flag', 'has_cache',     true),
('ac000000-0004-0001-0000-000000000001', 'ca000000-0000-0000-0004-000000000001', 'flag', 'has_websocket', true)
ON CONFLICT (id) DO NOTHING;
