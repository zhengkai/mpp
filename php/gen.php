#! /usr/bin/env php
<?php
require __DIR__ . '/common.inc.php';

$s = implode('', range('a', 'z')) . implode('', range(0, 9));
$s = str_repeat($s, ceil(65537 / strlen($s)));
foreach ([1, 31, 32, 33, 65535, 65536, 65537] as $i) {
	save('str/len-' . $i, substr($s, 0, $i));
}

$list = [128];
foreach (range(0, 6) as $i) {
	$i = pow(2, pow(2, $i));
	$i = min($i, PHP_INT_MAX);
	$list[] = $i;
}

foreach ($list as $i) {
	foreach ([-1, 0, 1] as $mod) {
		$j = $i + $mod;
		if (!is_integer($j) || $j < 0 || $j > PHP_INT_MAX) {
			continue;
		}
		save('int/n' . $j, $j);
	}
}

save('float-pi', M_PI);

save('string', ['abc' => 'def', 'a1' => 'a2', 'foo' => 'bar']);
save('int1', 1);
save('int-1', -1);
save('int128', 128);
save('int109', 109);
save('array-a', [1,2,'a']);
save('array-b', ['a','b',['x','y'],'c']);
save('array-c', ['a','b',['foo' => 'x', 'bar' => 'y'],'c']);
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

foreach (range(1, 10) as $i) {
	$a = makeRandomArray();
	save('json/' . $i . '.json', json_encode($a, JSON_UNESCAPED_SLASHES | JSON_UNESCAPED_UNICODE), false);
	save('json/' . $i . '.jsonf', json_encode($a, JSON_PRETTY_PRINT | JSON_UNESCAPED_SLASHES | JSON_UNESCAPED_UNICODE), false);
	save('json/' . $i, $a);
}
