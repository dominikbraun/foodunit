<?php

dispatch();

/**
 * @return bool
 */
function dispatch()
{
    if (isLoggedIn()) {
        render('app');
    } else {
        render('login');
    }
    return true;
}

/**
 * @return bool
 */
function isLoggedIn()
{
    return $_COOKIE['foodunit_key'] !== null;
}

/**
 * @param string $file
 */
function render(string $file)
{
    $dir = 'static/templates/';
    $ext = '.html';

    /** @noinspection PhpIncludeInspection */
    include_once $dir . $file . $ext;
}
