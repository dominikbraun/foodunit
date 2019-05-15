<?php

namespace foodunit\services;

/**
 * Class SupplierService
 * @package foodunit\services
 */
class SupplierService
{
    /**
     * @var int|int
     */
    private $supplierId;

    /**
     * SupplierService constructor.
     * @param int $supplierId
     */
    public function __construct(int $supplierId)
    {
        $this->supplierId = $supplierId;
    }

    /**
     * @return array
     */
    public function getDishes()
    {
        return [];
    }
}