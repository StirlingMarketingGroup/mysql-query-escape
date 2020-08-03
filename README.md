# MySQL Query Escape/Unescape

A small MySQL UDF library for escaping and unescaping query strings written in Golang. Also known as urlencode/urldecode, encodeURIComponent/decodeURIComponent.

## Usage

### `query_escape`

Escapes the string so it can be safely placed inside a URL query.

```sql
`query_escape` ( `string` )
```

```sql
select`query_escape`('test?');
-- 'test%3F'
select`query_escape`('шеллы');
-- '%D1%88%D0%B5%D0%BB%D0%BB%D1%8B'
```

### `query_unescape`

Does the inverse transformation of `query_escape`, converting each 3-byte encoded substring of the form "%AB" into the hex-decoded byte 0xAB. It returns null if any % is not followed by two hexadecimal digits.

```sql
`query_unescape` ( `string` )
```

```sql
select`query_unescape`('test%3F');
-- 'test?'
select`query_unescape`('%D1%88%D0%B5%D0%BB%D0%BB%D1%8B');
-- 'шеллы'
```

## Dependencies

You will need Golang, which you can get from here https://golang.org/doc/install.

Debian / Ubuntu

```shell
sudo apt update
sudo apt install libmysqlclient-dev
```

## Installing

You can find your MySQL plugin directory by running this MySQL query

```sql
select @@plugin_dir;
```

then replace `/usr/lib/mysql/plugin` below with your MySQL plugin directory.

```shell
cd ~ # or wherever you store your git projects
git clone https://github.com/StirlingMarketingGroup/mysql-query-escape.git
cd mysql-query-escape
go build -buildmode=c-shared -o mysql_query_escape.so
sudo cp mysql_query_escape.so /usr/lib/mysql/plugin/mysql_query_escape.so # replace plugin dir here if needed
```

Enable the functions in MySQL by running these MySQL queries

```sql
create function`query_escape`returns string soname'mysql_query_escape.so';
create function`query_unescape`returns string soname'mysql_query_escape.so';
```