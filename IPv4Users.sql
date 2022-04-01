-- write your code in PostgreSQL 9.4
SELECT a.web_page_ipv4, max(a.users_cnt) users_cnt FROM (
SELECT
    v.web_page_ipv4, 
    count(v.user_ipv4) users_cnt 
FROM visits v
GROUP BY v.web_page_ipv4, v.user_ipv4
ORDER BY 2 DESC
) a
GROUP BY a.web_page_ipv4;