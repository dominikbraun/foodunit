<?php

namespace foodunit\conf;

class Conf
{
    private static $conf;

    public static function init()
    {
        self::$conf = parse_ini_file('conf.ini');
    }

    public static function get(string $key)
    {
        if (!isset(self::$conf[$key])) {
            return false;
        }
        return self::$conf[$key];
    }
}
Conf::init();
