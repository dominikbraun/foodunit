<?php

namespace foodunit\services;

/**
 * Class OfferService
 * @package foodunit\services
 */
class OfferService
{
    /**
     * OfferService constructor.
     */
    public function __construct()
    {
    }

    /**
     * @return array
     */
    public function getActiveOffers()
    {
        return [];
    }

    /**
     * @param int $offerId
     * @return array
     */
    public function getAllOrders(int $offerId)
    {
        return [];
    }

    /**
     * @param int $offerId
     * @param string $userEmail
     * @return array
     */
    public function getUserOrder(int $offerId, string $userEmail)
    {
        return [];
    }

    /**
     * @param int $offerId
     * @param int $dishId
     * @param string $userEmail
     * @return bool
     */
    public function addDishToOrder(int $offerId, int $dishId, string $userEmail)
    {
        return true;
    }
}