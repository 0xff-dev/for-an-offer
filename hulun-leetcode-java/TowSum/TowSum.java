import java.util.HashMap;
import java.util.Map;

public class TowSum {
    public static void main(String[] args){
        int[] result;
        int[] a = new int[]{2,7,11,15};
        Solution solution = new Solution();
        result = solution.twoSum(a,9);
        System.out.print("[");
        for(int i=0;i<result.length;i++){
            System.out.print(result[i]);
            if(i==1){
                break;
            }
            System.out.print(",");
        }
        System.out.print("]");
    }
}
class Solution {
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> map = new HashMap<>();
        int[] result = {-1, -1};
        for (int i = 0; i < nums.length; ++i) {
            if (map.containsKey(target - nums[i])) {
                result[0] = map.get(target - nums[i]);
                result[1] = i;
                break;
            }
            map.put(nums[i], i);
        }
        return result;
    }
}