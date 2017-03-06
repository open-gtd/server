#!/usr/bin/env php
<?php

$composerFlag = '';
$npmFlag = '--production';
$bowerFlag = '--production';
if ($argv > 1 && $argv[1] == "--dev")
{
    $composerFlag = "--dev";
    $npmFlag = '';
    $bowerFlag = '';
}

passthru("composer install {$composerFlag}");
passthru("npm install {$npmFlag}");
passthru("bower install {$bowerFlag}");
passthru("grunt");