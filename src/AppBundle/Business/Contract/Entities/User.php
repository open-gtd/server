<?php


namespace AppBundle\Business\Contract\Entities;

/**
 * Class User
 */
class User
{
    /** @var string */
    protected $username;

    /** @var  string */
    protected $email;

    /**
     * @return string
     */
    public function getUsername()
    {
        return $this->username;
    }

    /**
     * @return string
     */
    public function getEmail()
    {
        return $this->email;
    }

    public function setUserName($username)
    {
        $this->username = $username;
    }

    public function setEmail($email)
    {
        $this->email = $email;
    }
}