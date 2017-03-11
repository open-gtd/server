<?php

namespace IOKI\PlatformBundle\FastEntities;

use Symfony\Component\Validator\Constraints\DateTime;

/**
 * Hydrators
 */
class Hydrators
{
    /**
     * @param $class
     * @param array $value
     *
     * @return mixed
     */
    public static function asEntity($class, array $value) {
        return new $class($value);
    }

    /**
     * @param $class
     * @param array $value
     *
     * @return mixed
     */
    public static function asEntityCollection($class, array $value) {
        $result = [];

        foreach ($value as $entity) {
            $result[] = self::asEntity($class, $entity);
        }

        return $result;
    }

    /**
     * @param $value
     *
     * @return string
     */
    public static function asString($value) {
        return (string)$value;
    }

    /**
     * @param $value
     *
     * @return int
     */
    public static function asInteger($value) {
        return (integer)$value;
    }

    /**
     * @param array $value
     *
     * @return array
     */
    public static function asArray(array $value) {
        return $value;
    }

    public static function asDateTime($value)
    {
        if ($value instanceof DateTime)
            return $value;

        return new DateTime($value);
    }
}