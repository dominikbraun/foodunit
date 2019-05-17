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
     * @var string
     */
    private static $email = 'foodunit_email';

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
     * @return string|bool
     */
    public static function email()
    {
        if (!isset($_COOKIE[self::$email])) {
            return false;
        }
        return $_COOKIE[self::$email];
    }
}
