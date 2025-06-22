--
-- @lc app=leetcode id=1757 lang=postgresql
--
-- [1757] Recyclable and Low Fat Products
--
-- @lc code=start
-- Write your PostgreSQL query statement below
SELECT
    product_id
FROM
    Products
WHERE
    low_fats = 'Y'
    AND recyclable = 'Y' -- @lc code=end