<div align=center>
<img src="https://jslx01.oss-cn-hangzhou.aliyuncs.com/common/20230806/1691260208998_6905.ico" width=300" height="300" />
</div>
<div align=center>

# JB Admin
</div>
<div align=center>
<img src="https://img.shields.io/badge/Golang-1.20-blue"/>
<img src="https://img.shields.io/badge/GoFrame-2.5-lightBlue"/>
<img src="https://img.shields.io/badge/PHP-8.2-blue"/>
<img src="https://img.shields.io/badge/Hyperf-3.0-lightBlue"/>
<img src="https://img.shields.io/badge/Vue-3.2.38-blue"/>
<img src="https://img.shields.io/badge/Element Plus-2.3.8-lightBlue"/>
</div>

# 重要提示

如果您将此项目用于商业用途，请遵守Apache2.0协议并保留作者技术支持声明。您需保留如下版权声明信息，其余信息功能不做任何限制。如需剔除需联系作者购买授权

<img src="https://jslx01.oss-cn-hangzhou.aliyuncs.com/common/20230806/1691262207933_9012.png" width="1000">

# 常用命令（GO和PHP任选一个）

#### GO服务端（推荐使用，有代码自动生成功能。所在目录：api）
```bash
# 服务启动
gf run main.go -a "http"

# dao层生成
gf gen dao

# service层生成
gf gen service

# 全自动前后端代码生成（必须先生成main文件。如修改过api/internal/cmd/my_gen.go代码，则需重新生成）
./main myGen
```

#### PHP服务端（无代码自动生成功能。所在目录：api_sw）
```bash
# 服务启动
php bin/hyperf.php start

# 服务启动（热更新）
php bin/hyperf.php server:watch

# 快速生成模型类（--pool对应哪个库）
php bin/hyperf.php gen:model --pool=default 

# 快速生成Dao类，需先修改AbstractDao继承自\Hyperf\DbConnection\Model\Model，再注释掉冲突的方法，生成后再修改
php bin/hyperf.php gen:model --pool=default --path=app/Module/Db/Dao --inheritance=AbstractDao --uses='App\Module\Db\Dao\AbstractDao'
```

#### VUE前端（所在目录：dev/platform）
```bash
# 服务启动
npm run dev
# 打包
npm run build
```

## 微信交流群
| 微信 |
|  :---:  | 
| <img width="150" src="https://jslx01.oss-cn-hangzhou.aliyuncs.com/common/20230805/1691250502303_7330.jpg"> 

添加微信，备注"加入JB Admin交流群"

## 捐献
| 请作者喝饮料 |
|  :---:  | 
| <img width="150" src="https://jslx01.oss-cn-hangzhou.aliyuncs.com/common/20230805/1691251000749_8189.jpg"> 

## 商用注意事项

如果您将此项目用于商业用途，请遵守Apache2.0协议并保留作者技术支持声明。