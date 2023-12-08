# 环境准备

- https://learnku.com/articles/65095

## 数据

## consul

## jaeger
- 查看追踪   http://srv.in:16686/search
![](./__markdown.images/README/README-1680028247494.png)


# 启动服务



## 1. user 服务

- 这里就注册了了服务

![](./.markdown.images/README/README-1692174211459.png)

![](./.markdown.images/README/README-1692174224898.png)

![](./.markdown.images/README/README-1692175946155.png)

![](./.markdown.images/README/README-1692175975187.png)

- 服务的注册

![](./.markdown.images/README/README-1692175988040.png)

- 服务的发现(使用)

![](./.markdown.images/README/README-1692176192009.png)


```bash 
cd service/user && kratos run 
```

## 2. shop
```bash 
cd shop && kratos run 
```


