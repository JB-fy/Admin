# 常用命令（GO和PHP任选一个）

#### GO服务端（推荐使用，有代码自动生成功能）
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

#### PHP服务端（无代码自动生成功能）
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

#### VUE前端
```bash
# 服务启动
npm run dev
# 打包
npm run build
```

### 微信交流群
| 微信 |
|  :---:  | 
| <img width="150" src="https://jslx01.oss-cn-hangzhou.aliyuncs.com/common/20230805/1691250502303_7330.jpg"> 

添加微信，备注"加入JB-Admin交流群"

### 捐献（如果你觉得这个项目对你有帮助，你可以请作者喝饮料）
|  :---:  | 
| <img width="150" src="https://jslx01.oss-cn-hangzhou.aliyuncs.com/common/20230805/1691251000749_8189.jpg"> 