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
     * @return string|null
     */
    public static function key()
    {
        if (!isset($_COOKIE[self::$key])) {
            return null;
        }
        return $_COOKIE[self::$key];
    }

    /**
     * @return string|null
     */
    public static function email()
    {
        if (!isset($_COOKIE[self::$email])) {
            return null;
        }
        return $_COOKIE[self::$email];
    }
}
