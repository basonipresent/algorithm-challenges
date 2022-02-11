function countWays(n, k) {
  var count = new Array(n + 1).fill(0);
  var keys = [];

  for (let i = 1; i <= k; i++) {
    keys.push(i)
  }

  count[0] = count[1] = count[2] = 1;
  count[3] = 2;

  keys.forEach(element => {
    console.log(element)
  })

  keys.forEach(element => {
    element++
    for (let i = 4; i <= n; i++) {
      count[i] += count[i - element]
    }
  })

  return count[n];
}

console.log(countWays(56, 23))

function ways(total, k) {
  // Write your code here
  let min = 1 
  if (total < 0) return 0
  else if (total == 0) return min
  else if (k < min) return 0
  else
      return ways(total - k, k) + ways(total, k - min)
}