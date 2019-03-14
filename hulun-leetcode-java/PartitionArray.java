public class PartitionArray {
    private void swap(int i, int j, int[] arr) {
        int tmp = arr[i];
        arr[i] = arr[j];
        arr[j] = tmp;
    }
    /**
     *@param nums: The integer array you should partition
     *@param k: As description
     *return: The index after partition
     */
    public int partitionArray(int[] nums, int k) {

        int pl = 0;
        int pr = nums.length - 1;
        while (pl <= pr) {
            while (pl <= pr && nums[pl] < k) {
                pl++;
            }
            while (pl <= pr && nums[pr] >= k) {
                pr--;
            }
            if (pl <= pr) {
                swap(pl, pr, nums);
                pl++;
                pr--;
            }
        }
        return pl;
    }
}

