use todo;
update individual_priority
set owner= 1,
    one='一级',
    two = '二级',
    three = '三级',
    four = '四级',
    five = '五级',
    six = '六级',
    seven = '七级',
    eight = '八级',
    nine = '九级',
    created_at ='2020-02-02 10:10:10',
    updated_at = '2020-02-02 10:10:10',
    version = 1
where id = 1;
update user_info
set nickname = 'default',
    phone = 18717899732,
    created_at = '2020-02-02 10:10:10',
    updated_at = '2020-02-02 10:10:10'
where  id = 1;
update individual_time
set parse_time = '{"list":[{"time":"08:00","parse":"上班","type":"short","sort":1},{"time":"19:00","parse":"下班","type":"short","sort":1},{"time":"05:00","parse":"起床","type":"short","sort":1},{"time":"24:00","parse":"休息","type":"short","sort":1},{"time":"07:00","parse":"早饭","type":"short","sort":1},{"time":"12:00","parse":"午饭","type":"short","sort":1},{"time":"18:00","parse":"晚饭","type":"short","sort":1},{"time":"10","type":"发工资","parse":"word","sort":2},{"time":"06","parse":"还京东","type":"short","sort":1},{"time":"01*10","type":"long","parse":"上旬","sort":4},{"time":"11*20","type":"long","parse":"中旬","sort":4},{"time":"21*31","type":"long","parse":"下旬","sort":4},{"time":"01-12*02-13","type":"long","parse":"春节","sort":3},{"time":"12*13","type":"long","parse":"word","sort":4}]}',
    owner= 1,
    created_at = '2020-02-02 10:10:10',
    updated_at = '2020-02-02 10:10:10',
    version = 1
where  id = 1