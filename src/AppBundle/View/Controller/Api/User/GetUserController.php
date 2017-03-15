<?php

namespace AppBundle\View\Controller\Api\User;


use AppBundle\Business\Contract\User\GetUserInterface;
use FOS\RestBundle\Controller\FOSRestController;
use Symfony\Component\DependencyInjection\ContainerInterface;
use Symfony\Component\HttpFoundation\Response;
use FOS\RestBundle\View\View;


class GetUserController extends FOSRestController
{
    /** @var GetUserInterface */
    private $getUser;

    /**
     * @param GetUserInterface $getUser
     * @param ContainerInterface $container
     */
    public function __construct(GetUserInterface $getUser, ContainerInterface $container)
    {
        $this->getUser = $getUser;
        $this->setContainer($container);
    }
    /**
     * @param int $userId
     *
     * @return mixed
     */
    public function run($userId)
    {
        $result = $this->getUser->getUser($userId);
        if ($result === null) {
            return new View("users does not exists", Response::HTTP_NOT_FOUND);
        }
        return $result;
    }
}