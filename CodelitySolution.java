import java.util.*;
import java.util.stream.Collectors;

public class CodelitySolution {
    public static void main(String[] args) {
        int[] A = new int[] { 1, 3, 6, 4, 1, 2 };
        int[] B = new int[] { 10, 0, 8, 2, -1, 12, 11, 3 };
        String S = "111";
        //System.out.println(smallerPositiveArray(A));
        //System.out.println(byciclePark(B));
        //System.out.println(nonNegative(S));
        int[] param = new int[] { 3, 6, 9 };
        System.out.println(solution(param));
        param = new int[] { 3, 7, 2, 5, 5 };
        System.out.println(solution(param));
    }

    public static int smallerPositiveArray(int[] A) {
        int N = A.length;
        Set<Integer> integerSet = new HashSet<>();
        for (int i : A) {
            if (i > 0) {
                integerSet.add(i);
            }
        }
        for (int i = 1; i <= N + 1; i++) {
            if (!integerSet.contains(i)) {
                return i;
            }
        }
        return 1;
    }

    public static int byciclePark(int[] A) {
        Arrays.sort(A);
        int ans = Integer.MIN_VALUE;
        if (A.length == 2)
            return (A[1] - A[0]) / 2;
        for (int i = 0; i < A.length - 1; i++) {
            // If the free position between two points
            if (A[i + 1] - A[i] > 1) {
                ans = Math.max(ans, A[i + 1] - A[i]);
            }
        }

        return ans / 2;
    }

    public static int nonNegative(String S) {
        int counter = 0;
        long V = Long.parseUnsignedLong(S, 2);
        while (V > 0) {
            if (V % 2 == 0) {
                V = V / 2;
            } else {
                V = V - 1;
            }
            counter++;
        }
        return counter;
    }

    public int solution(int A, int B) {
        // write your code in Java SE 8
        String binary = "";
        for (long i = (1L << B - 1); i > 0; i = i / 2) {
            binary += (A & i) != 0 ? "1" : "0";
        }
        return Integer.parseInt(binary, 2);
    }

    public static String solution(int[] A) {
        String finalOP = "impossible";
        String[] colored = new String[A.length];
        int k = 3;
        List<Integer> numbers = Arrays.stream(A).boxed().collect(Collectors.toList());
        Integer arrSum = numbers.stream().reduce(0, ((a, b) -> a + b));
        
        System.out.println(arrSum % k == 0 && A.length >= k);
        if (arrSum % k == 0 && A.length >= k) {
            int requiredSumOfSubset = arrSum / k;
            colored = fillColor(colored, A, requiredSumOfSubset, 0, "R");
            colored = fillColor(colored, A, requiredSumOfSubset, 0, "G");
            colored = fillColor(colored, A, requiredSumOfSubset, 0, "B");
            if (colored.length > 0) {
                StringBuilder sb = new StringBuilder();
                for (String str : colored)
                    sb.append(str);
                finalOP = sb.toString();
            }
        }
        return finalOP;
    }

    private static String[] fillColor(String[] colored, int[] arr, int requiredSum, int currVal, String color) {
        boolean added = false;
        for (int idx = 0; idx < colored.length; idx++) {
            if (colored[idx] == null) {
                int tempSum = currVal + arr[idx];
                if (tempSum == requiredSum) {
                    added = true;
                    colored[idx] = color;
                    break;
                } else if (tempSum < requiredSum) {
                    String[] newColored = colored;
                    newColored[idx] = color;
                    String[] retColored = fillColor(newColored, arr, requiredSum, tempSum, color);
                    if (retColored.length > 0) {
                        added = true;
                        colored = retColored;
                        break;
                    }
                }
            }
        }

        String[] arrNew = new String[0];
        return added ? colored : arrNew;
    }
}
