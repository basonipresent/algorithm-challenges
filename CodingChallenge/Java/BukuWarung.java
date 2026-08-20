import java.util.LinkedList;
import java.util.List;

public class BukuWarung {
    public static String largestMagical(String binString) {
        // Write your code here
        List<Integer> list = new LinkedList<>();
        int length = binString.length();
        for (int i = 0; i < length; i++) {
            if (binString.charAt(i) == '1') {
                list.add(1);
            } else {
                list.add(-1);
            }
        }
        for (int i = 0; i < length - 1; i++) {
            WrapperInteger wrapperI = new WrapperInteger(i);
            for (int j = i + 1; j < length - 1; j++) {
                WrapperInteger wrapperJ = new WrapperInteger(j);
                WrapperInteger prefixSum = new WrapperInteger(0);
                WrapperInteger suffixSum = new WrapperInteger(0);
                WrapperInteger end = new WrapperInteger(-1);

                boolean isPrefixMagical = checkMagical(wrapperI, wrapperJ, prefixSum, list);
                boolean isSuffixMagical = checkMagical(wrapperJ, end, suffixSum, list);

                if (isPrefixMagical && isSuffixMagical && suffixSum.getIntValue() > prefixSum.getIntValue()) {
                    StringBuilder stringBuilder = new StringBuilder();
                    String tempStr = stringBuilder.append(binString.substring(0, wrapperI.getIntValue()))
                        .append(binString.substring(wrapperJ.getIntValue(), end.getIntValue()))
                        .append(binString.substring(wrapperI.getIntValue(), wrapperJ.getIntValue() - 1)).toString();
                    if (end.getIntValue() < length) {
                        tempStr = stringBuilder.append(binString.substring(end.getIntValue())).toString();
                    }
                    binString = tempStr;
                }
            }

        }
        return binString;
    }

    private static boolean checkMagical(WrapperInteger start, WrapperInteger end,
            WrapperInteger sum,
            List<Integer> list) {
        if (end.getIntValue() == -1) {
            for (int i = start.getIntValue() + 1; i < list.size(); i++) {
                WrapperInteger tempWrapper = new WrapperInteger(i);
                if (checkMagical(start, tempWrapper, sum, list)) {
                    end.setIntValue(i);
                    return true;
                }
            }
            return false;
        }

        int check = 0;
        int counter = end.getIntValue() - start.getIntValue();
        for (int i = start.getIntValue(); i < end.getIntValue(); i++) {
            check = check + list.get(i);
            if (list.get(i).equals(1)) {
                sum.setIntValue((int) (sum.getIntValue() + Math.pow(2, counter)));
            }
            counter--;
            if (check < 0) {
                return false;
            }
        }
        return check == 0;
    }

    private static class WrapperInteger {
        private int intValue;

        public WrapperInteger(int intValue) {
            this.intValue = intValue;
        }

        public int getIntValue() {
            return intValue;
        }

        public void setIntValue(int intValue) {
            this.intValue = intValue;
        }
    }
}
