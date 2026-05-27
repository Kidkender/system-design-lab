-- =============================================================================
-- Seed: Chat App Scenario
-- Scenario ID : ca000000-0000-0000-0000-000000000000
-- Steps       : 5 steps (scale → storage → cache → realtime → scaling)
-- Safe to re-run: ON CONFLICT (id) DO NOTHING on every insert
-- =============================================================================
-- -----------------------------------------------------------------------------
-- Seed user (dùng để test StartSession)
-- -----------------------------------------------------------------------------

INSERT INTO users (id, username, email)
VALUES ('seed0000-0000-0000-0000-000000000001', 'seed_user', 'seed@example.com');
-- -----------------------------------------------------------------------------
-- Scenario
-- -----------------------------------------------------------------------------

INSERT INTO scenarios (id, title, description, difficulty)
VALUES (
        'ca000000-0000-0000-0000-000000000000',
        'Chat App',
        'Design a real-time chat application that can scale to 100k concurrent users and handle 1 million messages per day.',
        'medium'
    );
-- -----------------------------------------------------------------------------
-- Steps (insert trước UPDATE start_step_id do FK)
-- -----------------------------------------------------------------------------

INSERT INTO steps (id, scenario_id, question, context, order_index)
VALUES (
        'ca000000-0000-0000-0000-000000000005',
        'ca000000-0000-0000-0000-000000000000',
        'The app has grown to 500k users and a single chat server is struggling. How do you scale?',
        'CPU on your chat server is at 90%. You have budget. How do you handle this growth?',
        5
    );
-- Set start_step_id after steps are inserted

UPDATE scenarios
SET start_step_id = 'ca000000-0000-0000-0000-000000000001'
WHERE id = 'ca000000-0000-0000-0000-000000000000';
-- -----------------------------------------------------------------------------
-- Choices
-- UUID scheme: ca000000-0000-0000-<step>-00000000000<choice_idx>
--   step 1 = 0001, step 2 = 0002, ...
--   correct = 1, wrong = 2
-- -----------------------------------------------------------------------------
-- Step 1: Scale planning

INSERT INTO choices (id, step_id, label, next_step_id, is_correct)
VALUES (
        'ca000000-0000-0000-0001-000000000002',
        'ca000000-0000-0000-0000-000000000001',
        'Skip planning and start building — we will scale when we need to',
        'ca000000-0000-0000-0000-000000000002',
        false
    );
-- Step 2: Message storage

INSERT INTO choices (id, step_id, label, next_step_id, is_correct)
VALUES (
        'ca000000-0000-0000-0002-000000000002',
        'ca000000-0000-0000-0000-000000000002',
        'Write each message directly to the database on every API request',
        'ca000000-0000-0000-0000-000000000003',
        false
    );
-- Step 3: Read traffic / cache

INSERT INTO choices (id, step_id, label, next_step_id, is_correct)
VALUES (
        'ca000000-0000-0000-0003-000000000002',
        'ca000000-0000-0000-0000-000000000003',
        'Query the database directly on every request for message history',
        'ca000000-0000-0000-0000-000000000004',
        false
    );
-- Step 4: Real-time delivery

INSERT INTO choices (id, step_id, label, next_step_id, is_correct)
VALUES (
        'ca000000-0000-0000-0004-000000000002',
        'ca000000-0000-0000-0000-000000000004',
        'Use HTTP polling — clients ask the server for new messages every 2 seconds',
        'ca000000-0000-0000-0000-000000000005',
        false
    );
-- Step 5: Scaling strategy (last step — next_step_id = NULL = end)

INSERT INTO choices (id, step_id, label, next_step_id, is_correct)
VALUES (
        'ca000000-0000-0000-0005-000000000002',
        'ca000000-0000-0000-0000-000000000005',
        'Upgrade the existing server to a bigger machine (vertical scaling)',
        NULL,
        false
    );
-- -----------------------------------------------------------------------------
-- Impacts
-- Metric scale: latency (ms delta), cost (1-10), scalability (1-10)
-- -----------------------------------------------------------------------------
-- Step 1 correct: planning improves scalability, no cost yet

INSERT INTO impacts (id, choice_id, metric, type, value)
VALUES ('ca000000-0001-0001-0001-000000000002', 'ca000000-0000-0000-0001-000000000001', 'latency',     'add', -5);
-- Step 1 wrong: no planning hurts future scalability

INSERT INTO impacts (id, choice_id, metric, type, value)
VALUES ('ca000000-0001-0002-0001-000000000002', 'ca000000-0000-0000-0001-000000000002', 'cost',        'add',  1);
-- Step 2 correct: queue reduces write latency, boosts scalability, small infra cost

INSERT INTO impacts (id, choice_id, metric, type, value)
VALUES ('ca000000-0002-0001-0001-000000000003', 'ca000000-0000-0000-0002-000000000001', 'cost',        'add',   1);
-- Step 2 wrong: direct DB write spikes latency, hurts scalability

INSERT INTO impacts (id, choice_id, metric, type, value)
VALUES ('ca000000-0002-0002-0001-000000000002', 'ca000000-0000-0000-0002-000000000002', 'scalability', 'add', -2);
-- Step 3 correct: Redis cache dramatically lowers read latency

INSERT INTO impacts (id, choice_id, metric, type, value)
VALUES ('ca000000-0003-0001-0001-000000000003', 'ca000000-0000-0000-0003-000000000001', 'cost',        'add',   1);
-- Step 3 wrong: DB query for every read is slow and unscalable

INSERT INTO impacts (id, choice_id, metric, type, value)
VALUES ('ca000000-0003-0002-0001-000000000002', 'ca000000-0000-0000-0003-000000000002', 'scalability', 'add', -2);
-- Step 4 correct: WebSocket = low-latency push, efficient

INSERT INTO impacts (id, choice_id, metric, type, value)
VALUES ('ca000000-0004-0001-0001-000000000003', 'ca000000-0000-0000-0004-000000000001', 'cost',        'add',   1);
-- Step 4 wrong: polling floods server with unnecessary requests

INSERT INTO impacts (id, choice_id, metric, type, value)
VALUES ('ca000000-0004-0002-0001-000000000003', 'ca000000-0000-0000-0004-000000000002', 'cost',        'add',  2);
-- Step 5 correct: horizontal scale is efficient and resilient

INSERT INTO impacts (id, choice_id, metric, type, value)
VALUES ('ca000000-0005-0001-0001-000000000003', 'ca000000-0000-0000-0005-000000000001', 'cost',        'add',   2);
-- Step 5 wrong: vertical scaling is expensive and hits a ceiling

INSERT INTO impacts (id, choice_id, metric, type, value)
VALUES ('ca000000-0005-0002-0001-000000000003', 'ca000000-0000-0000-0005-000000000002', 'latency',     'add', 15);
-- -----------------------------------------------------------------------------
-- Explanations (3 levels per choice)
-- -----------------------------------------------------------------------------
-- Step 1 correct

INSERT INTO explanations (id, choice_id, level, content)
VALUES ('caex0001-0001-0000-0000-000000000003', 'ca000000-0000-0000-0001-000000000001', 'advanced',
     '100k concurrent users × 1 message/min ≈ 1,700 msg/s. Với message size 1 KB, bandwidth ~1.7 MB/s write. Số liệu này quyết định partition count của Kafka, số replica của DB, và memory footprint của Redis.');
-- Step 1 wrong

INSERT INTO explanations (id, choice_id, level, content)
VALUES ('caex0001-0002-0000-0000-000000000003', 'ca000000-0000-0000-0001-000000000002', 'advanced',
     'Premature optimization is the root of all evil — nhưng premature under-planning là thứ giết chết startups. Không có load estimate, bạn không thể set resource limits đúng, dẫn đến OOM kills và cascading failures dưới production load.');
-- Step 2 correct

INSERT INTO explanations (id, choice_id, level, content)
VALUES ('caex0002-0001-0000-0000-000000000003', 'ca000000-0000-0000-0002-000000000001', 'advanced',
     'Kafka với replication factor 3 đảm bảo at-least-once delivery. Consumer group cho phép horizontal scale DB writers. Retention policy cho phép replay messages khi consumer fails. Throughput Kafka ~ 1M msg/s trên commodity hardware.');
-- Step 2 wrong

INSERT INTO explanations (id, choice_id, level, content)
VALUES ('caex0002-0002-0000-0000-000000000003', 'ca000000-0000-0000-0002-000000000002', 'advanced',
     'Write amplification + connection pool exhaustion: với 100k concurrent users và mỗi write mất 5ms, cần 500k connections/giây — không DB nào chịu được. P99 latency tăng exponentially do queue buildup tại DB layer.');
-- Step 3 correct

INSERT INTO explanations (id, choice_id, level, content)
VALUES ('caex0003-0001-0000-0000-000000000003', 'ca000000-0000-0000-0003-000000000001', 'advanced',
     'Redis throughput > 1M ops/s với sub-millisecond latency. Cache hot conversations (top 10% chiếm 90% reads — Pareto principle). Invalidation strategy: append-on-write với max list size cap, tránh thundering herd khi cache expires.');
-- Step 3 wrong

INSERT INTO explanations (id, choice_id, level, content)
VALUES ('caex0003-0002-0000-0000-000000000003', 'ca000000-0000-0000-0003-000000000002', 'advanced',
     'DB reads at scale: mỗi query với ORDER BY + LIMIT trên conversation_id cần index scan. Với 100M messages và 100k concurrent reads, DB IOPS bị saturate. Read replica giúp nhưng thêm replication lag và complexity. Cache là giải pháp đúng đắn nhất.');
-- Step 4 correct

INSERT INTO explanations (id, choice_id, level, content)
VALUES ('caex0004-0001-0000-0000-000000000003', 'ca000000-0000-0000-0004-000000000001', 'advanced',
     'WebSocket multiplexing qua single TCP connection. Với 100k concurrent users, cần 100k open WebSocket connections — sử dụng event-driven server (Node.js, Go goroutines) thay vì thread-per-connection. Sticky session hoặc pub/sub (Redis pub/sub) để route messages qua multiple chat server instances.');
-- Step 4 wrong

INSERT INTO explanations (id, choice_id, level, content)
VALUES ('caex0004-0002-0000-0000-000000000003', 'ca000000-0000-0000-0004-000000000002', 'advanced',
     'Long polling cải thiện polling nhưng vẫn có overhead của HTTP handshake mỗi response. Server-Sent Events (SSE) là middle ground — unidirectional push nhưng không cần custom protocol. WebSocket vẫn tốt hơn cho bidirectional real-time communication.');
-- Step 5 correct

INSERT INTO explanations (id, choice_id, level, content)
VALUES ('caex0005-0001-0000-0000-000000000003', 'ca000000-0000-0000-0005-000000000001', 'advanced',
     'Horizontal scale với consistent hashing để route WebSocket connections đến đúng server instance. Redis pub/sub làm message bus giữa các instances. Health check + circuit breaker đảm bảo traffic không bị route đến unhealthy instances. Cost-effective hơn vertical: n × small instances rẻ hơn 1 × large instance.');
-- Step 5 wrong

INSERT INTO explanations (id, choice_id, level, content)
VALUES ('caex0005-0002-0000-0000-000000000003', 'ca000000-0000-0000-0005-000000000002', 'advanced',
     'Amdahl''s Law: speedup từ vertical scaling bị giới hạn bởi phần serial của code. Moore''s Law chậm lại — CPU frequency không tăng nhiều nữa. Cloud pricing: reserved instance 96-core đắt hơn 8 × 12-core instances và kém flexible hơn nhiều.');
-- -----------------------------------------------------------------------------
-- Consequences (flags set bởi choices — dùng cho conditional logic)
-- -----------------------------------------------------------------------------
-- Chọn Queue → flag has_queue = true (dùng trong scenario phức tạp hơn)

INSERT INTO consequences (id, choice_id, type, keys, value)
VALUES ('cacons001-0000-0000-0000-000000000001', 'ca000000-0000-0000-0002-000000000001', 'flag', 'has_queue', true);
-- Chọn Redis → flag has_cache = true

INSERT INTO consequences (id, choice_id, type, keys, value)
VALUES ('cacons002-0000-0000-0000-000000000001', 'ca000000-0000-0000-0003-000000000001', 'flag', 'has_cache', true);
-- Chọn WebSocket → flag has_websocket = true

INSERT INTO consequences (id, choice_id, type, keys, value)
VALUES ('cacons003-0000-0000-0000-000000000001', 'ca000000-0000-0000-0004-000000000001', 'flag', 'has_websocket', true);
