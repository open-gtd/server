<?php
namespace AppBundle\DataAccess\Business\Factories;


use AppBundle\Business\Contract\Entities\User;
use AppBundle\DataAccess\Contract\Entities\User as DbUser;


class UserFactory
{

    /**
     * @param DbUser $dbUser

     * @return User
     */
    public static function fromUser(DbUser $dbUser)
    {
        $user = new User();

        $user->setUserName($dbUser->getUsername());
        $user->setEmail($dbUser->getEmail());

        return $user;
    }


    /**
     * @param DbUser[] $users
     *
     * @return User[]
     */
    public static function fromUsers(array $users)
    {
        $mapFunc = function (DbUser $item) {
            return self::fromUser($item);
        };

        return array_map($mapFunc, $users);
    }
}