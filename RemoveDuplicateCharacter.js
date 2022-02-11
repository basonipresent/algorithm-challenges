function solution(input) {
    var arrSort = input.split('').sort();
    let result = "";
    for (let i = 0; i < arrSort.length; i++) {
        if (arrSort[i - 1] == undefined) {
            continue;
        }
        if (arrSort[i - 1] == arrSort[i]) {
            result += arrSort[i];
        }
    }
    return removeDuplicate(result);
}

function removeDuplicate(input) {
    return input.split('').filter((item, pos, self) => {
        return self.indexOf(item) == pos;
    }).join('');
}

console.log(solution("BABBBCAAADED"));