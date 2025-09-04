<?php

$start = time();
for ($i = 0; $i < 1000000000; $i++) {
    continue;
}
$elapsed = time() - $start;
echo "Elapsed time: $elapsed seconds\n";