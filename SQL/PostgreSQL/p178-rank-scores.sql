-- Write your PostgreSQL query statement below
-- Window Functions https://www.postgresql.org/docs/17/functions-window.html#FUNCTIONS-WINDOW-TABLE
SELECT
    score, dense_rank() OVER(ORDER BY score DESC) AS rank
FROM
    Scores
