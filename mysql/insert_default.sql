insert into individual_priority(id, owner, one, two, three,
                                four, five, six, seven, eight,
                                nine, created_at, updated_at,
                                version)
VALUES (1, 1, '一级', '二级', '三级',
        '四级', '五级', '六级', '七级',
        '八级', '九级', '2020-02-02 10:10:10',
        '2020-02-02 10:10:10', 1);
insert into user_info(id,nickname, phone, created_at, updated_at)
VALUES(1,'default',18717899732,'2020-02-02 10:10:10','2020-02-02 10:10:10');
insert into individual_time(id,owner, parse_time, created_at, updated_at, version)
values (1,1,'{"list":[{"time":"08:00","parse":"上班","type":"short","sort":1},{"time":"19:00","parse":"下班","type":"short","sort":1},{"time":"05:00","parse":"起床","type":"short","sort":1},{"time":"24:00","parse":"休息","type":"short","sort":1},{"time":"07:00","parse":"早饭","type":"short","sort":1},{"time":"12:00","parse":"午饭","type":"short","sort":1},{"time":"18:00","parse":"晚饭","type":"short","sort":1},{"time":"10","type":"发工资","parse":"word","sort":2},{"time":"06","parse":"还京东","type":"short","sort":1},{"time":"01*10","type":"long","parse":"上旬","sort":4},{"time":"11*20","type":"long","parse":"中旬","sort":4},{"time":"21*31","type":"long","parse":"下旬","sort":4},{"time":"01-12*02-13","type":"long","parse":"春节","sort":3},{"time":"12*13","type":"long","parse":"word","sort":4}]}',
        '2020-02-02 10:10:10','2020-02-02 10:10:10',1);
insert into user_info_expand(id,page_size)
VALUES(1,'20');