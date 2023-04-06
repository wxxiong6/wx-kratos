#  wx  template for kratos
这个项目是基于 [kratos](https://github.com/go-kratos/kratos) 的项目模板，使用 [cookiecutter](https://github.com/cookiecutter/cookiecutter) 生成项目模板，使用 [wire](https://github.com/google/wire) 生成依赖注入代码。

在做项目中经常会遇到一些重复性的工作，比如创建项目，创建数据库连接，创建路由，创建配置文件等等，这些工作都是重复性的，而且很多时候我们都是按照自己的习惯来创建项目，这样就会导致项目之间的风格不一致，这个项目就是为了解决这些问题，基于平时开发项目提炼出来的一些常用基础组件为模板，以标准规范助力企业研发提速增效。

## Config
```json
{
    "repo_name": "greeter",
    "service_name": "Greeter",
    "api_dir_name": "helloworld",
    "module_name": "github.com/go-kratos/kratos-layout"
}
```
配置文件参数：
- repo_name: 项目名称 (小写)
- service_name: 服务名称 （首字母大写）
- api_dir_name: api目录名称 (小写)
- module_name: 模块名称 (小写)
  
  根据自己的项目修改这些配置，然后执行下面的命令就可以生成项目模板了。

## Usage:

```shell
pip3 install cookiecutter
```

```shell
cookiecutter https://github.com/wxxiong6/wx-kratos
```

```shell
make init
make all
make wire
```
