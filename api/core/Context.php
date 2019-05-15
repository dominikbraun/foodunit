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
    private static $userCookie = 'foodunit_user';

    /**
     * @var string
     */
    private static $sessCookie = 'foodunit_sess';

    /**
     * @return string|null
     */
    public static function user()
    {
        if (!isset($_COOKIE[self::$userCookie])) {
            return null;
        }
        return $_COOKIE[self::$userCookie];
    }

    /**
     * @return string|null
     */
    public static function sess()
    {
        if (!isset($_COOKIE[self::$sessCookie])) {
            return null;
        }
        return $_COOKIE[self::$sessCookie];
    }
}