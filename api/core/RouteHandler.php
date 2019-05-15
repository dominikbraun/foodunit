<?php

namespace foodunit\core;

require_once 'services/OfferService.php';
require_once 'services/SupplierService.php';

use foodunit\services\OfferService;
use foodunit\services\SupplierService;
use Slim\Http\Request;
use Slim\Http\Response;

/**
 * Class RouteHandler
 * @package foodunit\core
 */
class RouteHandler
{
    /**
     * RouteHandler constructor.
     */
    public function __construct()
    {
    }

    /**
     * @param Request $req
     * @param Response $res
     * @param array $args
     */
    public function offers(Request $req, Response $res, array $args)
    {
        $offers = (new OfferService())->getActiveOffers();
    }

    /**
     * @param Request $req
     * @param Response $res
     * @param array $args
     */
    public function dishes(Request $req, Response $res, array $args)
    {
        $supplierId = $args['supplier'];
        $dishes = (new SupplierService($supplierId))->getDishes();
    }

    /**
     * @param Request $req
     * @param Response $res
     * @param array $args
     */
    public function orders(Request $req, Response $res, array $args)
    {
    }

    /**
     * @param Request $req
     * @param Response $res
     * @param array $args
     */
    public function userOrder(Request $req, Response $res, array $args)
    {
    }

    /**
     * @param Request $req
     * @param Response $res
     * @param array $args
     */
    public function add(Request $req, Response $res, array $args)
    {
    }

    /**
     * @param Request $req
     * @param Response $res
     * @param array $args
     */
    public function del(Request $req, Response $res, array $args)
    {
    }
}
