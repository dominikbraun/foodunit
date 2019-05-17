<?php

namespace foodunit\core;

require_once 'ResponseDispatcher.php';
require_once 'Context.php';
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
     * @var ResponseDispatcher
     */
    private $dispatcher;

    /**
     * RouteHandler constructor.
     */
    public function __construct()
    {
        $this->dispatcher = new ResponseDispatcher();
    }

    /**
     * @param Request $req
     * @param Response $res
     * @param array $args
     */
    public function offers(Request $req, Response $res, array $args)
    {
        $offers = (new OfferService())->getActiveOffers();
        $this->dispatcher->run($offers);
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

        $this->dispatcher->run($dishes);
    }

    /**
     * @param Request $req
     * @param Response $res
     * @param array $args
     */
    public function orders(Request $req, Response $res, array $args)
    {
        $offerId = $args['offer'];
        $orders = (new OfferService())->getAllOrders($offerId);

        $this->dispatcher->run($orders);
    }

    /**
     * @param Request $req
     * @param Response $res
     * @param array $args
     */
    public function userOrder(Request $req, Response $res, array $args)
    {
        $offerId = $args['offer'];
        $userKey = Context::key();

        $userOrder = (new OfferService())->getUserOrder($offerId, $userKey);

        $this->dispatcher->run($userOrder);
    }

    /**
     * @param Request $req
     * @param Response $res
     * @param array $args
     */
    public function add(Request $req, Response $res, array $args)
    {
        $offerId = $args['offer'];
        $dishId = $args['dish'];
        $userKey = Context::key();

        $res = (new OfferService())->addDishToOrder($offerId, $dishId, $userKey);

        $this->dispatcher->run([$res]);
    }

    /**
     * @param Request $req
     * @param Response $res
     * @param array $args
     */
    public function delete(Request $req, Response $res, array $args)
    {
        $offerId = $args['offer'];
        $dishId = $args['dish'];
        $userKey = Context::key();

        $res = (new OfferService())->deleteDishFromOrder($offerId, $dishId, $userKey);

        $this->dispatcher->run([$res]);
    }

    /**
     * @param Request $req
     * @param Response $res
     * @param array $args
     */
    public function sso(Request $req, Response $res, array $args)
    {
        $this->dispatcher->run([]);
    }

    /**
     * @param Request $req
     * @param Response $res
     * @param array $args
     */
    public function confirm(Request $req, Response $res, array $args)
    {
        $this->dispatcher->run([]);
    }
}
