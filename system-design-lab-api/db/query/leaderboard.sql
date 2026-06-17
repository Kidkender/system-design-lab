-- name: GetLeaderboardByScenario :many
SELECT
    us.id         AS session_id,
    us.user_id,
    u.username,
    us.created_at,
    us.completed_at,
    COUNT(uc.id)::int                                       AS total_choices,
    COUNT(CASE WHEN c.is_correct THEN 1 END)::int           AS correct_choices
FROM user_sessions us
JOIN users u        ON u.id  = us.user_id
JOIN user_choices uc ON uc.session_id = us.id
JOIN choices c       ON c.id  = uc.choice_id
WHERE us.scenario_id = @scenario_id::uuid
  AND us.status      = 'completed'
GROUP BY us.id, us.user_id, u.username
ORDER BY correct_choices DESC, us.completed_at ASC
LIMIT sqlc.arg(top_n)::int;

-- name: GetSessionScore :one
SELECT
    COUNT(uc.id)::int                             AS total_choices,
    COUNT(CASE WHEN c.is_correct THEN 1 END)::int AS correct_choices
FROM user_choices uc
JOIN choices c ON c.id = uc.choice_id
WHERE uc.session_id = @session_id::uuid;

-- name: GetScenarioProgressByUser :many
SELECT
    s.id                                                            AS scenario_id,
    s.title,
    s.difficulty,
    COUNT(us.id)::int                                               AS attempts,
    COUNT(CASE WHEN us.status = 'completed' THEN 1 END)::int        AS completions,
    MAX(us.completed_at)                                            AS last_completed_at
FROM scenarios s
LEFT JOIN user_sessions us ON us.scenario_id = s.id AND us.user_id = @user_id::uuid
GROUP BY s.id, s.title, s.difficulty
ORDER BY s.created_at;
