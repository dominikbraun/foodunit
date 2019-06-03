<?php

namespace foodunit\core;

/**
 * Class Context
 * @package foodunit\core
 */
class Context
{
    /**
     * @var string
     */
    private static $key = 'foodunit_key';

    /**
     * @return string|bool
     */
    public static function key()
    {
        if (!isset($_COOKIE[self::$key])) {
            return false;
        }
        return $_COOKIE[self::$key];
    }

    /**
     * @param string $key
     */
    public static function set(string $key)
    {
        setcookie(self::$key, $key, 0, '/', '', false, true);
    }
}
