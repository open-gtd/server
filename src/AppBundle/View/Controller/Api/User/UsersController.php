<?php

namespace AppBundle\View\Controller\Api\User;


use AppBundle\Business\Contract\User\UserListInterface;
use FOS\RestBundle\Controller\Annotations\QueryParam;
use FOS\RestBundle\Controller\FOSRestController;
use FOS\RestBundle\Request\ParamFetcherInterface;
use Symfony\Component\DependencyInjection\ContainerInterface;
use Symfony\Component\HttpFoundation\Response;


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
     * @return Response
     */
    public function run()//ParamFetcherInterface $paramFetcher)
    {
        $data = $this->userList->getAllUsers();
        $view = $this->view($data, 200)
            ->setTemplate("MyBundle:Users:getUsers.html.twig")
            ->setTemplateVar('users');

        return $this->handleView($view);
    }
}