Overview
========

``Masquerade`` allows different locations as sources / output. Here you can find the complete list.

RabbitMQ
========

To read from Rabbit use:

.. code-block:: console

    > maskrabbitin -dial amqp://guest:guest@localhost:5672/ -channel test

This command will consume the queue and output the content into Standard Output.

To write on Rabbit use:

.. code-block:: console

    > cat data | maskrabbitout -dial amqp://guest:guest@localhost:5672/ -channel test

This command will send lines from data file into RabbitMQ.

You can copy a queue using this commands together:

.. code-block:: console

    > maskrabbitin -dial amqp://guest:guest@localhost:5672/ -channel topicA | maskrabbitout -dial amqp://guest:guest@localhost:5672/ -channel topicB

HDFS
====

HDFS has stdio support, just use is as follows.

To read:

.. code-block:: console

    > hdfs dfs -cat data.csv

To write:

.. code-block:: console

    > cat data.csv | hdfs dfs -put - data.csv

S3
==

This services can be accesed with stdio support thru `minio-cli <https://github.com/minio/mc#add-a-cloud-storage-service>`_

GCS (Google Cloud Storage)
==========================

This services can be accesed with stdio support thru `minio-cli <https://github.com/minio/mc#add-a-cloud-storage-service>`_

