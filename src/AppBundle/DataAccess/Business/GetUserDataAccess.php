<?php
namespace AppBundle\DataAccess\Business;

use AppBundle\Business\Contract\DataAccess\GetUserDataAccessInterface;
use AppBundle\DataAccess\Business\Factories\UserFactory;
use AppBundle\DataAccess\Contract\Repositories\UserRepositoryInterface;

class GetUserDataAccess implements GetUserDataAccessInterface
{
    /** @var  UserRepositoryInterface */
    protected $userRepository;

    /**
     * @param UserRepositoryInterface $userRepository
     */
    public function __construct(UserRepositoryInterface $userRepository)
    {
        $this->userRepository = $userRepository;
    }

    /**
     * {@inheritdoc}
     */
    public function getUser($userName)
    {
        return UserFactory::fromUser($this->userRepository->getUserByName($userName));
    }
}