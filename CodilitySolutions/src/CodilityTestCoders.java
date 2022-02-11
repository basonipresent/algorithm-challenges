package CodilitySolutions.src;

public class CodilityTestCoders {
    public void solution(int N) {
        // write your code in Java SE 8
        int i = N;
        String result2 = "Codility";
        String result3 = "Test";
        String result5 = "Coders";

        for (int n = 1; n <= i; n++) {

            if (n % 2 == 0 && n % 3 == 0 && n % 5 == 0) {
                System.out.println(result2 + result3 + result5);
            } else if (n % 3 == 0 && n % 5 == 0) {
                System.out.println(result3 + result5);
            } else if (n % 2 == 0 && n % 3 == 0) {
                System.out.println(result2 + result3);
            } else if (n % 2 == 0 && n % 5 == 0) {
                System.out.println(result2 + result5);
            } else if (n % 2 == 0) {
                System.out.println(result2);
            } else if (n % 3 == 0) {
                System.out.println(result3);
            } else if (n % 5 == 0) {
                System.out.println(result5);
            } else {
                System.out.println(n);
            }
        }
    }
}
