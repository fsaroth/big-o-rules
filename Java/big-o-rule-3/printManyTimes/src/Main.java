import org.jetbrains.annotations.NotNull;

public class Main {

    private static final int SCALING_FACTOR = 1000000;

    public static void main(String[] args) {

        printFirstItemThenFirstHalfThenSayHiXTimes(new Integer[]{1, 2, 3, 4, 5, 6, 7, 8, 9, 10});
    }
    /**
     * Prints the first item of an array, then the first half of the array, and finally prints "hi" a number of times
     * defined by a scaling factor.
     * This method performs three distinct operations:<br>
     * 1. Prints the first element of the provided integer array to the console.<br>
     * 2. Iterates through the first half of the array (up to the middle index, exclusive) and prints each element
     * to the console.<br>
     * 3. Prints the string "hi" to the console a number of times equal to the predefined SCALING_FACTOR constant.<br>
     * 4. Time complexity: O(1 + n/2 + 2) = O(n)<br>
     * 5. Space complexity: O(1)<br>
     * 6. Returns: void<br>
     * @param items The array of integers to be processed.
     *              Assumes the array is not null and has at least one element.<br>
 */
    private static void printFirstItemThenFirstHalfThenSayHiXTimes(@NotNull Integer [] items) {
        System.out.println(items[0]);

        int middleIndex = items.length / 2;
        int index = 0;
        while (index < middleIndex) {
            System.out.println(items[index]);
            index++;
        }

        for (int i = 0; i < SCALING_FACTOR; i++) {
            System.out.println("hi");
        }
    }
}