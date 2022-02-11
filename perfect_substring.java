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
     * Complete the 'perfectSubstring' function below.
     *
     * The function is expected to return an INTEGER.
     * The function accepts following parameters:
     *  1. STRING s
     *  2. INTEGER k
     */
     
    static int MAX_CHAR = 10;
    
    static boolean check(int freq[], int k){
        for(int i = 0; i < MAX_CHAR; i++){
            if(freq[i] != 0 && freq[i] != k){
                return false;
            }
        }
        return true;
    }

    public static int perfectSubstring(String s, int k) {
        // Write your code here
        int res = 0;
        for(int i = 0; i < s.length(); i++){
            int freq[] = new int[MAX_CHAR];
            for(int j = i; j < s.length(); j++){
                int index = s.charAt(j) - '0';
                freq[index]++;
                
                if(freq[index] > k){
                    break;
                }
                else if (freq[index] == k && check(freq, k) == true){
                    res++;
                }
            }
        }
        return res;
    }
}
public class Solution {
    public static void main(String[] args) throws IOException {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        String s = bufferedReader.readLine();

        int k = Integer.parseInt(bufferedReader.readLine().trim());

        int result = Result.perfectSubstring(s, k);

        bufferedWriter.write(String.valueOf(result));
        bufferedWriter.newLine();

        bufferedReader.close();
        bufferedWriter.close();
    }
}
