#! /usr/bin/env php
<?php
require __DIR__ . '/common.inc.php';

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
