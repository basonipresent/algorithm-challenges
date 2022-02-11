const { string, array } = require("yargs");

const palindromeOld = (str) => {
  let reverse = '';//str.toLowerCase().split('').reverse().join('');

  let strArray = str.toLowerCase().split('');

  for (let i = strArray.length; i > 0; i--) {
    reverse += strArray[i - 1];
  }

  if (reverse === str) {
    return reverse;
  } else {
    return 'not palindrome word';
  }
}

const palindromeNew = (str) => {
  if (typeof str === 'string' || typeof str === 'number') {
    if (typeof str === 'number') {
      str = new Number(str).toString();
    }
    //replace empty string and to lower case
    str = str.toLowerCase().replace(/\s/gm, "");

    let charLength = str.length;
    let charMid = Math.floor(charLength / 2);
    let i = 0;
    while (i < charMid) {
      if (str[i] !== str[charLength - (1 + i)]) {
        return false;
      }
      i++;
    }
  } else {
    throw TypeError(`was expected an argumentof type string or number, received an argument of type ${typeof str}`);
  }
  return true;
}

const oneSwapPalindrome = (n, letters) => {
  let diffCount = 0;
  let diff = [
    ['', ''],
    ['', '']
  ];
  let i = 0;
  for (i = 0; i < Math.floor(n / 2); i++) {
    if (letters[i] != letters[n - (i + 1)]) {
      if (diffCount == 2) {
        return false;
      }
      diff[diffCount][0] = letters[i];
      diff[diffCount++][1] = letters[n - (i + 1)];
    }
  }
  switch (diffCount) {
    case 0:
      return true;
    case 1:
      var midChar = letters[i];
      if (n % 2 != 0 && (diff[0][0] == midChar || diff[0][1] == midChar))
        return true;
    case 2:
      if ((diff[0][0] == diff[1][0] && diff[0][1] == diff[1][1])
        || (diff[0][0] == diff[1][1] && diff[0][1] == diff[1][0]))
        return true;
  }
  return false;
}

const countPalindrome = (n, str) => {
  let start = 0;
  let end = n - 1;
  let count = 0;
  let flag = true;
  let f_start = false;
  let f_end = false;
  while (start < end) {
    if (f_start) {
      if (str[start - 1] != str[end]) {
        if (str[start - 1] == str[end - 1] && start < end - 1) {
          count++; 
          f_end = true; 
          f_start = false;
        }
        else {
          flag = false; 
          break;
        }
      }
      else
        f_start = false;
    }
    else if (f_end) {
      if (str[start] != str[end + 1]) {
        if (str[start + 1] == str[end + 1] && start + 1 < end) {
          count++; 
          f_end = false; 
          f_start = true;
        }
        else {
          flag = false; 
          break;
        }
      }
      else
        f_end = false;
    }
    else if (str[start] != str[end]) {
      if (str[start + 1] == str[end] && start + 1 < end) {
        count++; f_start = true;
      }
      else if (str[start] == str[end - 1] && start < end - 1) {
        count++; f_end = true;
      }
      else {
        flag = false; break;
      }
    }
    start++;
    end--;
  }
  if (flag)
    return `YES ${count}`;
  else
    return `NO ${count}`;
}

console.log(palindromeNew('A Man a Plan a Canal Panama'));
console.log(palindromeNew(1234567654321));
console.log(oneSwapPalindrome(3, "bbg"));
console.log(oneSwapPalindrome(7, "bdababd"));
console.log(oneSwapPalindrome(6, "gcagac"));
console.log(oneSwapPalindrome(6, "gcagca"));
console.log(oneSwapPalindrome(21, "amanaplanacanalpanama"));
console.log(countPalindrome(7, "bdababd"));
console.log(countPalindrome(6, "maddma"));
console.log(countPalindrome(7, "bdababd"));
