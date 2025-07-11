--
-- @lc app=leetcode id=1683 lang=postgresql
--
-- [1683] Invalid Tweets
--

-- @lc code=start
SELECT
    twee
    t_id
FROM
    Tweets
WHERE
    char_length(content) > 15
-- @lc code=end

