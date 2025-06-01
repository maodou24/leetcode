SELECT
    person_name
FROM (
    SELECT 
        person_name,
        SUM(weight) OVER ( order by turn) AS total_weight
    FROM Queue
)
WHERE total_weight <= 1000
ORDER BY total_weight DESC
LIMIT 1