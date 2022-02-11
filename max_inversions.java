import java.io.*;
import java.math.*;
import java.security.*;
import java.text.*;
import java.util.*;
import java.util.concurrent.*;
import java.util.function.*;
import java.util.regex.*;
import java.util.stream.*;
import static java.util.stream.Collectors.joining;
import static java.util.stream.Collectors.toList;



class Result {

    /*
     * Complete the 'maxInversions' function below.
     *
     * The function is expected to return a LONG_INTEGER.
     * The function accepts INTEGER_ARRAY arr as parameter.
     */

    public static long maxInversions(List<Integer> arr) {
        // Write your code here
        long inversions = 0;
        int n = arr.size();
        
        for(int i = 0; i < n - 1; i++){
            int smaller = 0;
            for(int j = i + 1; j < n; j++){
                if(arr.get(i) > arr.get(j)){
                    smaller++;
                }
            }
            int greater = 0;
            for(int j = i-1; j >=0; j--){
                if(arr.get(i) < arr.get(j)) {
                    greater++;
                }
            }
            inversions += greater * smaller;
        }
        
        return inversions;
    }

}

public class Solution {