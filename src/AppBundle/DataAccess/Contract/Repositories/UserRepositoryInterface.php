<?php

namespace AppBundle\DataAccess\Contract\Repositories;


use AppBundle\DataAccess\Contract\Entities\User;

interface UserRepositoryInterface
{
    /**
     * @return User[]
     */
    public function getUsers();

    public function getUserByName($userName);
}