<?php
/**
 * Created by PhpStorm.
 * User: bg
 * Date: 05.03.17
 * Time: 01:28
 */

namespace AppBundle\DataAccess\Contract\Entities;


class Flags
{
    public static function flagIsSet($flag, $value) {
        return (bool)($flag == $flag & $value);
    }

    public static function setFlag($flag, $value) {
        return $flag | $value;
    }

    public static function unsetFlag($flag, $value) {
        return ~$flag & $value;
    }
}