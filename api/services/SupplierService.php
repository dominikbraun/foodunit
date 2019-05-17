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
        $dishes = $this->db->query(/** @lang sql */'
            SELECT      id, cat_id, name, description, price
            FROM        dishes
            WHERE       supplier_id = :supplier_id
            ORDER BY    id DESC
        ', $bindings);

        $cats = [];

        foreach ($dishes as $row) {
            $catId = $row['cat_id'];

            if (!isset($cats[$catId])) {
                $bindings = [
                    'id' => $catId
                ];
                $cat = $this->db->query(/** @lang sql */'
                    SELECT      id, name, img
                    FROM        cats
                    WHERE       id = :id      
                ', $bindings);
                $cat = $cat[0];

                $cats[$catId] = [
                    'id' => $cat['id'],
                    'name' => $cat['name'],
                    'img' => $cat['img'],
                    'dishes' => []
                ];
            }
            $dish = [
                'id' => $row['id'],
                'name' => $row['name'],
                'desc' => $row['description'],
                'price' => $row['price']
            ];
            $cats[$catId]['dishes'][] = $dish;
        }
        return $cats;
    }
}
