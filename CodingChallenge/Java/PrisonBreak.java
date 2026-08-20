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



class PrisonBreakResult {

    /*
     * Complete the 'prison' function below.
     *
     * The function is expected to return a LONG_INTEGER.
     * The function accepts following parameters:
     *  1. INTEGER n
     *  2. INTEGER m
     *  3. INTEGER_ARRAY h
     *  4. INTEGER_ARRAY v
     */

    public static long prison(int n, int m, List<Integer> h, List<Integer> v) {
    // Write your code here
        ArrayList<ArrayList<Long>> prison = new ArrayList<ArrayList<Long>>();
        for (int r = 0; r <= n; r++){
            ArrayList<Long> temp = new ArrayList<Long>();
            for(int c = 0; c <= m; c++){
                temp.add((long) 1);
            }
            prison.add(temp);
        }
        
        List<Integer> x = h;
        int xnum = x.size();
        Collections.sort(x);
        
        List<Integer> y = v;
        int ynum = y.size();
        Collections.sort(y);
        
        for(int a = xnum - 1; a >= 0; a--){
            int i = x.get(a);
            for(int cell = 0; cell < prison.get(i).size(); cell++) {
                prison.get(i).set(cell, prison.get(i).get(cell) + prison.get(i-1).get(cell));
            }
            prison.remove(i - 1);
        }
        
        ArrayList<ArrayList<Long>> newPrison = new ArrayList<ArrayList<Long>>();
        
        for(int col = 0; col < prison.get(0).size(); col++){
            ArrayList<Long> temp = new ArrayList<Long>();
            for(int row = 0; row < prison.size(); row++){
                temp.add(prison.get(row).get(col));
            }
            newPrison.add(temp);
        }
        
        for(int b = ynum - 1; b >= 0; b--){
            int i = y.get(b);
            for(int cell = 0; cell < newPrison.get(i).size(); cell++){
                newPrison.get(i).set(cell, newPrison.get(i).get(cell) + newPrison.get(i-1).get(cell));
            }
            newPrison.remove(i - 1);
        }
        
        long max = 1;
        for(ArrayList<Long> arr : newPrison) {
            for(long num : arr) {
                if(num > max) {
                    max = num;
                }
            }
        }
        
        return max;
    }

}