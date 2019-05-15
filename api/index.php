<?php

use Psr\Http\Message\ServerRequestInterface as Request;
use Psr\Http\Message\ResponseInterface as Response;
use \foodunit\core\RouteManager as RouteManager;

require_once 'vendor/autoload.php';
require_once 'core/RouteManager.php';

$app = new Slim\App();
$manager = new RouteManager();

foreach ($manager->mappings() as $route => $f) {
    $app->get($route, $f);
}

$app->run();
