SELECT
    activity_date as day,
    COUNT(DISTINCT user_id) as active_users
FROM
    Activity
WHERE
    '2019-07-27' - activity_date >= 0
    AND '2019-07-27' - activity_date < 30
GROUP BY
    day -- date opertions  http://www.postgres.cn/docs/9.4/functions-datetime.html