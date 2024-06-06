from profile import run as performance_run

NEMO = ['nemo']

EVERYONE = [
    'dory', 'bruce', 'marlin', 'gill', 'bloat', 'nigel', 'squirt',
    'darla', 'hank', 'nemo'
]  # Worst Case Scenario BIG O Rule 1.

LARGE = ['nemo'] * 100000


def find_nemo(array):
    """
    Searches for the string 'nemo' in the given array.
    Time complexity: O(n) -> Linear Time

    Parameters:
    array (list): The array to search in.

    Returns:
    None
    """
    for i in range(len(array)):
        if array[i] == 'nemo':
            print('Found NEMO!')
            # break
            # Even if you find nemo before the end of the array,
            # the time complexity is still O(n). (Worst Case)


performance_run('find_nemo(LARGE)')
# find_nemo(LARGE)
