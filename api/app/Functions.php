<?php

declare(strict_types=1);

/*----------------基于业务逻辑封装的函数  开始----------------*/
if (!function_exists('getRequestScene')) {
    /**
     * 获取当前请求场景
     *
     * @return string|null
     */
    function getRequestScene(): ?string
    {
        return getContainer()->get(\Hyperf\HttpServer\Contract\RequestInterface::class)->header('Scene');
    }
}

/**
 * 获取当前请求是http还是https
 *
 * @return string
 */
/* function getHttpScheme(): string
{
    //必须在nginx配置文件中设置转发该头部
    return request()->header('x-forwarded-proto', '');
} */
/*----------------基于业务逻辑封装的函数  结束----------------*/


/*----------------基于当前框架封装的函数  开始----------------*/
if (!function_exists('throwSuccessJson')) {
    /**
     * 抛出错误（利用错误处理返回结果json格式。好处：深层次调用无需反复return）
     *
     * @param array $data
     * @param string $code
     * @param string $msg
     * @throws \App\Exception\Json
     * @return void
     */
    function throwSuccessJson(array $data = [], string $code = '000000', string $msg = '')
    {
        throw make(\App\Exception\Json::class, ['code' => $code, 'msg' => $msg, 'data' => $data]);
    }
}


if (!function_exists('throwFailJson')) {
    /**
     * 抛出错误（利用错误处理返回结果json格式。好处：深层次调用无需反复return）
     *
     * @param string $code
     * @param string $msg
     * @param array $data
     * @throws \App\Exception\Json
     * @return void
     */
    function throwFailJson(string $code = '999999', string $msg = '', array $data = [])
    {
        throw make(\App\Exception\Json::class, ['code' => $code, 'msg' => $msg, 'data' => $data]);
    }
}

if (!function_exists('throwRaw')) {
    /**
     * 抛出错误（利用错误处理返回结果raw格式。好处：深层次调用无需反复return）
     *
     * @param string $raw
     * @throws \App\Exception\Raw
     * @return void
     */
    function throwRaw(string $raw)
    {
        throw make(\App\Exception\Raw::class, ['raw' => $raw]);
    }
}

if (!function_exists('getContainer')) {
    /**
     * 获取容器
     *
     * @return \Psr\Container\ContainerInterface
     */
    function getContainer(): \Psr\Container\ContainerInterface
    {
        return \Hyperf\Utils\ApplicationContext::getContainer();
    }
}

if (!function_exists('getDao')) {
    /**
     * 获取Dao对象
     *  注意：
     *      app\Module\Db\Dao文件夹内的类统一使用此方法实例化。防止误使用容器获取，容器获取的实例状态改变会污染框架环境
     *
     * @param string $className
     * @return object
     */
    function getDao(string $className): object
    {
        return make($className);
    }
}

if (!function_exists('getCache')) {
    /**
     * 获取Cache对象
     *  注意：
     *      当确定使用的缓存对象一定不会切换连接库时（即不改变app\Module\Cache\AbstractCache类的$cache变量），可使用容器缓存获取
     * @param string $className
     * @return object
     */
    function getCache(string $className): object
    {
        //return make($className);
        return getContainer()->get($className);
    }
}

if (!function_exists('getLogger')) {
    /**
     * 获取Logger对象
     * 
     * @param string $name  对应文件内开头的name
     * @param string $group 对应config/autoload/logger.php内的key
     * @return \Psr\Log\LoggerInterface
     */
    function getLogger(string $name = 'hyperf', string $group = 'default'): \Psr\Log\LoggerInterface
    {
        return getContainer()->get(\Hyperf\Logger\LoggerFactory::class)->get($name, $group);
    }
}
/*----------------基于当前框架封装的函数  结束----------------*/


/*----------------基于PHP封装的函数  开始----------------*/
if (!function_exists('randStr')) {
    /**
     * 随机字符串
     *
     * @param int $length
     * @param bool $number
     * @param bool $lower
     * @param bool $upper
     * @return string
     */
    function randStr(int $length, bool $number = true, bool $lower = true, bool $upper = true): string
    {
        mt_srand();
        $str = '';
        $number ? $str .= '0123456789' : null;
        $lower ? $str .= 'abcdefghijklmnopqrstuvwxyz' : null;
        $upper ? $str .= 'ABCDEFGHIJKLMNOPQRSTUVWXYZ' : null;
        $strLen = strlen($str);
        $randStr = '';
        for ($i = 0; $i < $length; $i++) {
            $index = mt_rand(0, $strLen - 1);
            $randStr .= $str[$index];
        }
        return $randStr;
    }
}
/*----------------基于PHP封装的函数  结束----------------*/
