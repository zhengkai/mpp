#! /usr/bin/env php
<?php

debug(msgpack_pack([12]));

function debug($s) {
	foreach (str_split($s) as $b) {
		echo sprintf("%01s %02s %08d", $b,  bin2hex($b), sprintf('%08d', decbin($b)));
	}
}
