from profile import run as performance_run

BOXES = [1, 2, 3, 4, 5]


def log_first_two_boxes(boxes):
    """
    Prints the first two boxes in the given list.
    Time complexity: O(1) -> Constant Time

    Parameters:
    boxes (list): The list of boxes.

    Returns:
    None
    """
    print(boxes[0])  # O(1)
    print(boxes[1])  # O(1)


performance_run('log_first_two_boxes(BOXES)')
