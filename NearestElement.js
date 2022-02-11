const nearestElement = (arr, K) => {
  let res = arr[0];
  let N = arr.length;

  // Traverse the array
  for (let i = 1; i < N; i++) {

    // If absolute difference
    // of K and res exceeds
    // absolute difference of K
    // and current element
    if (Math.abs(K - res) >
      Math.abs(K - arr[i])) {
      res = arr[i];
    }
  }

  // Return the closest
  // array element
  return res;
}

let arr = [4, 2, 8, 11, 7];
let K = 12;
console.log(nearestElement(arr, K));