<?php

$startTime = microtime(true);

$file = fopen('php-test.csv', 'w');

for ($i=0; $i < 9000000; $i++) {
fputcsv(
    $file, 
    [
        $i,
        'Bear',
        'Honey',
        'Snow'
    ]
    );
}

$execution_time = (microtime(true) - $startTime);
echo "Processing time is: ".$execution_time*1000;