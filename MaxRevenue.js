function maxRevenue(arr) {
  let max = 0;

  for(let i = 0; i < arr.length; i++){
    let j = i + 1;
    while(j < arr.length){
      if(max > parseInt(arr[j] - arr[i])){
        max = parseInt(max);
      } else {
        max = parseInt(arr[j] - arr[i]);
      }
      j++;
    }
  }

  return max
}

console.log(maxRevenue([3,5,1,7,2,9]));
console.log(maxRevenue([2,5,3,7,8]));
console.log(maxRevenue([3,9,1,3,4,6]));
console.log(maxRevenue([9,8,7,6,5]));
console.log(maxRevenue([9,9,9,9]));
console.log(maxRevenue([9,9,9,7]));