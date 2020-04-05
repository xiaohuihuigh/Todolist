DROP TABLE task_info;
CREATE TABLE task_info
(
    id          integer PRIMARY KEY AUTOINCREMENT, --唯一标示
    title       text NOT NULL,                     --标题
    context     text    DEFAULT NULL,              --内容
    type        integer DEFAULT 1,                 --task的类型 1,2 todo|task
    priority    integer DEFAULT 1,                 --优先级 1-9  映射列表
    sub_id      text    DEFAULT ',',              --子task的所有id用`,`隔开
    parent_id   integer  default -1,              --父task的id
    attention   text    DEFAULT NULL,              --提醒时间，时间类型 09:11 只有小时维度| 自定义提醒时间，与创建时间相同，但是只有可以解析的可以作为提醒
    create_time text    DEFAULT NULL,              --创建时间
    operator       integer NOT NULL,              -- 操作人
    created_at  text    DEFAULT NULL,
    updated_at  text    DEFAULT NULL,
    version     integer DEFAULT NULL
);
CREATE TABLE task_info_expand_for_type2
(
    id              integer PRIMARY KEY,  --唯一标示对应task_info中的id
    progress        integer DEFAULT 0,    --type2使用 进度  0-99
    status          integer DEFAULT 1,    --type2使用 是否完结 1未开始，2进行中，3已结束
    start_time      text    DEFAULT NULL, --type2使用 开始时间
    plan_end_time   text    DEFAULT NULL, --type2使用  计划结束时间
    actual_end_time text    DEFAULT NULL, --type2使用 实际结束时间
    created_at      text    DEFAULT NULL,
    updated_at      text    DEFAULT NULL,
    version         integer DEFAULT NULL
);
CREATE TABLE task_info_expand_for_multi_user
(
    id         integer PRIMARY KEY,  --唯一标示对应task_info中的id
    processer  text    DEFAULT NULL, --多人状态使用  保存user_info 的id，用,隔开，多人处理
    creater    integer DEFAULT NULL, --多人状态使用  保存user_info 的id
    updater    integer DEFAULT NULL, --多人状态使用  保存user_info 的id
    created_at text    DEFAULT NULL,
    updated_at text    DEFAULT NULL,
    version    integer DEFAULT NULL
);
CREATE TABLE task_log
(
    id         integer PRIMARY KEY AUTOINCREMENT, --唯一标示对应task_info中的id
    task_id    integer,                           --对应task_info中的id
    title      varchar(255),                      --更改task的标题：
    log_type   varchar(255),                      --log的类型
    updater    integer,                           --修改人 对应user_info的id
    befor      text,                              --之前的字段
    after      text,                              --变化后的字段
    created_at text DEFAULT NULL
);
CREATE TABLE user_info
(
    id         integer      NOT NULL PRIMARY KEY AUTOINCREMENT, --唯一标示
    nickname   varchar(128) NOT NULL,
    phone      varchar(16) DEFAULT NULL,
    created_at text        DEFAULT NULL,
    updated_at text        DEFAULT NULL,
    version    integer     DEFAULT NULL
);
drop table user_info_expand;
CREATE TABLE user_info_expand
(
    id                 integer NOT NULL PRIMARY KEY, --唯一标示
    avatar             text,
    page_size integer default 10,
    weixin_nickname    varchar(128) DEFAULT NULL,
    weixin_avatar      text,
    continue_login_day tinyint(1)   DEFAULT '0',
    department_name    varchar(255) DEFAULT null,
    Position           varchar(255) DEFAULT null,
    `last_login_date`  TEXT         DEFAULT '0000-00-00 00:00:00'
);

CREATE TABLE individual_priority
(
    id         integer PRIMARY KEY AUTOINCREMENT,--唯一标识
    owner      integer NOT NULL,                 --user_id
    one        text DEFAULT NULL,
    two        text DEFAULT NULL,
    three      text DEFAULT NULL,
    four       text DEFAULT NULL,
    five       text DEFAULT NULL,
    six        text DEFAULT NULL,
    seven      text DEFAULT NULL,
    eight      text DEFAULT NULL,
    nine       text DEFAULT NULL,
    created_at text DEFAULT NULL,
    updated_at text DEFAULT NULL,
    version    integer
);
CREATE TABLE individual_time
(
    id         integer PRIMARY KEY AUTOINCREMENT, --唯一标识
    owner      integer NOT NULL,                  --user_id
    parse_time text    DEFAULT NULL,
    created_at text    DEFAULT NULL,
    updated_at text    DEFAULT NULL,
    version    integer DEFAULT NULL
)

