/**
 * @param {number[]} nums
 * @return {number}
 */
var findPeakElement = function (nums) {
    //Return index of peak
    let left = 0;
    let right = nums.length - 1;
    while(left<right){
        const mid = Math.floor((left+right)/2);
        if(nums[mid] < nums[mid+1]){
            left = mid + 1; //must be on right side
        }else{
            right = mid; //peak must be on left side
        }
    }
    return left;
};
console.log(findPeakElement([1,2,3,1]));
console.log(findPeakElement([1,2,1,3,5,6,4]));
console.log(findPeakElement([1,7,6,6,5,4]));
