# Link

db diagram: [link](https://dbdiagram.io/d/Simple-bank-6637896f5b24a634d08b305b)

## SQL query

### Show db lock

```sql
SELECT  a.application_name,
        l.relation::regclass,
        l.transactionid,
        l.mode,
        l.lockType,
        l.GRANTED,
        a.usename,
        a.query,
        a.pid
FROM pg_stat_activity a
JOIN pg_locks l ON l.pid = a.pid
where a.application_name = 'psql'
ORDER BY a.pid;
```
