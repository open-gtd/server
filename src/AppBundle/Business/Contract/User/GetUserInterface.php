<?php


namespace AppBundle\Business\Contract\User;


use AppBundle\Business\Contract\Entities\User;

interface GetUserInterface
{
    /**
     * @param int $userId
     *
     * @return User
     */
    public function getUser($userId);
}