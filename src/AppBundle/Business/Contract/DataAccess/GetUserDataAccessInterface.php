<?php
namespace AppBundle\Business\Contract\DataAccess;


use AppBundle\DataAccess\Contract\Entities\User;


interface GetUserDataAccessInterface
{
    /**
     * @param string $userName
     *
     * @return User
     */
    public function getUser($userName);
}