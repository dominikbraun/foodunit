<?php

dispatch();

/**
 * @return bool
 */
function dispatch()
{
    if (hasKey()) {
        render('app');
    } else {
        render('login');
    }
    return true;
}

/**
 * @return bool
 */
function hasKey()
{
    return $_COOKIE['foodunit_key'] !== null;
}

/**
 * @param string $tpl
 */
function render(string $tpl)
{
    $dir = 'static/templates/';
    $ext = '.html';

    /** @noinspection PhpIncludeInspection */
    include_once $dir . $tpl . $ext;
}
