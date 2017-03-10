<?php


namespace AppBundle\Business\Contract\User;


use AppBundle\Business\Contract\Entities\User;

interface UserListInterface
{
    /**
     * @return User[]
     */
    public function getAllUsers();
}