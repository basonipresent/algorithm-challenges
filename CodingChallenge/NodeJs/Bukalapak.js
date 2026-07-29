let a = true
let b = true
let c = false

function result(a, b, c) {
    console.log(((a || c) && (!b || c)));
    console.log(!(a && b && !c));
    console.log((!b || c || (!a && b)) || c);
    console.log((!c && a) && !(b || c));
    console.log((a && b) || (!b && c));
}

console.log(result(true, true, false));