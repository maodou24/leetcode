--
-- @lc app=leetcode id=1148 lang=postgresql
--
-- [1148] Article Views I
--

-- @lc code=start
-- Write your PostgreSQL query statement below
SELECT DISTINCT author_id as id FROM Views WHERE Views.author_id = Views.viewer_id ORDER BY id
-- @lc code=end

