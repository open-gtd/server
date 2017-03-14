<?php

namespace AppBundle\View\Controller\Api\User;


use AppBundle\Business\Contract\User\UserListInterface;
use FOS\RestBundle\Controller\FOSRestController;
use Symfony\Component\DependencyInjection\ContainerInterface;
use Symfony\Component\HttpFoundation\Response;
use FOS\RestBundle\View\View;


class UsersController extends FOSRestController
{
    /** @var UserListInterface */
    private $userList;

    /**
     * @param UserListInterface $userList
     * @param ContainerInterface $container
     */
    public function __construct(UserListInterface $userList, ContainerInterface $container)
    {
        $this->userList = $userList;
        $this->setContainer($container);
    }
    /**
     * Get the list of users.
     *
//     * @param ParamFetcherInterface $paramFetcher
     *
     * @return mixed
     */
    public function run()//ParamFetcherInterface $paramFetcher)
    {
        $result = $this->userList->getAllUsers();
        if ($result === null) {
            return new View("there are no users exist", Response::HTTP_NOT_FOUND);
        }
        return $result;
    }
}