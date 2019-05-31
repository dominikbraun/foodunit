<?php

namespace foodunit\core;

/**
 * Class ResponseDispatcher
 * @package foodunit\core
 */
class ResponseDispatcher
{
    /**
     * ResponseDispatcher constructor.
     */
    public function __construct()
    {
    }

    /**
     * @param mixed $response
     * @param int $status
     */
    public function run($response, $status = 200)
    {
        echo json_encode($response);
    }

    /**
     * @param string $url
     */
    public function redirect(string $url)
    {
        header('Location: ' . $url);
        exit();
    }
}
