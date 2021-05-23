行数（初期）

```
mysql> select table_name, table_rows from information_schema.TABLES where table_schema = "ishocon2";
+------------+------------+
| table_name | table_rows |
+------------+------------+
| candidates |         30 |
| users      |    3973877 |
| votes      |          0 |
+------------+------------+
3 rows in set (0.00 sec)
```

最新スコア

```
2021/05/23 12:38:30 {"score": 26086, "success": 19518, "failure": 0}
```