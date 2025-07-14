-- Write your PostgreSQL query statement below
SELECT
    d.name AS Department,
    t.name AS Employee,
    t.salary AS Salary
FROM
    (
        SELECT
            departmentId,
            name,
            salary,
            rank() OVER (
                PARTITION BY departmentid
                ORDER BY
                    salary DESC
            ) AS ranking
        FROM
            Employee
    ) AS t
    JOIN Department AS d ON t.departmentId = d.id
WHERE
    t.ranking = 1
ORDER BY
    salary DESC