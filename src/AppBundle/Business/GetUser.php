<?php


namespace AppBundle\Business;


use AppBundle\Business\Contract\DataAccess\GetUserDataAccessInterface;
use AppBundle\Business\Contract\User\GetUserInterface;

class GetUser implements GetUserInterface
{
    private $getUserDataAccess;

    /**
     * @param $getUserDataAccess
     */
    public function __construct(GetUserDataAccessInterface $getUserDataAccess)
    {
        $this->getUserDataAccess = $getUserDataAccess;
    }

    /**
     * {@inheritdoc}
     */
    public function getUser($userName)
    {
        return $this->getUserDataAccess->getUser($userName);
    }
}