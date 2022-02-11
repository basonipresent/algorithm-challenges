// you can write to stdout for debugging purposes, e.g.
// console.log('this is a debug message');

function solution(A) {
  // write your code in JavaScript (Node.js 8.9.4)
  if (isAesthetic(A)) {
    return 0;
  }

  let aestheticPatternCount = 0;
  for (let j = 0; j < A.length; j++) {
    let newA = copyArrWithoutElement(A, j);
    if (isAesthetic(newA)) {
      aestheticPatternCount++;
    }
  }

  if (aestheticPatternCount == 0) {
    return -1;
  } else {
    return aestheticPatternCount;
  }
}

function copyArrWithoutElement(arr, index) {
  let arrayLength = arr.length;
  var newArr = new Array([arrayLength - 1]);
  let tempK = 0;
  for (let k = 0; k < arrayLength; k++) {
    if (k != index) {
      newArr[tempK++] = arr[k];
    }
  }
  return newArr;
}

function isAesthetic(arr) {
  let arrayLength = arr.length;
  let increastFlag = 0;
  for (let i = 0; i < arrayLength; i++) {
    if (increastFlag == 0) {
      if (arr[i] < arr[i + 1]) {
        increastFlag = 1;
      } else {
        increastFlag = 2;
      }
    } else {
      if (increastFlag == 1) {
        if (i % 2 == 1 && arr[i] > arr[i - 1]) {
        } else if (i % 2 == 0 && arr[i] < arr[i - 1]) {
        } else {
          return false;
        }
      } else {
        if (i % 2 == 1 && arr[i] < arr[i - 1]) {
          // do nothing
        } else if (i % 2 == 0 && arr[i] > arr[i - 1]) {
          // do nothing
        } else {
          return false;
        }
      }
    }
  }
  return true;
}

console.log(solution([1,2,3,4]))