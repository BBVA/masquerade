Feature: RabbitMQ masquerade connector
    Source and Destination from RabbitMQ to mask data and scale out.
    For testing proposes source will be interrupted 1 second after start,
    otherwise test will never finish. So exit code will be -1.

    Scenario: No config on source
        Given No parameters
        When Invoke "maskrabbitin"
        Then exit code must be 1
        And Error message should contain "Rabbit Dial expected"

    Scenario: Valid Source config but no channel
        Given Dial parameter "amqp://guest:guest@localhost:5672/"
        And No Channel
        When Invoke "maskrabbitin"
        Then exit code must be 1
        And Error message should contain "channel expected"
    
    @dev
    Scenario: Valid Source config read lines
        Given Dial parameter "amqp://guest:guest@localhost:5672/"
        And Channel "test" with lines:
        """
        hello,Tom
        hi,John
        """
        When Invoke "maskrabbitin"
        Then exit code must be 0
        And StdOut should contain lines:
        """
        hello,Tom
        hi,John
        """

    Scenario: No config on destination
        Given No parameters
        When Invoke "maskrabbitout"
        Then exit code must be 1
        And Error message should contain "Rabbit Dial expected"
    
    Scenario: Valid destination but no channel
        Given Dial parameter "amqp://guest:guest@localhost:5672/"
        And No Channel
        When Invoke "maskrabbitout"
        Then exit code must be 1
        And Error message should contain "channel expected"
    
    Scenario: Valid destination config write lines
        Given Dial parameter "amqp://guest:guest@localhost:5672/"
        And Channel "test"
        When pass thru StdIn lines:
        """
        hi,John
        hello,Tom
        """
        When Invoke "maskrabbitout"
        Then Channel "test" contains:
        """
        hi,John
        hello,Tom
        """
        And exit code must be 0

