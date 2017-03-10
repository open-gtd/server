<?php

namespace AppBundle\Presentation\Controller\User;

use AppBundle\Business\Contract\User\UserListInterface;
use Symfony\Bundle\FrameworkBundle\Controller\Controller;
use Symfony\Component\DependencyInjection\ContainerInterface;

/**
 * User controller.
 */
class UserIndexController extends Controller
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
     * @return \Symfony\Component\HttpFoundation\Response
     */
    public function run()
    {
        $users = $this->userList->getAllUsers();

        return $this->render('user/index.html.twig', array(
            'users' => $users,
        ));
    }
}
