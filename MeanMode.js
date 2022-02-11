function mode(arr) {
    var c = {};
    for (var i = 0; i < arr.length; i++) {
        c[arr[i]] = (c[arr[i]] || 0) + 1;
    }
    var mode;
    var max = 0;
    for (e in c) {
        if (c[e] > max) {
            mode = e;
            max = c[e];
        }
    }
    return mode * 1;
}