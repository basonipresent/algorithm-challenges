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


//Array to store position of Horizontal Cuts as Flags
int x[] = new int[n];
for(int i=0;i<x.length;i++)
    x[i] = 1;


//Array to store position of Vertical Cuts  as Flags
int y[] = new int[m];
for(int i=0;i<y.length;i++)
    y[i] = 1;


//Start storing the flags in the Positions
int cx=0, ox= Integer.MIN_VALUE, cy=0, oy=Integer.MIN_VALUE;
for (int i=0;i<h.length;++i)
    x[h[i]]= 0;//x has 1 where there is a bar, 0 when the bar is removed
for (int i=0;i<v.length;++i)
    y[v[i]]= 0;//y has 1 where there is a bar, 0 when the bar is removed


//the below 2 for loops can be replaced by DFS OR BFS Alogorithm
for (int i=0;i<=n-1;i++) {//loop to find the maximum gap horizontally
    if (x[i]>0)
        cx= 0;
    else {
        cx++;
        ox= Math.max(ox, cx);
        System.out.println("Horizontal max is...."+ox);
    }
}

for (int i=1;i<=m-1;i++) {//loop to find the maximum gap horizontally
    if (y[i]>0)
        cy= 0;
    else {
        cy++;
        oy= Math.max(oy, cy);
        System.out.println("vertical max is...."+ox);
    }
}

return (ox+1)*(oy+1);