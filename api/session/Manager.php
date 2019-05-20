<?php

namespace foodunit\session;

require_once 'util/Mail.php';

use foodunit\conf\Conf;
use foodunit\core\Context;
use foodunit\database\Connection;
use foodunit\util\Mail;

/**
 * Class Manager
 * @package foodunit\session
 */
class Manager
{
    /**
     * @var Connection
     */
    private $db;

    /**
     * Manager constructor.
     */
    public function __construct()
    {
        $this->db = new Connection();
    }

    /**
     * @param string $email
     * @return bool
     */
    public function startSession(string $email)
    {
        $from = Conf::get('mail_from');
        $subject = Conf::get('mail_subject');

        $token = self::generateUniqueString();

        if ($token === false) {
            return false;
        }
        self::createSession($email, $token);
        $url = self::getConfirmationUrl($token);

        $success = (new Mail($from, $email))
            ->setSubject($subject)
            ->setBody($url)
            ->send();

        return $success;
    }

    /**
     * @param string $token
     * @return string
     */
    private function getConfirmationUrl(string $token)
    {
        $api = Conf::get('api_url');
        $url = $api . '/confirm/' . $token;

        return $url;
    }

    /**
     * @param string $email
     * @param string $token
     * @return bool
     */
    private function createSession(string $email, string $token)
    {
        $bindings = [
            'email' => $email,
            'confirmation_token' => $token
        ];
        $success = $this->db->exec(/** @lang sql */'
            INSERT INTO sessions
                        (_key, email, valid, confirmation_token, confirmed)
            VALUES      (\'\', :email, 0, :confirmation_token, 0)
        ', $bindings);

        return $success;
    }

    /**
     * @param string $token
     * @return bool
     */
    public function confirmSession(string $token)
    {
        $key = self::generateUniqueString();

        if ($key === false) {
            return false;
        }
        $bindings = [
            'key' => $key,
            'confirmation_token' => $token
        ];
        $success = $this->db->exec(/** @lang sql */'
            UPDATE  sessions
            SET     _key = :key,
                    valid = 1,
                    confirmed = 1
            WHERE   confirmation_token = :confirmation_token
        ', $bindings);

        if ($success) {
            Context::set($key);
        }
        return $success;
    }

    /**
     * @return bool|string
     */
    private function generateUniqueString()
    {
        try {
            $token = bin2hex(random_bytes(16));
            return $token;
        } catch (\Exception $e) {
            return false;
        }
    }
}
