class Solution {
public:
    void Move(int num1, int num2, int &step){
        if(num1 == 0 || num2 ==0){
            return ;
        }
        if(num1 >= num2){
            num1 = num1 - num2;
        }else{
            num2 = num2 - num1;
        }
        step++;
        Move(num1, num2, step);
    }
    int countOperations(int num1, int num2) {
        int step = 0;
        Move(num1, num2, step);
        return step;
    }
};