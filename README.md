# tada-server

## run server

```shell
rm -f ./data.db && go run main.go
```

## test code

```shell
rm -f ./api/v1/data.db && go test ./api/v1
```

## gulp-aglio

### install

```shell
npm install -g gulp aglio gulp-aglio gulp-rename
```

### run server

```shell
 aglio -i api/all.apib --server
```

## make html

```shell
gulp
```

input: api/all.apib

output: doc/out.html

https://www.lapis-zero09.xyz/tada-server/doc/out.html
