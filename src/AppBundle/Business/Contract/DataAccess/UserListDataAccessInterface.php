<?php

namespace AppBundle\Business\Contract\DataAccess;


use AppBundle\Business\Contract\Entities\User;

interface UserListDataAccessInterface
{
    /**
     * @return User[]
     */
    public function getUsers();
}