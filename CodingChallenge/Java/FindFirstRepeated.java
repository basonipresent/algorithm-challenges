import java.util.HashSet;
import java.util.Set;

public class FindFirstRepeated {
  static String findFirstRepeated(String sentence) {
    String splittedSentence[] = sentence.split("[-,:;.\\s]");
    Set<String> seenWords = new HashSet<>();

    for (String word : splittedSentence) {
        if (!seenWords.add(word)) {
            return word;
        }
    }
    return "";
}
}
