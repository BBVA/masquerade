Installation
============

To install ``Masquerade`` you need `GoLang <https://golang.org>`_ compilation tools.

.. _installation:

Just make, it will put their binaries into $GOPATH/bin.

.. code-block:: console

    > make

Ensure that $GOPATH/bin it's in your path:

.. code-block:: console

    > export PATH=$(go env GOPATH)/bin:$(go env GOROOT)/bin:$PATH

