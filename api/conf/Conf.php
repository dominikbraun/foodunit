<?php

namespace foodunit\conf;

/**
 * Class Conf
 * @package foodunit\conf
 */
class Conf
{
    /**
     * @var array
     */
    private static $conf;

    /**
     * Conf initializer.
     */
    public static function init()
    {
        self::$conf = parse_ini_file('conf.ini');
    }

    /**
     * @param string $key
     * @return bool|mixed
     */
    public static function get(string $key)
    {
        if (!isset(self::$conf[$key])) {
            return false;
        }
        return self::$conf[$key];
    }
}
Conf::init();
