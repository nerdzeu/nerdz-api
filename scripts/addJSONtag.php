<?php
$r = file_get_contents("types.go");
$e = explode("\n", $r);

$out = '';
foreach($e as $line) {
    $out .= preg_replace_callback("#^(\s+\w+)(\s+.+?)(\s+`.+?`)?$#i", function($m) {
        if(strpos($m[1], 'return') !== false) {
            return $m[0];
        }
        $name = lcfirst(trim($m[1]));
        if(strlen($name) == 2) {
            $name = strtolower($name);
        }

        if(isset($m[3])) {
            return $m[1].$m[2].substr($m[3],0,-1).' json:"'.$name.'"`';
        } else {
            $comment = ' ';
            if(($pos = strpos($m[2], '//')) !== false) {
                $comment .= substr($m[2],$pos);
            }
            return rtrim($m[1].$m[2].' `json:"'.$name.'"`'.$comment);
        }
    }, $line);
    $out .= "\n";
}

echo $out;
?>
