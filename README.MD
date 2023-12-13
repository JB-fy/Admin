<div align=center>
<img src="./文档/logo.png" width=300" height="300" />
</div>
<div align=center>

# JB Admin
</div>
<div align=center>
<img src="https://img.shields.io/badge/Golang-1.20-blue"/>
<img src="https://img.shields.io/badge/GoFrame-2.5-lightBlue"/>
<img src="https://img.shields.io/badge/PHP-8.2-blue"/>
<img src="https://img.shields.io/badge/Hyperf-3.0-lightBlue"/>
<img src="https://img.shields.io/badge/Node-18.16.0-blue"/>
<img src="https://img.shields.io/badge/Vue-3.2.38-brightgreen"/>
<img src="https://img.shields.io/badge/Element Plus-2.3.8-green"/>
</div>

# 重要提示

如果您将此项目用于商业用途，请遵守[Apache2.0协议](./LICENSE)并保留作者技术支持声明。您需保留如下版权声明信息，其它功能不做任何限制。如需剔除请联系作者购买授权（往下翻有作者微信）

<img src="./文档/bqsm.png" width="1000">

# [文档目录](./文档)说明

项目内[文档目录](./文档)含有本人记录的学习资料，提供给有兴趣的同学学习使用。

其内含有环境搭建，配置示例等文件。例如：[项目的nginx配置文件示例](./文档/配置示例/nginx/admin.conf)

# 常用命令（服务端GO和PHP任选一个。推荐使用GO服务端，功能更多，如全自动代码生成）

#### GO服务端（推荐。代码目录：[api](./api)）
```bash
# 服务启动（--gf.gcfg.file用于指定配置文件，默认config.yaml）
gf run main.go -a "http --gf.gcfg.file=config.prod.yaml"

# dao层生成
gf gen dao

# service层生成
gf gen service

# 全自动代码生成（需先编译生成可执行文件。如修改api/internal/cmd/my_gen.go代码，需重新编译才会生效）
./main myGen
```

#### PHP服务端（不推荐。代码目录：[api_sw](./api_sw)）
```bash
# 服务启动
php bin/hyperf.php start

# 服务启动（热更新）
php bin/hyperf.php server:watch

# 快速生成模型类（--pool对应哪个库）
php bin/hyperf.php gen:model --pool=default 

# 快速生成Dao类（需先修改AbstractDao继承自\Hyperf\DbConnection\Model\Model，再注释掉冲突的方法，生成后再还原）
php bin/hyperf.php gen:model --pool=default --path=app/Module/Db/Dao --inheritance=AbstractDao --uses='App\Module\Db\Dao\AbstractDao'
```

#### VUE前端（代码目录：[view/platform](./view/platform)）
```bash
# 服务启动
npm run dev
# 打包
npm run build
```

## 微信交流群
| 微信 |
|  :---:  | 
| <img width="150" src="./文档/zzwx.jpg"> 

添加微信，备注"加入JB Admin交流群"

## 捐献
| 请作者喝饮料 |
|  :---:  | 
| <img width="150" src="./文档/zzwxjx.jpg"> 

## 商用注意事项

如果您将此项目用于商业用途，请遵守[Apache2.0协议](./LICENSE)并保留作者技术支持声明。