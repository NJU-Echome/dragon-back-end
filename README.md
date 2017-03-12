course+ 后台代码  

##环境配置

###安装Glide

```
go get github.com/Masterminds/glide
```


###安装依赖

```
glide install
```

运行 `glide install` 会根据 `glide.lock` 文件中的依赖信息安装依赖，如果需要更新依赖版本，请使用

```
glide update
```

```
glide get ...
```
安装新的依赖

会将依赖更新到新版本并更新 `glide.lock` 文件。


##运行

###安装Bee命令行工具
```
go get github.com/beego/bee
```

###运行
```
bee run
```