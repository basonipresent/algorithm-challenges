package CodilitySolutions.src;

import java.util.Random;

public class ArbitraryInteger {
    public int solution(int N) {
        // write your code in Java SE 8
        boolean done = false;
        int arbitrary = 0;
        while(!done) {
            arbitrary = arbitrary();
            if(arbitrary > N && arbitrary < 1000000000 && arbitrary % 10 == 0) {
                done = true;
                return arbitrary;
            }
        }
        return 0;
    }

    private int arbitrary(){
        Random arbitrary = new Random();
        int i  = arbitrary.nextInt();
        return i;
    }
}
