# Garder server

Receive garden information from arduino sensors

# API

## Parameters 

### List All Parameters [GET] [/garden/v1/parameters/list]


+ Response 200 (application/json)

        [
            {
                "_id": "5931dbb333121abffc1d88cd",
                "createdat": "2017-06-02 18:42:11.309830755 -0300 -03",
                "measure": " percent",
                "name": " humidity",
                "value": 14
            },
            {
                "_id": "5931e3c733121abffc1d8a6e",
                "createdat": "2017-06-02 19:16:39.554790641 -0300 -03",
                "measure": " percent",
                "name": " humidity",
                "value": 11
            }
        ]
        

### Create a New Parameter [POST] [/garden/v1/parameters/save]

You may create your own parameter using this action. It takes a JSON
object containing a parameters values.

+ Request (application/json)

        {
            "name": "humidity",
            "value": 80,
            "measure": "percent",
        }

+ Response 201 (application/json)

    + Headers

            Location: /parameter/UUID

    + Body

            {
                "created": "2015-08-05T08:40:51.620Z"
            }
