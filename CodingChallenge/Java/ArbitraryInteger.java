public class ArbitraryInteger {
    private static final int UPPER_BOUND = 1000000000;

    public int solution(int N) {
        int candidate = (N / 10 + 1) * 10;
        return candidate < UPPER_BOUND ? candidate : 0;
    }
}
