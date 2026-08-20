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



class MaxInversionsResult {

    /*
     * Complete the 'maxInversions' function below.
     *
     * The function is expected to return a LONG_INTEGER.
     * The function accepts INTEGER_ARRAY arr as parameter.
     */

    public static long maxInversions(List<Integer> arr) {
        int n = arr.size();

        // Coordinate-compress values to dense ranks 1..n so they fit a Fenwick tree.
        Integer[] sorted = arr.toArray(new Integer[0]);
        Arrays.sort(sorted);
        Map<Integer, Integer> rank = new HashMap<>();
        int nextRank = 1;
        for (int v : sorted) {
            if (!rank.containsKey(v)) {
                rank.put(v, nextRank++);
            }
        }

        // suffixSmaller[i] = number of elements after i with a smaller value.
        int[] suffixSmaller = new int[n];
        FenwickTree fromRight = new FenwickTree(n);
        for (int i = n - 1; i >= 0; i--) {
            int r = rank.get(arr.get(i));
            suffixSmaller[i] = fromRight.prefixCount(r - 1);
            fromRight.add(r);
        }

        // For each i, count elements before i with a greater value, times elements
        // after i with a smaller value, and sum: same result as the original O(n^2)
        // scan, computed in O(n log n) via two Fenwick trees.
        long inversions = 0;
        FenwickTree fromLeft = new FenwickTree(n);
        for (int i = 0; i < n; i++) {
            int r = rank.get(arr.get(i));
            int prefixGreater = i - fromLeft.prefixCount(r);
            inversions += (long) prefixGreater * suffixSmaller[i];
            fromLeft.add(r);
        }

        return inversions;
    }

    private static class FenwickTree {
        private final int[] tree;

        FenwickTree(int size) {
            tree = new int[size + 1];
        }

        void add(int i) {
            for (; i < tree.length; i += i & (-i)) {
                tree[i]++;
            }
        }

        int prefixCount(int i) {
            int sum = 0;
            for (; i > 0; i -= i & (-i)) {
                sum += tree[i];
            }
            return sum;
        }
    }

}

public class MaxInversions {
    public static void main(String[] args) throws IOException {
        BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bufferedWriter = new BufferedWriter(new FileWriter(System.getenv("OUTPUT_PATH")));

        int arrCount = Integer.parseInt(bufferedReader.readLine().trim());

        List<Integer> arr = new ArrayList<>();
        for (int i = 0; i < arrCount; i++) {
            arr.add(Integer.parseInt(bufferedReader.readLine().trim()));
        }

        long result = MaxInversionsResult.maxInversions(arr);

        bufferedWriter.write(String.valueOf(result));
        bufferedWriter.newLine();

        bufferedReader.close();
        bufferedWriter.close();
    }
}