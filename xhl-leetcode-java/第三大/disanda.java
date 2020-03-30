
public class disanda {
    public int thirdMax(int[] nums) {
        Integer maxOne = null;
        Integer maxTwo = null;
        Integer maxThree = null;
        for (Integer i : nums) {
            if (i.equals(maxOne) || i.equals(maxTwo) || i.equals(maxThree)) {
                continue;
            }
            if (maxOne == null || i > maxOne) {
                maxThree = maxTwo;
                maxTwo = maxOne;
                maxOne = i;
            } else if (maxTwo == null || i > maxTwo) {
                maxThree = maxTwo;
                maxTwo = i;
            } else if (maxThree == null || i > maxThree) {
                maxThree = i;
            }
        }
        return maxThree != null ? maxThree : maxOne;
    }
}

