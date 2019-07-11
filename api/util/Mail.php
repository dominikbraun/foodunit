<?php

namespace foodunit\util;

use foodunit\conf\Conf;
use SendGrid\Mail\TypeException;

/**
 * Class Mail
 * @package foodunit\util
 */
class Mail
{
    /**
     * @var \SendGrid\Mail\Mail
     */
    private $mail;

    /**
     * @var int
     */
    private $okCode;

    /**
     * Mail constructor.
     * @param string $from
     * @param string $to
     */
    public function __construct(string $from, string $to)
    {
        $this->mail = new \SendGrid\Mail\Mail();
        $this->okCode = 202;

        try {
            $this->mail->setFrom($from);
            $this->mail->addTo($to);
        } catch (\Exception $e) {
        }
    }

    /**
     * @param string $subject
     * @return $this
     */
    public function setSubject(string $subject)
    {
        $this->mail->setSubject($subject);
        return $this;
    }

    /**
     * @param string $body
     * @return $this
     */
    public function setBody(string $body)
    {
        $this->mail->addContent('text/html', $body);
        return $this;
    }

    /**
     * @return bool
     */
    public function send()
    {
        $sendGrid = new \SendGrid(Conf::get('sg_api_key'));
        $res = $sendGrid->send($this->mail);

        return $res->statusCode() === $this->okCode;
    }
}
