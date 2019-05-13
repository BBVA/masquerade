Overview
========

``Masquerade`` has the power to read from many different locations as source, and export the obfuscated data to a different location **in streaming** and with a very high performance.

.. image:: /_static/images/masquerade_working_mode.png


**Examples:**

1. You can read data from a **S3 Bucket**, obfuscate data, and export results to **local file system**.
2. You can read data from a **S3 Bucket**, obfuscate data, and export results to **HDFS File System**.
3. You can read data from a **HDFS Filesystem**, obfuscate data, and export results to **Google Cloud Storage**.
4. You can read data from a **S3 bucket**, obfuscate data, and export results to another **S3 Bucket**.