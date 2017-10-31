#! /usr/bin/env php
<?php

function debug($s) {
	foreach (str_split($s) as $b) {
		echo sprintf("%01s %02s %08d\n", ' ', bin2hex($b), sprintf('%08d', decbin(ord($b))));
	}
	echo "\n";
}

function save($file, $v) {
	file_put_contents(__DIR__ . '/demo/' . $file . '.bin', msgpack_pack($v));
}

function bin($s) {
	printf("%s = %3d, %s\n", $s, bindec($s), dechex(bindec($s)));
}

bin('11100000');
bin('11110000');
bin('11111000');
bin('11111100');
bin('11111110');
bin('11111111');

bin('01111111');
bin('00111111');
bin('00011111');
bin('00001111');

bin('00110111');
bin('00011111');
bin('00010111');

echo "\n";

echo 0b01110111 & 0b00011111, "\n";

// debug(json_encode(['abc']));

debug(msgpack_pack(1));
debug(msgpack_pack(-1));

save('string', ['abc' => 'def', 'a1' => 'a2', 'foo' => 'bar']);
save('int1', 1);
save('int-1', -1);
save('int128', 128);
save('int109', 109);
save('array1', [1,2,'a']);
save('arrays', ['a','b','c']);
save('multi', [
	[
		1,
		2,
		[
			3,
			4,
		],
	],
	[
		5,
		6,
	],
]);
