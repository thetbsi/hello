"""Not yo grandma's hello world module.

This module is a collection of functions that can be called
from programs or other modules that import it.

Notice there is no shebang line because there is no ``if __name__
== '__main__'`` section in this module, although it could have one
for testing, etc.

This documentation can be displayed with ``pydoc3 hello``
or converted to web page with any number of tools.

"""

import skilstak.colors as c
import time

def print_plain(message):
    """ Prints <messeage> to screen.

    Args:
        message (str): message to print

    """
    print(message)

def print_stripes(message):
    """ Prints <message> to screen forever as stripes

    Args:
        message (str: message to print

    """
    while True:
        print(message,end="                   ")

def print_multi(message):
    """ Prints <message> in multiple colors like Christmas tree.

    Args:
        message (str): message to print

    """
    while True:
        print(c.clear + c.multi(message))
        time.sleep(0.5)

def print_merica(message):
    """Fills the screen with Merica colors <message>.

    Args:
        message (str): message to print

    """
    print(c.clear)
    while True:
        print(c.merica(message), end=" ")

def parse_args():
    """Parse and return properties (name, option)

    Returns:
        dict: A dictionary of properties (name, oprion)

    """

    from sys import argv

    name = "world"
    option = ""

    if len(argv) == 2:
        if argv[1].startswith("-"):
            option = argv[1]
        else:
            name = argv[1]
    elif len(argv > 2:
        option = argv[1]
        name = argv[2]

    return name, option
