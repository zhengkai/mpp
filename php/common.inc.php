<?php
define('TYPE_INT', 1);
define('TYPE_STR', 2);
define('TYPE_ARRAY', 3);

function debug($s) {
	foreach (str_split($s) as $b) {
		echo sprintf("%01s %02s %08d\n", ' ', bin2hex($b), sprintf('%08d', decbin(ord($b))));
	}
	echo "\n";
}

function save($file, $v, $msgpack = true) {
	$file = __DIR__ . '/demo/' . $file . ($msgpack ? '.bin' : '');

	$dir = dirname($file);
	if (!file_exists($dir)) {
		mkdir($dir, 0777, TRUE);
	}

	if ($msgpack) {
		$v = msgpack_pack($v);
	}
	file_put_contents($file, $v);
	echo $file, "\n";
}

function bin($s) {
	printf("%s = %3d, %s\n", $s, bindec($s), dechex(bindec($s)));
}

function makeRandomArray($depth = 0) {

	$r = [];

	$l = range(1, mt_rand(3, 5));

	foreach ($l as $i) {

		if ($depth < 5) {
			$type = mt_rand(1, 3);
		} else {
			$type = mt_rand(1, 2);
		}

		switch ($type) {

		case TYPE_INT:
			$r[] = mt_rand(1, 1000000);
			break;

		case TYPE_STR:
			$r[] = base64_encode(random_bytes(mt_rand(1, 10)));
			break;

		case TYPE_ARRAY:
			$r[] = makeRandomArray($depth + 1);
			break;
		}
	}

	return $r;
}
