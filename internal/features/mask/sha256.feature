Feature: Mask sha256

   Check that we can perform basic sha256 masking.

   Scenario: call empty
        Given No parameters
        When Invoke "masquerade"
        Then exit code must be 1
        And Error message should contain "Mask fields"
   
   Scenario: no mask
        Given Fields:
        ||
        ||
        When pass thru StdIn msgpack:
        | hello | John Smith |
        And Invoke "masquerade"
        Then exit code must be 0
        And StdOut should be msgpack:
        | hello | John Smith |

   Scenario: mask one field
        Given Fields:
        ||
        |sha256|
        When pass thru StdIn msgpack:
        | hello | John Smith |
        And Invoke "masquerade"
        Then exit code must be 0
        And StdOut should be msgpack:
        | hello | ef61a579c907bbed674c0dbcbcf7f7af8f851538eef7b8e58c5bee0b8cfdac4a |

   Scenario: mask one field starting number
        Given Fields:
        ||
        |sha256|
        When pass thru StdIn msgpack:
        | 1 | John Smith |
        And Invoke "masquerade"
        Then exit code must be 0
        And StdOut should be msgpack:
        | 1 | ef61a579c907bbed674c0dbcbcf7f7af8f851538eef7b8e58c5bee0b8cfdac4a |
