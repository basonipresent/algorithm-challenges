const rearrangeNumber = (arr, n) => {
  let j = 0;
  for (let i = 0; i < n; i++) {
    if (arr[i] < 0) {
      if (i != j) {
        let temp = arr[i];
        arr[i] = arr[j];
        arr[j] = temp;
      }
      j++;
    }
  }
  return arr;
}

let arr = [-12, 11, -13, -5, 6, -7, 5, -3, -6];
let n = arr.length;
console.log(rearrangeNumber(arr, n));