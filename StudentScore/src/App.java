import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class App {
    public static void main(String[] args) throws Exception {
        System.out.println("Hello, World!");
        menu();
    }

    public static void menu() throws Exception {
        System.out.println("Menu :: \n");
        System.out.println("1.\n");
        System.out.println("2.\n");
        System.out.println("3.\n");
        System.out.println("4.\n");
        BufferedReader reader = new BufferedReader(
            new InputStreamReader(System.in));
        String choice = reader.readLine();
        switch (choice){
            case "1" :
            System.out.println(1);
            break;
            case "2" :
            System.out.println(2);
            break;
            case "3" :
            System.out.println(3);
            break;
            case "4" :
            System.out.println(4);
            break;
            default :
            System.out.println("default");
            break;
        }
        menu();
    }
}
