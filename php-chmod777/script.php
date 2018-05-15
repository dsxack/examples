<?php
/**
 * @author Smotrov Dmitriy <smotrov@worksolutions.ru>
 */

$iterator = new RecursiveIteratorIterator(
    new RecursiveDirectoryIterator(__DIR__),
    RecursiveIteratorIterator::SELF_FIRST
);

foreach ($iterator as $path) {
    /** @var SplFileInfo $path */
    chmod($path->getPathname(), 0777);
}

echo 'ok';
