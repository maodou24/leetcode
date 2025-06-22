--
-- @lc app=leetcode id=585 lang=postgresql
--
-- [585] Investments in 2016
--
-- @lc code=start
-- Write your PostgreSQL query statement below
SELECT
    ROUND(CAST(SUM(tiv_2016) AS numeric), 2) AS tiv_2016
FROM
    Insurance AS i1
WHERE
    pid IN (
        SELECT
            MIN(pid)
        FROM
            Insurance
        GROUP BY
            (lat, lon)
        HAVING
            COUNT(pid) = 1
    )
    AND (
        (
            SELECT
                COUNT(*)
            FROM
                Insurance AS i2
            WHERE
                i1.tiv_2015 = i2.tiv_2015
        ) > 1
    ) -- @lc code=end