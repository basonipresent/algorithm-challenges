public class find_first_repeated {
  static String findFirstRepeated(String sentence) {
    String splittedSentence[] = sentence.split("[-,:;.\\s]");
    HashMap<String, Integer> setOfWords = new HashMap<String, Integer>();
    String firstRepeatWord = "";

    for (int i = 0; i < splittedSentence.length; i++) {
        if (setOfWords.containsKey(splittedSentence[i])) {
            setOfWords.put(splittedSentence[i], setOfWords.get(splittedSentence[i]) + 1);
            firstRepeatWord = splittedSentence[i];
            break;
        } else {
            setOfWords.put(splittedSentence[i], 1);
        }
    }
    return firstRepeatWord;
}
}
