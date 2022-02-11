// 9,5,1,7,2
const merge = (left, right, type) => {
  let sortArr = [];
  while(left.length && right.length){
    if(type === 'descending'){
      if(left[0] > right[0]){
        sortArr.push(left.shift());
      } else {
        sortArr.push(right.shift());
      }
    } else {
      if(left[0] < right[0]){
        sortArr.push(left.shift());
      } else {
        sortArr.push(right.shift());
      }
    }
  }
  return [...sortArr, ...left, ...right];
}

const mergeSort = (arr, type) => {
  let middle = Math.floor(arr.length / 2);
  if(arr.length <= 1){
    return arr;
  } 
  let left = arr.splice(0, middle);
  let right = arr;
  return merge(mergeSort(left, type), mergeSort(right, type), type);
}

console.log(mergeSort([9, 5, 1, 7, 2, 10], 'descending'));