-- =============================================================================
-- Seed: URL Shortener Scenario
-- Scenario ID : us000000-0000-0000-0000-000000000000
-- Steps       : 5 steps (hash → storage → redirect cache → analytics → abuse)
-- Safe to re-run: ON CONFLICT (id) DO NOTHING on every insert
-- =============================================================================

-- -----------------------------------------------------------------------------
-- Scenario
-- -----------------------------------------------------------------------------
INSERT INTO scenarios (id, title, description, difficulty) VALUES
    (
        'us000000-0000-0000-0000-000000000000',
        'URL Shortener',
        'Design a URL shortening service (like bit.ly) that handles 100M shortened URLs, 10B redirects/month, and < 10ms redirect latency.',
        'easy'
    )
ON CONFLICT (id) DO NOTHING;

-- -----------------------------------------------------------------------------
-- Steps
-- -----------------------------------------------------------------------------
INSERT INTO steps (id, scenario_id, question, context, order_index) VALUES
    (
        'us000000-0000-0000-0000-000000000001',
        'us000000-0000-0000-0000-000000000000',
        'How will you generate the short code for each URL?',
        'Users submit a long URL (e.g. https://example.com/very/long/path?query=value) and expect a short code like "abc123". This code must be unique across 100 million URLs.',
        1
    ),
    (
        'us000000-0000-0000-0000-000000000002',
        'us000000-0000-0000-0000-000000000000',
        'Which database will you use to store the short-code → original-URL mapping?',
        'You need to store 100M URL mappings and serve 10B lookups per month. The access pattern is almost entirely key-value: given a short code, return the original URL.',
        2
    ),
    (
        'us000000-0000-0000-0000-000000000003',
        'us000000-0000-0000-0000-000000000000',
        'How will you make redirects as fast as possible?',
        'A redirect is the most frequent operation: user clicks a short link, your service must respond with a 301/302 in under 10ms. At 10B redirects/month, that is ~3,800 requests/second on average.',
        3
    ),
    (
        'us000000-0000-0000-0000-000000000004',
        'us000000-0000-0000-0000-000000000000',
        'How will you track how many times each short URL was clicked?',
        'The marketing team wants real-time click counts per short URL. Every redirect must be counted. This is 10B events per month.',
        4
    ),
    (
        'us000000-0000-0000-0000-000000000005',
        'us000000-0000-0000-0000-000000000000',
        'Users are abusing the service to shorten malicious URLs. How do you prevent this?',
        'Some short URLs redirect to phishing or malware sites. Users can also flood the system with thousands of shortening requests per second.',
        5
    )
ON CONFLICT (id) DO NOTHING;

-- Set start_step_id
UPDATE scenarios
SET start_step_id = 'us000000-0000-0000-0000-000000000001'
WHERE id = 'us000000-0000-0000-0000-000000000000';

-- -----------------------------------------------------------------------------
-- Choices
-- UUID scheme: us000000-0000-0000-<step>-00000000000<idx>
-- -----------------------------------------------------------------------------

-- Step 1: Short code generation
INSERT INTO choices (id, step_id, label, next_step_id, is_correct) VALUES
    (
        'us000000-0000-0000-0001-000000000001',
        'us000000-0000-0000-0000-000000000001',
        'Hash the original URL with MD5, take the first 7 characters of the hex output',
        'us000000-0000-0000-0000-000000000002',
        true
    ),
    (
        'us000000-0000-0000-0001-000000000002',
        'us000000-0000-0000-0000-000000000001',
        'Use a random UUID (e.g. "550e8400") as the short code',
        'us000000-0000-0000-0000-000000000002',
        false
    )
ON CONFLICT (id) DO NOTHING;

-- Step 2: Storage
INSERT INTO choices (id, step_id, label, next_step_id, is_correct) VALUES
    (
        'us000000-0000-0000-0002-000000000001',
        'us000000-0000-0000-0000-000000000002',
        'Use a NoSQL key-value store (e.g. DynamoDB or Cassandra) — short code as the key',
        'us000000-0000-0000-0000-000000000003',
        true
    ),
    (
        'us000000-0000-0000-0002-000000000002',
        'us000000-0000-0000-0000-000000000002',
        'Use a relational database (PostgreSQL) with an index on the short_code column',
        'us000000-0000-0000-0000-000000000003',
        false
    )
ON CONFLICT (id) DO NOTHING;

-- Step 3: Redirect caching
INSERT INTO choices (id, step_id, label, next_step_id, is_correct) VALUES
    (
        'us000000-0000-0000-0003-000000000001',
        'us000000-0000-0000-0000-000000000003',
        'Cache the short-code → URL mapping in Redis (TTL 1 hour) and use a CDN for global edge caching',
        'us000000-0000-0000-0000-000000000004',
        true
    ),
    (
        'us000000-0000-0000-0003-000000000002',
        'us000000-0000-0000-0000-000000000003',
        'Query the database on every redirect request with no caching layer',
        'us000000-0000-0000-0000-000000000004',
        false
    )
ON CONFLICT (id) DO NOTHING;

-- Step 4: Click analytics
INSERT INTO choices (id, step_id, label, next_step_id, is_correct) VALUES
    (
        'us000000-0000-0000-0004-000000000001',
        'us000000-0000-0000-0000-000000000004',
        'Emit a click event to a message queue (Kafka) asynchronously — a separate consumer updates counters',
        'us000000-0000-0000-0000-000000000005',
        true
    ),
    (
        'us000000-0000-0000-0004-000000000002',
        'us000000-0000-0000-0000-000000000004',
        'Run UPDATE clicks = clicks + 1 in the database synchronously on every redirect',
        'us000000-0000-0000-0000-000000000005',
        false
    )
ON CONFLICT (id) DO NOTHING;

-- Step 5: Abuse prevention (last step)
INSERT INTO choices (id, step_id, label, next_step_id, is_correct) VALUES
    (
        'us000000-0000-0000-0005-000000000001',
        'us000000-0000-0000-0000-000000000005',
        'Implement rate limiting per IP (token bucket) and check submitted URLs against a malware/phishing blocklist API',
        NULL,
        true
    ),
    (
        'us000000-0000-0000-0005-000000000002',
        'us000000-0000-0000-0000-000000000005',
        'Add a CAPTCHA to the URL shortening form and rely on user reports for malicious links',
        NULL,
        false
    )
ON CONFLICT (id) DO NOTHING;

-- -----------------------------------------------------------------------------
-- Impacts
-- -----------------------------------------------------------------------------

-- Step 1 correct: hash is deterministic, fast, compact
INSERT INTO impacts (id, choice_id, metric, type, value) VALUES
    ('us000000-0001-0001-0001-000000000001', 'us000000-0000-0000-0001-000000000001', 'latency',     'add', -5),
    ('us000000-0001-0001-0001-000000000002', 'us000000-0000-0000-0001-000000000001', 'scalability', 'add',  2)
ON CONFLICT (id) DO NOTHING;

-- Step 1 wrong: UUID is long (36 chars), ugly, not URL-friendly
INSERT INTO impacts (id, choice_id, metric, type, value) VALUES
    ('us000000-0001-0002-0001-000000000001', 'us000000-0000-0000-0001-000000000002', 'scalability', 'add', -1),
    ('us000000-0001-0002-0001-000000000002', 'us000000-0000-0000-0001-000000000002', 'cost',        'add',  1)
ON CONFLICT (id) DO NOTHING;

-- Step 2 correct: NoSQL is designed for key-value at massive scale
INSERT INTO impacts (id, choice_id, metric, type, value) VALUES
    ('us000000-0002-0001-0001-000000000001', 'us000000-0000-0000-0002-000000000001', 'latency',     'add', -10),
    ('us000000-0002-0001-0001-000000000002', 'us000000-0000-0000-0002-000000000001', 'scalability', 'add',   3),
    ('us000000-0002-0001-0001-000000000003', 'us000000-0000-0000-0002-000000000001', 'cost',        'add',   1)
ON CONFLICT (id) DO NOTHING;

-- Step 2 wrong: PostgreSQL works but struggles at 10B queries/month
INSERT INTO impacts (id, choice_id, metric, type, value) VALUES
    ('us000000-0002-0002-0001-000000000001', 'us000000-0000-0000-0002-000000000002', 'latency',     'add', 15),
    ('us000000-0002-0002-0001-000000000002', 'us000000-0000-0000-0002-000000000002', 'scalability', 'add', -2)
ON CONFLICT (id) DO NOTHING;

-- Step 3 correct: Redis + CDN = < 5ms global redirect latency
INSERT INTO impacts (id, choice_id, metric, type, value) VALUES
    ('us000000-0003-0001-0001-000000000001', 'us000000-0000-0000-0003-000000000001', 'latency',     'add', -40),
    ('us000000-0003-0001-0001-000000000002', 'us000000-0000-0000-0003-000000000001', 'scalability', 'add',   3),
    ('us000000-0003-0001-0001-000000000003', 'us000000-0000-0000-0003-000000000001', 'cost',        'add',   2)
ON CONFLICT (id) DO NOTHING;

-- Step 3 wrong: DB per redirect is slow and will explode under load
INSERT INTO impacts (id, choice_id, metric, type, value) VALUES
    ('us000000-0003-0002-0001-000000000001', 'us000000-0000-0000-0003-000000000002', 'latency',     'add', 50),
    ('us000000-0003-0002-0001-000000000002', 'us000000-0000-0000-0003-000000000002', 'scalability', 'add', -3)
ON CONFLICT (id) DO NOTHING;

-- Step 4 correct: async event emission adds no latency to redirect
INSERT INTO impacts (id, choice_id, metric, type, value) VALUES
    ('us000000-0004-0001-0001-000000000001', 'us000000-0000-0000-0004-000000000001', 'latency',     'add',  0),
    ('us000000-0004-0001-0001-000000000002', 'us000000-0000-0000-0004-000000000001', 'scalability', 'add',  2),
    ('us000000-0004-0001-0001-000000000003', 'us000000-0000-0000-0004-000000000001', 'cost',        'add',  1)
ON CONFLICT (id) DO NOTHING;

-- Step 4 wrong: synchronous counter update adds latency and creates hot row
INSERT INTO impacts (id, choice_id, metric, type, value) VALUES
    ('us000000-0004-0002-0001-000000000001', 'us000000-0000-0000-0004-000000000002', 'latency',     'add', 25),
    ('us000000-0004-0002-0001-000000000002', 'us000000-0000-0000-0004-000000000002', 'scalability', 'add', -2)
ON CONFLICT (id) DO NOTHING;

-- Step 5 correct: rate limiting + blocklist is the right defense
INSERT INTO impacts (id, choice_id, metric, type, value) VALUES
    ('us000000-0005-0001-0001-000000000001', 'us000000-0000-0000-0005-000000000001', 'scalability', 'add',  2),
    ('us000000-0005-0001-0001-000000000002', 'us000000-0000-0000-0005-000000000001', 'cost',        'add',  1),
    ('us000000-0005-0001-0001-000000000003', 'us000000-0000-0000-0005-000000000001', 'latency',     'add', -5)
ON CONFLICT (id) DO NOTHING;

-- Step 5 wrong: CAPTCHA is UX-hostile and does not stop bots
INSERT INTO impacts (id, choice_id, metric, type, value) VALUES
    ('us000000-0005-0002-0001-000000000001', 'us000000-0000-0000-0005-000000000002', 'latency',     'add', 10),
    ('us000000-0005-0002-0001-000000000002', 'us000000-0000-0000-0005-000000000002', 'scalability', 'add', -1)
ON CONFLICT (id) DO NOTHING;

-- -----------------------------------------------------------------------------
-- Explanations (3 levels per choice)
-- -----------------------------------------------------------------------------

-- Step 1 correct
INSERT INTO explanations (id, choice_id, level, content) VALUES
    ('usex0001-0001-0000-0000-000000000001', 'us000000-0000-0000-0001-000000000001', 'beginner',
     'MD5 hash giống máy xay sinh tố: dù URL dài đến đâu, output luôn là chuỗi 32 ký tự. Lấy 7 ký tự đầu là đủ cho 268 triệu unique codes.'),
    ('usex0001-0001-0000-0000-000000000002', 'us000000-0000-0000-0001-000000000001', 'intermediate',
     '7 hex chars = 16^7 = 268M combinations — đủ cho 100M URLs với collision probability thấp. Khi collision xảy ra, append counter suffix và retry. MD5 compute time < 1μs, deterministic (same URL → same code).'),
    ('usex0001-0001-0000-0000-000000000003', 'us000000-0000-0000-0001-000000000001', 'advanced',
     'Birthday paradox: với 100M URLs và 268M codes, collision probability ≈ 1 - e^(-n²/2m) ≈ 17%. Giải pháp: check-before-insert với optimistic locking, hoặc dùng distributed ID generator (Snowflake) cho guaranteed uniqueness. MD5 không collision-resistant nhưng đủ tốt cho URL hashing không phải cryptographic context.')
ON CONFLICT (id) DO NOTHING;

-- Step 1 wrong
INSERT INTO explanations (id, choice_id, level, content) VALUES
    ('usex0001-0002-0000-0000-000000000001', 'us000000-0000-0000-0001-000000000002', 'beginner',
     'UUID dài 36 ký tự (ffffffff-ffff-...) — không ai muốn share một URL như vậy. Short URL cần thực sự ngắn: 6-8 ký tự.'),
    ('usex0001-0002-0000-0000-000000000002', 'us000000-0000-0000-0001-000000000002', 'intermediate',
     'UUID v4 là random và guaranteed unique, nhưng 36 chars không URL-friendly và không compressible. Base62 encoding (a-z, A-Z, 0-9) của sequential ID cho 6 chars = 62^6 = 56B combinations — tốt hơn nhiều.'),
    ('usex0001-0002-0000-0000-000000000003', 'us000000-0000-0000-0001-000000000002', 'advanced',
     'UUID v4 dùng 122 bits random → storage cost lớn hơn, không sortable (UUID v7 fix điều này). Cho URL shortener, Base62(sequential_id) tốt nhất: compact, sortable, human-readable, zero collision. Distributed: dùng Snowflake ID (41-bit timestamp + 10-bit machine + 12-bit sequence) → 64-bit integer → ~11 chars Base62.')
ON CONFLICT (id) DO NOTHING;

-- Step 2 correct
INSERT INTO explanations (id, choice_id, level, content) VALUES
    ('usex0002-0001-0000-0000-000000000001', 'us000000-0000-0000-0002-000000000001', 'beginner',
     'NoSQL giống một cuốn từ điển khổng lồ: tra "abc123" ra URL ngay lập tức. Không cần join, không cần tính toán phức tạp — chỉ cần lookup.'),
    ('usex0002-0001-0000-0000-000000000002', 'us000000-0000-0000-0002-000000000001', 'intermediate',
     'DynamoDB/Cassandra được tối ưu cho key-value lookups: O(1) read với consistent hash partitioning. Auto-sharding theo key space, horizontal scale tự động, built-in replication. 100M rows là bình thường.'),
    ('usex0002-0001-0000-0000-000000000003', 'us000000-0000-0000-0002-000000000001', 'advanced',
     'DynamoDB: single-digit millisecond latency at any scale, 99.999% SLA, automatic multi-AZ replication. Partition key = short_code → even distribution. Avoid hot partitions bằng cách KHÔNG dùng timestamp as partition key. Conditional writes prevent race conditions khi tạo short URL.')
ON CONFLICT (id) DO NOTHING;

-- Step 2 wrong
INSERT INTO explanations (id, choice_id, level, content) VALUES
    ('usex0002-0002-0000-0000-000000000001', 'us000000-0000-0000-0002-000000000002', 'beginner',
     'PostgreSQL hoạt động tốt nhưng sẽ chậm dần khi bảng có 100 triệu rows và hàng tỷ queries mỗi tháng.'),
    ('usex0002-0002-0000-0000-000000000002', 'us000000-0000-0000-0002-000000000002', 'intermediate',
     'B-tree index trên short_code giúp lookup O(log n) nhưng vẫn kém NoSQL. Connection pool bottleneck: PostgreSQL handle ~500 concurrent connections tốt, nhưng 3,800 req/s cần > 500 connections. PgBouncer giúp nhưng thêm complexity.'),
    ('usex0002-0002-0000-0000-000000000003', 'us000000-0000-0000-0002-000000000002', 'advanced',
     'PostgreSQL cho URL shortener không sai về correctness nhưng sai về operational simplicity at scale. Sharding PostgreSQL phức tạp (Citus, manual routing). Replication lag trên read replica có thể gây stale redirect. NoSQL là right tool for this access pattern.')
ON CONFLICT (id) DO NOTHING;

-- Step 3 correct
INSERT INTO explanations (id, choice_id, level, content) VALUES
    ('usex0003-0001-0000-0000-000000000001', 'us000000-0000-0000-0003-000000000001', 'beginner',
     'Redis nhớ kết quả tra cứu trong RAM — nhanh hơn DB 100 lần. CDN nhớ kết quả tại server gần người dùng nhất, nhanh hơn cả Redis vì không cần về trung tâm.'),
    ('usex0003-0001-0000-0000-000000000002', 'us000000-0000-0000-0003-000000000001', 'intermediate',
     'Redis: sub-millisecond lookup, cache hot URLs (top 20% URLs chiếm 80% traffic — Pareto). CDN edge cache: 301 Permanent redirect được browser cache, reducing repeat lookups to zero network calls. TTL strategy: popular URLs TTL 1h, new URLs TTL 5m.'),
    ('usex0003-0001-0000-0000-000000000003', 'us000000-0000-0000-0003-000000000001', 'advanced',
     '301 vs 302: 301 Permanent được browser cache → zero server hit cho repeat users. 302 Temporary → always hits server but allows analytics. Trade-off: 301 for static URLs (better UX/perf), 302 for trackable URLs. CDN + Redis two-tier cache: CDN handles ~80% traffic at edge, Redis handles remaining cache misses, DB is last resort.')
ON CONFLICT (id) DO NOTHING;

-- Step 3 wrong
INSERT INTO explanations (id, choice_id, level, content) VALUES
    ('usex0003-0002-0000-0000-000000000001', 'us000000-0000-0000-0003-000000000002', 'beginner',
     'Mỗi lần click vào link là 1 lần truy vấn DB. Với tỷ lượt click mỗi tháng, DB sẽ quá tải ngay.'),
    ('usex0003-0002-0000-0000-000000000002', 'us000000-0000-0000-0003-000000000002', 'intermediate',
     '10B redirects/month = 3,800 req/s average, peak có thể 10-20x = 38,000-76,000 req/s. Không DB nào chịu được 76,000 lookups/s mà không có cache. Thậm chí với SSD và index, PostgreSQL max ~10,000 simple queries/s per instance.'),
    ('usex0003-0002-0000-0000-000000000003', 'us000000-0000-0000-0003-000000000002', 'advanced',
     'DB connection exhaustion: 76,000 req/s × 1ms latency = 76 concurrent connections minimum. Thực tế latency cao hơn → nhiều concurrent connections hơn. Read replica farm giúp nhưng tốn kém hơn cache nhiều. Cache hit rate > 95% là realistic với URL access pattern (power law distribution).')
ON CONFLICT (id) DO NOTHING;

-- Step 4 correct
INSERT INTO explanations (id, choice_id, level, content) VALUES
    ('usex0004-0001-0000-0000-000000000001', 'us000000-0000-0000-0004-000000000001', 'beginner',
     'Ghi nhận click vào hàng đợi Kafka như ghi vào sổ tay — nhanh và không ảnh hưởng đến việc chuyển hướng. Sau đó ai đó đọc sổ tay và cập nhật số liệu.'),
    ('usex0004-0001-0000-0000-000000000002', 'us000000-0000-0000-0004-000000000001', 'intermediate',
     'Fire-and-forget event emission: redirect handler emit event đến Kafka < 1ms, không block response. Separate consumer service batch-processes click events và update analytics DB (ClickHouse/BigQuery). Decoupled, scalable, durable.'),
    ('usex0004-0001-0000-0000-000000000003', 'us000000-0000-0000-0004-000000000001', 'advanced',
     'Lambda architecture: Kafka stream → (1) real-time counter in Redis INCR for live stats, (2) batch aggregation to data warehouse for historical analysis. Exactly-once semantics với Kafka transactions. Clickhouse ingest ~1B rows/day easily. Alternative: AWS Kinesis + Lambda for serverless analytics pipeline.')
ON CONFLICT (id) DO NOTHING;

-- Step 4 wrong
INSERT INTO explanations (id, choice_id, level, content) VALUES
    ('usex0004-0002-0000-0000-000000000001', 'us000000-0000-0000-0004-000000000002', 'beginner',
     'UPDATE trên cùng 1 row mỗi khi có click — nếu link viral và 10,000 người click cùng lúc, tất cả đều tranh nhau sửa cùng 1 dòng trong DB.'),
    ('usex0004-0002-0000-0000-000000000002', 'us000000-0000-0000-0004-000000000002', 'intermediate',
     '"Hot row" problem: row-level lock trên counter row tạo serialization bottleneck. 10,000 concurrent UPDATE trên same row = 9,999 rows chờ. Thêm latency vào redirect path — điều không bao giờ được xảy ra.'),
    ('usex0004-0002-0000-0000-000000000003', 'us000000-0000-0000-0004-000000000002', 'advanced',
     'MVCC trong PostgreSQL: mỗi UPDATE tạo new row version → write amplification + dead tuple accumulation → VACUUM pressure. Tại 76,000 req/s, counter table sẽ bloat nhanh. Redis INCR là O(1) atomic và non-blocking — dùng Redis cho real-time counter, sync to DB periodically (eventual consistency acceptable for analytics).')
ON CONFLICT (id) DO NOTHING;

-- Step 5 correct
INSERT INTO explanations (id, choice_id, level, content) VALUES
    ('usex0005-0001-0000-0000-000000000001', 'us000000-0000-0000-0005-000000000001', 'beginner',
     'Rate limiting giới hạn mỗi IP chỉ được tạo bao nhiêu link mỗi phút. Blocklist kiểm tra xem URL có độc hại không trước khi tạo link.'),
    ('usex0005-0001-0000-0000-000000000002', 'us000000-0000-0000-0005-000000000001', 'intermediate',
     'Token bucket per IP: 10 requests/minute với burst tối đa 20. Malware check: Google Safe Browsing API hoặc VirusTotal API synchronous check khi tạo URL (chấp nhận 100-200ms overhead vì đây là write path, không phải read path). Block và log suspicious IPs.'),
    ('usex0005-0001-0000-0000-000000000003', 'us000000-0000-0000-0005-000000000001', 'advanced',
     'Defense in depth: (1) Rate limiting tại API Gateway (e.g. AWS WAF) — trước khi request vào application. (2) Async re-scan: schedule periodic re-check của stored URLs against updated blocklists. (3) Reputation scoring: IP reputation database (e.g. MaxMind, AbuseIPDB). (4) Domain allowlist/denylist. (5) ML-based URL classifier cho phishing detection với high precision.')
ON CONFLICT (id) DO NOTHING;

-- Step 5 wrong
INSERT INTO explanations (id, choice_id, level, content) VALUES
    ('usex0005-0002-0000-0000-000000000001', 'us000000-0000-0000-0005-000000000002', 'beginner',
     'CAPTCHA phiền người dùng thật và bot hiện đại vẫn có thể giải được. Chờ báo cáo từ người dùng thì link đã phát tán rồi.'),
    ('usex0005-0002-0000-0000-000000000002', 'us000000-0000-0000-0005-000000000002', 'intermediate',
     'CAPTCHA v2 bị bypass bởi CAPTCHA-solving services ($1/1000 solves). Reactive approach (chờ user report) nghĩa là link độc hại tồn tại hàng giờ/ngày trước khi bị xóa. Reputation damage cho service là không thể chấp nhận được.'),
    ('usex0005-0002-0000-0000-000000000003', 'us000000-0000-0000-0005-000000000002', 'advanced',
     'CAPTCHA là friction cho legitimate users nhưng không stop determined attackers. reCAPTCHA Enterprise cải thiện nhưng vẫn không đủ. Proactive defense tốt hơn reactive: rate limiting + blocklist + behavioral analysis (velocity check, user agent fingerprinting) chặn 99% abuse trước khi nó xảy ra.')
ON CONFLICT (id) DO NOTHING;

-- -----------------------------------------------------------------------------
-- Consequences
-- -----------------------------------------------------------------------------

INSERT INTO consequences (id, choice_id, type, keys, value) VALUES
    ('uscons001-0000-0000-0000-000000000001', 'us000000-0000-0000-0003-000000000001', 'flag', 'has_cache', true)
ON CONFLICT (id) DO NOTHING;

INSERT INTO consequences (id, choice_id, type, keys, value) VALUES
    ('uscons002-0000-0000-0000-000000000001', 'us000000-0000-0000-0004-000000000001', 'flag', 'has_async_analytics', true)
ON CONFLICT (id) DO NOTHING;
