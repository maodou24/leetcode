import java.util.HashMap;

class Solution {
    public int[] twoSum(int[] nums, int target) {
        HashMap m = new HashMap(nums.length);
        for (int i = 0; i < nums.length; i++) {
            int tmp = target - nums[i];
            if (m.containsKey(tmp)) {
                return new int[]{(Integer)m.get(tmp), i};
            }
            m.put(nums[i], i);
        }
        return new int[]{};
    }
}