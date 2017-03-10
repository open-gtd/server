<?php

namespace IOKI\PlatformBundle\FastEntities;


trait Entity
{
    /** @var array */
    private $fields = [];

    /** @var array */
    private $cache = [];

    /** @var array */
    protected $filter = [];

    /**
     * Entity constructor.
     * @param array $fields
     */
    public function init(array $fields)
    {
        foreach ($fields as $field => $value) {
            if (isset($this->filter[$field])) {
                $this->fields[$field] = $value;
            }
        }
    }

    public function __get($name)
    {
        if (!isset($this->filter[$name])) {
            throw new \Exception('Invalid property');
        }

        if (!isset($this->cache[$name])) {
            $this->cache[$name] = $this->getFieldValue($name);
        }

        return $this->cache[$name];
    }

    /**
     * @param $name
     * @return mixed
     */
    protected function getFieldValue($name)
    {
        if (!isset($this->fields[$name])) {
            return null;
        }

        $hydration = $this->filter[$name];
        $value = $hydration($this->fields[$name]);

        return $value;
    }

    protected function toArray(array $filter)
    {
        $result = [];

        foreach ($filter as $field) {
            if (isset($this->filter[$field])) {
                $result[$field] = $this->{$field};
            }
        }

        return $result;
    }
}