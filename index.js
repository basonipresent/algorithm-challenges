
const merge = (arr1, arr2) => {
  return [...arr1, ...arr2];
}

const mergeSort = (arr) => {
  const middle = Math.floor(arr.length / 2);
  if(arr.length <= 1){
    return arr;
  }

  const left = arr.splice(0, middle);
  const right = arr;
  return _sortMerge(mergeSort(left), mergeSort(right));
}

const _sortMerge = (left, right) => {
  const sortArr = [];
  while(left.length && right.length){
    if(left[0] > right[0]){
      sortArr.push(left.shift());
    } else {
      sortArr.push(right.shift());
    }
  }
  return [...sortArr, ...left, ...right];
}

//fibonanci
const fibonanciGenerator = (index) => { 
  let result = [];
  result[0] = 0;
  result[1] = 1;

  for(i = 2; i <= index; i++){
    result[i] = result[i - 2] + result[i - 1];
  }

  console.log(result);
  return result[index];
}

const swap = (a, b) => {
  a = a + b;
  b = a - b;
  a = a - b;
  console.log(`a = ${a}`);
  console.log(`b = ${b}`);
  return '';
}

console.log(swap(5, 6))
