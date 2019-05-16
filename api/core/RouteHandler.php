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
        $user = Context::user();

        $userOrder = (new OfferService())->getUserOrder($offerId, $user);

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
        $user = Context::user();

        $res = (new OfferService())->addDishToOrder($offerId, $dishId, $user);

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
        $user = Context::user();

        $res = (new OfferService())->deleteDishFromOrder($offerId, $dishId, $user);

        $this->dispatcher->run([$res]);
    }
}
