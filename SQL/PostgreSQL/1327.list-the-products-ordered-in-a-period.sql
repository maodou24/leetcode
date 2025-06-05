--
-- @lc app=leetcode id=1327 lang=postgresql
--
-- [1327] List the Products Ordered in a Period
--
-- @lc code=start
-- Write your PostgreSQL query statement below
SELECT
    p.product_name,
    SUM(o.unit) AS unit
FROM
    Orders AS o
    JOIN Products AS p ON o.product_id = p.product_id
WHERE
    o.order_date >= '2020-02-01'
    AND o.order_date < '2020-03-01'
GROUP BY
    p.product_name
HAVING
    SUM(o.unit) >= 100 -- @lc code=end