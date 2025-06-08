--
-- @lc app=leetcode id=602 lang=mysql
--
-- [602] Friend Requests II: Who Has the Most Friends
--
-- @lc code=start
-- Write your PostgreSQL query statement below
WITH RS AS (
    SELECT
        requester_id AS id
    FROM
        RequestAccepted
    UNION
    ALL
    SELECT
        accepter_id AS id
    FROM
        RequestAccepted
)
SELECT
    id,
    COUNT(id) AS num
FROM
    RS
GROUP BY
    id
ORDER BY
    num DESC
LIMIT
    1 -- @lc code=end