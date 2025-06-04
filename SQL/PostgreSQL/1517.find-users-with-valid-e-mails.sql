--
-- @lc app=leetcode id=1517 lang=postgresql
--
-- [1517] Find Users With Valid E-Mails
--

-- @lc code=start
-- Write your PostgreSQL query statement below
SELECT
    *
FROM
    Users
WHERE
    regexp_like(mail, '^[A-Za-z][A-Za-z0-9_.-]*@leetcode\.com$', 'i')
-- @lc code=end

