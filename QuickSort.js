const quickSort = (arr) => {
  if (arr.length <= 1) {
    return arr;
  }

  let pivot = arr[0];
  let left = [];
  let right = [];
  let lengthArr = arr.length;

  for (let i = 1; i < lengthArr; i++) {
    if (arr[i] < pivot) {
      left.push(arr[i]);
    } else {
      right.push(arr[i]);
    }
  }

  return [pivot, ...quickSort(left), ...quickSort(right)];
}

console.log(quickSort([12, 27, 45, 41, 55, 34, 63, 72]));