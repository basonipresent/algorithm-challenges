import java.util.Arrays;

public class App {
    public static void main(String[] args) throws Exception {
        System.out.println("Hello, World!");

        Chart chart = new Chart();
        int[] input = new int[5];
        System.out.println(chart.solution(input));

        MaxValue maxValue = new MaxValue();
        System.out.println(maxValue.func(10));

        TricoloringPossible tricoloringPossible = new TricoloringPossible();
        int[] param = new int[3];
        System.out.println(tricoloringPossible.solution(param));

        ArbitraryInteger arbitraryInteger = new ArbitraryInteger();
        System.out.println(arbitraryInteger.solution(23));

        CodilityTestCoders codilityTestCoders = new CodilityTestCoders();
        codilityTestCoders.solution(15);

        Incorrect incorrect = new Incorrect();
        System.out.println(incorrect.solution("codility"));

        System.out.println(CodelitySolution.smallerPositiveArray(new int[] { 1, 3, 6, 4, 1, 2 }));
        System.out.println(CodelitySolution.byciclePark(new int[] { 10, 0, 8, 2, -1, 12, 11, 3 }));
        System.out.println(CodelitySolution.nonNegative("111"));
        System.out.println(new CodelitySolution().solution(5, 4));
        System.out.println(CodelitySolution.solution(new int[] { 3, 7, 2, 5, 5 }));

        System.out.println(CollisonCourse.collision(Arrays.asList(2, 4, 1, 1), 0));

        System.out.println(FindFirstRepeated.findFirstRepeated("the quick brown fox jumps over the lazy dog"));

        System.out.println(MaxInversionsResult.maxInversions(Arrays.asList(1, 1, 1, 2, 2)));

        System.out.println(PerfectSubstringResult.perfectSubstring("212345", 1));

        System.out.println(PrisonBreakResult.prison(4, 4, Arrays.asList(2), Arrays.asList(2)));

        System.out.println(SwapPalindrome.isPalindromePossible("bdababd"));

        System.out.println(TwoPower.isPowerOfTwoBinary("0010"));

        System.out.println(BukuWarung.largestMagical("1001"));

        TreeNode root = new TreeNode(3);
        root.left = new TreeNode(5);
        root.right = new TreeNode(1);
        root.left.left = new TreeNode(6);
        root.left.right = new TreeNode(2);
        root.right.left = new TreeNode(0);
        root.right.right = new TreeNode(8);
        root.left.right.left = new TreeNode(7);
        root.left.right.right = new TreeNode(4);
        System.out.println(Solution.distanceK(root, root.left, 2));
    }
}
