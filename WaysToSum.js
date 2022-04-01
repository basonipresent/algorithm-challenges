function countWays(n, k) {
  var count = new Array(n + 1).fill(0);
  var keys = [];

  for (let i = 1; i <= k; i++) {
    keys.push(i)
  }

  count[0] = count[1] = count[2] = 1;
  count[3] = 2;

  keys.forEach(element => {
    //console.log(element)
  })

  keys.forEach(element => {
    element++
    for (let i = 4; i <= n; i++) {
      count[i] += count[i - element]
    }
  })

  return count[n];
}

console.log(NumberOfways(842, 91))

function ways(total, k) {
  // Write your code here
  let min = 1 
  if (total < 0) return 0
  else if (total == 0) return min
  else if (k < min) return 0
  else
      return ways(total - k, k) + ways(total, k - min)
}

function NumberOfways(N, K)
{
     
    // Initialize a list
    let dp = Array.from({length: N +1}, (_, i) => 0);
   
    // Update dp[0] to 1
    dp[0] = 1;
 
    // Iterate over the range [1, K + 1]
    for(let row = 1; row < K + 1; row++)
    {
 
        // Iterate over the range [1, N + 1]
        for(let col = 1; col < N + 1; col++)
        {
             
            // If col is greater
            // than or equal to row
            if (col >= row)
               
                // Update current
                // dp[col] state
                dp[col] = dp[col] + dp[col - row];
          }
    }
 
    // Return the total number of ways
    return dp[N];
}

static int NumberOfways(int N, int K)
{
     
    // Initialize a list
    int[] dp = new int[N + 1];
   
    // Update dp[0] to 1
    dp[0] = 1;
 
    // Iterate over the range [1, K + 1]
    for(int row = 1; row < K + 1; row++)
    {
 
        // Iterate over the range [1, N + 1]
        for(int col = 1; col < N + 1; col++)
        {
             
            // If col is greater
            // than or equal to row
            if (col >= row)
               
                // Update current
                // dp[col] state
                dp[col] = dp[col] + dp[col - row];
          }
    }
 
    // Return the total number of ways
    return(dp[N]);
}