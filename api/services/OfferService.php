<?php

namespace foodunit\services;

require_once 'database/Connection.php';

use foodunit\database\Connection;

/**
 * Class OfferService
 * @package foodunit\services
 */
class OfferService
{
    /**
     * @var Connection
     */
    private $db;

    /**
     * OfferService constructor.
     */
    public function __construct()
    {
        $this->db = new Connection();
    }

    /**
     * @return array
     */
    public function getActiveOffers()
    {
        $res = $this->db->query(/** @lang text */'
            SELECT  *
            FROM    offers
            WHERE   active = 1
        ');
        return $res;
    }

    /**
     * @param int $offerId
     * @return array
     */
    public function getAllOrders(\int $offerId)
    {
        return [];
    }

    /**
     * @param int $offerId
     * @param string $userKey
     * @return array
     */
    public function getUserOrder(\int $offerId, \string $userKey)
    {
        return [];
    }

    /**
     * @param int $offerId
     * @param int $dishId
     * @param string $userKey
     * @return bool
     */
    public function addDishToOrder(\int $offerId, \int $dishId, \string $userKey)
    {
        return true;
    }

    /**
     * @param int $offerId
     * @param int $dishId
     * @param string $userKey
     * @return bool
     */
    public function deleteDishFromOrder(\int $offerId, \int $dishId, \string $userKey)
    {
        return true;
    }
}
