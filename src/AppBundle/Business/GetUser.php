<?php


namespace AppBundle\Business;


use AppBundle\Business\Contract\Entities\User;
use AppBundle\Business\Contract\User\GetUserInterface;
use AppBundle\Business\Contract\User\UserListInterface;
use AppBundle\DataAccess\Contract\Entities\User as DbUser;
use AppBundle\DataAccess\Contract\Repositories\UserRepositoryInterface;

class GetUser implements GetUserInterface
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
     * {@inheritdoc}
     */
    public function getUser($userId)
    {
        $user = $this->userRepository->getUser($userId);
        return new User([
            User::USER_NAME => $user->getUsername(),
            User::EMAIL => $user->getEmail(),
            User::ENABLED => $user->isEnabled(),
            User::LAST_LOGIN => $user->getLastLogin(),
            User::ROLES => $user->getRoles()
        ]);
    }
}