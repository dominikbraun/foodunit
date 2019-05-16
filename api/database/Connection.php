<?php

namespace foodunit\database;

use foodunit\conf\Conf;

/**
 * Class Connection
 * @package foodunit\database
 */
class Connection
{
    /**
     * @var \PDO
     */
    private $pdo;

    /**
     * Connection constructor.
     */
    public function __construct()
    {
        $con = Conf::get('db_con_str');
        $this->pdo = new \PDO($con, Conf::get('user'), Conf::get('pass'));
    }
}
