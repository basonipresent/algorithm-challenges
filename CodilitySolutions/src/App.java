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
    }
}
