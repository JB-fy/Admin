<?php

declare(strict_types=1);

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
    function throwFailJson(string $code, string $msg = '', array $data = [])
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
