--
-- @lc app=leetcode id=1907 lang=postgresql
--
-- [1907] Count Salary Categories
--
-- @lc code=start
SELECT
    c.category,
    COALESCE(q.accounts_count, 0) as accounts_count
FROM
    (
        VALUES
            ('Low Salary'),
            ('Average Salary'),
            ('High Salary')
    ) AS c(category)
    LEFT JOIN (
        SELECT
            (
                CASE
                    WHEN income < 20000 THEN 'Low Salary'
                    WHEN income >= 20000
                    AND income <= 50000 THEN 'Average Salary'
                    WHEN income > 50000 THEN 'High Salary'
                END
            ) as category,
            COUNT(*) as accounts_count
        FROM
            Accounts
        GROUP BY
            category
    ) as q ON c.category = q.category -- @lc code=end