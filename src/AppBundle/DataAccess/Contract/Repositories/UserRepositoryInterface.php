<?php

namespace AppBundle\DataAccess\Contract\Repositories;


use AppBundle\DataAccess\Contract\Entities\User;

interface UserRepositoryInterface
{
    /**
     * @return User[]
     */
    public function getUsers();

    /**
     * @param int $userId
     *
     * @return User
     */
    public function getUser($userId);
}