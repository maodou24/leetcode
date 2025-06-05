--
-- @lc app=leetcode id=1164 lang=mysql
--
-- [1164] Product Price at a Given Date
--
-- @lc code=start
-- Write your PostgreSQL query statement below
WITH P AS (
    SELECT
        product_id,
        MAX(change_date) as change_date
    FROM
        Products
    WHERE
        change_date <= '2019-08-16'
    GROUP BY
        product_id
)
SELECT
    product_id,
    new_price as price
FROM
    Products
WHERE
    (product_id, change_date) IN (
        SELECT
            product_id,
            change_date
        FROM
            P
    )
UNION
SELECT
    product_id,
    10 as price
FROM
    Products
WHERE
    product_id NOT IN (
        SELECT
            product_id
        FROM
            P
    )
GROUP BY
    product_id -- @lc code=end