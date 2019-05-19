<?php

namespace foodunit\session;

require_once 'util/Mail.php';

use foodunit\conf\Conf;
use foodunit\util\Mail;

/**
 * Class Manager
 * @package foodunit\session
 */
class Manager
{
    /**
     * Manager constructor.
     */
    public function __construct()
    {
    }

    /**
     * @param string $email
     * @return bool
     */
    public function newSession(string $email)
    {
        $from = Conf::get('mail_from');
        $subject = Conf::get('mail_subject');

        $url = self::confirmationUrl($email);

        $success = (new Mail($from, $email))
            ->setSubject($subject)
            ->setBody($url)
            ->send();

        return $success;
    }

    /**
     * @param string $email
     * @return string
     */
    private function confirmationUrl(string $email)
    {
        return '...';
    }
}
