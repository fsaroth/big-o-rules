package br.com.fsaroth.firsttwoboxes;

public class App {
    public static void main(String[] args) {
        int[] boxes = new int[] { 1, 2, 3, 4, 5};
        logFirstTwoBoxes(boxes);
    }


    /**
     * Prints the first two elements of the given array of boxes.
     * Time complexity: O(1)
     *
     * @param boxes the array of boxes
     */
    public static void logFirstTwoBoxes(int[] boxes) {
        System.out.println(boxes[0]); // O(1)
        System.out.println(boxes[1]); // O(1)
    }
}
