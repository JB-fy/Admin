# 常用命令

1.GO服务端
```bash
# 服务启动
gf run main.go -a "http"

# dao层生成
gf gen dao

# service层生成
gf gen service

# 全自动代码生成（必须在项目api文件夹下且已经生成main文件）
./main myGen
```

2.PHP服务端
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

3.VUE前端
```bash
# 服务启动
npm run dev
# 打包
npm run build
```