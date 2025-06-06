--
-- @lc app=leetcode id=1341 lang=postgresql
--
-- [1341] Movie Rating
--
-- @lc code=start
-- Write your PostgreSQL query statement below
WITH MR AS (
    SELECT
        u.name,
        m.title,
        mr.rating,
        mr.created_at
    FROM
        MovieRating AS mr
        JOIN Movies m ON m.movie_id = mr.movie_id
        JOIN Users u ON u.user_id = mr.user_id
) (
    SELECT
        name as results
    FROM
        MR
    GROUP BY
        name
    ORDER BY
        COUNT(rating) DESC,
        name
    LIMIT
        1
)
UNION
ALL (
    SELECT
        title as results
    FROM
        MR
    WHERE
        created_at >= '2020-02-01'
        AND created_at < '2020-03-01'
    GROUP BY
        title
    ORDER BY
        AVG(rating) DESC,
        title
    LIMIT
        1
) -- @lc code=end