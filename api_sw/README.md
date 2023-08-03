# Introduction

This is a skeleton application using the Hyperf framework. This application is meant to be used as a starting place for those looking to get their feet wet with Hyperf Framework.

# Requirements

Hyperf has some requirements for the system environment, it can only run under Linux and Mac environment, but due to the development of Docker virtualization technology, Docker for Windows can also be used as the running environment under Windows.

The various versions of Dockerfile have been prepared for you in the [hyperf/hyperf-docker](https://github.com/hyperf/hyperf-docker) project, or directly based on the already built [hyperf/hyperf](https://hub.docker.com/r/hyperf/hyperf) Image to run.

When you don't want to use Docker as the basis for your running environment, you need to make sure that your operating environment meets the following requirements:  

 - PHP >= 8.0
 - Any of the following network engines
   - Swoole PHP extension >= 4.5，with `swoole.use_shortname` set to `Off` in your `php.ini`
   - Swow PHP extension (Beta)
 - JSON PHP extension
 - Pcntl PHP extension
 - OpenSSL PHP extension （If you need to use the HTTPS）
 - PDO PHP extension （If you need to use the MySQL Client）
 - Redis PHP extension （If you need to use the Redis Client）
 - Protobuf PHP extension （If you need to use the gRPC Server or Client）

# Installation using Composer

The easiest way to create a new Hyperf project is to use Composer. If you don't have it already installed, then please install as per the documentation.

To create your new Hyperf project:

$ composer create-project hyperf/hyperf-skeleton path/to/install

Once installed, you can run the server immediately using the command below.

$ cd path/to/install
$ php bin/hyperf.php start

This will start the cli-server on port `9501`, and bind it to all network interfaces. You can then visit the site at `http://localhost:9501/`

which will bring up Hyperf default home page.

框架使用说明：
    启动服务
        php bin/hyperf.php start
        php bin/hyperf.php server:watch //热更新
    快速生成模型类（--pool对应哪个库）
        php bin/hyperf.php gen:model --pool=default 
    快速生成Dao类，需先修改AbstractDao继承自\Hyperf\DbConnection\Model\Model，再注释掉冲突的方法，生成后再修改
        php bin/hyperf.php gen:model --pool=default --path=app/Module/Db/Dao --inheritance=AbstractDao --uses='App\Module\Db\Dao\AbstractDao'
    
框架使用规范：
    1：使用容器和依赖注入功能时，需要特别注意对象是否含有状态。以下带有状态的类严禁使用容器获取实例，如使用则必须以new className()或make(className::class)的方式生成实例
        app\Module\Db\Model内的类（当切换连接或表时，$connection和$table带有状态），建议统一使用getModel方法生成实例
        app\Module\Db\Dao内的类（内部的属性几乎都含有状态），建议统一使用getDao方法生成实例
        app\Module\Cache内的类（当要切换连接时，$cache带有状态），建议统一使用getCache方法生成实例
    2：数据库处理建议使用app\Module\Db\Dao或Hyperf\DbConnection\Db，尽量不使用app\Module\Db\Model