package br.com.fsaroth.findingnemo;

import java.util.Arrays;

public class App {

    public static void main(String[] args) {
        long start = System.currentTimeMillis();
        String[] nemo = {"Nemo"};
        String[] everyone = {"Dory", "Bruce", "Marlin", "Nemo", "Gill", "Bloat", "Nigel", "Squirt", "Darla", "Hank"};
        String[] large = Arrays.asList(new String[1000000]).stream().map(e -> "Nemo").toArray(String[]::new);
        findNemo(large);
        long end = System.currentTimeMillis();
        System.out.println("Time taken: " + (end - start) + "ms");
    }

    /**
        * Finds the fish named "Nemo" in the given array of fishes.
        * Complexity: O(n)
        * 
        * @param fishes an array of fishes
        * @return void
        */
    public static void findNemo(String[] fishes) {
        for (int i = 0; i < fishes.length; i++) {
            if (fishes[i].equals("Nemo")) {
                System.out.println("Found Nemo!");
                // break;
                // Even if you find Nemo before the end of the array,
                // the time complexity is still O(n) (Worst Case).
            }
        }
    }
}
