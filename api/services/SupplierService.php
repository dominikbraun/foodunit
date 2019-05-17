<?php

namespace foodunit\services;

require_once 'database/Connection.php';

use foodunit\database\Connection;

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
     * @var Connection
     */
    private $db;

    /**
     * SupplierService constructor.
     * @param int $supplierId
     */
    public function __construct(int $supplierId)
    {
        $this->supplierId = $supplierId;
        $this->db = new Connection();
    }

    /**
     * @return array
     */
    public function getDishes()
    {
        $bindings = [
            'supplier_id' => $this->supplierId
        ];
        $res = $this->db->query(/** @lang sql */'
            SELECT      id, cat_id, name, description, price
            FROM        dishes
            WHERE       supplier_id = :supplier_id
            ORDER BY    id DESC
        ', $bindings);

        return $res;
    }
}
