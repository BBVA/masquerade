Feature: Masquerade csv formater
    Source and Destination from csv files to translate from to csv.

    Scenario: simple csv to binary
        Given No parameters
        When pass thru StdIn lines:
        """
        hi,tom
        """
        And Invoke "maskcsvin"
        Then exit code must be 0
        And StdOut should be msgpack:
        | hi     | tom     |

    Scenario: simple binary to csv
        Given No parameters
        When pass thru StdIn msgpack:
        | hi     | tom     |
        And Invoke "maskcsvout"
        Then exit code must be 0
        And StdOut should contain lines:
        """
        "hi","tom"
        """
    
    Scenario: multiline csv to binary
        Given No parameters
        When pass thru StdIn lines:
        """
        hi,tom
        hello,tim
        ciao,giovanni
        """
        And Invoke "maskcsvin"
        Then exit code must be 0
        And StdOut should be msgpack:
        | hi     | tom      |
        | hello  | tim      |
        | ciao   | giovanni |

    Scenario: simple binary to csv
        Given No parameters
        When pass thru StdIn msgpack:
        | hi     | tom     |
        | hello  | tim      |
        | ciao   | giovanni |
        And Invoke "maskcsvout"
        Then exit code must be 0
        And StdOut should contain lines:
        """
        "hi","tom"
        "hello","tim"
        "ciao","giovanni"
        """

    Scenario: csv to binary with separator
        Given separator "|"
        When pass thru StdIn lines:
        """
        hi|tom
        """
        And Invoke "maskcsvin"
        Then exit code must be 0
        And StdOut should be msgpack:
        | hi     | tom     |

    Scenario: binary to csv with separator
        Need pass order of fields allways to test, otherwise order 
        is random. Also, by the moment, csv output it's allways quoted.
        Given separator "|"
        When pass thru StdIn msgpack:
        | hi     | tom     |
        And Invoke "maskcsvout"
        Then exit code must be 0
        And StdOut should contain lines:
        """
        "hi"|"tom"
        """