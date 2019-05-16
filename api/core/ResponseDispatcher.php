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
     */
    public function run(array $response)
    {
        echo json_encode($response);
    }
}