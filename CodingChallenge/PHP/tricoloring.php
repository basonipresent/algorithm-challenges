function solution($arr) {
    $sum = array_sum($arr);
    if ($sum % 3) return "impossible";
    $division = $sum / 3;

    function recurse(&$arr, &$result, $division, $color, $sum, $i) {
        if ($sum === $division) {
            if ($color === "G") return true;
            $first = array_search("R", $result);
            return recurse($arr, $result, $division, "G", 0, $first+1);
        }
        if ($i >= count($arr)) return false;
        $success = recurse($arr, $result, $division, $color, $sum, $i+1);
        if ($success) return true;
        if ($result[$i] === "B") {
            $result[$i] = $color;
            $success = recurse($arr, $result, $division, $color, $sum + $arr[$i], $i+1);
            if ($success) return true;
            $result[$i] = "B";
        }
        return false;
    }

    $result = array_fill(0, count($arr), "B"); 
    if ($division === 0) return $result;

    $success = recurse($arr, $result, $division, "R", 0, 1);
    return $success ? implode($result, "") : "impossible";
}