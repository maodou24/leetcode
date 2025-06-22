--
-- @lc app=leetcode id=185 lang=postgresql
--
-- [185] Department Top Three Salaries
--
-- @lc code=start
-- Write your PostgreSQL query statement below
SELECT
    Department,
    Employee,
    Salary
FROM
    (
        SELECT
            d.name AS Department,
            e.name AS Employee,
            e.salary AS Salary,
            DENSE_RANK() OVER (
                PARTITION BY d.name
                ORDER BY
                    e.salary DESC
            ) AS ranking
        FROM
            Employee e
            JOIN Department d ON e.departmentId = d.id
    ) AS r
WHERE
    r.ranking <= 3 -- @lc code=end