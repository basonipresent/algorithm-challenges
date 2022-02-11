import java.util.Arrays;
import java.util.Set;
import java.util.HashSet;

public class CodelitySolution {
    public static void main(String[] args) {
        int[] A = new int[]{1, 3, 6, 4, 1, 2};
        int[] B = new int[]{10, 0, 8, 2, -1, 12, 11, 3};
        String S = "111";
        System.out.println(smallerPositiveArray(A));
        System.out.println(byciclePark(B));
        System.out.println(nonNegative(S));
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
        if (A.length == 2) return (A[1] - A[0]) / 2;
        for (int i = 0; i < A.length - 1; i++) {
            //If the free position between two points
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
}
