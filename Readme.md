# MyHTTP Tool

## Table of contents
1. [Description](#Description)
2. [Examples](#Example)   
3. [Testing](#Testing)


## 1. Description

A tool which makes http requests and prints the address of the request along with the MD5 hash of the response.

## 2. Examples

> 1\. First build myhttp tool

```shell
go build cmd/myhttp.go
```

> 2\. Check the command-line arguments available to use

```shell
./myhttp -h
```

that will print:

```shell
Usage of ./myhttp:
  -limit int
        int that indicates the limit of concurrent request possible (default 10)
```

> 3\. Execute multiple http requests and print response md5 hashed


```shell
./myhttp http://www.adjust.com http://google.com
```

> Note: Providing only the hostname is sufficient.

 ```shell
./myhttp adjust.com google.com
```

> 3\. Execute multiple http requests with providing limit flag that indicates the possible numb of concurrent request 

```shell
./myhttp -limit=8 adjust.com google.com facebook.com yahoo.com yandex.com twitter.com reddit.com/r/funny reddit.com/r/notfunny baroquemusiclibrary.com
```

## 3. Testing

> Execute all test cases with coverage

```shell
go test -cover -v ./...
```

