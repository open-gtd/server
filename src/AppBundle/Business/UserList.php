<?php


namespace AppBundle\Business\Contract\User;


use AppBundle\Business\Contract\Entities\User;
use AppBundle\DataAccess\Contract\Entities\User as DbUser;
use AppBundle\DataAccess\Contract\Repositories\UserRepositoryInterface;

class UserList implements UserListInterface
{
    private $userRepository;

    /**
     * @param $userRepository
     */
    public function __construct(UserRepositoryInterface $userRepository)
    {
        $this->userRepository = $userRepository;
    }

    /**
     * @return User[]
     */
    public function getAllUsers()
    {
        $result = [];

        /** @var DbUser[] $allUsers */
        $allUsers = $this->userRepository->getUsers();
        foreach ($allUsers as $user) {
            $result[] = new User([
                User::USER_NAME => $user->getUsername(),
                User::EMAIL => $user->getEmail(),
                User::ENABLED => $user->isEnabled(),
                User::LAST_LOGIN => $user->getLastLogin(),
                User::ROLES => $user->getRoles()
            ]);
        }

        return $result;
    }
}