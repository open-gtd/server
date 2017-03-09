<?php
/**
 * Created by PhpStorm.
 * User: bg
 * Date: 05.03.17
 * Time: 01:21
 */

namespace AppBundle\DataAccess\Contract\Entities;


class TaskStates
{
    const INBOX = 0;
    const ACTIONABLE = 1;
    const PLANNING = 2;
    const INCUBATED = 4;
    const REFERENCED = 8;
    const PLANNED = 16;
    const DELEGATED = 32;
    const NEXT = 64;
}