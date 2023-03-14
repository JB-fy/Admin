<?php

/**
 * This file is part of webman.
 *
 * Licensed under The MIT License
 * For full copyright and license information, please see the MIT-LICENSE.txt
 * Redistributions of files must retain the above copyright notice.
 *
 * @author    walkor<walkor@workerman.net>
 * @copyright walkor<walkor@workerman.net>
 * @link      http://www.workerman.net/
 * @license   http://www.opensource.org/licenses/mit-license.php MIT License
 */

namespace support;

/**
 * Class Response
 * @package support
 */
class Response extends \Webman\Http\Response
{
    public function __construct($status = 200, $headers = array(), $body = '')
    {
        // $this->_status = $status;
        // $this->_header = $headers;
        // $this->_body   = $body;
        parent::__construct($status, $headers, $body);
        //自定义头部，全局设置
        $this->_header['Server'] = getenv('APP_NAME');  //修改名称。隐藏是用workerman做的
        $this->_header['Access-Control-Allow-Credentials'] = 'true';
        //$this->_header['Access-Control-Allow-Origin'] = request()->header('Origin', '*');
        //$this->_header['Access-Control-Allow-Origin'] = 'http://www.xxxx.com';
        $this->_header['Access-Control-Allow-Origin'] = '*';
        //$this->_header['Access-Control-Allow-Methods'] = 'GET, POST, PUT, DELETE, PATCH, OPTIONS';
        $this->_header['Access-Control-Allow-Methods'] = '*';
        //$this->_header['Access-Control-Allow-Headers'] = 'X-Requested-With, Content-Type, Accept, Origin, Authorization';   //如果有自定义头部此处需要加上,为方便直接使用*不限制
        $this->_header['Access-Control-Allow-Headers'] = '*';
    }
}
