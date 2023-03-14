<?php

declare(strict_types=1);

/**
 * 抛出错误（利用错误处理返回结果json格式。好处：深层次调用无需反复return）
 *
 * @param array $data
 * @param string $code
 * @param string $msg
 * @throws \app\exception\Json
 * @return void
 */
function throwSuccessJson(array $data = [], string $code = '000000', string $msg = '')
{
    throw container(\app\exception\Json::class, true, ['code' => $code, 'msg' => $msg, 'data' => $data]);
}

/**
 * 抛出错误（利用错误处理返回结果json格式。好处：深层次调用无需反复return）
 *
 * @param string $code
 * @param string $msg
 * @param array $data
 * @throws \app\exception\Json
 * @return void
 */
function throwFailJson(string $code, string $msg = '', array $data = [])
{
    throw container(\app\exception\Json::class, true, ['code' => $code, 'msg' => $msg, 'data' => $data]);
}

/**
 * 抛出错误（利用错误处理返回结果raw格式。好处：深层次调用无需反复return）
 *
 * @param string $raw
 * @throws \app\exception\Raw
 * @return void
 */
function throwRaw(string $raw)
{
    throw container(\app\exception\Raw::class, true, ['raw' => $raw]);
}

/**
 * 利用容器获取对象（容器拥有依赖注入，注解注入，缓存对象等功能。但是利用容器注入的对象必须是无状态。建议对象都使用这个方法获取）
 *
 * @template T
 * @param string|class-string<T> $name
 * @param boolean $isMake   当类带有状态时，用于创建短生命周期的对象
 * @param array $parameters
 * @return mixed|T
 */
function container(string $name, bool $isMake = false, array $parameters = [])
{
    if ($isMake) {
        /**
         * make每次都会重新创建对象
         * 使用场景：当对象带有状态，且需在当前请求使用后销毁时，建议使用
         * 
         * 带有状态的类：
         *      \app\module\cache\文件夹下的类
         *      \app\module\db\dao\文件夹下的类
         *      \app\module\validate\文件夹下的类继承\think\Validate，带有以下状态：当前场景currentScene，是否批量batch，是否抛出错误failException
         */
        return \support\Container::make($name, $parameters);
    }
    /**
     * get会复用容器内缓存的对象
     * 使用场景：当对象没有状态，可供进程内所有请求共同使用时，建议使用
     * 
     * 每个请求都会共用这些对象，减少对象的创建，销毁。增加服务器负载能力，但其实作用很小，只有在对象初始化需要做某些耗时操作，作用才会突显。
     * 如module/db/table文件夹下的类创建对象时需要调用数据库获取全部字段，这种情况很适合使用容器缓存对象。
     */
    return \support\Container::get($name);
}

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
