import java.util.ArrayList;
import java.util.List;

public class Main {
  class Solution {
    public static List<Integer> spiralOrder(int[][] matrix) {
      List<Integer> result = new ArrayList<>();
      if (matrix.length == 0) {
        return result;
      }

      int m = matrix[0].length - 1;
      int n = matrix.length - 1;
      int top = 0;
      int left = 0;
      int right = m;
      int bottom = n;

      while(top <= bottom && left <= right) {
        // go from left to right
        for (int i = left; i <= right; i++) {
          result.add(matrix[top][i]);
        }
        top++;

        // go from top right to bottom
        for (int i = top; i <= bottom; i++) {
          result.add(matrix[i][right]);
        }
        right--;

        // go from bottom right to left
        for (int i = right; i >= left; i--) {
          result.add(matrix[bottom][i]);
        }
        bottom--;

        // go from bottom left to top
        for (int i = bottom; i >= top; i--) {
          result.add(matrix[i][left]);
        }
        left++;
      }

      return result;
    }
  }

  public static void main(String[] args) {
    int[][] matrix = {
            {1,2,3},
            {4,5,6},
            {7,8,9}
    };
    System.out.println(Solution.spiralOrder(matrix));
  }
}