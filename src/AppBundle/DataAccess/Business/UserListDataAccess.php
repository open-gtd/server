<?php
namespace AppBundle\DataAccess\Business;


use AppBundle\Business\Contract\DataAccess\UserListDataAccessInterface;
use AppBundle\DataAccess\Business\Factories\UserFactory;
use AppBundle\DataAccess\Contract\Repositories\UserRepositoryInterface;

class UserListDataAccess implements UserListDataAccessInterface
{
    /** @var UserRepositoryInterface */
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
    public function getUsers()
    {
        return UserFactory::fromUsers($this->userRepository->getUsers());
    }
}