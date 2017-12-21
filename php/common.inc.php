<?php
define('TYPE_INT', 1);
define('TYPE_STR', 2);
define('TYPE_NULL', 3);
define('TYPE_BOOL', 4);
define('TYPE_FLOAT', 5);

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

	$l = range(1, mt_rand(5, 8));

	foreach ($l as $i) {

		$max = 5;
		if ($depth < 4) {
			$max += 1;
		}

		$type = mt_rand(1, $max);

		switch ($type) {

		case TYPE_INT:
			$i = mt_rand(0, PHP_INT_MAX);
			$i = substr($i, 0, mt_rand(1, 12));
			$r[] = intval($i);
			break;

		case TYPE_STR:
			$r[] = makeRandStr();
			break;

		case TYPE_NULL:
			$r[] = NULL;
			break;

		case TYPE_BOOL:
			$r[] = mt_rand(0, 1) === 1;
			break;

		case TYPE_FLOAT:
			$r[] = mt_rand(1, 1000000000) / mt_rand(1, 1000000000);
			break;

		default:
			$r[] = makeRandomArray($depth + 1);
			break;
		}
	}

	if (mt_rand(0, 1)) {
		$map = [];
		foreach ($r as $row) {
			$map[makeRandStr()] = $row;
		}
		$r = $map;
	}

	return $r;
}

function makeRandStr() {
	$s = base64_encode(random_bytes(50));
	$s = substr($s, 0, mt_rand(1, 20));
	$s = str_replace('=', '', $s);
	$s = str_replace('a', '"', $s);
	$s = str_replace('b', '\\', $s);
	return $s;
}
