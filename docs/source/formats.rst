Overview
========

``Masquerade`` allows different formats as input / output. Here you can find the complete list.

CSV
===

The simplest format.

There is two executables: :samp:`maskcsvin` and :samp:`maskcsvout`.

CSV -> MsgPack
===================

.. code-block:: console

    > echo hello,World | maskcsvin > binary.out

Will return binary format of the csv.

From MsgPack -> CSV
===================

.. code-block:: console

    > cat binary.out | maskcsvout

Will return "hello","World".

A complete usage may be:

.. code-block:: console

    > echo hello,World | maskcsvin | maskcsvout

Will return "hello","World". Notice that our process add quotes, thats because our binary don't know how looks the original csv, so try to build the most "correct" one.

Custom separator
================

You can use another separator like '|' or '@'. Just provide separator param.

.. code-block:: console

    > echo hello@World | maskcsvin -separator '@' | maskcsvout -separator '|'

Will return "hello"|"World".