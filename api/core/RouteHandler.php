<?php

namespace foodunit\core;

require_once 'ResponseDispatcher.php';
require_once 'Context.php';
require_once 'services/OfferService.php';
require_once 'services/SupplierService.php';
require_once 'session/Manager.php';

use foodunit\services\OfferService;
use foodunit\services\SupplierService;
use foodunit\session\Manager;
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
        $key = Context::key();

        $userOrder = (new OfferService())->getUserOrder($offerId, $key);

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
        $key = Context::key();

        $res = (new OfferService())->addDishToOrder($offerId, $dishId, $key);

        $this->dispatcher->run($res);
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
        $key = Context::key();

        $res = (new OfferService())->deleteDishFromOrder($offerId, $dishId, $key);

        $this->dispatcher->run($res);
    }

    /**
     * @param Request $req
     * @param Response $res
     * @param array $args
     */
    public function getRemark(Request $req, Response $res, array $args)
    {
        $offerId = $args['offer'];
        $key = Context::key();

        $remark = (new OfferService())->getRemark($offerId, $key);

        $this->dispatcher->run($remark);
    }

    /**
     * @param Request $req
     * @param Response $res
     * @param array $args
     */
    public function insertRemark(Request $req, Response $res, array $args)
    {
        $offerId = $args['offer'];
        $remark = $args['remark'];
        $key = Context::key();

        $res = (new OfferService())->insertRemark($offerId, $remark, $key);

        $this->dispatcher->run($res);
    }

    /**
     * @param Request $req
     * @param Response $res
     * @param array $args
     */
    public function sso(Request $req, Response $res, array $args)
    {
        $email = $args['email'];
        $res = (new Manager())->startSession($email);

        $this->dispatcher->run($res);
    }

    /**
     * @param Request $req
     * @param Response $res
     * @param array $args
     */
    public function confirmSession(Request $req, Response $res, array $args)
    {
        $token = $args['token'];
        $res = (new Manager())->confirmSession($token);

        $this->dispatcher->run($res);
    }

    /**
     * @param Request $req
     * @param Response $res
     * @param array $args
     */
    public function email(Request $req, Response $res, array $args)
    {
        $email = Context::email();
        $this->dispatcher->run($email);
    }
}
