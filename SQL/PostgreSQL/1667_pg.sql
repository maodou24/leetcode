-- string function and operation  https://www.postgresql.org/docs/17/functions-string.html
SELECT
    user_id, format('%s%s', upper(substring(name, 1, 1)), lower(substring(name, 2))) as name
FROM
    Users
ORDER BY
    user_id