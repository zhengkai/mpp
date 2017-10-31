#! /usr/bin/env php
<?php
require __DIR__ . '/common.inc.php';

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
