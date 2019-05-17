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
     * @param array $response
     * @param int $status
     */
    public function run(array $response, $status = 200)
    {
        echo json_encode($response);
    }
}