<?php


namespace AppBundle\Business;


use AppBundle\Business\Contract\DataAccess\UserListDataAccessInterface;
use AppBundle\Business\Contract\Entities\User;
use AppBundle\Business\Contract\User\UserListInterface;

class UserList implements UserListInterface
{
    private $userListDataAccess;

    /**
     * @param $userListDataAccess
     */
    public function __construct(UserListDataAccessInterface $userListDataAccess)
    {
        $this->userListDataAccess = $userListDataAccess;
    }

    /**
     * @return User[]
     */
    public function getAllUsers()
    {
        return $this->userListDataAccess->getUsers();
    }
}