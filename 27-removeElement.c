#include <stdio.h>

int removeElement(int* nums, int numsSize, int val);

int removeElement(int* nums, int numsSize, int val){
    int y = 0;

    for (int i = 0; i < numsSize; ++i) {
        if (nums[i] != val) {
            nums[y++] = nums[i];
        }
    }
    return y;
}

int main(int argc, char *argv[])
{
    int arr[4] = {3,2,2,3};
    int res = removeElement(arr, 4, 3);

    printf("res: %d\n", res);
    return 0;
}

