<?php

namespace AppBundle\DataAccess\Contract\Entities;

use Doctrine\Common\Collections\ArrayCollection;
use Doctrine\ORM\Mapping as ORM;
use Doctrine\ORM\Mapping\JoinColumn;
use Doctrine\ORM\Mapping\JoinTable;
use Doctrine\ORM\Mapping\ManyToMany;
use Doctrine\ORM\Mapping\OneToMany;
use Doctrine\ORM\Mapping\OneToOne;

/**
 * Task
 *
 * @ORM\Table(name="task")
 * @ORM\Entity(repositoryClass="AppBundle\DataAccess\Orm\TaskRepository")
 */
class Task
{
    /**
     * @var int
     *
     * @ORM\Column(name="id", type="integer")
     * @ORM\Id
     * @ORM\GeneratedValue(strategy="AUTO")
     */
    private $id;
    /**
     * @var string
     *
     * @ORM\Column(name="title", type="string", length=255)
     */
    private $title;

    /**
     * @var string
     *
     * @ORM\Column(name="description", type="string", length=4000)
     */
    private $description;

    /**
     * @var int
     *
     * @ORM\Column(name="state", type="integer")
     */
    private $state;

    /**
     * @var int
     *
     * @ORM\Column(name="star", type="integer")
     */
    private $star;

    /**
     * @var ArrayCollection
     *
     * @OneToMany(targetEntity="Checklist", mappedBy="task")
     */
    private $checklist;

    /**
     * @var Context
     *
     * @OneToOne(targetEntity="Context")
     * @JoinColumn(name="context_id", referencedColumnName="id")
     */
    private $context;

    /**
     * @var Project
     *
     * @OneToOne(targetEntity="Project")
     * @JoinColumn(name="project_id", referencedColumnName="id")
     */
    private $project;

    /**
     * @var ArrayCollection
     *
     * @ManyToMany(targetEntity="Tag")
     * @JoinTable(name="tasks_tags",
     *      joinColumns={@JoinColumn(name="task_id", referencedColumnName="id")},
     *      inverseJoinColumns={@JoinColumn(name="tag_id", referencedColumnName="id")}
     *      )
     */
    private $tags;

    /**
     * Task constructor.
     */
    public function __construct()
    {
        $this->checklist = new ArrayCollection();
        $this->tags = new ArrayCollection();
    }


    /**
     * Get id
     *
     * @return int
     */
    public function getId()
    {
        return $this->id;
    }

    /**
     * Set title
     *
     * @param string $title
     *
     * @return Task
     */
    public function setTitle($title)
    {
        $this->title = $title;

        return $this;
    }

    /**
     * Get title
     *
     * @return string
     */
    public function getTitle()
    {
        return $this->title;
    }

    /**
     * Set description
     *
     * @param string $description
     *
     * @return Task
     */
    public function setDescription($description)
    {
        $this->description = $description;

        return $this;
    }

    /**
     * Get description
     *
     * @return string
     */
    public function getDescription()
    {
        return $this->description;
    }

    /**
     * Set state
     *
     * @param integer $state
     *
     * @return Task
     */
    protected function setState($state)
    {
        $this->state = $state;

        return $this;
    }

    /**
     * Get state
     *
     * @return int
     */
    protected function getState()
    {
        return $this->state;
    }

    /**
     * @param bool $inbox
     *
     * @return Task
     */
    public function setInbox($inbox)
    {
        $inbox
            ? Flags::setFlag(TaskStates::INBOX, $this->state)
            : Flags::unsetFlag(TaskStates::INBOX, $this->state);

        return $this;
    }

    /**
     * @return bool
     */
    public function inInbox()
    {
        return Flags::flagIsSet(TaskStates::INBOX, $this->state);
    }

    /**
     * @param bool $actionable
     *
     * @return Task
     */
    public function setActionable($actionable)
    {
        $actionable
            ? Flags::setFlag(TaskStates::ACTIONABLE, $this->state)
            : Flags::unsetFlag(TaskStates::ACTIONABLE, $this->state);

        return $this;
    }

    /**
     * @return bool
     */
    public function isActionable()
    {
        return Flags::flagIsSet(TaskStates::ACTIONABLE, $this->state);
    }

    /**
     * @param bool $planning
     *
     * @return Task
     */
    public function setPlanning($planning)
    {
        $planning
            ? Flags::setFlag(TaskStates::PLANNING, $this->state)
            : Flags::unsetFlag(TaskStates::PLANNING, $this->state);

        return $this;
    }

    /**
     * @return bool
     */
    public function inPlanning()
    {
        return Flags::flagIsSet(TaskStates::PLANNING, $this->state);
    }

    /**
     * @param bool $incubated
     *
     * @return Task
     */
    public function setIncubated($incubated)
    {
        $incubated
            ? Flags::setFlag(TaskStates::INCUBATED, $this->state)
            : Flags::unsetFlag(TaskStates::INCUBATED, $this->state);

        return $this;
    }

    /**
     * @return bool
     */
    public function isIncubated()
    {
        return Flags::flagIsSet(TaskStates::INCUBATED, $this->state);
    }

    /**
     * @param bool $referenced
     *
     * @return Task
     */
    public function setReferenced($referenced)
    {
        $referenced
            ? Flags::setFlag(TaskStates::REFERENCED, $this->state)
            : Flags::unsetFlag(TaskStates::REFERENCED, $this->state);

        return $this;
    }

    /**
     * @return bool
     */
    public function isReferenced()
    {
        return Flags::flagIsSet(TaskStates::REFERENCED, $this->state);
    }

    /**
     * @param bool $delegated
     *
     * @return Task
     */
    public function setDelegated($delegated)
    {
        $delegated
            ? Flags::setFlag(TaskStates::DELEGATED, $this->state)
            : Flags::unsetFlag(TaskStates::DELEGATED, $this->state);

        return $this;
    }

    /**
     * @return bool
     */
    public function isDelegated()
    {
        return Flags::flagIsSet(TaskStates::DELEGATED, $this->state);
    }

    /**
     * @param bool $next
     *
     * @return Task
     */
    public function setNext($next)
    {
        $next
            ? Flags::setFlag(TaskStates::NEXT, $this->state)
            : Flags::unsetFlag(TaskStates::NEXT, $this->state);

        return $this;
    }

    /**
     * @return bool
     */
    public function isNext()
    {
        return Flags::flagIsSet(TaskStates::NEXT, $this->state);
    }

    /**
     * @return ArrayCollection
     */
    public function getChecklist()
    {
        return $this->checklist;
    }

    /**
     * @param ArrayCollection $checklist
     *
     * @return Task
     */
    public function setChecklist($checklist)
    {
        $this->checklist = $checklist;

        return $this;
    }

    /**
     * @return Context
     */
    public function getContext()
    {
        return $this->context;
    }

    /**
     * @param Context $context
     *
     * @return Task
     */
    public function setContext($context)
    {
        $this->context = $context;

        return $this;
    }

    /**
     * @return Project
     */
    public function getProject()
    {
        return $this->project;
    }

    /**
     * @param Project $project
     *
     * @return Task
     */
    public function setProject($project)
    {
        $this->project = $project;

        return $this;
    }

    /**
     * @return ArrayCollection
     */
    public function getTags()
    {
        return $this->tags;
    }

    /**
     * @param ArrayCollection $tags
     *
     * @return $this
     */
    public function setTags($tags)
    {
        $this->tags = $tags;

        return $this;
    }

    /**
     * @return int
     */
    public function getStar()
    {
        return $this->star;
    }

    /**
     * @param int $star
     *
     * @return Task
     */
    public function setStar($star)
    {
        $this->star = $star;

        return $this;
    }
}

