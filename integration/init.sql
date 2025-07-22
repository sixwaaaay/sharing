CREATE TABLE users (
    id bigserial,
    name character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    code character varying(6) NOT NULL DEFAULT ''::character varying,
    code_time timestamptz NOT NULL DEFAULT to_timestamp('2020-01-01 00:00:00'::text, 'YYYY-MM-DD HH24:MI:SS'::text),
    code_try integer NOT NULL DEFAULT 0,
    bio character varying(255) NOT NULL DEFAULT ''::character varying,
    avatar character varying(255) NOT NULL DEFAULT ''::character varying,
    created_at timestamptz NOT NULL  DEFAULT now(),
    background character varying(255) NOT NULL DEFAULT ''::character varying,
    PRIMARY KEY (id),
    UNIQUE (email)
);


-- 一个脚本，生成 10 万 条数据
INSERT INTO users (name, email, code, code_time, code_try, bio, avatar, created_at, background)
SELECT
    'user_' || i::text,
    'user_' || i::text || '@example.com',
    '123456',
    now(),
    0,
    'This is a bio for user ' || i::text,
    'https://example.com/avatar_' || i::text || '.jpg',
    now(),
    'https://example.com/background_' || i::text || '.jpg'
FROM generate_series(1, 100000) AS s(i);

CREATE TABLE graph (
    id bigserial,
    relation integer NOT NULL,
    subject_id bigint NOT NULL,
    object_id bigint NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    PRIMARY KEY (id),
    UNIQUE (relation, subject_id, object_id)
);

-- 一个脚本，生成 10 万 条数据
INSERT INTO graph (relation, subject_id, object_id)
SELECT-- 从序列获取下一个值作为 id
    floor(random() * 3) + 1, -- 生成 1 到 3 之间的随机整数
    floor(random() * 100000) + 1, -- 生成 1 到 100000 之间的随机整数
    floor(random() * 100000) + 1 -- 生成 1 到 100000 之间的随机整数
FROM generate_series(1, 100000) AS s(i)
ON CONFLICT (relation, subject_id, object_id) DO NOTHING;



CREATE TABLE videos (
    id bigserial,
    user_id bigint NOT NULL,
    title character varying(255) NOT NULL,
    des text NOT NULL,
    cover_url character varying(255) DEFAULT ''::character varying NOT NULL,
    video_url character varying(255) DEFAULT ''::character varying NOT NULL,
    duration integer,
    view_count integer DEFAULT 0 NOT NULL,
    like_count integer DEFAULT 0 NOT NULL,
    created_at timestamptz DEFAULT now() NOT NULL,
    updated_at timestamptz DEFAULT now() NOT NULL,
    processed integer DEFAULT 0 NOT NULL,
    PRIMARY KEY (id)
);

-- 一个脚本，生成 10 万 条数据
INSERT INTO videos (user_id, title, des, cover_url, video_url, duration, view_count, like_count, created_at, updated_at, processed)
SELECT
    floor(random() * 100000) + 1, -- 生成 1 到 100000 之间的随机整数
    'Video ' || i::text,
    'This is a description for video ' || i::text,
    'https://example.com/cover_' || i::text || '.jpg',
    'https://example.com/video_' || i::text || '.mp4',
    floor(random() * 3600) + 1, -- 生成 1 到 3600 之间的随机整数
    floor(random() * 10000) + 1, -- 生成 1 到 10000 之间的随机整数
    floor(random() * 1000) + 1, -- 生成 1 到 1000 之间的随机整数
    now(),
    now(),
    floor(random() * 2) -- 生成 0 或 1
FROM generate_series(1, 100000) AS s(i);