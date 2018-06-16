namespace Sum;

%{
    #include "../../libsum.h"
}%

class Sum
{
    public static function getSum(int a, int b)
    {
        int res = 0;

        %{
            res = Sum(a, b);
        }%

        return res;
    }
}
