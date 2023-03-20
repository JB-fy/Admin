<?php

var_dump(openssl_encrypt('123456', 'AES-128-ECB', 'e26e543ce222cc0f', 0));
var_dump(openssl_decrypt('dasfsdfds', 'AES-128-ECB', 'e26e543ce222cc0f', 0));
var_dump(strlen('哈哈'));

var_dump(json_decode(openssl_decrypt(base64_decode('dFdWa1N0VklJUm9LUWVMY0dRa0FtY0VnOWxFYzlKS0lQUXUwQktDU09iMDNqZjVhZnJDT1ZnNXY1cnhqUmU1dG91OURDOVlhUkNZVGxUWXB1MjVTOHBtaG9vb09NRm13clV1OVlnMGNNUWFldWsxbFVjNkc5S1crQTc2SmxvblhuQVB2Tm9ld2xWa3JLRVhlcExKZ0RPNDJUdUlEaHBJaDZpTVJsOXllMTRpRkdkeTZkS2xXWklEcjYwRGw3S3ZjeXcyZ3dpK3NxZVlYMDhOeWgzZnhvM1E2MHowMzFnWS9hREdpS3BRSVg1clF4aEJpdkJmMWJNVThobENWM295M1BueGhndkV0WnFuUXFPSlZZUTZTajVjVHlJdzRwZkpjQzRNUlBUU0FOUlI5akJJVWFOYkNqYUZJakx6aWRSY09kdTJycVljTUR5T3ZHSEcxU1lnbzBZQUVPei9UMW92d21jUW00bjVCM1hnUVhDRTBHYzg2STFjbHEyZlUyUUpyYkRQQUJOK3RXMXVZQWQrZi9yY0xiNFFnU2FFU2huRkR4ZDJ1Q3FmNitrNDdtbXFNNHRrRU9VRVJuWDg4NTFvODBLcHdzR0RYWjRITTVHYzVSN2xOVzg4SDJCY0s2LzRJM3lWb25rUkR1T1loM2drZElPdjROMjgyLzNhWnQrcEtFbU9HdU8rWHhXTGFJVFdQUjhMb0hjdm9BcElWTnE5MmRsZ1NaVGlpTTlHbEN2OGorZDlnQ2QzcUVIY0F5Z0FS'), 'AES-128-ECB', 'e26e543ce222cc0f', 0), true));


/* 
SQL语法顺序：select->from->where->group by->having->order by->limit
SQL执行顺序：from->where->group by->having->select->order by->limit

url组成部分：scheme://host:port/path?query#anchor

框架使用说明：
    启动服务
        php bin/hyperf.php start
        php bin/hyperf.php server:watch //热更新
    快速生成模型类，需先修改AbstractModel继承自\Hyperf\DbConnection\Model\Model（--pool对应哪个库）
        php bin/hyperf.php gen:model --pool=default 
    快速生成Dao类，需先修改AbstractDao继承自\Hyperf\DbConnection\Model\Model，再注释掉冲突的方法，生成后再修改
        php bin/hyperf.php gen:model --pool=default --path=app/Module/Db/Dao --inheritance=AbstractDao --uses='App\Module\Db\Dao\AbstractDao'
    
框架使用规范：
    1：使用容器和依赖注入功能时，需要特别注意对象是否含有状态。以下为带有状态的类
        app\Module\Db\Dao内的类（内部的属性几乎都含有状态）
        app\Module\Db\Model内的类（如果继承\Hyperf\DbConnection\Model\Model模型的话，且要切换连接或表时，$connection和$table带有状态）
        app\Module\Cache内的类（当要切换连接时，$cache带有状态）
    2：禁止在任何地方使用app\Module\Db\Model类，防止改变其状态，否则会对app\Module\Db\Dao的使用造成影响
    3：只使用app\Module\Db\Dao或Hyperf\DbConnection\Db做数据库处理
    4：app\Module\Db\Dao文件夹内的类统一使用getDao方法实例化
    5：app\Module\Cache文件夹内的类统一使用getCache方法实例化
 */