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

    public function getSupplierInfo()
    {
        $bindings = [
            'id' => $this->supplierId
        ];
        $supplierInfo = $this->db->query(/** @lang sql */'
            SELECT      id, name, address, phone, mon, tue, wed, thu, fri
            FROM        suppliers
            WHERE       id = :id
        ', $bindings);

        $supplierInfo = count($supplierInfo) ? $supplierInfo[0] : [];

        return $supplierInfo;
    }

    /**
     * @return array
     */
    public function getMenu()
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
            $dish = [
                'id' => $row['id'],
                'name' => $row['name'],
                'desc' => $row['description'],
                'price' => (float) $row['price']
            ];
            $found = false;

            foreach ($cats as &$c) {
                if ($c['id'] === $catId) {
                    $c['dishes'][] = $dish;
                    $found = true;
                }
            }
            if (!$found) {
                $bindings = [
                    'id' => $catId
                ];
                $cat = $this->db->query(/** @lang sql */'
                    SELECT      id, name, img
                    FROM        cats
                    WHERE       id = :id      
                ', $bindings);

                $cat = $cat[0];

                $cats[] = [
                    'id' => $cat['id'],
                    'name' => $cat['name'],
                    'img' => $cat['img'],
                    'dishes' => [
                        $dish
                    ]
                ];
            }
        }
        return $cats;
    }
}
