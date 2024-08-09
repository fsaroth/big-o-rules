SCALING_FACTOR = 1000000


def print_first_item_then_first_half_then_say_hi_x_times(items):
    """
        This function prints the first item of the list,
        then the first half of the list, and finally prints "Hi" by
        a factor of SCALING_FACTOR.

        Parameters:
            - items (list): A list of items to be printed.
        Returns:
            - None
        Time Complexity:
            - O(n/2 + 101) => O(n)
        Space Complexity:
            - O(1)
    """
    print(items[0])
    middleIndex = len(items) // 2
    index = 0
    while index < middleIndex:
        print(items[index])
        index += 1
    for i in range(SCALING_FACTOR):
        print("Hi")


if __name__ == "__main__":
    print_first_item_then_first_half_then_say_hi_x_times([1, 2, 3, 4, 5, 6])
