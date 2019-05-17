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
     * @var string
     */
    private $okCode;

    /**
     * Connection constructor.
     */
    public function __construct()
    {
        $con = Conf::get('db_con_str');
        try {
            $this->pdo = new \PDO($con, Conf::get('user'), Conf::get('pass'));
        } catch (\PDOException $e) {
        }
        $this->okCode = '00000';
    }

    /**
     * @param string $sql
     * @param array $bindings
     * @return array
     */
    public function query(string $sql, array $bindings = [])
    {
        $stmt = $this->pdo->prepare($sql);
        $stmt->execute($bindings);

        return $stmt->fetchAll();
    }

    /**
     * @param string $sql
     * @param array $bindings
     * @return bool
     */
    public function exec(string $sql, array $bindings = [])
    {
        $stmt = $this->pdo->prepare($sql);
        $stmt->execute($bindings);

        return $stmt->errorCode() === $this->okCode;
    }

    /**
     * @return bool
     */
    public function close()
    {
        $this->pdo = null;
        return true;
    }
}
