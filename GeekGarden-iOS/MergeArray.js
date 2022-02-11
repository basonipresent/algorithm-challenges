function merge(left, right) {
  let sortArr = [];
  while (left.length && right.length) {
    if (left[0] < right[0]) {
      sortArr.push(left.shift());
    } else {
      sortArr.push(right.shift());
    }
  }
  return [...sortArr, ...left, ...right];
}

function mergeSort(arr) {
  let middle = Math.floor(arr.length / 2);
  if (arr.length <= 1) {
    return arr;
  }
  let left = arr.splice(0, middle);
  let right = arr;
  return merge(mergeSort(left), mergeSort(right));
}

function concatArraySort(A, B) {
  arr = [...A, ...B]
  return mergeSort(arr)
}

// normal test
console.log(concatArraySort([1, 2, 5, 7], [3, 4, 8]));
// null array test
console.log(concatArraySort([], [3, 4, 8]));
// double digit test
console.log(concatArraySort([11, 11, 23, 56, 79], [31, 10, 12, 800]));