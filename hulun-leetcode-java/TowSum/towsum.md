### 题目描述
> Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
#### 例子
>Given nums = [2, 7, 11, 15], target = 9,
Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

### 第一次提交
暴力解法 时间复杂度 O(n^2)
```java
class Solution {
        public int[] twoSum(int[] nums, int target) {
            int[] a = new int[2];
            for(int i = 0 ; i < nums.length - 1; i++){
                for(int j = i+1; j < nums.length ; j++){
                    if(nums[i] + nums[j] == target){
                        a[0] = i;
                        a[1] = j;
                        break;
                    }
                }
            }
            return a;
        }
    }
```
### 第二次提交
利用哈希表 时间复杂度 O(n)
```java
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
```
