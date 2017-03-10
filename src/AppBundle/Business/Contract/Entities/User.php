<?php


namespace AppBundle\Business\Contract\Entities;


use IOKI\PlatformBundle\FastEntities\Entity;
use IOKI\PlatformBundle\FastEntities\Hydrators;

/**
 * User
 *
 * @property string username
 */
class User
{
    use Entity;

    const USER_NAME = 'username';
    const EMAIL = 'email';
    const ENABLED = 'enabled';
    const LAST_LOGIN = 'last_login';
    const ROLES = 'roles';

    /**
     */
    public function __construct(array $fields)
    {
        $this->filter = [
            self::USER_NAME => function($value) { return Hydrators::asString($value); },
            self::EMAIL => function($value) { return Hydrators::asString($value); },
            self::ENABLED => function($value) { return Hydrators::asString($value); },
            self::LAST_LOGIN => function($value) { return Hydrators::asDateTime($value); },
            self::ROLES => function($value) { return Hydrators::asArray($value); }
        ];
        $this->init($fields);
    }
}