-- write your code in PostgreSQL 9.4
SELECT 
    d.dept_id,
    COUNT(e.emp_id) as count,
    SUM(e.salary) as sum_of_salary 
FROM 
    department as d 
    inner join employee as e
        on e.dept_id = d.dept_id
GROUP BY
    d.dept_id
ORDER BY 
    d.dept_id