# go-docker-judger

## Introduction

An OJ judger server by docker and golang, now this project is developing.

if you are interested in it, you can contract with me by email:

- 1437876073@qq.com

- pushy.zhengzuqin@gmail.com

## Usage

Firstly, you need build image from docker directory:

```shell
$ docker image build -t pushyzheng/go-docker-judger .
```

Then run following command to start container, the output will show judgement result:

```shell
$ docker run -v "e:/usr/pushy":/usr/src/oj/code -v "e:/usr/cases":/usr/src/oj/cases pushyzheng/go-docker-judger
```

But your computer must have `Main.java` in `e:/usr/pushy`, if don't have this path, you can change it in top command.

In addition, you can put case file to `/e:/usr/cases`, the form of test case file as following:

```
1
2
3
```