<?php

var_dump(openssl_encrypt('123456', 'AES-128-ECB', 'e26e543ce222cc0f', 0));
var_dump(openssl_decrypt('dasfsdfds', 'AES-128-ECB', 'e26e543ce222cc0f', 0));
var_dump(strlen('哈哈'));
