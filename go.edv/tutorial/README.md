
### [Accessing a relational database](https://go.dev/doc/tutorial/database-access)

需要准备的数据:

```sql
DROP TABLE IF EXISTS album;
CREATE TABLE album (
  id         INT AUTO_INCREMENT NOT NULL,
  title      VARCHAR(128) NOT NULL,
  artist     VARCHAR(255) NOT NULL,
  price      DECIMAL(5,2) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO album
  (title, artist, price)
VALUES
  ('Blue Train', 'John Coltrane', 56.99),
  ('Giant Steps', 'John Coltrane', 63.99),
  ('Jeru', 'Gerry Mulligan', 17.99),
  ('Sarah Vaughan', 'Sarah Vaughan', 34.98);
```

| section | description | note |
|---|---|---|
| [Find and import a database driver](https://go.dev/doc/tutorial/database-access#import_driver) | 在 [SQLDrivers](https://go.dev/wiki/SQLDrivers) 中查找对应 database driver 的 module. | import 后使用 `go get .` 下载和引入该 module. |
| [Get a database handle and connect](https://go.dev/doc/tutorial/database-access#get_handle) | | 获取 database handle 时不一定会 connect, 所以需要 ping. |
| [Query for multiple rows](https://go.dev/doc/tutorial/database-access#multiple_rows) | 见 function [`albumsByArtist`](https://github.com/mikurisan/golang/blob/main/go.edv/tutorial/access-a-relational-database/main.go#L72) | 学而时习之. |
| [Query for a single row](https://go.dev/doc/tutorial/database-access#single_row) | 见 function [`albumByID`](https://github.com/mikurisan/golang/blob/main/go.edv/tutorial/access-a-relational-database/main.go#L96) | 学而时习之. |
| [Add data](https://go.dev/doc/tutorial/database-access#add_data) | 见 function [`addAlbum`](https://github.com/mikurisan/golang/blob/main/go.edv/tutorial/access-a-relational-database/main.go#L111) | 学而时习之. |
