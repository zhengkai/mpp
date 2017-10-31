<?php
function debug($s) {
	foreach (str_split($s) as $b) {
		echo sprintf("%01s %02s %08d\n", ' ', bin2hex($b), sprintf('%08d', decbin(ord($b))));
	}
	echo "\n";
}

function save($file, $v) {
	$file = __DIR__ . '/demo/' . $file . '.bin';
	file_put_contents($file, msgpack_pack($v));
	echo $file, "\n";
}

function bin($s) {
	printf("%s = %3d, %s\n", $s, bindec($s), dechex(bindec($s)));
}
