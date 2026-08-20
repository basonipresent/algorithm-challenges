public class MaxValue {
    public int func(int N) {
        int insertedDigit = 5;

        int negative = N / Math.abs(N);
        N = Math.abs(N);
        int num = N;

        int maxVal = Integer.MIN_VALUE;
        int counter = 0;
        int position = 1;

        while (num > 0) {
            counter++;
            num = num / 10;
        }

        if (num == 0)
        {
          return insertedDigit * 10;
        }

        for (int i = 0; i <= counter; i++) {
            int newVal = ((N / position) * (position * 10)) + (insertedDigit * position) + (N % position);

            if (newVal * negative > maxVal) {
                maxVal = newVal * negative;
            }

            position = position * 10;
        }

        return maxVal;
    }
}
