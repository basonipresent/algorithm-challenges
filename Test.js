function solution(input){
    let lengthInput = input.length;
    let partLength = lengthInput / 2;

    input = input.sort().join('');
    let found = 0
    for(let i = 0; i < input.length; i++){
        
        if(input[i - 1] == undefined){
            continue
        }
        
        if(input[i - 1] == input[i]){
            found++
            if(found > partLength){
                return input[i]
            }
        }
    }
    return 0;
}

console.log(solution([2,2,1,3,3,3]));
