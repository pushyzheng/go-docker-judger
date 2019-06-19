# go-docker-judger

## 介绍

一个使用Docker实现的 `OJ Judger Server`，该项目目前还在开发中。

如果你感兴趣的话，可以通过邮箱来联系我：

- 1437876073@qq.com

- pushy.zhengzuqin@gmail.com

## 使用

首先，你需要构建Docker镜像从`docker`文件夹中的`DockerFile`文件：

```shell
$ docker image build -t pushyzheng/go-docker-judger .
```

然后运行下面的命令启动容器，输出的内容将会是判题的结果：

```shell
$ docker run -v "e:/usr/pushy":/usr/src/oj/code -v "e:/usr/cases":/usr/src/oj/cases pushyzheng/go-docker-judger
```

但是你的系统上在`e:/usr/pushy`路径下必须存在`Main.java`文件，如果没有该路径，可以在上面的命令中修改。


另外，你还可以添加测试样例文件到`/e:/usr/cases`中，样例文件的格式如下：


```
1
2
3
```