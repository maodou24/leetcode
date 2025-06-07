--
-- @lc app=leetcode id=1321 lang=postgresql
--
-- [1321] Restaurant Growth
--
-- @lc code=start
-- Write your PostgreSQL query statement below
SELECT
    DISTINCT visited_on,
    amount,
    ROUND(amount :: numeric / 7 :: numeric, 2) AS average_amount
FROM
    (
        SELECT
            visited_on,
            SUM(amount) OVER (
                ORDER BY
                    visited_on RANGE BETWEEN INTERVAL '6 day' PRECEDING
                    AND CURRENT ROW
            ) AS amount
        FROM
            Customer
    ) AS t
ORDER BY
    visited_on OFFSET 6 -- @lc code=end