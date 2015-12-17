import skilstak.colors as c

def print_plain(message):
    """prints <message> in plain text>"""
    print(message)

def print_color(message):
    """Prints hello in a random color each time"""
    print(c.rc() + message + "!" + c.reset)

def print_nyan(message):
    """Prints <message> forever"""
    while True:
        print(c.rc() + message , end=" ")
