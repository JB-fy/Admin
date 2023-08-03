# 框架使用说明：
1.启动服务
- php bin/hyperf.php start
- php bin/hyperf.php server:watch //热更新
2.快速生成模型类（--pool对应哪个库）
- php bin/hyperf.php gen:model --pool=default 
3.快速生成Dao类，需先修改AbstractDao继承自\Hyperf\DbConnection\Model\Model，再注释掉冲突的方法，生成后再修改
- php bin/hyperf.php gen:model --pool=default --path=app/Module/Db/Dao --inheritance=AbstractDao --uses='App\Module\Db\Dao\AbstractDao'
    
# 框架使用规范：
1.使用容器和依赖注入功能时，需要特别注意对象是否含有状态。以下带有状态的类严禁使用容器获取实例，如使用则必须以new className()或make(className::class)的方式生成实例
- app\Module\Db\Model内的类（当切换连接或表时，$connection和$table带有状态），建议统一使用getModel方法生成实例
- app\Module\Db\Dao内的类（内部的属性几乎都含有状态），建议统一使用getDao方法生成实例
- app\Module\Cache内的类（当要切换连接时，$cache带有状态），建议统一使用getCache方法生成实例
2.数据库处理建议使用app\Module\Db\Dao或Hyperf\DbConnection\Db，尽量不使用app\Module\Db\Model