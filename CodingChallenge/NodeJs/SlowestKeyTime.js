function slowestKey(keyTimes) {
  // Write your code here
  var slowestAscii = 0
  var previousIndex = 0
  var previousItem = []
  var maxTimeArr = []
  const asciiDict = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z']

  for (let i = 0; i < keyTimes.length; i++) {
    if (i == 0) {
      previousIndex = 0
      previousItem = [0, 0]
    } else {
      previousIndex = i - 1
      previousItem = keyTimes[previousIndex]
    }

    var currentArr = keyTimes[i]
    var diffTime = [keyTimes[i][0], currentArr[1] - previousItem[1]]
    maxTimeArr.push(diffTime)
    // slowestAscii = maxTimeArr[0]
  }

  return maxTimeArr.reverse(1) // asciiDict[slowestAscii]
}

console.log(slowestKey([[0, 1], [0, 3], [4, 5], [5, 6], [4, 10]]))

import Foundation



/*
 * Complete the 'slowestKey' function below.
 *
 * The function is expected to return a CHARACTER.
 * The function accepts 2D_INTEGER_ARRAY keyTimes as parameter.
 */

func slowestKey(keyTimes: [[Int]]) -> Character {
    // Write your code here
    let str = "abcdefghijklmnopqrstuvwxyz"
    let alphabet = Array(str);
    
    var tStart = 0;
    var maxT = (0, "")
    
    keyTimes.map{
        let T = $0[1] - tStart
        let c = alphabet[$0[0]]
        if T > maxT.0 {maxT = (T, String(c))}
        tStart = $0[1]
    }
    return Character(String(maxT.1))
}

let stdout = ProcessInfo.processInfo.environment["OUTPUT_PATH"]!